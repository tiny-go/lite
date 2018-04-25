package user

// UserModel is a user model.
type UserModel struct {
	Email    string `json:"email,omitempty" xml:"Email,omitempty"`
	Password string `json:"password,omitempty" xml:"Password,omitempty"`
	Token    string `json:"token,omitempty" xml:"Token,omitempty"`
	Refresh  string `json:"refresh,omitempty" xml:"Refresh,omitempty"`
}
