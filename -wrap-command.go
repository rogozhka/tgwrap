package tgwrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type GenericResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

//
// Generic method wrapper for any command
//
func (p *bot) sendCommand(methodName string, bodyStruct interface{}) ([]byte, error) {

	// empty result to return with errors
	res := []byte{}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", p.token, methodName)

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(bodyStruct); err != nil {
		return res, err
	}

	resp, err := http.Post(url, "application/json", &buf)
	if resp != nil {
		defer resp.Body.Close()
		resp.Close = true
	}
	if err != nil {
		return res, fmt.Errorf("POST error:%v", err)
	}

	if res, err := ioutil.ReadAll(resp.Body); err != nil {
		return res, fmt.Errorf("ReadAll error:%v", err)
	} else {
		return res, nil
	}
}

func (p *bot) getResponse(methodName string, bodyStruct interface{}, resultStruct interface{}) error {

	data, err := p.sendCommand(methodName, bodyStruct)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, resultStruct); err != nil {
		return fmt.Errorf("Unmarshal error:%v", err)
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
