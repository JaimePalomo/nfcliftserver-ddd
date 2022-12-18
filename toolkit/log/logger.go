package log

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type logger struct {
	mutex   sync.Mutex
	out     io.Writer
	err     io.Writer
	buffer  []byte
	level   int
	service string
}

// Message represents a log line to be sent to the log service
type Message struct {
	Service   string `json:"service"`
	Position  string `json:"position"`
	Level     int    `json:"level"`
	Timestamp string `json:"timestamp"`
	Text      string `json:"text"`
}

const (
	OFF    = 0
	PANIC  = 1
	ERROR  = 2
	INFO   = 3
	DEBUG  = 4
	TRACE  = 5
	ALL    = 10
	FORMAT = "2006/01/02 15:04:05"
)

var instance *logger

func init() {

	level := getLevelValue()
	if level == -1 {
		level = INFO
	}

	service := os.Getenv("SERVICE_NAME")
	if service == "" {
		service = "UNKNOWN"
	}

	instance = &logger{
		mutex:   sync.Mutex{},
		out:     os.Stdout,
		err:     os.Stderr,
		level:   level,
		service: service,
	}

}

// Panic logs a message with Panic level using the default logger before calling panic()
func Panic(v ...interface{}) {
	instance.panic(v...)
}

// Error logs a message with Error level using the default logger
func Error(v ...interface{}) {
	instance.error(v...)
}

// Debug logs a message with Debug level using the default logger
func Debug(v ...interface{}) {
	instance.debug(v...)
}

// Info logs a message with Info level using the default logger
func Info(v ...interface{}) {
	instance.info(v...)
}

// Trace logs a message with Trace level using the default logger
func Trace(v ...interface{}) {
	instance.trace(v...)
}

func (logger *logger) panic(v ...interface{}) {
	if PANIC <= logger.level {
		logger.output(PANIC, fmt.Sprintln(v...))
	}
	panic(fmt.Sprint(v...))
}

func (logger *logger) error(v ...interface{}) {
	if ERROR <= logger.level {
		logger.output(ERROR, fmt.Sprintln(v...))
	}
}

func (logger *logger) info(v ...interface{}) {
	if INFO <= logger.level {
		logger.output(INFO, fmt.Sprintln(v...))
	}
}

func (logger *logger) debug(v ...interface{}) {
	if DEBUG <= logger.level {
		logger.output(DEBUG, fmt.Sprintln(v...))
	}
}

func (logger *logger) trace(v ...interface{}) {
	if TRACE <= logger.level {
		logger.output(TRACE, fmt.Sprintln(v...))
	}
}

func (logger *logger) output(level int, text string) {
	message := Message{
		Service:   logger.service,
		Position:  getPosition(),
		Level:     level,
		Timestamp: time.Now().Format(FORMAT),
		Text:      text,
	}

	logger.consoleOutput(message)

}

func (logger *logger) consoleOutput(message Message) {
	//text := fmt.Sprintf("%s - %s - %s || %s", aurora.Green(message.Timestamp), aurora.Blue(message.Uuid), GetColorLevelText(ctx, message.Level), message.Text)
	text := fmt.Sprintf("%s - %s - %s - %s || %s", aurora.Green(message.Timestamp), aurora.Blue(message.Position), GetColorLevelText(message.Level), message.Text)

	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.buffer = logger.buffer[:0]
	logger.buffer = append(logger.buffer, text...)
	if message.Level > ERROR {
		_, _ = logger.out.Write(logger.buffer)
	} else {
		_, _ = logger.err.Write(logger.buffer)
	}
}

func getPosition() (position string) {

	var funcName string

	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		return "Unknown caller -"
	}

	function := runtime.FuncForPC(pc)
	if function != nil {
		funcName = function.Name() + "()"
		funcName = strings.Split(funcName, ".")[len(strings.Split(funcName, "."))-1]
	}

	position = fmt.Sprintf("%s:%d %s", file, line, funcName)

	return position
}

// GetLevelText returns the string corresponding to the input log level
func GetColorLevelText(level int) (result aurora.Value) {
	switch level {
	case PANIC:
		result = aurora.BgRed(aurora.Black("PANIC"))
	case ERROR:
		result = aurora.BgRed(aurora.Black("ERROR"))
	case INFO:
		result = aurora.White("INFO")
	case DEBUG:
		result = aurora.Cyan("DEBUG")
	case TRACE:
		result = aurora.Cyan("TRACE")
	}

	return result
}

func getLevelValue() (result int) {

	level := os.Getenv("LOG_LEVEL")

	switch level {
	case "OFF":
		return OFF
	case "PANIC":
		return PANIC
	case "ERROR":
		return ERROR
	case "INFO":
		return INFO
	case "DEBUG":
		return DEBUG
	case "TRACE":
		return TRACE
	case "ALL":
		return ALL
	default:
		return -1
	}
}
