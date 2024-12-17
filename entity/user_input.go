package entity

type (
	LoginEmailInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	DataUserInput struct {
		ID       string
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Roles    string `json:"roles"`
	}
)
