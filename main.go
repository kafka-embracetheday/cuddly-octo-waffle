package main

import (
	"github.com/kafka-embracetheday/cuddly-octo-waffle/common/db"
	"github.com/kafka-embracetheday/cuddly-octo-waffle/common/logger"
	"github.com/kafka-embracetheday/cuddly-octo-waffle/config"
)

func main() {
	config.Init()
	logger.Init()
	db.Init()
}
