package models

type AddUser struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type DelUser struct {
	Username string `json:"username"`
}

type Resp struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}