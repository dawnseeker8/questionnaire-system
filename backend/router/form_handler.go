package router

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/dawnseeker8/questionnaire-system/app/dto"
	"github.com/dawnseeker8/questionnaire-system/app/model"
)

// ListForms godoc
// @Summary List all forms
// @Description Get a list of all questionnaire forms
// @Tags forms
// @Accept json
// @Produce json
// @Success 200 {array} model.Form
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms [get]
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

// CreateForm godoc
// @Summary Create a new form
// @Description Create a new questionnaire form
// @Tags forms
// @Accept json
// @Produce json
// @Param form body dto.FormCreateRequest true "Form data"
// @Success 201 {object} model.Form
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms [post]
func CreateForm(ctx context.Context, c *app.RequestContext) {
	var formRequest dto.FormCreateRequest
	if err := c.BindAndValidate(&formRequest); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Create a new form model from the request DTO
	form := &model.Form{
		Title:       formRequest.Title,
		Description: formRequest.Description,
	}

	if err := formService.CreateForm(form); err != nil {
		hlog.Error("Failed to create form:", err)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create form",
		})
		return
	}

	c.JSON(consts.StatusCreated, form)
}

// GetForm godoc
// @Summary Get a form by ID
// @Description Get a questionnaire form by its ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path int true "Form ID"
// @Success 200 {object} model.Form
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/forms/{id} [get]
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

// UpdateForm godoc
// @Summary Update a form
// @Description Update a questionnaire form by its ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path int true "Form ID"
// @Param form body dto.FormUpdateRequest true "Form data"
// @Success 200 {object} model.Form
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms/{id} [put]
func UpdateForm(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID format",
		})
		return
	}

	var formRequest dto.FormUpdateRequest
	if err := c.BindAndValidate(&formRequest); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Get the existing form first
	existingForm, err := formService.GetFormByID(uint(id))
	if err != nil {
		hlog.Error("Failed to get form for update:", err)
		c.JSON(consts.StatusNotFound, map[string]interface{}{
			"error": "Form not found",
		})
		return
	}

	// Update only the fields that should be updated
	existingForm.Title = formRequest.Title
	existingForm.Description = formRequest.Description

	if err := formService.UpdateForm(existingForm); err != nil {
		hlog.Error("Failed to update form:", err)
		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update form",
		})
		return
	}

	c.JSON(consts.StatusOK, existingForm)
}

// DeleteForm godoc
// @Summary Delete a form
// @Description Delete a questionnaire form by its ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path int true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/forms/{id} [delete]
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
