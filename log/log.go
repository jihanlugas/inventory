package log

import (
	"github.com/jihanlugas/inventory/config"
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
	if config.Environment == config.Production {
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
	if config.Environment == config.Production {
		now := time.Now()
		sqlErrorLogFilePath := config.LogPath + "/" + sqlErrorFileName + now.Format(config.FormatTime) + ".log"
		systemLogFilePath := config.LogPath + "/" + systemLogFileName + now.Format(config.FormatTime) + ".log"

		sqlLogFile.switchNewFile(sqlErrorLogFilePath)
		sysLogFile.switchNewFile(systemLogFilePath)
	}
}

func Run() {
	if config.Environment != config.Production {
		out := zerolog.ConsoleWriter{Out: os.Stdout}
		System = zerolog.New(out).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		Sql = zerolog.New(out).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	} else {

	}
}
