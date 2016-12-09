package main

import (
	"github.com/MounirOnGithub/go-rest-service/utils"
	"github.com/Sirupsen/logrus"
	"net/http"
)

func main() {
	r := utils.NewRouter()
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		logrus.Error(err)
	}
}
