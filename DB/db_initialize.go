package DB

import (
	"fmt"
	"log"

	"DB_GORM/models"
	"DB_GORM/utils"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
)

var DB *gorm.DB

var osExit = os.Exit

func Initialize() {
	loadYMLConfig()

	utils.InitLogger()

	username := viper.GetString("prod.username")
	password := viper.GetString("prod.password")
	dbHost := viper.GetString("prod.db_host")
	dbPort := viper.GetInt("prod.db_port")
	dbName := viper.GetString("prod.db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, dbHost, dbPort, dbName)

	utils.InfoLog.Println("Connecting to DB with DSN:", dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Job{}, &models.Recruiter{}, &models.Application{}, &models.Interview{})
	if err != nil {
		utils.ErrorLog.Printf("AutoMigrate failed: %v", err)
	}

	utils.InfoLog.Println("Database connected successfully")
}

func loadYMLConfig() {
	viper.AddConfigPath(".")
viper.SetConfigName("config")
viper.SetConfigType("yml")
if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("read config: %v", err)
}


}
