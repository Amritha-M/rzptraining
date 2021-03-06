
import (
	Config "first-api"
	"fmt"
)

package main

import (
"day4-5/Config"
"day4-5/Models"
"day4-5/Routes"
"fmt"
"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Transaction{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}