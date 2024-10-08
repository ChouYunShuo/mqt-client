/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PaymentSourceStatusEnum Current status of the payment source.
type PaymentSourceStatusEnum string

// List of PaymentSourceStatusEnum
const (
	PAYMENTSOURCESTATUSENUM_ACTIVE PaymentSourceStatusEnum = "ACTIVE"
	PAYMENTSOURCESTATUSENUM_PENDING PaymentSourceStatusEnum = "PENDING"
	PAYMENTSOURCESTATUSENUM_INACTIVE PaymentSourceStatusEnum = "INACTIVE"
)

// All allowed values of PaymentSourceStatusEnum enum
var AllowedPaymentSourceStatusEnumEnumValues = []PaymentSourceStatusEnum{
	"ACTIVE",
	"PENDING",
	"INACTIVE",
}

func (v *PaymentSourceStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentSourceStatusEnum(value)
	for _, existing := range AllowedPaymentSourceStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentSourceStatusEnum", value)
}

// NewPaymentSourceStatusEnumFromValue returns a pointer to a valid PaymentSourceStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentSourceStatusEnumFromValue(v string) (*PaymentSourceStatusEnum, error) {
	ev := PaymentSourceStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentSourceStatusEnum: valid values are %v", v, AllowedPaymentSourceStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentSourceStatusEnum) IsValid() bool {
	for _, existing := range AllowedPaymentSourceStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentSourceStatusEnum value
func (v PaymentSourceStatusEnum) Ptr() *PaymentSourceStatusEnum {
	return &v
}

type NullablePaymentSourceStatusEnum struct {
	value *PaymentSourceStatusEnum
	isSet bool
}

func (v NullablePaymentSourceStatusEnum) Get() *PaymentSourceStatusEnum {
	return v.value
}

func (v *NullablePaymentSourceStatusEnum) Set(val *PaymentSourceStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSourceStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSourceStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSourceStatusEnum(val *PaymentSourceStatusEnum) *NullablePaymentSourceStatusEnum {
	return &NullablePaymentSourceStatusEnum{value: val, isSet: true}
}

func (v NullablePaymentSourceStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSourceStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

