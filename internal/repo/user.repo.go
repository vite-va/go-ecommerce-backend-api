package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // ur: user repo
// func (ur *UserRepo) GetInfoUser() string {
// 	return "Tipjs"
// }

// INTERFACE_VERSION
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	// panic("unimplemented")
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
