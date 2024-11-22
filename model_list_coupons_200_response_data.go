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

// checks if the ListCoupons200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCoupons200ResponseData{}

// ListCoupons200ResponseData struct for ListCoupons200ResponseData
type ListCoupons200ResponseData struct {
	Coupons []CouponListing `json:"coupons,omitempty"`
}

// NewListCoupons200ResponseData instantiates a new ListCoupons200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCoupons200ResponseData() *ListCoupons200ResponseData {
	this := ListCoupons200ResponseData{}
	return &this
}

// NewListCoupons200ResponseDataWithDefaults instantiates a new ListCoupons200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCoupons200ResponseDataWithDefaults() *ListCoupons200ResponseData {
	this := ListCoupons200ResponseData{}
	return &this
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *ListCoupons200ResponseData) GetCoupons() []CouponListing {
	if o == nil || IsNil(o.Coupons) {
		var ret []CouponListing
		return ret
	}
	return o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCoupons200ResponseData) GetCouponsOk() ([]CouponListing, bool) {
	if o == nil || IsNil(o.Coupons) {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *ListCoupons200ResponseData) HasCoupons() bool {
	if o != nil && !IsNil(o.Coupons) {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given []CouponListing and assigns it to the Coupons field.
func (o *ListCoupons200ResponseData) SetCoupons(v []CouponListing) {
	o.Coupons = v
}

func (o ListCoupons200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCoupons200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coupons) {
		toSerialize["coupons"] = o.Coupons
	}
	return toSerialize, nil
}

type NullableListCoupons200ResponseData struct {
	value *ListCoupons200ResponseData
	isSet bool
}

func (v NullableListCoupons200ResponseData) Get() *ListCoupons200ResponseData {
	return v.value
}

func (v *NullableListCoupons200ResponseData) Set(val *ListCoupons200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListCoupons200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListCoupons200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCoupons200ResponseData(val *ListCoupons200ResponseData) *NullableListCoupons200ResponseData {
	return &NullableListCoupons200ResponseData{value: val, isSet: true}
}

func (v NullableListCoupons200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCoupons200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


