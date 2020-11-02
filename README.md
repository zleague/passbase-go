![banner](https://passbase-sdk-banner.netlify.app/go.png)

# Go API client for passbase

# Introduction

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.2
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./passbase"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.passbase.com/verification/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IdentityApi* | [**GetIdentityById**](docs/IdentityApi.md#getidentitybyid) | **Get** /identities/{id} | Get identity
*IdentityApi* | [**GetIdentityResourceById**](docs/IdentityApi.md#getidentityresourcebyid) | **Get** /identity/{id}/resources/{resource_id} | Get resource
*IdentityApi* | [**ListIdentities**](docs/IdentityApi.md#listidentities) | **Get** /identities | List identities
*IdentityApi* | [**ListIdentityResources**](docs/IdentityApi.md#listidentityresources) | **Get** /identity/{id}/resources | List resources
*ProjectApi* | [**GetSettings**](docs/ProjectApi.md#getsettings) | **Get** /settings | Get project settings

## Documentation For Models

 - [Cursor](docs/Cursor.md)
 - [DataPoints](docs/DataPoints.md)
 - [Identity](docs/Identity.md)
 - [IdentityResource](docs/IdentityResource.md)
 - [PaginatedIdentities](docs/PaginatedIdentities.md)
 - [PaginatedResources](docs/PaginatedResources.md)
 - [ProjectSettings](docs/ProjectSettings.md)
 - [ProjectSettingsCustomizations](docs/ProjectSettingsCustomizations.md)
 - [ProjectSettingsVerificationSteps](docs/ProjectSettingsVerificationSteps.md)
 - [Resource](docs/Resource.md)
 - [ResourceFile](docs/ResourceFile.md)
 - [ResourceFilesInner](docs/ResourceFilesInner.md)
 - [ResourceFilesInputInner](docs/ResourceFilesInputInner.md)
 - [ResourceInput](docs/ResourceInput.md)
 - [User](docs/User.md)
 - [WatchlistResponse](docs/WatchlistResponse.md)

## Documentation For Authorization

## IdentityAccessToken
## PublishableApiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## SecretApiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


