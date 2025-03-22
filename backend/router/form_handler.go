package router

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/dawnseeker8/questionnaire-system/app/model"
)

// ListForms handles GET /api/v1/forms
func ListForms(ctx context.Context, c *app.RequestContext) {
	forms, err := formService.GetAllForms()
	if err != nil {
		hlog.Error("Failed to get forms:", err)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get forms",
		})
		return
	}

	c.JSON(consts.StatusOK, forms)
}

// CreateForm handles POST /api/v1/forms
func CreateForm(ctx context.Context, c *app.RequestContext) {
	var form model.Form
	if err := c.BindAndValidate(&form); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err := formService.CreateForm(&form); err != nil {
		hlog.Error("Failed to create form:", err)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create form",
		})
		return
	}

	c.JSON(consts.StatusCreated, form)
}

// GetForm handles GET /api/v1/forms/:id
func GetForm(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID format",
		})
		return
	}

	form, err := formService.GetFormByID(uint(id))
	if err != nil {
		hlog.Error("Failed to get form:", err)
		c.JSON(consts.StatusNotFound, map[string]interface{}{
			"error": "Form not found",
		})
		return
	}

	c.JSON(consts.StatusOK, form)
}

// UpdateForm handles PUT /api/v1/forms/:id
func UpdateForm(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID format",
		})
		return
	}

	var form model.Form
	if err := c.BindAndValidate(&form); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	form.ID = uint(id)
	if err := formService.UpdateForm(&form); err != nil {
		hlog.Error("Failed to update form:", err)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update form",
		})
		return
	}

	c.JSON(consts.StatusOK, form)
}

// DeleteForm handles DELETE /api/v1/forms/:id
func DeleteForm(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID format",
		})
		return
	}

	if err := formService.DeleteForm(uint(id)); err != nil {
		hlog.Error("Failed to delete form:", err)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete form",
		})
		return
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "Form deleted successfully",
	})
}
