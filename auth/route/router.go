package route

import (
	"net/http"
	"time"
	"auth/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ "http://localhost:8080" },
		AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
	}))

	e.POST("/auth", auth)

	return e;
}

func auth(c echo.Context) error {
	db, err := gorm.Open("mysql", "vgo:password@tcp(db:3306)/vgo?parseTime=true")
	if err != nil {
		return err
	}

	var user dto.User
	db.Where("name = ? AND pass = ?", c.FormValue("username"), c.FormValue("password")).First(&user)

	if user.ID == 0 {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	exp := time.Now().Add(time.Hour * 72)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["admin"] = user.Administrator
	claims["exp"] = exp.Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	user.Token = t
	user.Token_Expire_Date = &exp
	db.Save(&user)

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
