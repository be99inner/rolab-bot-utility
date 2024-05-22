package networking

import (
	"log"

	"github.com/gorilla/websocket"
)

// connects to a WebSocket server at the given URL
func ConnectWebSocket(url string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return nil, err
	}
	return conn, nil
}

// sends serialized GameData to the WebSocket server
func SendData(conn *websocket.Conn, data GameData) error {
	jsonData, err := Serialize(data)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Printf("Error writing message: %v\n", err)
		return err
	}

	return nil
}

// receives data from the WebSocket server and deserializes it into GameData
func ReceiveData(conn *websocket.Conn) (GameData, error) {
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading message: %v\n", err)
		return GameData{}, err
	}

	data, err := Deserialize(message)
	if err != nil {
		return GameData{}, err
	}

	return data, nil
}
