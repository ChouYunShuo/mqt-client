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
)

// checks if the Acquirer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Acquirer{}

// Acquirer Contains information about the merchant's financial institution.
type Acquirer struct {
	// Country code of the merchant's financial institution.
	InstitutionCountry *string `json:"institution_country,omitempty"`
	// Identifier code of the merchant's financial institution.
	InstitutionIdCode *string `json:"institution_id_code,omitempty"`
	// The international network identifier.
	NetworkInternationalId *string `json:"network_international_id,omitempty"`
	// Retrieval reference number of the merchant's financial institution.
	RetrievalReferenceNumber *string `json:"retrieval_reference_number,omitempty"`
	// System trace audit number of the merchant's financial institution.
	SystemTraceAuditNumber *string `json:"system_trace_audit_number,omitempty"`
}

// NewAcquirer instantiates a new Acquirer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquirer() *Acquirer {
	this := Acquirer{}
	return &this
}

// NewAcquirerWithDefaults instantiates a new Acquirer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquirerWithDefaults() *Acquirer {
	this := Acquirer{}
	return &this
}

// GetInstitutionCountry returns the InstitutionCountry field value if set, zero value otherwise.
func (o *Acquirer) GetInstitutionCountry() string {
	if o == nil || IsNil(o.InstitutionCountry) {
		var ret string
		return ret
	}
	return *o.InstitutionCountry
}

// GetInstitutionCountryOk returns a tuple with the InstitutionCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Acquirer) GetInstitutionCountryOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionCountry) {
		return nil, false
	}
	return o.InstitutionCountry, true
}

// HasInstitutionCountry returns a boolean if a field has been set.
func (o *Acquirer) HasInstitutionCountry() bool {
	if o != nil && !IsNil(o.InstitutionCountry) {
		return true
	}

	return false
}

// SetInstitutionCountry gets a reference to the given string and assigns it to the InstitutionCountry field.
func (o *Acquirer) SetInstitutionCountry(v string) {
	o.InstitutionCountry = &v
}

// GetInstitutionIdCode returns the InstitutionIdCode field value if set, zero value otherwise.
func (o *Acquirer) GetInstitutionIdCode() string {
	if o == nil || IsNil(o.InstitutionIdCode) {
		var ret string
		return ret
	}
	return *o.InstitutionIdCode
}

// GetInstitutionIdCodeOk returns a tuple with the InstitutionIdCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Acquirer) GetInstitutionIdCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionIdCode) {
		return nil, false
	}
	return o.InstitutionIdCode, true
}

// HasInstitutionIdCode returns a boolean if a field has been set.
func (o *Acquirer) HasInstitutionIdCode() bool {
	if o != nil && !IsNil(o.InstitutionIdCode) {
		return true
	}

	return false
}

// SetInstitutionIdCode gets a reference to the given string and assigns it to the InstitutionIdCode field.
func (o *Acquirer) SetInstitutionIdCode(v string) {
	o.InstitutionIdCode = &v
}

// GetNetworkInternationalId returns the NetworkInternationalId field value if set, zero value otherwise.
func (o *Acquirer) GetNetworkInternationalId() string {
	if o == nil || IsNil(o.NetworkInternationalId) {
		var ret string
		return ret
	}
	return *o.NetworkInternationalId
}

// GetNetworkInternationalIdOk returns a tuple with the NetworkInternationalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Acquirer) GetNetworkInternationalIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkInternationalId) {
		return nil, false
	}
	return o.NetworkInternationalId, true
}

// HasNetworkInternationalId returns a boolean if a field has been set.
func (o *Acquirer) HasNetworkInternationalId() bool {
	if o != nil && !IsNil(o.NetworkInternationalId) {
		return true
	}

	return false
}

// SetNetworkInternationalId gets a reference to the given string and assigns it to the NetworkInternationalId field.
func (o *Acquirer) SetNetworkInternationalId(v string) {
	o.NetworkInternationalId = &v
}

// GetRetrievalReferenceNumber returns the RetrievalReferenceNumber field value if set, zero value otherwise.
func (o *Acquirer) GetRetrievalReferenceNumber() string {
	if o == nil || IsNil(o.RetrievalReferenceNumber) {
		var ret string
		return ret
	}
	return *o.RetrievalReferenceNumber
}

// GetRetrievalReferenceNumberOk returns a tuple with the RetrievalReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Acquirer) GetRetrievalReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RetrievalReferenceNumber) {
		return nil, false
	}
	return o.RetrievalReferenceNumber, true
}

// HasRetrievalReferenceNumber returns a boolean if a field has been set.
func (o *Acquirer) HasRetrievalReferenceNumber() bool {
	if o != nil && !IsNil(o.RetrievalReferenceNumber) {
		return true
	}

	return false
}

// SetRetrievalReferenceNumber gets a reference to the given string and assigns it to the RetrievalReferenceNumber field.
func (o *Acquirer) SetRetrievalReferenceNumber(v string) {
	o.RetrievalReferenceNumber = &v
}

// GetSystemTraceAuditNumber returns the SystemTraceAuditNumber field value if set, zero value otherwise.
func (o *Acquirer) GetSystemTraceAuditNumber() string {
	if o == nil || IsNil(o.SystemTraceAuditNumber) {
		var ret string
		return ret
	}
	return *o.SystemTraceAuditNumber
}

// GetSystemTraceAuditNumberOk returns a tuple with the SystemTraceAuditNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Acquirer) GetSystemTraceAuditNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SystemTraceAuditNumber) {
		return nil, false
	}
	return o.SystemTraceAuditNumber, true
}

// HasSystemTraceAuditNumber returns a boolean if a field has been set.
func (o *Acquirer) HasSystemTraceAuditNumber() bool {
	if o != nil && !IsNil(o.SystemTraceAuditNumber) {
		return true
	}

	return false
}

// SetSystemTraceAuditNumber gets a reference to the given string and assigns it to the SystemTraceAuditNumber field.
func (o *Acquirer) SetSystemTraceAuditNumber(v string) {
	o.SystemTraceAuditNumber = &v
}

func (o Acquirer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Acquirer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstitutionCountry) {
		toSerialize["institution_country"] = o.InstitutionCountry
	}
	if !IsNil(o.InstitutionIdCode) {
		toSerialize["institution_id_code"] = o.InstitutionIdCode
	}
	if !IsNil(o.NetworkInternationalId) {
		toSerialize["network_international_id"] = o.NetworkInternationalId
	}
	if !IsNil(o.RetrievalReferenceNumber) {
		toSerialize["retrieval_reference_number"] = o.RetrievalReferenceNumber
	}
	if !IsNil(o.SystemTraceAuditNumber) {
		toSerialize["system_trace_audit_number"] = o.SystemTraceAuditNumber
	}
	return toSerialize, nil
}

type NullableAcquirer struct {
	value *Acquirer
	isSet bool
}

func (v NullableAcquirer) Get() *Acquirer {
	return v.value
}

func (v *NullableAcquirer) Set(val *Acquirer) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquirer) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquirer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquirer(val *Acquirer) *NullableAcquirer {
	return &NullableAcquirer{value: val, isSet: true}
}

func (v NullableAcquirer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquirer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


