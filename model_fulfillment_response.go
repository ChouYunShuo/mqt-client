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

// checks if the FulfillmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentResponse{}

// FulfillmentResponse Specifies certain physical characteristics of a card, as well as bulk shipment information.  This object is returned if it exists in the resource.
type FulfillmentResponse struct {
	CardPersonalization CardPersonalization `json:"card_personalization"`
	Shipping *ShippingInformationResponse `json:"shipping,omitempty"`
}

type _FulfillmentResponse FulfillmentResponse

// NewFulfillmentResponse instantiates a new FulfillmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentResponse(cardPersonalization CardPersonalization) *FulfillmentResponse {
	this := FulfillmentResponse{}
	this.CardPersonalization = cardPersonalization
	return &this
}

// NewFulfillmentResponseWithDefaults instantiates a new FulfillmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentResponseWithDefaults() *FulfillmentResponse {
	this := FulfillmentResponse{}
	return &this
}

// GetCardPersonalization returns the CardPersonalization field value
func (o *FulfillmentResponse) GetCardPersonalization() CardPersonalization {
	if o == nil {
		var ret CardPersonalization
		return ret
	}

	return o.CardPersonalization
}

// GetCardPersonalizationOk returns a tuple with the CardPersonalization field value
// and a boolean to check if the value has been set.
func (o *FulfillmentResponse) GetCardPersonalizationOk() (*CardPersonalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardPersonalization, true
}

// SetCardPersonalization sets field value
func (o *FulfillmentResponse) SetCardPersonalization(v CardPersonalization) {
	o.CardPersonalization = v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *FulfillmentResponse) GetShipping() ShippingInformationResponse {
	if o == nil || IsNil(o.Shipping) {
		var ret ShippingInformationResponse
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentResponse) GetShippingOk() (*ShippingInformationResponse, bool) {
	if o == nil || IsNil(o.Shipping) {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *FulfillmentResponse) HasShipping() bool {
	if o != nil && !IsNil(o.Shipping) {
		return true
	}

	return false
}

// SetShipping gets a reference to the given ShippingInformationResponse and assigns it to the Shipping field.
func (o *FulfillmentResponse) SetShipping(v ShippingInformationResponse) {
	o.Shipping = &v
}

func (o FulfillmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfillmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_personalization"] = o.CardPersonalization
	if !IsNil(o.Shipping) {
		toSerialize["shipping"] = o.Shipping
	}
	return toSerialize, nil
}

func (o *FulfillmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_personalization",
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

	varFulfillmentResponse := _FulfillmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFulfillmentResponse)

	if err != nil {
		return err
	}

	*o = FulfillmentResponse(varFulfillmentResponse)

	return err
}

type NullableFulfillmentResponse struct {
	value *FulfillmentResponse
	isSet bool
}

func (v NullableFulfillmentResponse) Get() *FulfillmentResponse {
	return v.value
}

func (v *NullableFulfillmentResponse) Set(val *FulfillmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentResponse(val *FulfillmentResponse) *NullableFulfillmentResponse {
	return &NullableFulfillmentResponse{value: val, isSet: true}
}

func (v NullableFulfillmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


