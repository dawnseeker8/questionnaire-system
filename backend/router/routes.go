package router

import (
	"github.com/cloudwego/hertz/pkg/route"
)

// setupQuestionnaireRoutes registers all form related routes
func setupQuestionnaireRoutes(group *route.RouterGroup) {
	forms := group.Group("/forms")

	// Form routes
	forms.GET("", ListForms)
	forms.POST("", CreateForm)
	forms.GET("/:id", GetForm)
	forms.PUT("/:id", UpdateForm)
	forms.DELETE("/:id", DeleteForm)

	// Section routes
	forms.GET("/:formId/sections", ListSections)
	forms.POST("/:formId/sections", CreateSection)
	forms.GET("/:formId/sections/:id", GetSection)
	forms.PUT("/:formId/sections/:id", UpdateSection)
	forms.DELETE("/:formId/sections/:id", DeleteSection)
}
