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

// checks if the ListedProductVolumeDiscountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListedProductVolumeDiscountsInner{}

// ListedProductVolumeDiscountsInner struct for ListedProductVolumeDiscountsInner
type ListedProductVolumeDiscountsInner struct {
	// Whether the discount value is a percentage or a fixed amount.
	Type *string `json:"type,omitempty"`
	// Value of a percentage or fixed discount.
	Value *int32 `json:"value,omitempty"`
	// How much of this product needs to be purchased in order to be eligible for this volume discount.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewListedProductVolumeDiscountsInner instantiates a new ListedProductVolumeDiscountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListedProductVolumeDiscountsInner() *ListedProductVolumeDiscountsInner {
	this := ListedProductVolumeDiscountsInner{}
	return &this
}

// NewListedProductVolumeDiscountsInnerWithDefaults instantiates a new ListedProductVolumeDiscountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListedProductVolumeDiscountsInnerWithDefaults() *ListedProductVolumeDiscountsInner {
	this := ListedProductVolumeDiscountsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListedProductVolumeDiscountsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductVolumeDiscountsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListedProductVolumeDiscountsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListedProductVolumeDiscountsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListedProductVolumeDiscountsInner) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductVolumeDiscountsInner) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListedProductVolumeDiscountsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ListedProductVolumeDiscountsInner) SetValue(v int32) {
	o.Value = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ListedProductVolumeDiscountsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductVolumeDiscountsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ListedProductVolumeDiscountsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ListedProductVolumeDiscountsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o ListedProductVolumeDiscountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListedProductVolumeDiscountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableListedProductVolumeDiscountsInner struct {
	value *ListedProductVolumeDiscountsInner
	isSet bool
}

func (v NullableListedProductVolumeDiscountsInner) Get() *ListedProductVolumeDiscountsInner {
	return v.value
}

func (v *NullableListedProductVolumeDiscountsInner) Set(val *ListedProductVolumeDiscountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListedProductVolumeDiscountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListedProductVolumeDiscountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedProductVolumeDiscountsInner(val *ListedProductVolumeDiscountsInner) *NullableListedProductVolumeDiscountsInner {
	return &NullableListedProductVolumeDiscountsInner{value: val, isSet: true}
}

func (v NullableListedProductVolumeDiscountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedProductVolumeDiscountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

