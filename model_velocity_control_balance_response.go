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

// checks if the VelocityControlBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VelocityControlBalanceResponse{}

// VelocityControlBalanceResponse struct for VelocityControlBalanceResponse
type VelocityControlBalanceResponse struct {
	// Indicates whether the velocity control is active.
	Active *bool `json:"active,omitempty"`
	// Maximum monetary sum that can be cleared within the time period defined by the `velocity_window` field.
	AmountLimit decimal.Decimal `json:"amount_limit"`
	// If set to `true`, only approved transactions are subject to control.
	ApprovalsOnly *bool `json:"approvals_only,omitempty"`
	Association *SpendControlAssociation `json:"association,omitempty"`
	Available Available `json:"available"`
	// Three-character ISO 4217 currency code.
	CurrencyCode string `json:"currency_code"`
	// If set to `true`, the cashback components of point-of-sale transactions are subject to control.
	IncludeCashback *bool `json:"include_cashback,omitempty"`
	// If set to `true`, original credit transactions (OCT) are subject to control.
	IncludeCredits *bool `json:"include_credits,omitempty"`
	// If set to `true`, the following transactions are subject to control:  * *Account funding:* All account funding transactions * *Cashback:* Only the purchase component of cashback transactions * *Purchase transactions:* All authorizations, PIN debit transactions, and incremental transactions * *Quasi-cash:* All quasi-cash transactions * *Refunds:* All refund transactions (see <</developer-guides/controlling-spending#_controls_to_limit_amount_and_frequency_of_spending, Controls to limit amount and frequency of spending>> for more information) * *Reversals:* All reversal transactions
	IncludePurchases *bool `json:"include_purchases,omitempty"`
	// If set to `true`, account-to-account transfers are subject to control. Account-to-account transfers are not currently supported.
	IncludeTransfers *bool `json:"include_transfers,omitempty"`
	// If set to `true`, ATM withdrawals are subject to control.
	IncludeWithdrawals *bool `json:"include_withdrawals,omitempty"`
	MerchantScope *MerchantScope `json:"merchant_scope,omitempty"`
	MoneyInTransaction *MoneyInTransaction `json:"money_in_transaction,omitempty"`
	// Description of how the velocity control restricts spending, for example, \"Max spend of $500 per day\" or \"Max spend of $5000 per month for non-exempt employees\".
	Name *string `json:"name,omitempty"`
	// Unique identifier of the velocity control.
	Token *string `json:"token,omitempty"`
	// Maximum number of times a card can be used within the time period defined by the `velocity_window` field.
	UsageLimit *int32 `json:"usage_limit,omitempty"`
	// Defines the time period to which the `amount_limit` and `usage_limit` fields apply:  * *DAY* – one day; days begin at 00:00:00 UTC. * *WEEK* – one week; weeks begin Sundays at 00:00:00 UTC. * *MONTH* – one month; months begin on the first day of month at 00:00:00 UTC. * *LIFETIME* – forever; time period never expires. * *TRANSACTION* – a single transaction.
	VelocityWindow string `json:"velocity_window"`
}

type _VelocityControlBalanceResponse VelocityControlBalanceResponse

// NewVelocityControlBalanceResponse instantiates a new VelocityControlBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVelocityControlBalanceResponse(amountLimit decimal.Decimal, available Available, currencyCode string, velocityWindow string) *VelocityControlBalanceResponse {
	this := VelocityControlBalanceResponse{}
	this.AmountLimit = amountLimit
	this.Available = available
	this.CurrencyCode = currencyCode
	this.VelocityWindow = velocityWindow
	return &this
}

// NewVelocityControlBalanceResponseWithDefaults instantiates a new VelocityControlBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelocityControlBalanceResponseWithDefaults() *VelocityControlBalanceResponse {
	this := VelocityControlBalanceResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *VelocityControlBalanceResponse) SetActive(v bool) {
	o.Active = &v
}

// GetAmountLimit returns the AmountLimit field value
func (o *VelocityControlBalanceResponse) GetAmountLimit() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.AmountLimit
}

// GetAmountLimitOk returns a tuple with the AmountLimit field value
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetAmountLimitOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountLimit, true
}

// SetAmountLimit sets field value
func (o *VelocityControlBalanceResponse) SetAmountLimit(v decimal.Decimal) {
	o.AmountLimit = v
}

// GetApprovalsOnly returns the ApprovalsOnly field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetApprovalsOnly() bool {
	if o == nil || IsNil(o.ApprovalsOnly) {
		var ret bool
		return ret
	}
	return *o.ApprovalsOnly
}

// GetApprovalsOnlyOk returns a tuple with the ApprovalsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetApprovalsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ApprovalsOnly) {
		return nil, false
	}
	return o.ApprovalsOnly, true
}

// HasApprovalsOnly returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasApprovalsOnly() bool {
	if o != nil && !IsNil(o.ApprovalsOnly) {
		return true
	}

	return false
}

// SetApprovalsOnly gets a reference to the given bool and assigns it to the ApprovalsOnly field.
func (o *VelocityControlBalanceResponse) SetApprovalsOnly(v bool) {
	o.ApprovalsOnly = &v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetAssociation() SpendControlAssociation {
	if o == nil || IsNil(o.Association) {
		var ret SpendControlAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetAssociationOk() (*SpendControlAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given SpendControlAssociation and assigns it to the Association field.
func (o *VelocityControlBalanceResponse) SetAssociation(v SpendControlAssociation) {
	o.Association = &v
}

// GetAvailable returns the Available field value
func (o *VelocityControlBalanceResponse) GetAvailable() Available {
	if o == nil {
		var ret Available
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetAvailableOk() (*Available, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *VelocityControlBalanceResponse) SetAvailable(v Available) {
	o.Available = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *VelocityControlBalanceResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *VelocityControlBalanceResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetIncludeCashback returns the IncludeCashback field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetIncludeCashback() bool {
	if o == nil || IsNil(o.IncludeCashback) {
		var ret bool
		return ret
	}
	return *o.IncludeCashback
}

// GetIncludeCashbackOk returns a tuple with the IncludeCashback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetIncludeCashbackOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCashback) {
		return nil, false
	}
	return o.IncludeCashback, true
}

// HasIncludeCashback returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasIncludeCashback() bool {
	if o != nil && !IsNil(o.IncludeCashback) {
		return true
	}

	return false
}

// SetIncludeCashback gets a reference to the given bool and assigns it to the IncludeCashback field.
func (o *VelocityControlBalanceResponse) SetIncludeCashback(v bool) {
	o.IncludeCashback = &v
}

// GetIncludeCredits returns the IncludeCredits field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetIncludeCredits() bool {
	if o == nil || IsNil(o.IncludeCredits) {
		var ret bool
		return ret
	}
	return *o.IncludeCredits
}

// GetIncludeCreditsOk returns a tuple with the IncludeCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetIncludeCreditsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCredits) {
		return nil, false
	}
	return o.IncludeCredits, true
}

// HasIncludeCredits returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasIncludeCredits() bool {
	if o != nil && !IsNil(o.IncludeCredits) {
		return true
	}

	return false
}

// SetIncludeCredits gets a reference to the given bool and assigns it to the IncludeCredits field.
func (o *VelocityControlBalanceResponse) SetIncludeCredits(v bool) {
	o.IncludeCredits = &v
}

// GetIncludePurchases returns the IncludePurchases field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetIncludePurchases() bool {
	if o == nil || IsNil(o.IncludePurchases) {
		var ret bool
		return ret
	}
	return *o.IncludePurchases
}

// GetIncludePurchasesOk returns a tuple with the IncludePurchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetIncludePurchasesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludePurchases) {
		return nil, false
	}
	return o.IncludePurchases, true
}

// HasIncludePurchases returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasIncludePurchases() bool {
	if o != nil && !IsNil(o.IncludePurchases) {
		return true
	}

	return false
}

// SetIncludePurchases gets a reference to the given bool and assigns it to the IncludePurchases field.
func (o *VelocityControlBalanceResponse) SetIncludePurchases(v bool) {
	o.IncludePurchases = &v
}

// GetIncludeTransfers returns the IncludeTransfers field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetIncludeTransfers() bool {
	if o == nil || IsNil(o.IncludeTransfers) {
		var ret bool
		return ret
	}
	return *o.IncludeTransfers
}

// GetIncludeTransfersOk returns a tuple with the IncludeTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetIncludeTransfersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTransfers) {
		return nil, false
	}
	return o.IncludeTransfers, true
}

// HasIncludeTransfers returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasIncludeTransfers() bool {
	if o != nil && !IsNil(o.IncludeTransfers) {
		return true
	}

	return false
}

// SetIncludeTransfers gets a reference to the given bool and assigns it to the IncludeTransfers field.
func (o *VelocityControlBalanceResponse) SetIncludeTransfers(v bool) {
	o.IncludeTransfers = &v
}

// GetIncludeWithdrawals returns the IncludeWithdrawals field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetIncludeWithdrawals() bool {
	if o == nil || IsNil(o.IncludeWithdrawals) {
		var ret bool
		return ret
	}
	return *o.IncludeWithdrawals
}

// GetIncludeWithdrawalsOk returns a tuple with the IncludeWithdrawals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetIncludeWithdrawalsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeWithdrawals) {
		return nil, false
	}
	return o.IncludeWithdrawals, true
}

// HasIncludeWithdrawals returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasIncludeWithdrawals() bool {
	if o != nil && !IsNil(o.IncludeWithdrawals) {
		return true
	}

	return false
}

// SetIncludeWithdrawals gets a reference to the given bool and assigns it to the IncludeWithdrawals field.
func (o *VelocityControlBalanceResponse) SetIncludeWithdrawals(v bool) {
	o.IncludeWithdrawals = &v
}

// GetMerchantScope returns the MerchantScope field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetMerchantScope() MerchantScope {
	if o == nil || IsNil(o.MerchantScope) {
		var ret MerchantScope
		return ret
	}
	return *o.MerchantScope
}

// GetMerchantScopeOk returns a tuple with the MerchantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetMerchantScopeOk() (*MerchantScope, bool) {
	if o == nil || IsNil(o.MerchantScope) {
		return nil, false
	}
	return o.MerchantScope, true
}

// HasMerchantScope returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasMerchantScope() bool {
	if o != nil && !IsNil(o.MerchantScope) {
		return true
	}

	return false
}

// SetMerchantScope gets a reference to the given MerchantScope and assigns it to the MerchantScope field.
func (o *VelocityControlBalanceResponse) SetMerchantScope(v MerchantScope) {
	o.MerchantScope = &v
}

// GetMoneyInTransaction returns the MoneyInTransaction field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetMoneyInTransaction() MoneyInTransaction {
	if o == nil || IsNil(o.MoneyInTransaction) {
		var ret MoneyInTransaction
		return ret
	}
	return *o.MoneyInTransaction
}

// GetMoneyInTransactionOk returns a tuple with the MoneyInTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetMoneyInTransactionOk() (*MoneyInTransaction, bool) {
	if o == nil || IsNil(o.MoneyInTransaction) {
		return nil, false
	}
	return o.MoneyInTransaction, true
}

// HasMoneyInTransaction returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasMoneyInTransaction() bool {
	if o != nil && !IsNil(o.MoneyInTransaction) {
		return true
	}

	return false
}

// SetMoneyInTransaction gets a reference to the given MoneyInTransaction and assigns it to the MoneyInTransaction field.
func (o *VelocityControlBalanceResponse) SetMoneyInTransaction(v MoneyInTransaction) {
	o.MoneyInTransaction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VelocityControlBalanceResponse) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *VelocityControlBalanceResponse) SetToken(v string) {
	o.Token = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *VelocityControlBalanceResponse) GetUsageLimit() int32 {
	if o == nil || IsNil(o.UsageLimit) {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetUsageLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UsageLimit) {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *VelocityControlBalanceResponse) HasUsageLimit() bool {
	if o != nil && !IsNil(o.UsageLimit) {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *VelocityControlBalanceResponse) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

// GetVelocityWindow returns the VelocityWindow field value
func (o *VelocityControlBalanceResponse) GetVelocityWindow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VelocityWindow
}

// GetVelocityWindowOk returns a tuple with the VelocityWindow field value
// and a boolean to check if the value has been set.
func (o *VelocityControlBalanceResponse) GetVelocityWindowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VelocityWindow, true
}

// SetVelocityWindow sets field value
func (o *VelocityControlBalanceResponse) SetVelocityWindow(v string) {
	o.VelocityWindow = v
}

func (o VelocityControlBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelocityControlBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["amount_limit"] = o.AmountLimit
	if !IsNil(o.ApprovalsOnly) {
		toSerialize["approvals_only"] = o.ApprovalsOnly
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	toSerialize["available"] = o.Available
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.IncludeCashback) {
		toSerialize["include_cashback"] = o.IncludeCashback
	}
	if !IsNil(o.IncludeCredits) {
		toSerialize["include_credits"] = o.IncludeCredits
	}
	if !IsNil(o.IncludePurchases) {
		toSerialize["include_purchases"] = o.IncludePurchases
	}
	if !IsNil(o.IncludeTransfers) {
		toSerialize["include_transfers"] = o.IncludeTransfers
	}
	if !IsNil(o.IncludeWithdrawals) {
		toSerialize["include_withdrawals"] = o.IncludeWithdrawals
	}
	if !IsNil(o.MerchantScope) {
		toSerialize["merchant_scope"] = o.MerchantScope
	}
	if !IsNil(o.MoneyInTransaction) {
		toSerialize["money_in_transaction"] = o.MoneyInTransaction
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UsageLimit) {
		toSerialize["usage_limit"] = o.UsageLimit
	}
	toSerialize["velocity_window"] = o.VelocityWindow
	return toSerialize, nil
}

func (o *VelocityControlBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_limit",
		"available",
		"currency_code",
		"velocity_window",
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

	varVelocityControlBalanceResponse := _VelocityControlBalanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVelocityControlBalanceResponse)

	if err != nil {
		return err
	}

	*o = VelocityControlBalanceResponse(varVelocityControlBalanceResponse)

	return err
}

type NullableVelocityControlBalanceResponse struct {
	value *VelocityControlBalanceResponse
	isSet bool
}

func (v NullableVelocityControlBalanceResponse) Get() *VelocityControlBalanceResponse {
	return v.value
}

func (v *NullableVelocityControlBalanceResponse) Set(val *VelocityControlBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityControlBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityControlBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityControlBalanceResponse(val *VelocityControlBalanceResponse) *NullableVelocityControlBalanceResponse {
	return &NullableVelocityControlBalanceResponse{value: val, isSet: true}
}

func (v NullableVelocityControlBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityControlBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


