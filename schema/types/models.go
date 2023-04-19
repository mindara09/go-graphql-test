package types

type Products struct {
	Id          int    `json:"id"`
	NameProduct string `json:"name_product"`
	TypeProduct int    `json:"type_product"`
	CreatedAt   string `json:"created_at"`
}

type TypeProducts struct {
	Id        int    `json:"id"`
	NameType  string `json:"name_type"`
	CreatedAt string `json:"created_at"`
}

type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
