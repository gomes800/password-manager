package model

type Credential struct {
	ID          int    `json: "id"`
	UserId      int    `json: "user_id"`
	ServiceName string `json: "service_name"`
	Username    string `json: "username"`
	Ciphertext  []byte `json: "ciphertext"`
	Nonce       []byte `json: "nonce"`
	Salt        []byte `json: "salt"`
}
