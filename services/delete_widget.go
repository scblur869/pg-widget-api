package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// api/v1/deleteById/:id
func (p *PgHandler) DeleteWidget(c *gin.Context) {

	id := c.Param("id")

	res, err := p.Connect.Exec(p.Ctx, "DELETE FROM widget WHERE id = $1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "delete failed: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusAccepted, res.RowsAffected())
}
