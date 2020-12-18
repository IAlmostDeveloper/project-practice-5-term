package mysql_storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/src/dto"
)

type UserRepository struct {
	db *sqlx.DB
}

func (repo *UserRepository) Create(user *dto.User) error {
	insertStatement := "insert into `users`(`Email`, `Login`, `HashedPassword`, " +
		"`FirstName`, `LastName`, `BirthDate`, `RegistrationDate`, `IsRegisteredWithGoogle`, " +
		"`GoogleAccountData`, `AvatarPicture`) " +
		"values(:Email, :Login, :HashedPassword, " +
		":FirstName, :LastName, :BirthDate,:RegistrationDate,  " +
		":IsRegisteredWithGoogle, :GoogleAccountData, :AvatarPicture)"
	if _, err := repo.db.NamedExec(insertStatement, user); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetById(userId string) (*dto.User, error) {
	selectStatement := "SELECT * FROM `users` WHERE userid = ?"
	user := &dto.User{}
	if err := repo.db.Get(user, selectStatement, userId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByLogin(login string) (*dto.User, error) {
	selectStatement := "SELECT * FROM `users` WHERE login = ?"
	user := &dto.User{}
	if err := repo.db.Get(user, selectStatement, login); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByLoginAndHashedPassword(login string, hashedPassword string) (*dto.User, error){
	selectStatement := "SELECT * FROM `users` WHERE login = ? AND hashedpassword = ?"
	user := &dto.User{}
	if err := repo.db.Get(user, selectStatement, login, hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Update(user *dto.User) error {
	updateStatement := "update `users` set `email` = :email, `hashedpassword` = :hashed_password" +
		"`firstname` = :first_name, `lastname` = :last_name, `birthdate` = :birth_date," +
		" `registrationdate` = :registration_date, `avatarpicture` = :avatar_picture" +
		" where `userid` = :user_id"
		if _, err := repo.db.NamedExec(updateStatement, user); err != nil{
			return err
		}
		return nil
}

func (repo *UserRepository) RemoveById(userId string) error {
	deleteStatement := "DELETE FROM `users` WHERE userid = ?"
	if _, err := repo.db.Exec(deleteStatement, userId); err != nil {
		return err
	}
	return nil
}
