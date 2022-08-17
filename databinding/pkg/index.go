package pkg

import "encoding/json"

// JSONMarshal
// JSON encoding
func JSONMarshal(instance interface{}) (bytes []byte, err error) {
	// Marshal
	bytes, err = json.Marshal(instance)
	if err != nil {
		return nil, err
	}

	// Returns
	return bytes, nil
}
