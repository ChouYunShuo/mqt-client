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

// checks if the OtherPoi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtherPoi{}

// OtherPoi Allows for configuration of points of interaction other than ecommerce or ATMs, such as points of sale (POS).
type OtherPoi struct {
	// If set to `true`, card transactions at points of interaction other than e-commerce or ATMs are allowed. This group includes points of sale (POS).
	Allow *bool `json:"allow,omitempty"`
	// If set to `true`, cards of this card product type are required to be present during the transaction, such as in IVR scenarios.
	CardPresenceRequired *bool `json:"card_presence_required,omitempty"`
	// If set to `true`, the cardholder is required to be present during the transaction, such as in a restaurant where the card is present but the cardholder might not be present when the card is swiped.
	CardholderPresenceRequired *bool `json:"cardholder_presence_required,omitempty"`
	Track1DiscretionaryData *string `json:"track1_discretionary_data,omitempty"`
	Track2DiscretionaryData *string `json:"track2_discretionary_data,omitempty"`
	UseStaticPin *bool `json:"use_static_pin,omitempty"`
}

// NewOtherPoi instantiates a new OtherPoi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherPoi() *OtherPoi {
	this := OtherPoi{}
	var allow bool = true
	this.Allow = &allow
	var cardPresenceRequired bool = false
	this.CardPresenceRequired = &cardPresenceRequired
	var cardholderPresenceRequired bool = false
	this.CardholderPresenceRequired = &cardholderPresenceRequired
	var useStaticPin bool = false
	this.UseStaticPin = &useStaticPin
	return &this
}

// NewOtherPoiWithDefaults instantiates a new OtherPoi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherPoiWithDefaults() *OtherPoi {
	this := OtherPoi{}
	var allow bool = true
	this.Allow = &allow
	var cardPresenceRequired bool = false
	this.CardPresenceRequired = &cardPresenceRequired
	var cardholderPresenceRequired bool = false
	this.CardholderPresenceRequired = &cardholderPresenceRequired
	var useStaticPin bool = false
	this.UseStaticPin = &useStaticPin
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *OtherPoi) GetAllow() bool {
	if o == nil || IsNil(o.Allow) {
		var ret bool
		return ret
	}
	return *o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherPoi) GetAllowOk() (*bool, bool) {
	if o == nil || IsNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *OtherPoi) HasAllow() bool {
	if o != nil && !IsNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given bool and assigns it to the Allow field.
func (o *OtherPoi) SetAllow(v bool) {
	o.Allow = &v
}

// GetCardPresenceRequired returns the CardPresenceRequired field value if set, zero value otherwise.
func (o *OtherPoi) GetCardPresenceRequired() bool {
	if o == nil || IsNil(o.CardPresenceRequired) {
		var ret bool
		return ret
	}
	return *o.CardPresenceRequired
}

// GetCardPresenceRequiredOk returns a tuple with the CardPresenceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherPoi) GetCardPresenceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.CardPresenceRequired) {
		return nil, false
	}
	return o.CardPresenceRequired, true
}

// HasCardPresenceRequired returns a boolean if a field has been set.
func (o *OtherPoi) HasCardPresenceRequired() bool {
	if o != nil && !IsNil(o.CardPresenceRequired) {
		return true
	}

	return false
}

// SetCardPresenceRequired gets a reference to the given bool and assigns it to the CardPresenceRequired field.
func (o *OtherPoi) SetCardPresenceRequired(v bool) {
	o.CardPresenceRequired = &v
}

// GetCardholderPresenceRequired returns the CardholderPresenceRequired field value if set, zero value otherwise.
func (o *OtherPoi) GetCardholderPresenceRequired() bool {
	if o == nil || IsNil(o.CardholderPresenceRequired) {
		var ret bool
		return ret
	}
	return *o.CardholderPresenceRequired
}

// GetCardholderPresenceRequiredOk returns a tuple with the CardholderPresenceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherPoi) GetCardholderPresenceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.CardholderPresenceRequired) {
		return nil, false
	}
	return o.CardholderPresenceRequired, true
}

// HasCardholderPresenceRequired returns a boolean if a field has been set.
func (o *OtherPoi) HasCardholderPresenceRequired() bool {
	if o != nil && !IsNil(o.CardholderPresenceRequired) {
		return true
	}

	return false
}

// SetCardholderPresenceRequired gets a reference to the given bool and assigns it to the CardholderPresenceRequired field.
func (o *OtherPoi) SetCardholderPresenceRequired(v bool) {
	o.CardholderPresenceRequired = &v
}

// GetTrack1DiscretionaryData returns the Track1DiscretionaryData field value if set, zero value otherwise.
func (o *OtherPoi) GetTrack1DiscretionaryData() string {
	if o == nil || IsNil(o.Track1DiscretionaryData) {
		var ret string
		return ret
	}
	return *o.Track1DiscretionaryData
}

// GetTrack1DiscretionaryDataOk returns a tuple with the Track1DiscretionaryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherPoi) GetTrack1DiscretionaryDataOk() (*string, bool) {
	if o == nil || IsNil(o.Track1DiscretionaryData) {
		return nil, false
	}
	return o.Track1DiscretionaryData, true
}

// HasTrack1DiscretionaryData returns a boolean if a field has been set.
func (o *OtherPoi) HasTrack1DiscretionaryData() bool {
	if o != nil && !IsNil(o.Track1DiscretionaryData) {
		return true
	}

	return false
}

// SetTrack1DiscretionaryData gets a reference to the given string and assigns it to the Track1DiscretionaryData field.
func (o *OtherPoi) SetTrack1DiscretionaryData(v string) {
	o.Track1DiscretionaryData = &v
}

// GetTrack2DiscretionaryData returns the Track2DiscretionaryData field value if set, zero value otherwise.
func (o *OtherPoi) GetTrack2DiscretionaryData() string {
	if o == nil || IsNil(o.Track2DiscretionaryData) {
		var ret string
		return ret
	}
	return *o.Track2DiscretionaryData
}

// GetTrack2DiscretionaryDataOk returns a tuple with the Track2DiscretionaryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherPoi) GetTrack2DiscretionaryDataOk() (*string, bool) {
	if o == nil || IsNil(o.Track2DiscretionaryData) {
		return nil, false
	}
	return o.Track2DiscretionaryData, true
}

// HasTrack2DiscretionaryData returns a boolean if a field has been set.
func (o *OtherPoi) HasTrack2DiscretionaryData() bool {
	if o != nil && !IsNil(o.Track2DiscretionaryData) {
		return true
	}

	return false
}

// SetTrack2DiscretionaryData gets a reference to the given string and assigns it to the Track2DiscretionaryData field.
func (o *OtherPoi) SetTrack2DiscretionaryData(v string) {
	o.Track2DiscretionaryData = &v
}

// GetUseStaticPin returns the UseStaticPin field value if set, zero value otherwise.
func (o *OtherPoi) GetUseStaticPin() bool {
	if o == nil || IsNil(o.UseStaticPin) {
		var ret bool
		return ret
	}
	return *o.UseStaticPin
}

// GetUseStaticPinOk returns a tuple with the UseStaticPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherPoi) GetUseStaticPinOk() (*bool, bool) {
	if o == nil || IsNil(o.UseStaticPin) {
		return nil, false
	}
	return o.UseStaticPin, true
}

// HasUseStaticPin returns a boolean if a field has been set.
func (o *OtherPoi) HasUseStaticPin() bool {
	if o != nil && !IsNil(o.UseStaticPin) {
		return true
	}

	return false
}

// SetUseStaticPin gets a reference to the given bool and assigns it to the UseStaticPin field.
func (o *OtherPoi) SetUseStaticPin(v bool) {
	o.UseStaticPin = &v
}

func (o OtherPoi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtherPoi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allow) {
		toSerialize["allow"] = o.Allow
	}
	if !IsNil(o.CardPresenceRequired) {
		toSerialize["card_presence_required"] = o.CardPresenceRequired
	}
	if !IsNil(o.CardholderPresenceRequired) {
		toSerialize["cardholder_presence_required"] = o.CardholderPresenceRequired
	}
	if !IsNil(o.Track1DiscretionaryData) {
		toSerialize["track1_discretionary_data"] = o.Track1DiscretionaryData
	}
	if !IsNil(o.Track2DiscretionaryData) {
		toSerialize["track2_discretionary_data"] = o.Track2DiscretionaryData
	}
	if !IsNil(o.UseStaticPin) {
		toSerialize["use_static_pin"] = o.UseStaticPin
	}
	return toSerialize, nil
}

type NullableOtherPoi struct {
	value *OtherPoi
	isSet bool
}

func (v NullableOtherPoi) Get() *OtherPoi {
	return v.value
}

func (v *NullableOtherPoi) Set(val *OtherPoi) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherPoi) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherPoi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherPoi(val *OtherPoi) *NullableOtherPoi {
	return &NullableOtherPoi{value: val, isSet: true}
}

func (v NullableOtherPoi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherPoi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


