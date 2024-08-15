package service

import "github.com/lauro-ss/api_with_goe/internal/data"

type userRepository struct {
	db *data.Database
}

func NewUserRepository(db *data.Database) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) List() (users []data.User, err error) {
	db := ur.db
	_, err = db.Select(db.User).Scan(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) Get(id uint) (*data.User, error) {
	db := ur.db
	var user data.User
	_, err := db.Select(db.User).Where(db.Equals(&db.User.Id, id)).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Create(user data.User) (uint, error) {
	db := ur.db
	_, err := db.Insert(db.User).Value(&user)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (ur *userRepository) Update(user data.User) (uint, error) {
	db := ur.db
	_, err := db.Update(db.User).Where(db.Equals(&db.User.Id, user.Id)).Value(&user)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (ur *userRepository) Delete(id uint) (bool, error) {
	db := ur.db
	_, err := db.Delete(db.User).Where(db.Equals(&db.User.Id, id))
	if err != nil {
		return false, err
	}
	return true, nil
}
