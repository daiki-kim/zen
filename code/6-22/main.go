package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Content string `json:"content"`
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	receive chan Message
}

func newHab() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.receive)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				// clientのreceiveチャンネルが開いていれば送信
				case client.receive <- message:
				// clientのreceiveチャンネルがブロックされていれば切断
				default:
					close(client.receive)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (c *Client) serveMessage() {
	// 接続が切れたらclientを削除
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	// clientからのmessageをbroadcastに送信
	for {
		var message Message
		err := c.conn.ReadJSON(&message)
		if err != nil {
			break
		}
		c.hub.broadcast <- message
	}
}

func (c *Client) listenMessage() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.receive:
			if !ok {
				return
			}

			err := c.conn.WriteJSON(message)
			if err != nil {
				return
			}
		}
	}
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		hub:     hub,
		conn:    conn,
		receive: make(chan Message),
	}
	client.hub.register <- client

	go client.listenMessage()
	go client.serveMessage()
}

func main() {
	hub := newHab()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	http.ListenAndServe(":8080", nil)
}
