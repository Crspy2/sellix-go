/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner{}

// AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner struct for AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner
type AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner struct {
	// Entity.
	Entity *string `json:"entity,omitempty"`
	RiskTriggers *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerRiskTriggers `json:"risk_triggers,omitempty"`
	ContributionValue *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue `json:"contribution_value,omitempty"`
	// Contribution percentage.
	ContributionPercentage *float64 `json:"contribution_percentage,omitempty"`
}

// NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner instantiates a new AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner {
	this := AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner{}
	return &this
}

// NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerWithDefaults instantiates a new AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerWithDefaults() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner {
	this := AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) SetEntity(v string) {
	o.Entity = &v
}

// GetRiskTriggers returns the RiskTriggers field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetRiskTriggers() AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerRiskTriggers {
	if o == nil || IsNil(o.RiskTriggers) {
		var ret AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerRiskTriggers
		return ret
	}
	return *o.RiskTriggers
}

// GetRiskTriggersOk returns a tuple with the RiskTriggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetRiskTriggersOk() (*AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerRiskTriggers, bool) {
	if o == nil || IsNil(o.RiskTriggers) {
		return nil, false
	}
	return o.RiskTriggers, true
}

// HasRiskTriggers returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) HasRiskTriggers() bool {
	if o != nil && !IsNil(o.RiskTriggers) {
		return true
	}

	return false
}

// SetRiskTriggers gets a reference to the given AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerRiskTriggers and assigns it to the RiskTriggers field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) SetRiskTriggers(v AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerRiskTriggers) {
	o.RiskTriggers = &v
}

// GetContributionValue returns the ContributionValue field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetContributionValue() AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue {
	if o == nil || IsNil(o.ContributionValue) {
		var ret AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue
		return ret
	}
	return *o.ContributionValue
}

// GetContributionValueOk returns a tuple with the ContributionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetContributionValueOk() (*AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue, bool) {
	if o == nil || IsNil(o.ContributionValue) {
		return nil, false
	}
	return o.ContributionValue, true
}

// HasContributionValue returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) HasContributionValue() bool {
	if o != nil && !IsNil(o.ContributionValue) {
		return true
	}

	return false
}

// SetContributionValue gets a reference to the given AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue and assigns it to the ContributionValue field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) SetContributionValue(v AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInnerContributionValue) {
	o.ContributionValue = &v
}

// GetContributionPercentage returns the ContributionPercentage field value if set, zero value otherwise.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetContributionPercentage() float64 {
	if o == nil || IsNil(o.ContributionPercentage) {
		var ret float64
		return ret
	}
	return *o.ContributionPercentage
}

// GetContributionPercentageOk returns a tuple with the ContributionPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) GetContributionPercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.ContributionPercentage) {
		return nil, false
	}
	return o.ContributionPercentage, true
}

// HasContributionPercentage returns a boolean if a field has been set.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) HasContributionPercentage() bool {
	if o != nil && !IsNil(o.ContributionPercentage) {
		return true
	}

	return false
}

// SetContributionPercentage gets a reference to the given float64 and assigns it to the ContributionPercentage field.
func (o *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) SetContributionPercentage(v float64) {
	o.ContributionPercentage = &v
}

func (o AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.RiskTriggers) {
		toSerialize["risk_triggers"] = o.RiskTriggers
	}
	if !IsNil(o.ContributionValue) {
		toSerialize["contribution_value"] = o.ContributionValue
	}
	if !IsNil(o.ContributionPercentage) {
		toSerialize["contribution_percentage"] = o.ContributionPercentage
	}
	return toSerialize, nil
}

type NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner struct {
	value *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner
	isSet bool
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) Get() *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner {
	return v.value
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) Set(val *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner(val *AmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner {
	return &NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner{value: val, isSet: true}
}

func (v NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmlWalletEvaluationDetailInnerSourceInnerMatchedElementsInnerContributionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

