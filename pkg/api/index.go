package api

import (
	"strings"

	"github.com/hooone/evening/pkg/models"
)

//Index  Main page
func (hs *HTTPServer) Index(c *models.ReqContext) {
	c.HTML(200, "index")
}

//NotFoundHandler 404 page
func (hs *HTTPServer) NotFoundHandler(c *models.ReqContext) {
	if strings.HasPrefix(c.Req.URL.Path, "/api") {
		c.JSON(404, "Not found")
		return
	}

	c.HTML(404, "index")
}
