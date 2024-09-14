package lsp

type TextDocumentDidChangeNotification struct {
	Notifictaion
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionTextDocumentIdentifier   `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEven `json:"contentChanges"`
}

type TextDocumentContentChangeEven struct {
	Text string `json:"text"`
}
