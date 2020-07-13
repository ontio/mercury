package packager

type MessageData struct {
	Data []byte `json:"data,omitempty"`
	Sign []byte `json:"sign,omitempty"`
}

type Envelope struct {
	Message *MessageData `json:"message,omitempty"`
	MsgType string       `json:"msgtype,omitempty"`
	FromDID string       `json:"fromdid,omitempty"`
	ToDID   string       `json:"todid,omitempty"`
}
