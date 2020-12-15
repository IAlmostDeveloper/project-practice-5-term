package services

import (
	"crypto/sha256"
	"encoding/base64"
)

type PasswordService struct{

}

func NewPasswordService() *PasswordService{
	return &PasswordService{}
}

func (service *PasswordService) EncodePassword (password string) string {
	sha := sha256.New()
	sha.Write([]byte(password))
	hash := base64.URLEncoding.EncodeToString(sha.Sum(nil))
	return hash
}
