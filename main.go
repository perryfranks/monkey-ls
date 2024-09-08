package main

import (
	"bufio"
	"encoding/json"
	"log"
	"monkeylsp/lsp"
	"monkeylsp/rpc"
	"os"
)

func main() {

	logger := getLogger("/home/perry/monkey-ls/log.txt")
	logger.Println("logger started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}

		handleMessage(logger, method, contents)

	}

}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	// logger.Println(msg)
	logger.Println("Recieved msg with method: ", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Failure to parse.", err)
			return
		}
		logger.Printf("connected to %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
	default:
		logger.Println("unhandled method recieved. method=%s", method)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[monkeylsp]", log.Ldate|log.Ltime|log.Lshortfile)

}
