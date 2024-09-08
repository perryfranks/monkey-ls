package lsp

type Request struct {
	RPC string `json:"jsonrpc"` // always just says 2.0
	// TODO: support strings here as well
	ID     int    `json:"id"` // neovim always sends int. protocol can provide strings. ignore for now
	Method string `json:"method"`

	// params will be in the sub messages

}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id, omitempty"`

	// result
	// error
}

type Notification struct {
	PRC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
