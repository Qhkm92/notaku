package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// func Login(c *gin.Context) {
// 	fmt.Println("success")
// 	c.JSONP(http.StatusOK, gin.H{
// 		"message": "login",
// 	})
// }

func Signup(c *gin.Context) {
	fmt.Println("success")
	c.JSONP(http.StatusOK, gin.H{
		"message": "signup",
	})
}

func Logout(c *gin.Context) {
	fmt.Println("success")
	c.JSONP(http.StatusOK, gin.H{
		"message": "logout",
	})
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json: "password"`
}

type JWT struct {
	Token string `json: "token"`
}

// func Signup(c *gin.Context) {
// 	// signup
// 	// handle email and password
// 	// firstname lastname
// 	// signup
// 	// auto login
// 	session := sessions.Default(c)
// 	if session.Get("hello") != "world" {
// 		session.Set("hello", "world")
// 		session.Save()
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "route exixst",
// 		"hello":   session.Get("hello"),
// 	})
// }

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ExampleData struct {
	email    string
	password string
}

func Login(c *gin.Context) {

	var qaiyyum ExampleData
	qaiyyum.email = "qaiyyum"
	qaiyyum.password = "test"

	fmt.Println(qaiyyum)
	// accept email
	var requestData LoginRequest

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "testing bad erro",
		})
		fmt.Println(err)
		// need to abort
		c.Abort()
		return
	}

	if requestData.Email == qaiyyum.email && requestData.Password == qaiyyum.password {
		fmt.Println("password is right")
		// if email & pw is valid
		// create session
		session := sessions.Default(c)
		v := session.Get("count")
		x := session.Get("mysession")
		fmt.Println(v)
		fmt.Println(x)
		session.Set("count", "world")
		session.Save()
		c.JSON(200, gin.H{
			"count": session.Get("count"),
		})

		// return cookies
	}

	// validate empty
	// if requestData.Email == http.NoBody || requestData.Password == http.NoBody {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "email or password not provided",
	// 	})
	// }

	//fmt.Println(requestData.Email)
	// fmt.Println(requestData.Password)
	// c.Get("password")
	// if(email && password !== nill) {
	// 	log bla2
	// }

	// / accept password

	// check email and password in DB exist
	// if exist generate token
	// pass token to client
	c.JSON(http.StatusOK, gin.H{
		"message": requestData,
	})
}
