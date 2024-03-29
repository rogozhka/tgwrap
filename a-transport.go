package tgwrap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// GenericResponse is a common part of all the API responses.
// Description and ErrorCode are provided in case of error.
type GenericResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

type commonRequestOptions struct {
	// Context is optional request context where applicable.
	// context.Background() is used by default if nil.
	Context context.Context
}

type fCommandSender func(ctx context.Context, methodName string, bodyStruct interface{}) ([]byte, error)

// Generic fCommandSender method for any command.
// Encodes bodyStruct w/ nested structures as JSON
func (p *bot) sendJSON(ctx context.Context, methodName string, bodyStruct interface{}) ([]byte, error) {

	if nil == ctx {
		ctx = context.Background()
	}

	// empty result to return with errors
	var res []byte

	url := fmt.Sprintf("%s%s/%s", p.apiURL, p.token, methodName)

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(bodyStruct); err != nil {
		return res, err
	}

	return p.postRequest(ctx, url, "application/json", &buf)
}

// postRequest makes request and reads result.
func (p *bot) postRequest(ctx context.Context, url string, contentType string, body io.Reader) ([]byte, error) {

	var res []byte
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req = req.WithContext(ctx)

	resp, err := p.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
		resp.Close = true
	}
	if err != nil {
		return res, p.maskToken(err)
	}
	res, err = io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read resp.Body | %w", err)
	}
	return res, nil
}

// sendFormData is fCommandSender method for file/media upload commands.
// Encodes body as multipart/form-data
//
// NOTE:
// 1) recommended for use w/ upload commands only,
// although should work with other API methods w/o
// nested structures
//
//  2. it's unknown how to encode request with
//     nested structures like ReplyMarkup so they are ignored.
func (p *bot) sendFormData(ctx context.Context, methodName string, bodyStruct interface{}) ([]byte, error) {

	// empty result to return with errors
	var res []byte
	url := fmt.Sprintf("%s%s/%s", p.apiURL, p.token, methodName)

	buf := bytes.Buffer{}
	mpw := multipart.NewWriter(&buf)

	t := reflect.TypeOf(bodyStruct)
	r := reflect.ValueOf(bodyStruct)

	// iterate over bodyStruct fields including contents
	// of the embedded struct
	for _, fieldT := range thestruct.Fields(t, thestruct.FieldsFilterDefault) {
		v := r.FieldByName(fieldT.Name)
		// ignore fields w/o struct tags
		if len(fieldT.Tag) < 1 {
			continue
		}
		stags, err := thestruct.ParseStructTagLiteral(string(fieldT.Tag))
		if err != nil {
			return res, err
		}
		jsonTag := stags.Tag("json")
		formTag := stags.Tag("form")

		// do not encode fields w/o "json" struct tag
		// it's form encoding but we use json tags
		// as universal encoding markup
		if jsonTag == nil {
			continue
		}
		if jsonTag.IsOpt("omitempty") && thestruct.IsEmptyValue(v) {
			continue
		}

		typeName := thestruct.Type(v.Type()).Name()
		if len(typeName) == 0 && formTag != nil && formTag.Value == "file" {

			path := fmt.Sprint(v.Interface())
			if len(path) < 1 {
				// file member has all the tags
				// but initialized as file_id || url
				// and should be treated as MarshalText()
				goto writeField
			}

			f, err := os.Open(path)
			if err != nil {
				return res, fmt.Errorf("cannot open | %v | %w", path, err)
			}
			if f != nil {
				// defer is not recommended in loop
				// but there can be not so many files in request struct ;)
				defer f.Close()
			}
			part, err := mpw.CreateFormFile(jsonTag.Value, filepath.Base(path))
			if err != nil {
				return res, err
			}
			if _, err = io.Copy(part, f); err != nil {
				return res, err
			}
			continue
		}
	writeField:
		mpw.WriteField(jsonTag.Value, fmt.Sprintf("%v", v.Interface()))
	}

	// write closing boundary into buf
	mpw.Close()

	return p.postRequest(ctx, url, mpw.FormDataContentType(), &buf)
}

// getAPIResponse makes request and decodes API result into GenericResponse-based object.
// returns error if API result decoded and not OK.
func (p *bot) getAPIResponse(
	ctx context.Context,
	methodName string,
	sender fCommandSender,
	bodyStruct interface{},
	resultStruct interface{},
) error {
	if nil == ctx {
		ctx = context.Background()
	}
	data, err := sender(ctx, methodName, bodyStruct)
	if err != nil {
		return fmt.Errorf("%s | %w", methodName, err)
	}
	if err := json.Unmarshal(data, resultStruct); err != nil {
		return fmt.Errorf("unmarshal | %w | %s", err, data)
	}

	// name of embedded struct with common fields (OK, ErrorCode, ...)
	const commonStruct = "GenericResponse"

	r := reflect.ValueOf(resultStruct).Elem()
	f := r.FieldByName(commonStruct)
	if !f.IsValid() {
		return fmt.Errorf("%v not found in target struct | %T", commonStruct, resultStruct)
	}
	// see "type assertion" topic ;)
	resp, ok := f.Interface().(GenericResponse)
	if !ok {
		// field is not embedded but has (commonStruct) name
		return fmt.Errorf("unmarshal target field is not %v | %T", commonStruct, resultStruct)
	}
	if !resp.OK {
		return fmt.Errorf("API | %v | %v | %v", methodName, resp.ErrorCode, resp.Description)
	}
	return nil
}
