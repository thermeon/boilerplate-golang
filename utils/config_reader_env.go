package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	envFieldTag = "env"
	optional    = "optional"
)

type envConfigReader struct{}

// NewEnvConfigReader create an instance of environment configuration reader.
func NewEnvConfigReader() ConfigReader {
	return &envConfigReader{}
}

// Read fill up configurations i pointer of the configuration struct.
// The fields of the structs should be as bellow.
func (e *envConfigReader) Read(i interface{}) (err error) {
	rv := reflect.ValueOf(i)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("given value isn't a pointer")
	}
	rt := reflect.TypeOf(i).Elem()
	for j := 0; j < rt.NumField(); j++ {
		f := rt.Field(j)
		if val, ok := f.Tag.Lookup(envFieldTag); ok {
			tagValue, isOptional := validateTags(val)
			if err = e.setValue(rv.Elem().Field(j), os.Getenv(tagValue), isOptional); err != nil {
				err = errors.New(fmt.Sprintf("invalid or empty value for %s environment variable", val))
				return
			}
		}
	}
	return
}

// setValue set values to the relevant fields with considering data type of the field.
func (e *envConfigReader) setValue(value reflect.Value, readValue string, isOptional bool) (err error) {
	if len(readValue) == 0 && isOptional {
		return
	}
	k := value.Kind()
	switch k {
	case reflect.Bool:
		var convertedVal bool
		convertedVal, err = strconv.ParseBool(readValue)
		value.SetBool(convertedVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var convertedVal int64
		convertedVal, err = strconv.ParseInt(readValue, 10, 64)
		value.SetInt(convertedVal)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var convertedVal uint64
		convertedVal, err = strconv.ParseUint(readValue, 10, 64)
		value.SetUint(convertedVal)

	case reflect.Float32, reflect.Float64:
		var convertedVal float64
		convertedVal, err = strconv.ParseFloat(readValue, 64)
		value.SetFloat(convertedVal)

	case reflect.Complex64, reflect.Complex128:
	case reflect.Array:
	case reflect.String:
		value.SetString(readValue)
	case reflect.Ptr:
		err = json.Unmarshal([]byte(readValue), value.Interface())
	}
	return
}

// validateTags validate the field tags.( example: with the optional field)
func validateTags(val string) (tagValue string, isOptional bool) {
	tagValue = ""
	isOptional = false
	arr := strings.Split(val, ";")
	if len(arr) == 0 {
		return
	}
	for i, s := range arr {
		if i == 0 {
			tagValue = s
			continue
		}
		if s == optional {
			isOptional = true
		}
	}
	return
}
