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
)

// checks if the ConfigFeeScheduleReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigFeeScheduleReq{}

// ConfigFeeScheduleReq Contains information relevant to configuring fees.
type ConfigFeeScheduleReq struct {
	// Contains one or more fee schedules.
	Schedule []ConfigFeeScheduleEntry `json:"schedule"`
	Type AccountProductFeeType `json:"type"`
}

type _ConfigFeeScheduleReq ConfigFeeScheduleReq

// NewConfigFeeScheduleReq instantiates a new ConfigFeeScheduleReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigFeeScheduleReq(schedule []ConfigFeeScheduleEntry, type_ AccountProductFeeType) *ConfigFeeScheduleReq {
	this := ConfigFeeScheduleReq{}
	this.Schedule = schedule
	this.Type = type_
	return &this
}

// NewConfigFeeScheduleReqWithDefaults instantiates a new ConfigFeeScheduleReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigFeeScheduleReqWithDefaults() *ConfigFeeScheduleReq {
	this := ConfigFeeScheduleReq{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *ConfigFeeScheduleReq) GetSchedule() []ConfigFeeScheduleEntry {
	if o == nil {
		var ret []ConfigFeeScheduleEntry
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ConfigFeeScheduleReq) GetScheduleOk() ([]ConfigFeeScheduleEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule, true
}

// SetSchedule sets field value
func (o *ConfigFeeScheduleReq) SetSchedule(v []ConfigFeeScheduleEntry) {
	o.Schedule = v
}

// GetType returns the Type field value
func (o *ConfigFeeScheduleReq) GetType() AccountProductFeeType {
	if o == nil {
		var ret AccountProductFeeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigFeeScheduleReq) GetTypeOk() (*AccountProductFeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfigFeeScheduleReq) SetType(v AccountProductFeeType) {
	o.Type = v
}

func (o ConfigFeeScheduleReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigFeeScheduleReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schedule"] = o.Schedule
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ConfigFeeScheduleReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schedule",
		"type",
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

	varConfigFeeScheduleReq := _ConfigFeeScheduleReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigFeeScheduleReq)

	if err != nil {
		return err
	}

	*o = ConfigFeeScheduleReq(varConfigFeeScheduleReq)

	return err
}

type NullableConfigFeeScheduleReq struct {
	value *ConfigFeeScheduleReq
	isSet bool
}

func (v NullableConfigFeeScheduleReq) Get() *ConfigFeeScheduleReq {
	return v.value
}

func (v *NullableConfigFeeScheduleReq) Set(val *ConfigFeeScheduleReq) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigFeeScheduleReq) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigFeeScheduleReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigFeeScheduleReq(val *ConfigFeeScheduleReq) *NullableConfigFeeScheduleReq {
	return &NullableConfigFeeScheduleReq{value: val, isSet: true}
}

func (v NullableConfigFeeScheduleReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigFeeScheduleReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


