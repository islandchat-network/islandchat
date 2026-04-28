package ws

import (
	"testing"
	"time"

	"github.com/islandchat-network/islandchat/backend/internal/models"
)

// TestHub_Broadcast tests that a message sent to the hub is distributed to all registered clients.
func TestHub_Broadcast(t *testing.T) {
	// 1. Arrange: Setup the Hub and Clients
	hub := NewHub()
	go hub.Run() // Start the hub in a separate goroutine

	// Create two dummy clients with buffered channels
	client1 := &Client{Hub: hub, Send: make(chan models.Message, 10)}
	client2 := &Client{Hub: hub, Send: make(chan models.Message, 10)}

	// Register clients
	hub.Register <- client1
	hub.Register <- client2

	// Give the hub a tiny moment to process the registrations
	time.Sleep(50 * time.Millisecond)

	// 2. Act: Broadcast a message
	testMsg := models.Message{
		ID:       "msg-123",
		Content:  "Hello from the automated test!",
		SenderID: "test-bot",
		Type:     0,
	}
	hub.Broadcast <- testMsg

	// 3. Assert: Check if both clients received the message
	verifyClientReceivesMessage(t, client1, testMsg, "Client 1")
	verifyClientReceivesMessage(t, client2, testMsg, "Client 2")
}

// TestHub_Unregister tests that an unregistered client stops receiving messages.
func TestHub_Unregister(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	client := &Client{Hub: hub, Send: make(chan models.Message, 10)}
	hub.Register <- client
	time.Sleep(50 * time.Millisecond) // Wait for registration

	// Act: Unregister the client
	hub.Unregister <- client
	time.Sleep(50 * time.Millisecond) // Wait for unregistration

	// Broadcast a message
	hub.Broadcast <- models.Message{Content: "Should not be received"}

	// Assert: The client's send channel should be closed or empty
	select {
	case msg, ok := <-client.Send:
		if ok {
			t.Errorf("Client received a message after being unregistered: %v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		// This is the expected path: timeout means no message was received
	}
}

// Helper function to keep tests clean
func verifyClientReceivesMessage(t *testing.T, client *Client, expectedMsg models.Message, clientName string) {
	t.Helper() // Marks this as a helper so error line numbers point to the actual test

	select {
	case receivedMsg := <-client.Send:
		if receivedMsg.Content != expectedMsg.Content {
			t.Errorf("%s received wrong content: got %v want %v", clientName, receivedMsg.Content, expectedMsg.Content)
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout: %s did not receive the broadcast message", clientName)
	}
}
