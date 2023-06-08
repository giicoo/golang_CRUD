package main

import (
	"log"

	"github.com/giicoo/golang_CRUD/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Server Start")
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
