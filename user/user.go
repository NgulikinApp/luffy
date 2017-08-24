package user

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	Source    string `json:"source"`
	Activated bool   `json:"activated"`
}
