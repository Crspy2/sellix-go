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

// checks if the AmlWalletClusterEntitiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmlWalletClusterEntitiesInner{}

// AmlWalletClusterEntitiesInner struct for AmlWalletClusterEntitiesInner
type AmlWalletClusterEntitiesInner struct {
	// Entity name.
	Name *string `json:"name,omitempty"`
	// Is a VASP.
	IsVasp *bool `json:"is_vasp,omitempty"`
	// Actor ID.
	ActorId *int32 `json:"actor_id,omitempty"`
	// Entity category.
	Category *string `json:"category,omitempty"`
	// Is the primary entity.
	IsPrimaryEntity *bool `json:"is_primary_entity,omitempty"`
	// Is after the sanction date.
	IsAfterSanctionDate *bool `json:"is_after_sanction_date,omitempty"`
}

// NewAmlWalletClusterEntitiesInner instantiates a new AmlWalletClusterEntitiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmlWalletClusterEntitiesInner() *AmlWalletClusterEntitiesInner {
	this := AmlWalletClusterEntitiesInner{}
	return &this
}

// NewAmlWalletClusterEntitiesInnerWithDefaults instantiates a new AmlWalletClusterEntitiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmlWalletClusterEntitiesInnerWithDefaults() *AmlWalletClusterEntitiesInner {
	this := AmlWalletClusterEntitiesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AmlWalletClusterEntitiesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletClusterEntitiesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AmlWalletClusterEntitiesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AmlWalletClusterEntitiesInner) SetName(v string) {
	o.Name = &v
}

// GetIsVasp returns the IsVasp field value if set, zero value otherwise.
func (o *AmlWalletClusterEntitiesInner) GetIsVasp() bool {
	if o == nil || IsNil(o.IsVasp) {
		var ret bool
		return ret
	}
	return *o.IsVasp
}

// GetIsVaspOk returns a tuple with the IsVasp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletClusterEntitiesInner) GetIsVaspOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVasp) {
		return nil, false
	}
	return o.IsVasp, true
}

// HasIsVasp returns a boolean if a field has been set.
func (o *AmlWalletClusterEntitiesInner) HasIsVasp() bool {
	if o != nil && !IsNil(o.IsVasp) {
		return true
	}

	return false
}

// SetIsVasp gets a reference to the given bool and assigns it to the IsVasp field.
func (o *AmlWalletClusterEntitiesInner) SetIsVasp(v bool) {
	o.IsVasp = &v
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *AmlWalletClusterEntitiesInner) GetActorId() int32 {
	if o == nil || IsNil(o.ActorId) {
		var ret int32
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletClusterEntitiesInner) GetActorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActorId) {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *AmlWalletClusterEntitiesInner) HasActorId() bool {
	if o != nil && !IsNil(o.ActorId) {
		return true
	}

	return false
}

// SetActorId gets a reference to the given int32 and assigns it to the ActorId field.
func (o *AmlWalletClusterEntitiesInner) SetActorId(v int32) {
	o.ActorId = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AmlWalletClusterEntitiesInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletClusterEntitiesInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AmlWalletClusterEntitiesInner) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AmlWalletClusterEntitiesInner) SetCategory(v string) {
	o.Category = &v
}

// GetIsPrimaryEntity returns the IsPrimaryEntity field value if set, zero value otherwise.
func (o *AmlWalletClusterEntitiesInner) GetIsPrimaryEntity() bool {
	if o == nil || IsNil(o.IsPrimaryEntity) {
		var ret bool
		return ret
	}
	return *o.IsPrimaryEntity
}

// GetIsPrimaryEntityOk returns a tuple with the IsPrimaryEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletClusterEntitiesInner) GetIsPrimaryEntityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimaryEntity) {
		return nil, false
	}
	return o.IsPrimaryEntity, true
}

// HasIsPrimaryEntity returns a boolean if a field has been set.
func (o *AmlWalletClusterEntitiesInner) HasIsPrimaryEntity() bool {
	if o != nil && !IsNil(o.IsPrimaryEntity) {
		return true
	}

	return false
}

// SetIsPrimaryEntity gets a reference to the given bool and assigns it to the IsPrimaryEntity field.
func (o *AmlWalletClusterEntitiesInner) SetIsPrimaryEntity(v bool) {
	o.IsPrimaryEntity = &v
}

// GetIsAfterSanctionDate returns the IsAfterSanctionDate field value if set, zero value otherwise.
func (o *AmlWalletClusterEntitiesInner) GetIsAfterSanctionDate() bool {
	if o == nil || IsNil(o.IsAfterSanctionDate) {
		var ret bool
		return ret
	}
	return *o.IsAfterSanctionDate
}

// GetIsAfterSanctionDateOk returns a tuple with the IsAfterSanctionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmlWalletClusterEntitiesInner) GetIsAfterSanctionDateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAfterSanctionDate) {
		return nil, false
	}
	return o.IsAfterSanctionDate, true
}

// HasIsAfterSanctionDate returns a boolean if a field has been set.
func (o *AmlWalletClusterEntitiesInner) HasIsAfterSanctionDate() bool {
	if o != nil && !IsNil(o.IsAfterSanctionDate) {
		return true
	}

	return false
}

// SetIsAfterSanctionDate gets a reference to the given bool and assigns it to the IsAfterSanctionDate field.
func (o *AmlWalletClusterEntitiesInner) SetIsAfterSanctionDate(v bool) {
	o.IsAfterSanctionDate = &v
}

func (o AmlWalletClusterEntitiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmlWalletClusterEntitiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsVasp) {
		toSerialize["is_vasp"] = o.IsVasp
	}
	if !IsNil(o.ActorId) {
		toSerialize["actor_id"] = o.ActorId
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.IsPrimaryEntity) {
		toSerialize["is_primary_entity"] = o.IsPrimaryEntity
	}
	if !IsNil(o.IsAfterSanctionDate) {
		toSerialize["is_after_sanction_date"] = o.IsAfterSanctionDate
	}
	return toSerialize, nil
}

type NullableAmlWalletClusterEntitiesInner struct {
	value *AmlWalletClusterEntitiesInner
	isSet bool
}

func (v NullableAmlWalletClusterEntitiesInner) Get() *AmlWalletClusterEntitiesInner {
	return v.value
}

func (v *NullableAmlWalletClusterEntitiesInner) Set(val *AmlWalletClusterEntitiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAmlWalletClusterEntitiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAmlWalletClusterEntitiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmlWalletClusterEntitiesInner(val *AmlWalletClusterEntitiesInner) *NullableAmlWalletClusterEntitiesInner {
	return &NullableAmlWalletClusterEntitiesInner{value: val, isSet: true}
}

func (v NullableAmlWalletClusterEntitiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmlWalletClusterEntitiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


