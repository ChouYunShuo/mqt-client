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
)

// checks if the AddressResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressResponseModel{}

// AddressResponseModel Address associated with the business.
type AddressResponseModel struct {
	// Street name and number of the address.  This field is returned if it exists in the resource.
	Address1 *string `json:"address1,omitempty"`
	// Additional address information.  This field is returned if it exists in the resource.
	Address2 *string `json:"address2,omitempty"`
	// City of the address.  This field is returned if it exists in the resource.
	City *string `json:"city,omitempty"`
	// Country of the address.  This field is returned if it exists in the resource.
	Country *string `json:"country,omitempty"`
	// Postal code of the address.  This field is returned if it exists in the resource.
	PostalCode *string `json:"postal_code,omitempty"`
	// State, province, or territory of the address.  This field is returned if it exists in the resource.
	State *string `json:"state,omitempty"`
	// United States ZIP code of the address.  This field is returned if it exists in the resource.
	Zip *string `json:"zip,omitempty"`
}

// NewAddressResponseModel instantiates a new AddressResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressResponseModel() *AddressResponseModel {
	this := AddressResponseModel{}
	return &this
}

// NewAddressResponseModelWithDefaults instantiates a new AddressResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressResponseModelWithDefaults() *AddressResponseModel {
	this := AddressResponseModel{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *AddressResponseModel) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *AddressResponseModel) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *AddressResponseModel) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *AddressResponseModel) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *AddressResponseModel) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *AddressResponseModel) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressResponseModel) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressResponseModel) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressResponseModel) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressResponseModel) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressResponseModel) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressResponseModel) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressResponseModel) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressResponseModel) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressResponseModel) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressResponseModel) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressResponseModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressResponseModel) SetState(v string) {
	o.State = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *AddressResponseModel) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseModel) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *AddressResponseModel) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *AddressResponseModel) SetZip(v string) {
	o.Zip = &v
}

func (o AddressResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableAddressResponseModel struct {
	value *AddressResponseModel
	isSet bool
}

func (v NullableAddressResponseModel) Get() *AddressResponseModel {
	return v.value
}

func (v *NullableAddressResponseModel) Set(val *AddressResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressResponseModel(val *AddressResponseModel) *NullableAddressResponseModel {
	return &NullableAddressResponseModel{value: val, isSet: true}
}

func (v NullableAddressResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


