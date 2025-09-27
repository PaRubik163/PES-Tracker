package main

import (
	"tracker/internal/app"
	"tracker/internal/config"

	"github.com/sirupsen/logrus"
)

func main(){
	conf := config.NewConfig()

	application, err := app.NewApp(conf)

	if err != nil{
		logrus.Fatal(err)
	}

	application.Run()
}