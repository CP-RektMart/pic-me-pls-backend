package chat

type ChatSystem struct {
	clients map[uint]chan string
}

var instance *ChatSystem

func NewChatSystem() *ChatSystem {
	if instance == nil {
		instance = &ChatSystem{
			clients: make(map[uint]chan string),
		}
	}
	return instance
}

func (c *ChatSystem) Register(userID uint) chan string {
	channel := make(chan string)
	c.clients[userID] = channel
	return channel
}

func (c *ChatSystem) Logout(userID uint) {
	if c.IsUserExist(userID) {
		delete(c.clients, userID)
	}
}

func (c *ChatSystem) SendMessage(receiverID uint, message string) {
	if c.IsUserExist(receiverID) {
		c.clients[receiverID] <- message
	}
}

func (c *ChatSystem) IsUserExist(userID uint) bool {
	_, ok := c.clients[userID]
	return ok
}
