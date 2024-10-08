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

// checks if the ProgramGatewayUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramGatewayUpdateReq{}

// ProgramGatewayUpdateReq Contains information relevant to updating a Program Gateway.
type ProgramGatewayUpdateReq struct {
	// Indicates whether the Program Gateway is active.
	Active *bool `json:"active,omitempty"`
	// Basic Authentication password for authenticating your environment.
	BasicAuthPassword *string `json:"basic_auth_password,omitempty" validate:"regexp=(?!^ +$)^.+$"`
	// Basic Authentication username for authenticating your environment.
	BasicAuthUsername *string `json:"basic_auth_username,omitempty" validate:"regexp=(?!^ +$)^.+$"`
	// Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a Program Gateway request. Custom headers also appear in the associated webhook's notifications.
	CustomHeader *map[string]string `json:"custom_header,omitempty"`
	// Indicates whether the Program Gateway uses mutual Transport Layer Security (mTLS).
	Mtls *bool `json:"mtls,omitempty"`
	// Name of the Program Gateway.
	Name *string `json:"name,omitempty"`
	// Total timeout for Program Gateway calls, in milliseconds.
	TimeoutMillis *int32 `json:"timeout_millis,omitempty"`
	// Unique identifier of the Program Gateway.
	Token *string `json:"token,omitempty"`
	// URL of the Program Gateway endpoint hosted in your environment and configured to receive authorization requests made by the Marqeta platform. Must be HTTPS.
	Url *string `json:"url,omitempty" validate:"regexp=(?!^ +$)^.+$"`
}

// NewProgramGatewayUpdateReq instantiates a new ProgramGatewayUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramGatewayUpdateReq() *ProgramGatewayUpdateReq {
	this := ProgramGatewayUpdateReq{}
	var active bool = true
	this.Active = &active
	var mtls bool = false
	this.Mtls = &mtls
	return &this
}

// NewProgramGatewayUpdateReqWithDefaults instantiates a new ProgramGatewayUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramGatewayUpdateReqWithDefaults() *ProgramGatewayUpdateReq {
	this := ProgramGatewayUpdateReq{}
	var active bool = true
	this.Active = &active
	var mtls bool = false
	this.Mtls = &mtls
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ProgramGatewayUpdateReq) SetActive(v bool) {
	o.Active = &v
}

// GetBasicAuthPassword returns the BasicAuthPassword field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetBasicAuthPassword() string {
	if o == nil || IsNil(o.BasicAuthPassword) {
		var ret string
		return ret
	}
	return *o.BasicAuthPassword
}

// GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetBasicAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthPassword) {
		return nil, false
	}
	return o.BasicAuthPassword, true
}

// HasBasicAuthPassword returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasBasicAuthPassword() bool {
	if o != nil && !IsNil(o.BasicAuthPassword) {
		return true
	}

	return false
}

// SetBasicAuthPassword gets a reference to the given string and assigns it to the BasicAuthPassword field.
func (o *ProgramGatewayUpdateReq) SetBasicAuthPassword(v string) {
	o.BasicAuthPassword = &v
}

// GetBasicAuthUsername returns the BasicAuthUsername field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetBasicAuthUsername() string {
	if o == nil || IsNil(o.BasicAuthUsername) {
		var ret string
		return ret
	}
	return *o.BasicAuthUsername
}

// GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetBasicAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthUsername) {
		return nil, false
	}
	return o.BasicAuthUsername, true
}

// HasBasicAuthUsername returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasBasicAuthUsername() bool {
	if o != nil && !IsNil(o.BasicAuthUsername) {
		return true
	}

	return false
}

// SetBasicAuthUsername gets a reference to the given string and assigns it to the BasicAuthUsername field.
func (o *ProgramGatewayUpdateReq) SetBasicAuthUsername(v string) {
	o.BasicAuthUsername = &v
}

// GetCustomHeader returns the CustomHeader field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetCustomHeader() map[string]string {
	if o == nil || IsNil(o.CustomHeader) {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeader
}

// GetCustomHeaderOk returns a tuple with the CustomHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetCustomHeaderOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomHeader) {
		return nil, false
	}
	return o.CustomHeader, true
}

// HasCustomHeader returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasCustomHeader() bool {
	if o != nil && !IsNil(o.CustomHeader) {
		return true
	}

	return false
}

// SetCustomHeader gets a reference to the given map[string]string and assigns it to the CustomHeader field.
func (o *ProgramGatewayUpdateReq) SetCustomHeader(v map[string]string) {
	o.CustomHeader = &v
}

// GetMtls returns the Mtls field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetMtls() bool {
	if o == nil || IsNil(o.Mtls) {
		var ret bool
		return ret
	}
	return *o.Mtls
}

// GetMtlsOk returns a tuple with the Mtls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetMtlsOk() (*bool, bool) {
	if o == nil || IsNil(o.Mtls) {
		return nil, false
	}
	return o.Mtls, true
}

// HasMtls returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasMtls() bool {
	if o != nil && !IsNil(o.Mtls) {
		return true
	}

	return false
}

// SetMtls gets a reference to the given bool and assigns it to the Mtls field.
func (o *ProgramGatewayUpdateReq) SetMtls(v bool) {
	o.Mtls = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProgramGatewayUpdateReq) SetName(v string) {
	o.Name = &v
}

// GetTimeoutMillis returns the TimeoutMillis field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetTimeoutMillis() int32 {
	if o == nil || IsNil(o.TimeoutMillis) {
		var ret int32
		return ret
	}
	return *o.TimeoutMillis
}

// GetTimeoutMillisOk returns a tuple with the TimeoutMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetTimeoutMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutMillis) {
		return nil, false
	}
	return o.TimeoutMillis, true
}

// HasTimeoutMillis returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasTimeoutMillis() bool {
	if o != nil && !IsNil(o.TimeoutMillis) {
		return true
	}

	return false
}

// SetTimeoutMillis gets a reference to the given int32 and assigns it to the TimeoutMillis field.
func (o *ProgramGatewayUpdateReq) SetTimeoutMillis(v int32) {
	o.TimeoutMillis = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProgramGatewayUpdateReq) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ProgramGatewayUpdateReq) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramGatewayUpdateReq) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ProgramGatewayUpdateReq) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ProgramGatewayUpdateReq) SetUrl(v string) {
	o.Url = &v
}

func (o ProgramGatewayUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramGatewayUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.BasicAuthPassword) {
		toSerialize["basic_auth_password"] = o.BasicAuthPassword
	}
	if !IsNil(o.BasicAuthUsername) {
		toSerialize["basic_auth_username"] = o.BasicAuthUsername
	}
	if !IsNil(o.CustomHeader) {
		toSerialize["custom_header"] = o.CustomHeader
	}
	if !IsNil(o.Mtls) {
		toSerialize["mtls"] = o.Mtls
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TimeoutMillis) {
		toSerialize["timeout_millis"] = o.TimeoutMillis
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableProgramGatewayUpdateReq struct {
	value *ProgramGatewayUpdateReq
	isSet bool
}

func (v NullableProgramGatewayUpdateReq) Get() *ProgramGatewayUpdateReq {
	return v.value
}

func (v *NullableProgramGatewayUpdateReq) Set(val *ProgramGatewayUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramGatewayUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramGatewayUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramGatewayUpdateReq(val *ProgramGatewayUpdateReq) *NullableProgramGatewayUpdateReq {
	return &NullableProgramGatewayUpdateReq{value: val, isSet: true}
}

func (v NullableProgramGatewayUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramGatewayUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


