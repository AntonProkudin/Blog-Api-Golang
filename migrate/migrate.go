package migrate

import (
	"fmt"
	"log"
	"test/initializers"
	"test/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comm{})
	if err != nil {
		return
	}
	fmt.Println("👍 Migration complete")
}
