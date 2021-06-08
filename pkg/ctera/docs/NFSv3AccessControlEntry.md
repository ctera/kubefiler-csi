# NFSv3AccessControlEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Netmask** | **string** |  | 
**Perm** | [**FileAccessMode**](FileAccessMode.md) |  | 

## Methods

### NewNFSv3AccessControlEntry

`func NewNFSv3AccessControlEntry(address string, netmask string, perm FileAccessMode, ) *NFSv3AccessControlEntry`

NewNFSv3AccessControlEntry instantiates a new NFSv3AccessControlEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSv3AccessControlEntryWithDefaults

`func NewNFSv3AccessControlEntryWithDefaults() *NFSv3AccessControlEntry`

NewNFSv3AccessControlEntryWithDefaults instantiates a new NFSv3AccessControlEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NFSv3AccessControlEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NFSv3AccessControlEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NFSv3AccessControlEntry) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetmask

`func (o *NFSv3AccessControlEntry) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *NFSv3AccessControlEntry) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *NFSv3AccessControlEntry) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetPerm

`func (o *NFSv3AccessControlEntry) GetPerm() FileAccessMode`

GetPerm returns the Perm field if non-nil, zero value otherwise.

### GetPermOk

`func (o *NFSv3AccessControlEntry) GetPermOk() (*FileAccessMode, bool)`

GetPermOk returns a tuple with the Perm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerm

`func (o *NFSv3AccessControlEntry) SetPerm(v FileAccessMode)`

SetPerm sets Perm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


