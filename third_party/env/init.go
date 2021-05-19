package env

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var once sync.Once

type Util interface {
	Get(key string) string
}

//model for environtment
type util struct {
	data interface{}
}

// Init load environtment from your json file
func Init() Util {
	var data interface{}

	//do this once for each goroutine
	once.Do(func() {

		//default profile
		var profile = ""

		//check if there is input flag, if exits set flag to profile
		if len(os.Args) > 1 {
			profile = "-" + os.Args[1]
		}

		//open file config
		absPath, _ := filepath.Abs("configs/config" + profile + ".json")
		environtmentFile, err := os.Open(absPath)
		if err != nil {
			panic("Error, config" + profile + ".json Is Missing In Directory, " + err.Error())
		}

		//defer close file
		defer func() {
			if err := environtmentFile.Close(); err != nil {
				panic(err.Error())
			}
		}()

		//decode json file
		jsonParser := json.NewDecoder(environtmentFile)
		err = jsonParser.Decode(&data)
		if err != nil {
			panic(err.Error())
		}

	})
	return &util{
		data: data,
	}
}

// InitPath load environtment from your json file
func InitPath(path string) Util {
	var data interface{}

	//do this once for each goroutine
	once.Do(func() {
		//open file config
		environtmentFile, err := os.Open(path)
		if err != nil {
			panic(err.Error())
		}

		//defer close file
		defer func() {
			if err := environtmentFile.Close(); err != nil {
				panic(err.Error())
			}
		}()

		//decode json file
		var temp interface{}
		jsonParser := json.NewDecoder(environtmentFile)
		err = jsonParser.Decode(&temp)
		if err != nil {
			panic(err.Error())
		}

	})

	return &util{
		data: data,
	}
}

//map environment to map str interface
func (c *util) doMapify() (map[string]interface{}, error) {
	if m, ok := c.data.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("can't type assert with map[str]interface{}")
}

func (c util) Get(key string) string {
	m, err := c.doMapify()
	if err == nil {
		if val, ok := m[key]; ok {
			c := &util{val}
			if s, ok := c.data.(string); ok {
				return s
			}
			panic("Error Conversion, Field Is Not String")
		}
	}
	return ""
}

func GetEnvPath() string {
	_, filename, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(filename, "../../"+string(filepath.Separator))))
	return apppath
}
