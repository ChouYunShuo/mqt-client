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

// checks if the AccountAdjustmentReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAdjustmentReq{}

// AccountAdjustmentReq Contains information relevant to creating an account adjustment.
type AccountAdjustmentReq struct {
	// Amount of the adjustment.  Value must be negative if `original_ledger_entry_token` is not passed.
	Amount decimal.Decimal `json:"amount"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Description of the adjustment.
	Description string `json:"description"`
	// Unique identifier you provide of an associated external adjustment that exists outside Marqeta's credit platform.
	ExternalAdjustmentId *string `json:"external_adjustment_id,omitempty"`
	// Additional information on the adjustment.
	Note *string `json:"note,omitempty"`
	// Unique identifier of the original journal entry needing the adjustment.  Required when adjusting an existing journal entry.
	OriginalLedgerEntryToken *string `json:"original_ledger_entry_token,omitempty"`
	// Reason for the adjustment.  * `DISPUTE` - The adjustment occurred because a dispute was initiated. * `DISPUTE_RESOLUTION` - The adjustment occurred because of the result of a dispute resolution. * `RETURNED_OR_CANCELED_PAYMENT` - The adjustment occurred because a payment was returned or canceled. * `OTHER` - Any other reason the adjustment occurred. For example, a waived fee or account write-off.
	Reason *string `json:"reason,omitempty"`
	// Unique identifier of the adjustment.
	Token *string `json:"token,omitempty"`
}

type _AccountAdjustmentReq AccountAdjustmentReq

// NewAccountAdjustmentReq instantiates a new AccountAdjustmentReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAdjustmentReq(amount decimal.Decimal, currencyCode CurrencyCode, description string) *AccountAdjustmentReq {
	this := AccountAdjustmentReq{}
	this.Amount = amount
	this.CurrencyCode = currencyCode
	this.Description = description
	return &this
}

// NewAccountAdjustmentReqWithDefaults instantiates a new AccountAdjustmentReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAdjustmentReqWithDefaults() *AccountAdjustmentReq {
	this := AccountAdjustmentReq{}
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	return &this
}

// GetAmount returns the Amount field value
func (o *AccountAdjustmentReq) GetAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AccountAdjustmentReq) SetAmount(v decimal.Decimal) {
	o.Amount = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *AccountAdjustmentReq) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *AccountAdjustmentReq) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value
func (o *AccountAdjustmentReq) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AccountAdjustmentReq) SetDescription(v string) {
	o.Description = v
}

// GetExternalAdjustmentId returns the ExternalAdjustmentId field value if set, zero value otherwise.
func (o *AccountAdjustmentReq) GetExternalAdjustmentId() string {
	if o == nil || IsNil(o.ExternalAdjustmentId) {
		var ret string
		return ret
	}
	return *o.ExternalAdjustmentId
}

// GetExternalAdjustmentIdOk returns a tuple with the ExternalAdjustmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetExternalAdjustmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalAdjustmentId) {
		return nil, false
	}
	return o.ExternalAdjustmentId, true
}

// HasExternalAdjustmentId returns a boolean if a field has been set.
func (o *AccountAdjustmentReq) HasExternalAdjustmentId() bool {
	if o != nil && !IsNil(o.ExternalAdjustmentId) {
		return true
	}

	return false
}

// SetExternalAdjustmentId gets a reference to the given string and assigns it to the ExternalAdjustmentId field.
func (o *AccountAdjustmentReq) SetExternalAdjustmentId(v string) {
	o.ExternalAdjustmentId = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *AccountAdjustmentReq) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *AccountAdjustmentReq) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *AccountAdjustmentReq) SetNote(v string) {
	o.Note = &v
}

// GetOriginalLedgerEntryToken returns the OriginalLedgerEntryToken field value if set, zero value otherwise.
func (o *AccountAdjustmentReq) GetOriginalLedgerEntryToken() string {
	if o == nil || IsNil(o.OriginalLedgerEntryToken) {
		var ret string
		return ret
	}
	return *o.OriginalLedgerEntryToken
}

// GetOriginalLedgerEntryTokenOk returns a tuple with the OriginalLedgerEntryToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetOriginalLedgerEntryTokenOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalLedgerEntryToken) {
		return nil, false
	}
	return o.OriginalLedgerEntryToken, true
}

// HasOriginalLedgerEntryToken returns a boolean if a field has been set.
func (o *AccountAdjustmentReq) HasOriginalLedgerEntryToken() bool {
	if o != nil && !IsNil(o.OriginalLedgerEntryToken) {
		return true
	}

	return false
}

// SetOriginalLedgerEntryToken gets a reference to the given string and assigns it to the OriginalLedgerEntryToken field.
func (o *AccountAdjustmentReq) SetOriginalLedgerEntryToken(v string) {
	o.OriginalLedgerEntryToken = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AccountAdjustmentReq) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AccountAdjustmentReq) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AccountAdjustmentReq) SetReason(v string) {
	o.Reason = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AccountAdjustmentReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdjustmentReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AccountAdjustmentReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AccountAdjustmentReq) SetToken(v string) {
	o.Token = &v
}

func (o AccountAdjustmentReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAdjustmentReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["description"] = o.Description
	if !IsNil(o.ExternalAdjustmentId) {
		toSerialize["external_adjustment_id"] = o.ExternalAdjustmentId
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.OriginalLedgerEntryToken) {
		toSerialize["original_ledger_entry_token"] = o.OriginalLedgerEntryToken
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *AccountAdjustmentReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency_code",
		"description",
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

	varAccountAdjustmentReq := _AccountAdjustmentReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountAdjustmentReq)

	if err != nil {
		return err
	}

	*o = AccountAdjustmentReq(varAccountAdjustmentReq)

	return err
}

type NullableAccountAdjustmentReq struct {
	value *AccountAdjustmentReq
	isSet bool
}

func (v NullableAccountAdjustmentReq) Get() *AccountAdjustmentReq {
	return v.value
}

func (v *NullableAccountAdjustmentReq) Set(val *AccountAdjustmentReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAdjustmentReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAdjustmentReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAdjustmentReq(val *AccountAdjustmentReq) *NullableAccountAdjustmentReq {
	return &NullableAccountAdjustmentReq{value: val, isSet: true}
}

func (v NullableAccountAdjustmentReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAdjustmentReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


