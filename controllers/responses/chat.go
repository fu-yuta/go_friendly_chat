package responses

type Chat struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type Chats struct {
	Chats []Chat `json:"chats"`
}
