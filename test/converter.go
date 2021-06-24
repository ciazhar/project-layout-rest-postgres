package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

func ToStruct(filename string, obj interface{}) {
	_, b, _, _ := runtime.Caller(0)
	path := fmt.Sprintf(filepath.Dir(b) + "/data/" + filename)
	expected, _ := ioutil.ReadFile(path)
	if err := json.Unmarshal(expected, &obj); err != nil {
		panic(err)
	}
}

func ToReader(obj interface{}) *strings.Reader {
	jsonPayload, _ := json.Marshal(obj)
	return strings.NewReader(string(jsonPayload))
}
