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

// login user
func (a authController) Login(c *gin.Context) {

	var loginModel models.LoginModel
	//get data off request body
	if err := c.ShouldBindJSON(&loginModel); err != nil {
		utils.CustomRepsonseWriter(c, http.StatusBadRequest, nil, "Error binding the data")
		return
	}

	//Check if user exist or not - then only we give access to login
	var user models.User
	userIns := a.db.Where("email = ?", loginModel.Email).First(&user)
	if userIns.RowsAffected == 0 {
		utils.CustomRepsonseWriter(c, http.StatusNotFound, nil, "User not found")
		return
	}

	//check if the user is verified or not
	if !user.IsActive {
		utils.CustomRepsonseWriter(c, http.StatusUnauthorized, nil, "Please verify your email")
	}

	//generate a basic auth token
	token := utils.GenerateBasicAuthToken(user.Email)

	res := map[string]any{
		"token": token,
	}

	//validate the password of existing user
	if user.IsActive {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password))
		if err != nil {
			utils.CustomRepsonseWriter(c, http.StatusUnauthorized, nil, "Incorrect password")
			return
		}

		utils.CustomRepsonseWriter(c, http.StatusFound, res, "Login successful")
	}
}

func (a authController) VerifyUser(c *gin.Context) {

	//get the token from the request url
	email := c.Param("email")

	//check if the user already exists
	var user models.User
	userIns := a.db.Where("email =?", email).First(&user)
	if userIns.RowsAffected == 0 {
		utils.CustomRepsonseWriter(c, http.StatusNotFound, nil, "User not found")
		return
	}
	//check if the user is already veridied/active
	if user.IsActive {
		utils.CustomRepsonseWriter(c, http.StatusConflict, nil, "User is already verified")
		return
	}
	//change the status to active
	user.IsActive = true
	if err := a.db.Save(&user).Error; err != nil {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, "error verifying user")
		return
	}

	utils.CustomRepsonseWriter(c, http.StatusOK, user, "user verified, status changed to active")

}
