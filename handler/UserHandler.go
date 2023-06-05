package handler

import (
	"final-project-2/dto"
	"final-project-2/pkg/errs"
	"final-project-2/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

// NewUserHandler membuat instance baru dari userHandler.
// Menerima parameter userService yang merupakan implementasi dari service.UserService.
// Mengembalikan userHandler yang baru dibuat.
func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Register digunakan untuk menangani permintaan pendaftaran pengguna baru.
func (uh *userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest

	// Membaca dan mengikat JSON permintaan ke struct newUserRequest.
	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		// Membuat kesalahan UnprocessableEntityError dengan pesan "invalid request body".
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		// Mengirim respons JSON dengan kode status dan pesan kesalahan yang sesuai.
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	// Memanggil userService untuk membuat pengguna baru.
	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		// Mengirim detail kesalahan ke klien
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(result.StatusCode, result)
}

// Login digunakan untuk menangani permintaan login pengguna.
func (uh *userHandler) Login(ctx *gin.Context) {
	var loginUserRequest dto.LoginUserRequest

	// Membaca dan mengikat JSON permintaan ke struct LoginRequest.
	if err := ctx.ShouldBindJSON(&loginUserRequest); err != nil {
		// Membuat kesalahan UnprocessableEntityError dengan pesan "invalid request body".
		errBindJSON := errs.NewUnprocessibleEntityError("invalid request body")

		// Mengirim respons JSON dengan kode status dan pesan kesalahan yang sesuai.
		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	// Memanggil userService untuk melakukan login pengguna.
	result, err := uh.userService.Login(loginUserRequest)

	if err != nil {
		// Mengirim detail kesalahan ke klien
		ctx.JSON(err.Status(), err)
		return
	}

	// Mengirim respons JSON dengan kode status dan hasil yang diterima dari userService.
	ctx.JSON(result.StatusCode, result)
}
