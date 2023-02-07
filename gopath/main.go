package main

import (
	//_ "github.com/go-redis/redis/v8"  // “_”为空导入
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Printf("thello module mode")
	logrus.Printf(uuid.NewString())

}

// go  mod init  github.com/funkypants/gomodule
//
