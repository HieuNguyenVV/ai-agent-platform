package controller

import (
	appErrors "agent/ai-agent-platform/internal/errors"
	"agent/ai-agent-platform/internal/server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}

// HandleGetUserByID is a handler for the GET /users/{id} endpoint
// It returns a user by ID
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 500 {object} AppError
// @Router /users/{id} [get]
// @Security AccessToken
func (c *UserController) HandleGetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handleError(ctx, appErrors.ErrBadRequest.WrapString("id is required"))
		return
	}
	user, err := c.userService.GetUserByID(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}