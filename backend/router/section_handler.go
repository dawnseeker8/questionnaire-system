package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListSections godoc
// @Summary List all sections of a form
// @Description Get a list of all sections for a specific form
// @Tags sections
// @Accept json
// @Produce json
// @Param formId path int true "Form ID"
// @Success 200 {array} model.Section
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/forms/{formId}/sections [get]
func ListSections(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "List sections endpoint",
	})
}

// CreateSection godoc
// @Summary Create a new section
// @Description Create a new section for a specific form
// @Tags sections
// @Accept json
// @Produce json
// @Param formId path int true "Form ID"
// @Param section body model.Section true "Section data"
// @Success 201 {object} model.Section
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms/{formId}/sections [post]
func CreateSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Create section endpoint",
	})
}

// GetSection godoc
// @Summary Get a section by ID
// @Description Get a section by its ID for a specific form
// @Tags sections
// @Accept json
// @Produce json
// @Param formId path int true "Form ID"
// @Param id path int true "Section ID"
// @Success 200 {object} model.Section
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/forms/{formId}/sections/{id} [get]
func GetSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Get section endpoint",
	})
}

// UpdateSection godoc
// @Summary Update a section
// @Description Update a section by its ID for a specific form
// @Tags sections
// @Accept json
// @Produce json
// @Param formId path int true "Form ID"
// @Param id path int true "Section ID"
// @Param section body model.Section true "Section data"
// @Success 200 {object} model.Section
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms/{formId}/sections/{id} [put]
func UpdateSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Update section endpoint",
	})
}

// DeleteSection godoc
// @Summary Delete a section
// @Description Delete a section by its ID for a specific form
// @Tags sections
// @Accept json
// @Produce json
// @Param formId path int true "Form ID"
// @Param id path int true "Section ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms/{formId}/sections/{id} [delete]
func DeleteSection(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Delete section endpoint",
	})
}
