package Logger

import (
	"fmt"
	"github.com/edwinkessels/datetime"
	"log"
	"os"
	"strings"
)

const (
	WINDOWS_LOGNAME = "c:\\temp\\Muhjee_GoEngine.log"
	LINUX_LOGNAME   = "/tmp/Muhjee_GoEngine.log"
)

var (
	//WarningLogger *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

func InitLogger(inOperatingSystem string) {

	logFileName := LINUX_LOGNAME

	fmt.Println("Log is written to " + logFileName)

	// Rotate the LogFile
	rotateLogFileName := strings.Replace(logFileName, ".log", "", 1) + "." + DateTime.GetTimeStamp() + ".log"

	// If the current Logfile exists, then this should be rotated
	if _, err := os.Stat(logFileName); err == nil {
		// Logfile exists and should be rotated
		fmt.Println("Rotating Logfile")
		os.Rename(logFileName, rotateLogFileName)
	}

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INF: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file, "DBG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
