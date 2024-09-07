package main

import (
	"encoding/json"
)

// Converts the interface of a class into another interface that a client expects

type jsonConverter interface {
	toJson() (map[string]interface{}, error)
}

type textAdapter struct {
	textData string
}

func (adapter textAdapter) toJson() (jData map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(adapter.textData), &jData)
	if err != nil {
		return nil, err
	}
	return jData, nil
}

/*
func main() {
	textData := "{\"app\": \"test-app\", \"host\": \"localhost\", \"ip\": \"172.31.100.12\"}"
	adapter := textAdapter{textData}
	jData, err := adapter.toJson()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jData)
}
*/
