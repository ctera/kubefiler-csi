# \SharesApi

All URIs are relative to *http://localhost:9090/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CteraGatewayOpenapiApiSharesAddAcl**](SharesApi.md#CteraGatewayOpenapiApiSharesAddAcl) | **Put** /shares/{name}/acl | Add the ACLs of the share
[**CteraGatewayOpenapiApiSharesAddScreenedFileTypes**](SharesApi.md#CteraGatewayOpenapiApiSharesAddScreenedFileTypes) | **Put** /shares/{name}/screened_file_types | Add to the screened files types of the share
[**CteraGatewayOpenapiApiSharesAddTrustedNfsClients**](SharesApi.md#CteraGatewayOpenapiApiSharesAddTrustedNfsClients) | **Put** /shares/{name}/trusted_nfs_clients | Add the Trusted NFS Clients of the share
[**CteraGatewayOpenapiApiSharesGetAcl**](SharesApi.md#CteraGatewayOpenapiApiSharesGetAcl) | **Get** /shares/{name}/acl | List the ACLs of the share
[**CteraGatewayOpenapiApiSharesGetScreenedFileTypes**](SharesApi.md#CteraGatewayOpenapiApiSharesGetScreenedFileTypes) | **Get** /shares/{name}/screened_file_types | List the screened file types of the share
[**CteraGatewayOpenapiApiSharesGetTrustedNfsClients**](SharesApi.md#CteraGatewayOpenapiApiSharesGetTrustedNfsClients) | **Get** /shares/{name}/trusted_nfs_clients | List the Trusted NFS Clients of the share
[**CteraGatewayOpenapiApiSharesRemoveAcl**](SharesApi.md#CteraGatewayOpenapiApiSharesRemoveAcl) | **Delete** /shares/{name}/acl | Remove the ACL of the share
[**CteraGatewayOpenapiApiSharesRemoveScreenedFileType**](SharesApi.md#CteraGatewayOpenapiApiSharesRemoveScreenedFileType) | **Delete** /shares/{name}/screened_file_types | Remove from the screened file type of the share
[**CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient**](SharesApi.md#CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient) | **Delete** /shares/{name}/trusted_nfs_clients | Remove the Trusted NFS Client of the share
[**CteraGatewayOpenapiApiSharesSetAcl**](SharesApi.md#CteraGatewayOpenapiApiSharesSetAcl) | **Post** /shares/{name}/acl | Set the ACLs of the share (override the current ACL list)
[**CteraGatewayOpenapiApiSharesSetScreenedFileTypes**](SharesApi.md#CteraGatewayOpenapiApiSharesSetScreenedFileTypes) | **Post** /shares/{name}/screened_file_types | Set the list of screened file types of the share (override the current list)
[**CteraGatewayOpenapiApiSharesSetTrustedNfsClients**](SharesApi.md#CteraGatewayOpenapiApiSharesSetTrustedNfsClients) | **Post** /shares/{name}/trusted_nfs_clients | Set the Trusted NFS Clients of the share (override the current list)
[**SharesGet**](SharesApi.md#SharesGet) | **Get** /shares | List all shares
[**SharesNameDelete**](SharesApi.md#SharesNameDelete) | **Delete** /shares/{name} | Delete a share
[**SharesNameGet**](SharesApi.md#SharesNameGet) | **Get** /shares/{name} | Get the specified share
[**SharesNamePut**](SharesApi.md#SharesNamePut) | **Put** /shares/{name} | Update existing share
[**SharesPost**](SharesApi.md#SharesPost) | **Post** /shares | Create a new Share



## CteraGatewayOpenapiApiSharesAddAcl

> CteraGatewayOpenapiApiSharesAddAcl(ctx, name).ShareAccessControlEntry(shareAccessControlEntry).Execute()

Add the ACLs of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    shareAccessControlEntry := []openapiclient.ShareAccessControlEntry{*openapiclient.NewShareAccessControlEntry(openapiclient.PrincipalType("LU"), "Name_example", openapiclient.FileAccessMode("RW"))} // []ShareAccessControlEntry | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesAddAcl(context.Background(), name).ShareAccessControlEntry(shareAccessControlEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesAddAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesAddAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shareAccessControlEntry** | [**[]ShareAccessControlEntry**](ShareAccessControlEntry.md) | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesAddScreenedFileTypes

> CteraGatewayOpenapiApiSharesAddScreenedFileTypes(ctx, name).RequestBody(requestBody).Execute()

Add to the screened files types of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    requestBody := []string{"Property_example"} // []string | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesAddScreenedFileTypes(context.Background(), name).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesAddScreenedFileTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesAddScreenedFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesAddTrustedNfsClients

> CteraGatewayOpenapiApiSharesAddTrustedNfsClients(ctx, name).NFSv3AccessControlEntry(nFSv3AccessControlEntry).Execute()

Add the Trusted NFS Clients of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    nFSv3AccessControlEntry := []openapiclient.NFSv3AccessControlEntry{*openapiclient.NewNFSv3AccessControlEntry("Address_example", "Netmask_example", openapiclient.FileAccessMode("RW"))} // []NFSv3AccessControlEntry | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesAddTrustedNfsClients(context.Background(), name).NFSv3AccessControlEntry(nFSv3AccessControlEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesAddTrustedNfsClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesAddTrustedNfsClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nFSv3AccessControlEntry** | [**[]NFSv3AccessControlEntry**](NFSv3AccessControlEntry.md) | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesGetAcl

> [][]ShareAccessControlEntry CteraGatewayOpenapiApiSharesGetAcl(ctx, name).Execute()

List the ACLs of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesGetAcl(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesGetAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CteraGatewayOpenapiApiSharesGetAcl`: [][]ShareAccessControlEntry
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.CteraGatewayOpenapiApiSharesGetAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesGetAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[][]ShareAccessControlEntry**](array.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesGetScreenedFileTypes

> []string CteraGatewayOpenapiApiSharesGetScreenedFileTypes(ctx, name).Execute()

List the screened file types of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesGetScreenedFileTypes(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesGetScreenedFileTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CteraGatewayOpenapiApiSharesGetScreenedFileTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.CteraGatewayOpenapiApiSharesGetScreenedFileTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesGetScreenedFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesGetTrustedNfsClients

> [][]NFSv3AccessControlEntry CteraGatewayOpenapiApiSharesGetTrustedNfsClients(ctx, name).Execute()

List the Trusted NFS Clients of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesGetTrustedNfsClients(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesGetTrustedNfsClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CteraGatewayOpenapiApiSharesGetTrustedNfsClients`: [][]NFSv3AccessControlEntry
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.CteraGatewayOpenapiApiSharesGetTrustedNfsClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesGetTrustedNfsClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[][]NFSv3AccessControlEntry**](array.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesRemoveAcl

> CteraGatewayOpenapiApiSharesRemoveAcl(ctx, name).PrincipalType(principalType).PrincipalName(principalName).Execute()

Remove the ACL of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    principalType := openapiclient.PrincipalType("LU") // PrincipalType | The principal type (optional)
    principalName := "principalName_example" // string | The principal name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesRemoveAcl(context.Background(), name).PrincipalType(principalType).PrincipalName(principalName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesRemoveAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesRemoveAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **principalType** | [**PrincipalType**](PrincipalType.md) | The principal type | 
 **principalName** | **string** | The principal name | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesRemoveScreenedFileType

> CteraGatewayOpenapiApiSharesRemoveScreenedFileType(ctx, name).FileType(fileType).Execute()

Remove from the screened file type of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    fileType := "fileType_example" // string | File type to remove (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesRemoveScreenedFileType(context.Background(), name).FileType(fileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesRemoveScreenedFileType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesRemoveScreenedFileTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileType** | **string** | File type to remove | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient

> CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient(ctx, name).Address(address).Netmask(netmask).Execute()

Remove the Trusted NFS Client of the share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    address := "address_example" // string | IP Address (optional)
    netmask := "netmask_example" // string | Netmask (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient(context.Background(), name).Address(address).Netmask(netmask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesRemoveTrustedNfsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **string** | IP Address | 
 **netmask** | **string** | Netmask | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesSetAcl

> CteraGatewayOpenapiApiSharesSetAcl(ctx, name).ShareAccessControlEntry(shareAccessControlEntry).Execute()

Set the ACLs of the share (override the current ACL list)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    shareAccessControlEntry := []openapiclient.ShareAccessControlEntry{*openapiclient.NewShareAccessControlEntry(openapiclient.PrincipalType("LU"), "Name_example", openapiclient.FileAccessMode("RW"))} // []ShareAccessControlEntry | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesSetAcl(context.Background(), name).ShareAccessControlEntry(shareAccessControlEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesSetAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesSetAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shareAccessControlEntry** | [**[]ShareAccessControlEntry**](ShareAccessControlEntry.md) | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesSetScreenedFileTypes

> CteraGatewayOpenapiApiSharesSetScreenedFileTypes(ctx, name).RequestBody(requestBody).Execute()

Set the list of screened file types of the share (override the current list)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    requestBody := []string{"Property_example"} // []string | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesSetScreenedFileTypes(context.Background(), name).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesSetScreenedFileTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesSetScreenedFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CteraGatewayOpenapiApiSharesSetTrustedNfsClients

> CteraGatewayOpenapiApiSharesSetTrustedNfsClients(ctx, name).NFSv3AccessControlEntry(nFSv3AccessControlEntry).Execute()

Set the Trusted NFS Clients of the share (override the current list)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share
    nFSv3AccessControlEntry := []openapiclient.NFSv3AccessControlEntry{*openapiclient.NewNFSv3AccessControlEntry("Address_example", "Netmask_example", openapiclient.FileAccessMode("RW"))} // []NFSv3AccessControlEntry | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.CteraGatewayOpenapiApiSharesSetTrustedNfsClients(context.Background(), name).NFSv3AccessControlEntry(nFSv3AccessControlEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.CteraGatewayOpenapiApiSharesSetTrustedNfsClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiCteraGatewayOpenapiApiSharesSetTrustedNfsClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nFSv3AccessControlEntry** | [**[]NFSv3AccessControlEntry**](NFSv3AccessControlEntry.md) | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesGet

> []Share SharesGet(ctx).Execute()

List all shares

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.SharesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.SharesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGet`: []Share
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.SharesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetRequest struct via the builder pattern


### Return type

[**[]Share**](Share.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesNameDelete

> SharesNameDelete(ctx, name).Execute()

Delete a share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.SharesNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.SharesNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesNameGet

> Share SharesNameGet(ctx, name).Execute()

Get the specified share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.SharesNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.SharesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesNameGet`: Share
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.SharesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Share**](Share.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesNamePut

> SharesNamePut(ctx, name).Share(share).Execute()

Update existing share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | The name of the share to update
    share := *openapiclient.NewShare("Name_example") // Share | Share parameters to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.SharesNamePut(context.Background(), name).Share(share).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.SharesNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the share to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **share** | [**Share**](Share.md) | Share parameters to update | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesPost

> SharesPost(ctx).Share(share).Execute()

Create a new Share

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    share := *openapiclient.NewShare("Name_example") // Share | Share to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesApi.SharesPost(context.Background()).Share(share).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.SharesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSharesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **share** | [**Share**](Share.md) | Share to add | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

