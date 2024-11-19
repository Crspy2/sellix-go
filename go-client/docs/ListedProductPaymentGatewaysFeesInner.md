# ListedProductPaymentGatewaysFeesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this resource belongs. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**PercentAmount** | Pointer to **int32** |  | [optional] 
**FixedAmount** | Pointer to **int32** |  | [optional] 
**FixedCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ActiveType** | Pointer to **string** | DEPRECATED | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Datetime for the creation of the resource | [optional] 

## Methods

### NewListedProductPaymentGatewaysFeesInner

`func NewListedProductPaymentGatewaysFeesInner() *ListedProductPaymentGatewaysFeesInner`

NewListedProductPaymentGatewaysFeesInner instantiates a new ListedProductPaymentGatewaysFeesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedProductPaymentGatewaysFeesInnerWithDefaults

`func NewListedProductPaymentGatewaysFeesInnerWithDefaults() *ListedProductPaymentGatewaysFeesInner`

NewListedProductPaymentGatewaysFeesInnerWithDefaults instantiates a new ListedProductPaymentGatewaysFeesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListedProductPaymentGatewaysFeesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListedProductPaymentGatewaysFeesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListedProductPaymentGatewaysFeesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShopId

`func (o *ListedProductPaymentGatewaysFeesInner) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ListedProductPaymentGatewaysFeesInner) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ListedProductPaymentGatewaysFeesInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetGateway

`func (o *ListedProductPaymentGatewaysFeesInner) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ListedProductPaymentGatewaysFeesInner) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ListedProductPaymentGatewaysFeesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPercentAmount

`func (o *ListedProductPaymentGatewaysFeesInner) GetPercentAmount() int32`

GetPercentAmount returns the PercentAmount field if non-nil, zero value otherwise.

### GetPercentAmountOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetPercentAmountOk() (*int32, bool)`

GetPercentAmountOk returns a tuple with the PercentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentAmount

`func (o *ListedProductPaymentGatewaysFeesInner) SetPercentAmount(v int32)`

SetPercentAmount sets PercentAmount field to given value.

### HasPercentAmount

`func (o *ListedProductPaymentGatewaysFeesInner) HasPercentAmount() bool`

HasPercentAmount returns a boolean if a field has been set.

### GetFixedAmount

`func (o *ListedProductPaymentGatewaysFeesInner) GetFixedAmount() int32`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetFixedAmountOk() (*int32, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *ListedProductPaymentGatewaysFeesInner) SetFixedAmount(v int32)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *ListedProductPaymentGatewaysFeesInner) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetFixedCurrency

`func (o *ListedProductPaymentGatewaysFeesInner) GetFixedCurrency() Currency`

GetFixedCurrency returns the FixedCurrency field if non-nil, zero value otherwise.

### GetFixedCurrencyOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetFixedCurrencyOk() (*Currency, bool)`

GetFixedCurrencyOk returns a tuple with the FixedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCurrency

`func (o *ListedProductPaymentGatewaysFeesInner) SetFixedCurrency(v Currency)`

SetFixedCurrency sets FixedCurrency field to given value.

### HasFixedCurrency

`func (o *ListedProductPaymentGatewaysFeesInner) HasFixedCurrency() bool`

HasFixedCurrency returns a boolean if a field has been set.

### GetActiveType

`func (o *ListedProductPaymentGatewaysFeesInner) GetActiveType() string`

GetActiveType returns the ActiveType field if non-nil, zero value otherwise.

### GetActiveTypeOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetActiveTypeOk() (*string, bool)`

GetActiveTypeOk returns a tuple with the ActiveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveType

`func (o *ListedProductPaymentGatewaysFeesInner) SetActiveType(v string)`

SetActiveType sets ActiveType field to given value.

### HasActiveType

`func (o *ListedProductPaymentGatewaysFeesInner) HasActiveType() bool`

HasActiveType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListedProductPaymentGatewaysFeesInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListedProductPaymentGatewaysFeesInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListedProductPaymentGatewaysFeesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListedProductPaymentGatewaysFeesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListedProductPaymentGatewaysFeesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListedProductPaymentGatewaysFeesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListedProductPaymentGatewaysFeesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


