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
	"github.com/shopspring/decimal"
)

// checks if the DepositDepositResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositDepositResponse{}

// DepositDepositResponse Contains information about a direct deposit.
type DepositDepositResponse struct {
	// Amount being debited or credited.
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// The unique identifier of the business account holder.
	BusinessToken *string `json:"business_token,omitempty"`
	// Company-specific data provided by the direct deposit originator.
	CompanyDiscretionaryData *string `json:"company_discretionary_data,omitempty"`
	// Company-specific data provided by the direct deposit originator.
	CompanyEntryDescription *string `json:"company_entry_description,omitempty"`
	// Alphanumeric code that identifies the direct deposit originator.
	CompanyIdentification *string `json:"company_identification,omitempty"`
	// Name of the direct deposit originator.
	CompanyName *string `json:"company_name,omitempty"`
	// Date and time when the direct deposit account was created.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// The unique identifier of the direct deposit account.
	DirectDepositAccountToken *string `json:"direct_deposit_account_token,omitempty"`
	// Accounting number by which the recipient is known to the direct deposit originator.
	IndividualIdentificationNumber *string `json:"individual_identification_number,omitempty"`
	// Name of the direct deposit recipient.
	IndividualName *string `json:"individual_name,omitempty"`
	// Date and time when the direct deposit account was last modified.
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Date and time when the credit or debit is applied.
	SettlementDate *time.Time `json:"settlement_date,omitempty"`
	// Three-letter code identifying the type of entry.
	StandardEntryClassCode *string `json:"standard_entry_class_code,omitempty"`
	// Current status of the direct deposit record.
	State *string `json:"state,omitempty"`
	// Explanation for why the direct deposit record is in the current state.
	StateReason *string `json:"state_reason,omitempty"`
	// Standard code describing the reason for the current state.
	StateReasonCode *string `json:"state_reason_code,omitempty"`
	// The unique identifier of the direct deposit record.
	Token *string `json:"token,omitempty"`
	// Determines whether funds are being debited from or credited to the account.
	Type *string `json:"type,omitempty"`
	// The unique identifier of the user account holder.
	UserToken *string `json:"user_token,omitempty"`
}

// NewDepositDepositResponse instantiates a new DepositDepositResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositDepositResponse() *DepositDepositResponse {
	this := DepositDepositResponse{}
	return &this
}

// NewDepositDepositResponseWithDefaults instantiates a new DepositDepositResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositDepositResponseWithDefaults() *DepositDepositResponse {
	this := DepositDepositResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetAmount() decimal.Decimal {
	if o == nil || IsNil(o.Amount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given decimal.Decimal and assigns it to the Amount field.
func (o *DepositDepositResponse) SetAmount(v decimal.Decimal) {
	o.Amount = &v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *DepositDepositResponse) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetCompanyDiscretionaryData() string {
	if o == nil || IsNil(o.CompanyDiscretionaryData) {
		var ret string
		return ret
	}
	return *o.CompanyDiscretionaryData
}

// GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetCompanyDiscretionaryDataOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyDiscretionaryData) {
		return nil, false
	}
	return o.CompanyDiscretionaryData, true
}

// HasCompanyDiscretionaryData returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasCompanyDiscretionaryData() bool {
	if o != nil && !IsNil(o.CompanyDiscretionaryData) {
		return true
	}

	return false
}

// SetCompanyDiscretionaryData gets a reference to the given string and assigns it to the CompanyDiscretionaryData field.
func (o *DepositDepositResponse) SetCompanyDiscretionaryData(v string) {
	o.CompanyDiscretionaryData = &v
}

// GetCompanyEntryDescription returns the CompanyEntryDescription field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetCompanyEntryDescription() string {
	if o == nil || IsNil(o.CompanyEntryDescription) {
		var ret string
		return ret
	}
	return *o.CompanyEntryDescription
}

// GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetCompanyEntryDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyEntryDescription) {
		return nil, false
	}
	return o.CompanyEntryDescription, true
}

// HasCompanyEntryDescription returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasCompanyEntryDescription() bool {
	if o != nil && !IsNil(o.CompanyEntryDescription) {
		return true
	}

	return false
}

// SetCompanyEntryDescription gets a reference to the given string and assigns it to the CompanyEntryDescription field.
func (o *DepositDepositResponse) SetCompanyEntryDescription(v string) {
	o.CompanyEntryDescription = &v
}

// GetCompanyIdentification returns the CompanyIdentification field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetCompanyIdentification() string {
	if o == nil || IsNil(o.CompanyIdentification) {
		var ret string
		return ret
	}
	return *o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyIdentification) {
		return nil, false
	}
	return o.CompanyIdentification, true
}

// HasCompanyIdentification returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasCompanyIdentification() bool {
	if o != nil && !IsNil(o.CompanyIdentification) {
		return true
	}

	return false
}

// SetCompanyIdentification gets a reference to the given string and assigns it to the CompanyIdentification field.
func (o *DepositDepositResponse) SetCompanyIdentification(v string) {
	o.CompanyIdentification = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *DepositDepositResponse) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *DepositDepositResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDirectDepositAccountToken returns the DirectDepositAccountToken field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetDirectDepositAccountToken() string {
	if o == nil || IsNil(o.DirectDepositAccountToken) {
		var ret string
		return ret
	}
	return *o.DirectDepositAccountToken
}

// GetDirectDepositAccountTokenOk returns a tuple with the DirectDepositAccountToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetDirectDepositAccountTokenOk() (*string, bool) {
	if o == nil || IsNil(o.DirectDepositAccountToken) {
		return nil, false
	}
	return o.DirectDepositAccountToken, true
}

// HasDirectDepositAccountToken returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasDirectDepositAccountToken() bool {
	if o != nil && !IsNil(o.DirectDepositAccountToken) {
		return true
	}

	return false
}

// SetDirectDepositAccountToken gets a reference to the given string and assigns it to the DirectDepositAccountToken field.
func (o *DepositDepositResponse) SetDirectDepositAccountToken(v string) {
	o.DirectDepositAccountToken = &v
}

// GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetIndividualIdentificationNumber() string {
	if o == nil || IsNil(o.IndividualIdentificationNumber) {
		var ret string
		return ret
	}
	return *o.IndividualIdentificationNumber
}

// GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetIndividualIdentificationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualIdentificationNumber) {
		return nil, false
	}
	return o.IndividualIdentificationNumber, true
}

// HasIndividualIdentificationNumber returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasIndividualIdentificationNumber() bool {
	if o != nil && !IsNil(o.IndividualIdentificationNumber) {
		return true
	}

	return false
}

// SetIndividualIdentificationNumber gets a reference to the given string and assigns it to the IndividualIdentificationNumber field.
func (o *DepositDepositResponse) SetIndividualIdentificationNumber(v string) {
	o.IndividualIdentificationNumber = &v
}

// GetIndividualName returns the IndividualName field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetIndividualName() string {
	if o == nil || IsNil(o.IndividualName) {
		var ret string
		return ret
	}
	return *o.IndividualName
}

// GetIndividualNameOk returns a tuple with the IndividualName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetIndividualNameOk() (*string, bool) {
	if o == nil || IsNil(o.IndividualName) {
		return nil, false
	}
	return o.IndividualName, true
}

// HasIndividualName returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasIndividualName() bool {
	if o != nil && !IsNil(o.IndividualName) {
		return true
	}

	return false
}

// SetIndividualName gets a reference to the given string and assigns it to the IndividualName field.
func (o *DepositDepositResponse) SetIndividualName(v string) {
	o.IndividualName = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *DepositDepositResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetSettlementDate returns the SettlementDate field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetSettlementDate() time.Time {
	if o == nil || IsNil(o.SettlementDate) {
		var ret time.Time
		return ret
	}
	return *o.SettlementDate
}

// GetSettlementDateOk returns a tuple with the SettlementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetSettlementDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SettlementDate) {
		return nil, false
	}
	return o.SettlementDate, true
}

// HasSettlementDate returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasSettlementDate() bool {
	if o != nil && !IsNil(o.SettlementDate) {
		return true
	}

	return false
}

// SetSettlementDate gets a reference to the given time.Time and assigns it to the SettlementDate field.
func (o *DepositDepositResponse) SetSettlementDate(v time.Time) {
	o.SettlementDate = &v
}

// GetStandardEntryClassCode returns the StandardEntryClassCode field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetStandardEntryClassCode() string {
	if o == nil || IsNil(o.StandardEntryClassCode) {
		var ret string
		return ret
	}
	return *o.StandardEntryClassCode
}

// GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetStandardEntryClassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StandardEntryClassCode) {
		return nil, false
	}
	return o.StandardEntryClassCode, true
}

// HasStandardEntryClassCode returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasStandardEntryClassCode() bool {
	if o != nil && !IsNil(o.StandardEntryClassCode) {
		return true
	}

	return false
}

// SetStandardEntryClassCode gets a reference to the given string and assigns it to the StandardEntryClassCode field.
func (o *DepositDepositResponse) SetStandardEntryClassCode(v string) {
	o.StandardEntryClassCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DepositDepositResponse) SetState(v string) {
	o.State = &v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetStateReason() string {
	if o == nil || IsNil(o.StateReason) {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetStateReasonOk() (*string, bool) {
	if o == nil || IsNil(o.StateReason) {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasStateReason() bool {
	if o != nil && !IsNil(o.StateReason) {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *DepositDepositResponse) SetStateReason(v string) {
	o.StateReason = &v
}

// GetStateReasonCode returns the StateReasonCode field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetStateReasonCode() string {
	if o == nil || IsNil(o.StateReasonCode) {
		var ret string
		return ret
	}
	return *o.StateReasonCode
}

// GetStateReasonCodeOk returns a tuple with the StateReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetStateReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StateReasonCode) {
		return nil, false
	}
	return o.StateReasonCode, true
}

// HasStateReasonCode returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasStateReasonCode() bool {
	if o != nil && !IsNil(o.StateReasonCode) {
		return true
	}

	return false
}

// SetStateReasonCode gets a reference to the given string and assigns it to the StateReasonCode field.
func (o *DepositDepositResponse) SetStateReasonCode(v string) {
	o.StateReasonCode = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DepositDepositResponse) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DepositDepositResponse) SetType(v string) {
	o.Type = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *DepositDepositResponse) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositDepositResponse) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *DepositDepositResponse) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *DepositDepositResponse) SetUserToken(v string) {
	o.UserToken = &v
}

func (o DepositDepositResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositDepositResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	if !IsNil(o.CompanyDiscretionaryData) {
		toSerialize["company_discretionary_data"] = o.CompanyDiscretionaryData
	}
	if !IsNil(o.CompanyEntryDescription) {
		toSerialize["company_entry_description"] = o.CompanyEntryDescription
	}
	if !IsNil(o.CompanyIdentification) {
		toSerialize["company_identification"] = o.CompanyIdentification
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DirectDepositAccountToken) {
		toSerialize["direct_deposit_account_token"] = o.DirectDepositAccountToken
	}
	if !IsNil(o.IndividualIdentificationNumber) {
		toSerialize["individual_identification_number"] = o.IndividualIdentificationNumber
	}
	if !IsNil(o.IndividualName) {
		toSerialize["individual_name"] = o.IndividualName
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if !IsNil(o.SettlementDate) {
		toSerialize["settlement_date"] = o.SettlementDate
	}
	if !IsNil(o.StandardEntryClassCode) {
		toSerialize["standard_entry_class_code"] = o.StandardEntryClassCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateReason) {
		toSerialize["state_reason"] = o.StateReason
	}
	if !IsNil(o.StateReasonCode) {
		toSerialize["state_reason_code"] = o.StateReasonCode
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

type NullableDepositDepositResponse struct {
	value *DepositDepositResponse
	isSet bool
}

func (v NullableDepositDepositResponse) Get() *DepositDepositResponse {
	return v.value
}

func (v *NullableDepositDepositResponse) Set(val *DepositDepositResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositDepositResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositDepositResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositDepositResponse(val *DepositDepositResponse) *NullableDepositDepositResponse {
	return &NullableDepositDepositResponse{value: val, isSet: true}
}

func (v NullableDepositDepositResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositDepositResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


