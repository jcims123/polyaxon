# PolyaxonSdk.AuthV1Api

Polyaxon sdk

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**changePassword**](AuthV1Api.md#changePassword) | **POST** /api/v1/auth/change-password | Change password
[**login**](AuthV1Api.md#login) | **POST** /api/v1/auth/token | Login
[**resetPassword**](AuthV1Api.md#resetPassword) | **POST** /api/v1/auth/reset-password | Reset password
[**signup**](AuthV1Api.md#signup) | **POST** /api/v1/auth/signup | Signup



## changePassword

> changePassword(body)

Change password

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.AuthV1Api();
let body = new PolyaxonSdk.V1PasswordChange(); // V1PasswordChange | 
apiInstance.changePassword(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1PasswordChange**](V1PasswordChange.md)|  | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## login

> V1Auth login(body)

Login

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.AuthV1Api();
let body = new PolyaxonSdk.V1Credentials(); // V1Credentials | 
apiInstance.login(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1Credentials**](V1Credentials.md)|  | 

### Return type

[**V1Auth**](V1Auth.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## resetPassword

> resetPassword(body)

Reset password

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.AuthV1Api();
let body = new PolyaxonSdk.V1UserEmail(); // V1UserEmail | 
apiInstance.resetPassword(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1UserEmail**](V1UserEmail.md)|  | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## signup

> signup(body)

Signup

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.AuthV1Api();
let body = new PolyaxonSdk.V1UserSingup(); // V1UserSingup | 
apiInstance.signup(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1UserSingup**](V1UserSingup.md)|  | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

