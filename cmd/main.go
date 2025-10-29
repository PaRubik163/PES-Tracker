package main

import (
	"fmt"
	"log"
	"tracker/internal/app"
	"tracker/internal/config"
	"tracker/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main(){
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	gin.SetMode(gin.ReleaseMode)

	conf := config.NewConfig()
	addr := fmt.Sprintf("%s:%s", conf.GRPCHost, conf.GRPPCPort)
	
	grpcClient, err := logger.NewClient(addr)
	loggerAdapter := logger.NewAdapter(grpcClient)

	if err != nil{
		log.Fatal(err)
	}

	application, err := app.NewApp(conf, loggerAdapter)

	if err != nil{
		logrus.Fatal(err)
	}

	application.Run()
}