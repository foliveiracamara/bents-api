package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
		Category: req.Category,
		Password: req.Password,
		Rank:     req.Rank,
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
	name := ctx.Param("name")

	user, err := uc.EateryService.GetEatery(name)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	model := response.EateriesResponse{}
	userResponse := model.ParseEateryDomainToResponse(user)

	log.Info().
		Str("journey", "eateryController.GetEatery").
		Msgf("Eateries found successfully.",)

	return ctx.JSON(http.StatusOK, userResponse)
}

func (uc *EateryController) FindEateries(ctx echo.Context) error {
	eateryCategory := ctx.QueryParam("type")
	eateryRank := ctx.QueryParam("rank")

	eateryRankInt, err := strconv.Atoi(eateryRank)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	filters := map[string]interface{}{
		"category": eateryCategory,
		"rank":     eateryRankInt,
	}

	eateries, appErr := uc.EateryService.FindEateries(filters)
	if err != nil {
		return ctx.JSON(appErr.Code, appErr)
	}

	model := &response.EateriesResponse{}
	eateriesResponse := model.ParseEateryDomainToResponse(eateries)

	fmt.Println("response: ", eateriesResponse)

	return ctx.JSON(http.StatusOK, eateriesResponse)
}
