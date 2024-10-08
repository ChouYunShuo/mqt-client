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

// AuthorizationAdviceEvent struct for AuthorizationAdviceEvent
type AuthorizationAdviceEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	PrecedingRelatedTransactionToken string `json:"preceding_related_transaction_token"`
}

// NewAuthorizationAdviceEvent instantiates a new AuthorizationAdviceEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationAdviceEvent(precedingRelatedTransactionToken string, amount float32) *AuthorizationAdviceEvent {
	this := AuthorizationAdviceEvent{}
	this.Amount = &amount
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	this.PrecedingRelatedTransactionToken = precedingRelatedTransactionToken
	return &this
}

// NewAuthorizationAdviceEventWithDefaults instantiates a new AuthorizationAdviceEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationAdviceEventWithDefaults() *AuthorizationAdviceEvent {
	this := AuthorizationAdviceEvent{}
	return &this
}

// GetPrecedingRelatedTransactionToken returns the PrecedingRelatedTransactionToken field value
func (o *AuthorizationAdviceEvent) GetPrecedingRelatedTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrecedingRelatedTransactionToken
}

// GetPrecedingRelatedTransactionTokenOk returns a tuple with the PrecedingRelatedTransactionToken field value
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceEvent) GetPrecedingRelatedTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecedingRelatedTransactionToken, true
}

// SetPrecedingRelatedTransactionToken sets field value
func (o *AuthorizationAdviceEvent) SetPrecedingRelatedTransactionToken(v string) {
	o.PrecedingRelatedTransactionToken = v
}

func (o AuthorizationAdviceEvent) MarshalJSON() ([]byte, error) {
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

type NullableAuthorizationAdviceEvent struct {
	value *AuthorizationAdviceEvent
	isSet bool
}

func (v NullableAuthorizationAdviceEvent) Get() *AuthorizationAdviceEvent {
	return v.value
}

func (v *NullableAuthorizationAdviceEvent) Set(val *AuthorizationAdviceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationAdviceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationAdviceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationAdviceEvent(val *AuthorizationAdviceEvent) *NullableAuthorizationAdviceEvent {
	return &NullableAuthorizationAdviceEvent{value: val, isSet: true}
}

func (v NullableAuthorizationAdviceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationAdviceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


