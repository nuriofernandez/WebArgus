package add

type Response struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}
