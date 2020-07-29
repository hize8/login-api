package userHandler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"

	"github.com/hize8/login-api/utils"
)

func SignUp(c echo.Context) error {
	u := new(User)
	err := c.Bind(u)
	utils.CheckError(err)
	userToSearch := getUserByEmail(u.Email)
	if userToSearch.ID > 0 {
		return c.String(http.StatusOK, "El usuario ya ha sido creado")
	}
	saveNewUser(User{Name: u.Name, Email: u.Email, Password: u.Password})
	return c.JSON(http.StatusOK, u)
}

func Login(c echo.Context) error {
	u := new(User)
	err := c.Bind(u)
	utils.CheckError(err)
	userToSearch := getUserByEmail(u.Email)
	if userToSearch.ID == 0 {
		return c.String(http.StatusOK, "El usuario no existe")
	}

	if u.Password != userToSearch.Password {
		return c.String(http.StatusOK, "Contraseña incorrecta")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userToSearch.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(3 * time.Hour).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func GetUsers(c echo.Context) error {
	allUsers := getAllUsers()
	return c.JSON(http.StatusOK, allUsers)
}
