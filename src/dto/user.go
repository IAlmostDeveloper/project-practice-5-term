package dto

type User struct {
	UserId                 string            `json:"user_id" db:"user_id"`
	Email                  string            `json:"email" db:"email"`
	Login                  string            `json:"login" db:"login"`
	HashedPassword         string            `json:"hashed_password" db:"hashed_password"`
	FirstName              string            `json:"first_name" db:"first_name"`
	LastName               string            `json:"last_name" db:"last_name"`
	BirthDate              *TimeJson         `json:"birth_date" db:"birth_date"`
	Registration           *TimeJson         `json:"registration_date" db:"registration_date"`
	IsRegisteredWithGoogle bool              `json:"is_registered_with_google"`
	AccountData            GoogleAccountData `json:"google_account_data" db:"google_account_data"`
	AvatarPicture          string            `json:"avatar_picture" db:"avatar_picture"`
}
