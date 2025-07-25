package ft

type ExtensionCommand struct {
	Action  string                 `json:"action"`  // Specific action to perform (e.g., "sign", "connect", "getAccounts")
	Payload map[string]interface{} `json:"payload"` // Data needed by the extension to execute the command
}

type BasicResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}