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

// checks if the FeeBreakdownServiceFeeBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeBreakdownServiceFeeBreakdown{}

// FeeBreakdownServiceFeeBreakdown struct for FeeBreakdownServiceFeeBreakdown
type FeeBreakdownServiceFeeBreakdown struct {
	Flat *FeeBreakdownServiceFeeBreakdownFlat `json:"flat,omitempty"`
	Percentage *FeeBreakdownServiceFeeBreakdownPercentage `json:"percentage,omitempty"`
}

// NewFeeBreakdownServiceFeeBreakdown instantiates a new FeeBreakdownServiceFeeBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeBreakdownServiceFeeBreakdown() *FeeBreakdownServiceFeeBreakdown {
	this := FeeBreakdownServiceFeeBreakdown{}
	return &this
}

// NewFeeBreakdownServiceFeeBreakdownWithDefaults instantiates a new FeeBreakdownServiceFeeBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeBreakdownServiceFeeBreakdownWithDefaults() *FeeBreakdownServiceFeeBreakdown {
	this := FeeBreakdownServiceFeeBreakdown{}
	return &this
}

// GetFlat returns the Flat field value if set, zero value otherwise.
func (o *FeeBreakdownServiceFeeBreakdown) GetFlat() FeeBreakdownServiceFeeBreakdownFlat {
	if o == nil || IsNil(o.Flat) {
		var ret FeeBreakdownServiceFeeBreakdownFlat
		return ret
	}
	return *o.Flat
}

// GetFlatOk returns a tuple with the Flat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeBreakdownServiceFeeBreakdown) GetFlatOk() (*FeeBreakdownServiceFeeBreakdownFlat, bool) {
	if o == nil || IsNil(o.Flat) {
		return nil, false
	}
	return o.Flat, true
}

// HasFlat returns a boolean if a field has been set.
func (o *FeeBreakdownServiceFeeBreakdown) HasFlat() bool {
	if o != nil && !IsNil(o.Flat) {
		return true
	}

	return false
}

// SetFlat gets a reference to the given FeeBreakdownServiceFeeBreakdownFlat and assigns it to the Flat field.
func (o *FeeBreakdownServiceFeeBreakdown) SetFlat(v FeeBreakdownServiceFeeBreakdownFlat) {
	o.Flat = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *FeeBreakdownServiceFeeBreakdown) GetPercentage() FeeBreakdownServiceFeeBreakdownPercentage {
	if o == nil || IsNil(o.Percentage) {
		var ret FeeBreakdownServiceFeeBreakdownPercentage
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeBreakdownServiceFeeBreakdown) GetPercentageOk() (*FeeBreakdownServiceFeeBreakdownPercentage, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *FeeBreakdownServiceFeeBreakdown) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given FeeBreakdownServiceFeeBreakdownPercentage and assigns it to the Percentage field.
func (o *FeeBreakdownServiceFeeBreakdown) SetPercentage(v FeeBreakdownServiceFeeBreakdownPercentage) {
	o.Percentage = &v
}

func (o FeeBreakdownServiceFeeBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeBreakdownServiceFeeBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flat) {
		toSerialize["flat"] = o.Flat
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableFeeBreakdownServiceFeeBreakdown struct {
	value *FeeBreakdownServiceFeeBreakdown
	isSet bool
}

func (v NullableFeeBreakdownServiceFeeBreakdown) Get() *FeeBreakdownServiceFeeBreakdown {
	return v.value
}

func (v *NullableFeeBreakdownServiceFeeBreakdown) Set(val *FeeBreakdownServiceFeeBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeBreakdownServiceFeeBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeBreakdownServiceFeeBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeBreakdownServiceFeeBreakdown(val *FeeBreakdownServiceFeeBreakdown) *NullableFeeBreakdownServiceFeeBreakdown {
	return &NullableFeeBreakdownServiceFeeBreakdown{value: val, isSet: true}
}

func (v NullableFeeBreakdownServiceFeeBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeBreakdownServiceFeeBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


