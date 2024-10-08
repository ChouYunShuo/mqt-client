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

// checks if the InterestCalculation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterestCalculation{}

// InterestCalculation Contains the configurations for interest calculation.
type InterestCalculation struct {
	ApplicationOfCredits ApplicationOfCredits `json:"application_of_credits"`
	// Day-count convention.
	DayCount string `json:"day_count"`
	// One or more transactions that are excluded from current billing period's interest charge, but included in next.
	ExcludeTranTypes []string `json:"exclude_tran_types,omitempty"`
	// Determines the last day of grace period based on which interest charges are calculated.
	GraceDaysApplication string `json:"grace_days_application"`
	// One or more balance types on which interest is applied.
	InterestApplication []string `json:"interest_application"`
	InterestOnGraceReactivation InterestOnGraceReactivationEnum `json:"interest_on_grace_reactivation"`
	// Method of interest calculation.
	Method string `json:"method"`
	// When interest is applied, this value determines the minimum amount of interest that can be charged.
	MinimumInterest decimal.Decimal `json:"minimum_interest"`
}

type _InterestCalculation InterestCalculation

// NewInterestCalculation instantiates a new InterestCalculation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterestCalculation(applicationOfCredits ApplicationOfCredits, dayCount string, graceDaysApplication string, interestApplication []string, interestOnGraceReactivation InterestOnGraceReactivationEnum, method string, minimumInterest decimal.Decimal) *InterestCalculation {
	this := InterestCalculation{}
	this.ApplicationOfCredits = applicationOfCredits
	this.DayCount = dayCount
	this.GraceDaysApplication = graceDaysApplication
	this.InterestApplication = interestApplication
	this.InterestOnGraceReactivation = interestOnGraceReactivation
	this.Method = method
	this.MinimumInterest = minimumInterest
	return &this
}

// NewInterestCalculationWithDefaults instantiates a new InterestCalculation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterestCalculationWithDefaults() *InterestCalculation {
	this := InterestCalculation{}
	return &this
}

// GetApplicationOfCredits returns the ApplicationOfCredits field value
func (o *InterestCalculation) GetApplicationOfCredits() ApplicationOfCredits {
	if o == nil {
		var ret ApplicationOfCredits
		return ret
	}

	return o.ApplicationOfCredits
}

// GetApplicationOfCreditsOk returns a tuple with the ApplicationOfCredits field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetApplicationOfCreditsOk() (*ApplicationOfCredits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationOfCredits, true
}

// SetApplicationOfCredits sets field value
func (o *InterestCalculation) SetApplicationOfCredits(v ApplicationOfCredits) {
	o.ApplicationOfCredits = v
}

// GetDayCount returns the DayCount field value
func (o *InterestCalculation) GetDayCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayCount
}

// GetDayCountOk returns a tuple with the DayCount field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetDayCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayCount, true
}

// SetDayCount sets field value
func (o *InterestCalculation) SetDayCount(v string) {
	o.DayCount = v
}

// GetExcludeTranTypes returns the ExcludeTranTypes field value if set, zero value otherwise.
func (o *InterestCalculation) GetExcludeTranTypes() []string {
	if o == nil || IsNil(o.ExcludeTranTypes) {
		var ret []string
		return ret
	}
	return o.ExcludeTranTypes
}

// GetExcludeTranTypesOk returns a tuple with the ExcludeTranTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetExcludeTranTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeTranTypes) {
		return nil, false
	}
	return o.ExcludeTranTypes, true
}

// HasExcludeTranTypes returns a boolean if a field has been set.
func (o *InterestCalculation) HasExcludeTranTypes() bool {
	if o != nil && !IsNil(o.ExcludeTranTypes) {
		return true
	}

	return false
}

// SetExcludeTranTypes gets a reference to the given []string and assigns it to the ExcludeTranTypes field.
func (o *InterestCalculation) SetExcludeTranTypes(v []string) {
	o.ExcludeTranTypes = v
}

// GetGraceDaysApplication returns the GraceDaysApplication field value
func (o *InterestCalculation) GetGraceDaysApplication() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GraceDaysApplication
}

// GetGraceDaysApplicationOk returns a tuple with the GraceDaysApplication field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetGraceDaysApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GraceDaysApplication, true
}

// SetGraceDaysApplication sets field value
func (o *InterestCalculation) SetGraceDaysApplication(v string) {
	o.GraceDaysApplication = v
}

// GetInterestApplication returns the InterestApplication field value
func (o *InterestCalculation) GetInterestApplication() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InterestApplication
}

// GetInterestApplicationOk returns a tuple with the InterestApplication field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetInterestApplicationOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterestApplication, true
}

// SetInterestApplication sets field value
func (o *InterestCalculation) SetInterestApplication(v []string) {
	o.InterestApplication = v
}

// GetInterestOnGraceReactivation returns the InterestOnGraceReactivation field value
func (o *InterestCalculation) GetInterestOnGraceReactivation() InterestOnGraceReactivationEnum {
	if o == nil {
		var ret InterestOnGraceReactivationEnum
		return ret
	}

	return o.InterestOnGraceReactivation
}

// GetInterestOnGraceReactivationOk returns a tuple with the InterestOnGraceReactivation field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetInterestOnGraceReactivationOk() (*InterestOnGraceReactivationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterestOnGraceReactivation, true
}

// SetInterestOnGraceReactivation sets field value
func (o *InterestCalculation) SetInterestOnGraceReactivation(v InterestOnGraceReactivationEnum) {
	o.InterestOnGraceReactivation = v
}

// GetMethod returns the Method field value
func (o *InterestCalculation) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *InterestCalculation) SetMethod(v string) {
	o.Method = v
}

// GetMinimumInterest returns the MinimumInterest field value
func (o *InterestCalculation) GetMinimumInterest() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.MinimumInterest
}

// GetMinimumInterestOk returns a tuple with the MinimumInterest field value
// and a boolean to check if the value has been set.
func (o *InterestCalculation) GetMinimumInterestOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumInterest, true
}

// SetMinimumInterest sets field value
func (o *InterestCalculation) SetMinimumInterest(v decimal.Decimal) {
	o.MinimumInterest = v
}

func (o InterestCalculation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterestCalculation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application_of_credits"] = o.ApplicationOfCredits
	toSerialize["day_count"] = o.DayCount
	if !IsNil(o.ExcludeTranTypes) {
		toSerialize["exclude_tran_types"] = o.ExcludeTranTypes
	}
	toSerialize["grace_days_application"] = o.GraceDaysApplication
	toSerialize["interest_application"] = o.InterestApplication
	toSerialize["interest_on_grace_reactivation"] = o.InterestOnGraceReactivation
	toSerialize["method"] = o.Method
	toSerialize["minimum_interest"] = o.MinimumInterest
	return toSerialize, nil
}

func (o *InterestCalculation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"application_of_credits",
		"day_count",
		"grace_days_application",
		"interest_application",
		"interest_on_grace_reactivation",
		"method",
		"minimum_interest",
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

	varInterestCalculation := _InterestCalculation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInterestCalculation)

	if err != nil {
		return err
	}

	*o = InterestCalculation(varInterestCalculation)

	return err
}

type NullableInterestCalculation struct {
	value *InterestCalculation
	isSet bool
}

func (v NullableInterestCalculation) Get() *InterestCalculation {
	return v.value
}

func (v *NullableInterestCalculation) Set(val *InterestCalculation) {
	v.value = val
	v.isSet = true
}

func (v NullableInterestCalculation) IsSet() bool {
	return v.isSet
}

func (v *NullableInterestCalculation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterestCalculation(val *InterestCalculation) *NullableInterestCalculation {
	return &NullableInterestCalculation{value: val, isSet: true}
}

func (v NullableInterestCalculation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterestCalculation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


