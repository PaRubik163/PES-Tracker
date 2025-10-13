package main

import (
	"tracker/internal/app"
	"tracker/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main(){
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	gin.SetMode(gin.ReleaseMode)

	conf := config.NewConfig()
	application, err := app.NewApp(conf)

	if err != nil{
		logrus.Fatal(err)
	}

	application.Run()
}