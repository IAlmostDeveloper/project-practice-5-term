package services

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"server/src/dto"
	"server/src/storage/interfaces"
	"time"
)

type UserService struct {
	storage        interfaces.StorageProvider
	accessSecret   string
	AccessTokenTTL time.Duration
	redis          *redis.Client
}

func NewUserService(storage interfaces.StorageProvider, redis *redis.Client) *UserService {
	return &UserService{
		storage: storage,
		redis:   redis,
	}
}

func (s *UserService) AuthorizeUser(accessToken string) (string, error) {
	claims := jwt.MapClaims{}
	if _, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.accessSecret), nil
	}); err != nil {
		return "", err
	}

	result := s.redis.Get(context.Background(), accessToken)
	storedId := result.Val()
	tokenId := claims["user_id"].(string)
	if storedId != tokenId {
		return "", errors.New("stored user id and token user id do not match")
	}
	return claims["user_id"].(string), nil
}

func (s *UserService) AuthenticateUser(user *dto.User) (string, error) {
	user, err := s.storage.UserRepository().GetByLoginAndHashedPassword(user.Login, user.HashedPassword)
	if err != nil {
		return "", errInvalidUserData
	}
	s.AccessTokenTTL = time.Second * 3600
	return s.GenerateAndSaveToken(user)
}

func (s *UserService) RegisterUser(user *dto.User) error {
	storageUser, err := s.storage.UserRepository().GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if storageUser != nil {
		return errUserAlreadyExists
	}
	user.RegistrationDate = new(dto.TimeJson)
	if err = user.RegistrationDate.UnmarshalJSON([]byte(time.Now().Add(5).Format(dto.DateFormat))); err != nil{
		return err
	}
	if err := s.storage.UserRepository().Create(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) register(user *dto.User) error {
	return s.storage.UserRepository().Create(user)
}

func (s *UserService) GenerateAndSaveToken(user *dto.User) (string, error) {
	accessToken, err := s.createAccessToken(user)
	if err != nil {
		return "", err
	}
	if err := s.saveToken(user, accessToken); err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *UserService) GetUserById(userId string) (*dto.User, error) {
	return s.storage.UserRepository().GetById(userId)
}

func (s *UserService) GetAccessTokenTTL() time.Duration {
	return s.AccessTokenTTL
}

func (s *UserService) createAccessToken(user *dto.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.UserId
	atClaims["exp"] = time.Now().Add(s.AccessTokenTTL).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(s.accessSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *UserService) saveToken(user *dto.User, accessToken string) error {
	s.redis.Set(context.Background(), accessToken, user.UserId, s.AccessTokenTTL)
	res := s.redis.Save(context.Background())
	return res.Err()
}
