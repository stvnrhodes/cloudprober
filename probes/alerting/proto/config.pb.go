// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/probes/alerting/proto/config.proto

package proto

import (
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

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Command to run when alert is fired. In the command line following fields
	// are substituted:
	//  @alert@: Alert name
	//  @probe@: Probe name
	//  @target@: Target name, or target and port if port is specified.
	//  @target.label.<label>@: Label <label> value, e.g. target.label.role.
	//  @value@: Value that triggered the alert.
	//  @threshold@: Threshold that was crossed.
	//  @since@: Time since the alert condition started.
	//
	// For example, if you want to send an email when an alert is fired, you can
	// use the following command:
	// command: "/usr/bin/mail -s 'Alert @alert@ fired for @target@' manu@a.b"
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Notify) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type AlertConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the alert. Default is to use the probe name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Thresholds for the alert.
	FailureThreshold float32 `protobuf:"fixed32,2,opt,name=failure_threshold,json=failureThreshold,proto3" json:"failure_threshold,omitempty"`
	// Duration threshold in seconds. If duration_threshold_sec is set, alert
	// will be fired only if alert condition is true for
	// duration_threshold_sec.
	DurationThresholdSec *int32 `protobuf:"varint,3,opt,name=duration_threshold_sec,json=durationThresholdSec,proto3,oneof" json:"duration_threshold_sec,omitempty"`
	// How to notify in case of alert.
	Notify *Notify `protobuf:"bytes,4,opt,name=notify,proto3" json:"notify,omitempty"`
}

func (x *AlertConf) Reset() {
	*x = AlertConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertConf) ProtoMessage() {}

func (x *AlertConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertConf.ProtoReflect.Descriptor instead.
func (*AlertConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *AlertConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertConf) GetFailureThreshold() float32 {
	if x != nil {
		return x.FailureThreshold
	}
	return 0
}

func (x *AlertConf) GetDurationThresholdSec() int32 {
	if x != nil && x.DurationThresholdSec != nil {
		return *x.DurationThresholdSec
	}
	return 0
}

func (x *AlertConf) GetNotify() *Notify {
	if x != nil {
		return x.Notify
	}
	return nil
}

var File_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x22, 0x22, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x10, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x39, 0x0a, 0x16, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x14, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x65, 0x63, 0x88, 0x01, 0x01,
	0x12, 0x39, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x42, 0x19, 0x0a, 0x17, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_goTypes = []interface{}{
	(*Notify)(nil),    // 0: cloudprober.probes.alerts.Notify
	(*AlertConf)(nil), // 1: cloudprober.probes.alerts.AlertConf
}
var file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_depIdxs = []int32{
	0, // 0: cloudprober.probes.alerts.AlertConf.notify:type_name -> cloudprober.probes.alerts.Notify
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
		file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertConf); i {
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
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_probes_alerting_proto_config_proto_depIdxs = nil
}
