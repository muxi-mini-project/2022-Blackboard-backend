package user

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/service/user"
	"encoding/base64"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//@Summary 登录
//@Tags user
//@Description 一站式登录
//@Accep applica/json
//@Produce applic/json
//@Param object body model.User
