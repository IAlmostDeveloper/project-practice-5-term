package interfaces

type StorageProvider interface {
	UserRepository() UserRepositoryProvider
	ExerciseRepository() ExerciseRepositoryProvider
	ArticleRepository() ArticleRepositoryProvider
	AchievementRepository() AchievementRepositoryProvider
	QuotesRepository() QuotesRepositoryProvider
}