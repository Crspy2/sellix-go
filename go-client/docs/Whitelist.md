# Whitelist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this whitelist belongs. | [optional] 
**Type** | Pointer to **string** | The type of data of this whitelist. | [optional] 
**Data** | Pointer to **string** | The value of the whitelist. | [optional] 
**Note** | Pointer to **string** | Additional note provided on whitelist creation. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the whitelist. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the whitelist has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the whitelist. | [optional] 

## Methods

### NewWhitelist

`func NewWhitelist() *Whitelist`

NewWhitelist instantiates a new Whitelist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhitelistWithDefaults

`func NewWhitelistWithDefaults() *Whitelist`

NewWhitelistWithDefaults instantiates a new Whitelist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Whitelist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Whitelist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Whitelist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Whitelist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Whitelist) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Whitelist) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Whitelist) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Whitelist) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *Whitelist) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Whitelist) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Whitelist) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Whitelist) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *Whitelist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Whitelist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Whitelist) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Whitelist) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *Whitelist) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Whitelist) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Whitelist) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Whitelist) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNote

`func (o *Whitelist) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Whitelist) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Whitelist) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Whitelist) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Whitelist) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Whitelist) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Whitelist) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Whitelist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Whitelist) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Whitelist) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Whitelist) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Whitelist) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Whitelist) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Whitelist) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Whitelist) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Whitelist) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


