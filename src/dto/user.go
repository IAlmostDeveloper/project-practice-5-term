package dto

type User struct {
	UserId                 string    `json:"user_id" db:"UserId"`
	Email                  string    `json:"email" db:"Email"`
	Login                  string    `json:"login" db:"Login"`
	HashedPassword         string    `json:"password" db:"HashedPassword"`
	FirstName              string    `json:"first_name" db:"FirstName"`
	LastName               string    `json:"last_name" db:"LastName"`
	BirthDate              *TimeJson `json:"birth_date" db:"BirthDate"`
	RegistrationDate       *TimeJson `json:"registration_date" db:"RegistrationDate"`
	IsRegisteredWithGoogle bool      `json:"is_registered_with_google" db:"IsRegisteredWithGoogle"`
	AccountData            string    `json:"google_account_data" db:"GoogleAccountData"`
	AvatarPicture          string    `json:"avatar_picture" db:"AvatarPicture"`
}
