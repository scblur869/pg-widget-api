package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// api/v1/updateWidgetColor/:id/:color
func (p *PgHandler) UpdateWidgetColor(c *gin.Context) {
	color := c.Param("color")
	id := c.Param("id")

	res, err := p.Connect.Exec(p.Ctx, "UPDATE widget set color=$1 WHERE id = $2", color, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "update failed: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusAccepted, res.RowsAffected())
}

// api/v1/updateWidgetCategory/:id/:category
func (p *PgHandler) UpdateWidgetCategory(c *gin.Context) {
	category := c.Param("category")
	id := c.Param("id")

	res, err := p.Connect.Exec(p.Ctx, "UPDATE widget set category=$1 WHERE id = $2", category, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "update failed: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusAccepted, res.RowsAffected())
}
