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

// checks if the AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner{}

// AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner struct for AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner
type AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner struct {
	// Category.
	Category *string `json:"category,omitempty"`
	Contributions []AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner `json:"contributions,omitempty"`
}

// NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner instantiates a new AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner {
	this := AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner{}
	return &this
}

// NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerWithDefaults instantiates a new AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerWithDefaults() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner {
	this := AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) SetCategory(v string) {
	o.Category = &v
}

// GetContributions returns the Contributions field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) GetContributions() []AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner {
	if o == nil || IsNil(o.Contributions) {
		var ret []AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner
		return ret
	}
	return o.Contributions
}

// GetContributionsOk returns a tuple with the Contributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) GetContributionsOk() ([]AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner, bool) {
	if o == nil || IsNil(o.Contributions) {
		return nil, false
	}
	return o.Contributions, true
}

// HasContributions returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) HasContributions() bool {
	if o != nil && !IsNil(o.Contributions) {
		return true
	}

	return false
}

// SetContributions gets a reference to the given []AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner and assigns it to the Contributions field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) SetContributions(v []AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) {
	o.Contributions = v
}

func (o AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Contributions) {
		toSerialize["contributions"] = o.Contributions
	}
	return toSerialize, nil
}

type NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner struct {
	value *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner
	isSet bool
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) Get() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner {
	return v.value
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) Set(val *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner(val *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner {
	return &NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner{value: val, isSet: true}
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

