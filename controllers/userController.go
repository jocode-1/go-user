package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jocode-1/go-user/models"
	"github.com/jocode-1/go-user/services"
)

type UserController struct {
	UserService services.UserService
}

func New(UserService services.UserService) UserController {
	return UserController{
		UserService: UserService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {

	userID := ctx.Param("id")
	user, err := uc.UserService.GetUser(&userID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {

	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {

	userroute := rg.Group("/API/V1/User")

	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get:name", uc.GetUser)
	userroute.GET("/getAll", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete", uc.DeleteUser)

}
