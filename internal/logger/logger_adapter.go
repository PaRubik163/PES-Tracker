package logger

import(
	pb "tracker/pkg/logger"
)

type Adapter struct{
	client *Client
}

func NewAdapter(client *Client) *Adapter {
	return &Adapter{
		client: client,
	}
}

func (a *Adapter) Users(level string, data *pb.LogData) {
	a.client.Log("Users", level, data)
}

func (a *Adapter) Income(level string, data *pb.LogData) {
	a.client.Log("Income", level, data)
}

func (a *Adapter) Expense(level string, data *pb.LogData) {
	a.client.Log("Expense", level, data)
}

func (a *Adapter) Subscription(level string, data *pb.LogData) {
	a.client.Log("Subscription", level, data)
}