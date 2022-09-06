package db

import (
	"log"
	"math/rand"
	"time"

	faker "github.com/eykn2z/go_faker"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/model"
)

var (
	db  *gorm.DB
	err error
)

func Get() *gorm.DB {
	return db
}

func Connect() *gorm.DB {
	db, err = gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Init() *gorm.DB {
	db := Connect()

	//data delete
	dropTable(db)

	//table migration(creation)
	db.AutoMigrate(&model.Group{}, &model.User{}, &model.SuperUser{}, &model.Message{})

	//insert test data
	insertTestData(db)

	return db
}

func dropTable(db *gorm.DB) {
	db.Migrator().DropTable(&model.Group{}, &model.User{}, &model.SuperUser{}, &model.Message{})
}

func insertTestData(db *gorm.DB) {
	//insert groups
	groups := []model.Group{{Name: "A"}, {Name: "B"}, {Name: "C"}}
	groupsResult := db.Create(&groups)
	if groupsResult.Error != nil {
		log.Fatal(groupsResult.Error)
	}

	//insert users
	people := faker.GetPersons(20)
	users := []model.User{}
	for _, person := range people {
		user := model.User{Name: person.Name, Group: groups[randomIndex(3)]}
		users = append(users, user)
	}
	usersResult := db.Create(&users)
	if usersResult.Error != nil {
		log.Fatal(usersResult.Error)
	}
}

func randomIndex(arrayLength int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(arrayLength)
}
