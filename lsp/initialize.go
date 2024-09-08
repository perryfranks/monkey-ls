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
