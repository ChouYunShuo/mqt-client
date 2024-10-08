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

// checks if the StatementInterestCharge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatementInterestCharge{}

// StatementInterestCharge Contains information on statement interest charges.
type StatementInterestCharge struct {
	// Amount of interest calculated for the billing period.
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// Type of APR.
	AprType *string `json:"apr_type,omitempty"`
	// Annual percentage rate.
	AprValue *decimal.Decimal `json:"apr_value,omitempty"`
	// Average daily balance used to calculate interest.
	BalanceSubjectToInterestRate *decimal.Decimal `json:"balance_subject_to_interest_rate,omitempty"`
	// Type of balance.  * `PURCHASE` - The balance on purchases.
	BalanceType *string `json:"balance_type,omitempty"`
}

// NewStatementInterestCharge instantiates a new StatementInterestCharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementInterestCharge() *StatementInterestCharge {
	this := StatementInterestCharge{}
	return &this
}

// NewStatementInterestChargeWithDefaults instantiates a new StatementInterestCharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementInterestChargeWithDefaults() *StatementInterestCharge {
	this := StatementInterestCharge{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *StatementInterestCharge) GetAmount() decimal.Decimal {
	if o == nil || IsNil(o.Amount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementInterestCharge) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *StatementInterestCharge) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given decimal.Decimal and assigns it to the Amount field.
func (o *StatementInterestCharge) SetAmount(v decimal.Decimal) {
	o.Amount = &v
}

// GetAprType returns the AprType field value if set, zero value otherwise.
func (o *StatementInterestCharge) GetAprType() string {
	if o == nil || IsNil(o.AprType) {
		var ret string
		return ret
	}
	return *o.AprType
}

// GetAprTypeOk returns a tuple with the AprType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementInterestCharge) GetAprTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AprType) {
		return nil, false
	}
	return o.AprType, true
}

// HasAprType returns a boolean if a field has been set.
func (o *StatementInterestCharge) HasAprType() bool {
	if o != nil && !IsNil(o.AprType) {
		return true
	}

	return false
}

// SetAprType gets a reference to the given string and assigns it to the AprType field.
func (o *StatementInterestCharge) SetAprType(v string) {
	o.AprType = &v
}

// GetAprValue returns the AprValue field value if set, zero value otherwise.
func (o *StatementInterestCharge) GetAprValue() decimal.Decimal {
	if o == nil || IsNil(o.AprValue) {
		var ret decimal.Decimal
		return ret
	}
	return *o.AprValue
}

// GetAprValueOk returns a tuple with the AprValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementInterestCharge) GetAprValueOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.AprValue) {
		return nil, false
	}
	return o.AprValue, true
}

// HasAprValue returns a boolean if a field has been set.
func (o *StatementInterestCharge) HasAprValue() bool {
	if o != nil && !IsNil(o.AprValue) {
		return true
	}

	return false
}

// SetAprValue gets a reference to the given decimal.Decimal and assigns it to the AprValue field.
func (o *StatementInterestCharge) SetAprValue(v decimal.Decimal) {
	o.AprValue = &v
}

// GetBalanceSubjectToInterestRate returns the BalanceSubjectToInterestRate field value if set, zero value otherwise.
func (o *StatementInterestCharge) GetBalanceSubjectToInterestRate() decimal.Decimal {
	if o == nil || IsNil(o.BalanceSubjectToInterestRate) {
		var ret decimal.Decimal
		return ret
	}
	return *o.BalanceSubjectToInterestRate
}

// GetBalanceSubjectToInterestRateOk returns a tuple with the BalanceSubjectToInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementInterestCharge) GetBalanceSubjectToInterestRateOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.BalanceSubjectToInterestRate) {
		return nil, false
	}
	return o.BalanceSubjectToInterestRate, true
}

// HasBalanceSubjectToInterestRate returns a boolean if a field has been set.
func (o *StatementInterestCharge) HasBalanceSubjectToInterestRate() bool {
	if o != nil && !IsNil(o.BalanceSubjectToInterestRate) {
		return true
	}

	return false
}

// SetBalanceSubjectToInterestRate gets a reference to the given decimal.Decimal and assigns it to the BalanceSubjectToInterestRate field.
func (o *StatementInterestCharge) SetBalanceSubjectToInterestRate(v decimal.Decimal) {
	o.BalanceSubjectToInterestRate = &v
}

// GetBalanceType returns the BalanceType field value if set, zero value otherwise.
func (o *StatementInterestCharge) GetBalanceType() string {
	if o == nil || IsNil(o.BalanceType) {
		var ret string
		return ret
	}
	return *o.BalanceType
}

// GetBalanceTypeOk returns a tuple with the BalanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementInterestCharge) GetBalanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BalanceType) {
		return nil, false
	}
	return o.BalanceType, true
}

// HasBalanceType returns a boolean if a field has been set.
func (o *StatementInterestCharge) HasBalanceType() bool {
	if o != nil && !IsNil(o.BalanceType) {
		return true
	}

	return false
}

// SetBalanceType gets a reference to the given string and assigns it to the BalanceType field.
func (o *StatementInterestCharge) SetBalanceType(v string) {
	o.BalanceType = &v
}

func (o StatementInterestCharge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatementInterestCharge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AprType) {
		toSerialize["apr_type"] = o.AprType
	}
	if !IsNil(o.AprValue) {
		toSerialize["apr_value"] = o.AprValue
	}
	if !IsNil(o.BalanceSubjectToInterestRate) {
		toSerialize["balance_subject_to_interest_rate"] = o.BalanceSubjectToInterestRate
	}
	if !IsNil(o.BalanceType) {
		toSerialize["balance_type"] = o.BalanceType
	}
	return toSerialize, nil
}

type NullableStatementInterestCharge struct {
	value *StatementInterestCharge
	isSet bool
}

func (v NullableStatementInterestCharge) Get() *StatementInterestCharge {
	return v.value
}

func (v *NullableStatementInterestCharge) Set(val *StatementInterestCharge) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementInterestCharge) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementInterestCharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementInterestCharge(val *StatementInterestCharge) *NullableStatementInterestCharge {
	return &NullableStatementInterestCharge{value: val, isSet: true}
}

func (v NullableStatementInterestCharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementInterestCharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


