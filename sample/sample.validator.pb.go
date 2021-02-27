// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sample/sample.proto

package sample

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SampleMethodRequest) Validate() error {
	if this.Foo == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Foo", fmt.Errorf(`value '%v' must not be an empty string`, this.Foo))
	}
	if nil == this.Sample {
		return github_com_mwitkow_go_proto_validators.FieldError("Sample", fmt.Errorf("message must exist"))
	}
	if this.Sample != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sample); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sample", err)
		}
	}
	return nil
}
func (this *SampleRequest) Validate() error {
	if len(this.Records) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("Records", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Records))
	}
	for _, item := range this.Records {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Records", err)
			}
		}
	}
	return nil
}
func (this *SampleRequest_Record) Validate() error {
	if this.UserIdentity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserIdentity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserIdentity", err)
		}
	}
	if this.RequestParameters != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RequestParameters); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RequestParameters", err)
		}
	}
	if this.ResponseElement != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResponseElement); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResponseElement", err)
		}
	}
	if this.S3 != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.S3); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("S3", err)
		}
	}
	return nil
}
func (this *SampleRequest_Record_UserIdentity) Validate() error {
	return nil
}
func (this *SampleRequest_Record_RequestParameters) Validate() error {
	return nil
}
func (this *SampleRequest_Record_ResponseElements) Validate() error {
	return nil
}
func (this *SampleRequest_Record_S3) Validate() error {
	if this.Bucket != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Bucket); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Bucket", err)
		}
	}
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *SampleRequest_Record_S3_Bucket) Validate() error {
	if this.OwnerIdentity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OwnerIdentity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OwnerIdentity", err)
		}
	}
	return nil
}
func (this *SampleRequest_Record_S3_Bucket_OwnerIdentity) Validate() error {
	return nil
}
func (this *SampleRequest_Record_S3_Object) Validate() error {
	return nil
}
func (this *SampleRequest_Record_GlacierEventData) Validate() error {
	if this.RestoreEventData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RestoreEventData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RestoreEventData", err)
		}
	}
	return nil
}
func (this *SampleRequest_Record_GlacierEventData_RestoreEventData) Validate() error {
	return nil
}
func (this *SampleMethodResponse) Validate() error {
	if this.ResultSet != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResultSet); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResultSet", err)
		}
	}
	return nil
}
func (this *ResultSet) Validate() error {
	for _, item := range this.Records {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Records", err)
			}
		}
	}
	return nil
}
func (this *EmptyResponse) Validate() error {
	return nil
}
