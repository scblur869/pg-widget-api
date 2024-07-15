package services

import (
	"fmt"
	s "local/pg-widget-api/structs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// api/v1/getByColor/:color
func (p *PgHandler) GetWidgetsByColor(c *gin.Context) {
	var widget s.Widget
	color := c.Param("color")

	err := p.Connect.QueryRow(p.Ctx, "SELECT id, name, category, color FROM widget WHERE color=$1", color).Scan(&widget.Id, &widget.Name, &widget.Category, &widget.Color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusAccepted, widget)
}

// api/v1/getByCategory/:category
func (p *PgHandler) GetWidgetsByCategory(c *gin.Context) {
	var widget s.Widget
	_type := c.Param("category")

	err := p.Connect.QueryRow(p.Ctx, "SELECT id, name, category, color FROM widget WHERE category=$1", _type).Scan(&widget.Id, &widget.Name, &widget.Category, &widget.Color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusAccepted, widget)
}
