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
	insertStatement := "insert into Users (`Email`, `Login`, `HashedPassword`, " +
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
	selectStatement := "SELECT * FROM Users WHERE UserId = ?"
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
	selectStatement := "SELECT * FROM Users WHERE Login = ?"
	user := &dto.User{}
	if err := repo.db.Get(user, selectStatement, login); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*dto.User, error) {
	selectStatement := "SELECT * FROM Users WHERE Email = ?"
	user := &dto.User{}
	if err := repo.db.Get(user, selectStatement, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByLoginAndHashedPassword(login string, hashedPassword string) (*dto.User, error) {
	selectStatement := "SELECT * FROM Users WHERE Login = ? AND HashedPassword = ?"
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
	updateStatement := "update Users set `email` = :email, `hashedpassword` = :hashed_password" +
		"`firstname` = :first_name, `lastname` = :last_name, `birthdate` = :birth_date," +
		" `registrationdate` = :registration_date, `avatarpicture` = :avatar_picture" +
		" where `userid` = :user_id"
	if _, err := repo.db.NamedExec(updateStatement, user); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) RemoveById(userId string) error {
	deleteStatement := "DELETE FROM Users WHERE UserId = ?"
	if _, err := repo.db.Exec(deleteStatement, userId); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserArticles(userId string) ([]*dto.Article, error) {
	selectStatement := "SELECT a.* FROM ArticlesSaved as s inner join Articles as a on a.ArticleId=s.Articleid where s.UserId=?"
	articles := &[]*dto.Article{}
	err := repo.db.Select(articles, selectStatement)
	if err != nil {
		return nil, err
	}
	return *articles, err
}

func (repo *UserRepository) GetUserFocusingExercises(userId string) ([]*dto.FocusingExercise, error) {
	selectStatement := "SELECT e.* FROM FocusingExercisesStarted as s inner join FocusingExercises as e on e.ExerciseId=s.ExerciseId where s.UserId=?"
	exercises := &[]*dto.FocusingExercise{}
	err := repo.db.Select(exercises, selectStatement, userId)
	if err != nil {
		return nil, err
	}
	return *exercises, err
}

func (repo *UserRepository) GetUserMeditationExercises(userId string) ([]*dto.MeditationExercise, error) {
	selectStatement := "SELECT e.* FROM MeditationExercisesStarted as s inner join MeditationExercises as e on e.ExerciseId=s.ExerciseId where s.UserId=?"
	exercises := &[]*dto.MeditationExercise{}
	err := repo.db.Select(exercises, selectStatement, userId)
	if err != nil {
		return nil, err
	}
	return *exercises, err
}

func (repo *UserRepository) GetUserAchievements(userId string) ([]*dto.Achievement, error) {
	selectStatement := "SELECT a.* FROM AchievementsAchieved as s inner join Achievements as a on a.AchievementId=s.AchievementId where s.UserId=?"
	exercises := &[]*dto.Achievement{}
	err := repo.db.Select(exercises, selectStatement, userId)
	if err != nil {
		return nil, err
	}
	return *exercises, err
}
