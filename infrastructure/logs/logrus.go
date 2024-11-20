package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger *logrus.Logger

func InitLogrus() {
	baseDir, _ := os.Getwd()
	logDir := baseDir + "/logs"
	_, err := os.Stat(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(logDir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}
	}
	logger = logrus.New()
	fileWriter := getLogWriter(logDir + "/app.log")
	//save the file and output the console
	mw := io.MultiWriter(fileWriter, os.Stdout)
	logger.SetOutput(mw)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: FmtDateTime,
	})
}

func GetLogger() *logrus.Logger {
	return logger
}
