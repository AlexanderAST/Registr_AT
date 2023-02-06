package registration

type User struct {
	Id           int    `json:"-" db:"id"`
	UserName     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	ReferralCode string `json:"referralcode" binding:"required"`
}
