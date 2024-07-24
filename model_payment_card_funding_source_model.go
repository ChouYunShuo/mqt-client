/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the PaymentCardFundingSourceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentCardFundingSourceModel{}

// PaymentCardFundingSourceModel struct for PaymentCardFundingSourceModel
type PaymentCardFundingSourceModel struct {
	FundingSourceModel
	AccountSuffix string `json:"account_suffix"`
	AccountType   string `json:"account_type"`
	// Required if 'user_token' is null
	BusinessToken *string `json:"business_token,omitempty"`
	ExpDate       string  `json:"exp_date"`
	// Required if 'business_token' is null
	UserToken *string `json:"user_token,omitempty"`
}

type _PaymentCardFundingSourceModel PaymentCardFundingSourceModel

// NewPaymentCardFundingSourceModel instantiates a new PaymentCardFundingSourceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCardFundingSourceModel(accountSuffix string, accountType string, expDate string, active bool, createdTime time.Time, isDefaultAccount bool, lastModifiedTime time.Time, token string, type_ string) *PaymentCardFundingSourceModel {
	this := PaymentCardFundingSourceModel{}
	this.Active = active
	this.CreatedTime = createdTime
	this.IsDefaultAccount = isDefaultAccount
	this.LastModifiedTime = lastModifiedTime
	this.Token = token
	this.Type = type_
	this.AccountSuffix = accountSuffix
	this.AccountType = accountType
	this.ExpDate = expDate
	return &this
}

// NewPaymentCardFundingSourceModelWithDefaults instantiates a new PaymentCardFundingSourceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCardFundingSourceModelWithDefaults() *PaymentCardFundingSourceModel {
	this := PaymentCardFundingSourceModel{}
	return &this
}

// GetAccountSuffix returns the AccountSuffix field value
func (o *PaymentCardFundingSourceModel) GetAccountSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSuffix
}

// GetAccountSuffixOk returns a tuple with the AccountSuffix field value
// and a boolean to check if the value has been set.
func (o *PaymentCardFundingSourceModel) GetAccountSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSuffix, true
}

// SetAccountSuffix sets field value
func (o *PaymentCardFundingSourceModel) SetAccountSuffix(v string) {
	o.AccountSuffix = v
}

// GetAccountType returns the AccountType field value
func (o *PaymentCardFundingSourceModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *PaymentCardFundingSourceModel) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *PaymentCardFundingSourceModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *PaymentCardFundingSourceModel) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCardFundingSourceModel) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *PaymentCardFundingSourceModel) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *PaymentCardFundingSourceModel) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetExpDate returns the ExpDate field value
func (o *PaymentCardFundingSourceModel) GetExpDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpDate
}

// GetExpDateOk returns a tuple with the ExpDate field value
// and a boolean to check if the value has been set.
func (o *PaymentCardFundingSourceModel) GetExpDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpDate, true
}

// SetExpDate sets field value
func (o *PaymentCardFundingSourceModel) SetExpDate(v string) {
	o.ExpDate = v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *PaymentCardFundingSourceModel) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCardFundingSourceModel) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *PaymentCardFundingSourceModel) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *PaymentCardFundingSourceModel) SetUserToken(v string) {
	o.UserToken = &v
}

func (o PaymentCardFundingSourceModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCardFundingSourceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFundingSourceModel, errFundingSourceModel := json.Marshal(o.FundingSourceModel)
	if errFundingSourceModel != nil {
		return map[string]interface{}{}, errFundingSourceModel
	}
	errFundingSourceModel = json.Unmarshal([]byte(serializedFundingSourceModel), &toSerialize)
	if errFundingSourceModel != nil {
		return map[string]interface{}{}, errFundingSourceModel
	}
	toSerialize["account_suffix"] = o.AccountSuffix
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["exp_date"] = o.ExpDate
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

func (o *PaymentCardFundingSourceModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_suffix",
		"account_type",
		"exp_date",
		"active",
		"created_time",
		"is_default_account",
		"last_modified_time",
		"token",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaymentCardFundingSourceModel := _PaymentCardFundingSourceModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentCardFundingSourceModel)

	if err != nil {
		return err
	}

	*o = PaymentCardFundingSourceModel(varPaymentCardFundingSourceModel)

	return err
}

type NullablePaymentCardFundingSourceModel struct {
	value *PaymentCardFundingSourceModel
	isSet bool
}

func (v NullablePaymentCardFundingSourceModel) Get() *PaymentCardFundingSourceModel {
	return v.value
}

func (v *NullablePaymentCardFundingSourceModel) Set(val *PaymentCardFundingSourceModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCardFundingSourceModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCardFundingSourceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCardFundingSourceModel(val *PaymentCardFundingSourceModel) *NullablePaymentCardFundingSourceModel {
	return &NullablePaymentCardFundingSourceModel{value: val, isSet: true}
}

func (v NullablePaymentCardFundingSourceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCardFundingSourceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
