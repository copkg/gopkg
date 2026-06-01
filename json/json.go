package json

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func Unmarshal(input []byte, data interface{}) error {
	return json.Unmarshal(input, &data)
}
