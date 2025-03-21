package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var accessToken = flag.String("acc", "", "access token")

func main() {
	flag.Parse()

	// Define WebSocket server URL
	serverURL := "ws://localhost:8000/api/v1/chats/ws"

	// Connect to the WebSocket server
	header := http.Header{}
	header.Add("Authorization", "Bearer "+*accessToken)

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, header)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Start a goroutine to read messages from the server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			fmt.Println("Received from server:", string(message))
		}
	}()

	// Read input from stdin and send to WebSocket
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages (Ctrl+C to exit):")
	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		msg := fmt.Sprintf(`{
			"receiverId": %s,
			"type": "TEXT",
			"content": "%s"
		}`, tokens[0], tokens[1])

		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner error:", err)
	}
}
