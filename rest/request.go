package rest

type Invitation struct {
	Type            string   `json:"@type"`
	Id              string   `json:"@id"`
	Lable           string   `json:"lable"`
	Did             string   `json:"did"`
	RecipientKeys   []string `json:"recipientKeys"`
	ServiceEndpoint string   `json:"serviceEndpoint"`
	routingKeys     string   `json:"routingKeys"`
}

type ConnectionRequest struct {
	Type            string   `json:"@type"`
	Id              string   `json:"@id"`
	Lable           string   `json:"lable"`
	Connection  Connection `json:"connection"`
}

type Connection struct {
	Did string `json:"did"`
}

type DIDDoc struct {
	Context []string `json:"@context"`
	Id string `json:"id"`
	PublicKey 
}

type PublicKey struct {
	Type            string   `json:"type"`
	Id              string   `json:"id"`
	Controller      string `json:"controller"`
	PublicKeyBase58 string `json:"publicKeyBase58"`
}