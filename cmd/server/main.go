package main

import (
	"github.com/giicoo/golang_CRUD/internal/app"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Server Start")
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}
