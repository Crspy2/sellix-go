# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this query belongs. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice this query is linked to, if specified by the customer. | [optional] 
**CustomerEmail** | Pointer to **string** | Email of the customer who created this query. | [optional] 
**Title** | Pointer to **string** | Query title, brief summary of what the customer needs. | [optional] 
**Status** | Pointer to **string** | Status of the query. PENDING if the query has been created and awaits a reply from the merchant. CLOSED if the query has been closed by the merchant or the customer. SHOP_REPLY if the query has been replied by the merchant, CUSTOMER_REPLY if the query has been replied by the customer. | [optional] 
**DayValue** | Pointer to **int32** | Day value, number. | [optional] 
**Month** | Pointer to **string** | First three letters of the month name. | [optional] 
**Year** | Pointer to **int32** | Year number. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the query. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the query has been updated. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the query. | [optional] 
**Messages** | Pointer to [**[]QueryMessage**](QueryMessage.md) |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqid

`func (o *Query) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Query) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Query) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Query) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *Query) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Query) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Query) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Query) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Query) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Query) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Query) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Query) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *Query) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *Query) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *Query) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *Query) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetTitle

`func (o *Query) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Query) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Query) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Query) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *Query) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Query) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Query) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Query) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDayValue

`func (o *Query) GetDayValue() int32`

GetDayValue returns the DayValue field if non-nil, zero value otherwise.

### GetDayValueOk

`func (o *Query) GetDayValueOk() (*int32, bool)`

GetDayValueOk returns a tuple with the DayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayValue

`func (o *Query) SetDayValue(v int32)`

SetDayValue sets DayValue field to given value.

### HasDayValue

`func (o *Query) HasDayValue() bool`

HasDayValue returns a boolean if a field has been set.

### GetMonth

`func (o *Query) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Query) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Query) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Query) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *Query) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Query) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Query) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *Query) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Query) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Query) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Query) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Query) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Query) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Query) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Query) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Query) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Query) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Query) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Query) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Query) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetMessages

`func (o *Query) GetMessages() []QueryMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Query) GetMessagesOk() (*[]QueryMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Query) SetMessages(v []QueryMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Query) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


