/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sellix

import (
	"encoding/json"
	"time"
)

// checks if the ListedProductPaymentGatewaysFeesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListedProductPaymentGatewaysFeesInner{}

// ListedProductPaymentGatewaysFeesInner struct for ListedProductPaymentGatewaysFeesInner
type ListedProductPaymentGatewaysFeesInner struct {
	Id *int32 `json:"id,omitempty"`
	// The shop ID to which this resource belongs.
	ShopId *int32 `json:"shop_id,omitempty"`
	Gateway *Gateway `json:"gateway,omitempty"`
	PercentAmount *int32 `json:"percent_amount,omitempty"`
	FixedAmount *int32 `json:"fixed_amount,omitempty"`
	FixedCurrency *Currency `json:"fixed_currency,omitempty"`
	// DEPRECATED
	ActiveType *string `json:"active_type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Datetime for the creation of the resource
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewListedProductPaymentGatewaysFeesInner instantiates a new ListedProductPaymentGatewaysFeesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListedProductPaymentGatewaysFeesInner() *ListedProductPaymentGatewaysFeesInner {
	this := ListedProductPaymentGatewaysFeesInner{}
	return &this
}

// NewListedProductPaymentGatewaysFeesInnerWithDefaults instantiates a new ListedProductPaymentGatewaysFeesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListedProductPaymentGatewaysFeesInnerWithDefaults() *ListedProductPaymentGatewaysFeesInner {
	this := ListedProductPaymentGatewaysFeesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListedProductPaymentGatewaysFeesInner) SetId(v int32) {
	o.Id = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetShopId() int32 {
	if o == nil || IsNil(o.ShopId) {
		var ret int32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetShopIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int32 and assigns it to the ShopId field.
func (o *ListedProductPaymentGatewaysFeesInner) SetShopId(v int32) {
	o.ShopId = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetGateway() Gateway {
	if o == nil || IsNil(o.Gateway) {
		var ret Gateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetGatewayOk() (*Gateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given Gateway and assigns it to the Gateway field.
func (o *ListedProductPaymentGatewaysFeesInner) SetGateway(v Gateway) {
	o.Gateway = &v
}

// GetPercentAmount returns the PercentAmount field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetPercentAmount() int32 {
	if o == nil || IsNil(o.PercentAmount) {
		var ret int32
		return ret
	}
	return *o.PercentAmount
}

// GetPercentAmountOk returns a tuple with the PercentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetPercentAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentAmount) {
		return nil, false
	}
	return o.PercentAmount, true
}

// HasPercentAmount returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasPercentAmount() bool {
	if o != nil && !IsNil(o.PercentAmount) {
		return true
	}

	return false
}

// SetPercentAmount gets a reference to the given int32 and assigns it to the PercentAmount field.
func (o *ListedProductPaymentGatewaysFeesInner) SetPercentAmount(v int32) {
	o.PercentAmount = &v
}

// GetFixedAmount returns the FixedAmount field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetFixedAmount() int32 {
	if o == nil || IsNil(o.FixedAmount) {
		var ret int32
		return ret
	}
	return *o.FixedAmount
}

// GetFixedAmountOk returns a tuple with the FixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetFixedAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.FixedAmount) {
		return nil, false
	}
	return o.FixedAmount, true
}

// HasFixedAmount returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasFixedAmount() bool {
	if o != nil && !IsNil(o.FixedAmount) {
		return true
	}

	return false
}

// SetFixedAmount gets a reference to the given int32 and assigns it to the FixedAmount field.
func (o *ListedProductPaymentGatewaysFeesInner) SetFixedAmount(v int32) {
	o.FixedAmount = &v
}

// GetFixedCurrency returns the FixedCurrency field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetFixedCurrency() Currency {
	if o == nil || IsNil(o.FixedCurrency) {
		var ret Currency
		return ret
	}
	return *o.FixedCurrency
}

// GetFixedCurrencyOk returns a tuple with the FixedCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetFixedCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.FixedCurrency) {
		return nil, false
	}
	return o.FixedCurrency, true
}

// HasFixedCurrency returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasFixedCurrency() bool {
	if o != nil && !IsNil(o.FixedCurrency) {
		return true
	}

	return false
}

// SetFixedCurrency gets a reference to the given Currency and assigns it to the FixedCurrency field.
func (o *ListedProductPaymentGatewaysFeesInner) SetFixedCurrency(v Currency) {
	o.FixedCurrency = &v
}

// GetActiveType returns the ActiveType field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetActiveType() string {
	if o == nil || IsNil(o.ActiveType) {
		var ret string
		return ret
	}
	return *o.ActiveType
}

// GetActiveTypeOk returns a tuple with the ActiveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetActiveTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveType) {
		return nil, false
	}
	return o.ActiveType, true
}

// HasActiveType returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasActiveType() bool {
	if o != nil && !IsNil(o.ActiveType) {
		return true
	}

	return false
}

// SetActiveType gets a reference to the given string and assigns it to the ActiveType field.
func (o *ListedProductPaymentGatewaysFeesInner) SetActiveType(v string) {
	o.ActiveType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ListedProductPaymentGatewaysFeesInner) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ListedProductPaymentGatewaysFeesInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedProductPaymentGatewaysFeesInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ListedProductPaymentGatewaysFeesInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ListedProductPaymentGatewaysFeesInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o ListedProductPaymentGatewaysFeesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListedProductPaymentGatewaysFeesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.PercentAmount) {
		toSerialize["percent_amount"] = o.PercentAmount
	}
	if !IsNil(o.FixedAmount) {
		toSerialize["fixed_amount"] = o.FixedAmount
	}
	if !IsNil(o.FixedCurrency) {
		toSerialize["fixed_currency"] = o.FixedCurrency
	}
	if !IsNil(o.ActiveType) {
		toSerialize["active_type"] = o.ActiveType
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableListedProductPaymentGatewaysFeesInner struct {
	value *ListedProductPaymentGatewaysFeesInner
	isSet bool
}

func (v NullableListedProductPaymentGatewaysFeesInner) Get() *ListedProductPaymentGatewaysFeesInner {
	return v.value
}

func (v *NullableListedProductPaymentGatewaysFeesInner) Set(val *ListedProductPaymentGatewaysFeesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListedProductPaymentGatewaysFeesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListedProductPaymentGatewaysFeesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedProductPaymentGatewaysFeesInner(val *ListedProductPaymentGatewaysFeesInner) *NullableListedProductPaymentGatewaysFeesInner {
	return &NullableListedProductPaymentGatewaysFeesInner{value: val, isSet: true}
}

func (v NullableListedProductPaymentGatewaysFeesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedProductPaymentGatewaysFeesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


