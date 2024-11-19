# GatewayFees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this gateway fee belongs. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**PercentAmount** | Pointer to **float32** | The percent amount of the gateway fee. | [optional] 
**FixedAmount** | Pointer to **float32** | The fixed amount of the gateway fee. | [optional] 
**FixedCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ActiveType** | Pointer to **string** | The active type of the gateway fee. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | DateTime, available if the gateway fee has been updated. | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime, available if the gateway fee has been created. | [optional] 
**Conversions** | Pointer to [**GatewayFeesConversions**](GatewayFeesConversions.md) |  | [optional] 

## Methods

### NewGatewayFees

`func NewGatewayFees() *GatewayFees`

NewGatewayFees instantiates a new GatewayFees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayFeesWithDefaults

`func NewGatewayFeesWithDefaults() *GatewayFees`

NewGatewayFeesWithDefaults instantiates a new GatewayFees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayFees) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayFees) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayFees) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayFees) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShopId

`func (o *GatewayFees) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *GatewayFees) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *GatewayFees) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *GatewayFees) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetGateway

`func (o *GatewayFees) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GatewayFees) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GatewayFees) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GatewayFees) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPercentAmount

`func (o *GatewayFees) GetPercentAmount() float32`

GetPercentAmount returns the PercentAmount field if non-nil, zero value otherwise.

### GetPercentAmountOk

`func (o *GatewayFees) GetPercentAmountOk() (*float32, bool)`

GetPercentAmountOk returns a tuple with the PercentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentAmount

`func (o *GatewayFees) SetPercentAmount(v float32)`

SetPercentAmount sets PercentAmount field to given value.

### HasPercentAmount

`func (o *GatewayFees) HasPercentAmount() bool`

HasPercentAmount returns a boolean if a field has been set.

### GetFixedAmount

`func (o *GatewayFees) GetFixedAmount() float32`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *GatewayFees) GetFixedAmountOk() (*float32, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *GatewayFees) SetFixedAmount(v float32)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *GatewayFees) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetFixedCurrency

`func (o *GatewayFees) GetFixedCurrency() Currency`

GetFixedCurrency returns the FixedCurrency field if non-nil, zero value otherwise.

### GetFixedCurrencyOk

`func (o *GatewayFees) GetFixedCurrencyOk() (*Currency, bool)`

GetFixedCurrencyOk returns a tuple with the FixedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCurrency

`func (o *GatewayFees) SetFixedCurrency(v Currency)`

SetFixedCurrency sets FixedCurrency field to given value.

### HasFixedCurrency

`func (o *GatewayFees) HasFixedCurrency() bool`

HasFixedCurrency returns a boolean if a field has been set.

### GetActiveType

`func (o *GatewayFees) GetActiveType() string`

GetActiveType returns the ActiveType field if non-nil, zero value otherwise.

### GetActiveTypeOk

`func (o *GatewayFees) GetActiveTypeOk() (*string, bool)`

GetActiveTypeOk returns a tuple with the ActiveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveType

`func (o *GatewayFees) SetActiveType(v string)`

SetActiveType sets ActiveType field to given value.

### HasActiveType

`func (o *GatewayFees) HasActiveType() bool`

HasActiveType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GatewayFees) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GatewayFees) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GatewayFees) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GatewayFees) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GatewayFees) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GatewayFees) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GatewayFees) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GatewayFees) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetConversions

`func (o *GatewayFees) GetConversions() GatewayFeesConversions`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *GatewayFees) GetConversionsOk() (*GatewayFeesConversions, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *GatewayFees) SetConversions(v GatewayFeesConversions)`

SetConversions sets Conversions field to given value.

### HasConversions

`func (o *GatewayFees) HasConversions() bool`

HasConversions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


