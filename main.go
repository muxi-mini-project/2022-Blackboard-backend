package main

import (
	// "blackboard/config"
	"blackboard/model"
	"blackboard/router"
	"blackboard/services/flag_handle"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/spf13/viper"
)

// @title BlackBoard API
// @version 1.0.0
// @description 黑板API
// @termsOfService http://swagger.io/terrms/
// @contact.name Wishiforpeace
// @contact.email 1903180340@qq.com
// @host 122.112.236.36:8080
// @BasePath /api/v1
// @Schemes http

var err error

func main() {
	// err := config.Init("./conf/config.yaml", "")
	// if err != nil {
	// 	panic(err)
	// }
	// dbMap := viper.GetStringMapString("db")
	// dsn := fmt.Sprintf("%s:%s@/%s?parseTime=True", dbMap["username"], dbMap["password"], dbMap["name"])
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

func init() {
	port := flag.String("port", "8080", "本地监听的端口")
	platform := flag.String("platform", "gitee", "平台名称，支持gitee/github")
	token := flag.String("token", "5e3014aedeaaf3e2c9b8dc58dc03d959", "Gitee/Github 的用户授权码")
	owner := flag.String("owner", "Wishforpeace", "仓库所属空间地址(企业、组织或个人的地址path)")
	repo := flag.String("repo", "blackboard", "仓库路径(path)")
	path := flag.String("path", "", "文件的路径")
	branch := flag.String("branch", "master", "分支")
	flag.Parse()
	flag_handle.PORT = *port
	flag_handle.OWNER = *owner
	flag_handle.REPO = *repo
	flag_handle.PATH = *path
	flag_handle.TOKEN = *token
	flag_handle.PLATFORM = *platform
	flag_handle.BRANCH = *branch

	if flag_handle.TOKEN == "" {
		panic("token 必须！")
	}
}
