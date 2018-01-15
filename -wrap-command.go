package tgwrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/rogozhka/tgwrap/internal/thestruct"
)

type GenericResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

type fCommandSender func(methodName string, bodyStruct interface{}) ([]byte, error)

//
// Generic method wrapper for any command
// encodes body as JSON
//
func (p *bot) sendJSON(methodName string, bodyStruct interface{}) ([]byte, error) {

	// empty result to return with errors
	res := []byte{}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", p.token, methodName)

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(bodyStruct); err != nil {
		return res, err
	}

	return p.postRequest(url, "application/json", &buf)
}

//
// postRequest makes request and reads result
//
func (p *bot) postRequest(url string, contentType string, body io.Reader) ([]byte, error) {

	var res []byte

	resp, err := http.Post(url, contentType, body)
	if resp != nil {
		defer resp.Body.Close()
		resp.Close = true
	}
	if err != nil {
		return res, fmt.Errorf("POST error:%v", err)
	}

	if res, err := ioutil.ReadAll(resp.Body); err != nil {
		return res, fmt.Errorf("POST ReadAll error:%v", err)
	} else {
		return res, nil
	}
}

//
// Wrapper for file upload commands
// encodes body as multipart/form-data
//
// NOTE:
// 1) recommended for use with file upload commands only;
//
// 2) it's unknown how to encode request with
//    nested structures like ReplyMarkup - they are ignored.
//
func (p *bot) sendFormData(methodName string, bodyStruct interface{}) ([]byte, error) {

	// empty result to return with errors
	res := []byte{}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", p.token, methodName)

	var buf bytes.Buffer
	mpw := multipart.NewWriter(&buf)

	t := reflect.TypeOf(bodyStruct)
	r := reflect.ValueOf(bodyStruct)

	arr := thestruct.Fields(t)

	for _, fieldT := range arr {

		v := r.FieldByName(fieldT.Name)

		if len(fieldT.Tag) < 1 {
			continue
		}
		typeName := thestruct.Type(v.Type()).Name()

		tags, err := thestruct.ParseLiteral(string(fieldT.Tag))
		if err != nil {
			return res, err
		}

		jsonTag := tags.Tag("json")
		formTag := tags.Tag("form")

		// do not encode fields w/o json: struct tag
		if jsonTag == nil {
			continue
		}

		if jsonTag.IsOpt("omitempty") && isEmptyValue(v) {
			continue
		}

		if len(typeName) == 0 && formTag != nil && formTag.Value == "file" {

			inputFile := v.Interface().(*InputFile)
			path := inputFile.Name()

			f, err := os.Open(path)
			if err != nil {
				return res, err
			}
			if f != nil {
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

		mpw.WriteField(jsonTag.Value, fmt.Sprintf("%v", v.Interface()))
	}

	// write closing boundary into buf
	mpw.Close()

	return p.postRequest(url, mpw.FormDataContentType(), &buf)
}

func isEmptyValue(v reflect.Value) bool {

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool() == false
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:

		return fmt.Sprintf("%v", v.Interface()) == "0"

	case reflect.String:
		return fmt.Sprintf("%v", v.Interface()) == ""

	case
		reflect.Uintptr,
		reflect.Interface,
		reflect.Ptr,
		reflect.UnsafePointer:

		return v.Interface() == nil
	}

	return false
}

//
// Makes request and decodes API result
// into GenericResponse-based object
//
// returns error if API result decoded and not OK
//
func (p *bot) getAPIResponse(methodName string, sender fCommandSender, bodyStruct interface{}, resultStruct interface{}) error {

	data, err := sender(methodName, bodyStruct)
	if err != nil {
		return err
	}

	//	log.Printf("[%s]%s", methodName, data)

	if err := json.Unmarshal(data, resultStruct); err != nil {
		return fmt.Errorf("Unmarshal error:%v | %s", err, data)
	}

	// name of embedded struct with common fields (OK, ErrorCode, ...)
	commonStruct := "GenericResponse"

	r := reflect.ValueOf(resultStruct).Elem()
	f := r.FieldByName(commonStruct)
	if !f.IsValid() {
		return fmt.Errorf("%v not found in target struct: %T", commonStruct, resultStruct)
	}

	// see "type assertion" topic ;)
	resp, ok := f.Interface().(GenericResponse)
	if !ok {
		// field is not embedded but has (commonStruct) name
		return fmt.Errorf("ERROR: unmarshal target field is not %v type:%T", commonStruct, resultStruct)
	}

	if !resp.OK {
		return fmt.Errorf("API ERROR(%v)[%v]:%v", methodName, resp.ErrorCode, resp.Description)
	}

	return nil
}
