package v1Handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"nguyenduy.com/golang-api/utils"
)

type GetUserByIdV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get list users is successful"})
}

func (u *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	var params GetUserByIdV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Get list users is successful"})
}

func (u *UserHandler) GetUsersByUuidV1(ctx *gin.Context) {
	idStr := ctx.Param("uuid")

	_, err := uuid.Parse(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid uuid"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Update users is successful"})
}

func (u *UserHandler) CreateUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create new user is successful"})
}

func (u *UserHandler) UpdateUserV1(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id must a number"})
		return
	}

	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": "Id must be positive"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Update users is successful"})
}

func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete user is successful"})
}
