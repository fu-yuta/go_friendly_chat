package requests

type Chat struct {
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type UpdateChat struct {
	Message string `json:"message"`
}
