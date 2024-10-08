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

// checks if the CustomerDueDiligenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDueDiligenceRequest{}

// CustomerDueDiligenceRequest struct for CustomerDueDiligenceRequest
type CustomerDueDiligenceRequest struct {
	Answer string `json:"answer"`
	Question string `json:"question"`
	Token *string `json:"token,omitempty"`
}

type _CustomerDueDiligenceRequest CustomerDueDiligenceRequest

// NewCustomerDueDiligenceRequest instantiates a new CustomerDueDiligenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDueDiligenceRequest(answer string, question string) *CustomerDueDiligenceRequest {
	this := CustomerDueDiligenceRequest{}
	this.Answer = answer
	this.Question = question
	return &this
}

// NewCustomerDueDiligenceRequestWithDefaults instantiates a new CustomerDueDiligenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDueDiligenceRequestWithDefaults() *CustomerDueDiligenceRequest {
	this := CustomerDueDiligenceRequest{}
	return &this
}

// GetAnswer returns the Answer field value
func (o *CustomerDueDiligenceRequest) GetAnswer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value
// and a boolean to check if the value has been set.
func (o *CustomerDueDiligenceRequest) GetAnswerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Answer, true
}

// SetAnswer sets field value
func (o *CustomerDueDiligenceRequest) SetAnswer(v string) {
	o.Answer = v
}

// GetQuestion returns the Question field value
func (o *CustomerDueDiligenceRequest) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *CustomerDueDiligenceRequest) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *CustomerDueDiligenceRequest) SetQuestion(v string) {
	o.Question = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CustomerDueDiligenceRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDueDiligenceRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CustomerDueDiligenceRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CustomerDueDiligenceRequest) SetToken(v string) {
	o.Token = &v
}

func (o CustomerDueDiligenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDueDiligenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answer"] = o.Answer
	toSerialize["question"] = o.Question
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *CustomerDueDiligenceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"answer",
		"question",
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

	varCustomerDueDiligenceRequest := _CustomerDueDiligenceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerDueDiligenceRequest)

	if err != nil {
		return err
	}

	*o = CustomerDueDiligenceRequest(varCustomerDueDiligenceRequest)

	return err
}

type NullableCustomerDueDiligenceRequest struct {
	value *CustomerDueDiligenceRequest
	isSet bool
}

func (v NullableCustomerDueDiligenceRequest) Get() *CustomerDueDiligenceRequest {
	return v.value
}

func (v *NullableCustomerDueDiligenceRequest) Set(val *CustomerDueDiligenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDueDiligenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDueDiligenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDueDiligenceRequest(val *CustomerDueDiligenceRequest) *NullableCustomerDueDiligenceRequest {
	return &NullableCustomerDueDiligenceRequest{value: val, isSet: true}
}

func (v NullableCustomerDueDiligenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDueDiligenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


