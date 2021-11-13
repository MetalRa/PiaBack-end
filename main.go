package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductoTerminado struct {
	gorm.Model
	Modelo      string `json:"Modelo"`
	NumeroParte uint   `json:"NumeroParte"`
	Stock       uint   `json:"Stock"`
}

/*type (
	ProductoTerminado struct {
	gorm.Model
	Modelo      string `json:"Modelo"`
	NumeroParte uint   `json:"NumeroParte"`
	Stock       uint   `json:"Stock"`
	}
	jwtCustomClaims struct {
		Name  string `json:"name"`
		Admin bool   `json:"admin"`
		jwt.StandardClaims
	}
)*/

type Soldadura struct {
	gorm.Model
	Tipo     string `json:"Tipo"`
	Cantidad uint   `json:"Cantidad"`
}

type Componente struct {
	gorm.Model
	TipoComponente string `json:"TipoComponente"`
	NumeroParte    uint   `json:"NumeroParte"`
	Cantidad       uint   `json:"Cantidad"`
	Stock          uint   `json:"Stock"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func main() {

	// Echo instance
	e := echo.New()

	//Login route
	e.POST("/login", login)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.net", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//CRUD
	//Producto Terminado
	e.GET("/ProductoTerminado", getProductoTermiando)
	e.POST("/ProductoTerminado", createProductoTerminado)
	e.PUT("/ProductoTerminado/:ID", updateProductoTerminado)

	//Soldadura
	e.POST("/soldadura", createSoldadura)
	e.GET("/soldadura", getSoldadura)
	e.PUT("/soldadura/:ID", updateSoldadura)

	//Componentes
	e.PUT("/componente/:id", updateComponente)
	e.POST("/componente", createComponente)
	e.GET("/componente", getComponente)
	e.DELETE("/componente/:id", deleteComponente)

	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)
	//r.GET("/ProductoTerminado", getProductoTermiando)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

func getProductoTermiando(c echo.Context) error {
	var cont []ProductoTerminado
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	db.Find(&cont)
	return c.JSON(http.StatusOK, cont)
}

func createProductoTerminado(c echo.Context) error {
	prod := new(ProductoTerminado)
	if err := c.Bind(prod); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	db.Create(&prod)
	return c.JSON(http.StatusOK, prod)
}

func updateProductoTerminado(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	productoT := new(ProductoTerminado)
	db.First(&productoT, id)
	if err := c.Bind(productoT); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Save(&productoT)
	return c.JSON(http.StatusOK, productoT)
}

func createSoldadura(c echo.Context) error {
	sold := new(Soldadura)
	if err := c.Bind(sold); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	db.Create(&sold)

	return c.JSON(http.StatusOK, sold)
}

func getSoldadura(c echo.Context) error {
	var con []Soldadura
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	db.Find(&con)
	return c.JSON(http.StatusOK, con)
}

func updateSoldadura(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	product := new(ProductoTerminado)
	db.First(&product, id)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Save(&product)
	return c.JSON(http.StatusOK, product)
}
func deleteSoldadura(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	var res ProductoTerminado
	db.Where("ID = ? ", id).Delete(&res)
	return c.NoContent(http.StatusNoContent)
}

func createComponente(c echo.Context) error {
	comp := new(Componente)
	if err := c.Bind(comp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	db.Create(&comp)

	return c.JSON(http.StatusOK, comp)
}

func getComponente(c echo.Context) error {
	var cont []Componente
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	db.Find(&cont)
	return c.JSON(http.StatusOK, cont)
}

func updateComponente(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	comp := new(Componente)
	db.First(&comp, id)
	if err := c.Bind(comp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db.Save(&comp)
	return c.JSON(http.StatusOK, comp)
}
func deleteComponente(c echo.Context) error {
	idc, _ := strconv.Atoi(c.Param("id"))
	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	var com Componente
	db.Where("ID = ? ", idc).Delete(&com)
	return c.NoContent(http.StatusNoContent)
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "adm" || password != "raul123" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Raul",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
