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

// checks if the AmlWalletBlockchainInfoClusterOutflowValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmlWalletBlockchainInfoClusterOutflowValue{}

// AmlWalletBlockchainInfoClusterOutflowValue struct for AmlWalletBlockchainInfoClusterOutflowValue
type AmlWalletBlockchainInfoClusterOutflowValue struct {
	// Inflow value in USD.
	Usd *float64 `json:"usd,omitempty"`
}

// NewAmlWalletBlockchainInfoClusterOutflowValue instantiates a new AmlWalletBlockchainInfoClusterOutflowValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmlWalletBlockchainInfoClusterOutflowValue() *AmlWalletBlockchainInfoClusterOutflowValue {
	this := AmlWalletBlockchainInfoClusterOutflowValue{}
	return &this
}

// NewAmlWalletBlockchainInfoClusterOutflowValueWithDefaults instantiates a new AmlWalletBlockchainInfoClusterOutflowValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmlWalletBlockchainInfoClusterOutflowValueWithDefaults() *AmlWalletBlockchainInfoClusterOutflowValue {
	this := AmlWalletBlockchainInfoClusterOutflowValue{}
	return &this
}

// GetUsd returns the Usd field value if set, zero value otherwise.
func (o *AmlWalletBlockchainInfoClusterOutflowValue) GetUsd() float64 {
	if o == nil || IsNil(o.Usd) {
		var ret float64
		return ret
	}
	return *o.Usd
}

// GetUsdOk returns a tuple with the Usd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletBlockchainInfoClusterOutflowValue) GetUsdOk() (*float64, bool) {
	if o == nil || IsNil(o.Usd) {
		return nil, false
	}
	return o.Usd, true
}

// HasUsd returns a boolean if a field has been set.
func (o *AmlWalletBlockchainInfoClusterOutflowValue) HasUsd() bool {
	if o != nil && !IsNil(o.Usd) {
		return true
	}

	return false
}

// SetUsd gets a reference to the given float64 and assigns it to the Usd field.
func (o *AmlWalletBlockchainInfoClusterOutflowValue) SetUsd(v float64) {
	o.Usd = &v
}

func (o AmlWalletBlockchainInfoClusterOutflowValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmlWalletBlockchainInfoClusterOutflowValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Usd) {
		toSerialize["usd"] = o.Usd
	}
	return toSerialize, nil
}

type NullableAmlWalletBlockchainInfoClusterOutflowValue struct {
	value *AmlWalletBlockchainInfoClusterOutflowValue
	isSet bool
}

func (v NullableAmlWalletBlockchainInfoClusterOutflowValue) Get() *AmlWalletBlockchainInfoClusterOutflowValue {
	return v.value
}

func (v *NullableAmlWalletBlockchainInfoClusterOutflowValue) Set(val *AmlWalletBlockchainInfoClusterOutflowValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAmlWalletBlockchainInfoClusterOutflowValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAmlWalletBlockchainInfoClusterOutflowValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmlWalletBlockchainInfoClusterOutflowValue(val *AmlWalletBlockchainInfoClusterOutflowValue) *NullableAmlWalletBlockchainInfoClusterOutflowValue {
	return &NullableAmlWalletBlockchainInfoClusterOutflowValue{value: val, isSet: true}
}

func (v NullableAmlWalletBlockchainInfoClusterOutflowValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmlWalletBlockchainInfoClusterOutflowValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

