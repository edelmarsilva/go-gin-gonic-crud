package request

type UserRequest struct {
	Email string `'json:"email"`
	Senha string `'json:"senha"`
	Nome  string `'json:"nome"`
	Idade int8   `'json:"idade"`
}
