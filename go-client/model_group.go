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

// checks if the Group type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Group{}

// Group struct for Group
type Group struct {
	// The ID of the resource
	Id *float32 `json:"id,omitempty"`
	// Unique ID of the resource, used as reference across the API.
	Uniqid *string `json:"uniqid,omitempty"`
	// The unique identifyer for the shop.
	ShopId *float32 `json:"shop_id,omitempty"`
	// Group title.
	Title *string `json:"title,omitempty"`
	// DEPRECATED: Unique id of the image attachment for this product.
	ImageAttachment *string `json:"image_attachment,omitempty"`
	// Whether or not the group is unlisted.
	Unlisted *bool `json:"unlisted,omitempty"`
	// The priority of the group on the storefront.
	SortPriority *float32 `json:"sort_priority,omitempty"`
	// Timestamp for the creation of the group.
	CreatedAt *int32 `json:"created_at,omitempty"`
	// Date, available if the group has been edited.
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// User ID of the user who updated the group.
	UpdatedBy *int32 `json:"updated_by,omitempty"`
	ProductsBound []Product `json:"products_bound,omitempty"`
	ImageAttachmentInfo *ImageAttachment `json:"image_attachment_info,omitempty"`
	// The name of the store the group is on.
	Name *string `json:"name,omitempty"`
	// The theme of the group.
	Theme *string `json:"theme,omitempty"`
	// Whether or not the group is in dark mode.
	DarkMode *int32 `json:"dark_mode,omitempty"`
	// Whether or not the group has a custom UI style configuration.
	UiStyleConfiguration *bool `json:"ui_style_configuration,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Group) SetId(v float32) {
	o.Id = &v
}

// GetUniqid returns the Uniqid field value if set, zero value otherwise.
func (o *Group) GetUniqid() string {
	if o == nil || IsNil(o.Uniqid) {
		var ret string
		return ret
	}
	return *o.Uniqid
}

// GetUniqidOk returns a tuple with the Uniqid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUniqidOk() (*string, bool) {
	if o == nil || IsNil(o.Uniqid) {
		return nil, false
	}
	return o.Uniqid, true
}

// HasUniqid returns a boolean if a field has been set.
func (o *Group) HasUniqid() bool {
	if o != nil && !IsNil(o.Uniqid) {
		return true
	}

	return false
}

// SetUniqid gets a reference to the given string and assigns it to the Uniqid field.
func (o *Group) SetUniqid(v string) {
	o.Uniqid = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *Group) GetShopId() float32 {
	if o == nil || IsNil(o.ShopId) {
		var ret float32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetShopIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *Group) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given float32 and assigns it to the ShopId field.
func (o *Group) SetShopId(v float32) {
	o.ShopId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Group) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Group) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Group) SetTitle(v string) {
	o.Title = &v
}

// GetImageAttachment returns the ImageAttachment field value if set, zero value otherwise.
func (o *Group) GetImageAttachment() string {
	if o == nil || IsNil(o.ImageAttachment) {
		var ret string
		return ret
	}
	return *o.ImageAttachment
}

// GetImageAttachmentOk returns a tuple with the ImageAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetImageAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.ImageAttachment) {
		return nil, false
	}
	return o.ImageAttachment, true
}

// HasImageAttachment returns a boolean if a field has been set.
func (o *Group) HasImageAttachment() bool {
	if o != nil && !IsNil(o.ImageAttachment) {
		return true
	}

	return false
}

// SetImageAttachment gets a reference to the given string and assigns it to the ImageAttachment field.
func (o *Group) SetImageAttachment(v string) {
	o.ImageAttachment = &v
}

// GetUnlisted returns the Unlisted field value if set, zero value otherwise.
func (o *Group) GetUnlisted() bool {
	if o == nil || IsNil(o.Unlisted) {
		var ret bool
		return ret
	}
	return *o.Unlisted
}

// GetUnlistedOk returns a tuple with the Unlisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUnlistedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unlisted) {
		return nil, false
	}
	return o.Unlisted, true
}

// HasUnlisted returns a boolean if a field has been set.
func (o *Group) HasUnlisted() bool {
	if o != nil && !IsNil(o.Unlisted) {
		return true
	}

	return false
}

// SetUnlisted gets a reference to the given bool and assigns it to the Unlisted field.
func (o *Group) SetUnlisted(v bool) {
	o.Unlisted = &v
}

// GetSortPriority returns the SortPriority field value if set, zero value otherwise.
func (o *Group) GetSortPriority() float32 {
	if o == nil || IsNil(o.SortPriority) {
		var ret float32
		return ret
	}
	return *o.SortPriority
}

// GetSortPriorityOk returns a tuple with the SortPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSortPriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.SortPriority) {
		return nil, false
	}
	return o.SortPriority, true
}

// HasSortPriority returns a boolean if a field has been set.
func (o *Group) HasSortPriority() bool {
	if o != nil && !IsNil(o.SortPriority) {
		return true
	}

	return false
}

// SetSortPriority gets a reference to the given float32 and assigns it to the SortPriority field.
func (o *Group) SetSortPriority(v float32) {
	o.SortPriority = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Group) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Group) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *Group) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Group) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Group) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *Group) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Group) GetUpdatedBy() int32 {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUpdatedByOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Group) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *Group) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

// GetProductsBound returns the ProductsBound field value if set, zero value otherwise.
func (o *Group) GetProductsBound() []Product {
	if o == nil || IsNil(o.ProductsBound) {
		var ret []Product
		return ret
	}
	return o.ProductsBound
}

// GetProductsBoundOk returns a tuple with the ProductsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetProductsBoundOk() ([]Product, bool) {
	if o == nil || IsNil(o.ProductsBound) {
		return nil, false
	}
	return o.ProductsBound, true
}

// HasProductsBound returns a boolean if a field has been set.
func (o *Group) HasProductsBound() bool {
	if o != nil && !IsNil(o.ProductsBound) {
		return true
	}

	return false
}

// SetProductsBound gets a reference to the given []Product and assigns it to the ProductsBound field.
func (o *Group) SetProductsBound(v []Product) {
	o.ProductsBound = v
}

// GetImageAttachmentInfo returns the ImageAttachmentInfo field value if set, zero value otherwise.
func (o *Group) GetImageAttachmentInfo() ImageAttachment {
	if o == nil || IsNil(o.ImageAttachmentInfo) {
		var ret ImageAttachment
		return ret
	}
	return *o.ImageAttachmentInfo
}

// GetImageAttachmentInfoOk returns a tuple with the ImageAttachmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetImageAttachmentInfoOk() (*ImageAttachment, bool) {
	if o == nil || IsNil(o.ImageAttachmentInfo) {
		return nil, false
	}
	return o.ImageAttachmentInfo, true
}

// HasImageAttachmentInfo returns a boolean if a field has been set.
func (o *Group) HasImageAttachmentInfo() bool {
	if o != nil && !IsNil(o.ImageAttachmentInfo) {
		return true
	}

	return false
}

// SetImageAttachmentInfo gets a reference to the given ImageAttachment and assigns it to the ImageAttachmentInfo field.
func (o *Group) SetImageAttachmentInfo(v ImageAttachment) {
	o.ImageAttachmentInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Group) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Group) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Group) SetName(v string) {
	o.Name = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *Group) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *Group) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *Group) SetTheme(v string) {
	o.Theme = &v
}

// GetDarkMode returns the DarkMode field value if set, zero value otherwise.
func (o *Group) GetDarkMode() int32 {
	if o == nil || IsNil(o.DarkMode) {
		var ret int32
		return ret
	}
	return *o.DarkMode
}

// GetDarkModeOk returns a tuple with the DarkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDarkModeOk() (*int32, bool) {
	if o == nil || IsNil(o.DarkMode) {
		return nil, false
	}
	return o.DarkMode, true
}

// HasDarkMode returns a boolean if a field has been set.
func (o *Group) HasDarkMode() bool {
	if o != nil && !IsNil(o.DarkMode) {
		return true
	}

	return false
}

// SetDarkMode gets a reference to the given int32 and assigns it to the DarkMode field.
func (o *Group) SetDarkMode(v int32) {
	o.DarkMode = &v
}

// GetUiStyleConfiguration returns the UiStyleConfiguration field value if set, zero value otherwise.
func (o *Group) GetUiStyleConfiguration() bool {
	if o == nil || IsNil(o.UiStyleConfiguration) {
		var ret bool
		return ret
	}
	return *o.UiStyleConfiguration
}

// GetUiStyleConfigurationOk returns a tuple with the UiStyleConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUiStyleConfigurationOk() (*bool, bool) {
	if o == nil || IsNil(o.UiStyleConfiguration) {
		return nil, false
	}
	return o.UiStyleConfiguration, true
}

// HasUiStyleConfiguration returns a boolean if a field has been set.
func (o *Group) HasUiStyleConfiguration() bool {
	if o != nil && !IsNil(o.UiStyleConfiguration) {
		return true
	}

	return false
}

// SetUiStyleConfiguration gets a reference to the given bool and assigns it to the UiStyleConfiguration field.
func (o *Group) SetUiStyleConfiguration(v bool) {
	o.UiStyleConfiguration = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uniqid) {
		toSerialize["uniqid"] = o.Uniqid
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.ImageAttachment) {
		toSerialize["image_attachment"] = o.ImageAttachment
	}
	if !IsNil(o.Unlisted) {
		toSerialize["unlisted"] = o.Unlisted
	}
	if !IsNil(o.SortPriority) {
		toSerialize["sort_priority"] = o.SortPriority
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
	if !IsNil(o.ProductsBound) {
		toSerialize["products_bound"] = o.ProductsBound
	}
	if !IsNil(o.ImageAttachmentInfo) {
		toSerialize["image_attachment_info"] = o.ImageAttachmentInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.DarkMode) {
		toSerialize["dark_mode"] = o.DarkMode
	}
	if !IsNil(o.UiStyleConfiguration) {
		toSerialize["ui_style_configuration"] = o.UiStyleConfiguration
	}
	return toSerialize, nil
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

