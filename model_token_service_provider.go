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

// checks if the TokenServiceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenServiceProvider{}

// TokenServiceProvider Contains information held and provided by the token service provider (card network).
type TokenServiceProvider struct {
	// Unique value representing a tokenization request (Mastercard only).
	CorrelationId *string `json:"correlation_id,omitempty"`
	// Unique identifier of the digital wallet token primary account number (PAN) within the card network.
	PanReferenceId string `json:"pan_reference_id"`
	// _(Mastercard only)_ Represents the confidence level in the digital wallet token.
	TokenAssuranceLevel *string `json:"token_assurance_level,omitempty"`
	// Digital wallet's decision as to whether the digital wallet token should be provisioned.
	TokenEligibilityDecision *string `json:"token_eligibility_decision,omitempty"`
	// Expiration date of the digital wallet token.
	TokenExpiration *string `json:"token_expiration,omitempty"`
	// Primary account number (PAN) of the digital wallet token.
	TokenPan *string `json:"token_pan,omitempty"`
	// Unique identifier of the digital wallet token within the card network.
	TokenReferenceId *string `json:"token_reference_id,omitempty"`
	// Unique numerical identifier of the token requestor within the card network. These ID numbers map to `token_requestor_name` field values as follows:  *Mastercard*  * 50110030273 – `APPLE_PAY` * 50120834693 – `ANDROID_PAY` * 50139059239 – `SAMSUNG_PAY`  *Visa*  * 40010030273 – `APPLE_PAY` * 40010075001 – `ANDROID_PAY` * 40010043095 – `SAMSUNG_PAY` * 40010075196 – `MICROSOFT_PAY` * 40010075338 – `VISA_CHECKOUT` * 40010075449 – `FACEBOOK` * 40010075839 – `NETFLIX` * 40010077056 – `FITBIT_PAY` * 40010069887 – `GARMIN_PAY`
	TokenRequestorId *string `json:"token_requestor_id,omitempty"`
	// Name of the token requestor within the card network.  *NOTE:* The list of example values for this field is maintained by the card networks and is subject to change.
	TokenRequestorName *string `json:"token_requestor_name,omitempty"`
	// Token score assigned by the digital wallet.
	TokenScore *string `json:"token_score,omitempty"`
	// Type of the digital wallet token.
	TokenType *string `json:"token_type,omitempty"`
}

type _TokenServiceProvider TokenServiceProvider

// NewTokenServiceProvider instantiates a new TokenServiceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenServiceProvider(panReferenceId string) *TokenServiceProvider {
	this := TokenServiceProvider{}
	this.PanReferenceId = panReferenceId
	return &this
}

// NewTokenServiceProviderWithDefaults instantiates a new TokenServiceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenServiceProviderWithDefaults() *TokenServiceProvider {
	this := TokenServiceProvider{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *TokenServiceProvider) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetPanReferenceId returns the PanReferenceId field value
func (o *TokenServiceProvider) GetPanReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PanReferenceId
}

// GetPanReferenceIdOk returns a tuple with the PanReferenceId field value
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetPanReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PanReferenceId, true
}

// SetPanReferenceId sets field value
func (o *TokenServiceProvider) SetPanReferenceId(v string) {
	o.PanReferenceId = v
}

// GetTokenAssuranceLevel returns the TokenAssuranceLevel field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenAssuranceLevel() string {
	if o == nil || IsNil(o.TokenAssuranceLevel) {
		var ret string
		return ret
	}
	return *o.TokenAssuranceLevel
}

// GetTokenAssuranceLevelOk returns a tuple with the TokenAssuranceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenAssuranceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.TokenAssuranceLevel) {
		return nil, false
	}
	return o.TokenAssuranceLevel, true
}

// HasTokenAssuranceLevel returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenAssuranceLevel() bool {
	if o != nil && !IsNil(o.TokenAssuranceLevel) {
		return true
	}

	return false
}

// SetTokenAssuranceLevel gets a reference to the given string and assigns it to the TokenAssuranceLevel field.
func (o *TokenServiceProvider) SetTokenAssuranceLevel(v string) {
	o.TokenAssuranceLevel = &v
}

// GetTokenEligibilityDecision returns the TokenEligibilityDecision field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenEligibilityDecision() string {
	if o == nil || IsNil(o.TokenEligibilityDecision) {
		var ret string
		return ret
	}
	return *o.TokenEligibilityDecision
}

// GetTokenEligibilityDecisionOk returns a tuple with the TokenEligibilityDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenEligibilityDecisionOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEligibilityDecision) {
		return nil, false
	}
	return o.TokenEligibilityDecision, true
}

// HasTokenEligibilityDecision returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenEligibilityDecision() bool {
	if o != nil && !IsNil(o.TokenEligibilityDecision) {
		return true
	}

	return false
}

// SetTokenEligibilityDecision gets a reference to the given string and assigns it to the TokenEligibilityDecision field.
func (o *TokenServiceProvider) SetTokenEligibilityDecision(v string) {
	o.TokenEligibilityDecision = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenExpiration() string {
	if o == nil || IsNil(o.TokenExpiration) {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.TokenExpiration) {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenExpiration() bool {
	if o != nil && !IsNil(o.TokenExpiration) {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *TokenServiceProvider) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetTokenPan returns the TokenPan field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenPan() string {
	if o == nil || IsNil(o.TokenPan) {
		var ret string
		return ret
	}
	return *o.TokenPan
}

// GetTokenPanOk returns a tuple with the TokenPan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenPanOk() (*string, bool) {
	if o == nil || IsNil(o.TokenPan) {
		return nil, false
	}
	return o.TokenPan, true
}

// HasTokenPan returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenPan() bool {
	if o != nil && !IsNil(o.TokenPan) {
		return true
	}

	return false
}

// SetTokenPan gets a reference to the given string and assigns it to the TokenPan field.
func (o *TokenServiceProvider) SetTokenPan(v string) {
	o.TokenPan = &v
}

// GetTokenReferenceId returns the TokenReferenceId field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenReferenceId() string {
	if o == nil || IsNil(o.TokenReferenceId) {
		var ret string
		return ret
	}
	return *o.TokenReferenceId
}

// GetTokenReferenceIdOk returns a tuple with the TokenReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenReferenceId) {
		return nil, false
	}
	return o.TokenReferenceId, true
}

// HasTokenReferenceId returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenReferenceId() bool {
	if o != nil && !IsNil(o.TokenReferenceId) {
		return true
	}

	return false
}

// SetTokenReferenceId gets a reference to the given string and assigns it to the TokenReferenceId field.
func (o *TokenServiceProvider) SetTokenReferenceId(v string) {
	o.TokenReferenceId = &v
}

// GetTokenRequestorId returns the TokenRequestorId field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenRequestorId() string {
	if o == nil || IsNil(o.TokenRequestorId) {
		var ret string
		return ret
	}
	return *o.TokenRequestorId
}

// GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenRequestorIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenRequestorId) {
		return nil, false
	}
	return o.TokenRequestorId, true
}

// HasTokenRequestorId returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenRequestorId() bool {
	if o != nil && !IsNil(o.TokenRequestorId) {
		return true
	}

	return false
}

// SetTokenRequestorId gets a reference to the given string and assigns it to the TokenRequestorId field.
func (o *TokenServiceProvider) SetTokenRequestorId(v string) {
	o.TokenRequestorId = &v
}

// GetTokenRequestorName returns the TokenRequestorName field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenRequestorName() string {
	if o == nil || IsNil(o.TokenRequestorName) {
		var ret string
		return ret
	}
	return *o.TokenRequestorName
}

// GetTokenRequestorNameOk returns a tuple with the TokenRequestorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenRequestorNameOk() (*string, bool) {
	if o == nil || IsNil(o.TokenRequestorName) {
		return nil, false
	}
	return o.TokenRequestorName, true
}

// HasTokenRequestorName returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenRequestorName() bool {
	if o != nil && !IsNil(o.TokenRequestorName) {
		return true
	}

	return false
}

// SetTokenRequestorName gets a reference to the given string and assigns it to the TokenRequestorName field.
func (o *TokenServiceProvider) SetTokenRequestorName(v string) {
	o.TokenRequestorName = &v
}

// GetTokenScore returns the TokenScore field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenScore() string {
	if o == nil || IsNil(o.TokenScore) {
		var ret string
		return ret
	}
	return *o.TokenScore
}

// GetTokenScoreOk returns a tuple with the TokenScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenScoreOk() (*string, bool) {
	if o == nil || IsNil(o.TokenScore) {
		return nil, false
	}
	return o.TokenScore, true
}

// HasTokenScore returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenScore() bool {
	if o != nil && !IsNil(o.TokenScore) {
		return true
	}

	return false
}

// SetTokenScore gets a reference to the given string and assigns it to the TokenScore field.
func (o *TokenServiceProvider) SetTokenScore(v string) {
	o.TokenScore = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenServiceProvider) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenServiceProvider) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenServiceProvider) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenServiceProvider) SetTokenType(v string) {
	o.TokenType = &v
}

func (o TokenServiceProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenServiceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlation_id"] = o.CorrelationId
	}
	toSerialize["pan_reference_id"] = o.PanReferenceId
	if !IsNil(o.TokenAssuranceLevel) {
		toSerialize["token_assurance_level"] = o.TokenAssuranceLevel
	}
	if !IsNil(o.TokenEligibilityDecision) {
		toSerialize["token_eligibility_decision"] = o.TokenEligibilityDecision
	}
	if !IsNil(o.TokenExpiration) {
		toSerialize["token_expiration"] = o.TokenExpiration
	}
	if !IsNil(o.TokenPan) {
		toSerialize["token_pan"] = o.TokenPan
	}
	if !IsNil(o.TokenReferenceId) {
		toSerialize["token_reference_id"] = o.TokenReferenceId
	}
	if !IsNil(o.TokenRequestorId) {
		toSerialize["token_requestor_id"] = o.TokenRequestorId
	}
	if !IsNil(o.TokenRequestorName) {
		toSerialize["token_requestor_name"] = o.TokenRequestorName
	}
	if !IsNil(o.TokenScore) {
		toSerialize["token_score"] = o.TokenScore
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	return toSerialize, nil
}

func (o *TokenServiceProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pan_reference_id",
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

	varTokenServiceProvider := _TokenServiceProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenServiceProvider)

	if err != nil {
		return err
	}

	*o = TokenServiceProvider(varTokenServiceProvider)

	return err
}

type NullableTokenServiceProvider struct {
	value *TokenServiceProvider
	isSet bool
}

func (v NullableTokenServiceProvider) Get() *TokenServiceProvider {
	return v.value
}

func (v *NullableTokenServiceProvider) Set(val *TokenServiceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenServiceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenServiceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenServiceProvider(val *TokenServiceProvider) *NullableTokenServiceProvider {
	return &NullableTokenServiceProvider{value: val, isSet: true}
}

func (v NullableTokenServiceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenServiceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


