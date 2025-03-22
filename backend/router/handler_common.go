package router

import (
	"github.com/dawnseeker8/questionnaire-system/app/service"
)

// Services used by handlers
var (
	formService *service.FormService
)

func init() {
	// Initialize services
	formService = service.NewFormService()
}
