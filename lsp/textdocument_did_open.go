package lsp

type DidOpenTextDocumentNotification struct {
	Notifictaion
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
