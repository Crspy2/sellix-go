# ListOrders200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnedInvoices** | Pointer to **float32** | The total number of invoices returned | [optional] 
**TotalInvoices** | Pointer to **float32** | The total number of invoices on the store | [optional] 
**TotalInvoicesFiltered** | Pointer to **float32** | The total number of invoices that matched the specified fitler | [optional] 
**Page** | Pointer to **float32** | The current page if pagination was specified | [optional] 
**Limit** | Pointer to **int32** | The total amount of invoices to display per page | [optional] 
**Filter** | Pointer to **string** | The filter to use against the invoices | [optional] 
**Sort** | Pointer to **string** | The fields to sort the invoices by | [optional] 
**Search** | Pointer to **string** | The fields to search for in invoices | [optional] 
**Fields** | Pointer to **[]string** | The total amount of fields returned | [optional] 
**TimeTakenMs** | Pointer to **float32** | THe amount of time, in milliseconds, it took to retrieve the orders | [optional] 
**Orders** | Pointer to [**[]InvoiceListing**](InvoiceListing.md) |  | [optional] 

## Methods

### NewListOrders200ResponseData

`func NewListOrders200ResponseData() *ListOrders200ResponseData`

NewListOrders200ResponseData instantiates a new ListOrders200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrders200ResponseDataWithDefaults

`func NewListOrders200ResponseDataWithDefaults() *ListOrders200ResponseData`

NewListOrders200ResponseDataWithDefaults instantiates a new ListOrders200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnedInvoices

`func (o *ListOrders200ResponseData) GetReturnedInvoices() float32`

GetReturnedInvoices returns the ReturnedInvoices field if non-nil, zero value otherwise.

### GetReturnedInvoicesOk

`func (o *ListOrders200ResponseData) GetReturnedInvoicesOk() (*float32, bool)`

GetReturnedInvoicesOk returns a tuple with the ReturnedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedInvoices

`func (o *ListOrders200ResponseData) SetReturnedInvoices(v float32)`

SetReturnedInvoices sets ReturnedInvoices field to given value.

### HasReturnedInvoices

`func (o *ListOrders200ResponseData) HasReturnedInvoices() bool`

HasReturnedInvoices returns a boolean if a field has been set.

### GetTotalInvoices

`func (o *ListOrders200ResponseData) GetTotalInvoices() float32`

GetTotalInvoices returns the TotalInvoices field if non-nil, zero value otherwise.

### GetTotalInvoicesOk

`func (o *ListOrders200ResponseData) GetTotalInvoicesOk() (*float32, bool)`

GetTotalInvoicesOk returns a tuple with the TotalInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvoices

`func (o *ListOrders200ResponseData) SetTotalInvoices(v float32)`

SetTotalInvoices sets TotalInvoices field to given value.

### HasTotalInvoices

`func (o *ListOrders200ResponseData) HasTotalInvoices() bool`

HasTotalInvoices returns a boolean if a field has been set.

### GetTotalInvoicesFiltered

`func (o *ListOrders200ResponseData) GetTotalInvoicesFiltered() float32`

GetTotalInvoicesFiltered returns the TotalInvoicesFiltered field if non-nil, zero value otherwise.

### GetTotalInvoicesFilteredOk

`func (o *ListOrders200ResponseData) GetTotalInvoicesFilteredOk() (*float32, bool)`

GetTotalInvoicesFilteredOk returns a tuple with the TotalInvoicesFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvoicesFiltered

`func (o *ListOrders200ResponseData) SetTotalInvoicesFiltered(v float32)`

SetTotalInvoicesFiltered sets TotalInvoicesFiltered field to given value.

### HasTotalInvoicesFiltered

`func (o *ListOrders200ResponseData) HasTotalInvoicesFiltered() bool`

HasTotalInvoicesFiltered returns a boolean if a field has been set.

### GetPage

`func (o *ListOrders200ResponseData) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListOrders200ResponseData) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListOrders200ResponseData) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListOrders200ResponseData) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLimit

`func (o *ListOrders200ResponseData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOrders200ResponseData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOrders200ResponseData) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListOrders200ResponseData) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilter

`func (o *ListOrders200ResponseData) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListOrders200ResponseData) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListOrders200ResponseData) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListOrders200ResponseData) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *ListOrders200ResponseData) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListOrders200ResponseData) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListOrders200ResponseData) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListOrders200ResponseData) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSearch

`func (o *ListOrders200ResponseData) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ListOrders200ResponseData) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ListOrders200ResponseData) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ListOrders200ResponseData) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetFields

`func (o *ListOrders200ResponseData) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ListOrders200ResponseData) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ListOrders200ResponseData) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ListOrders200ResponseData) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetTimeTakenMs

`func (o *ListOrders200ResponseData) GetTimeTakenMs() float32`

GetTimeTakenMs returns the TimeTakenMs field if non-nil, zero value otherwise.

### GetTimeTakenMsOk

`func (o *ListOrders200ResponseData) GetTimeTakenMsOk() (*float32, bool)`

GetTimeTakenMsOk returns a tuple with the TimeTakenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenMs

`func (o *ListOrders200ResponseData) SetTimeTakenMs(v float32)`

SetTimeTakenMs sets TimeTakenMs field to given value.

### HasTimeTakenMs

`func (o *ListOrders200ResponseData) HasTimeTakenMs() bool`

HasTimeTakenMs returns a boolean if a field has been set.

### GetOrders

`func (o *ListOrders200ResponseData) GetOrders() []InvoiceListing`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ListOrders200ResponseData) GetOrdersOk() (*[]InvoiceListing, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ListOrders200ResponseData) SetOrders(v []InvoiceListing)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ListOrders200ResponseData) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


