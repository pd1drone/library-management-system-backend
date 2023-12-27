package models

type LoginResponsedb struct {
	ID       int64  `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success  bool             `json:"Success"`
	Message  string           `json:"Message"`
	Response *LoginResponsedb `json:"Response"`
}
