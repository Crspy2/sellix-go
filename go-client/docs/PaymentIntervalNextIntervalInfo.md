# PaymentIntervalNextIntervalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of interval | [optional] 
**EndDate** | Pointer to **time.Time** | The timestamp for when the payment interval ended | [optional] 
**StartDate** | Pointer to **time.Time** | The timestamp for when the payment interval started | [optional] 
**PaymentDate** | Pointer to **time.Time** | The timestamp for when the payment was made | [optional] 
**PaymentType** | Pointer to **string** | The billing type of the payment interval | [optional] 

## Methods

### NewPaymentIntervalNextIntervalInfo

`func NewPaymentIntervalNextIntervalInfo() *PaymentIntervalNextIntervalInfo`

NewPaymentIntervalNextIntervalInfo instantiates a new PaymentIntervalNextIntervalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntervalNextIntervalInfoWithDefaults

`func NewPaymentIntervalNextIntervalInfoWithDefaults() *PaymentIntervalNextIntervalInfo`

NewPaymentIntervalNextIntervalInfoWithDefaults instantiates a new PaymentIntervalNextIntervalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentIntervalNextIntervalInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentIntervalNextIntervalInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentIntervalNextIntervalInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentIntervalNextIntervalInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEndDate

`func (o *PaymentIntervalNextIntervalInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PaymentIntervalNextIntervalInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PaymentIntervalNextIntervalInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PaymentIntervalNextIntervalInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *PaymentIntervalNextIntervalInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PaymentIntervalNextIntervalInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PaymentIntervalNextIntervalInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PaymentIntervalNextIntervalInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetPaymentDate

`func (o *PaymentIntervalNextIntervalInfo) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentIntervalNextIntervalInfo) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentIntervalNextIntervalInfo) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *PaymentIntervalNextIntervalInfo) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetPaymentType

`func (o *PaymentIntervalNextIntervalInfo) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentIntervalNextIntervalInfo) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentIntervalNextIntervalInfo) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentIntervalNextIntervalInfo) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


