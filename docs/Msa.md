# Msa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignToken** | **string** |  | 
**ReloadAmount** | **decimal.Decimal** |  | 
**TriggerAmount** | **decimal.Decimal** |  | 

## Methods

### NewMsa

`func NewMsa(campaignToken string, reloadAmount decimal.Decimal, triggerAmount decimal.Decimal, ) *Msa`

NewMsa instantiates a new Msa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaWithDefaults

`func NewMsaWithDefaults() *Msa`

NewMsaWithDefaults instantiates a new Msa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignToken

`func (o *Msa) GetCampaignToken() string`

GetCampaignToken returns the CampaignToken field if non-nil, zero value otherwise.

### GetCampaignTokenOk

`func (o *Msa) GetCampaignTokenOk() (*string, bool)`

GetCampaignTokenOk returns a tuple with the CampaignToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignToken

`func (o *Msa) SetCampaignToken(v string)`

SetCampaignToken sets CampaignToken field to given value.


### GetReloadAmount

`func (o *Msa) GetReloadAmount() decimal.Decimal`

GetReloadAmount returns the ReloadAmount field if non-nil, zero value otherwise.

### GetReloadAmountOk

`func (o *Msa) GetReloadAmountOk() (*decimal.Decimal, bool)`

GetReloadAmountOk returns a tuple with the ReloadAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloadAmount

`func (o *Msa) SetReloadAmount(v decimal.Decimal)`

SetReloadAmount sets ReloadAmount field to given value.


### GetTriggerAmount

`func (o *Msa) GetTriggerAmount() decimal.Decimal`

GetTriggerAmount returns the TriggerAmount field if non-nil, zero value otherwise.

### GetTriggerAmountOk

`func (o *Msa) GetTriggerAmountOk() (*decimal.Decimal, bool)`

GetTriggerAmountOk returns a tuple with the TriggerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAmount

`func (o *Msa) SetTriggerAmount(v decimal.Decimal)`

SetTriggerAmount sets TriggerAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


