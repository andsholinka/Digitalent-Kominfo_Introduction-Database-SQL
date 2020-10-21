package main

import (
	"fmt"

	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1-main/sql-orm/config"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1-main/sql-orm/database"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	database.InsertCustomer(database.CustomerORM{
		FirstName:    "Andrey",
		LastName:     "Sholinka",
		NpwpId:       "npwp123",
		Age:          22,
		CustomerType: "VVIP",
		Street:       "Kabupaten",
		City:         "Sleman",
		State:        "Indonesia",
		ZipCode:      "55293",
		PhoneNumber:  "08980704148",
		AccountORM: []database.AccountORM{
			{
				Balance:     19000,
				AccountType: "Premium",
			},
			{
				Balance:     3000,
				AccountType: "Deposito",
			},
		},
	}, db)
	// database.GetCustomer(db)
	//database.DeleteCustomer(1,db)
	//database.UpdateCustomer(database.CustomerORM{PhoneNumber: "0812314"},2,db)
	//
	//database.InsertAccount(database.AccountORM{
	//	Balance:         10000,
	//	AccountType:     "Deposit",
	//},2,db)

}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&database.CustomerORM{}, &database.AccountORM{})

	log.Println("db successfully connected")

	return db, nil
}
