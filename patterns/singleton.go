package patterns

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var (
	instance *Log
	once     sync.Once
)

type Log struct {
}

func NewLog() *Log {
	once.Do(func() {
		instance = &Log{}
	})
	return instance
}

func (l Log) LogExecution(mes string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Loger(mes, file)
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func Loger(logMessage string, w io.Writer) {
	_, err := fmt.Fprint(w, "\r\nLog Entry : ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v %v\n", time.Now(), time.Now().Unix())
	fmt.Fprintf(w, "Действие: %v\n", logMessage)
	fmt.Println("-------------------------------")
}

func RunOperation(operationCode rune, operand int) float64 {
	log := Log{}
	rez := 0.0
	switch operationCode {
	case '+':
		rez += float64(operand)
		log.LogExecution(fmt.Sprintf("Сложение %v", operand))
	case '-':
		rez -= float64(operand)
		log.LogExecution(fmt.Sprintf("Вычитание %v", operand))
	case '*':
		rez *= float64(operand)
	case ':':
		fallthrough
	case '/':
		rez /= float64(operand)
	}
	return rez
}
