package handlers

import (
	"context"
	"github/RegisterTask/Api/api/models"
	_ "github/RegisterTask/Api/api/models"
	_ "github/RegisterTask/Api/api/token"
	"github/RegisterTask/Api/etc"
	pb "github/RegisterTask/Api/genproto/register_service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

// Register godoc
// @Summary Create new user
// @Description This API for creating a new user
// @Tags Reg
// @Accept json
// @Param body body models.User true "body"
// @Produce json
// @Success 201 {object} models.User
// @Router /reg [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// Creating access and refresh tokens
	accessTokenString, refreshTokenString, err := h.jwtHandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusConflict, models.StandardErrorModel{
			Error: models.Error{
				Message: "Error while generating tokens",
			},
		})
		return
	}

	// Creating hash of a password
	hashedPassword, err := etc.GeneratePasswordHash(body.Password)
	if err != nil {
		c.JSON(http.StatusConflict, models.StandardErrorModel{
			Error: models.Error{
				Message: "Error while generating hash for password",
			},
		})
		return
	}

	response, err := h.serviceManager.RegisterService().Create(ctx, &pb.User{
		Id:           body.Id,
		Name:         body.Name,
		Password:     string(hashedPassword),
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user")
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser godoc
// @Summary GetUser
// @Schemes
// @Description  Get User
// @Security 	BearerAuth
// @Tags Reg
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func (h *handlerV1) Get(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.RegisterService().Get(
		ctx, &pb.ById{
			Id: guid,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User")
		return
	}

	c.JSON(http.StatusOK, response)

}

// Login godoc
// @Summary Login
// @Schemes
// @Description  Get My Profile
// @Security 	BearerAuth
// @Tags Reg
// @Accept json
// @Param body body models.Login true "body"
// @Produce json
// @Success 200 {object} models.User
// @Router /login [post]
func (h *handlerV1) Login(c *gin.Context) {
	var (
		body        pb.Loginn
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.RegisterService().Login(ctx, &pb.Loginn{
		Password: string(body.Password),
		Name:     body.Name,
	})

	compareErr := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(body.Password))

	if compareErr != nil {
		c.JSON(http.StatusConflict, models.StandardErrorModel{
			Error: models.Error{
				Message: "Password Error",
			},
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User Profile")
		return
	}

	c.JSON(http.StatusOK, response)

}
