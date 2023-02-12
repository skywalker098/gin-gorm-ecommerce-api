package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/net-http/database"
	"github.com/net-http/models"
	"github.com/net-http/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authController struct {
	db *gorm.DB
}

func NewAuthController() *authController {
	return &authController{
		db: database.Db,
	}
}

// Sign up the user
func (a *authController) Signup(c *gin.Context) {

	var authModel models.AuthModel

	//get data off request body
	if err := c.ShouldBindJSON(&authModel); err != nil {
		utils.CustomRepsonseWriter(c, http.StatusBadRequest, nil, "Error binding the data")
		return
	}

	//Check if user already exists
	var user models.User
	userIns := a.db.Where("email = ?", authModel.Email).First(&user)
	if userIns.RowsAffected != 0 {
		utils.CustomRepsonseWriter(c, http.StatusConflict, nil, "User already exists")
		return
	}

	//password validation
	if len(authModel.Password) < 6 {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, "Password is short")
		return
	}

	//Encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authModel.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	user.Password = string(hashedPassword)
	user.Email = authModel.Email
	user.FirstName = authModel.FirstName
	user.LastName = authModel.LastName

	//sending verification email to user email
	log.Println("sending verification email to: ", authModel.Email)

	//create the user
	if err := a.db.Create(&user).Error; err != nil {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}
	utils.CustomRepsonseWriter(c, http.StatusCreated, user, "User created")

}
