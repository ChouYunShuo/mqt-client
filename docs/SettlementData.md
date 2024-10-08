# SettlementData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **decimal.Decimal** | The settled amount. | [optional] 
**ConversionRate** | Pointer to **decimal.Decimal** | Returned when the transaction currency is different from the origination currency.  Conversion rate between the origination currency and the settlement currency. | [optional] 
**CurrencyCode** | Pointer to **string** | The ISO 4217 code of the currency used in the transaction. | [optional] 

## Methods

### NewSettlementData

`func NewSettlementData() *SettlementData`

NewSettlementData instantiates a new SettlementData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementDataWithDefaults

`func NewSettlementDataWithDefaults() *SettlementData`

NewSettlementDataWithDefaults instantiates a new SettlementData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SettlementData) GetAmount() decimal.Decimal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SettlementData) GetAmountOk() (*decimal.Decimal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SettlementData) SetAmount(v decimal.Decimal)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SettlementData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetConversionRate

`func (o *SettlementData) GetConversionRate() decimal.Decimal`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *SettlementData) GetConversionRateOk() (*decimal.Decimal, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *SettlementData) SetConversionRate(v decimal.Decimal)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *SettlementData) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SettlementData) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SettlementData) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SettlementData) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SettlementData) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


