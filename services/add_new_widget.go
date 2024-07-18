package services

import (
	"fmt"
	"net/http"
	"os"

	s "github.com/scblur869/pg-widget-api/structs"

	"github.com/gin-gonic/gin"
)

// api/v1/addNew
func (p *PgHandler) AddNewWidget(c *gin.Context) {
	var widget s.Widget
	if err := c.ShouldBindJSON(&widget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println(widget)
	res, err := p.Connect.Exec(p.Ctx, "INSERT INTO widget(name, category, color) VALUES($1,$2,$3)", widget.Name, widget.Category, widget.Color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	c.JSON(http.StatusAccepted, res.RowsAffected())
}
