package controller

import (
	"fmt"
	"net/http"

	inmemory "github.com/foliveiracamara/bents-api/adapter/driven/in_memory"
	"github.com/foliveiracamara/bents-api/adapter/driver/lambda/model/request"
	"github.com/foliveiracamara/bents-api/adapter/driver/lambda/model/response"
	"github.com/foliveiracamara/bents-api/application/entity"
	port "github.com/foliveiracamara/bents-api/application/port/driver"
	"github.com/foliveiracamara/bents-api/application/service"
	"github.com/foliveiracamara/bents-api/configuration/apperr"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	UserService port.UserService
	Validate    *request.CustomValidator
	Model       *response.UserResponse
}

func newUserController(svc port.UserService, cv *request.CustomValidator) *UserController {
	return &UserController{
		UserService: svc,
		Validate:    cv,
	}
}

func InitUserController(ctx echo.Context) (controller *UserController) {
	userRepo := inmemory.NewInMemoryUserAdapter()
	userSvc := service.NewUserService(userRepo)
	cv := request.NewCustomValidator()
	controller = newUserController(userSvc, cv)

	return controller
}

func (uc *UserController) CreateUser(ctx echo.Context) error {
	var user request.UserRequest
	err := ctx.Bind(&user)
	if err != nil {
		appErr := apperr.NewBadRequestError("Error binding request.")
		return ctx.JSON(appErr.Code, appErr)
	}

	if err = uc.Validate.Validate(user); err != nil {
		log.Error().
			Str("journey", "userController.CreateUser").
			Msgf(err.Error())

		// appErr := validation.ValidateUserError(err)
		appErr := apperr.NewBadRequestError("Some fields are incorrect.")
		return ctx.JSON(appErr.Code, appErr)
	}

	userDomain := &entity.User{
		UUID:       uuid.New().String(),
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		Age:        user.Age,
		Favorites:  []string{},
		FirstLogin: true,
	}

	svc, appErr := uc.UserService.CreateUser(userDomain)
	if appErr != nil {
		log.Error().
			Str("journey", "userController.CreateUser").
			Msg(appErr.Error())
		return ctx.JSON(appErr.Code, appErr)
	}

	msg := fmt.Sprintf("User %s created successfully.", svc.UUID)
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": msg,
		"code":    http.StatusCreated,
	})
}

func (uc *UserController) GetUser(ctx echo.Context) error {
	uuid := ctx.Param("uuid")

	user, err := uc.UserService.GetUser(uuid)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	userResponse := uc.Model.ParseUserDomainToResponse(user)

	return ctx.JSON(http.StatusOK, userResponse)
}

func (uc *UserController) LoginUser(ctx echo.Context) error {
	var user request.UserLoginRequest

	err := ctx.Bind(&user)
	if err != nil {
		appErr := apperr.NewBadRequestError("Error binding request.")
		return ctx.JSON(appErr.Code, appErr)
	}

	if err = uc.Validate.Validate(user); err != nil {
		log.Error().
			Str("journey", "userController.LoginUser").
			Msgf(err.Error())

		// appErr := validation.ValidateUserError(err) 
		appErr := apperr.NewBadRequestError("Some fields are incorrect.")
		return ctx.JSON(appErr.Code, appErr)
	}

	userDomain := entity.User{
		Email:    user.Email,
		Password: user.Password,
	}

	appErr := uc.UserService.LoginUser(userDomain.Email, userDomain.Password)
	if appErr != nil {
		return ctx.JSON(appErr.Code, appErr)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "OK",
		"code":    http.StatusOK,
	})
}
