package user

import "time"

type User struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"`
	Fullname      string    `json:"fullname"`
	University    string    `json:"university"`
	Major         string    `json:"major"`
	CitizenID     string    `json:"citizen_id"`
	StudentID     string    `json:"student_id"`
	Password      string    `json:"password"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	DOB           string    `json:"dob"`
	Gender        string    `json:"gender"`
	Manual        bool      `json:"manual"`
	Secret        string    `json:"secret"`
	Source        string    `json:"source"`
	Socmed        string    `json:"socmed"`
	CitizenCard   string    `json:"citizen_card"`
	CitizenSelfie string    `json:"citizen_selfie"`
	StudentCard   string    `json:"student_card"`
	Merchant      bool      `json:"merchant"`
	TimeSignup    time.Time `json:"time_signup"`
	TimeSignin    time.Time `json:"time_signin"`
	Activated     bool      `json:"activated"`
	SocmedID      string    `json:"socmed_id"`
}
