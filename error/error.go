package error

import "go.uber.org/zap"



func Getlogger () *zap.Logger {	
	logger , _ := zap.NewProduction()
	return logger
}



var Logger *zap.Logger = Getlogger();



//logger.Info("Hello Starting the logger engine");
