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

// checks if the AccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountResponse{}

// AccountResponse Contains information on a credit account.
type AccountResponse struct {
	// Date and time when the credit account was activated on Marqeta's credit platform, in UTC.
	ActivationTime *time.Time `json:"activation_time,omitempty"`
	// Amount of credit available for use on the credit account.
	AvailableCredit decimal.Decimal `json:"available_credit"`
	// Unique identifier of the associated bundle product.
	BundleToken *string `json:"bundle_token,omitempty"`
	// Unique identifier of the parent business program.  Either a `user_token` or `business_token` is returned, not both.
	BusinessToken *string `json:"business_token,omitempty"`
	Config AccountConfigResponse `json:"config"`
	// Date and time when the credit account was created on Marqeta's credit platform, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Maximum balance the credit account can carry.
	CreditLimit decimal.Decimal `json:"credit_limit"`
	// Unique identifier of the associated credit product.
	CreditProductToken *string `json:"credit_product_token,omitempty"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Current purchase balance on the credit account.
	CurrentBalance decimal.Decimal `json:"current_balance"`
	// Description for the credit account.
	Description *string `json:"description,omitempty"`
	// Unique identifier you provide of the associated external credit offer.
	ExternalOfferId *string `json:"external_offer_id,omitempty"`
	LatestStatementCycleType *CycleType `json:"latest_statement_cycle_type,omitempty"`
	// Contains `max_apr_schedule` objects, which provide information about any temporary overrides of the APRs on the credit account. This could include special APR rates due to account/user sub status changes.
	MaxAprSchedules []MaxAPRSchedulesResponse `json:"max_apr_schedules,omitempty"`
	// Name of the credit account.
	Name *string `json:"name,omitempty"`
	// Amount remaining on the latest statement's minimum payment, after it's adjusted for payments, returned payments, and applicable credits that occurred after the latest statement's closing date.
	RemainingMinPaymentDue decimal.Decimal `json:"remaining_min_payment_due"`
	// Amount remaining on the latest statement's balance, after it's adjusted for payments, returned payments, and applicable credits that occurred after the latest statement's closing date.
	RemainingStatementBalance decimal.Decimal `json:"remaining_statement_balance"`
	Status AccountStatusEnum `json:"status"`
	// Substatuses of the credit account.
	Substatuses []string `json:"substatuses,omitempty"`
	// Unique identifier of the credit account.
	Token string `json:"token"`
	Type *AccountType `json:"type,omitempty"`
	// Date and time when the credit account was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime time.Time `json:"updated_time"`
	// Contains one or more `usages` objects that contain information on how a credit account is used and what types of balances are permitted on the account.
	Usages []AccountUsageResponse `json:"usages"`
	// Substatuses of the users under the credit account.
	UserSubstatuses []string `json:"user_substatuses,omitempty"`
	// Unique identifier of the primary account holder.  Either a `user_token` or `business_token` is returned, not both.
	UserToken *string `json:"user_token,omitempty"`
}

type _AccountResponse AccountResponse

// NewAccountResponse instantiates a new AccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResponse(availableCredit decimal.Decimal, config AccountConfigResponse, createdTime time.Time, creditLimit decimal.Decimal, currencyCode CurrencyCode, currentBalance decimal.Decimal, remainingMinPaymentDue decimal.Decimal, remainingStatementBalance decimal.Decimal, status AccountStatusEnum, token string, updatedTime time.Time, usages []AccountUsageResponse) *AccountResponse {
	this := AccountResponse{}
	this.AvailableCredit = availableCredit
	this.Config = config
	this.CreatedTime = createdTime
	this.CreditLimit = creditLimit
	this.CurrencyCode = currencyCode
	this.CurrentBalance = currentBalance
	this.RemainingMinPaymentDue = remainingMinPaymentDue
	this.RemainingStatementBalance = remainingStatementBalance
	this.Status = status
	this.Token = token
	this.UpdatedTime = updatedTime
	this.Usages = usages
	return &this
}

// NewAccountResponseWithDefaults instantiates a new AccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResponseWithDefaults() *AccountResponse {
	this := AccountResponse{}
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	return &this
}

// GetActivationTime returns the ActivationTime field value if set, zero value otherwise.
func (o *AccountResponse) GetActivationTime() time.Time {
	if o == nil || IsNil(o.ActivationTime) {
		var ret time.Time
		return ret
	}
	return *o.ActivationTime
}

// GetActivationTimeOk returns a tuple with the ActivationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetActivationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActivationTime) {
		return nil, false
	}
	return o.ActivationTime, true
}

// HasActivationTime returns a boolean if a field has been set.
func (o *AccountResponse) HasActivationTime() bool {
	if o != nil && !IsNil(o.ActivationTime) {
		return true
	}

	return false
}

// SetActivationTime gets a reference to the given time.Time and assigns it to the ActivationTime field.
func (o *AccountResponse) SetActivationTime(v time.Time) {
	o.ActivationTime = &v
}

// GetAvailableCredit returns the AvailableCredit field value
func (o *AccountResponse) GetAvailableCredit() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.AvailableCredit
}

// GetAvailableCreditOk returns a tuple with the AvailableCredit field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetAvailableCreditOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableCredit, true
}

// SetAvailableCredit sets field value
func (o *AccountResponse) SetAvailableCredit(v decimal.Decimal) {
	o.AvailableCredit = v
}

// GetBundleToken returns the BundleToken field value if set, zero value otherwise.
func (o *AccountResponse) GetBundleToken() string {
	if o == nil || IsNil(o.BundleToken) {
		var ret string
		return ret
	}
	return *o.BundleToken
}

// GetBundleTokenOk returns a tuple with the BundleToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetBundleTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BundleToken) {
		return nil, false
	}
	return o.BundleToken, true
}

// HasBundleToken returns a boolean if a field has been set.
func (o *AccountResponse) HasBundleToken() bool {
	if o != nil && !IsNil(o.BundleToken) {
		return true
	}

	return false
}

// SetBundleToken gets a reference to the given string and assigns it to the BundleToken field.
func (o *AccountResponse) SetBundleToken(v string) {
	o.BundleToken = &v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *AccountResponse) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *AccountResponse) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *AccountResponse) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetConfig returns the Config field value
func (o *AccountResponse) GetConfig() AccountConfigResponse {
	if o == nil {
		var ret AccountConfigResponse
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetConfigOk() (*AccountConfigResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *AccountResponse) SetConfig(v AccountConfigResponse) {
	o.Config = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *AccountResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *AccountResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCreditLimit returns the CreditLimit field value
func (o *AccountResponse) GetCreditLimit() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.CreditLimit
}

// GetCreditLimitOk returns a tuple with the CreditLimit field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCreditLimitOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditLimit, true
}

// SetCreditLimit sets field value
func (o *AccountResponse) SetCreditLimit(v decimal.Decimal) {
	o.CreditLimit = v
}

// GetCreditProductToken returns the CreditProductToken field value if set, zero value otherwise.
func (o *AccountResponse) GetCreditProductToken() string {
	if o == nil || IsNil(o.CreditProductToken) {
		var ret string
		return ret
	}
	return *o.CreditProductToken
}

// GetCreditProductTokenOk returns a tuple with the CreditProductToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCreditProductTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CreditProductToken) {
		return nil, false
	}
	return o.CreditProductToken, true
}

// HasCreditProductToken returns a boolean if a field has been set.
func (o *AccountResponse) HasCreditProductToken() bool {
	if o != nil && !IsNil(o.CreditProductToken) {
		return true
	}

	return false
}

// SetCreditProductToken gets a reference to the given string and assigns it to the CreditProductToken field.
func (o *AccountResponse) SetCreditProductToken(v string) {
	o.CreditProductToken = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *AccountResponse) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *AccountResponse) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetCurrentBalance returns the CurrentBalance field value
func (o *AccountResponse) GetCurrentBalance() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.CurrentBalance
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCurrentBalanceOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBalance, true
}

// SetCurrentBalance sets field value
func (o *AccountResponse) SetCurrentBalance(v decimal.Decimal) {
	o.CurrentBalance = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExternalOfferId returns the ExternalOfferId field value if set, zero value otherwise.
func (o *AccountResponse) GetExternalOfferId() string {
	if o == nil || IsNil(o.ExternalOfferId) {
		var ret string
		return ret
	}
	return *o.ExternalOfferId
}

// GetExternalOfferIdOk returns a tuple with the ExternalOfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetExternalOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalOfferId) {
		return nil, false
	}
	return o.ExternalOfferId, true
}

// HasExternalOfferId returns a boolean if a field has been set.
func (o *AccountResponse) HasExternalOfferId() bool {
	if o != nil && !IsNil(o.ExternalOfferId) {
		return true
	}

	return false
}

// SetExternalOfferId gets a reference to the given string and assigns it to the ExternalOfferId field.
func (o *AccountResponse) SetExternalOfferId(v string) {
	o.ExternalOfferId = &v
}

// GetLatestStatementCycleType returns the LatestStatementCycleType field value if set, zero value otherwise.
func (o *AccountResponse) GetLatestStatementCycleType() CycleType {
	if o == nil || IsNil(o.LatestStatementCycleType) {
		var ret CycleType
		return ret
	}
	return *o.LatestStatementCycleType
}

// GetLatestStatementCycleTypeOk returns a tuple with the LatestStatementCycleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetLatestStatementCycleTypeOk() (*CycleType, bool) {
	if o == nil || IsNil(o.LatestStatementCycleType) {
		return nil, false
	}
	return o.LatestStatementCycleType, true
}

// HasLatestStatementCycleType returns a boolean if a field has been set.
func (o *AccountResponse) HasLatestStatementCycleType() bool {
	if o != nil && !IsNil(o.LatestStatementCycleType) {
		return true
	}

	return false
}

// SetLatestStatementCycleType gets a reference to the given CycleType and assigns it to the LatestStatementCycleType field.
func (o *AccountResponse) SetLatestStatementCycleType(v CycleType) {
	o.LatestStatementCycleType = &v
}

// GetMaxAprSchedules returns the MaxAprSchedules field value if set, zero value otherwise.
func (o *AccountResponse) GetMaxAprSchedules() []MaxAPRSchedulesResponse {
	if o == nil || IsNil(o.MaxAprSchedules) {
		var ret []MaxAPRSchedulesResponse
		return ret
	}
	return o.MaxAprSchedules
}

// GetMaxAprSchedulesOk returns a tuple with the MaxAprSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetMaxAprSchedulesOk() ([]MaxAPRSchedulesResponse, bool) {
	if o == nil || IsNil(o.MaxAprSchedules) {
		return nil, false
	}
	return o.MaxAprSchedules, true
}

// HasMaxAprSchedules returns a boolean if a field has been set.
func (o *AccountResponse) HasMaxAprSchedules() bool {
	if o != nil && !IsNil(o.MaxAprSchedules) {
		return true
	}

	return false
}

// SetMaxAprSchedules gets a reference to the given []MaxAPRSchedulesResponse and assigns it to the MaxAprSchedules field.
func (o *AccountResponse) SetMaxAprSchedules(v []MaxAPRSchedulesResponse) {
	o.MaxAprSchedules = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountResponse) SetName(v string) {
	o.Name = &v
}

// GetRemainingMinPaymentDue returns the RemainingMinPaymentDue field value
func (o *AccountResponse) GetRemainingMinPaymentDue() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.RemainingMinPaymentDue
}

// GetRemainingMinPaymentDueOk returns a tuple with the RemainingMinPaymentDue field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetRemainingMinPaymentDueOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingMinPaymentDue, true
}

// SetRemainingMinPaymentDue sets field value
func (o *AccountResponse) SetRemainingMinPaymentDue(v decimal.Decimal) {
	o.RemainingMinPaymentDue = v
}

// GetRemainingStatementBalance returns the RemainingStatementBalance field value
func (o *AccountResponse) GetRemainingStatementBalance() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.RemainingStatementBalance
}

// GetRemainingStatementBalanceOk returns a tuple with the RemainingStatementBalance field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetRemainingStatementBalanceOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingStatementBalance, true
}

// SetRemainingStatementBalance sets field value
func (o *AccountResponse) SetRemainingStatementBalance(v decimal.Decimal) {
	o.RemainingStatementBalance = v
}

// GetStatus returns the Status field value
func (o *AccountResponse) GetStatus() AccountStatusEnum {
	if o == nil {
		var ret AccountStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetStatusOk() (*AccountStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountResponse) SetStatus(v AccountStatusEnum) {
	o.Status = v
}

// GetSubstatuses returns the Substatuses field value if set, zero value otherwise.
func (o *AccountResponse) GetSubstatuses() []string {
	if o == nil || IsNil(o.Substatuses) {
		var ret []string
		return ret
	}
	return o.Substatuses
}

// GetSubstatusesOk returns a tuple with the Substatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetSubstatusesOk() ([]string, bool) {
	if o == nil || IsNil(o.Substatuses) {
		return nil, false
	}
	return o.Substatuses, true
}

// HasSubstatuses returns a boolean if a field has been set.
func (o *AccountResponse) HasSubstatuses() bool {
	if o != nil && !IsNil(o.Substatuses) {
		return true
	}

	return false
}

// SetSubstatuses gets a reference to the given []string and assigns it to the Substatuses field.
func (o *AccountResponse) SetSubstatuses(v []string) {
	o.Substatuses = v
}

// GetToken returns the Token field value
func (o *AccountResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AccountResponse) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountResponse) GetType() AccountType {
	if o == nil || IsNil(o.Type) {
		var ret AccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AccountType and assigns it to the Type field.
func (o *AccountResponse) SetType(v AccountType) {
	o.Type = &v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *AccountResponse) GetUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *AccountResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = v
}

// GetUsages returns the Usages field value
func (o *AccountResponse) GetUsages() []AccountUsageResponse {
	if o == nil {
		var ret []AccountUsageResponse
		return ret
	}

	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetUsagesOk() ([]AccountUsageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usages, true
}

// SetUsages sets field value
func (o *AccountResponse) SetUsages(v []AccountUsageResponse) {
	o.Usages = v
}

// GetUserSubstatuses returns the UserSubstatuses field value if set, zero value otherwise.
func (o *AccountResponse) GetUserSubstatuses() []string {
	if o == nil || IsNil(o.UserSubstatuses) {
		var ret []string
		return ret
	}
	return o.UserSubstatuses
}

// GetUserSubstatusesOk returns a tuple with the UserSubstatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetUserSubstatusesOk() ([]string, bool) {
	if o == nil || IsNil(o.UserSubstatuses) {
		return nil, false
	}
	return o.UserSubstatuses, true
}

// HasUserSubstatuses returns a boolean if a field has been set.
func (o *AccountResponse) HasUserSubstatuses() bool {
	if o != nil && !IsNil(o.UserSubstatuses) {
		return true
	}

	return false
}

// SetUserSubstatuses gets a reference to the given []string and assigns it to the UserSubstatuses field.
func (o *AccountResponse) SetUserSubstatuses(v []string) {
	o.UserSubstatuses = v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *AccountResponse) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *AccountResponse) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *AccountResponse) SetUserToken(v string) {
	o.UserToken = &v
}

func (o AccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivationTime) {
		toSerialize["activation_time"] = o.ActivationTime
	}
	toSerialize["available_credit"] = o.AvailableCredit
	if !IsNil(o.BundleToken) {
		toSerialize["bundle_token"] = o.BundleToken
	}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["config"] = o.Config
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["credit_limit"] = o.CreditLimit
	if !IsNil(o.CreditProductToken) {
		toSerialize["credit_product_token"] = o.CreditProductToken
	}
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["current_balance"] = o.CurrentBalance
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalOfferId) {
		toSerialize["external_offer_id"] = o.ExternalOfferId
	}
	if !IsNil(o.LatestStatementCycleType) {
		toSerialize["latest_statement_cycle_type"] = o.LatestStatementCycleType
	}
	if !IsNil(o.MaxAprSchedules) {
		toSerialize["max_apr_schedules"] = o.MaxAprSchedules
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["remaining_min_payment_due"] = o.RemainingMinPaymentDue
	toSerialize["remaining_statement_balance"] = o.RemainingStatementBalance
	toSerialize["status"] = o.Status
	if !IsNil(o.Substatuses) {
		toSerialize["substatuses"] = o.Substatuses
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["updated_time"] = o.UpdatedTime
	toSerialize["usages"] = o.Usages
	if !IsNil(o.UserSubstatuses) {
		toSerialize["user_substatuses"] = o.UserSubstatuses
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

func (o *AccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"available_credit",
		"config",
		"created_time",
		"credit_limit",
		"currency_code",
		"current_balance",
		"remaining_min_payment_due",
		"remaining_statement_balance",
		"status",
		"token",
		"updated_time",
		"usages",
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

	varAccountResponse := _AccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountResponse)

	if err != nil {
		return err
	}

	*o = AccountResponse(varAccountResponse)

	return err
}

type NullableAccountResponse struct {
	value *AccountResponse
	isSet bool
}

func (v NullableAccountResponse) Get() *AccountResponse {
	return v.value
}

func (v *NullableAccountResponse) Set(val *AccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResponse(val *AccountResponse) *NullableAccountResponse {
	return &NullableAccountResponse{value: val, isSet: true}
}

func (v NullableAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


