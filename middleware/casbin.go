package middleware

import (
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Casbin() gin.HandlerFunc {
	ErrorMessage := uility.ErrorMessage{}
	return func(c *gin.Context) {
		e := uility.E
		RuleId := c.MustGet("RuleId")
		obj := c.FullPath()
		act := c.Request.Method
		fmt.Println(RuleId, obj, act)
		ok, err := e.Enforce(RuleId, obj, act)
		if err != nil {
			// 处理err
			log.Println(e)
			return
		}
		fmt.Println(ok)
		if !ok {
			// 拒绝请求，抛出异常
			log.Println("不通过")
			c.Abort()
			ErrorMessage.ErrorDescription = "权限验证不通过"
			panic(ErrorMessage)

		}
		log.Println("权限通过")
		c.Next()
	}
}
