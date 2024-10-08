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
)

// checks if the AchResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AchResponseModel{}

// AchResponseModel struct for AchResponseModel
type AchResponseModel struct {
	// ACH account identifier appended to the bank account number.
	AccountSuffix string `json:"account_suffix"`
	// Type of account.
	AccountType string `json:"account_type"`
	// Specifies whether the account is active.
	Active bool `json:"active"`
	// Name of the bank holding the account. This field is returned if it exists in the resource.
	BankName *string `json:"bank_name,omitempty"`
	// Unique identifier of the business account holder. This token is returned if a `user_token` is not specified.
	BusinessToken *string `json:"business_token,omitempty"`
	// Date and time when the resource was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Date and time in UTC when either the request for account validation was sent to the third-party partner, or when the funding source was verified by microdeposits.  This field is returned if it exists in the resource.
	DateSentForVerification *time.Time `json:"date_sent_for_verification,omitempty"`
	// Date and time when the account was verified, in UTC.  This field is returned if it exists in the resource.
	DateVerified *time.Time `json:"date_verified,omitempty"`
	// If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.  This field is returned if it exists in the resource.
	IsDefaultAccount *bool `json:"is_default_account,omitempty"`
	// Date and time when the resource was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Name on the ACH account.
	NameOnAccount string `json:"name_on_account"`
	// Name of the partner who validated the account holder. Returned when a third-party partner was used for account validation.
	Partner *string `json:"partner,omitempty"`
	// Supplied by the account validation partner, this value is a reference to the account holder's details, such as the account number and routing number. Returned when a third-party partner was used for account validation.
	PartnerAccountLinkReferenceToken *string `json:"partner_account_link_reference_token,omitempty"`
	// Unique identifier of the funding source.
	Token string `json:"token"`
	// Unique identifier of the user account holder. This token is returned if a `business_token` is not specified.
	UserToken *string `json:"user_token,omitempty"`
	// Free-form text field for holding notes about verification. This field is returned only if `verification_override = true`.
	VerificationNotes *string `json:"verification_notes,omitempty"`
	// Allows the ACH funding source to be used regardless of its verification status. This field is returned if it exists in the resource.  *NOTE:* When using `PLAID` to validate a funding source, this field is always set to `true`.
	VerificationOverride *bool `json:"verification_override,omitempty"`
	// Account verification status. This field is returned if it exists in the resource.
	VerificationStatus *string `json:"verification_status,omitempty"`
}

type _AchResponseModel AchResponseModel

// NewAchResponseModel instantiates a new AchResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchResponseModel(accountSuffix string, accountType string, active bool, createdTime time.Time, lastModifiedTime time.Time, nameOnAccount string, token string) *AchResponseModel {
	this := AchResponseModel{}
	this.AccountSuffix = accountSuffix
	this.AccountType = accountType
	this.Active = active
	this.CreatedTime = createdTime
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	this.LastModifiedTime = lastModifiedTime
	this.NameOnAccount = nameOnAccount
	this.Token = token
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// NewAchResponseModelWithDefaults instantiates a new AchResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchResponseModelWithDefaults() *AchResponseModel {
	this := AchResponseModel{}
	var active bool = false
	this.Active = active
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// GetAccountSuffix returns the AccountSuffix field value
func (o *AchResponseModel) GetAccountSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSuffix
}

// GetAccountSuffixOk returns a tuple with the AccountSuffix field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetAccountSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSuffix, true
}

// SetAccountSuffix sets field value
func (o *AchResponseModel) SetAccountSuffix(v string) {
	o.AccountSuffix = v
}

// GetAccountType returns the AccountType field value
func (o *AchResponseModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *AchResponseModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetActive returns the Active field value
func (o *AchResponseModel) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *AchResponseModel) SetActive(v bool) {
	o.Active = v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *AchResponseModel) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *AchResponseModel) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *AchResponseModel) SetBankName(v string) {
	o.BankName = &v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *AchResponseModel) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *AchResponseModel) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *AchResponseModel) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *AchResponseModel) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *AchResponseModel) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetDateSentForVerification returns the DateSentForVerification field value if set, zero value otherwise.
func (o *AchResponseModel) GetDateSentForVerification() time.Time {
	if o == nil || IsNil(o.DateSentForVerification) {
		var ret time.Time
		return ret
	}
	return *o.DateSentForVerification
}

// GetDateSentForVerificationOk returns a tuple with the DateSentForVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetDateSentForVerificationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateSentForVerification) {
		return nil, false
	}
	return o.DateSentForVerification, true
}

// HasDateSentForVerification returns a boolean if a field has been set.
func (o *AchResponseModel) HasDateSentForVerification() bool {
	if o != nil && !IsNil(o.DateSentForVerification) {
		return true
	}

	return false
}

// SetDateSentForVerification gets a reference to the given time.Time and assigns it to the DateSentForVerification field.
func (o *AchResponseModel) SetDateSentForVerification(v time.Time) {
	o.DateSentForVerification = &v
}

// GetDateVerified returns the DateVerified field value if set, zero value otherwise.
func (o *AchResponseModel) GetDateVerified() time.Time {
	if o == nil || IsNil(o.DateVerified) {
		var ret time.Time
		return ret
	}
	return *o.DateVerified
}

// GetDateVerifiedOk returns a tuple with the DateVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetDateVerifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateVerified) {
		return nil, false
	}
	return o.DateVerified, true
}

// HasDateVerified returns a boolean if a field has been set.
func (o *AchResponseModel) HasDateVerified() bool {
	if o != nil && !IsNil(o.DateVerified) {
		return true
	}

	return false
}

// SetDateVerified gets a reference to the given time.Time and assigns it to the DateVerified field.
func (o *AchResponseModel) SetDateVerified(v time.Time) {
	o.DateVerified = &v
}

// GetIsDefaultAccount returns the IsDefaultAccount field value if set, zero value otherwise.
func (o *AchResponseModel) GetIsDefaultAccount() bool {
	if o == nil || IsNil(o.IsDefaultAccount) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccount
}

// GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetIsDefaultAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccount) {
		return nil, false
	}
	return o.IsDefaultAccount, true
}

// HasIsDefaultAccount returns a boolean if a field has been set.
func (o *AchResponseModel) HasIsDefaultAccount() bool {
	if o != nil && !IsNil(o.IsDefaultAccount) {
		return true
	}

	return false
}

// SetIsDefaultAccount gets a reference to the given bool and assigns it to the IsDefaultAccount field.
func (o *AchResponseModel) SetIsDefaultAccount(v bool) {
	o.IsDefaultAccount = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *AchResponseModel) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *AchResponseModel) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetNameOnAccount returns the NameOnAccount field value
func (o *AchResponseModel) GetNameOnAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameOnAccount
}

// GetNameOnAccountOk returns a tuple with the NameOnAccount field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetNameOnAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameOnAccount, true
}

// SetNameOnAccount sets field value
func (o *AchResponseModel) SetNameOnAccount(v string) {
	o.NameOnAccount = v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *AchResponseModel) GetPartner() string {
	if o == nil || IsNil(o.Partner) {
		var ret string
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetPartnerOk() (*string, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *AchResponseModel) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given string and assigns it to the Partner field.
func (o *AchResponseModel) SetPartner(v string) {
	o.Partner = &v
}

// GetPartnerAccountLinkReferenceToken returns the PartnerAccountLinkReferenceToken field value if set, zero value otherwise.
func (o *AchResponseModel) GetPartnerAccountLinkReferenceToken() string {
	if o == nil || IsNil(o.PartnerAccountLinkReferenceToken) {
		var ret string
		return ret
	}
	return *o.PartnerAccountLinkReferenceToken
}

// GetPartnerAccountLinkReferenceTokenOk returns a tuple with the PartnerAccountLinkReferenceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetPartnerAccountLinkReferenceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerAccountLinkReferenceToken) {
		return nil, false
	}
	return o.PartnerAccountLinkReferenceToken, true
}

// HasPartnerAccountLinkReferenceToken returns a boolean if a field has been set.
func (o *AchResponseModel) HasPartnerAccountLinkReferenceToken() bool {
	if o != nil && !IsNil(o.PartnerAccountLinkReferenceToken) {
		return true
	}

	return false
}

// SetPartnerAccountLinkReferenceToken gets a reference to the given string and assigns it to the PartnerAccountLinkReferenceToken field.
func (o *AchResponseModel) SetPartnerAccountLinkReferenceToken(v string) {
	o.PartnerAccountLinkReferenceToken = &v
}

// GetToken returns the Token field value
func (o *AchResponseModel) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AchResponseModel) SetToken(v string) {
	o.Token = v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *AchResponseModel) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *AchResponseModel) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *AchResponseModel) SetUserToken(v string) {
	o.UserToken = &v
}

// GetVerificationNotes returns the VerificationNotes field value if set, zero value otherwise.
func (o *AchResponseModel) GetVerificationNotes() string {
	if o == nil || IsNil(o.VerificationNotes) {
		var ret string
		return ret
	}
	return *o.VerificationNotes
}

// GetVerificationNotesOk returns a tuple with the VerificationNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetVerificationNotesOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationNotes) {
		return nil, false
	}
	return o.VerificationNotes, true
}

// HasVerificationNotes returns a boolean if a field has been set.
func (o *AchResponseModel) HasVerificationNotes() bool {
	if o != nil && !IsNil(o.VerificationNotes) {
		return true
	}

	return false
}

// SetVerificationNotes gets a reference to the given string and assigns it to the VerificationNotes field.
func (o *AchResponseModel) SetVerificationNotes(v string) {
	o.VerificationNotes = &v
}

// GetVerificationOverride returns the VerificationOverride field value if set, zero value otherwise.
func (o *AchResponseModel) GetVerificationOverride() bool {
	if o == nil || IsNil(o.VerificationOverride) {
		var ret bool
		return ret
	}
	return *o.VerificationOverride
}

// GetVerificationOverrideOk returns a tuple with the VerificationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetVerificationOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.VerificationOverride) {
		return nil, false
	}
	return o.VerificationOverride, true
}

// HasVerificationOverride returns a boolean if a field has been set.
func (o *AchResponseModel) HasVerificationOverride() bool {
	if o != nil && !IsNil(o.VerificationOverride) {
		return true
	}

	return false
}

// SetVerificationOverride gets a reference to the given bool and assigns it to the VerificationOverride field.
func (o *AchResponseModel) SetVerificationOverride(v bool) {
	o.VerificationOverride = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AchResponseModel) GetVerificationStatus() string {
	if o == nil || IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchResponseModel) GetVerificationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AchResponseModel) HasVerificationStatus() bool {
	if o != nil && !IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AchResponseModel) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o AchResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AchResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_suffix"] = o.AccountSuffix
	toSerialize["account_type"] = o.AccountType
	toSerialize["active"] = o.Active
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.DateSentForVerification) {
		toSerialize["date_sent_for_verification"] = o.DateSentForVerification
	}
	if !IsNil(o.DateVerified) {
		toSerialize["date_verified"] = o.DateVerified
	}
	if !IsNil(o.IsDefaultAccount) {
		toSerialize["is_default_account"] = o.IsDefaultAccount
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["name_on_account"] = o.NameOnAccount
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.PartnerAccountLinkReferenceToken) {
		toSerialize["partner_account_link_reference_token"] = o.PartnerAccountLinkReferenceToken
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	if !IsNil(o.VerificationNotes) {
		toSerialize["verification_notes"] = o.VerificationNotes
	}
	if !IsNil(o.VerificationOverride) {
		toSerialize["verification_override"] = o.VerificationOverride
	}
	if !IsNil(o.VerificationStatus) {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return toSerialize, nil
}

func (o *AchResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_suffix",
		"account_type",
		"active",
		"created_time",
		"last_modified_time",
		"name_on_account",
		"token",
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

	varAchResponseModel := _AchResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAchResponseModel)

	if err != nil {
		return err
	}

	*o = AchResponseModel(varAchResponseModel)

	return err
}

type NullableAchResponseModel struct {
	value *AchResponseModel
	isSet bool
}

func (v NullableAchResponseModel) Get() *AchResponseModel {
	return v.value
}

func (v *NullableAchResponseModel) Set(val *AchResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAchResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAchResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchResponseModel(val *AchResponseModel) *NullableAchResponseModel {
	return &NullableAchResponseModel{value: val, isSet: true}
}

func (v NullableAchResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


