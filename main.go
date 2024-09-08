package main

import (
	"bufio"
	"log"
	"monkeylsp/rpc"
	"os"
)

func main() {

	logger := getLogger("/home/perry/monkey-ls/log.txt")
	logger.Println("logger started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)

	}

}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[monkeylsp]", log.Ldate|log.Ltime|log.Lshortfile)

}
