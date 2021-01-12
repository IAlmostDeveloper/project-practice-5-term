-- 00001_init_database.sql

-- +goose Up

CREATE TABLE IF NOT EXISTS Users
(
    UserId                 INT PRIMARY KEY AUTO_INCREMENT,
    Email                  VARCHAR(40) UNIQUE NOT NULL,
    Login                  VARCHAR(20) UNIQUE NOT NULL,
    HashedPassword         NVARCHAR(100),
    FirstName              NVARCHAR(30)       NOT NULL,
    LastName               NVARCHAR(30)       NOT NULL,
    BirthDate              DATETIME,
    RegistrationDate       DATETIME           NOT NULL,
    IsRegisteredWithGoogle NUMERIC            NOT NULL,
    GoogleAccountData      VARCHAR(500),
    AvatarPicture          VARCHAR(150)
);

CREATE TABLE IF NOT EXISTS UserPreferences
(
    Id     INT PRIMARY KEY AUTO_INCREMENT,
    Name   VARCHAR(100) NOT NULL,
    UserId INT         NOT NULL
);

ALTER TABLE UserPreferences
    ADD UNIQUE UserPreference (Name, UserId);

CREATE TABLE IF NOT EXISTS MeditationExercises
(
    ExerciseId   INT PRIMARY KEY AUTO_INCREMENT,
    Name         VARCHAR(100) NOT NULL,
    ExerciseTime INT          NOT NULL
);

CREATE TABLE IF NOT EXISTS FocusingExercises
(
    ExerciseId   INT PRIMARY KEY AUTO_INCREMENT,
    Name         VARCHAR(100) NOT NULL,
    ExerciseTime INT          NOT NULL
);

CREATE TABLE IF NOT EXISTS MeditationExercisesStarted
(
    Id           INT PRIMARY KEY AUTO_INCREMENT,
    ExerciseId   INT      NOT NULL,
    UserId       INT      NOT NULL,
    IsCompleted  BIT      NOT NULL,
    StartDate    DATETIME NOT NULL,
    CompleteDate DATETIME
);

CREATE TABLE IF NOT EXISTS FocusingExercisesStarted
(
    Id           INT PRIMARY KEY AUTO_INCREMENT,
    ExerciseId   INT      NOT NULL,
    UserId       INT      NOT NULL,
    IsCompleted  BIT      NOT NULL,
    StartDate    DATETIME NOT NULL,
    CompleteDate DATETIME
);

CREATE TABLE IF NOT EXISTS Achievements
(
    AchievementId INT PRIMARY KEY AUTO_INCREMENT,
    Name          VARCHAR(50)  NOT NULL,
    Description   VARCHAR(150) NOT NULL
);

CREATE TABLE IF NOT EXISTS AchievementsAchieved
(
    Id            INT PRIMARY KEY AUTO_INCREMENT,
    AchievementId INT      NOT NULL,
    UserId        INT      NOT NULL,
    AchieveDate   DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS Articles
(
    ArticleId  INT PRIMARY KEY AUTO_INCREMENT,
    Name       VARCHAR(100) NOT NULL,
    Content    LONGTEXT     NOT NULL,
    CreateDate DATETIME     NOT NULL
);

CREATE TABLE IF NOT EXISTS ArticleCategories
(
    Id        INT PRIMARY KEY AUTO_INCREMENT,
    ArticleId INT          NOT NULL,
    Name      VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS ArticleMedia
(
    MediaId   INT PRIMARY KEY AUTO_INCREMENT,
    Type      VARCHAR(20)  NOT NULL,
    ArticleId INT          NOT NULL,
    FTPLink   VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS ArticlesSaved
(
    Id        INT PRIMARY KEY AUTO_INCREMENT,
    ArticleId INT      NOT NULL,
    UserId    INT      NOT NULL,
    SaveDate  DATETIME NOT NULL
);

-- +goose Down
DROP DATABASE ProjectServer;
