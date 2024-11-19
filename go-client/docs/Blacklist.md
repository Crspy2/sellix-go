# Blacklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**Scope** | Pointer to **string** | Whether it is a PRIVATE or SHARED blacklist. SHARED blacklists are created by Sellix&#39;s Fraud Shield to be used across Business and Enterprise merchants. | [optional] 
**ScopeDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this blacklist belongs. | [optional] 
**Type** | Pointer to **string** | The type of data of this blacklist. | [optional] 
**Data** | Pointer to **string** | The value of the blacklist. | [optional] 
**Note** | Pointer to **string** | Additional note provided on blacklist creation. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the blacklist. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the blacklist has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the blaclist. | [optional] 

## Methods

### NewBlacklist

`func NewBlacklist() *Blacklist`

NewBlacklist instantiates a new Blacklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlacklistWithDefaults

`func NewBlacklistWithDefaults() *Blacklist`

NewBlacklistWithDefaults instantiates a new Blacklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Blacklist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Blacklist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Blacklist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Blacklist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Blacklist) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Blacklist) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Blacklist) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Blacklist) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetScope

`func (o *Blacklist) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Blacklist) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Blacklist) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Blacklist) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeDetails

`func (o *Blacklist) GetScopeDetails() map[string]interface{}`

GetScopeDetails returns the ScopeDetails field if non-nil, zero value otherwise.

### GetScopeDetailsOk

`func (o *Blacklist) GetScopeDetailsOk() (*map[string]interface{}, bool)`

GetScopeDetailsOk returns a tuple with the ScopeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDetails

`func (o *Blacklist) SetScopeDetails(v map[string]interface{})`

SetScopeDetails sets ScopeDetails field to given value.

### HasScopeDetails

`func (o *Blacklist) HasScopeDetails() bool`

HasScopeDetails returns a boolean if a field has been set.

### GetShopId

`func (o *Blacklist) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Blacklist) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Blacklist) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Blacklist) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *Blacklist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Blacklist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Blacklist) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Blacklist) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *Blacklist) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Blacklist) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Blacklist) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Blacklist) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNote

`func (o *Blacklist) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Blacklist) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Blacklist) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Blacklist) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Blacklist) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Blacklist) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Blacklist) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Blacklist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Blacklist) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Blacklist) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Blacklist) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Blacklist) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Blacklist) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Blacklist) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Blacklist) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Blacklist) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


