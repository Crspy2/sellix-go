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

// checks if the AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue{}

// AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue struct for AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue
type AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue struct {
	// Contribution value in USD.
	Usd *float64 `json:"usd,omitempty"`
}

// NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue instantiates a new AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue {
	this := AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue{}
	return &this
}

// NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValueWithDefaults instantiates a new AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValueWithDefaults() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue {
	this := AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue{}
	return &this
}

// GetUsd returns the Usd field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) GetUsd() float64 {
	if o == nil || IsNil(o.Usd) {
		var ret float64
		return ret
	}
	return *o.Usd
}

// GetUsdOk returns a tuple with the Usd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) GetUsdOk() (*float64, bool) {
	if o == nil || IsNil(o.Usd) {
		return nil, false
	}
	return o.Usd, true
}

// HasUsd returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) HasUsd() bool {
	if o != nil && !IsNil(o.Usd) {
		return true
	}

	return false
}

// SetUsd gets a reference to the given float64 and assigns it to the Usd field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) SetUsd(v float64) {
	o.Usd = &v
}

func (o AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Usd) {
		toSerialize["usd"] = o.Usd
	}
	return toSerialize, nil
}

type NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue struct {
	value *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue
	isSet bool
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) Get() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue {
	return v.value
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) Set(val *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue(val *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue {
	return &NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue{value: val, isSet: true}
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


