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

// checks if the ProductDownload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDownload{}

// ProductDownload struct for ProductDownload
type ProductDownload struct {
	// ID of the resource.
	Id *int32 `json:"id,omitempty"`
	// Unique ID of the invoice this download is linked to.
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The IP of the customer that downloaded the product.
	CustomerIp *string `json:"customer_ip,omitempty"`
	// The ISP of the customer that downloaded the product.
	CustomerIsp *string `json:"customer_isp,omitempty"`
	// The timezone of the customer that downloaded the product.
	CustomerTimezone *string `json:"customer_timezone,omitempty"`
	// The country of the customer that downloaded the product.
	CustomerCountry *string `json:"customer_country,omitempty"`
	// Timestamp, available if the download has been created.
	CreatedAt *int32 `json:"created_at,omitempty"`
}

// NewProductDownload instantiates a new ProductDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDownload() *ProductDownload {
	this := ProductDownload{}
	return &this
}

// NewProductDownloadWithDefaults instantiates a new ProductDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDownloadWithDefaults() *ProductDownload {
	this := ProductDownload{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductDownload) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductDownload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProductDownload) SetId(v int32) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *ProductDownload) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *ProductDownload) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *ProductDownload) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *ProductDownload) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *ProductDownload) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *ProductDownload) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetCustomerIsp returns the CustomerIsp field value if set, zero value otherwise.
func (o *ProductDownload) GetCustomerIsp() string {
	if o == nil || IsNil(o.CustomerIsp) {
		var ret string
		return ret
	}
	return *o.CustomerIsp
}

// GetCustomerIspOk returns a tuple with the CustomerIsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetCustomerIspOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIsp) {
		return nil, false
	}
	return o.CustomerIsp, true
}

// HasCustomerIsp returns a boolean if a field has been set.
func (o *ProductDownload) HasCustomerIsp() bool {
	if o != nil && !IsNil(o.CustomerIsp) {
		return true
	}

	return false
}

// SetCustomerIsp gets a reference to the given string and assigns it to the CustomerIsp field.
func (o *ProductDownload) SetCustomerIsp(v string) {
	o.CustomerIsp = &v
}

// GetCustomerTimezone returns the CustomerTimezone field value if set, zero value otherwise.
func (o *ProductDownload) GetCustomerTimezone() string {
	if o == nil || IsNil(o.CustomerTimezone) {
		var ret string
		return ret
	}
	return *o.CustomerTimezone
}

// GetCustomerTimezoneOk returns a tuple with the CustomerTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetCustomerTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerTimezone) {
		return nil, false
	}
	return o.CustomerTimezone, true
}

// HasCustomerTimezone returns a boolean if a field has been set.
func (o *ProductDownload) HasCustomerTimezone() bool {
	if o != nil && !IsNil(o.CustomerTimezone) {
		return true
	}

	return false
}

// SetCustomerTimezone gets a reference to the given string and assigns it to the CustomerTimezone field.
func (o *ProductDownload) SetCustomerTimezone(v string) {
	o.CustomerTimezone = &v
}

// GetCustomerCountry returns the CustomerCountry field value if set, zero value otherwise.
func (o *ProductDownload) GetCustomerCountry() string {
	if o == nil || IsNil(o.CustomerCountry) {
		var ret string
		return ret
	}
	return *o.CustomerCountry
}

// GetCustomerCountryOk returns a tuple with the CustomerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetCustomerCountryOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerCountry) {
		return nil, false
	}
	return o.CustomerCountry, true
}

// HasCustomerCountry returns a boolean if a field has been set.
func (o *ProductDownload) HasCustomerCountry() bool {
	if o != nil && !IsNil(o.CustomerCountry) {
		return true
	}

	return false
}

// SetCustomerCountry gets a reference to the given string and assigns it to the CustomerCountry field.
func (o *ProductDownload) SetCustomerCountry(v string) {
	o.CustomerCountry = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductDownload) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDownload) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductDownload) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *ProductDownload) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

func (o ProductDownload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDownload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.CustomerIp) {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if !IsNil(o.CustomerIsp) {
		toSerialize["customer_isp"] = o.CustomerIsp
	}
	if !IsNil(o.CustomerTimezone) {
		toSerialize["customer_timezone"] = o.CustomerTimezone
	}
	if !IsNil(o.CustomerCountry) {
		toSerialize["customer_country"] = o.CustomerCountry
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableProductDownload struct {
	value *ProductDownload
	isSet bool
}

func (v NullableProductDownload) Get() *ProductDownload {
	return v.value
}

func (v *NullableProductDownload) Set(val *ProductDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDownload(val *ProductDownload) *NullableProductDownload {
	return &NullableProductDownload{value: val, isSet: true}
}

func (v NullableProductDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

