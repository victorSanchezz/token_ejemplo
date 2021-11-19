package estructura

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

var Hash = "$2a$10$C96TnfOQ56XQsnxHJTkji.XLVKr.rIerZIHxnfeKh5/RIMQvNp6Ve"
var User1 = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Name:     "jose",
	Lastname: "jose",
	Email:    "jose@gmail.com",
	Phone:    "1234567890",
}
