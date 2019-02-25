package main

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	_ "github.com/fsnotify/fsnotify"
	_ "github.com/jinzhu/gorm"
	_ "go.uber.org/zap"
	_ "go.uber.org/zap/zapcore"
	_ "github.com/json-iterator/go"
	_ "github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/sqs"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/nu7hatch/gouuid"
	_ "github.com/rs/xid"
)

func main() {

}