package outist

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const logFolderName = "log"

type TTextLog struct {
	FilePath              string
	logger                *log.Logger
	StandardOutputEnabled bool
}

var GlobalLog *TTextLog

func CreateTextLog() *TTextLog {
	return &TTextLog{}
}

func (this *TTextLog) Prepare() {
	if this.FilePath != "" {
		var file, openFileResult = os.OpenFile(this.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if openFileResult != nil {
			fmt.Println("Could not create log file at path '" + this.FilePath + "'")
		}
		this.logger = &log.Logger{}
		this.logger.SetOutput(file)
	}
}

func (this *TTextLog) Write(text string) {
	this.WriteL(1, text)
}

func (this *TTextLog) WriteL(level int, text string) {
	var extendedText = FormatFullTime(time.Now()) + " th" + UInt64ToStr(GetThreadId()) + " " + GetCallerName(1+level) + ": " + text
	if this.StandardOutputEnabled {
		fmt.Println(extendedText)
	}
	if this.logger != nil {
		this.logger.Print(extendedText)
	}
}

func StartGlobalLog(appDirectory string) {
	GlobalLog = CreateTextLog()
	GlobalLog.FilePath = appDirectory + "/" + logFolderName + "/" + FormateDateForFileName(time.Now()) + ".txt"
	GlobalLog.StandardOutputEnabled = true
	GlobalLog.Prepare()
}

func GetCallerName(level int) string {
	callers := make([]uintptr, 16)
	runtime.Callers(2+level, callers)
	f := runtime.FuncForPC(callers[0])
	return f.Name()
}

func init() {
	StartGlobalLog()
}
