package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"simple_lsp/analysis"
	"simple_lsp/lsp"
	"simple_lsp/rpc"
)

func main() {
	fmt.Println("HELLO")

	logger := getLogger("C:\\Users\\Admin\\Desktop\\Work\\Go\\Simple_LSP\\log.txt")
	logger.Println("Started.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Recived an error: %s\n", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents)
	}

}

func handleMessage(logger *log.Logger, writer io.Writer,
	state analysis.State, method string, contents []byte) {

	logger.Printf("Recived a message with method %s:\n", method)

	switch method {
	case "initialize":
		var request lsp.InitRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse this: %s\n", err)
			return
		}

		logger.Printf(
			"Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitResponse(request.ID)
		writeResponse(writer, msg)

		logger.Print("Sent the reply.")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s\n", err)
			return
		}

		logger.Printf("Opened: %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s\n", err)
			return
		}

		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		response := state.Hover(
			request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition: %s", err)
			return
		}

		response := state.Definition(
			request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)

	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("ERROR: no file found")
	}

	return log.New(logfile, "[simple_lsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
