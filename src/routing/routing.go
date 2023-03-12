package routing

import (
    "golang.org/x/crypto/bcrypt"
    "static_auth/database"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

//  ──────────────────────────────────────────────────────────────────────────

func InitRouter() *echo.Echo {

    router := echo.New()
    router.Static("/", "website/html")
    router.Use(middleware.BasicAuth(login))

	return router
}


func login(username, password string, c echo.Context) (bool, error) {

    storePassword := database.GetFromBucket("Users", username)

    if err := bcrypt.CompareHashAndPassword([]byte(storePassword), []byte(password)); err != nil {
        return false, nil
    } else {
        return true, nil
    }
}
