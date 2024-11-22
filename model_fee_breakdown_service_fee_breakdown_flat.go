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

// checks if the FeeBreakdownServiceFeeBreakdownFlat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeBreakdownServiceFeeBreakdownFlat{}

// FeeBreakdownServiceFeeBreakdownFlat struct for FeeBreakdownServiceFeeBreakdownFlat
type FeeBreakdownServiceFeeBreakdownFlat struct {
	// The amount of the service fee.
	Amount *float32 `json:"amount,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
}

// NewFeeBreakdownServiceFeeBreakdownFlat instantiates a new FeeBreakdownServiceFeeBreakdownFlat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeBreakdownServiceFeeBreakdownFlat() *FeeBreakdownServiceFeeBreakdownFlat {
	this := FeeBreakdownServiceFeeBreakdownFlat{}
	return &this
}

// NewFeeBreakdownServiceFeeBreakdownFlatWithDefaults instantiates a new FeeBreakdownServiceFeeBreakdownFlat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeBreakdownServiceFeeBreakdownFlatWithDefaults() *FeeBreakdownServiceFeeBreakdownFlat {
	this := FeeBreakdownServiceFeeBreakdownFlat{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *FeeBreakdownServiceFeeBreakdownFlat) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeBreakdownServiceFeeBreakdownFlat) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *FeeBreakdownServiceFeeBreakdownFlat) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *FeeBreakdownServiceFeeBreakdownFlat) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *FeeBreakdownServiceFeeBreakdownFlat) GetCurrency() Currency {
	if o == nil || IsNil(o.Currency) {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeBreakdownServiceFeeBreakdownFlat) GetCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *FeeBreakdownServiceFeeBreakdownFlat) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *FeeBreakdownServiceFeeBreakdownFlat) SetCurrency(v Currency) {
	o.Currency = &v
}

func (o FeeBreakdownServiceFeeBreakdownFlat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeBreakdownServiceFeeBreakdownFlat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableFeeBreakdownServiceFeeBreakdownFlat struct {
	value *FeeBreakdownServiceFeeBreakdownFlat
	isSet bool
}

func (v NullableFeeBreakdownServiceFeeBreakdownFlat) Get() *FeeBreakdownServiceFeeBreakdownFlat {
	return v.value
}

func (v *NullableFeeBreakdownServiceFeeBreakdownFlat) Set(val *FeeBreakdownServiceFeeBreakdownFlat) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeBreakdownServiceFeeBreakdownFlat) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeBreakdownServiceFeeBreakdownFlat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeBreakdownServiceFeeBreakdownFlat(val *FeeBreakdownServiceFeeBreakdownFlat) *NullableFeeBreakdownServiceFeeBreakdownFlat {
	return &NullableFeeBreakdownServiceFeeBreakdownFlat{value: val, isSet: true}
}

func (v NullableFeeBreakdownServiceFeeBreakdownFlat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeBreakdownServiceFeeBreakdownFlat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


