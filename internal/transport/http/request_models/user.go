package request_models

type User struct {
	FirstName   string `validate:"required,min=3,max=32" json:"first_name"`
	LastName    string `validate:"required,min=3,max=32" json:"last_name"`
	permissions []string
	Avatar      string
	Email       string `validate:"required,email,min=6,max=32"`
	Password    string `validate:"required,min=6,max=32"`
	BirthDate   string `validate:"required" json:"birth_date"`
}
