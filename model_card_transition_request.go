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
)

// checks if the CardTransitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardTransitionRequest{}

// CardTransitionRequest struct for CardTransitionRequest
type CardTransitionRequest struct {
	// Identifies the card whose state will transition.
	CardToken string `json:"card_token"`
	// The mechanism by which the transition was initiated.  * *ADMIN* - Indicates that the card transition was initiated through the Marqeta Dashboard. * *API* - Indicates that the card transition was initiated by you through the Core API. Use this value when creating a card transition with an API `POST` request. * *FRAUD* - Indicates that either Marqeta or the card network has determined that the card is fraudulent. * *IVR* - Indicates that the card transition was initiated through your Interactive Voice Response system. * *SYSTEM* - Indicates that the card transition was initiated by Marqeta. For example, Marqeta suspended the card due to excessive failed personal identification number (PIN) entries.
	Channel string `json:"channel"`
	// Additional information about the state change.
	Reason *string `json:"reason,omitempty"`
	// Standard code describing the reason for the transition.  *NOTE:* This field is required if your program uses v2 of the `user_card_state_version`, which is a program-specific configuration value that is managed by Marqeta and cannot be accessed via the API. To learn more about the `user_card_state_version` program configuration, contact your Marqeta representative.  * *00:* Object activated for the first time * *01:* Requested by you * *02:* Inactivity over time * *03:* This address cannot accept mail or the addressee is unknown * *04:* Negative account balance * *05:* Account under review * *06:* Suspicious activity was identified * *07:* Activity outside the program parameters was identified * *08:* Confirmed fraud was identified * *09:* Matched with an Office of Foreign Assets Control list * *10:* Card was reported lost * *11:* Card information was cloned * *12:* Account or card information was compromised * *13:* Temporary status change while on hold/leave * *14:* Initiated by Marqeta * *15:* Initiated by issuer * *16:* Card expired * *17:* Failed KYC * *18:* Changed to `ACTIVE` because information was properly validated * *19:* Changed to `ACTIVE` because account activity was properly validated * *20:* Change occurred prior to the normalization of reason codes * *21:* Initiated by a third party, often a digital wallet provider * *22:* PIN retry limit reached * *23:* Card was reported stolen * *24:* Address issue * *25:* Name issue * *26:* SSN issue * *27:* DOB issue * *28:* Email issue * *29:* Phone issue * *30:* Account/fulfillment mismatch * *31:* Other reason
	ReasonCode *string `json:"reason_code,omitempty"`
	// Specifies the new state.
	State string `json:"state"`
	// Set this field to `true` to synchronize the state of the card's associated token(s) with the card's new state. The digital wallet tokens must be in a valid starting state for the given transition, which will reflect the card's state transition. For example, if the card is transitioning from the `ACTIVE` state to the `SUSPENDED` state, only digital wallet tokens in the `ACTIVE` state will be synchronized with the card state transition and therefore be transitioned to the `SUSPENDED` state.  Leave this field blank or set it to `false` to keep the states of the card and its digital wallet tokens independent.
	SyncStateWithDwts *bool `json:"sync_state_with_dwts,omitempty"`
	// Unique identifier of the card transition.  If you do not include a token, the system will generate one automatically. This token is referenced in other API calls, so we recommend that you define a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	Validations *ValidationsRequest `json:"validations,omitempty"`
}

type _CardTransitionRequest CardTransitionRequest

// NewCardTransitionRequest instantiates a new CardTransitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTransitionRequest(cardToken string, channel string, state string) *CardTransitionRequest {
	this := CardTransitionRequest{}
	this.CardToken = cardToken
	this.Channel = channel
	this.State = state
	return &this
}

// NewCardTransitionRequestWithDefaults instantiates a new CardTransitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTransitionRequestWithDefaults() *CardTransitionRequest {
	this := CardTransitionRequest{}
	return &this
}

// GetCardToken returns the CardToken field value
func (o *CardTransitionRequest) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *CardTransitionRequest) SetCardToken(v string) {
	o.CardToken = v
}

// GetChannel returns the Channel field value
func (o *CardTransitionRequest) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CardTransitionRequest) SetChannel(v string) {
	o.Channel = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CardTransitionRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CardTransitionRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CardTransitionRequest) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *CardTransitionRequest) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *CardTransitionRequest) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *CardTransitionRequest) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetState returns the State field value
func (o *CardTransitionRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CardTransitionRequest) SetState(v string) {
	o.State = v
}

// GetSyncStateWithDwts returns the SyncStateWithDwts field value if set, zero value otherwise.
func (o *CardTransitionRequest) GetSyncStateWithDwts() bool {
	if o == nil || IsNil(o.SyncStateWithDwts) {
		var ret bool
		return ret
	}
	return *o.SyncStateWithDwts
}

// GetSyncStateWithDwtsOk returns a tuple with the SyncStateWithDwts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetSyncStateWithDwtsOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncStateWithDwts) {
		return nil, false
	}
	return o.SyncStateWithDwts, true
}

// HasSyncStateWithDwts returns a boolean if a field has been set.
func (o *CardTransitionRequest) HasSyncStateWithDwts() bool {
	if o != nil && !IsNil(o.SyncStateWithDwts) {
		return true
	}

	return false
}

// SetSyncStateWithDwts gets a reference to the given bool and assigns it to the SyncStateWithDwts field.
func (o *CardTransitionRequest) SetSyncStateWithDwts(v bool) {
	o.SyncStateWithDwts = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CardTransitionRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CardTransitionRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CardTransitionRequest) SetToken(v string) {
	o.Token = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *CardTransitionRequest) GetValidations() ValidationsRequest {
	if o == nil || IsNil(o.Validations) {
		var ret ValidationsRequest
		return ret
	}
	return *o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransitionRequest) GetValidationsOk() (*ValidationsRequest, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *CardTransitionRequest) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given ValidationsRequest and assigns it to the Validations field.
func (o *CardTransitionRequest) SetValidations(v ValidationsRequest) {
	o.Validations = &v
}

func (o CardTransitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardTransitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_token"] = o.CardToken
	toSerialize["channel"] = o.Channel
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	toSerialize["state"] = o.State
	if !IsNil(o.SyncStateWithDwts) {
		toSerialize["sync_state_with_dwts"] = o.SyncStateWithDwts
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	return toSerialize, nil
}

func (o *CardTransitionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_token",
		"channel",
		"state",
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

	varCardTransitionRequest := _CardTransitionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardTransitionRequest)

	if err != nil {
		return err
	}

	*o = CardTransitionRequest(varCardTransitionRequest)

	return err
}

type NullableCardTransitionRequest struct {
	value *CardTransitionRequest
	isSet bool
}

func (v NullableCardTransitionRequest) Get() *CardTransitionRequest {
	return v.value
}

func (v *NullableCardTransitionRequest) Set(val *CardTransitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTransitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTransitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTransitionRequest(val *CardTransitionRequest) *NullableCardTransitionRequest {
	return &NullableCardTransitionRequest{value: val, isSet: true}
}

func (v NullableCardTransitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTransitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


