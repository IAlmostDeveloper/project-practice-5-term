-- 00001_init_database.sql

-- +goose Up
CREATE DATABASE IF NOT EXISTS ProjectServer;
USE ProjectServer;

CREATE TABLE IF NOT EXISTS Users (
    UserId INT PRIMARY KEY AUTO_INCREMENT,
    Email VARCHAR(40) UNIQUE NOT NULL,
    Login VARCHAR(20) UNIQUE NOT NULL,
    HashedPassword NVARCHAR(32),
    FirstName NVARCHAR(30) NOT NULL,
    LastName NVARCHAR(30) NOT NULL,
    BirthDate DATE,
    RegistrationDate DATETIME NOT NULL,
    IsRegisteredWithGoogle BIT NOT NULL,
    GoogleAccountData VARCHAR(500),
    AvatarPicture VARCHAR(150)
);

CREATE TABLE IF NOT EXISTS MeditationExercises(
    ExerciseId INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100) NOT NULL,
    ExerciseTime INT NOT NULL
);

CREATE TABLE IF NOT EXISTS FocusingExercises(
    ExerciseId INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100) NOT NULL,
    ExerciseTime INT NOT NULL
);

CREATE TABLE IF NOT EXISTS MeditationExercisesStarted(
    ExerciseId INT PRIMARY KEY AUTO_INCREMENT,
    UserId INT NOT NULL,
    IsCompleted BIT NOT NULL,
    StartDate DATETIME NOT NULL,
    CompleteDate DATETIME
);

CREATE TABLE IF NOT EXISTS FocusingExercisesStarted(
   ExerciseId INT PRIMARY KEY AUTO_INCREMENT,
   UserId INT NOT NULL,
   IsCompleted BIT NOT NULL,
   StartDate DATETIME NOT NULL,
   CompleteDate DATETIME
);

CREATE TABLE IF NOT EXISTS Achievements(
    AchievementId INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50) NOT NULL,
    Description VARCHAR(150) NOT NULL
);

CREATE TABLE IF NOT EXISTS AchievementsAchieved(
    AchievementId INT PRIMARY KEY AUTO_INCREMENT,
    UserId INT NOT NULL,
    AchieveDate DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS Articles(
    ArticleId INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100) NOT NULL,
    Content LONGTEXT NOT NULL,
    CreateDate DATETIME NOT NULL

);

CREATE TABLE IF NOT EXISTS ArticleMedia(
    MediaId INT PRIMARY KEY AUTO_INCREMENT,
    Type VARCHAR(20) NOT NULL,
    ArticleId INT NOT NULL,
    FTPLink VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS ArticlesSaved(
    ArticleId INT PRIMARY KEY AUTO_INCREMENT,
    UserId INT NOT NULL,
    SaveDate DATETIME NOT NULL
);


-- +goose Down
DROP DATABASE ProjectServer;
