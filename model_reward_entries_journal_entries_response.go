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
	"github.com/shopspring/decimal"
)

// checks if the RewardEntriesJournalEntriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewardEntriesJournalEntriesResponse{}

// RewardEntriesJournalEntriesResponse struct for RewardEntriesJournalEntriesResponse
type RewardEntriesJournalEntriesResponse struct {
	// Unique identifier of the related journal entry to which the reward rule was applied to trigger the reward entry.
	RelatedJournalEntryToken *string `json:"related_journal_entry_token,omitempty"`
	// The transaction amount to which the reward rule was applied. Used to determine the value of the reward entry.
	TransactionAmount *decimal.Decimal `json:"transaction_amount,omitempty"`
	// Value of the reward entry.
	Value *decimal.Decimal `json:"value,omitempty"`
}

// NewRewardEntriesJournalEntriesResponse instantiates a new RewardEntriesJournalEntriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardEntriesJournalEntriesResponse() *RewardEntriesJournalEntriesResponse {
	this := RewardEntriesJournalEntriesResponse{}
	return &this
}

// NewRewardEntriesJournalEntriesResponseWithDefaults instantiates a new RewardEntriesJournalEntriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardEntriesJournalEntriesResponseWithDefaults() *RewardEntriesJournalEntriesResponse {
	this := RewardEntriesJournalEntriesResponse{}
	return &this
}

// GetRelatedJournalEntryToken returns the RelatedJournalEntryToken field value if set, zero value otherwise.
func (o *RewardEntriesJournalEntriesResponse) GetRelatedJournalEntryToken() string {
	if o == nil || IsNil(o.RelatedJournalEntryToken) {
		var ret string
		return ret
	}
	return *o.RelatedJournalEntryToken
}

// GetRelatedJournalEntryTokenOk returns a tuple with the RelatedJournalEntryToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardEntriesJournalEntriesResponse) GetRelatedJournalEntryTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedJournalEntryToken) {
		return nil, false
	}
	return o.RelatedJournalEntryToken, true
}

// HasRelatedJournalEntryToken returns a boolean if a field has been set.
func (o *RewardEntriesJournalEntriesResponse) HasRelatedJournalEntryToken() bool {
	if o != nil && !IsNil(o.RelatedJournalEntryToken) {
		return true
	}

	return false
}

// SetRelatedJournalEntryToken gets a reference to the given string and assigns it to the RelatedJournalEntryToken field.
func (o *RewardEntriesJournalEntriesResponse) SetRelatedJournalEntryToken(v string) {
	o.RelatedJournalEntryToken = &v
}

// GetTransactionAmount returns the TransactionAmount field value if set, zero value otherwise.
func (o *RewardEntriesJournalEntriesResponse) GetTransactionAmount() decimal.Decimal {
	if o == nil || IsNil(o.TransactionAmount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.TransactionAmount
}

// GetTransactionAmountOk returns a tuple with the TransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardEntriesJournalEntriesResponse) GetTransactionAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.TransactionAmount) {
		return nil, false
	}
	return o.TransactionAmount, true
}

// HasTransactionAmount returns a boolean if a field has been set.
func (o *RewardEntriesJournalEntriesResponse) HasTransactionAmount() bool {
	if o != nil && !IsNil(o.TransactionAmount) {
		return true
	}

	return false
}

// SetTransactionAmount gets a reference to the given decimal.Decimal and assigns it to the TransactionAmount field.
func (o *RewardEntriesJournalEntriesResponse) SetTransactionAmount(v decimal.Decimal) {
	o.TransactionAmount = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RewardEntriesJournalEntriesResponse) GetValue() decimal.Decimal {
	if o == nil || IsNil(o.Value) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardEntriesJournalEntriesResponse) GetValueOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RewardEntriesJournalEntriesResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given decimal.Decimal and assigns it to the Value field.
func (o *RewardEntriesJournalEntriesResponse) SetValue(v decimal.Decimal) {
	o.Value = &v
}

func (o RewardEntriesJournalEntriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardEntriesJournalEntriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelatedJournalEntryToken) {
		toSerialize["related_journal_entry_token"] = o.RelatedJournalEntryToken
	}
	if !IsNil(o.TransactionAmount) {
		toSerialize["transaction_amount"] = o.TransactionAmount
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRewardEntriesJournalEntriesResponse struct {
	value *RewardEntriesJournalEntriesResponse
	isSet bool
}

func (v NullableRewardEntriesJournalEntriesResponse) Get() *RewardEntriesJournalEntriesResponse {
	return v.value
}

func (v *NullableRewardEntriesJournalEntriesResponse) Set(val *RewardEntriesJournalEntriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardEntriesJournalEntriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardEntriesJournalEntriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardEntriesJournalEntriesResponse(val *RewardEntriesJournalEntriesResponse) *NullableRewardEntriesJournalEntriesResponse {
	return &NullableRewardEntriesJournalEntriesResponse{value: val, isSet: true}
}

func (v NullableRewardEntriesJournalEntriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardEntriesJournalEntriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


