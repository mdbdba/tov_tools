package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tov_tools/pkg/dice"
)

func RollDice(c *gin.Context) {
	sidesParam := c.Query("sides")
	timesParam := c.Query("timesToRoll")
	optionsParams := c.QueryArray("options")

	if sidesParam == "" || timesParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	sides, err := strconv.Atoi(sidesParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sides parameter"})
		return
	}

	timesToRoll, err := strconv.Atoi(timesParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timesToRoll parameter"})
		return
	}

	ctxRef := "dice_handler"

	result, err := dice.Perform(sides, timesToRoll, ctxRef, optionsParams...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
