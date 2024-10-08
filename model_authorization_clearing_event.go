/*
Marqeta Core API

Simplified management of your payment programs

API version: 3.0.0
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AuthorizationClearingEvent struct for AuthorizationClearingEvent
type AuthorizationClearingEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	PrecedingRelatedTransactionToken string `json:"preceding_related_transaction_token"`
}

// NewAuthorizationClearingEvent instantiates a new AuthorizationClearingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationClearingEvent(precedingRelatedTransactionToken string, amount float32) *AuthorizationClearingEvent {
	this := AuthorizationClearingEvent{}
	this.Amount = &amount
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	this.PrecedingRelatedTransactionToken = precedingRelatedTransactionToken
	return &this
}

// NewAuthorizationClearingEventWithDefaults instantiates a new AuthorizationClearingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationClearingEventWithDefaults() *AuthorizationClearingEvent {
	this := AuthorizationClearingEvent{}
	return &this
}

// GetPrecedingRelatedTransactionToken returns the PrecedingRelatedTransactionToken field value
func (o *AuthorizationClearingEvent) GetPrecedingRelatedTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrecedingRelatedTransactionToken
}

// GetPrecedingRelatedTransactionTokenOk returns a tuple with the PrecedingRelatedTransactionToken field value
// and a boolean to check if the value has been set.
func (o *AuthorizationClearingEvent) GetPrecedingRelatedTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecedingRelatedTransactionToken, true
}

// SetPrecedingRelatedTransactionToken sets field value
func (o *AuthorizationClearingEvent) SetPrecedingRelatedTransactionToken(v string) {
	o.PrecedingRelatedTransactionToken = v
}

func (o AuthorizationClearingEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCardTransactionSimulation, errCardTransactionSimulation := json.Marshal(o.CardTransactionSimulation)
	if errCardTransactionSimulation != nil {
		return []byte{}, errCardTransactionSimulation
	}
	errCardTransactionSimulation = json.Unmarshal([]byte(serializedCardTransactionSimulation), &toSerialize)
	if errCardTransactionSimulation != nil {
		return []byte{}, errCardTransactionSimulation
	}
	if true {
		toSerialize["preceding_related_transaction_token"] = o.PrecedingRelatedTransactionToken
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationClearingEvent struct {
	value *AuthorizationClearingEvent
	isSet bool
}

func (v NullableAuthorizationClearingEvent) Get() *AuthorizationClearingEvent {
	return v.value
}

func (v *NullableAuthorizationClearingEvent) Set(val *AuthorizationClearingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationClearingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationClearingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationClearingEvent(val *AuthorizationClearingEvent) *NullableAuthorizationClearingEvent {
	return &NullableAuthorizationClearingEvent{value: val, isSet: true}
}

func (v NullableAuthorizationClearingEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationClearingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


