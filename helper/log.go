package helper

import (
	"fmt"
	"time"
	"os"
)

var Log = Logger{
	logPath: "default.log",
	debugLevel:DebugLevelHigh,
}

const (
	DebugLevelHigh = 10
	DebugLevelLow = 1
	DebugLevelOff = 0
)


type Logger struct {
	logPath string
	debugLevel uint8
}

func (log Logger)Info(msg string){
	msg = time.Now().String() + " - " + "Info: " + msg
	fmt.Println(msg)
	log.writeToLogFile(msg)
}

func (log Logger)writeToLogFile(msg string){
	logFile, err := os.OpenFile(log.logPath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {

		fmt.Println("FATAL ERROR: could't write to logfile at location '" + log.logPath + "' because " + err.Error())
		return
	}

	defer logFile.Close()

	if _, err = logFile.WriteString(msg + "\n"); err != nil {
		fmt.Println("FATAL ERROR: could't write to logfile at location '" + log.logPath + "' because" + err.Error())
	}
}



