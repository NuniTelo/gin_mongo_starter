package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {

		var code = 0
		token := c.Query("token")

		if token == "" {
			fmt.Println("lol error")
			code = 12

		} else {

			//here validate the token

			//_, err := util.ParseToken(token)  //we need to parse the token here
			//if err != nil {
			//	switch err.(*jwt.ValidationError).Errors {
			//	case jwt.ValidationErrorExpired:
			//		code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT  //if false return false
			//	default:
			//		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL //chck this shit out
			//	}
			//}
			code = 11

		}

		if code != 11 {  //if success is not

		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Unauthorized!",
		})
				c.Abort()  //abort the request
				return
		}

		c.Next()   //if succcess then go
	}
}
