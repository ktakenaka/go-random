package service

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	usersKey       string = "usersKey"
	channelsKey    string = "channelsKey"
	userChannelFmt string = "userChannelFmt-%v"
)

// ChatUser user for chat function
// TODO: refactor
type ChatUser struct {
	name            string
	channelsHandler *redis.PubSub

	stopListnerChan chan struct{}
	listening       bool

	MessageChan chan redis.Message
}

// ChatHelper should be moved somewhere
type ChatHelper struct{}

//Connect connect user to user channels on redis
// TODO: (refactor) define as ChatUser method
func (c ChatHelper) Connect(rdb *redis.Client, name string) (*ChatUser, error) {
	if _, err := rdb.SAdd(usersKey, name).Result(); err != nil {
		return nil, err
	}

	u := &ChatUser{
		name:            name,
		stopListnerChan: make(chan struct{}),
		MessageChan:     make(chan redis.Message),
	}

	if err := u.connect(rdb); err != nil {
		return nil, err
	}
	return u, nil
}

// Disconnect from websocket
func (u *ChatUser) Disconnect() error {
	if u.channelsHandler != nil {
		if err := u.channelsHandler.Unsubscribe(); err != nil {
			return err
		}
		if err := u.channelsHandler.Close(); err != nil {
			return err
		}
	}
	if u.listening {
		u.stopListnerChan <- struct{}{}
	}

	close(u.MessageChan)

	return nil
}

func getUserChannelsKey(name /*chat user name*/ string) string {
	return fmt.Sprintf(userChannelFmt, name)
}

// Subscribe redis
func (u *ChatUser) Subscribe(rdb *redis.Client, channel string) error {
	userChannelsKey := getUserChannelsKey(u.name)

	if rdb.SIsMember(userChannelsKey, channel).Val() {
		return nil
	}
	if err := rdb.SAdd(userChannelsKey, channel).Err(); err != nil {
		return err
	}

	return u.connect(rdb)
}

// Unsubscribe redis
func (u *ChatUser) Unsubscribe(rdb *redis.Client, channel string) error {
	userChannelsKey := getUserChannelsKey(u.name)

	if !rdb.SIsMember(userChannelsKey, channel).Val() {
		return nil
	}
	if err := rdb.SRem(userChannelsKey, channel).Err(); err != nil {
		return nil
	}

	return u.connect(rdb)
}

func (u *ChatUser) connect(rdb *redis.Client) error {
	var c []string

	c1, err := rdb.SMembers(channelsKey).Result()
	if err != nil {
		return err
	}
	c = append(c, c1...)

	c2, err := rdb.SMembers(getUserChannelsKey(u.name)).Result()
	if err != nil {
		return err
	}
	c = append(c, c2...)

	if len(c) == 0 {
		fmt.Println("no channels to connect to for user: ", u.name)
		return nil
	}

	if u.channelsHandler != nil {
		if err := u.channelsHandler.Unsubscribe(); err != nil {
			return err
		}
		if err := u.channelsHandler.Close(); err != nil {
			return err
		}
	}

	if u.listening {
		u.stopListnerChan <- struct{}{}
	}

	return u.doConnect(rdb, c...)
}

func (u *ChatUser) doConnect(rdb *redis.Client, channels ...string) error {
	pubSub := rdb.Subscribe(channels...)

	u.channelsHandler = pubSub

	go func() {
		u.listening = true
		fmt.Println("starting the listener for user:", u.name, "on channels:", channels)

	loop:
		for {
			select {
			case msg, ok := <-pubSub.Channel():
				if !ok {
					break loop
				}
				u.MessageChan <- *msg
			case <-u.stopListnerChan:
				fmt.Println("stopping the listener for user:", u.name)
				break loop
			}
		}
	}()

	return nil
}

// Chat send message
func (u *ChatUser) Chat(rdb *redis.Client, channel string, content string) error {
	return rdb.Publish(channel, content).Err()
}
