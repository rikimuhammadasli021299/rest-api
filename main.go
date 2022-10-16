package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//car
// {
// 	"id": "1",
// 	"brand": "Honda",
// 	"car_type": "Civic"
// }

type car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"car_type"`
}

var cars = []car{
	{ID: "1", Brand: "Honda", Type: "City"},
	{ID: "2", Brand: "Toyota", Type: "Avanza"},
	{ID: "3", Brand: "Toyota", Type: "corolla"},
}

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK"})
	})

	//GET /cars - list cars
	r.GET("/cars", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cars)
	})

	//POST /cars - create car
	r.POST("/cars", func(ctx *gin.Context) {
		var car car
		if err := ctx.ShouldBindJSON(&car); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})

			return
		}
		cars = append(cars, car)
		ctx.JSON(http.StatusCreated, car)
	})

	//PUT /cars/:id - update car
	r.PUT("/cars/:car_id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")

		var updateCar car
		if error := ctx.BindJSON(&updateCar); error != nil {
			return
		}
		for i, car := range cars {
			if car.ID == id {
				cars[i] = updateCar
				ctx.IndentedJSON(http.StatusOK, cars[i])
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"errors": "Car Not Found",
		})
	})

	//get /cars/:id ambil satu data
	r.GET("/cars/:car_id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for _, car := range cars {
			if car.ID == id {
				ctx.IndentedJSON(http.StatusOK, car)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"errors": "Car Not Found",
		})
	})

	//DELETE /cars/:id - delete car
	r.DELETE("/cars/:car_id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				cars = append(cars[:i], cars[i+1:]...)
				break
			}
		}
		ctx.Status(http.StatusNoContent)
	})
	r.Run()

}
