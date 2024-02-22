package model

type User struct {
	UserId    string `json:"userId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Work_auth string `json:"workAuth"`
}
