package networking

import (
	"encoding/json"
	"log"
)

// GameData represents the structure of data being sent and received
type GameData struct {
	EventType string `json:"eventType"`
	Payload   string `json:"payload"`
}

// Serialize serialized GameData in JSON
func Serialize(data GameData) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error serializing data: %v\n", err)
		return nil, err
	}
	return jsonData, nil
}

// Deserialize deserialize JSON into GameData
func Deserialize(jsonData []byte) (GameData, error) {
	var data GameData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("Error deserializing data: %v\n", err)
		return GameData{}, err
	}
	return data, nil
}
