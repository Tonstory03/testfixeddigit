package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomNumeric string

func (n CustomNumeric) MarshalJSON() ([]byte, error) {
	var result []byte = nil

	nn := string(n)

	convertData, err := strconv.ParseFloat(nn, 64)

	if err != nil {
		panic("CustomNumeric is n't numeric")
	}

	if strings.Contains(nn, ".") {
		result = []byte(fmt.Sprintf("%.2f", convertData))
	} else {
		result = []byte(fmt.Sprintf("%.0f", convertData))
	}

	return result, nil
}

type ResponseData struct {
	Data CustomNumeric `json:"data"`
}

func main() {
	router := gin.New()

	router.GET("/:number", func(c *gin.Context) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":       false,
					"errorMessage": r,
				})
			}
		}()

		c.Next()

	}, func(c *gin.Context) {

		number := c.Param("number")
		c.JSON(http.StatusOK, ResponseData{Data: CustomNumeric(number)})
	})

	router.Run(":8080")
}
