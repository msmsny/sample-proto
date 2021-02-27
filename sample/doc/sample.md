# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [sample/sample.proto](#sample/sample.proto)
    - [EmptyResponse](#sample.EmptyResponse)
    - [ResultSet](#sample.ResultSet)
    - [SampleMethodRequest](#sample.SampleMethodRequest)
    - [SampleMethodResponse](#sample.SampleMethodResponse)
    - [SampleRequest](#sample.SampleRequest)
    - [SampleRequest.Record](#sample.SampleRequest.Record)
    - [SampleRequest.Record.GlacierEventData](#sample.SampleRequest.Record.GlacierEventData)
    - [SampleRequest.Record.GlacierEventData.RestoreEventData](#sample.SampleRequest.Record.GlacierEventData.RestoreEventData)
    - [SampleRequest.Record.RequestParameters](#sample.SampleRequest.Record.RequestParameters)
    - [SampleRequest.Record.ResponseElements](#sample.SampleRequest.Record.ResponseElements)
    - [SampleRequest.Record.S3](#sample.SampleRequest.Record.S3)
    - [SampleRequest.Record.S3.Bucket](#sample.SampleRequest.Record.S3.Bucket)
    - [SampleRequest.Record.S3.Bucket.OwnerIdentity](#sample.SampleRequest.Record.S3.Bucket.OwnerIdentity)
    - [SampleRequest.Record.S3.Object](#sample.SampleRequest.Record.S3.Object)
    - [SampleRequest.Record.UserIdentity](#sample.SampleRequest.Record.UserIdentity)
  
    - [Sample](#sample.Sample)
  
- [Scalar Value Types](#scalar-value-types)



<a name="sample/sample.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## sample/sample.proto



<a name="sample.EmptyResponse"></a>

### EmptyResponse







<a name="sample.ResultSet"></a>

### ResultSet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int32](#int32) |  |  |
| count | [int32](#int32) |  |  |
| records | [SampleRequest.Record](#sample.SampleRequest.Record) | repeated |  |






<a name="sample.SampleMethodRequest"></a>

### SampleMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| foo | [string](#string) |  |  |
| sample | [SampleRequest](#sample.SampleRequest) |  |  |






<a name="sample.SampleMethodResponse"></a>

### SampleMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResultSet | [ResultSet](#sample.ResultSet) |  |  |






<a name="sample.SampleRequest"></a>

### SampleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| records | [SampleRequest.Record](#sample.SampleRequest.Record) | repeated |  |






<a name="sample.SampleRequest.Record"></a>

### SampleRequest.Record



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventVersion | [string](#string) |  |  |
| eventSource | [string](#string) |  |  |
| awsRegion | [string](#string) |  |  |
| eventTime | [string](#string) |  |  |
| eventName | [string](#string) |  |  |
| userIdentity | [SampleRequest.Record.UserIdentity](#sample.SampleRequest.Record.UserIdentity) |  |  |
| requestParameters | [SampleRequest.Record.RequestParameters](#sample.SampleRequest.Record.RequestParameters) |  |  |
| responseElement | [SampleRequest.Record.ResponseElements](#sample.SampleRequest.Record.ResponseElements) |  |  |
| s3 | [SampleRequest.Record.S3](#sample.SampleRequest.Record.S3) |  |  |






<a name="sample.SampleRequest.Record.GlacierEventData"></a>

### SampleRequest.Record.GlacierEventData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| restoreEventData | [SampleRequest.Record.GlacierEventData.RestoreEventData](#sample.SampleRequest.Record.GlacierEventData.RestoreEventData) |  |  |






<a name="sample.SampleRequest.Record.GlacierEventData.RestoreEventData"></a>

### SampleRequest.Record.GlacierEventData.RestoreEventData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lifecycleRestorationExpiryTime | [string](#string) |  |  |
| lifecycleRestoreStorageClass | [string](#string) |  |  |






<a name="sample.SampleRequest.Record.RequestParameters"></a>

### SampleRequest.Record.RequestParameters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceIPAddress | [string](#string) |  |  |






<a name="sample.SampleRequest.Record.ResponseElements"></a>

### SampleRequest.Record.ResponseElements



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| x_amz_request_id | [string](#string) |  |  |
| x_amz_id_2 | [string](#string) |  |  |






<a name="sample.SampleRequest.Record.S3"></a>

### SampleRequest.Record.S3



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| s3SchemaVersion | [string](#string) |  |  |
| configurationId | [string](#string) |  |  |
| bucket | [SampleRequest.Record.S3.Bucket](#sample.SampleRequest.Record.S3.Bucket) |  |  |
| object | [SampleRequest.Record.S3.Object](#sample.SampleRequest.Record.S3.Object) |  |  |






<a name="sample.SampleRequest.Record.S3.Bucket"></a>

### SampleRequest.Record.S3.Bucket



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| ownerIdentity | [SampleRequest.Record.S3.Bucket.OwnerIdentity](#sample.SampleRequest.Record.S3.Bucket.OwnerIdentity) |  |  |
| arn | [string](#string) |  |  |






<a name="sample.SampleRequest.Record.S3.Bucket.OwnerIdentity"></a>

### SampleRequest.Record.S3.Bucket.OwnerIdentity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| principalId | [string](#string) |  |  |






<a name="sample.SampleRequest.Record.S3.Object"></a>

### SampleRequest.Record.S3.Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| size | [int64](#int64) |  | empty if DELETE |
| eTag | [string](#string) |  | empty if DELETE |
| versionId | [string](#string) |  | empty if not versioning |
| sequencer | [string](#string) |  |  |






<a name="sample.SampleRequest.Record.UserIdentity"></a>

### SampleRequest.Record.UserIdentity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| principalId | [string](#string) |  |  |





 

 

 


<a name="sample.Sample"></a>

### Sample


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SampleMethod | [SampleMethodRequest](#sample.SampleMethodRequest) | [SampleMethodResponse](#sample.SampleMethodResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

