package entitiy

type Authorization struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
