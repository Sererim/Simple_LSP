package lsp

type InitRequest struct {
	Request
	Params InitReqParams `json:"params"`
}

type InitReqParams struct {
	ClientInfo *ClientInfo `jsong:"clientInfo"`
	// ... More params for real LSP
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitResponse struct {
	Response
	Result InitResult `json:"result"`
}

type InitResult struct {
	Capabilities ServerCpbl `json:"capabilities"`
	ServerInfo   ServerInfo `json:"serverInfo"`
}

type ServerCpbl struct {
	TextDocSync int `json:"textDocumentSync"`

	HoverProvider      bool `json:"hoverProvider"`
	DefinitionProvider bool `json:"definitionProvider"`
	CodeActionProvider bool `json:"codeActionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitResponse(id int) InitResponse {
	return InitResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitResult{
			Capabilities: ServerCpbl{
				TextDocSync:        1,
				HoverProvider:      true,
				DefinitionProvider: true,
			},
			ServerInfo: ServerInfo{
				Name:    "simple_lsp",
				Version: "0.0.1",
			},
		},
	}
}
