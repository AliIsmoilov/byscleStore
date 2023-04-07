package models

type User struct {
	UserId     		string 	 `json:"user_id"`
	FirstName   	string 	`json:"first_name"`
	LastName    	string 	`json:"last_name"`
	Login			string	`json:"login"`
	Password		string	`json:"password"`
	Phone_number    string 	`json:"phone"`
	Created_at		string	`json:"created_at"`
	Updated_at		string	`json:"updated_at"`
}

type UserPrimaryKey struct {
	UserId	string	`json:"user_id"`
}

type CreateUser struct {
	FirstName   	string 	`json:"first_name"`
	LastName    	string 	`json:"last_name"`
	Login			string	`json:"login"`
	Password		string	`json:"password"`
	Phone_number    string 	`json:"phone"`
}

type UpdateUser struct {
	UserId     		string 	 `json:"user_id"`
	FirstName   	string 	`json:"first_name"`
	LastName    	string 	`json:"last_name"`
	Login			string	`json:"login"`
	Password		string	`json:"password"`
	Phone_number    string 	`json:"phone"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count  int      `json:"count"`
	Users []*User 	`json:"users"`
}
