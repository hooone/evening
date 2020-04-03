package api

import (
	"strings"

	"github.com/hooone/evening/pkg/api/dtos"
)

//Index  Main page
func (hs *HTTPServer) Index(c *dtos.ReqContext) {
	c.HTML(200, "index")
}

//NotFoundHandler 404 page
func (hs *HTTPServer) NotFoundHandler(c *dtos.ReqContext) {
	if strings.HasPrefix(c.Req.URL.Path, "/api") {
		c.JSON(404, "Not found")
		return
	}

	c.HTML(404, "index")
}
