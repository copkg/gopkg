package beanstalk

import (
	"github.com/iwanbk/gobeanstalk"
	"log"
)

var Client *gobeanstalk.Conn

func Connect(host string) *gobeanstalk.Conn {
	c, err := gobeanstalk.Dial(host)
	if err != nil {
		log.Panicf("dial beanstalk host:%s,get err:%s", host, err.Error())
		return nil
	}
	return c
}
