//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: ctl/storage.proto

package ctl

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

type StorageScanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nvme *ScanNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm  *ScanScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
}

func (x *StorageScanReq) Reset() {
	*x = StorageScanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageScanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageScanReq) ProtoMessage() {}

func (x *StorageScanReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageScanReq.ProtoReflect.Descriptor instead.
func (*StorageScanReq) Descriptor() ([]byte, []int) {
	return file_ctl_storage_proto_rawDescGZIP(), []int{0}
}

func (x *StorageScanReq) GetNvme() *ScanNvmeReq {
	if x != nil {
		return x.Nvme
	}
	return nil
}

func (x *StorageScanReq) GetScm() *ScanScmReq {
	if x != nil {
		return x.Scm
	}
	return nil
}

type StorageScanResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nvme *ScanNvmeResp `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm  *ScanScmResp  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
}

func (x *StorageScanResp) Reset() {
	*x = StorageScanResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageScanResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageScanResp) ProtoMessage() {}

func (x *StorageScanResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageScanResp.ProtoReflect.Descriptor instead.
func (*StorageScanResp) Descriptor() ([]byte, []int) {
	return file_ctl_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StorageScanResp) GetNvme() *ScanNvmeResp {
	if x != nil {
		return x.Nvme
	}
	return nil
}

func (x *StorageScanResp) GetScm() *ScanScmResp {
	if x != nil {
		return x.Scm
	}
	return nil
}

type StorageFormatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nvme     *FormatNvmeReq `protobuf:"bytes,1,opt,name=nvme,proto3" json:"nvme,omitempty"`
	Scm      *FormatScmReq  `protobuf:"bytes,2,opt,name=scm,proto3" json:"scm,omitempty"`
	Reformat bool           `protobuf:"varint,3,opt,name=reformat,proto3" json:"reformat,omitempty"`
}

func (x *StorageFormatReq) Reset() {
	*x = StorageFormatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageFormatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageFormatReq) ProtoMessage() {}

func (x *StorageFormatReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageFormatReq.ProtoReflect.Descriptor instead.
func (*StorageFormatReq) Descriptor() ([]byte, []int) {
	return file_ctl_storage_proto_rawDescGZIP(), []int{2}
}

func (x *StorageFormatReq) GetNvme() *FormatNvmeReq {
	if x != nil {
		return x.Nvme
	}
	return nil
}

func (x *StorageFormatReq) GetScm() *FormatScmReq {
	if x != nil {
		return x.Scm
	}
	return nil
}

func (x *StorageFormatReq) GetReformat() bool {
	if x != nil {
		return x.Reformat
	}
	return false
}

type StorageFormatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crets []*NvmeControllerResult `protobuf:"bytes,1,rep,name=crets,proto3" json:"crets,omitempty"` // One per controller format attempt
	Mrets []*ScmMountResult       `protobuf:"bytes,2,rep,name=mrets,proto3" json:"mrets,omitempty"` // One per scm format and mount attempt
}

func (x *StorageFormatResp) Reset() {
	*x = StorageFormatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageFormatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageFormatResp) ProtoMessage() {}

func (x *StorageFormatResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageFormatResp.ProtoReflect.Descriptor instead.
func (*StorageFormatResp) Descriptor() ([]byte, []int) {
	return file_ctl_storage_proto_rawDescGZIP(), []int{3}
}

func (x *StorageFormatResp) GetCrets() []*NvmeControllerResult {
	if x != nil {
		return x.Crets
	}
	return nil
}

func (x *StorageFormatResp) GetMrets() []*ScmMountResult {
	if x != nil {
		return x.Mrets
	}
	return nil
}

var File_ctl_storage_proto protoreflect.FileDescriptor

var file_ctl_storage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x74, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x74, 0x6c, 0x1a, 0x16, 0x63, 0x74, 0x6c, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x76, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x63, 0x74, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x76, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x4e, 0x76, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x6e, 0x76, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x03, 0x73, 0x63, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x74, 0x6c, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x63, 0x6d, 0x52, 0x65, 0x71, 0x52, 0x03, 0x73,
	0x63, 0x6d, 0x22, 0x5c, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x76, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x76,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x6e, 0x76, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x03,
	0x73, 0x63, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x74, 0x6c, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x53, 0x63, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x52, 0x03, 0x73, 0x63, 0x6d,
	0x22, 0x7b, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x76, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4e,
	0x76, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x6e, 0x76, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x03,
	0x73, 0x63, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x74, 0x6c, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x63, 0x6d, 0x52, 0x65, 0x71, 0x52, 0x03, 0x73, 0x63,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x6f, 0x0a,
	0x11, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x4e, 0x76, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x6d, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x63, 0x6d, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x6d, 0x72, 0x65, 0x74, 0x73, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6f,
	0x73, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ctl_storage_proto_rawDescOnce sync.Once
	file_ctl_storage_proto_rawDescData = file_ctl_storage_proto_rawDesc
)

func file_ctl_storage_proto_rawDescGZIP() []byte {
	file_ctl_storage_proto_rawDescOnce.Do(func() {
		file_ctl_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_ctl_storage_proto_rawDescData)
	})
	return file_ctl_storage_proto_rawDescData
}

var file_ctl_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ctl_storage_proto_goTypes = []interface{}{
	(*StorageScanReq)(nil),       // 0: ctl.StorageScanReq
	(*StorageScanResp)(nil),      // 1: ctl.StorageScanResp
	(*StorageFormatReq)(nil),     // 2: ctl.StorageFormatReq
	(*StorageFormatResp)(nil),    // 3: ctl.StorageFormatResp
	(*ScanNvmeReq)(nil),          // 4: ctl.ScanNvmeReq
	(*ScanScmReq)(nil),           // 5: ctl.ScanScmReq
	(*ScanNvmeResp)(nil),         // 6: ctl.ScanNvmeResp
	(*ScanScmResp)(nil),          // 7: ctl.ScanScmResp
	(*FormatNvmeReq)(nil),        // 8: ctl.FormatNvmeReq
	(*FormatScmReq)(nil),         // 9: ctl.FormatScmReq
	(*NvmeControllerResult)(nil), // 10: ctl.NvmeControllerResult
	(*ScmMountResult)(nil),       // 11: ctl.ScmMountResult
}
var file_ctl_storage_proto_depIdxs = []int32{
	4,  // 0: ctl.StorageScanReq.nvme:type_name -> ctl.ScanNvmeReq
	5,  // 1: ctl.StorageScanReq.scm:type_name -> ctl.ScanScmReq
	6,  // 2: ctl.StorageScanResp.nvme:type_name -> ctl.ScanNvmeResp
	7,  // 3: ctl.StorageScanResp.scm:type_name -> ctl.ScanScmResp
	8,  // 4: ctl.StorageFormatReq.nvme:type_name -> ctl.FormatNvmeReq
	9,  // 5: ctl.StorageFormatReq.scm:type_name -> ctl.FormatScmReq
	10, // 6: ctl.StorageFormatResp.crets:type_name -> ctl.NvmeControllerResult
	11, // 7: ctl.StorageFormatResp.mrets:type_name -> ctl.ScmMountResult
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_ctl_storage_proto_init() }
func file_ctl_storage_proto_init() {
	if File_ctl_storage_proto != nil {
		return
	}
	file_ctl_storage_nvme_proto_init()
	file_ctl_storage_scm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ctl_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageScanReq); i {
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
		file_ctl_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageScanResp); i {
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
		file_ctl_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageFormatReq); i {
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
		file_ctl_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageFormatResp); i {
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
			RawDescriptor: file_ctl_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ctl_storage_proto_goTypes,
		DependencyIndexes: file_ctl_storage_proto_depIdxs,
		MessageInfos:      file_ctl_storage_proto_msgTypes,
	}.Build()
	File_ctl_storage_proto = out.File
	file_ctl_storage_proto_rawDesc = nil
	file_ctl_storage_proto_goTypes = nil
	file_ctl_storage_proto_depIdxs = nil
}