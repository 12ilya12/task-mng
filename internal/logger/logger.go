package logger

import (
	"log"
	"sync"
)

type LogEntry struct {
	LogType string
	Message string
}

type Logger struct {
	logChan chan LogEntry

	//WaitGroup используем, чтобы дождаться обработки всех сообщений при остановке
	wg sync.WaitGroup
}

func NewLogger() *Logger {
	return &Logger{
		logChan: make(chan LogEntry, 10),
	}
}

func (l *Logger) Log(logType, message string) {
	l.logChan <- LogEntry{LogType: logType, Message: message}
}

func (l *Logger) Start() {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for entry := range l.logChan {
			log.Printf("%s: %s", entry.LogType, entry.Message)
		}
	}()
}

func (l *Logger) Stop() {
	close(l.logChan)
	l.wg.Wait()
}
