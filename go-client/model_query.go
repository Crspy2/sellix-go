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

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query struct for Query
type Query struct {
	// Unique ID of the resource, used as reference across the API.
	Uniqid *string `json:"uniqid,omitempty"`
	// The shop ID to which this query belongs.
	ShopId *int32 `json:"shop_id,omitempty"`
	// Unique ID of the invoice this query is linked to, if specified by the customer.
	InvoiceId *string `json:"invoice_id,omitempty"`
	// Email of the customer who created this query.
	CustomerEmail *string `json:"customer_email,omitempty"`
	// Query title, brief summary of what the customer needs.
	Title *string `json:"title,omitempty"`
	// Status of the query. PENDING if the query has been created and awaits a reply from the merchant. CLOSED if the query has been closed by the merchant or the customer. SHOP_REPLY if the query has been replied by the merchant, CUSTOMER_REPLY if the query has been replied by the customer.
	Status *string `json:"status,omitempty"`
	// Day value, number.
	DayValue *int32 `json:"day_value,omitempty"`
	// First three letters of the month name.
	Month *string `json:"month,omitempty"`
	// Year number.
	Year *int32 `json:"year,omitempty"`
	// Timestamp for the creation of the query.
	CreatedAt *int32 `json:"created_at,omitempty"`
	// Date, available if the query has been updated.
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// User ID of the user who updated the query.
	UpdatedBy *int32 `json:"updated_by,omitempty"`
	Messages []QueryMessage `json:"messages,omitempty"`
}

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery() *Query {
	this := Query{}
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetUniqid returns the Uniqid field value if set, zero value otherwise.
func (o *Query) GetUniqid() string {
	if o == nil || IsNil(o.Uniqid) {
		var ret string
		return ret
	}
	return *o.Uniqid
}

// GetUniqidOk returns a tuple with the Uniqid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetUniqidOk() (*string, bool) {
	if o == nil || IsNil(o.Uniqid) {
		return nil, false
	}
	return o.Uniqid, true
}

// HasUniqid returns a boolean if a field has been set.
func (o *Query) HasUniqid() bool {
	if o != nil && !IsNil(o.Uniqid) {
		return true
	}

	return false
}

// SetUniqid gets a reference to the given string and assigns it to the Uniqid field.
func (o *Query) SetUniqid(v string) {
	o.Uniqid = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *Query) GetShopId() int32 {
	if o == nil || IsNil(o.ShopId) {
		var ret int32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetShopIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *Query) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int32 and assigns it to the ShopId field.
func (o *Query) SetShopId(v int32) {
	o.ShopId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Query) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Query) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Query) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *Query) GetCustomerEmail() string {
	if o == nil || IsNil(o.CustomerEmail) {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetCustomerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerEmail) {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *Query) HasCustomerEmail() bool {
	if o != nil && !IsNil(o.CustomerEmail) {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *Query) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Query) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Query) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Query) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Query) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Query) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Query) SetStatus(v string) {
	o.Status = &v
}

// GetDayValue returns the DayValue field value if set, zero value otherwise.
func (o *Query) GetDayValue() int32 {
	if o == nil || IsNil(o.DayValue) {
		var ret int32
		return ret
	}
	return *o.DayValue
}

// GetDayValueOk returns a tuple with the DayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetDayValueOk() (*int32, bool) {
	if o == nil || IsNil(o.DayValue) {
		return nil, false
	}
	return o.DayValue, true
}

// HasDayValue returns a boolean if a field has been set.
func (o *Query) HasDayValue() bool {
	if o != nil && !IsNil(o.DayValue) {
		return true
	}

	return false
}

// SetDayValue gets a reference to the given int32 and assigns it to the DayValue field.
func (o *Query) SetDayValue(v int32) {
	o.DayValue = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *Query) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *Query) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *Query) SetMonth(v string) {
	o.Month = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *Query) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *Query) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *Query) SetYear(v int32) {
	o.Year = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Query) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Query) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *Query) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Query) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Query) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *Query) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Query) GetUpdatedBy() int32 {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetUpdatedByOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Query) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *Query) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Query) GetMessages() []QueryMessage {
	if o == nil || IsNil(o.Messages) {
		var ret []QueryMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetMessagesOk() ([]QueryMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Query) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []QueryMessage and assigns it to the Messages field.
func (o *Query) SetMessages(v []QueryMessage) {
	o.Messages = v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uniqid) {
		toSerialize["uniqid"] = o.Uniqid
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.DayValue) {
		toSerialize["day_value"] = o.DayValue
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

