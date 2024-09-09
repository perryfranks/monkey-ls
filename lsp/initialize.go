package lsp

type InitializeRequest struct {
	Request                         // append to the request struct
	Params  InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`

	// ... alot more goes here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Response
type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabailities ServerCapabilitilies `json:"capabilities"`
	ServerInfo    ServerInfo           `json:"severInfo"`
}

type ServerCapabilitilies struct {
	TextDocumentSync int  `json:"textDocumentSync"`
	HoverProvider    bool `json:"hoverProvider"`
}
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializseResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabailities: ServerCapabilitilies{
				TextDocumentSync: 1, // sync by sending full content of document
				HoverProvider:    true,
			},
			ServerInfo: ServerInfo{
				Name:    "monkey-ls",
				Version: "0.0.0",
			},
		},
	}
}
