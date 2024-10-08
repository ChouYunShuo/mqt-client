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

// checks if the GatewayProgramFundingSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayProgramFundingSourceRequest{}

// GatewayProgramFundingSourceRequest struct for GatewayProgramFundingSourceRequest
type GatewayProgramFundingSourceRequest struct {
	// Indicates whether the program gateway funding source is active.
	Active *bool `json:"active,omitempty"`
	// Password for authenticating your environment.
	BasicAuthPassword string `json:"basic_auth_password"`
	// Username for authenticating your environment.
	BasicAuthUsername string `json:"basic_auth_username"`
	// Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a JIT Funding request. Custom headers also appear in the associated webhook's notifications.
	CustomHeader *map[string]string `json:"custom_header,omitempty"`
	// Name of the program gateway funding source.
	Name string `json:"name"`
	// Total timeout in milliseconds for gateway processing.
	TimeoutMillis *int64 `json:"timeout_millis,omitempty"`
	// Unique identifier of the program gateway funding source. If you do not include a token, the system will generate one automatically. As this token is necessary for use in other calls, we recommend that you define a simple and easy to remember string rather than letting the system generate a token for you. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// URL of the gateway endpoint hosted in your environment, to which `POST` requests are submitted by the Marqeta platform.
	Url string `json:"url"`
	// Specifies whether or not to use mutual transport layer security (mTLS) authentication for the funding request.
	UseMtls *bool `json:"use_mtls,omitempty"`
}

type _GatewayProgramFundingSourceRequest GatewayProgramFundingSourceRequest

// NewGatewayProgramFundingSourceRequest instantiates a new GatewayProgramFundingSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayProgramFundingSourceRequest(basicAuthPassword string, basicAuthUsername string, name string, url string) *GatewayProgramFundingSourceRequest {
	this := GatewayProgramFundingSourceRequest{}
	this.BasicAuthPassword = basicAuthPassword
	this.BasicAuthUsername = basicAuthUsername
	this.Name = name
	this.Url = url
	var useMtls bool = false
	this.UseMtls = &useMtls
	return &this
}

// NewGatewayProgramFundingSourceRequestWithDefaults instantiates a new GatewayProgramFundingSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayProgramFundingSourceRequestWithDefaults() *GatewayProgramFundingSourceRequest {
	this := GatewayProgramFundingSourceRequest{}
	var useMtls bool = false
	this.UseMtls = &useMtls
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GatewayProgramFundingSourceRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GatewayProgramFundingSourceRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GatewayProgramFundingSourceRequest) SetActive(v bool) {
	o.Active = &v
}

// GetBasicAuthPassword returns the BasicAuthPassword field value
func (o *GatewayProgramFundingSourceRequest) GetBasicAuthPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasicAuthPassword
}

// GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetBasicAuthPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicAuthPassword, true
}

// SetBasicAuthPassword sets field value
func (o *GatewayProgramFundingSourceRequest) SetBasicAuthPassword(v string) {
	o.BasicAuthPassword = v
}

// GetBasicAuthUsername returns the BasicAuthUsername field value
func (o *GatewayProgramFundingSourceRequest) GetBasicAuthUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasicAuthUsername
}

// GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetBasicAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicAuthUsername, true
}

// SetBasicAuthUsername sets field value
func (o *GatewayProgramFundingSourceRequest) SetBasicAuthUsername(v string) {
	o.BasicAuthUsername = v
}

// GetCustomHeader returns the CustomHeader field value if set, zero value otherwise.
func (o *GatewayProgramFundingSourceRequest) GetCustomHeader() map[string]string {
	if o == nil || IsNil(o.CustomHeader) {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeader
}

// GetCustomHeaderOk returns a tuple with the CustomHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetCustomHeaderOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomHeader) {
		return nil, false
	}
	return o.CustomHeader, true
}

// HasCustomHeader returns a boolean if a field has been set.
func (o *GatewayProgramFundingSourceRequest) HasCustomHeader() bool {
	if o != nil && !IsNil(o.CustomHeader) {
		return true
	}

	return false
}

// SetCustomHeader gets a reference to the given map[string]string and assigns it to the CustomHeader field.
func (o *GatewayProgramFundingSourceRequest) SetCustomHeader(v map[string]string) {
	o.CustomHeader = &v
}

// GetName returns the Name field value
func (o *GatewayProgramFundingSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayProgramFundingSourceRequest) SetName(v string) {
	o.Name = v
}

// GetTimeoutMillis returns the TimeoutMillis field value if set, zero value otherwise.
func (o *GatewayProgramFundingSourceRequest) GetTimeoutMillis() int64 {
	if o == nil || IsNil(o.TimeoutMillis) {
		var ret int64
		return ret
	}
	return *o.TimeoutMillis
}

// GetTimeoutMillisOk returns a tuple with the TimeoutMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetTimeoutMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeoutMillis) {
		return nil, false
	}
	return o.TimeoutMillis, true
}

// HasTimeoutMillis returns a boolean if a field has been set.
func (o *GatewayProgramFundingSourceRequest) HasTimeoutMillis() bool {
	if o != nil && !IsNil(o.TimeoutMillis) {
		return true
	}

	return false
}

// SetTimeoutMillis gets a reference to the given int64 and assigns it to the TimeoutMillis field.
func (o *GatewayProgramFundingSourceRequest) SetTimeoutMillis(v int64) {
	o.TimeoutMillis = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayProgramFundingSourceRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayProgramFundingSourceRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayProgramFundingSourceRequest) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value
func (o *GatewayProgramFundingSourceRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GatewayProgramFundingSourceRequest) SetUrl(v string) {
	o.Url = v
}

// GetUseMtls returns the UseMtls field value if set, zero value otherwise.
func (o *GatewayProgramFundingSourceRequest) GetUseMtls() bool {
	if o == nil || IsNil(o.UseMtls) {
		var ret bool
		return ret
	}
	return *o.UseMtls
}

// GetUseMtlsOk returns a tuple with the UseMtls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceRequest) GetUseMtlsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMtls) {
		return nil, false
	}
	return o.UseMtls, true
}

// HasUseMtls returns a boolean if a field has been set.
func (o *GatewayProgramFundingSourceRequest) HasUseMtls() bool {
	if o != nil && !IsNil(o.UseMtls) {
		return true
	}

	return false
}

// SetUseMtls gets a reference to the given bool and assigns it to the UseMtls field.
func (o *GatewayProgramFundingSourceRequest) SetUseMtls(v bool) {
	o.UseMtls = &v
}

func (o GatewayProgramFundingSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayProgramFundingSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["basic_auth_password"] = o.BasicAuthPassword
	toSerialize["basic_auth_username"] = o.BasicAuthUsername
	if !IsNil(o.CustomHeader) {
		toSerialize["custom_header"] = o.CustomHeader
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.TimeoutMillis) {
		toSerialize["timeout_millis"] = o.TimeoutMillis
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.UseMtls) {
		toSerialize["use_mtls"] = o.UseMtls
	}
	return toSerialize, nil
}

func (o *GatewayProgramFundingSourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"basic_auth_password",
		"basic_auth_username",
		"name",
		"url",
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

	varGatewayProgramFundingSourceRequest := _GatewayProgramFundingSourceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGatewayProgramFundingSourceRequest)

	if err != nil {
		return err
	}

	*o = GatewayProgramFundingSourceRequest(varGatewayProgramFundingSourceRequest)

	return err
}

type NullableGatewayProgramFundingSourceRequest struct {
	value *GatewayProgramFundingSourceRequest
	isSet bool
}

func (v NullableGatewayProgramFundingSourceRequest) Get() *GatewayProgramFundingSourceRequest {
	return v.value
}

func (v *NullableGatewayProgramFundingSourceRequest) Set(val *GatewayProgramFundingSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayProgramFundingSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayProgramFundingSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayProgramFundingSourceRequest(val *GatewayProgramFundingSourceRequest) *NullableGatewayProgramFundingSourceRequest {
	return &NullableGatewayProgramFundingSourceRequest{value: val, isSet: true}
}

func (v NullableGatewayProgramFundingSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayProgramFundingSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


