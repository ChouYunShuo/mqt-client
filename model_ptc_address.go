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

// checks if the PTCAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PTCAddress{}

// PTCAddress struct for PTCAddress
type PTCAddress struct {
	City string `json:"city"`
	Country *string `json:"country,omitempty"`
	County string `json:"county"`
	Line1 string `json:"line1"`
	Line2 *string `json:"line2,omitempty"`
	PostalCode string `json:"postal_code"`
	State string `json:"state"`
}

type _PTCAddress PTCAddress

// NewPTCAddress instantiates a new PTCAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPTCAddress(city string, county string, line1 string, postalCode string, state string) *PTCAddress {
	this := PTCAddress{}
	this.City = city
	this.County = county
	this.Line1 = line1
	this.PostalCode = postalCode
	this.State = state
	return &this
}

// NewPTCAddressWithDefaults instantiates a new PTCAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPTCAddressWithDefaults() *PTCAddress {
	this := PTCAddress{}
	return &this
}

// GetCity returns the City field value
func (o *PTCAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *PTCAddress) SetCity(v string) {
	o.City = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PTCAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PTCAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PTCAddress) SetCountry(v string) {
	o.Country = &v
}

// GetCounty returns the County field value
func (o *PTCAddress) GetCounty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.County
}

// GetCountyOk returns a tuple with the County field value
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetCountyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.County, true
}

// SetCounty sets field value
func (o *PTCAddress) SetCounty(v string) {
	o.County = v
}

// GetLine1 returns the Line1 field value
func (o *PTCAddress) GetLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line1, true
}

// SetLine1 sets field value
func (o *PTCAddress) SetLine1(v string) {
	o.Line1 = v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *PTCAddress) GetLine2() string {
	if o == nil || IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *PTCAddress) HasLine2() bool {
	if o != nil && !IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *PTCAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetPostalCode returns the PostalCode field value
func (o *PTCAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *PTCAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetState returns the State field value
func (o *PTCAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PTCAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PTCAddress) SetState(v string) {
	o.State = v
}

func (o PTCAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PTCAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	toSerialize["county"] = o.County
	toSerialize["line1"] = o.Line1
	if !IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	toSerialize["postal_code"] = o.PostalCode
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *PTCAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"county",
		"line1",
		"postal_code",
		"state",
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

	varPTCAddress := _PTCAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPTCAddress)

	if err != nil {
		return err
	}

	*o = PTCAddress(varPTCAddress)

	return err
}

type NullablePTCAddress struct {
	value *PTCAddress
	isSet bool
}

func (v NullablePTCAddress) Get() *PTCAddress {
	return v.value
}

func (v *NullablePTCAddress) Set(val *PTCAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePTCAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePTCAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePTCAddress(val *PTCAddress) *NullablePTCAddress {
	return &NullablePTCAddress{value: val, isSet: true}
}

func (v NullablePTCAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePTCAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


