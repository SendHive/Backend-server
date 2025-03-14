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
}

type CreateUserRequest struct {
	Name string `json:"name"`
}
