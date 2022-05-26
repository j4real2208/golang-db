package error

import "go.uber.org/zap"



func Getlogger () *zap.Logger {	
	logger , _ := zap.NewProduction()
	return logger
}