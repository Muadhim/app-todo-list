package seed

import (
	"log"

	"github.com/Muadhim/app-todo-list/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Username: "ucub victor",
		Email: "ucub@gmail.com",
		Password: "ucubPassword",
	},
	{
		Username: "abdol",
		Email: "abdol@gmail.com",
		Password: "abdolPassword",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}).Error // can add other model
	if err != nil {
		log.Fatalf("cannot drop table %v", err)
	}
	
	err = db.Debug().AutoMigrate(&models.User{}).Error // can add other model
	if err != nil {
		log.Fatalf("cannot migrate table %v", err)
	}

	// // example add foreignkey
	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		// other example
		// posts[i].AuthorID = users[i].ID

		// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}
}