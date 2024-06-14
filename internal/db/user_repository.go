package database

func (ctx *DbContext) GetUsers() (*[]User, error) {
	var users []User
	if err := ctx.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (ctx *DbContext) GetUserByID(id uint) (*User, error) {
	var user User
	if err := ctx.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ctx *DbContext) CreateUser(user *User) error {
	return ctx.Db.Create(user).Error
}

func (ctx *DbContext) DeleteUser(id uint) error {
	return ctx.Db.Delete(&User{}, id).Error
}
