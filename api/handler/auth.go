package handler

import (
	"app/api/models"
	"app/config"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Create Register
// @Description Create Register
// @Tags Register
// @Accept json
// @Produce json
// @Param Regester body models.Register true "RegisterRequestBody"
// @Success 201 {object} models.RegisterResponse "GetRegisterBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Register(c *gin.Context) {
	
	var register models.Register

	err := c.ShouldBindJSON(&register)
	if err != nil{
		h.handlerResponse(c, "Register create", 400, err.Error())
		return
	}

	id, err := h.storages.User().Create(context.Background(), 
		&models.CreateUser{
			FirstName: register.FirstName,
			LastName: register.LastName,
			Login: register.Login,
			Password: register.Password,
			Phone_number: register.PhoneNumber,
		},
	)

	if err != nil{
		h.handlerResponse(c, "Register create storage", 500, err.Error())
		return
	}

	user, err := h.storages.User().GetByID(context.Background(), &models.UserPrimaryKey{UserId: id})
	if err != nil{
		h.handlerResponse(c, "Register Get By Id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Register User", 200, user)
}

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Create Login
// @Description Create Login
// @Tags Login
// @Accept json
// @Produce json
// @Param Login body models.Login true "LoginRequestBody"
// @Success 201 {object} models.LoginResponse "GetLoginBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil{
		h.handlerResponse(c, "Auth Login", 400, err.Error())
		return
	}

	user, err := h.storages.User().GetByID_Login(context.Background(), &models.Login{Login: login.Login, Password: login.Password})
	if err != nil{
		h.handlerResponse(c, "Login Storage Get By id", 500, err.Error())
		return
	}
	
	data := map[string]interface{}{
		"user_id": user.UserId,
		"first_name":user.FirstName,
		"last_name":user.LastName,
		"login":user.Login,
		"password":user.Password,
		"phone_number":user.Phone_number,
		"created_at":user.Created_at,
		"updated_at":user.Updated_at,
	}

	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SecretKey)
	if err != nil {
		h.handlerResponse(c, "Login Generate token", 500, err.Error())
		return
	}

	// h.handlerResponse(c, "Login User", 200, models.LoginResponse{AccessToken: token})
	c.JSON(http.StatusCreated, models.RegisterResponse{AccessToken: token})
}