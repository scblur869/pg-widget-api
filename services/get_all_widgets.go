package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/scblur869/pg-widget-api/structs"

	"github.com/gin-gonic/gin"
)

// api/v1/getAll
func (p *PgHandler) GetAllWidgets(c *gin.Context) {
	var name string
	var color string
	var category string
	var id int64
	var w structs.Widget
	var widgets []structs.Widget
	rows, err := p.Connect.Query(p.Ctx, "SELECT id, name, category, color FROM widget")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err.Error())
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &category, &color)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
			c.JSON(http.StatusUnprocessableEntity, err)
		}
		w.Id = id
		w.Name = name
		w.Category = category
		w.Color = color
		widgets = append(widgets, w)
	}
	c.JSON(http.StatusAccepted, widgets)

}
