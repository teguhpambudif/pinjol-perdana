package controller

import (
	"net/http"
	"pinjol-perdana/model"
	"pinjol-perdana/usecase"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	router     *gin.Engine
	accUsecase usecase.AccountUseCase
}

func (ac *AccountController) RegisterNewAccount(ctx *gin.Context) {
	var newAccount *model.Account
	err := ctx.BindJSON(&newAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ac.accUsecase.RegisterNewAccount(newAccount)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"account": newAccount,
		})
	}
}

func (ac *AccountController) GetAccountById(ctx *gin.Context) {
	idAccount := ctx.Param("id")
	responseUc, err := ac.accUsecase.GetAccountById(idAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		if (responseUc == model.Account{}) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Id tidak ditemukan",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    responseUc,
			})
		}
	}
}

func (ac *AccountController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := "assets/" + file.Filename
	err = ctx.SaveUploadedFile(file, filename)
	if err != nil {
		ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "berhasil diupload",
			"file_ktp": filename,
		})
	}
}

func NewAccountController(router *gin.Engine, accUseCase usecase.AccountUseCase) *AccountController {
	newAccController := AccountController{
		router:     router,
		accUsecase: accUseCase,
	}

	account := router.Group("/account")
	account.POST("", newAccController.RegisterNewAccount)
	account.GET("/:id", newAccController.GetAccountById)
	account.PUT("/:id/upload", newAccController.Upload)
	return &newAccController
}
