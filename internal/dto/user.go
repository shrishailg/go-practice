package dto

type (
	User struct {
		ID           string `json:"id"`
		Username     string `json:"user_name"`
		UserLastName string `json:"last_name"`
		Age          int    `json:"age"`
	}
)
