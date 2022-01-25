package main

import (
	"blackboard/model"
	"blackboard/router"
	"fmt"

	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/spf13/viper"
)

// @title BlackBoard API
// @version 1.0.0
// @description 黑££板API
// @termsOfService http://swagger.io/terrms/
// @contact.name Wishiforpeace
// @contact.email 1903180340@qq.com
// @host localhost
// @BasePath /api/v1
// @Schemes http

var err error

func main() {

	// //init db
	// //获得“db”的map
	// dbMap := viper.GetStringMapString("db")
	// dbConfig := fmt.Sprintf("%s:%s@%s?parseTime=Ture",dbMap["username"],dbMap["password",dbMap["name"]])
	// model.DB, err = gorm.Open("mysql",dbConfig)
	// if err !=nil{
	// 	panic(err)
	// }

	// go concurent
	//本地连接 init db
	dsn := "root:root&1234@tcp(127.0.0.1:3306)/blackboard?charset=utf8mb4&parseTime=True&loc=Local"
	model.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	model.Migrate(model.DB)
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
	defer model.DB.Close()
}
