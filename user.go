package registration

type User struct {
	Id             int    `json:"-" db:"id"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	EnteredRefCode string `json:"enteredrefcode" binding:"required"`
}
