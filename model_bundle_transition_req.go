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

// checks if the BundleTransitionReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleTransitionReq{}

// BundleTransitionReq Contains information on a bundle transition.
type BundleTransitionReq struct {
	Status BundleResourceStatus `json:"status"`
}

type _BundleTransitionReq BundleTransitionReq

// NewBundleTransitionReq instantiates a new BundleTransitionReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleTransitionReq(status BundleResourceStatus) *BundleTransitionReq {
	this := BundleTransitionReq{}
	this.Status = status
	return &this
}

// NewBundleTransitionReqWithDefaults instantiates a new BundleTransitionReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleTransitionReqWithDefaults() *BundleTransitionReq {
	this := BundleTransitionReq{}
	return &this
}

// GetStatus returns the Status field value
func (o *BundleTransitionReq) GetStatus() BundleResourceStatus {
	if o == nil {
		var ret BundleResourceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BundleTransitionReq) GetStatusOk() (*BundleResourceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BundleTransitionReq) SetStatus(v BundleResourceStatus) {
	o.Status = v
}

func (o BundleTransitionReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleTransitionReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *BundleTransitionReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varBundleTransitionReq := _BundleTransitionReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBundleTransitionReq)

	if err != nil {
		return err
	}

	*o = BundleTransitionReq(varBundleTransitionReq)

	return err
}

type NullableBundleTransitionReq struct {
	value *BundleTransitionReq
	isSet bool
}

func (v NullableBundleTransitionReq) Get() *BundleTransitionReq {
	return v.value
}

func (v *NullableBundleTransitionReq) Set(val *BundleTransitionReq) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleTransitionReq) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleTransitionReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleTransitionReq(val *BundleTransitionReq) *NullableBundleTransitionReq {
	return &NullableBundleTransitionReq{value: val, isSet: true}
}

func (v NullableBundleTransitionReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleTransitionReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


