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

// checks if the GatewayFees type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayFees{}

// GatewayFees struct for GatewayFees
type GatewayFees struct {
	// ID of the resource.
	Id *int32 `json:"id,omitempty"`
	// The shop ID to which this gateway fee belongs.
	ShopId *int32 `json:"shop_id,omitempty"`
	Gateway *Gateway `json:"gateway,omitempty"`
	// The percent amount of the gateway fee.
	PercentAmount *float32 `json:"percent_amount,omitempty"`
	// The fixed amount of the gateway fee.
	FixedAmount *float32 `json:"fixed_amount,omitempty"`
	FixedCurrency *Currency `json:"fixed_currency,omitempty"`
	// The active type of the gateway fee.
	ActiveType *string `json:"active_type,omitempty"`
	// DateTime, available if the gateway fee has been updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// DateTime, available if the gateway fee has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Conversions *GatewayFeesConversions `json:"conversions,omitempty"`
}

// NewGatewayFees instantiates a new GatewayFees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayFees() *GatewayFees {
	this := GatewayFees{}
	return &this
}

// NewGatewayFeesWithDefaults instantiates a new GatewayFees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayFeesWithDefaults() *GatewayFees {
	this := GatewayFees{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayFees) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayFees) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GatewayFees) SetId(v int32) {
	o.Id = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *GatewayFees) GetShopId() int32 {
	if o == nil || IsNil(o.ShopId) {
		var ret int32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetShopIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *GatewayFees) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int32 and assigns it to the ShopId field.
func (o *GatewayFees) SetShopId(v int32) {
	o.ShopId = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GatewayFees) GetGateway() Gateway {
	if o == nil || IsNil(o.Gateway) {
		var ret Gateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetGatewayOk() (*Gateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GatewayFees) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given Gateway and assigns it to the Gateway field.
func (o *GatewayFees) SetGateway(v Gateway) {
	o.Gateway = &v
}

// GetPercentAmount returns the PercentAmount field value if set, zero value otherwise.
func (o *GatewayFees) GetPercentAmount() float32 {
	if o == nil || IsNil(o.PercentAmount) {
		var ret float32
		return ret
	}
	return *o.PercentAmount
}

// GetPercentAmountOk returns a tuple with the PercentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetPercentAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentAmount) {
		return nil, false
	}
	return o.PercentAmount, true
}

// HasPercentAmount returns a boolean if a field has been set.
func (o *GatewayFees) HasPercentAmount() bool {
	if o != nil && !IsNil(o.PercentAmount) {
		return true
	}

	return false
}

// SetPercentAmount gets a reference to the given float32 and assigns it to the PercentAmount field.
func (o *GatewayFees) SetPercentAmount(v float32) {
	o.PercentAmount = &v
}

// GetFixedAmount returns the FixedAmount field value if set, zero value otherwise.
func (o *GatewayFees) GetFixedAmount() float32 {
	if o == nil || IsNil(o.FixedAmount) {
		var ret float32
		return ret
	}
	return *o.FixedAmount
}

// GetFixedAmountOk returns a tuple with the FixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetFixedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.FixedAmount) {
		return nil, false
	}
	return o.FixedAmount, true
}

// HasFixedAmount returns a boolean if a field has been set.
func (o *GatewayFees) HasFixedAmount() bool {
	if o != nil && !IsNil(o.FixedAmount) {
		return true
	}

	return false
}

// SetFixedAmount gets a reference to the given float32 and assigns it to the FixedAmount field.
func (o *GatewayFees) SetFixedAmount(v float32) {
	o.FixedAmount = &v
}

// GetFixedCurrency returns the FixedCurrency field value if set, zero value otherwise.
func (o *GatewayFees) GetFixedCurrency() Currency {
	if o == nil || IsNil(o.FixedCurrency) {
		var ret Currency
		return ret
	}
	return *o.FixedCurrency
}

// GetFixedCurrencyOk returns a tuple with the FixedCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetFixedCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.FixedCurrency) {
		return nil, false
	}
	return o.FixedCurrency, true
}

// HasFixedCurrency returns a boolean if a field has been set.
func (o *GatewayFees) HasFixedCurrency() bool {
	if o != nil && !IsNil(o.FixedCurrency) {
		return true
	}

	return false
}

// SetFixedCurrency gets a reference to the given Currency and assigns it to the FixedCurrency field.
func (o *GatewayFees) SetFixedCurrency(v Currency) {
	o.FixedCurrency = &v
}

// GetActiveType returns the ActiveType field value if set, zero value otherwise.
func (o *GatewayFees) GetActiveType() string {
	if o == nil || IsNil(o.ActiveType) {
		var ret string
		return ret
	}
	return *o.ActiveType
}

// GetActiveTypeOk returns a tuple with the ActiveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetActiveTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveType) {
		return nil, false
	}
	return o.ActiveType, true
}

// HasActiveType returns a boolean if a field has been set.
func (o *GatewayFees) HasActiveType() bool {
	if o != nil && !IsNil(o.ActiveType) {
		return true
	}

	return false
}

// SetActiveType gets a reference to the given string and assigns it to the ActiveType field.
func (o *GatewayFees) SetActiveType(v string) {
	o.ActiveType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GatewayFees) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GatewayFees) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GatewayFees) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GatewayFees) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GatewayFees) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GatewayFees) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetConversions returns the Conversions field value if set, zero value otherwise.
func (o *GatewayFees) GetConversions() GatewayFeesConversions {
	if o == nil || IsNil(o.Conversions) {
		var ret GatewayFeesConversions
		return ret
	}
	return *o.Conversions
}

// GetConversionsOk returns a tuple with the Conversions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFees) GetConversionsOk() (*GatewayFeesConversions, bool) {
	if o == nil || IsNil(o.Conversions) {
		return nil, false
	}
	return o.Conversions, true
}

// HasConversions returns a boolean if a field has been set.
func (o *GatewayFees) HasConversions() bool {
	if o != nil && !IsNil(o.Conversions) {
		return true
	}

	return false
}

// SetConversions gets a reference to the given GatewayFeesConversions and assigns it to the Conversions field.
func (o *GatewayFees) SetConversions(v GatewayFeesConversions) {
	o.Conversions = &v
}

func (o GatewayFees) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayFees) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Conversions) {
		toSerialize["conversions"] = o.Conversions
	}
	return toSerialize, nil
}

type NullableGatewayFees struct {
	value *GatewayFees
	isSet bool
}

func (v NullableGatewayFees) Get() *GatewayFees {
	return v.value
}

func (v *NullableGatewayFees) Set(val *GatewayFees) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayFees) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayFees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayFees(val *GatewayFees) *NullableGatewayFees {
	return &NullableGatewayFees{value: val, isSet: true}
}

func (v NullableGatewayFees) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayFees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

