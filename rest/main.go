package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculation = []Calculation{}

func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression) //делаем выражение (55+55)
	if err != nil {
		return "", err // передали 55++55
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err
}

func getCalculation(c echo.Context) error {
	return c.JSON(http.StatusOK, calculation)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.RequestLogger())

	e.POST("/calculations", getCalculation)

	e.Start("localhost:8080")
}

