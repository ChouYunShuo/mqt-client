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

// checks if the KycRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KycRequest{}

// KycRequest struct for KycRequest
type KycRequest struct {
	// Specifies the business account holder on which to perform the identity check. Do not pass this field if your request includes the `user_token` field.  Send a `GET` request to `/businesses` to retrieve business tokens.
	BusinessToken *string `json:"business_token,omitempty"`
	// Set to `true` to designate a user account holder as having passed a verification check without Marqeta performing the check through one of its KYC providers.  Use this override when you perform verification through another mechanism, such as an alternative KYC provider or directly with the account holder.  You must obtain explicit, written approval from Marqeta before using the `manual_override` field for KYC verification. This feature is only available to programs that are enabled to perform their own Customer Identification Program (CIP) KYC verification. Consult your Marqeta representative for more information.
	ManualOverride *bool `json:"manual_override,omitempty"`
	// Notes pertaining to a manual override. This field is returned in the response only when the `manual_override` field is set to `true`.
	Notes *string `json:"notes,omitempty"`
	// Can be specified only if `manual_override=true`. If you verified a user account holder's identity by performing a KYC verification outside of the Marqeta platform, you can use the `reference_id` field to store the reference number returned by that KYC provider.  *NOTE:* When you submit a KYC verification request with `manual_override=false`, the Marqeta platform performs the verification through one of its KYC providers. If the KYC provider responds with a reference identifier, that identifier is passed to you by way of this field in the response.
	ReferenceId *string `json:"reference_id,omitempty"`
	// The unique identifier of the identity check.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// Specifies the user account holder on which to perform the identity check. Do not pass this field if your request includes the `business_token` field.  Send a `GET` request to `/users` to retrieve user tokens.
	UserToken *string `json:"user_token,omitempty"`
}

// NewKycRequest instantiates a new KycRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKycRequest() *KycRequest {
	this := KycRequest{}
	var manualOverride bool = false
	this.ManualOverride = &manualOverride
	return &this
}

// NewKycRequestWithDefaults instantiates a new KycRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKycRequestWithDefaults() *KycRequest {
	this := KycRequest{}
	var manualOverride bool = false
	this.ManualOverride = &manualOverride
	return &this
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *KycRequest) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycRequest) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *KycRequest) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *KycRequest) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetManualOverride returns the ManualOverride field value if set, zero value otherwise.
func (o *KycRequest) GetManualOverride() bool {
	if o == nil || IsNil(o.ManualOverride) {
		var ret bool
		return ret
	}
	return *o.ManualOverride
}

// GetManualOverrideOk returns a tuple with the ManualOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycRequest) GetManualOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualOverride) {
		return nil, false
	}
	return o.ManualOverride, true
}

// HasManualOverride returns a boolean if a field has been set.
func (o *KycRequest) HasManualOverride() bool {
	if o != nil && !IsNil(o.ManualOverride) {
		return true
	}

	return false
}

// SetManualOverride gets a reference to the given bool and assigns it to the ManualOverride field.
func (o *KycRequest) SetManualOverride(v bool) {
	o.ManualOverride = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *KycRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *KycRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *KycRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *KycRequest) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *KycRequest) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *KycRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KycRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KycRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KycRequest) SetToken(v string) {
	o.Token = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *KycRequest) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *KycRequest) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *KycRequest) SetUserToken(v string) {
	o.UserToken = &v
}

func (o KycRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KycRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	if !IsNil(o.ManualOverride) {
		toSerialize["manual_override"] = o.ManualOverride
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

type NullableKycRequest struct {
	value *KycRequest
	isSet bool
}

func (v NullableKycRequest) Get() *KycRequest {
	return v.value
}

func (v *NullableKycRequest) Set(val *KycRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKycRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKycRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKycRequest(val *KycRequest) *NullableKycRequest {
	return &NullableKycRequest{value: val, isSet: true}
}

func (v NullableKycRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKycRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


