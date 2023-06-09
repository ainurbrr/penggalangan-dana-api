package routes

import (
	"net/http"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/constants"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/controllers"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/middlewares"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Routes(e *echo.Echo, db *gorm.DB) {

	middlewares.LogMiddleware(e)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(mid.CORS())
	e.Pre(mid.RemoveTrailingSlash())

	e.Static("/images/avatar", "./images/avatar")
	e.Static("/images/campaignImages", "./images/campaignImages")

	e.POST("/users", controllers.RegisterUserController)
	e.POST("/login", controllers.LoginUserController)
	e.PUT("/avatar", controllers.UploadAvatarController, mid.JWT([]byte(constants.SECRET_JWT)))

	e.GET("/campaigns", controllers.GetCampaignsController)
	e.GET("/campaigns/:id", controllers.GetCampaignDetailController)
	e.POST("/campaign", controllers.CreateCampaignController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.PUT("/campaigns/:id", controllers.UpdateCampaignController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.DELETE("/campaigns/:id", controllers.DeleteCampaignController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.POST("/campaign-images/:id", controllers.UploadCampaignImageController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.GET("/campaigns/:id/transactions", controllers.GetCampaignTransactionsController, mid.JWT([]byte(constants.SECRET_JWT)))

	transactions := e.Group("/transactions", mid.JWT([]byte(constants.SECRET_JWT)))
	transactions.GET("", controllers.GetUserTransactionsController)
	transactions.POST("", controllers.CreateTransactionController)
	transactions.POST("/:id", controllers.ProcessPaymentController)
}
