# QueryListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this query belongs. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice this query is linked to, if specified by the customer. | [optional] 
**CustomerEmail** | Pointer to **string** | Email of the customer who created this query. | [optional] 
**Title** | Pointer to **string** | Query title, brief summary of what the customer needs. | [optional] 
**Status** | Pointer to **string** | Status of the query. PENDING if the query has been created and awaits a reply from the merchant. CLOSED if the query has been closed by the merchant or the customer. SHOP_REPLY if the query has been replied by the merchant, CUSTOMER_REPLY if the query has been replied by the customer. | [optional] 
**Country** | Pointer to **string** | The country the query was made from | [optional] 
**Priority** | Pointer to **int32** | The priority of the support query | [optional] 
**Category** | Pointer to **string** | The category of the support query | [optional] 
**DayValue** | Pointer to **int32** | Day value, number. | [optional] 
**Month** | Pointer to **string** | First three letters of the month name. | [optional] 
**Year** | Pointer to **int32** | Year number. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the query. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the query has been updated. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the query. | [optional] 
**LastMessage** | Pointer to **string** | The last message in the support query | [optional] 

## Methods

### NewQueryListing

`func NewQueryListing() *QueryListing`

NewQueryListing instantiates a new QueryListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryListingWithDefaults

`func NewQueryListingWithDefaults() *QueryListing`

NewQueryListingWithDefaults instantiates a new QueryListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryListing) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryListing) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryListing) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QueryListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *QueryListing) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *QueryListing) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *QueryListing) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *QueryListing) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *QueryListing) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *QueryListing) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *QueryListing) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *QueryListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *QueryListing) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *QueryListing) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *QueryListing) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *QueryListing) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *QueryListing) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *QueryListing) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *QueryListing) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *QueryListing) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetTitle

`func (o *QueryListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QueryListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QueryListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QueryListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *QueryListing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryListing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryListing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryListing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCountry

`func (o *QueryListing) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QueryListing) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QueryListing) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *QueryListing) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPriority

`func (o *QueryListing) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueryListing) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueryListing) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueryListing) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCategory

`func (o *QueryListing) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *QueryListing) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *QueryListing) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *QueryListing) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDayValue

`func (o *QueryListing) GetDayValue() int32`

GetDayValue returns the DayValue field if non-nil, zero value otherwise.

### GetDayValueOk

`func (o *QueryListing) GetDayValueOk() (*int32, bool)`

GetDayValueOk returns a tuple with the DayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayValue

`func (o *QueryListing) SetDayValue(v int32)`

SetDayValue sets DayValue field to given value.

### HasDayValue

`func (o *QueryListing) HasDayValue() bool`

HasDayValue returns a boolean if a field has been set.

### GetMonth

`func (o *QueryListing) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *QueryListing) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *QueryListing) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *QueryListing) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *QueryListing) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *QueryListing) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *QueryListing) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *QueryListing) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QueryListing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueryListing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueryListing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QueryListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QueryListing) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QueryListing) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QueryListing) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QueryListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *QueryListing) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *QueryListing) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *QueryListing) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *QueryListing) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetLastMessage

`func (o *QueryListing) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *QueryListing) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *QueryListing) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *QueryListing) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


