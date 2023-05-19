// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UserDataPropertiesGender - MALE, FEMALE, or OTHER
type UserDataPropertiesGender string

const (
	UserDataPropertiesGenderMale   UserDataPropertiesGender = "male"
	UserDataPropertiesGenderFemale UserDataPropertiesGender = "female"
	UserDataPropertiesGenderOther  UserDataPropertiesGender = "other"
)

func (e UserDataPropertiesGender) ToPointer() *UserDataPropertiesGender {
	return &e
}

func (e *UserDataPropertiesGender) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "male":
		fallthrough
	case "female":
		fallthrough
	case "other":
		*e = UserDataPropertiesGender(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserDataPropertiesGender: %v", v)
	}
}

// UserDataProperties - Personal data - name, birthday, SSN etc. None of the fields in user_data are required. You may send in an empty array: `"user_data": {}`.
type UserDataProperties struct {
	// Day of user's date of birth (e.g. 23)
	DayOfBirth *int64 `json:"day_of_birth,omitempty"`
	// User's associated first name
	FirstName *string `json:"first_name,omitempty"`
	// MALE, FEMALE, or OTHER
	Gender *UserDataPropertiesGender `json:"gender,omitempty"`
	// User's associated last name
	LastName *string `json:"last_name,omitempty"`
	// User's associated middle name
	MiddleName *string `json:"middle_name,omitempty"`
	// Month of user's date of birth (e.g. 12 for December)
	MonthOfBirth *int64 `json:"month_of_birth,omitempty"`
	// Social security number of the user, in the format xxx-xx-xxxx
	Ssn *string `json:"ssn,omitempty"`
	// Year of the user's date of birth (e.g. 1990)
	YearOfBirth *int64 `json:"year_of_birth,omitempty"`
}
