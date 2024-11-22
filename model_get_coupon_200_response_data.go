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

// checks if the GetCoupon200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCoupon200ResponseData{}

// GetCoupon200ResponseData struct for GetCoupon200ResponseData
type GetCoupon200ResponseData struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

// NewGetCoupon200ResponseData instantiates a new GetCoupon200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCoupon200ResponseData() *GetCoupon200ResponseData {
	this := GetCoupon200ResponseData{}
	return &this
}

// NewGetCoupon200ResponseDataWithDefaults instantiates a new GetCoupon200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCoupon200ResponseDataWithDefaults() *GetCoupon200ResponseData {
	this := GetCoupon200ResponseData{}
	return &this
}

// GetCoupon returns the Coupon field value if set, zero value otherwise.
func (o *GetCoupon200ResponseData) GetCoupon() Coupon {
	if o == nil || IsNil(o.Coupon) {
		var ret Coupon
		return ret
	}
	return *o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCoupon200ResponseData) GetCouponOk() (*Coupon, bool) {
	if o == nil || IsNil(o.Coupon) {
		return nil, false
	}
	return o.Coupon, true
}

// HasCoupon returns a boolean if a field has been set.
func (o *GetCoupon200ResponseData) HasCoupon() bool {
	if o != nil && !IsNil(o.Coupon) {
		return true
	}

	return false
}

// SetCoupon gets a reference to the given Coupon and assigns it to the Coupon field.
func (o *GetCoupon200ResponseData) SetCoupon(v Coupon) {
	o.Coupon = &v
}

func (o GetCoupon200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCoupon200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coupon) {
		toSerialize["coupon"] = o.Coupon
	}
	return toSerialize, nil
}

type NullableGetCoupon200ResponseData struct {
	value *GetCoupon200ResponseData
	isSet bool
}

func (v NullableGetCoupon200ResponseData) Get() *GetCoupon200ResponseData {
	return v.value
}

func (v *NullableGetCoupon200ResponseData) Set(val *GetCoupon200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCoupon200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCoupon200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCoupon200ResponseData(val *GetCoupon200ResponseData) *NullableGetCoupon200ResponseData {
	return &NullableGetCoupon200ResponseData{value: val, isSet: true}
}

func (v NullableGetCoupon200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCoupon200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


