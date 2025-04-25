package models

type CreateSmtpEntryRequest struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateSmtpEntryRequest struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateJobRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateFileRequest struct {
	Name string `json:"name"`
}

type UserAuthenticationRequest struct {
	Code string `json:"code"`
}

type CreateRequestBodyRequest struct {
	Name string `json:"name"`
	Promo_Text string `json:"promo_text"`
}

