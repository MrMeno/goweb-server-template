package Utils

import (
	"MetaWebServer/DataReflect/Config"
	"database/sql/driver"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	ERROR_PRX = "[ERROR] "
	AUTH_PRX  = "[AUTH] "
	INFO_PRX  = "[INFO] "
	WARN_PRX  = "[WARN] "
)

type CFG[T Config.DBConfig | Config.AuthConfig | Config.PropsConfig] struct {
	TYPE T
}

func (s CFG[T]) ReadFromYaml(path string) (*T, error) {
	var err error
	config := s.TYPE
	fileByte, err := os.Open(path)
	if err != nil {
		fileByte.Close()
		log.Printf("config read err:%#v\n", err.Error())
		return nil, err
	}
	if file, err := io.ReadAll(fileByte); err != nil {
		log.Printf("config read err:%#v\n", err.Error())
		fileByte.Close()
		return nil, err
	} else {
		defer yaml.Unmarshal(file, &config)
		fileByte.Close()
	}
	return &config, err
}

func IsNilOrEmpty(s interface{}) bool {
	if s == "" || s == nil {
		return true
	}
	return false
}

type BitBool bool

func (b BitBool) Value() (driver.Value, error) {
	result := make([]byte, 1)
	if b {
		result[0] = byte(1)
	} else {
		result[0] = byte(0)
	}
	return result, nil
}

func (b BitBool) Scan(v interface{}) error {
	bytes := v.([]byte)
	if bytes[0] == 0 {
		b = false
	} else {
		b = true
	}
	return nil
}
