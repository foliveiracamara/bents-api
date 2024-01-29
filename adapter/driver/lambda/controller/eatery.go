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

type EateryController struct {
	EateryService port.EateryService
	Validate      *request.CustomValidator
	Model         *response.EateryResponse
}

func newEateryController(svc port.EateryService, cv *request.CustomValidator) *EateryController {
	return &EateryController{
		EateryService: svc,
		Validate:      cv,
	}
}

func InitEateryController(ctx echo.Context) (controller *EateryController) {
	eateryRepo := inmemory.NewInMemoryEateryAdapter()
	eaterySvc := service.NewEateryService(eateryRepo)
	cv := request.NewCustomValidator()
	controller = newEateryController(eaterySvc, cv)

	return controller
}

func (uc *EateryController) CreateEatery(ctx echo.Context) error {
	var req request.EateryRequest
	err := ctx.Bind(&req)
	if err != nil {
		appErr := apperr.NewBadRequestError("Error binding request.")
		return ctx.JSON(appErr.Code, appErr)
	}

	if err = uc.Validate.Validate(req); err != nil {
		log.Error().
			Str("journey", "eateryController.CreateEatery").
			Msgf(err.Error())

		appErr := apperr.NewBadRequestError("Some fields are incorrect.")
		return ctx.JSON(appErr.Code, appErr)
	}

	eateryDomain := &entity.Eatery{
		UUID:     uuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		Type:     req.Type,
		Password: req.Password,
	}

	svc, appErr := uc.EateryService.CreateEatery(eateryDomain)
	if appErr != nil {
		log.Error().
			Str("journey", "eateryController.CreateEatery").
			Msg(appErr.Error())
		return ctx.JSON(appErr.Code, appErr)
	}

	msg := fmt.Sprintf("Eatery %s created successfully.", svc.UUID)
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": msg,
		"code":    http.StatusCreated,
	})
}

func (uc *EateryController) GetEatery(ctx echo.Context) error {
	uuid := ctx.Param("uuid")

	user, err := uc.EateryService.GetEatery(uuid)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	userResponse := uc.Model.ParseEateryDomainToResponse(user)

	return ctx.JSON(http.StatusOK, userResponse)
}

func (uc *EateryController) FindEateries(ctx echo.Context) error {
	eateryType := ctx.QueryParam("type")
	rankType := ctx.QueryParam("rank")

	fmt.Printf("eatery type: %s, rank type: %s", eateryType, rankType)

	eateries, err := uc.EateryService.FindEateries(eateryType, rankType)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	// eateryResponse := uc.Model.ParseEateryDomainToResponse(user)

	return ctx.JSON(http.StatusOK, eateries)
}
