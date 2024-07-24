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

// checks if the ProductCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductCreateReq{}

// ProductCreateReq Specifies shared details for a credit program.  Once set to `ACTIVE`, cannot be edited or deleted.
type ProductCreateReq struct {
	// One or more associated card product tokens.
	CardProductTokens []string `json:"card_product_tokens"`
	Classification ProductClassification `json:"classification"`
	Config ProductConfig `json:"config"`
	CreditLine ProductCreditLine `json:"credit_line"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Description of the credit product.
	Description *string `json:"description,omitempty"`
	InterestCalculation InterestCalculation `json:"interest_calculation"`
	MinPaymentCalculation *ProductMinPaymentCalculation `json:"min_payment_calculation,omitempty"`
	// Minimum payment, expressed as a flat amount, due on the payment due day.
	MinPaymentFlatAmount decimal.Decimal `json:"min_payment_flat_amount"`
	// Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day.
	MinPaymentPercentage decimal.Decimal `json:"min_payment_percentage"`
	// Name of the credit product.
	Name string `json:"name"`
	// Ordered list of balance types to which payments are allocated, from first to last.
	PaymentAllocationOrder []PaymentAllocationOrderEnum `json:"payment_allocation_order"`
	ProductSubType ProductSubType `json:"product_sub_type"`
	ProductType ProductType `json:"product_type"`
	Status *ResourceStatus `json:"status,omitempty"`
	// Unique identifier of the credit product.
	Token *string `json:"token,omitempty" validate:"regexp=(?!^ +$)^.+$"`
	// One or more usage types for the credit product.
	Usage []BalanceType `json:"usage"`
}

type _ProductCreateReq ProductCreateReq

// NewProductCreateReq instantiates a new ProductCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCreateReq(cardProductTokens []string, classification ProductClassification, config ProductConfig, creditLine ProductCreditLine, currencyCode CurrencyCode, interestCalculation InterestCalculation, minPaymentFlatAmount decimal.Decimal, minPaymentPercentage decimal.Decimal, name string, paymentAllocationOrder []PaymentAllocationOrderEnum, productSubType ProductSubType, productType ProductType, usage []BalanceType) *ProductCreateReq {
	this := ProductCreateReq{}
	this.CardProductTokens = cardProductTokens
	this.Classification = classification
	this.Config = config
	this.CreditLine = creditLine
	this.CurrencyCode = currencyCode
	this.InterestCalculation = interestCalculation
	this.MinPaymentFlatAmount = minPaymentFlatAmount
	this.MinPaymentPercentage = minPaymentPercentage
	this.Name = name
	this.PaymentAllocationOrder = paymentAllocationOrder
	this.ProductSubType = productSubType
	this.ProductType = productType
	this.Usage = usage
	return &this
}

// NewProductCreateReqWithDefaults instantiates a new ProductCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCreateReqWithDefaults() *ProductCreateReq {
	this := ProductCreateReq{}
	var classification ProductClassification = PRODUCTCLASSIFICATION_CONSUMER
	this.Classification = classification
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	var productSubType ProductSubType = PRODUCTSUBTYPE_CREDIT_CARD
	this.ProductSubType = productSubType
	var productType ProductType = PRODUCTTYPE_REVOLVING
	this.ProductType = productType
	return &this
}

// GetCardProductTokens returns the CardProductTokens field value
func (o *ProductCreateReq) GetCardProductTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CardProductTokens
}

// GetCardProductTokensOk returns a tuple with the CardProductTokens field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetCardProductTokensOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardProductTokens, true
}

// SetCardProductTokens sets field value
func (o *ProductCreateReq) SetCardProductTokens(v []string) {
	o.CardProductTokens = v
}

// GetClassification returns the Classification field value
func (o *ProductCreateReq) GetClassification() ProductClassification {
	if o == nil {
		var ret ProductClassification
		return ret
	}

	return o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetClassificationOk() (*ProductClassification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classification, true
}

// SetClassification sets field value
func (o *ProductCreateReq) SetClassification(v ProductClassification) {
	o.Classification = v
}

// GetConfig returns the Config field value
func (o *ProductCreateReq) GetConfig() ProductConfig {
	if o == nil {
		var ret ProductConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetConfigOk() (*ProductConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ProductCreateReq) SetConfig(v ProductConfig) {
	o.Config = v
}

// GetCreditLine returns the CreditLine field value
func (o *ProductCreateReq) GetCreditLine() ProductCreditLine {
	if o == nil {
		var ret ProductCreditLine
		return ret
	}

	return o.CreditLine
}

// GetCreditLineOk returns a tuple with the CreditLine field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetCreditLineOk() (*ProductCreditLine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditLine, true
}

// SetCreditLine sets field value
func (o *ProductCreateReq) SetCreditLine(v ProductCreditLine) {
	o.CreditLine = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *ProductCreateReq) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *ProductCreateReq) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductCreateReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductCreateReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductCreateReq) SetDescription(v string) {
	o.Description = &v
}

// GetInterestCalculation returns the InterestCalculation field value
func (o *ProductCreateReq) GetInterestCalculation() InterestCalculation {
	if o == nil {
		var ret InterestCalculation
		return ret
	}

	return o.InterestCalculation
}

// GetInterestCalculationOk returns a tuple with the InterestCalculation field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetInterestCalculationOk() (*InterestCalculation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterestCalculation, true
}

// SetInterestCalculation sets field value
func (o *ProductCreateReq) SetInterestCalculation(v InterestCalculation) {
	o.InterestCalculation = v
}

// GetMinPaymentCalculation returns the MinPaymentCalculation field value if set, zero value otherwise.
func (o *ProductCreateReq) GetMinPaymentCalculation() ProductMinPaymentCalculation {
	if o == nil || IsNil(o.MinPaymentCalculation) {
		var ret ProductMinPaymentCalculation
		return ret
	}
	return *o.MinPaymentCalculation
}

// GetMinPaymentCalculationOk returns a tuple with the MinPaymentCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetMinPaymentCalculationOk() (*ProductMinPaymentCalculation, bool) {
	if o == nil || IsNil(o.MinPaymentCalculation) {
		return nil, false
	}
	return o.MinPaymentCalculation, true
}

// HasMinPaymentCalculation returns a boolean if a field has been set.
func (o *ProductCreateReq) HasMinPaymentCalculation() bool {
	if o != nil && !IsNil(o.MinPaymentCalculation) {
		return true
	}

	return false
}

// SetMinPaymentCalculation gets a reference to the given ProductMinPaymentCalculation and assigns it to the MinPaymentCalculation field.
func (o *ProductCreateReq) SetMinPaymentCalculation(v ProductMinPaymentCalculation) {
	o.MinPaymentCalculation = &v
}

// GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field value
func (o *ProductCreateReq) GetMinPaymentFlatAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.MinPaymentFlatAmount
}

// GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetMinPaymentFlatAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPaymentFlatAmount, true
}

// SetMinPaymentFlatAmount sets field value
func (o *ProductCreateReq) SetMinPaymentFlatAmount(v decimal.Decimal) {
	o.MinPaymentFlatAmount = v
}

// GetMinPaymentPercentage returns the MinPaymentPercentage field value
func (o *ProductCreateReq) GetMinPaymentPercentage() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.MinPaymentPercentage
}

// GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetMinPaymentPercentageOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPaymentPercentage, true
}

// SetMinPaymentPercentage sets field value
func (o *ProductCreateReq) SetMinPaymentPercentage(v decimal.Decimal) {
	o.MinPaymentPercentage = v
}

// GetName returns the Name field value
func (o *ProductCreateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductCreateReq) SetName(v string) {
	o.Name = v
}

// GetPaymentAllocationOrder returns the PaymentAllocationOrder field value
func (o *ProductCreateReq) GetPaymentAllocationOrder() []PaymentAllocationOrderEnum {
	if o == nil {
		var ret []PaymentAllocationOrderEnum
		return ret
	}

	return o.PaymentAllocationOrder
}

// GetPaymentAllocationOrderOk returns a tuple with the PaymentAllocationOrder field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetPaymentAllocationOrderOk() ([]PaymentAllocationOrderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAllocationOrder, true
}

// SetPaymentAllocationOrder sets field value
func (o *ProductCreateReq) SetPaymentAllocationOrder(v []PaymentAllocationOrderEnum) {
	o.PaymentAllocationOrder = v
}

// GetProductSubType returns the ProductSubType field value
func (o *ProductCreateReq) GetProductSubType() ProductSubType {
	if o == nil {
		var ret ProductSubType
		return ret
	}

	return o.ProductSubType
}

// GetProductSubTypeOk returns a tuple with the ProductSubType field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetProductSubTypeOk() (*ProductSubType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductSubType, true
}

// SetProductSubType sets field value
func (o *ProductCreateReq) SetProductSubType(v ProductSubType) {
	o.ProductSubType = v
}

// GetProductType returns the ProductType field value
func (o *ProductCreateReq) GetProductType() ProductType {
	if o == nil {
		var ret ProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetProductTypeOk() (*ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ProductCreateReq) SetProductType(v ProductType) {
	o.ProductType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductCreateReq) GetStatus() ResourceStatus {
	if o == nil || IsNil(o.Status) {
		var ret ResourceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetStatusOk() (*ResourceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductCreateReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ResourceStatus and assigns it to the Status field.
func (o *ProductCreateReq) SetStatus(v ResourceStatus) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProductCreateReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProductCreateReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProductCreateReq) SetToken(v string) {
	o.Token = &v
}

// GetUsage returns the Usage field value
func (o *ProductCreateReq) GetUsage() []BalanceType {
	if o == nil {
		var ret []BalanceType
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *ProductCreateReq) GetUsageOk() ([]BalanceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *ProductCreateReq) SetUsage(v []BalanceType) {
	o.Usage = v
}

func (o ProductCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_product_tokens"] = o.CardProductTokens
	toSerialize["classification"] = o.Classification
	toSerialize["config"] = o.Config
	toSerialize["credit_line"] = o.CreditLine
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["interest_calculation"] = o.InterestCalculation
	if !IsNil(o.MinPaymentCalculation) {
		toSerialize["min_payment_calculation"] = o.MinPaymentCalculation
	}
	toSerialize["min_payment_flat_amount"] = o.MinPaymentFlatAmount
	toSerialize["min_payment_percentage"] = o.MinPaymentPercentage
	toSerialize["name"] = o.Name
	toSerialize["payment_allocation_order"] = o.PaymentAllocationOrder
	toSerialize["product_sub_type"] = o.ProductSubType
	toSerialize["product_type"] = o.ProductType
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

func (o *ProductCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_product_tokens",
		"classification",
		"config",
		"credit_line",
		"currency_code",
		"interest_calculation",
		"min_payment_flat_amount",
		"min_payment_percentage",
		"name",
		"payment_allocation_order",
		"product_sub_type",
		"product_type",
		"usage",
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

	varProductCreateReq := _ProductCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductCreateReq)

	if err != nil {
		return err
	}

	*o = ProductCreateReq(varProductCreateReq)

	return err
}

type NullableProductCreateReq struct {
	value *ProductCreateReq
	isSet bool
}

func (v NullableProductCreateReq) Get() *ProductCreateReq {
	return v.value
}

func (v *NullableProductCreateReq) Set(val *ProductCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCreateReq(val *ProductCreateReq) *NullableProductCreateReq {
	return &NullableProductCreateReq{value: val, isSet: true}
}

func (v NullableProductCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


