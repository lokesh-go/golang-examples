package pkg

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
)

// JSON encode
func JSONMarshal(instance interface{}) (bytes []byte, err error) {
	// Marshal
	bytes, err = json.Marshal(instance)
	if err != nil {
		return nil, err
	}

	// Returns
	return bytes, nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by instance
func JSONUnmarshal(bytes []byte, instance interface{}) (err error) {
	err = json.Unmarshal(bytes, instance)
	if err != nil {
		return err
	}

	// Returns
	return nil
}

// BSON encode
func BSONMarshal(instance interface{}) (bytes []byte, err error) {
	// Marshal
	bytes, err = bson.Marshal(instance)
	if err != nil {
		return nil, err
	}

	// Returns
	return bytes, nil
}

// Unmarshal parses the BSON-encoded data and stores the result in the value
// pointed to by instance
func BSONUnmarshal(bytes []byte, instance interface{}) (err error) {
	// Unmarshal
	err = bson.Unmarshal(bytes, instance)
	if err != nil {
		return err
	}

	// Returns
	return nil
}

// Mapstructure is a Go library for decoding generic map values to structures and vice versa
func MapStructureDecode(input interface{}, output interface{}) (err error) {
	// Decodes
	err = mapstructure.Decode(input, output)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
