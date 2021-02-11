// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateMedicineRequest create medicine request
//
// swagger:model CreateMedicineRequest
type CreateMedicineRequest struct {

	// Effective until n months after the full vaccination schedule is completed
	EffectiveUntil float64 `json:"effectiveUntil,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Indicative price if fixed or max price available.
	Price float64 `json:"price,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// schedule
	Schedule *CreateMedicineRequestSchedule `json:"schedule,omitempty"`

	// status
	// Enum: [Active Inactive Blocked]
	Status string `json:"status,omitempty"`

	// vaccination mode
	// Enum: [muscular injection oral nasal]
	VaccinationMode string `json:"vaccinationMode,omitempty"`
}

// Validate validates this create medicine request
func (m *CreateMedicineRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccinationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMedicineRequest) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

var createMedicineRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive","Blocked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createMedicineRequestTypeStatusPropEnum = append(createMedicineRequestTypeStatusPropEnum, v)
	}
}

const (

	// CreateMedicineRequestStatusActive captures enum value "Active"
	CreateMedicineRequestStatusActive string = "Active"

	// CreateMedicineRequestStatusInactive captures enum value "Inactive"
	CreateMedicineRequestStatusInactive string = "Inactive"

	// CreateMedicineRequestStatusBlocked captures enum value "Blocked"
	CreateMedicineRequestStatusBlocked string = "Blocked"
)

// prop value enum
func (m *CreateMedicineRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createMedicineRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateMedicineRequest) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var createMedicineRequestTypeVaccinationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["muscular injection","oral","nasal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createMedicineRequestTypeVaccinationModePropEnum = append(createMedicineRequestTypeVaccinationModePropEnum, v)
	}
}

const (

	// CreateMedicineRequestVaccinationModeMuscularInjection captures enum value "muscular injection"
	CreateMedicineRequestVaccinationModeMuscularInjection string = "muscular injection"

	// CreateMedicineRequestVaccinationModeOral captures enum value "oral"
	CreateMedicineRequestVaccinationModeOral string = "oral"

	// CreateMedicineRequestVaccinationModeNasal captures enum value "nasal"
	CreateMedicineRequestVaccinationModeNasal string = "nasal"
)

// prop value enum
func (m *CreateMedicineRequest) validateVaccinationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createMedicineRequestTypeVaccinationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateMedicineRequest) validateVaccinationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.VaccinationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateVaccinationModeEnum("vaccinationMode", "body", m.VaccinationMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create medicine request based on the context it is used
func (m *CreateMedicineRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMedicineRequest) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMedicineRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMedicineRequest) UnmarshalBinary(b []byte) error {
	var res CreateMedicineRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateMedicineRequestSchedule create medicine request schedule
//
// swagger:model CreateMedicineRequestSchedule
type CreateMedicineRequestSchedule struct {

	// osid
	Osid string `json:"osid,omitempty"`

	// Number of times the vaccination should be taken.
	RepeatInterval float64 `json:"repeatInterval,omitempty"`

	// How many times vaccination should be taken
	RepeatTimes float64 `json:"repeatTimes,omitempty"`
}

// Validate validates this create medicine request schedule
func (m *CreateMedicineRequestSchedule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create medicine request schedule based on context it is used
func (m *CreateMedicineRequestSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateMedicineRequestSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMedicineRequestSchedule) UnmarshalBinary(b []byte) error {
	var res CreateMedicineRequestSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
