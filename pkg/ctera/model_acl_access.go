/*
 * CTERA Gateway
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ctera

import (
	"encoding/json"
	"fmt"
)

// AclAccess the model 'AclAccess'
type AclAccess string

// List of AclAccess
const (
	WINDOWS_NT               AclAccess = "WindowsNT"
	ONLY_AUTHENTICATED_USERS AclAccess = "OnlyAuthenticatedUsers"
)

var allowedAclAccessEnumValues = []AclAccess{
	"WindowsNT",
	"OnlyAuthenticatedUsers",
}

func (v *AclAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclAccess(value)
	for _, existing := range allowedAclAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclAccess", value)
}

// NewAclAccessFromValue returns a pointer to a valid AclAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclAccessFromValue(v string) (*AclAccess, error) {
	ev := AclAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclAccess: valid values are %v", v, allowedAclAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclAccess) IsValid() bool {
	for _, existing := range allowedAclAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclAccess value
func (v AclAccess) Ptr() *AclAccess {
	return &v
}

type NullableAclAccess struct {
	value *AclAccess
	isSet bool
}

func (v NullableAclAccess) Get() *AclAccess {
	return v.value
}

func (v *NullableAclAccess) Set(val *AclAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableAclAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAclAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclAccess(val *AclAccess) *NullableAclAccess {
	return &NullableAclAccess{value: val, isSet: true}
}

func (v NullableAclAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
