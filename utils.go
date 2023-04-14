package utils

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func ConvertYAMLtoJSON(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = ConvertYAMLtoJSON(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertYAMLtoJSON(v)
		}
	}
	return i
}

func GetJSONfromYAML(i interface{}) ([]byte, error) {
	yamlObj := ConvertYAMLtoJSON(i)
	var err error
	var returnJSON []byte

	returnJSON, err = json.Marshal(yamlObj)
	return returnJSON, err
}

func GoTemplateFunc(t *template.Template) map[string]interface{} {
	f := sprig.TxtFuncMap()

	f["include"] = func(name string, data interface{}) (string, error) {
		buf := bytes.NewBuffer(nil)
		if err := t.ExecuteTemplate(buf, name, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	f["toYaml"] = func(v interface{}) string {
		data, err := yaml.Marshal(v)
		if err != nil {
			// Swallow errors inside of a template.
			return ""
		}
		return string(data)
	}

	return f
}
