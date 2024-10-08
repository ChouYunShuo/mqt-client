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

// checks if the Funding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Funding{}

// Funding Contains funding information for the transaction, including funding amount, type, and time.
type Funding struct {
	// Amount of funds loaded into the GPA.
	Amount *decimal.Decimal `json:"amount,omitempty"`
	GatewayLog *GatewayLogModel `json:"gateway_log,omitempty"`
	Source FundingSourceModel `json:"source"`
	SourceAddress *CardholderAddressResponse `json:"source_address,omitempty"`
}

type _Funding Funding

// NewFunding instantiates a new Funding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunding(source FundingSourceModel) *Funding {
	this := Funding{}
	this.Source = source
	return &this
}

// NewFundingWithDefaults instantiates a new Funding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingWithDefaults() *Funding {
	this := Funding{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Funding) GetAmount() decimal.Decimal {
	if o == nil || IsNil(o.Amount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Funding) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Funding) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given decimal.Decimal and assigns it to the Amount field.
func (o *Funding) SetAmount(v decimal.Decimal) {
	o.Amount = &v
}

// GetGatewayLog returns the GatewayLog field value if set, zero value otherwise.
func (o *Funding) GetGatewayLog() GatewayLogModel {
	if o == nil || IsNil(o.GatewayLog) {
		var ret GatewayLogModel
		return ret
	}
	return *o.GatewayLog
}

// GetGatewayLogOk returns a tuple with the GatewayLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Funding) GetGatewayLogOk() (*GatewayLogModel, bool) {
	if o == nil || IsNil(o.GatewayLog) {
		return nil, false
	}
	return o.GatewayLog, true
}

// HasGatewayLog returns a boolean if a field has been set.
func (o *Funding) HasGatewayLog() bool {
	if o != nil && !IsNil(o.GatewayLog) {
		return true
	}

	return false
}

// SetGatewayLog gets a reference to the given GatewayLogModel and assigns it to the GatewayLog field.
func (o *Funding) SetGatewayLog(v GatewayLogModel) {
	o.GatewayLog = &v
}

// GetSource returns the Source field value
func (o *Funding) GetSource() FundingSourceModel {
	if o == nil {
		var ret FundingSourceModel
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Funding) GetSourceOk() (*FundingSourceModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Funding) SetSource(v FundingSourceModel) {
	o.Source = v
}

// GetSourceAddress returns the SourceAddress field value if set, zero value otherwise.
func (o *Funding) GetSourceAddress() CardholderAddressResponse {
	if o == nil || IsNil(o.SourceAddress) {
		var ret CardholderAddressResponse
		return ret
	}
	return *o.SourceAddress
}

// GetSourceAddressOk returns a tuple with the SourceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Funding) GetSourceAddressOk() (*CardholderAddressResponse, bool) {
	if o == nil || IsNil(o.SourceAddress) {
		return nil, false
	}
	return o.SourceAddress, true
}

// HasSourceAddress returns a boolean if a field has been set.
func (o *Funding) HasSourceAddress() bool {
	if o != nil && !IsNil(o.SourceAddress) {
		return true
	}

	return false
}

// SetSourceAddress gets a reference to the given CardholderAddressResponse and assigns it to the SourceAddress field.
func (o *Funding) SetSourceAddress(v CardholderAddressResponse) {
	o.SourceAddress = &v
}

func (o Funding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Funding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.GatewayLog) {
		toSerialize["gateway_log"] = o.GatewayLog
	}
	toSerialize["source"] = o.Source
	if !IsNil(o.SourceAddress) {
		toSerialize["source_address"] = o.SourceAddress
	}
	return toSerialize, nil
}

func (o *Funding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
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

	varFunding := _Funding{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunding)

	if err != nil {
		return err
	}

	*o = Funding(varFunding)

	return err
}

type NullableFunding struct {
	value *Funding
	isSet bool
}

func (v NullableFunding) Get() *Funding {
	return v.value
}

func (v *NullableFunding) Set(val *Funding) {
	v.value = val
	v.isSet = true
}

func (v NullableFunding) IsSet() bool {
	return v.isSet
}

func (v *NullableFunding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunding(val *Funding) *NullableFunding {
	return &NullableFunding{value: val, isSet: true}
}

func (v NullableFunding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


