package user

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    int(user.ID),
		Name:  user.Name,
		Phone: user.Phone,
		Token: token,
	}
	return formatter
}
