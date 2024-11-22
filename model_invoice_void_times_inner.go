/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sellix

import (
	"encoding/json"
)

// checks if the InvoiceVoidTimesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceVoidTimesInner{}

// InvoiceVoidTimesInner struct for InvoiceVoidTimesInner
type InvoiceVoidTimesInner struct {
	Gateways []string `json:"gateways,omitempty"`
	Conf *InvoiceVoidTimesInnerConf `json:"conf,omitempty"`
}

// NewInvoiceVoidTimesInner instantiates a new InvoiceVoidTimesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceVoidTimesInner() *InvoiceVoidTimesInner {
	this := InvoiceVoidTimesInner{}
	return &this
}

// NewInvoiceVoidTimesInnerWithDefaults instantiates a new InvoiceVoidTimesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceVoidTimesInnerWithDefaults() *InvoiceVoidTimesInner {
	this := InvoiceVoidTimesInner{}
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *InvoiceVoidTimesInner) GetGateways() []string {
	if o == nil || IsNil(o.Gateways) {
		var ret []string
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceVoidTimesInner) GetGatewaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *InvoiceVoidTimesInner) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []string and assigns it to the Gateways field.
func (o *InvoiceVoidTimesInner) SetGateways(v []string) {
	o.Gateways = v
}

// GetConf returns the Conf field value if set, zero value otherwise.
func (o *InvoiceVoidTimesInner) GetConf() InvoiceVoidTimesInnerConf {
	if o == nil || IsNil(o.Conf) {
		var ret InvoiceVoidTimesInnerConf
		return ret
	}
	return *o.Conf
}

// GetConfOk returns a tuple with the Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceVoidTimesInner) GetConfOk() (*InvoiceVoidTimesInnerConf, bool) {
	if o == nil || IsNil(o.Conf) {
		return nil, false
	}
	return o.Conf, true
}

// HasConf returns a boolean if a field has been set.
func (o *InvoiceVoidTimesInner) HasConf() bool {
	if o != nil && !IsNil(o.Conf) {
		return true
	}

	return false
}

// SetConf gets a reference to the given InvoiceVoidTimesInnerConf and assigns it to the Conf field.
func (o *InvoiceVoidTimesInner) SetConf(v InvoiceVoidTimesInnerConf) {
	o.Conf = &v
}

func (o InvoiceVoidTimesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceVoidTimesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	if !IsNil(o.Conf) {
		toSerialize["conf"] = o.Conf
	}
	return toSerialize, nil
}

type NullableInvoiceVoidTimesInner struct {
	value *InvoiceVoidTimesInner
	isSet bool
}

func (v NullableInvoiceVoidTimesInner) Get() *InvoiceVoidTimesInner {
	return v.value
}

func (v *NullableInvoiceVoidTimesInner) Set(val *InvoiceVoidTimesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceVoidTimesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceVoidTimesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceVoidTimesInner(val *InvoiceVoidTimesInner) *NullableInvoiceVoidTimesInner {
	return &NullableInvoiceVoidTimesInner{value: val, isSet: true}
}

func (v NullableInvoiceVoidTimesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceVoidTimesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


