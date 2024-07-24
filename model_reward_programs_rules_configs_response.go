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
	"time"
	"bytes"
	"fmt"
	"github.com/shopspring/decimal"
)

// checks if the RewardProgramsRulesConfigsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewardProgramsRulesConfigsResponse{}

// RewardProgramsRulesConfigsResponse struct for RewardProgramsRulesConfigsResponse
type RewardProgramsRulesConfigsResponse struct {
	AccrualType AccrualType `json:"accrual_type"`
	// Date and time when the reward rules config was created on the Marqeta platform, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Minimum amount that the balance for a billing cycle can be to apply the specified reward percentage. For example, if the `greater_than` value is `500`, the account holder earns _x_% of the account balance if they spend over $500 during a billing cycle.
	GreaterThan *decimal.Decimal `json:"greater_than,omitempty"`
	// A value of `true` indicates that the reward rules config is active.
	IsActive bool `json:"is_active"`
	// Maximum amount that the balance for a billing cycle can be to apply the specified reward percentage. For example, if the `less_than` value is `1500`, the account holder earns _x_% of the account balance if they spend under $1500 during a billing cycle.
	LessThan *decimal.Decimal `json:"less_than,omitempty"`
	// Merchant category code (MCC) of the related journal entry.
	Mcc *string `json:"mcc,omitempty"`
	// The reward percentage applied when the balance for a billing cycle is within the range specified in the `less_than` and `greater_than` fields. For example, if the `percentage` is `1`, the account holder earns 1% of the account balance if they spend between the `less_than` and `greater_than` amounts during a billing cycle.
	Percentage decimal.Decimal `json:"percentage"`
	// Unique identifier of the reward program on which the rules config is applied.
	RewardProgramToken string `json:"reward_program_token"`
	// Unique identifier of the reward rules config.
	Token string `json:"token"`
	// Date and time when the reward rules config was last updated on the Marqeta platform, in UTC.
	UpdatedTime time.Time `json:"updated_time"`
}

type _RewardProgramsRulesConfigsResponse RewardProgramsRulesConfigsResponse

// NewRewardProgramsRulesConfigsResponse instantiates a new RewardProgramsRulesConfigsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardProgramsRulesConfigsResponse(accrualType AccrualType, createdTime time.Time, isActive bool, percentage decimal.Decimal, rewardProgramToken string, token string, updatedTime time.Time) *RewardProgramsRulesConfigsResponse {
	this := RewardProgramsRulesConfigsResponse{}
	this.AccrualType = accrualType
	this.CreatedTime = createdTime
	this.IsActive = isActive
	this.Percentage = percentage
	this.RewardProgramToken = rewardProgramToken
	this.Token = token
	this.UpdatedTime = updatedTime
	return &this
}

// NewRewardProgramsRulesConfigsResponseWithDefaults instantiates a new RewardProgramsRulesConfigsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardProgramsRulesConfigsResponseWithDefaults() *RewardProgramsRulesConfigsResponse {
	this := RewardProgramsRulesConfigsResponse{}
	return &this
}

// GetAccrualType returns the AccrualType field value
func (o *RewardProgramsRulesConfigsResponse) GetAccrualType() AccrualType {
	if o == nil {
		var ret AccrualType
		return ret
	}

	return o.AccrualType
}

// GetAccrualTypeOk returns a tuple with the AccrualType field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetAccrualTypeOk() (*AccrualType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccrualType, true
}

// SetAccrualType sets field value
func (o *RewardProgramsRulesConfigsResponse) SetAccrualType(v AccrualType) {
	o.AccrualType = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *RewardProgramsRulesConfigsResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *RewardProgramsRulesConfigsResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetGreaterThan returns the GreaterThan field value if set, zero value otherwise.
func (o *RewardProgramsRulesConfigsResponse) GetGreaterThan() decimal.Decimal {
	if o == nil || IsNil(o.GreaterThan) {
		var ret decimal.Decimal
		return ret
	}
	return *o.GreaterThan
}

// GetGreaterThanOk returns a tuple with the GreaterThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetGreaterThanOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.GreaterThan) {
		return nil, false
	}
	return o.GreaterThan, true
}

// HasGreaterThan returns a boolean if a field has been set.
func (o *RewardProgramsRulesConfigsResponse) HasGreaterThan() bool {
	if o != nil && !IsNil(o.GreaterThan) {
		return true
	}

	return false
}

// SetGreaterThan gets a reference to the given decimal.Decimal and assigns it to the GreaterThan field.
func (o *RewardProgramsRulesConfigsResponse) SetGreaterThan(v decimal.Decimal) {
	o.GreaterThan = &v
}

// GetIsActive returns the IsActive field value
func (o *RewardProgramsRulesConfigsResponse) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *RewardProgramsRulesConfigsResponse) SetIsActive(v bool) {
	o.IsActive = v
}

// GetLessThan returns the LessThan field value if set, zero value otherwise.
func (o *RewardProgramsRulesConfigsResponse) GetLessThan() decimal.Decimal {
	if o == nil || IsNil(o.LessThan) {
		var ret decimal.Decimal
		return ret
	}
	return *o.LessThan
}

// GetLessThanOk returns a tuple with the LessThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetLessThanOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.LessThan) {
		return nil, false
	}
	return o.LessThan, true
}

// HasLessThan returns a boolean if a field has been set.
func (o *RewardProgramsRulesConfigsResponse) HasLessThan() bool {
	if o != nil && !IsNil(o.LessThan) {
		return true
	}

	return false
}

// SetLessThan gets a reference to the given decimal.Decimal and assigns it to the LessThan field.
func (o *RewardProgramsRulesConfigsResponse) SetLessThan(v decimal.Decimal) {
	o.LessThan = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *RewardProgramsRulesConfigsResponse) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *RewardProgramsRulesConfigsResponse) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *RewardProgramsRulesConfigsResponse) SetMcc(v string) {
	o.Mcc = &v
}

// GetPercentage returns the Percentage field value
func (o *RewardProgramsRulesConfigsResponse) GetPercentage() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetPercentageOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *RewardProgramsRulesConfigsResponse) SetPercentage(v decimal.Decimal) {
	o.Percentage = v
}

// GetRewardProgramToken returns the RewardProgramToken field value
func (o *RewardProgramsRulesConfigsResponse) GetRewardProgramToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardProgramToken
}

// GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetRewardProgramTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardProgramToken, true
}

// SetRewardProgramToken sets field value
func (o *RewardProgramsRulesConfigsResponse) SetRewardProgramToken(v string) {
	o.RewardProgramToken = v
}

// GetToken returns the Token field value
func (o *RewardProgramsRulesConfigsResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RewardProgramsRulesConfigsResponse) SetToken(v string) {
	o.Token = v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *RewardProgramsRulesConfigsResponse) GetUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsRulesConfigsResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *RewardProgramsRulesConfigsResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = v
}

func (o RewardProgramsRulesConfigsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardProgramsRulesConfigsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accrual_type"] = o.AccrualType
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.GreaterThan) {
		toSerialize["greater_than"] = o.GreaterThan
	}
	toSerialize["is_active"] = o.IsActive
	if !IsNil(o.LessThan) {
		toSerialize["less_than"] = o.LessThan
	}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	toSerialize["percentage"] = o.Percentage
	toSerialize["reward_program_token"] = o.RewardProgramToken
	toSerialize["token"] = o.Token
	toSerialize["updated_time"] = o.UpdatedTime
	return toSerialize, nil
}

func (o *RewardProgramsRulesConfigsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accrual_type",
		"created_time",
		"is_active",
		"percentage",
		"reward_program_token",
		"token",
		"updated_time",
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

	varRewardProgramsRulesConfigsResponse := _RewardProgramsRulesConfigsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRewardProgramsRulesConfigsResponse)

	if err != nil {
		return err
	}

	*o = RewardProgramsRulesConfigsResponse(varRewardProgramsRulesConfigsResponse)

	return err
}

type NullableRewardProgramsRulesConfigsResponse struct {
	value *RewardProgramsRulesConfigsResponse
	isSet bool
}

func (v NullableRewardProgramsRulesConfigsResponse) Get() *RewardProgramsRulesConfigsResponse {
	return v.value
}

func (v *NullableRewardProgramsRulesConfigsResponse) Set(val *RewardProgramsRulesConfigsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardProgramsRulesConfigsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardProgramsRulesConfigsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardProgramsRulesConfigsResponse(val *RewardProgramsRulesConfigsResponse) *NullableRewardProgramsRulesConfigsResponse {
	return &NullableRewardProgramsRulesConfigsResponse{value: val, isSet: true}
}

func (v NullableRewardProgramsRulesConfigsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardProgramsRulesConfigsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


