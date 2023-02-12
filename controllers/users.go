package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/net-http/database"
	"github.com/net-http/models"
	"github.com/net-http/utils"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController() *userController {
	return &userController{
		db: database.Db,
	}
}

// Get all Users
func (t *userController) GetAllUsers(c *gin.Context) {
	var users []models.User
	t.db.Find(&users)

	utils.CustomRepsonseWriter(c, http.StatusOK, users, "users found")

}

// Get one user
func (t *userController) GetOneUser(c *gin.Context) {
	//get id off...url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.CustomRepsonseWriter(c, http.StatusBadRequest, nil, "Error")
		return
	}

	var user models.User
	if err := t.db.First(&user, id).Error; err != nil {
		utils.CustomRepsonseWriter(c, http.StatusNotFound, nil, "User not found")
		return
	}

	utils.CustomRepsonseWriter(c, http.StatusFound, user, "user found")
}

// Create a user
func (t *userController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		utils.CustomRepsonseWriter(c, http.StatusBadRequest, nil, "Error binding user")
		return
	}
	if err := t.db.Create(&user).Error; err != nil {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	utils.CustomRepsonseWriter(c, http.StatusCreated, user, "user created")
}

// Delete a user
func (t *userController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.CustomRepsonseWriter(c, http.StatusNotFound, nil, "User not found")
		return
	}
	if err := t.db.Delete(&models.User{}, id).Error; err != nil {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, "Error deleting user")
		return
	}

	utils.CustomRepsonseWriter(c, http.StatusOK, nil, "user deleted successfully")
}

// update user
func (t *userController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.CustomRepsonseWriter(c, http.StatusNotFound, nil, "User not found")
		return
	}
	//get data off request body
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		utils.CustomRepsonseWriter(c, http.StatusBadRequest, nil, "Error binding user")
		return
	}
	if err := t.db.Model(&models.User{}).Where("id =?", id).Updates(user).Error; err != nil {
		utils.CustomRepsonseWriter(c, http.StatusInternalServerError, nil, "Error updating user")
		return
	}

	utils.CustomRepsonseWriter(c, http.StatusOK, user, "user Updated Successfully")
}
