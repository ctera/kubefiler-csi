# Share

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Directory** | Pointer to **string** |  | [optional] 
**Acl** | Pointer to [**[]ShareAccessControlEntry**](ShareAccessControlEntry.md) |  | [optional] 
**Access** | Pointer to [**AclAccess**](AclAccess.md) |  | [optional] 
**ClientSideCaching** | Pointer to [**ClientSideCaching**](ClientSideCaching.md) |  | [optional] 
**DirPermissions** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ExportToAfp** | Pointer to **bool** |  | [optional] 
**ExportToFtp** | Pointer to **bool** |  | [optional] 
**ExportToNfs** | Pointer to **bool** |  | [optional] 
**ExportToPcAgent** | Pointer to **bool** |  | [optional] 
**ExportToRsync** | Pointer to **bool** |  | [optional] 
**Indexed** | Pointer to **bool** |  | [optional] 
**TrustedNfsClients** | Pointer to [**[]NFSv3AccessControlEntry**](NFSv3AccessControlEntry.md) |  | [optional] 

## Methods

### NewShare

`func NewShare(name string, ) *Share`

NewShare instantiates a new Share object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareWithDefaults

`func NewShareWithDefaults() *Share`

NewShareWithDefaults instantiates a new Share object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Share) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Share) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Share) SetName(v string)`

SetName sets Name field to given value.


### GetDirectory

`func (o *Share) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *Share) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *Share) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *Share) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetAcl

`func (o *Share) GetAcl() []ShareAccessControlEntry`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *Share) GetAclOk() (*[]ShareAccessControlEntry, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *Share) SetAcl(v []ShareAccessControlEntry)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *Share) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetAccess

`func (o *Share) GetAccess() AclAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Share) GetAccessOk() (*AclAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Share) SetAccess(v AclAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Share) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetClientSideCaching

`func (o *Share) GetClientSideCaching() ClientSideCaching`

GetClientSideCaching returns the ClientSideCaching field if non-nil, zero value otherwise.

### GetClientSideCachingOk

`func (o *Share) GetClientSideCachingOk() (*ClientSideCaching, bool)`

GetClientSideCachingOk returns a tuple with the ClientSideCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideCaching

`func (o *Share) SetClientSideCaching(v ClientSideCaching)`

SetClientSideCaching sets ClientSideCaching field to given value.

### HasClientSideCaching

`func (o *Share) HasClientSideCaching() bool`

HasClientSideCaching returns a boolean if a field has been set.

### GetDirPermissions

`func (o *Share) GetDirPermissions() int32`

GetDirPermissions returns the DirPermissions field if non-nil, zero value otherwise.

### GetDirPermissionsOk

`func (o *Share) GetDirPermissionsOk() (*int32, bool)`

GetDirPermissionsOk returns a tuple with the DirPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirPermissions

`func (o *Share) SetDirPermissions(v int32)`

SetDirPermissions sets DirPermissions field to given value.

### HasDirPermissions

`func (o *Share) HasDirPermissions() bool`

HasDirPermissions returns a boolean if a field has been set.

### GetComment

`func (o *Share) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Share) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Share) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Share) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExportToAfp

`func (o *Share) GetExportToAfp() bool`

GetExportToAfp returns the ExportToAfp field if non-nil, zero value otherwise.

### GetExportToAfpOk

`func (o *Share) GetExportToAfpOk() (*bool, bool)`

GetExportToAfpOk returns a tuple with the ExportToAfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToAfp

`func (o *Share) SetExportToAfp(v bool)`

SetExportToAfp sets ExportToAfp field to given value.

### HasExportToAfp

`func (o *Share) HasExportToAfp() bool`

HasExportToAfp returns a boolean if a field has been set.

### GetExportToFtp

`func (o *Share) GetExportToFtp() bool`

GetExportToFtp returns the ExportToFtp field if non-nil, zero value otherwise.

### GetExportToFtpOk

`func (o *Share) GetExportToFtpOk() (*bool, bool)`

GetExportToFtpOk returns a tuple with the ExportToFtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToFtp

`func (o *Share) SetExportToFtp(v bool)`

SetExportToFtp sets ExportToFtp field to given value.

### HasExportToFtp

`func (o *Share) HasExportToFtp() bool`

HasExportToFtp returns a boolean if a field has been set.

### GetExportToNfs

`func (o *Share) GetExportToNfs() bool`

GetExportToNfs returns the ExportToNfs field if non-nil, zero value otherwise.

### GetExportToNfsOk

`func (o *Share) GetExportToNfsOk() (*bool, bool)`

GetExportToNfsOk returns a tuple with the ExportToNfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToNfs

`func (o *Share) SetExportToNfs(v bool)`

SetExportToNfs sets ExportToNfs field to given value.

### HasExportToNfs

`func (o *Share) HasExportToNfs() bool`

HasExportToNfs returns a boolean if a field has been set.

### GetExportToPcAgent

`func (o *Share) GetExportToPcAgent() bool`

GetExportToPcAgent returns the ExportToPcAgent field if non-nil, zero value otherwise.

### GetExportToPcAgentOk

`func (o *Share) GetExportToPcAgentOk() (*bool, bool)`

GetExportToPcAgentOk returns a tuple with the ExportToPcAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToPcAgent

`func (o *Share) SetExportToPcAgent(v bool)`

SetExportToPcAgent sets ExportToPcAgent field to given value.

### HasExportToPcAgent

`func (o *Share) HasExportToPcAgent() bool`

HasExportToPcAgent returns a boolean if a field has been set.

### GetExportToRsync

`func (o *Share) GetExportToRsync() bool`

GetExportToRsync returns the ExportToRsync field if non-nil, zero value otherwise.

### GetExportToRsyncOk

`func (o *Share) GetExportToRsyncOk() (*bool, bool)`

GetExportToRsyncOk returns a tuple with the ExportToRsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToRsync

`func (o *Share) SetExportToRsync(v bool)`

SetExportToRsync sets ExportToRsync field to given value.

### HasExportToRsync

`func (o *Share) HasExportToRsync() bool`

HasExportToRsync returns a boolean if a field has been set.

### GetIndexed

`func (o *Share) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *Share) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *Share) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *Share) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetTrustedNfsClients

`func (o *Share) GetTrustedNfsClients() []NFSv3AccessControlEntry`

GetTrustedNfsClients returns the TrustedNfsClients field if non-nil, zero value otherwise.

### GetTrustedNfsClientsOk

`func (o *Share) GetTrustedNfsClientsOk() (*[]NFSv3AccessControlEntry, bool)`

GetTrustedNfsClientsOk returns a tuple with the TrustedNfsClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedNfsClients

`func (o *Share) SetTrustedNfsClients(v []NFSv3AccessControlEntry)`

SetTrustedNfsClients sets TrustedNfsClients field to given value.

### HasTrustedNfsClients

`func (o *Share) HasTrustedNfsClients() bool`

HasTrustedNfsClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


