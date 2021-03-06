// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: protos/dispatchers/generic.proto

package dispatchers

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Report file
type ReportFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path and filename where the generated report is stored temporally
	ReportFilepath string `protobuf:"bytes,1,opt,name=report_filepath,json=reportFilepath,proto3" json:"report_filepath,omitempty"`
}

func (x *ReportFile) Reset() {
	*x = ReportFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dispatchers_generic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFile) ProtoMessage() {}

func (x *ReportFile) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dispatchers_generic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFile.ProtoReflect.Descriptor instead.
func (*ReportFile) Descriptor() ([]byte, []int) {
	return file_protos_dispatchers_generic_proto_rawDescGZIP(), []int{0}
}

func (x *ReportFile) GetReportFilepath() string {
	if x != nil {
		return x.ReportFilepath
	}
	return ""
}

// Email dispatcher to send the report(s) attached in an email.
// This dispatcher can be used as notifier too, in conjuntion with other dispatchers
type EmailDispatcherTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instance name, used by monitoring
	Instance string `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// Email subject
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// Email body
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Email body rendered as HTML
	IsHtml bool `protobuf:"varint,4,opt,name=is_html,json=isHtml,proto3" json:"is_html,omitempty"`
	// Recipient list
	Recipient []string `protobuf:"bytes,5,rep,name=recipient,proto3" json:"recipient,omitempty"`
	// CCs list
	Cc []string `protobuf:"bytes,6,rep,name=cc,proto3" json:"cc,omitempty"`
	// BCCs list
	Bcc []string `protobuf:"bytes,7,rep,name=bcc,proto3" json:"bcc,omitempty"`
	// List of reports to be attached in the email
	Report []*ReportFile `protobuf:"bytes,8,rep,name=report,proto3" json:"report,omitempty"`
}

func (x *EmailDispatcherTask) Reset() {
	*x = EmailDispatcherTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dispatchers_generic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailDispatcherTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailDispatcherTask) ProtoMessage() {}

func (x *EmailDispatcherTask) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dispatchers_generic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailDispatcherTask.ProtoReflect.Descriptor instead.
func (*EmailDispatcherTask) Descriptor() ([]byte, []int) {
	return file_protos_dispatchers_generic_proto_rawDescGZIP(), []int{1}
}

func (x *EmailDispatcherTask) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *EmailDispatcherTask) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailDispatcherTask) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailDispatcherTask) GetIsHtml() bool {
	if x != nil {
		return x.IsHtml
	}
	return false
}

func (x *EmailDispatcherTask) GetRecipient() []string {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *EmailDispatcherTask) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *EmailDispatcherTask) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *EmailDispatcherTask) GetReport() []*ReportFile {
	if x != nil {
		return x.Report
	}
	return nil
}

type S3DispatcherTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instance name, used by monitoring
	Instance string `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// AWS secret key
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// AWS access key
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// AWS tocken
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	// AWS region
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	// Bucket name
	Bucket string `protobuf:"bytes,6,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// List of reports to be uploaded to the S3 bucket
	Report []*ReportFile `protobuf:"bytes,7,rep,name=report,proto3" json:"report,omitempty"`
}

func (x *S3DispatcherTask) Reset() {
	*x = S3DispatcherTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dispatchers_generic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3DispatcherTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3DispatcherTask) ProtoMessage() {}

func (x *S3DispatcherTask) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dispatchers_generic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3DispatcherTask.ProtoReflect.Descriptor instead.
func (*S3DispatcherTask) Descriptor() ([]byte, []int) {
	return file_protos_dispatchers_generic_proto_rawDescGZIP(), []int{2}
}

func (x *S3DispatcherTask) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *S3DispatcherTask) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *S3DispatcherTask) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *S3DispatcherTask) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *S3DispatcherTask) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *S3DispatcherTask) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3DispatcherTask) GetReport() []*ReportFile {
	if x != nil {
		return x.Report
	}
	return nil
}

var File_protos_dispatchers_generic_proto protoreflect.FileDescriptor

var file_protos_dispatchers_generic_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22,
	0x35, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0xe9, 0x01, 0x0a, 0x13, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x68,
	0x74, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x48, 0x74, 0x6d,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x63, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x62, 0x63,
	0x63, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x10, 0x53, 0x33, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_dispatchers_generic_proto_rawDescOnce sync.Once
	file_protos_dispatchers_generic_proto_rawDescData = file_protos_dispatchers_generic_proto_rawDesc
)

func file_protos_dispatchers_generic_proto_rawDescGZIP() []byte {
	file_protos_dispatchers_generic_proto_rawDescOnce.Do(func() {
		file_protos_dispatchers_generic_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_dispatchers_generic_proto_rawDescData)
	})
	return file_protos_dispatchers_generic_proto_rawDescData
}

var file_protos_dispatchers_generic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_dispatchers_generic_proto_goTypes = []interface{}{
	(*ReportFile)(nil),          // 0: dispatchers.ReportFile
	(*EmailDispatcherTask)(nil), // 1: dispatchers.EmailDispatcherTask
	(*S3DispatcherTask)(nil),    // 2: dispatchers.S3DispatcherTask
}
var file_protos_dispatchers_generic_proto_depIdxs = []int32{
	0, // 0: dispatchers.EmailDispatcherTask.report:type_name -> dispatchers.ReportFile
	0, // 1: dispatchers.S3DispatcherTask.report:type_name -> dispatchers.ReportFile
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_dispatchers_generic_proto_init() }
func file_protos_dispatchers_generic_proto_init() {
	if File_protos_dispatchers_generic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_dispatchers_generic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_dispatchers_generic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailDispatcherTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_dispatchers_generic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3DispatcherTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_dispatchers_generic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_dispatchers_generic_proto_goTypes,
		DependencyIndexes: file_protos_dispatchers_generic_proto_depIdxs,
		MessageInfos:      file_protos_dispatchers_generic_proto_msgTypes,
	}.Build()
	File_protos_dispatchers_generic_proto = out.File
	file_protos_dispatchers_generic_proto_rawDesc = nil
	file_protos_dispatchers_generic_proto_goTypes = nil
	file_protos_dispatchers_generic_proto_depIdxs = nil
}
