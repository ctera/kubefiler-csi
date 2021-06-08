/*
 * CTERA Gateway
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ctera

import (
	"encoding/json"
)

// Share struct for Share
type Share struct {
	Name              string                     `json:"name"`
	Directory         *string                    `json:"directory,omitempty"`
	Acl               *[]ShareAccessControlEntry `json:"acl,omitempty"`
	Access            *AclAccess                 `json:"access,omitempty"`
	ClientSideCaching *ClientSideCaching         `json:"client_side_caching,omitempty"`
	DirPermissions    *int32                     `json:"dir_permissions,omitempty"`
	Comment           *string                    `json:"comment,omitempty"`
	ExportToAfp       *bool                      `json:"export_to_afp,omitempty"`
	ExportToFtp       *bool                      `json:"export_to_ftp,omitempty"`
	ExportToNfs       *bool                      `json:"export_to_nfs,omitempty"`
	ExportToPcAgent   *bool                      `json:"export_to_pc_agent,omitempty"`
	ExportToRsync     *bool                      `json:"export_to_rsync,omitempty"`
	Indexed           *bool                      `json:"indexed,omitempty"`
	TrustedNfsClients *[]NFSv3AccessControlEntry `json:"trusted_nfs_clients,omitempty"`
}

// NewShare instantiates a new Share object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShare(name string) *Share {
	this := Share{}
	this.Name = name
	return &this
}

// NewShareWithDefaults instantiates a new Share object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareWithDefaults() *Share {
	this := Share{}
	return &this
}

// GetName returns the Name field value
func (o *Share) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Share) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Share) SetName(v string) {
	o.Name = v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *Share) GetDirectory() string {
	if o == nil || o.Directory == nil {
		var ret string
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetDirectoryOk() (*string, bool) {
	if o == nil || o.Directory == nil {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *Share) HasDirectory() bool {
	if o != nil && o.Directory != nil {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given string and assigns it to the Directory field.
func (o *Share) SetDirectory(v string) {
	o.Directory = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *Share) GetAcl() []ShareAccessControlEntry {
	if o == nil || o.Acl == nil {
		var ret []ShareAccessControlEntry
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetAclOk() (*[]ShareAccessControlEntry, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *Share) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given []ShareAccessControlEntry and assigns it to the Acl field.
func (o *Share) SetAcl(v []ShareAccessControlEntry) {
	o.Acl = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *Share) GetAccess() AclAccess {
	if o == nil || o.Access == nil {
		var ret AclAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetAccessOk() (*AclAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *Share) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AclAccess and assigns it to the Access field.
func (o *Share) SetAccess(v AclAccess) {
	o.Access = &v
}

// GetClientSideCaching returns the ClientSideCaching field value if set, zero value otherwise.
func (o *Share) GetClientSideCaching() ClientSideCaching {
	if o == nil || o.ClientSideCaching == nil {
		var ret ClientSideCaching
		return ret
	}
	return *o.ClientSideCaching
}

// GetClientSideCachingOk returns a tuple with the ClientSideCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetClientSideCachingOk() (*ClientSideCaching, bool) {
	if o == nil || o.ClientSideCaching == nil {
		return nil, false
	}
	return o.ClientSideCaching, true
}

// HasClientSideCaching returns a boolean if a field has been set.
func (o *Share) HasClientSideCaching() bool {
	if o != nil && o.ClientSideCaching != nil {
		return true
	}

	return false
}

// SetClientSideCaching gets a reference to the given ClientSideCaching and assigns it to the ClientSideCaching field.
func (o *Share) SetClientSideCaching(v ClientSideCaching) {
	o.ClientSideCaching = &v
}

// GetDirPermissions returns the DirPermissions field value if set, zero value otherwise.
func (o *Share) GetDirPermissions() int32 {
	if o == nil || o.DirPermissions == nil {
		var ret int32
		return ret
	}
	return *o.DirPermissions
}

// GetDirPermissionsOk returns a tuple with the DirPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetDirPermissionsOk() (*int32, bool) {
	if o == nil || o.DirPermissions == nil {
		return nil, false
	}
	return o.DirPermissions, true
}

// HasDirPermissions returns a boolean if a field has been set.
func (o *Share) HasDirPermissions() bool {
	if o != nil && o.DirPermissions != nil {
		return true
	}

	return false
}

// SetDirPermissions gets a reference to the given int32 and assigns it to the DirPermissions field.
func (o *Share) SetDirPermissions(v int32) {
	o.DirPermissions = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Share) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Share) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Share) SetComment(v string) {
	o.Comment = &v
}

// GetExportToAfp returns the ExportToAfp field value if set, zero value otherwise.
func (o *Share) GetExportToAfp() bool {
	if o == nil || o.ExportToAfp == nil {
		var ret bool
		return ret
	}
	return *o.ExportToAfp
}

// GetExportToAfpOk returns a tuple with the ExportToAfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetExportToAfpOk() (*bool, bool) {
	if o == nil || o.ExportToAfp == nil {
		return nil, false
	}
	return o.ExportToAfp, true
}

// HasExportToAfp returns a boolean if a field has been set.
func (o *Share) HasExportToAfp() bool {
	if o != nil && o.ExportToAfp != nil {
		return true
	}

	return false
}

// SetExportToAfp gets a reference to the given bool and assigns it to the ExportToAfp field.
func (o *Share) SetExportToAfp(v bool) {
	o.ExportToAfp = &v
}

// GetExportToFtp returns the ExportToFtp field value if set, zero value otherwise.
func (o *Share) GetExportToFtp() bool {
	if o == nil || o.ExportToFtp == nil {
		var ret bool
		return ret
	}
	return *o.ExportToFtp
}

// GetExportToFtpOk returns a tuple with the ExportToFtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetExportToFtpOk() (*bool, bool) {
	if o == nil || o.ExportToFtp == nil {
		return nil, false
	}
	return o.ExportToFtp, true
}

// HasExportToFtp returns a boolean if a field has been set.
func (o *Share) HasExportToFtp() bool {
	if o != nil && o.ExportToFtp != nil {
		return true
	}

	return false
}

// SetExportToFtp gets a reference to the given bool and assigns it to the ExportToFtp field.
func (o *Share) SetExportToFtp(v bool) {
	o.ExportToFtp = &v
}

// GetExportToNfs returns the ExportToNfs field value if set, zero value otherwise.
func (o *Share) GetExportToNfs() bool {
	if o == nil || o.ExportToNfs == nil {
		var ret bool
		return ret
	}
	return *o.ExportToNfs
}

// GetExportToNfsOk returns a tuple with the ExportToNfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetExportToNfsOk() (*bool, bool) {
	if o == nil || o.ExportToNfs == nil {
		return nil, false
	}
	return o.ExportToNfs, true
}

// HasExportToNfs returns a boolean if a field has been set.
func (o *Share) HasExportToNfs() bool {
	if o != nil && o.ExportToNfs != nil {
		return true
	}

	return false
}

// SetExportToNfs gets a reference to the given bool and assigns it to the ExportToNfs field.
func (o *Share) SetExportToNfs(v bool) {
	o.ExportToNfs = &v
}

// GetExportToPcAgent returns the ExportToPcAgent field value if set, zero value otherwise.
func (o *Share) GetExportToPcAgent() bool {
	if o == nil || o.ExportToPcAgent == nil {
		var ret bool
		return ret
	}
	return *o.ExportToPcAgent
}

// GetExportToPcAgentOk returns a tuple with the ExportToPcAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetExportToPcAgentOk() (*bool, bool) {
	if o == nil || o.ExportToPcAgent == nil {
		return nil, false
	}
	return o.ExportToPcAgent, true
}

// HasExportToPcAgent returns a boolean if a field has been set.
func (o *Share) HasExportToPcAgent() bool {
	if o != nil && o.ExportToPcAgent != nil {
		return true
	}

	return false
}

// SetExportToPcAgent gets a reference to the given bool and assigns it to the ExportToPcAgent field.
func (o *Share) SetExportToPcAgent(v bool) {
	o.ExportToPcAgent = &v
}

// GetExportToRsync returns the ExportToRsync field value if set, zero value otherwise.
func (o *Share) GetExportToRsync() bool {
	if o == nil || o.ExportToRsync == nil {
		var ret bool
		return ret
	}
	return *o.ExportToRsync
}

// GetExportToRsyncOk returns a tuple with the ExportToRsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetExportToRsyncOk() (*bool, bool) {
	if o == nil || o.ExportToRsync == nil {
		return nil, false
	}
	return o.ExportToRsync, true
}

// HasExportToRsync returns a boolean if a field has been set.
func (o *Share) HasExportToRsync() bool {
	if o != nil && o.ExportToRsync != nil {
		return true
	}

	return false
}

// SetExportToRsync gets a reference to the given bool and assigns it to the ExportToRsync field.
func (o *Share) SetExportToRsync(v bool) {
	o.ExportToRsync = &v
}

// GetIndexed returns the Indexed field value if set, zero value otherwise.
func (o *Share) GetIndexed() bool {
	if o == nil || o.Indexed == nil {
		var ret bool
		return ret
	}
	return *o.Indexed
}

// GetIndexedOk returns a tuple with the Indexed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetIndexedOk() (*bool, bool) {
	if o == nil || o.Indexed == nil {
		return nil, false
	}
	return o.Indexed, true
}

// HasIndexed returns a boolean if a field has been set.
func (o *Share) HasIndexed() bool {
	if o != nil && o.Indexed != nil {
		return true
	}

	return false
}

// SetIndexed gets a reference to the given bool and assigns it to the Indexed field.
func (o *Share) SetIndexed(v bool) {
	o.Indexed = &v
}

// GetTrustedNfsClients returns the TrustedNfsClients field value if set, zero value otherwise.
func (o *Share) GetTrustedNfsClients() []NFSv3AccessControlEntry {
	if o == nil || o.TrustedNfsClients == nil {
		var ret []NFSv3AccessControlEntry
		return ret
	}
	return *o.TrustedNfsClients
}

// GetTrustedNfsClientsOk returns a tuple with the TrustedNfsClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetTrustedNfsClientsOk() (*[]NFSv3AccessControlEntry, bool) {
	if o == nil || o.TrustedNfsClients == nil {
		return nil, false
	}
	return o.TrustedNfsClients, true
}

// HasTrustedNfsClients returns a boolean if a field has been set.
func (o *Share) HasTrustedNfsClients() bool {
	if o != nil && o.TrustedNfsClients != nil {
		return true
	}

	return false
}

// SetTrustedNfsClients gets a reference to the given []NFSv3AccessControlEntry and assigns it to the TrustedNfsClients field.
func (o *Share) SetTrustedNfsClients(v []NFSv3AccessControlEntry) {
	o.TrustedNfsClients = &v
}

func (o Share) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Directory != nil {
		toSerialize["directory"] = o.Directory
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.ClientSideCaching != nil {
		toSerialize["client_side_caching"] = o.ClientSideCaching
	}
	if o.DirPermissions != nil {
		toSerialize["dir_permissions"] = o.DirPermissions
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.ExportToAfp != nil {
		toSerialize["export_to_afp"] = o.ExportToAfp
	}
	if o.ExportToFtp != nil {
		toSerialize["export_to_ftp"] = o.ExportToFtp
	}
	if o.ExportToNfs != nil {
		toSerialize["export_to_nfs"] = o.ExportToNfs
	}
	if o.ExportToPcAgent != nil {
		toSerialize["export_to_pc_agent"] = o.ExportToPcAgent
	}
	if o.ExportToRsync != nil {
		toSerialize["export_to_rsync"] = o.ExportToRsync
	}
	if o.Indexed != nil {
		toSerialize["indexed"] = o.Indexed
	}
	if o.TrustedNfsClients != nil {
		toSerialize["trusted_nfs_clients"] = o.TrustedNfsClients
	}
	return json.Marshal(toSerialize)
}

type NullableShare struct {
	value *Share
	isSet bool
}

func (v NullableShare) Get() *Share {
	return v.value
}

func (v *NullableShare) Set(val *Share) {
	v.value = val
	v.isSet = true
}

func (v NullableShare) IsSet() bool {
	return v.isSet
}

func (v *NullableShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShare(val *Share) *NullableShare {
	return &NullableShare{value: val, isSet: true}
}

func (v NullableShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}