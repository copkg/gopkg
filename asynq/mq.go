package asynq

import (
	"github.com/hibiken/asynq"
)

var Client *asynq.Client
var Server *asynq.Server

func NewClient(redisAddr string) *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	//defer client.Close()
	return client
}
func NewServer(redisAddr string) *asynq.Server {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		})

	return srv
}
