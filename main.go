package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"monkeylsp/analysis"
	"monkeylsp/lsp"
	"monkeylsp/rpc"
	"os"
)

func main() {

	logger := getLogger("/home/perry/monkey-ls/log.txt")
	logger.Println("logger started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error:", err)
		}

		handleMessage(logger, writer, state, method, contents)

	}

}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	// logger.Println(msg)
	logger.Println("Recieved msg with method:", method)

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

		msg := lsp.NewInitializseResponse(request.ID)
		writeResponse(writer, msg)

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Failure to parse.", err)
			return
		}

		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Failure to parse textdocument/didchange.", err)
			return
		}

		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)

		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("Failure to parse textdocument/hover.", err)
			return
		}

		// create a response
		response := lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &request.ID,
			},
			Result: lsp.HoverResult{
				Contents: "Hello, from LSP",
			},
		}
		writeResponse(writer, response)

	default:
		logger.Println("unhandled method recieved. method=", method)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[monkeylsp]", log.Ldate|log.Ltime|log.Lshortfile)

}
