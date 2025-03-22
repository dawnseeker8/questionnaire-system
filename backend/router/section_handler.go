package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListSections handles GET /api/v1/forms/:formId/sections
func ListSections(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "List sections endpoint",
	})
}

// CreateSection handles POST /api/v1/forms/:formId/sections
func CreateSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Create section endpoint",
	})
}

// GetSection handles GET /api/v1/forms/:formId/sections/:id
func GetSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Get section endpoint",
	})
}

// UpdateSection handles PUT /api/v1/forms/:formId/sections/:id
func UpdateSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Update section endpoint",
	})
}

// DeleteSection handles DELETE /api/v1/forms/:formId/sections/:id
func DeleteSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Delete section endpoint",
	})
}
