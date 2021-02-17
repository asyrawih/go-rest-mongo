package model
// User request
// @todo Make Password Must Encrypt Before Persisting On Database
type CreateUserRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
// Create Response For User Entity
type CreateUserResponse struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
//Get User Response For User
type GetUserResponse struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
// Find One User
type FindUserResponse struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
