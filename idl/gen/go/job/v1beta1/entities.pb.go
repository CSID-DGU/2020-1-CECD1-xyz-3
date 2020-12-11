// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: job/v1beta1/entities.proto

package v1beta1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CompanySummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CompanySummary) Reset() {
	*x = CompanySummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1beta1_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanySummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanySummary) ProtoMessage() {}

func (x *CompanySummary) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1beta1_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanySummary.ProtoReflect.Descriptor instead.
func (*CompanySummary) Descriptor() ([]byte, []int) {
	return file_job_v1beta1_entities_proto_rawDescGZIP(), []int{0}
}

func (x *CompanySummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompanySummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JobSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company *CompanySummary `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *JobSummary) Reset() {
	*x = JobSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1beta1_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSummary) ProtoMessage() {}

func (x *JobSummary) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1beta1_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSummary.ProtoReflect.Descriptor instead.
func (*JobSummary) Descriptor() ([]byte, []int) {
	return file_job_v1beta1_entities_proto_rawDescGZIP(), []int{1}
}

func (x *JobSummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobSummary) GetCompany() *CompanySummary {
	if x != nil {
		return x.Company
	}
	return nil
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company     *CompanySummary      `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	CreateTime  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	OpenTime    *timestamp.Timestamp `protobuf:"bytes,5,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=close_time,json=closeTime,proto3,oneof" json:"close_time,omitempty"`
	Description *JobDescription      `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1beta1_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1beta1_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_job_v1beta1_entities_proto_rawDescGZIP(), []int{2}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetCompany() *CompanySummary {
	if x != nil {
		return x.Company
	}
	return nil
}

func (x *Job) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Job) GetOpenTime() *timestamp.Timestamp {
	if x != nil {
		return x.OpenTime
	}
	return nil
}

func (x *Job) GetCloseTime() *timestamp.Timestamp {
	if x != nil {
		return x.CloseTime
	}
	return nil
}

func (x *Job) GetDescription() *JobDescription {
	if x != nil {
		return x.Description
	}
	return nil
}

type JobDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location    string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *JobDescription) Reset() {
	*x = JobDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1beta1_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDescription) ProtoMessage() {}

func (x *JobDescription) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1beta1_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDescription.ProtoReflect.Descriptor instead.
func (*JobDescription) Descriptor() ([]byte, []int) {
	return file_job_v1beta1_entities_proto_rawDescGZIP(), []int{3}
}

func (x *JobDescription) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *JobDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_job_v1beta1_entities_proto protoreflect.FileDescriptor

var file_job_v1beta1_entities_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6a, 0x6f,
	0x62, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x67, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0xe4, 0x02, 0x0a, 0x03, 0x4a, 0x6f,
	0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x3b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x4e, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6f, 0x6f, 0x67, 0x65, 0x75, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_v1beta1_entities_proto_rawDescOnce sync.Once
	file_job_v1beta1_entities_proto_rawDescData = file_job_v1beta1_entities_proto_rawDesc
)

func file_job_v1beta1_entities_proto_rawDescGZIP() []byte {
	file_job_v1beta1_entities_proto_rawDescOnce.Do(func() {
		file_job_v1beta1_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_v1beta1_entities_proto_rawDescData)
	})
	return file_job_v1beta1_entities_proto_rawDescData
}

var file_job_v1beta1_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_job_v1beta1_entities_proto_goTypes = []interface{}{
	(*CompanySummary)(nil),      // 0: job.v1beta1.CompanySummary
	(*JobSummary)(nil),          // 1: job.v1beta1.JobSummary
	(*Job)(nil),                 // 2: job.v1beta1.Job
	(*JobDescription)(nil),      // 3: job.v1beta1.JobDescription
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_job_v1beta1_entities_proto_depIdxs = []int32{
	0, // 0: job.v1beta1.JobSummary.company:type_name -> job.v1beta1.CompanySummary
	0, // 1: job.v1beta1.Job.company:type_name -> job.v1beta1.CompanySummary
	4, // 2: job.v1beta1.Job.create_time:type_name -> google.protobuf.Timestamp
	4, // 3: job.v1beta1.Job.open_time:type_name -> google.protobuf.Timestamp
	4, // 4: job.v1beta1.Job.close_time:type_name -> google.protobuf.Timestamp
	3, // 5: job.v1beta1.Job.description:type_name -> job.v1beta1.JobDescription
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_job_v1beta1_entities_proto_init() }
func file_job_v1beta1_entities_proto_init() {
	if File_job_v1beta1_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_v1beta1_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanySummary); i {
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
		file_job_v1beta1_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSummary); i {
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
		file_job_v1beta1_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_job_v1beta1_entities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDescription); i {
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
	file_job_v1beta1_entities_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_v1beta1_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_v1beta1_entities_proto_goTypes,
		DependencyIndexes: file_job_v1beta1_entities_proto_depIdxs,
		MessageInfos:      file_job_v1beta1_entities_proto_msgTypes,
	}.Build()
	File_job_v1beta1_entities_proto = out.File
	file_job_v1beta1_entities_proto_rawDesc = nil
	file_job_v1beta1_entities_proto_goTypes = nil
	file_job_v1beta1_entities_proto_depIdxs = nil
}
