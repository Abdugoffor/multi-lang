package config

import (
	"log"
	language_model "project/module/admin/language/model"
	post_model "project/module/admin/post/model"
)

func RunMigrations() {

	err := DB.AutoMigrate(
		&language_model.Language{},
		&post_model.Post{},
	)

	if err != nil {
		log.Fatal("❌ Migration failed: ", err)
	}

	log.Println("✅ Database migrated successfully")
}
