# ShareAccessControlEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalType** | [**PrincipalType**](PrincipalType.md) |  | 
**Name** | **string** |  | 
**Perm** | [**FileAccessMode**](FileAccessMode.md) |  | 

## Methods

### NewShareAccessControlEntry

`func NewShareAccessControlEntry(principalType PrincipalType, name string, perm FileAccessMode, ) *ShareAccessControlEntry`

NewShareAccessControlEntry instantiates a new ShareAccessControlEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareAccessControlEntryWithDefaults

`func NewShareAccessControlEntryWithDefaults() *ShareAccessControlEntry`

NewShareAccessControlEntryWithDefaults instantiates a new ShareAccessControlEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalType

`func (o *ShareAccessControlEntry) GetPrincipalType() PrincipalType`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *ShareAccessControlEntry) GetPrincipalTypeOk() (*PrincipalType, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *ShareAccessControlEntry) SetPrincipalType(v PrincipalType)`

SetPrincipalType sets PrincipalType field to given value.


### GetName

`func (o *ShareAccessControlEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShareAccessControlEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShareAccessControlEntry) SetName(v string)`

SetName sets Name field to given value.


### GetPerm

`func (o *ShareAccessControlEntry) GetPerm() FileAccessMode`

GetPerm returns the Perm field if non-nil, zero value otherwise.

### GetPermOk

`func (o *ShareAccessControlEntry) GetPermOk() (*FileAccessMode, bool)`

GetPermOk returns a tuple with the Perm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerm

`func (o *ShareAccessControlEntry) SetPerm(v FileAccessMode)`

SetPerm sets Perm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


