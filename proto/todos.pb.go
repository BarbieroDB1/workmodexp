// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: todos.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task        string                 `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CustomField string                 `protobuf:"bytes,4,opt,name=custom_field,json=customField,proto3" json:"custom_field,omitempty"`
	NewField    string                 `protobuf:"bytes,5,opt,name=new_field,json=newField,proto3" json:"new_field,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Todo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Todo) GetCustomField() string {
	if x != nil {
		return x.CustomField
	}
	return ""
}

func (x *Todo) GetNewField() string {
	if x != nil {
		return x.NewField
	}
	return ""
}

var File_todos_proto protoreflect.FileDescriptor

var file_todos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x04, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x66, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x10, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x29, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x69, 0x65, 0x72,
	0x6f, 0x64, 0x62, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6d, 0x6f, 0x64, 0x65, 0x78, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todos_proto_rawDescOnce sync.Once
	file_todos_proto_rawDescData = file_todos_proto_rawDesc
)

func file_todos_proto_rawDescGZIP() []byte {
	file_todos_proto_rawDescOnce.Do(func() {
		file_todos_proto_rawDescData = protoimpl.X.CompressGZIP(file_todos_proto_rawDescData)
	})
	return file_todos_proto_rawDescData
}

var file_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_todos_proto_goTypes = []interface{}{
	(*Todo)(nil),                  // 0: experiment.Todo
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_todos_proto_depIdxs = []int32{
	1, // 0: experiment.Todo.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: experiment.TodoService.Create:input_type -> experiment.Todo
	0, // 2: experiment.TodoService.Get:input_type -> experiment.Todo
	0, // 3: experiment.TodoService.Create:output_type -> experiment.Todo
	0, // 4: experiment.TodoService.Get:output_type -> experiment.Todo
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todos_proto_init() }
func file_todos_proto_init() {
	if File_todos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_todos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todos_proto_goTypes,
		DependencyIndexes: file_todos_proto_depIdxs,
		MessageInfos:      file_todos_proto_msgTypes,
	}.Build()
	File_todos_proto = out.File
	file_todos_proto_rawDesc = nil
	file_todos_proto_goTypes = nil
	file_todos_proto_depIdxs = nil
}
