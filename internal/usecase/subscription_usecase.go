package usecase

import (
	"fmt"
	"time"
	"tracker/internal/entity"
	"tracker/internal/logger"
	"tracker/internal/repository"
	pb "tracker/pkg/logger"
)

type SubscriptionUseCase struct {
	subscriptionRepo *repository.SubscriptionRepo
	loggerAdapter *logger.Adapter
}

func NewSubscriptionUseCase(sr *repository.SubscriptionRepo, loggerAdapter *logger.Adapter) *SubscriptionUseCase {
	return &SubscriptionUseCase{
		subscriptionRepo: sr,
		loggerAdapter: loggerAdapter,
	}
}

func (su *SubscriptionUseCase) CreateSubscription(sub *entity.Subscription) error {
	start := time.Now()

	err := sub.CheckNewSubscription()

	if err != nil{
		return err
	}

	err = su.subscriptionRepo.Create(sub)

	if err != nil{
		su.loggerAdapter.Subscription("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", sub.UserID),
			Token: sub.User.Login,
			Method: "POST",
			Path: "/new_subscription",
			Status: logger.Internal,
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v",time.Since(start)),
			Error: err.Error(),
		})
		return err
	}

	su.loggerAdapter.Subscription("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", sub.UserID),
		Token: sub.User.Login,
		Method: "POST",
		Path: "/new_subscription",
		Status: logger.Created,
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v",time.Since(start)),
		Error: "",
	})
	return nil
}

func (su *SubscriptionUseCase) GetAllSubscriptions(userID int) ([]entity.Subscription, error) {
	start := time.Now()

	subs, err := su.subscriptionRepo.GetAll(userID)

	if err != nil{
		su.loggerAdapter.Subscription("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "GET",
			Path: "/subscriptions",
			Status: logger.Internal,
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v",time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}
	
	su.loggerAdapter.Subscription("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d",userID),
		Token: "",
		Method: "GET",
		Path: "/subscriptions",
		Status: logger.Ok,
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v",time.Since(start)),
		Error: "",
	})

	return subs, nil
}

func (su *SubscriptionUseCase) DeleteSubscription(subID, userID int) error {
	start := time.Now()

	if err := su.subscriptionRepo.DeleteByID(subID, userID); err != nil{
		su.loggerAdapter.Subscription("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", userID),
			Token: "",
			Method: "DELETE",
			Path: "/delete_subscription",
			Status: logger.Internal,
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v",time.Since(start)),
			Error: err.Error(),
		})
		return err
	}

	su.loggerAdapter.Subscription("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", userID),
		Token: "",
		Method: "DELETE",
		Path: fmt.Sprintf("/subscription/%d", subID),
		Status: logger.Ok,
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v",time.Since(start)),
		Error: "",
	})
	return nil
}