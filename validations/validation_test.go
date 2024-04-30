package validations_test

import (
	validation "cheat-sheets-golang-project/validations"
	"testing"
)

func TestExampleErrValidation(t *testing.T) {
	// This example demonstrates how to use the ErrValidation struct.

	type record struct {
		Name string
		Age  int
	} // This is a sample struct.

	// This is a sample function that validates the record.
	validate := func(r record) []validation.ErrValidation {
		var errs []validation.ErrValidation
		if r.Name == "" {
			errs = append(errs, validation.ErrValidation{
				Field:   "Name",
				Message: "Name is required",
			})
		}
		if r.Age < 0 {
			errs = append(errs, validation.ErrValidation{
				Field:   "Age",
				Message: "Age must be greater than zero",
			})
		}
		return errs
	}

	t.Run("example while Name is empty", func(t *testing.T) {
		r := record{Name: "", Age: 20}
		errs := validate(r)
		if len(errs) != 1 {
			t.Errorf("expected 1 error, got %d", len(errs))
		}
		if errs[0].Field != "Name" {
			t.Errorf("expected field Name, got %s", errs[0].Field)
		}
		if errs[0].Message != "Name is required" {
			t.Errorf("expected message Name is required, got %s", errs[0].Message)
		}
	})

	t.Run("example while Age is less than zero", func(t *testing.T) {
		r := record{Name: "John", Age: -1}
		errs := validate(r)
		if len(errs) != 1 {
			t.Errorf("expected 1 error, got %d", len(errs))
		}
		if errs[0].Field != "Age" {
			t.Errorf("expected field Age, got %s", errs[0].Field)
		}
		if errs[0].Message != "Age must be greater than zero" {
			t.Errorf("expected message Age must be greater than zero, got %s", errs[0].Message)
		}
	})

	t.Run("example while both Name is empty and Age is less than zero", func(t *testing.T) {
		r := record{Name: "", Age: -1}
		errs := validate(r)
		if len(errs) != 2 {
			t.Errorf("expected 2 errors, got %d", len(errs))
		}
		if errs[0].Field != "Name" {
			t.Errorf("expected field Name, got %s", errs[0].Field)
		}
		if errs[0].Message != "Name is required" {
			t.Errorf("expected message Name is required, got %s", errs[0].Message)
		}
		if errs[1].Field != "Age" {
			t.Errorf("expected field Age, got %s", errs[1].Field)
		}
		if errs[1].Message != "Age must be greater than zero" {
			t.Errorf("expected message Age must be greater than zero, got %s", errs[1].Message)
		}
	})
}
