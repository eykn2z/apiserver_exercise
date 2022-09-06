package controller

import (
	"log"

	"server/db"
	"server/model"
)

type UserController struct{}

func (con *UserController) Create(name string) (model.User, error) {
	db := db.Get()
	var user model.User
	// user := model.User{Name: name}
	// result := db.Create(&user)
	result := db.FirstOrCreate(&user, model.User{Name: name})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	db.First(&user)
	return user, nil
}

func (con *UserController) GetByID(id int) (model.User, error) {
	var user model.User
	db := db.Get()
	result := db.First(&user, id)
	if result.RowsAffected != 0 {
		return user, nil
	}
	return user, result.Error
}

func (con *UserController) Update(user model.User, name string) (model.User, error) {
	//TODO: input: User, gormフィールド以外のフィールドリストを取得して比較して相違があるものをupdate
	db := db.Get()
	user.Name = name
	db.Save(&user)
	return user, nil
}

func (con *UserController) Delete(user model.User) error {
	db := db.Get()
	db.Delete(&user)
	return nil
}

type UsersController struct {
}

func (con *UsersController) Read() ([]model.User, error) {
	var users []model.User
	db := db.Get()
	db.Find(&users)
	return users, nil
}
