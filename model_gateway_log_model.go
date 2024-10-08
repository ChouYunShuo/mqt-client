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

// checks if the GatewayLogModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayLogModel{}

// GatewayLogModel Contains information from the JIT Funding gateway in response to a funding request.
type GatewayLogModel struct {
	// Length of time in milliseconds that the gateway took to respond to a funding request.
	Duration *int64 `json:"duration,omitempty"`
	// Message about the status of the funding request. Useful for determining whether it was approved and completed successfully, declined by the gateway, or timed out.
	Message string `json:"message"`
	// Customer order number, same value as `transaction.token`.
	OrderNumber string `json:"order_number"`
	Response *GatewayResponse `json:"response,omitempty"`
	// Whether the gateway sent a response (`true`) or timed out (`false`).
	TimedOut *bool `json:"timed_out,omitempty"`
	// Customer-defined identifier for the transaction.
	TransactionId string `json:"transaction_id"`
}

type _GatewayLogModel GatewayLogModel

// NewGatewayLogModel instantiates a new GatewayLogModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayLogModel(message string, orderNumber string, transactionId string) *GatewayLogModel {
	this := GatewayLogModel{}
	this.Message = message
	this.OrderNumber = orderNumber
	var timedOut bool = false
	this.TimedOut = &timedOut
	this.TransactionId = transactionId
	return &this
}

// NewGatewayLogModelWithDefaults instantiates a new GatewayLogModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayLogModelWithDefaults() *GatewayLogModel {
	this := GatewayLogModel{}
	var timedOut bool = false
	this.TimedOut = &timedOut
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GatewayLogModel) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLogModel) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GatewayLogModel) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *GatewayLogModel) SetDuration(v int64) {
	o.Duration = &v
}

// GetMessage returns the Message field value
func (o *GatewayLogModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GatewayLogModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GatewayLogModel) SetMessage(v string) {
	o.Message = v
}

// GetOrderNumber returns the OrderNumber field value
func (o *GatewayLogModel) GetOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value
// and a boolean to check if the value has been set.
func (o *GatewayLogModel) GetOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNumber, true
}

// SetOrderNumber sets field value
func (o *GatewayLogModel) SetOrderNumber(v string) {
	o.OrderNumber = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *GatewayLogModel) GetResponse() GatewayResponse {
	if o == nil || IsNil(o.Response) {
		var ret GatewayResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLogModel) GetResponseOk() (*GatewayResponse, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *GatewayLogModel) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given GatewayResponse and assigns it to the Response field.
func (o *GatewayLogModel) SetResponse(v GatewayResponse) {
	o.Response = &v
}

// GetTimedOut returns the TimedOut field value if set, zero value otherwise.
func (o *GatewayLogModel) GetTimedOut() bool {
	if o == nil || IsNil(o.TimedOut) {
		var ret bool
		return ret
	}
	return *o.TimedOut
}

// GetTimedOutOk returns a tuple with the TimedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLogModel) GetTimedOutOk() (*bool, bool) {
	if o == nil || IsNil(o.TimedOut) {
		return nil, false
	}
	return o.TimedOut, true
}

// HasTimedOut returns a boolean if a field has been set.
func (o *GatewayLogModel) HasTimedOut() bool {
	if o != nil && !IsNil(o.TimedOut) {
		return true
	}

	return false
}

// SetTimedOut gets a reference to the given bool and assigns it to the TimedOut field.
func (o *GatewayLogModel) SetTimedOut(v bool) {
	o.TimedOut = &v
}

// GetTransactionId returns the TransactionId field value
func (o *GatewayLogModel) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *GatewayLogModel) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *GatewayLogModel) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o GatewayLogModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayLogModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["message"] = o.Message
	toSerialize["order_number"] = o.OrderNumber
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.TimedOut) {
		toSerialize["timed_out"] = o.TimedOut
	}
	toSerialize["transaction_id"] = o.TransactionId
	return toSerialize, nil
}

func (o *GatewayLogModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"order_number",
		"transaction_id",
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

	varGatewayLogModel := _GatewayLogModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGatewayLogModel)

	if err != nil {
		return err
	}

	*o = GatewayLogModel(varGatewayLogModel)

	return err
}

type NullableGatewayLogModel struct {
	value *GatewayLogModel
	isSet bool
}

func (v NullableGatewayLogModel) Get() *GatewayLogModel {
	return v.value
}

func (v *NullableGatewayLogModel) Set(val *GatewayLogModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayLogModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayLogModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayLogModel(val *GatewayLogModel) *NullableGatewayLogModel {
	return &NullableGatewayLogModel{value: val, isSet: true}
}

func (v NullableGatewayLogModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayLogModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


