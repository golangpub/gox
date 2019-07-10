package gox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gopub/log"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

func JSONUnmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

func JSONUnmarshalFile(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return JSONUnmarshal(data, v)
}

func JSONMarshalStr(i interface{}) string {
	if i == nil {
		return ""
	}

	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if val.IsNil() {
			return ""
		}
	}
	d, err := json.Marshal(i)
	if err != nil {
		log.Error(err)
		return ""
	}

	if len(d) == 0 {
		return ""
	}
	return string(d)
}

func JSONMarshalStrIndent(i interface{}, indent string) string {
	if i == nil {
		return ""
	}

	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if val.IsNil() {
			return ""
		}
	}
	d, err := json.MarshalIndent(i, "", indent)
	if err != nil {
		log.Error(err)
		return ""
	}

	if len(d) == 0 {
		return ""
	}
	return string(d)
}

func ToMap(i interface{}) (m M, err error) {
	var data []byte

	data, err = json.Marshal(i)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Error(err)
	}

	return
}

func FetchJSON(url string, httpMethod string, header http.Header, parameters M) M {
	if len(httpMethod) == 0 {
		httpMethod = "GET"
	}

	if len(url) == 0 {
		return nil
	}

	//LogDebug(httpMethod, url, parameters)

	var reader io.Reader
	if parameters == nil {
		parameters = M{}
	}

	body, err := json.Marshal(parameters)
	if err != nil {
		log.Error(url, err)
		return nil
	}

	reader = bytes.NewReader(body)

	var req *http.Request
	var resp *http.Response

	req, err = http.NewRequest(httpMethod, url, reader)
	if err != nil {
		log.Error(url, err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header[k] = v
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Error(url, err)
		return nil
	} else if resp != nil && resp.StatusCode != http.StatusOK {
		log.Error(url, resp.StatusCode)
		return nil
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result M
	if err = JSONUnmarshal(data, &result); err != nil {
		log.Error(url, err, string(data))
		return nil
	}

	return result
}

func IsEmptyJSON(jsonData []byte) bool {
	if len(jsonData) > 4 {
		return false
	}

	dataStr := string(jsonData)
	return dataStr == "{}" || dataStr == "[]" || dataStr == "null" || dataStr == "NULL"
}
