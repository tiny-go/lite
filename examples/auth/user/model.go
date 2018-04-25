package user

// User is a user model.
type User struct {
	Email    string `json:"email,omitempty" xml:"Email,omitempty"`
	Password string `json:"password,omitempty" xml:"Password,omitempty"`
	Token    string `json:"token,omitempty" xml:"Token,omitempty"`
	Refresh  string `json:"refresh,omitempty" xml:"Refresh,omitempty"`
}

func (u *User) Login() error {
	// TODO: add some logic
	return nil
}

func (u *User) RefreshToken() error {
	// TODO: add some logic
	return nil
}
