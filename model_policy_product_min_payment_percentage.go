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
	"github.com/shopspring/decimal"
)

// checks if the PolicyProductMinPaymentPercentage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyProductMinPaymentPercentage{}

// PolicyProductMinPaymentPercentage Contains information used to calculate the minimum payment amount when expressed as a percentage.
type PolicyProductMinPaymentPercentage struct {
	// One or more fee types to include when calculating the minimum payment.
	IncludeFeesCharged []string `json:"include_fees_charged,omitempty"`
	// Whether to include the amount of interest charged when calculating the minimum payment.
	IncludeInterestCharged bool `json:"include_interest_charged"`
	// Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day.
	PercentageOfBalance decimal.Decimal `json:"percentage_of_balance"`
}

type _PolicyProductMinPaymentPercentage PolicyProductMinPaymentPercentage

// NewPolicyProductMinPaymentPercentage instantiates a new PolicyProductMinPaymentPercentage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyProductMinPaymentPercentage(includeInterestCharged bool, percentageOfBalance decimal.Decimal) *PolicyProductMinPaymentPercentage {
	this := PolicyProductMinPaymentPercentage{}
	this.IncludeInterestCharged = includeInterestCharged
	this.PercentageOfBalance = percentageOfBalance
	return &this
}

// NewPolicyProductMinPaymentPercentageWithDefaults instantiates a new PolicyProductMinPaymentPercentage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyProductMinPaymentPercentageWithDefaults() *PolicyProductMinPaymentPercentage {
	this := PolicyProductMinPaymentPercentage{}
	return &this
}

// GetIncludeFeesCharged returns the IncludeFeesCharged field value if set, zero value otherwise.
func (o *PolicyProductMinPaymentPercentage) GetIncludeFeesCharged() []string {
	if o == nil || IsNil(o.IncludeFeesCharged) {
		var ret []string
		return ret
	}
	return o.IncludeFeesCharged
}

// GetIncludeFeesChargedOk returns a tuple with the IncludeFeesCharged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentPercentage) GetIncludeFeesChargedOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFeesCharged) {
		return nil, false
	}
	return o.IncludeFeesCharged, true
}

// HasIncludeFeesCharged returns a boolean if a field has been set.
func (o *PolicyProductMinPaymentPercentage) HasIncludeFeesCharged() bool {
	if o != nil && !IsNil(o.IncludeFeesCharged) {
		return true
	}

	return false
}

// SetIncludeFeesCharged gets a reference to the given []string and assigns it to the IncludeFeesCharged field.
func (o *PolicyProductMinPaymentPercentage) SetIncludeFeesCharged(v []string) {
	o.IncludeFeesCharged = v
}

// GetIncludeInterestCharged returns the IncludeInterestCharged field value
func (o *PolicyProductMinPaymentPercentage) GetIncludeInterestCharged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeInterestCharged
}

// GetIncludeInterestChargedOk returns a tuple with the IncludeInterestCharged field value
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentPercentage) GetIncludeInterestChargedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeInterestCharged, true
}

// SetIncludeInterestCharged sets field value
func (o *PolicyProductMinPaymentPercentage) SetIncludeInterestCharged(v bool) {
	o.IncludeInterestCharged = v
}

// GetPercentageOfBalance returns the PercentageOfBalance field value
func (o *PolicyProductMinPaymentPercentage) GetPercentageOfBalance() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.PercentageOfBalance
}

// GetPercentageOfBalanceOk returns a tuple with the PercentageOfBalance field value
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentPercentage) GetPercentageOfBalanceOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentageOfBalance, true
}

// SetPercentageOfBalance sets field value
func (o *PolicyProductMinPaymentPercentage) SetPercentageOfBalance(v decimal.Decimal) {
	o.PercentageOfBalance = v
}

func (o PolicyProductMinPaymentPercentage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyProductMinPaymentPercentage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeFeesCharged) {
		toSerialize["include_fees_charged"] = o.IncludeFeesCharged
	}
	toSerialize["include_interest_charged"] = o.IncludeInterestCharged
	toSerialize["percentage_of_balance"] = o.PercentageOfBalance
	return toSerialize, nil
}

func (o *PolicyProductMinPaymentPercentage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"include_interest_charged",
		"percentage_of_balance",
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

	varPolicyProductMinPaymentPercentage := _PolicyProductMinPaymentPercentage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyProductMinPaymentPercentage)

	if err != nil {
		return err
	}

	*o = PolicyProductMinPaymentPercentage(varPolicyProductMinPaymentPercentage)

	return err
}

type NullablePolicyProductMinPaymentPercentage struct {
	value *PolicyProductMinPaymentPercentage
	isSet bool
}

func (v NullablePolicyProductMinPaymentPercentage) Get() *PolicyProductMinPaymentPercentage {
	return v.value
}

func (v *NullablePolicyProductMinPaymentPercentage) Set(val *PolicyProductMinPaymentPercentage) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyProductMinPaymentPercentage) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyProductMinPaymentPercentage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyProductMinPaymentPercentage(val *PolicyProductMinPaymentPercentage) *NullablePolicyProductMinPaymentPercentage {
	return &NullablePolicyProductMinPaymentPercentage{value: val, isSet: true}
}

func (v NullablePolicyProductMinPaymentPercentage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyProductMinPaymentPercentage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


