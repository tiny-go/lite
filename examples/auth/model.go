package auth

type Model struct {
	Email    string `json:"omitempty" xml:"omitempty"`
	Password string `json:"omitempty" xml:"omitempty"`
	Token    string `json:"omitempty" xml:"omitempty"`
	Refresh  string `json:"omitempty" xml:"omitempty"`
}
