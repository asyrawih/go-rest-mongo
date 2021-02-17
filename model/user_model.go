package model

type CreateUserRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type GetUserResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type FindUserResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
