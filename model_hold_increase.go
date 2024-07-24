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
	"bytes"
	"fmt"
	"github.com/shopspring/decimal"
)

// checks if the HoldIncrease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HoldIncrease{}

// HoldIncrease Controls automatic increases to the authorization amount for MCCs specified in this group.
type HoldIncrease struct {
	// Controls whether the `value` field represents a fixed amount or a percentage of the authorization amount.
	Type string `json:"type"`
	// Specifies the amount of the automatic increase to the authorization amount.  The `type` field controls whether this amount is a fixed amount or a percentage.
	Value decimal.Decimal `json:"value"`
}

type _HoldIncrease HoldIncrease

// NewHoldIncrease instantiates a new HoldIncrease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldIncrease(type_ string, value decimal.Decimal) *HoldIncrease {
	this := HoldIncrease{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewHoldIncreaseWithDefaults instantiates a new HoldIncrease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldIncreaseWithDefaults() *HoldIncrease {
	this := HoldIncrease{}
	var type_ string = "AMOUNT"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *HoldIncrease) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HoldIncrease) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HoldIncrease) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *HoldIncrease) GetValue() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HoldIncrease) GetValueOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HoldIncrease) SetValue(v decimal.Decimal) {
	o.Value = v
}

func (o HoldIncrease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HoldIncrease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *HoldIncrease) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHoldIncrease := _HoldIncrease{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHoldIncrease)

	if err != nil {
		return err
	}

	*o = HoldIncrease(varHoldIncrease)

	return err
}

type NullableHoldIncrease struct {
	value *HoldIncrease
	isSet bool
}

func (v NullableHoldIncrease) Get() *HoldIncrease {
	return v.value
}

func (v *NullableHoldIncrease) Set(val *HoldIncrease) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldIncrease) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldIncrease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldIncrease(val *HoldIncrease) *NullableHoldIncrease {
	return &NullableHoldIncrease{value: val, isSet: true}
}

func (v NullableHoldIncrease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldIncrease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


