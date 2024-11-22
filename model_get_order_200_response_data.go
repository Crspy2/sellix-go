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

// checks if the GetOrder200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrder200ResponseData{}

// GetOrder200ResponseData struct for GetOrder200ResponseData
type GetOrder200ResponseData struct {
	Order *Invoice `json:"order,omitempty"`
}

// NewGetOrder200ResponseData instantiates a new GetOrder200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrder200ResponseData() *GetOrder200ResponseData {
	this := GetOrder200ResponseData{}
	return &this
}

// NewGetOrder200ResponseDataWithDefaults instantiates a new GetOrder200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrder200ResponseDataWithDefaults() *GetOrder200ResponseData {
	this := GetOrder200ResponseData{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GetOrder200ResponseData) GetOrder() Invoice {
	if o == nil || IsNil(o.Order) {
		var ret Invoice
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrder200ResponseData) GetOrderOk() (*Invoice, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GetOrder200ResponseData) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Invoice and assigns it to the Order field.
func (o *GetOrder200ResponseData) SetOrder(v Invoice) {
	o.Order = &v
}

func (o GetOrder200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrder200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableGetOrder200ResponseData struct {
	value *GetOrder200ResponseData
	isSet bool
}

func (v NullableGetOrder200ResponseData) Get() *GetOrder200ResponseData {
	return v.value
}

func (v *NullableGetOrder200ResponseData) Set(val *GetOrder200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrder200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrder200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrder200ResponseData(val *GetOrder200ResponseData) *NullableGetOrder200ResponseData {
	return &NullableGetOrder200ResponseData{value: val, isSet: true}
}

func (v NullableGetOrder200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrder200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


