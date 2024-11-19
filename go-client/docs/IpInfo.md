# IpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | IP info retrieved successfully. | [optional] 
**Message** | Pointer to **string** | Any other details linked to this IP. | [optional] 
**FraudScore** | Pointer to **int32** | IP Fraud Score details. | [optional] 
**CountryCode** | Pointer to **string** | IP country code. | [optional] 
**Region** | Pointer to **string** | IP region. | [optional] 
**City** | Pointer to **string** | IP city. | [optional] 
**ISP** | Pointer to **string** | IP ISP | [optional] 
**ASN** | Pointer to **int32** | ISP ASN | [optional] 
**OperatingSystem** | Pointer to **string** | Customer device operating system. | [optional] 
**Browser** | Pointer to **string** | Customer device browser. | [optional] 
**Organization** | Pointer to **string** | IP organization. | [optional] 
**IsCrawler** | Pointer to **bool** | If true, the IP has been recently detected as a crawler. | [optional] 
**Timezone** | Pointer to **string** | Customer timezone. | [optional] 
**Mobile** | Pointer to **bool** | If true, the customer used a mobile device. | [optional] 
**Host** | Pointer to **string** | ISP host. | [optional] 
**Proxy** | Pointer to **bool** | If true, the IP has been recently detected using a proxy. | [optional] 
**Vpn** | Pointer to **bool** | If true, the IP has been recently detected using a VPN. | [optional] 
**Tor** | Pointer to **bool** | If true, the IP has been recently detected using TOR. | [optional] 
**ActiveVpn** | Pointer to **bool** | If true, the IP has an active VPN connection. | [optional] 
**ActiveTor** | Pointer to **bool** | If true, the IP has an active TOR connection. | [optional] 
**DeviceBrand** | Pointer to **string** | Customer device brand. | [optional] 
**DeviceModel** | Pointer to **string** | Customer device model. | [optional] 
**RecentAbuse** | Pointer to **bool** | If true, the IP has been recently detected in an online abuse. | [optional] 
**BotStatus** | Pointer to **bool** | If true, the IP has been recently detected as a BOT. | [optional] 
**ConnectionType** | Pointer to **string** | Customer connection type. | [optional] 
**AbuseVelocity** | Pointer to **string** | IP abuse velocity. | [optional] 
**ZipCode** | Pointer to **string** | If detected, customer ZIP code. | [optional] 
**Latitude** | Pointer to **float64** | IP latitude. | [optional] 
**Longitude** | Pointer to **float64** | IP longitude. | [optional] 
**RequestId** | Pointer to **string** | IP request ID used for internal reference. | [optional] 

## Methods

### NewIpInfo

`func NewIpInfo() *IpInfo`

NewIpInfo instantiates a new IpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpInfoWithDefaults

`func NewIpInfoWithDefaults() *IpInfo`

NewIpInfoWithDefaults instantiates a new IpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *IpInfo) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *IpInfo) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *IpInfo) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *IpInfo) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *IpInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IpInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IpInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IpInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFraudScore

`func (o *IpInfo) GetFraudScore() int32`

GetFraudScore returns the FraudScore field if non-nil, zero value otherwise.

### GetFraudScoreOk

`func (o *IpInfo) GetFraudScoreOk() (*int32, bool)`

GetFraudScoreOk returns a tuple with the FraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudScore

`func (o *IpInfo) SetFraudScore(v int32)`

SetFraudScore sets FraudScore field to given value.

### HasFraudScore

`func (o *IpInfo) HasFraudScore() bool`

HasFraudScore returns a boolean if a field has been set.

### GetCountryCode

`func (o *IpInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *IpInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *IpInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *IpInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetRegion

`func (o *IpInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IpInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IpInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IpInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCity

`func (o *IpInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IpInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IpInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IpInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetISP

`func (o *IpInfo) GetISP() string`

GetISP returns the ISP field if non-nil, zero value otherwise.

### GetISPOk

`func (o *IpInfo) GetISPOk() (*string, bool)`

GetISPOk returns a tuple with the ISP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISP

`func (o *IpInfo) SetISP(v string)`

SetISP sets ISP field to given value.

### HasISP

`func (o *IpInfo) HasISP() bool`

HasISP returns a boolean if a field has been set.

### GetASN

`func (o *IpInfo) GetASN() int32`

GetASN returns the ASN field if non-nil, zero value otherwise.

### GetASNOk

`func (o *IpInfo) GetASNOk() (*int32, bool)`

GetASNOk returns a tuple with the ASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASN

`func (o *IpInfo) SetASN(v int32)`

SetASN sets ASN field to given value.

### HasASN

`func (o *IpInfo) HasASN() bool`

HasASN returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *IpInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *IpInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *IpInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *IpInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetBrowser

`func (o *IpInfo) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *IpInfo) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *IpInfo) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *IpInfo) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetOrganization

`func (o *IpInfo) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IpInfo) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IpInfo) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IpInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetIsCrawler

`func (o *IpInfo) GetIsCrawler() bool`

GetIsCrawler returns the IsCrawler field if non-nil, zero value otherwise.

### GetIsCrawlerOk

`func (o *IpInfo) GetIsCrawlerOk() (*bool, bool)`

GetIsCrawlerOk returns a tuple with the IsCrawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCrawler

`func (o *IpInfo) SetIsCrawler(v bool)`

SetIsCrawler sets IsCrawler field to given value.

### HasIsCrawler

`func (o *IpInfo) HasIsCrawler() bool`

HasIsCrawler returns a boolean if a field has been set.

### GetTimezone

`func (o *IpInfo) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IpInfo) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IpInfo) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *IpInfo) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetMobile

`func (o *IpInfo) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *IpInfo) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *IpInfo) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *IpInfo) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetHost

`func (o *IpInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IpInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IpInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IpInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetProxy

`func (o *IpInfo) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IpInfo) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IpInfo) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IpInfo) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetVpn

`func (o *IpInfo) GetVpn() bool`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *IpInfo) GetVpnOk() (*bool, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *IpInfo) SetVpn(v bool)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *IpInfo) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetTor

`func (o *IpInfo) GetTor() bool`

GetTor returns the Tor field if non-nil, zero value otherwise.

### GetTorOk

`func (o *IpInfo) GetTorOk() (*bool, bool)`

GetTorOk returns a tuple with the Tor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTor

`func (o *IpInfo) SetTor(v bool)`

SetTor sets Tor field to given value.

### HasTor

`func (o *IpInfo) HasTor() bool`

HasTor returns a boolean if a field has been set.

### GetActiveVpn

`func (o *IpInfo) GetActiveVpn() bool`

GetActiveVpn returns the ActiveVpn field if non-nil, zero value otherwise.

### GetActiveVpnOk

`func (o *IpInfo) GetActiveVpnOk() (*bool, bool)`

GetActiveVpnOk returns a tuple with the ActiveVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVpn

`func (o *IpInfo) SetActiveVpn(v bool)`

SetActiveVpn sets ActiveVpn field to given value.

### HasActiveVpn

`func (o *IpInfo) HasActiveVpn() bool`

HasActiveVpn returns a boolean if a field has been set.

### GetActiveTor

`func (o *IpInfo) GetActiveTor() bool`

GetActiveTor returns the ActiveTor field if non-nil, zero value otherwise.

### GetActiveTorOk

`func (o *IpInfo) GetActiveTorOk() (*bool, bool)`

GetActiveTorOk returns a tuple with the ActiveTor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTor

`func (o *IpInfo) SetActiveTor(v bool)`

SetActiveTor sets ActiveTor field to given value.

### HasActiveTor

`func (o *IpInfo) HasActiveTor() bool`

HasActiveTor returns a boolean if a field has been set.

### GetDeviceBrand

`func (o *IpInfo) GetDeviceBrand() string`

GetDeviceBrand returns the DeviceBrand field if non-nil, zero value otherwise.

### GetDeviceBrandOk

`func (o *IpInfo) GetDeviceBrandOk() (*string, bool)`

GetDeviceBrandOk returns a tuple with the DeviceBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBrand

`func (o *IpInfo) SetDeviceBrand(v string)`

SetDeviceBrand sets DeviceBrand field to given value.

### HasDeviceBrand

`func (o *IpInfo) HasDeviceBrand() bool`

HasDeviceBrand returns a boolean if a field has been set.

### GetDeviceModel

`func (o *IpInfo) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IpInfo) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *IpInfo) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *IpInfo) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetRecentAbuse

`func (o *IpInfo) GetRecentAbuse() bool`

GetRecentAbuse returns the RecentAbuse field if non-nil, zero value otherwise.

### GetRecentAbuseOk

`func (o *IpInfo) GetRecentAbuseOk() (*bool, bool)`

GetRecentAbuseOk returns a tuple with the RecentAbuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentAbuse

`func (o *IpInfo) SetRecentAbuse(v bool)`

SetRecentAbuse sets RecentAbuse field to given value.

### HasRecentAbuse

`func (o *IpInfo) HasRecentAbuse() bool`

HasRecentAbuse returns a boolean if a field has been set.

### GetBotStatus

`func (o *IpInfo) GetBotStatus() bool`

GetBotStatus returns the BotStatus field if non-nil, zero value otherwise.

### GetBotStatusOk

`func (o *IpInfo) GetBotStatusOk() (*bool, bool)`

GetBotStatusOk returns a tuple with the BotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotStatus

`func (o *IpInfo) SetBotStatus(v bool)`

SetBotStatus sets BotStatus field to given value.

### HasBotStatus

`func (o *IpInfo) HasBotStatus() bool`

HasBotStatus returns a boolean if a field has been set.

### GetConnectionType

`func (o *IpInfo) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IpInfo) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IpInfo) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *IpInfo) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetAbuseVelocity

`func (o *IpInfo) GetAbuseVelocity() string`

GetAbuseVelocity returns the AbuseVelocity field if non-nil, zero value otherwise.

### GetAbuseVelocityOk

`func (o *IpInfo) GetAbuseVelocityOk() (*string, bool)`

GetAbuseVelocityOk returns a tuple with the AbuseVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseVelocity

`func (o *IpInfo) SetAbuseVelocity(v string)`

SetAbuseVelocity sets AbuseVelocity field to given value.

### HasAbuseVelocity

`func (o *IpInfo) HasAbuseVelocity() bool`

HasAbuseVelocity returns a boolean if a field has been set.

### GetZipCode

`func (o *IpInfo) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *IpInfo) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *IpInfo) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *IpInfo) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetLatitude

`func (o *IpInfo) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *IpInfo) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *IpInfo) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *IpInfo) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *IpInfo) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *IpInfo) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *IpInfo) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *IpInfo) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetRequestId

`func (o *IpInfo) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IpInfo) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IpInfo) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *IpInfo) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


