package testdata

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

func ToStruct(filename string, obj interface{}) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	expected, _ := ioutil.ReadFile(basepath + "/" + filename)

	if err := json.Unmarshal(expected, &obj); err != nil {
		panic(err)
	}
}

func ToReader(obj interface{}) *strings.Reader {
	jsonPayload, _ := json.Marshal(obj)
	return strings.NewReader(string(jsonPayload))
}
