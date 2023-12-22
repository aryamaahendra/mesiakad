package account

type Repository interface {
	GetAllUsers() ([]Account, error)
	GetUserByUsername(username string) (Account, error)
	CreateUser(acc *Account) error
	UpdateUser(username string, acc *Account) error
	DeleteUser(username string) error
	CreateOrUpdateProfile(acc *Account) error
	GetProfileByUserID(ID uint64) (Account, error)
}
