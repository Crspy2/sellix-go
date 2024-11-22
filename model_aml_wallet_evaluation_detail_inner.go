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

// checks if the AmlWalletEvaluationDetailInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmlWalletEvaluationDetailInner{}

// AmlWalletEvaluationDetailInner struct for AmlWalletEvaluationDetailInner
type AmlWalletEvaluationDetailInner struct {
	Source []AmlWalletEvaluationDetailInnerSourceInner `json:"source,omitempty"`
}

// NewAmlWalletEvaluationDetailInner instantiates a new AmlWalletEvaluationDetailInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmlWalletEvaluationDetailInner() *AmlWalletEvaluationDetailInner {
	this := AmlWalletEvaluationDetailInner{}
	return &this
}

// NewAmlWalletEvaluationDetailInnerWithDefaults instantiates a new AmlWalletEvaluationDetailInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmlWalletEvaluationDetailInnerWithDefaults() *AmlWalletEvaluationDetailInner {
	this := AmlWalletEvaluationDetailInner{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInner) GetSource() []AmlWalletEvaluationDetailInnerSourceInner {
	if o == nil || IsNil(o.Source) {
		var ret []AmlWalletEvaluationDetailInnerSourceInner
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInner) GetSourceOk() ([]AmlWalletEvaluationDetailInnerSourceInner, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInner) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given []AmlWalletEvaluationDetailInnerSourceInner and assigns it to the Source field.
func (o *AmlWalletEvaluationDetailInner) SetSource(v []AmlWalletEvaluationDetailInnerSourceInner) {
	o.Source = v
}

func (o AmlWalletEvaluationDetailInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmlWalletEvaluationDetailInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableAmlWalletEvaluationDetailInner struct {
	value *AmlWalletEvaluationDetailInner
	isSet bool
}

func (v NullableAmlWalletEvaluationDetailInner) Get() *AmlWalletEvaluationDetailInner {
	return v.value
}

func (v *NullableAmlWalletEvaluationDetailInner) Set(val *AmlWalletEvaluationDetailInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAmlWalletEvaluationDetailInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAmlWalletEvaluationDetailInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmlWalletEvaluationDetailInner(val *AmlWalletEvaluationDetailInner) *NullableAmlWalletEvaluationDetailInner {
	return &NullableAmlWalletEvaluationDetailInner{value: val, isSet: true}
}

func (v NullableAmlWalletEvaluationDetailInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmlWalletEvaluationDetailInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


