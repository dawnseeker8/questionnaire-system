package service

import (
	"github.com/dawnseeker8/questionnaire-system/app/model"
)

// FormService handles business logic for forms
type FormService struct{}

// NewFormService creates a new form service
func NewFormService() *FormService {
	return &FormService{}
}

// GetAllForms retrieves all forms
func (s *FormService) GetAllForms() ([]model.Form, error) {
	var forms []model.Form
	result := model.DB.Find(&forms)
	return forms, result.Error
}

// GetFormByID retrieves a form by ID with all related data
func (s *FormService) GetFormByID(id uint) (*model.Form, error) {
	var form model.Form
	result := model.DB.Preload("Sections.Questions.Options").First(&form, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &form, nil
}

// CreateForm creates a new form
func (s *FormService) CreateForm(form *model.Form) error {
	return model.DB.Create(form).Error
}

// UpdateForm updates an existing form
func (s *FormService) UpdateForm(form *model.Form) error {
	return model.DB.Save(form).Error
}

// DeleteForm deletes a form by ID
func (s *FormService) DeleteForm(id uint) error {
	return model.DB.Delete(&model.Form{}, id).Error
}

// GetSectionByID retrieves a section by ID
func (s *FormService) GetSectionByID(id uint) (*model.Section, error) {
	var section model.Section
	result := model.DB.Preload("Questions.Options").First(&section, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &section, nil
}

// CreateSection creates a new section
func (s *FormService) CreateSection(section *model.Section) error {
	return model.DB.Create(section).Error
}

// UpdateSection updates an existing section
func (s *FormService) UpdateSection(section *model.Section) error {
	return model.DB.Save(section).Error
}

// DeleteSection deletes a section by ID
func (s *FormService) DeleteSection(id uint) error {
	return model.DB.Delete(&model.Section{}, id).Error
}
