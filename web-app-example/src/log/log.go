package log

import (
	"log"
	"os"
)

const Problem string = "Problem"
const UsualLog string = "Log"
const UndefinedType string = "Undefined"

func LogWorkApp(msg string, typeMsg string) {
	var logMessage string
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	if typeMsg == Problem {
		logMessage += Problem
	} else if typeMsg == UsualLog {
		logMessage += UsualLog
	} else {
		logMessage += UndefinedType
	}

	logMessage += " : " + msg
	log.SetOutput(f)
	log.Println(logMessage)

}
