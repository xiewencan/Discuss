package request

type RegisterRequest struct {
	Email     string `josn:"email"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	EmailCode string `json:"email_code"`
}
