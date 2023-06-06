package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"hospital-admin-api/global"
)

func LogJsonInfo(v interface{}) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}

	jsonString := string(jsonData)
	fmt.Println(jsonString)

	global.Logger.WithFields(logrus.Fields{
		"data": jsonString,
	}).Info("info")
}
