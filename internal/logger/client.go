package logger

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "tracker/pkg/logger"

	"google.golang.org/grpc"
)

type Client struct{
	client pb.TelegramLoggerClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil{
		return nil, fmt.Errorf("failed to connect to server %v", err)
	}

	client := pb.NewTelegramLoggerClient(conn)

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Log(topic, level string, data *pb.LogData){
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.client.SendLog(ctx, &pb.LogRequest{
		Topic: topic,
		Level: level,
		Data: data,
	})

	if err != nil || !resp.Ok {
		log.Print("failed to send log %v", err)
	}
}