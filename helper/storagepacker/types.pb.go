// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: helper/storagepacker/types.proto

package storagepacker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Item represents an entry that gets inserted into the storage packer
type Item struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID must be provided by the caller; the same value, if used with GetItem,
	// can be used to fetch the item. However, when iterating through a bucket,
	// this ID will be an internal ID. In other words, outside of the use-case
	// described above, the caller *must not* rely on this value to be
	// consistent with what they passed in.
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// message is the contents of the item
	Message       *anypb.Any `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_helper_storagepacker_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_helper_storagepacker_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_helper_storagepacker_types_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Item) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

// Bucket is a construct to hold multiple items within itself. This
// abstraction contains multiple buckets of the same kind within itself and
// shares amont them the items that get inserted. When the bucket as a whole
// gets too big to hold more items, the contained buckets gets pushed out only
// to become independent buckets. Hence, this can grow infinitely in terms of
// storage space for items that get inserted.
type Bucket struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Key is the storage path where the bucket gets stored
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Items holds the items contained within this bucket. Used by v1.
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// ItemMap stores a mapping of item ID to message. Used by v2.
	ItemMap       map[string]*anypb.Any `protobuf:"bytes,3,rep,name=item_map,json=itemMap,proto3" json:"item_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	mi := &file_helper_storagepacker_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_helper_storagepacker_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_helper_storagepacker_types_proto_rawDescGZIP(), []int{1}
}

func (x *Bucket) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Bucket) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Bucket) GetItemMap() map[string]*anypb.Any {
	if x != nil {
		return x.ItemMap
	}
	return nil
}

var File_helper_storagepacker_types_proto protoreflect.FileDescriptor

const file_helper_storagepacker_types_proto_rawDesc = "" +
	"\n" +
	" helper/storagepacker/types.proto\x12\rstoragepacker\x1a\x19google/protobuf/any.proto\"F\n" +
	"\x04Item\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12.\n" +
	"\amessage\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\amessage\"\xd6\x01\n" +
	"\x06Bucket\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12)\n" +
	"\x05items\x18\x02 \x03(\v2\x13.storagepacker.ItemR\x05items\x12=\n" +
	"\bitem_map\x18\x03 \x03(\v2\".storagepacker.Bucket.ItemMapEntryR\aitemMap\x1aP\n" +
	"\fItemMapEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05value:\x028\x01B1Z/github.com/openbao/openbao/helper/storagepackerb\x06proto3"

var (
	file_helper_storagepacker_types_proto_rawDescOnce sync.Once
	file_helper_storagepacker_types_proto_rawDescData []byte
)

func file_helper_storagepacker_types_proto_rawDescGZIP() []byte {
	file_helper_storagepacker_types_proto_rawDescOnce.Do(func() {
		file_helper_storagepacker_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_helper_storagepacker_types_proto_rawDesc), len(file_helper_storagepacker_types_proto_rawDesc)))
	})
	return file_helper_storagepacker_types_proto_rawDescData
}

var file_helper_storagepacker_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_helper_storagepacker_types_proto_goTypes = []any{
	(*Item)(nil),      // 0: storagepacker.Item
	(*Bucket)(nil),    // 1: storagepacker.Bucket
	nil,               // 2: storagepacker.Bucket.ItemMapEntry
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_helper_storagepacker_types_proto_depIDxs = []int32{
	3, // 0: storagepacker.Item.message:type_name -> google.protobuf.Any
	0, // 1: storagepacker.Bucket.items:type_name -> storagepacker.Item
	2, // 2: storagepacker.Bucket.item_map:type_name -> storagepacker.Bucket.ItemMapEntry
	3, // 3: storagepacker.Bucket.ItemMapEntry.value:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_helper_storagepacker_types_proto_init() }
func file_helper_storagepacker_types_proto_init() {
	if File_helper_storagepacker_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_helper_storagepacker_types_proto_rawDesc), len(file_helper_storagepacker_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helper_storagepacker_types_proto_goTypes,
		DependencyIndexes: file_helper_storagepacker_types_proto_depIDxs,
		MessageInfos:      file_helper_storagepacker_types_proto_msgTypes,
	}.Build()
	File_helper_storagepacker_types_proto = out.File
	file_helper_storagepacker_types_proto_goTypes = nil
	file_helper_storagepacker_types_proto_depIDxs = nil
}
