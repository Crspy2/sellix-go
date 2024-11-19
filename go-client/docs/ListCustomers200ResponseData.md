# ListCustomers200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | Pointer to [**[]CustomerListing**](CustomerListing.md) |  | [optional] 

## Methods

### NewListCustomers200ResponseData

`func NewListCustomers200ResponseData() *ListCustomers200ResponseData`

NewListCustomers200ResponseData instantiates a new ListCustomers200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomers200ResponseDataWithDefaults

`func NewListCustomers200ResponseDataWithDefaults() *ListCustomers200ResponseData`

NewListCustomers200ResponseDataWithDefaults instantiates a new ListCustomers200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *ListCustomers200ResponseData) GetCustomers() []CustomerListing`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *ListCustomers200ResponseData) GetCustomersOk() (*[]CustomerListing, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *ListCustomers200ResponseData) SetCustomers(v []CustomerListing)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *ListCustomers200ResponseData) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


