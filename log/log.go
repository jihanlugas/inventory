package log

import (
	"fmt"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/constant"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

type fileLock struct {
	mu sync.Mutex // 8
	f  *os.File   // 8
}

func (fl *fileLock) Close() (err error) {
	fl.mu.Lock()
	err = fl.f.Close()
	fl.mu.Unlock()
	return
}

func (fl *fileLock) Write(p []byte) (n int, err error) {
	fl.mu.Lock()
	n, err = fl.f.Write(p)
	fl.mu.Unlock()
	return
}

func (fl *fileLock) switchNewFile(filePath string) {
	fl.mu.Lock()
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		_ = fl.f.Close()
		fl.f = f
	}
	fl.mu.Unlock()
}

func CloseAll() {
	if !config.Debug {
		_ = sqlLogFile.Close()
		_ = sysLogFile.Close()
	}
}

var (
	sqlLogFile        fileLock
	sysLogFile        fileLock
	System            zerolog.Logger
	Sql               zerolog.Logger
	sqlErrorFileName  string
	systemLogFileName string
)

func init() {
	sqlErrorFileName = "sql_error."
	systemLogFileName = "system_logger."
}

func ChangeDay() {
	if !config.Debug {
		now := time.Now()
		sqlErrorLogFilePath := config.LogPath + "/" + sqlErrorFileName + now.Format(constant.FormatDateLayout) + ".log"
		systemLogFilePath := config.LogPath + "/" + systemLogFileName + now.Format(constant.FormatDateLayout) + ".log"

		sqlLogFile.switchNewFile(sqlErrorLogFilePath)
		sysLogFile.switchNewFile(systemLogFilePath)
	}
}

func Run() {
	if config.Debug {
		out := zerolog.ConsoleWriter{Out: os.Stdout}
		System = zerolog.New(out).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		Sql = zerolog.New(out).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	} else {
		var err error
		now := time.Now()

		err = os.MkdirAll(config.LogPath, 0755)
		if err != nil {
			fmt.Println("Directory log path is not writeable")
			os.Exit(1)
		}

		sqlErrorLogFilePath := config.LogPath + "/" + sqlErrorFileName + now.Format("2006-01-02") + ".log"
		systemLogFilePath := config.LogPath + "/" + systemLogFileName + now.Format("2006-01-02") + ".log"

		// If the file doesn't exist, create it, or append to the file
		sqlLogFile.f, err = os.OpenFile(sqlErrorLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error open file: ", err)
			os.Exit(1)
		}
		sysLogFile.f, err = os.OpenFile(systemLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error open file: ", err)
			os.Exit(1)
		}

		System = zerolog.New(&sysLogFile).Level(zerolog.WarnLevel).With().Timestamp().Logger()
		Sql = zerolog.New(&sqlLogFile).Level(zerolog.WarnLevel).With().Timestamp().Logger()
	}
}
