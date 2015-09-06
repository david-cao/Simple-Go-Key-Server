package types

type Get struct {
	FacebookID string
}

type User struct {
	PublicKey  string
	FacebookID string
}

type ErrorResponse struct {
	Error            bool
	ErrorDescription string
}

type GetPublicKeyResponse struct {
	*ErrorResponse
	PublicKey string
}
