package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(ctx *gin.Context) {
	user := new(User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := AddUser(user); err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func signIn(ctx *gin.Context) {
	user := new(User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	user, err := Authenticate(user.Username, user.Password)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Name or Password Error"})
	}
	//for _, u := range storage.Users {
	//	if u.Username == user.Username && u.Password == user.Password {
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"msg": "Signed in successfully.",
	//			"jwt": "123456789",
	//		})
	//		return
	//	}
	//}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}