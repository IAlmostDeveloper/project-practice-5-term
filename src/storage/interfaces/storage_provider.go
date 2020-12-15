package interfaces

type StorageProvider interface {
	UserRepository() UserRepositoryProvider
}