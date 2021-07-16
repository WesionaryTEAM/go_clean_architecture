package controllers

import (
	"clean-architecture/constants"
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/services"
	"clean-architecture/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service services.UserService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger lib.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	user, err := u.service.GetOneUser(lib.ParseUUID(paramID))

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, gin.H{"data": users})
}

// SaveUser saves the user
func (u UserController) SaveUser(c *gin.Context) {
	// user := models.User{
	// 	Age:   1,
	// 	Email: "dineshsilwal@gmail.com",
	// 	Name:  "Dinesh Silwal",
	// }

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.Create(user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// UpdateUser updates user
func (u UserController) UpdateUser(c *gin.Context) {
	paramID := lib.ParseUUID(c.Param("id"))

	user, err := u.service.GetOneUser(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.CustomBind(c.Request, &user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	metadata, _ := c.MustGet(constants.File).(lib.UploadedFiles)
	fmt.Println("---->")
	fmt.Println(metadata)
	user.ProfilePic = lib.SignedURL(metadata.GetFile("file").URL)

	if err := u.service.UpdateUser(user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes user
func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	if err := u.service.DeleteUser(lib.ParseUUID(paramID)); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
