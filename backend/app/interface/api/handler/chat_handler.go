package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"

	"github.com/ktakenaka/go-random/backend/app/domain/service"
	"github.com/ktakenaka/go-random/backend/app/interface/api/presenter"
)

// TODO: move
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// TODO: associate with authentication
var (
	chatHelper     = service.ChatHelper{}
	connectedUsers = make(map[string]*service.ChatUser)
)

const (
	commandSubscribe = iota
	commandUnsubscribe
	commandChat
)

// ChatHandler handler for chat
type ChatHandler struct {
	BaseHandler
}

// NewChatHandler constructor
func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

// GetWS web socket handler for chatting
func GetWS(ctx *gin.Context) {
	var (
		rdb *redis.Client
	)
	r := ctx.Request
	w := ctx.Writer

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		handleWSError(err, conn)
		return
	}

	err = onConnect(r, conn, rdb)
	if err != nil {
		handleWSError(err, conn)
	}

	closeCh := onDisconnect(r, conn, rdb)

	onChannelMessage(conn, r)

loop:
	for {
		select {
		case <-closeCh:
			break loop
		default:
			onUserMessage(conn, r, rdb)
		}
	}
}

func handleWSError(err error, conn *websocket.Conn) {

}

func onConnect(r *http.Request, conn *websocket.Conn, rdb *redis.Client) error {
	username := r.URL.Query()["username"][0]
	fmt.Println("connected from:", conn.RemoteAddr(), "user:", username)

	u, err := chatHelper.Connect(rdb, username)
	if err != nil {
		return err
	}
	connectedUsers[username] = u
	return nil
}

func onDisconnect(r *http.Request, conn *websocket.Conn, rdb *redis.Client) chan struct{} {
	closeChan := make(chan struct{})
	username := r.URL.Query()["username"][0]

	// use rdb connection
	fmt.Println(rdb)

	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed for user", username)

		u := connectedUsers[username]
		if err := u.Disconnect(); err != nil {
			return err
		}
		delete(connectedUsers, username)
		close(closeChan)
		return nil
	})

	return closeChan
}

func onChannelMessage(conn *websocket.Conn, r *http.Request) {

}

func onUserMessage(conn *websocket.Conn, r *http.Request, rdb *redis.Client) {
	var msg presenter.ChatMessage

	if err := conn.ReadJSON(&msg); err != nil {
		handleWSError(err, conn)
		return
	}

	username := r.URL.Query()["username"][0]
	u := connectedUsers[username]

	switch msg.Command {
	case commandSubscribe:
		if err := u.Subscribe(rdb, msg.Channel); err != nil {
			handleWSError(err, conn)
		}
	case commandUnsubscribe:
		if err := u.Unsubscribe(rdb, msg.Channel); err != nil {
			handleWSError(err, conn)
		}
	case commandChat:
		if err := u.Chat(rdb, msg.Channel, msg.Content); err != nil {
			handleWSError(err, conn)
		}
	}
}
