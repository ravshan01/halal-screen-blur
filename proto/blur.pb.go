// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: blur.proto

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

type BlurErrorCode int32

const (
	BlurErrorCode_BadRequest          BlurErrorCode = 0
	BlurErrorCode_MaxRequestsExceeded BlurErrorCode = 1
	BlurErrorCode_MaxImagesExceeded   BlurErrorCode = 2
	BlurErrorCode_InternalError       BlurErrorCode = 10
	BlurErrorCode_Unexpected          BlurErrorCode = 99
)

// Enum value maps for BlurErrorCode.
var (
	BlurErrorCode_name = map[int32]string{
		0:  "BadRequest",
		1:  "MaxRequestsExceeded",
		2:  "MaxImagesExceeded",
		10: "InternalError",
		99: "Unexpected",
	}
	BlurErrorCode_value = map[string]int32{
		"BadRequest":          0,
		"MaxRequestsExceeded": 1,
		"MaxImagesExceeded":   2,
		"InternalError":       10,
		"Unexpected":          99,
	}
)

func (x BlurErrorCode) Enum() *BlurErrorCode {
	p := new(BlurErrorCode)
	*p = x
	return p
}

func (x BlurErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlurErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_blur_proto_enumTypes[0].Descriptor()
}

func (BlurErrorCode) Type() protoreflect.EnumType {
	return &file_blur_proto_enumTypes[0]
}

func (x BlurErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlurErrorCode.Descriptor instead.
func (BlurErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{0}
}

type BlurredImage_ErrorCode int32

const (
	BlurredImage_InvalidContent    BlurredImage_ErrorCode = 0
	BlurredImage_InvalidPercentage BlurredImage_ErrorCode = 1
	BlurredImage_InternalError     BlurredImage_ErrorCode = 2
	BlurredImage_Unexpected        BlurredImage_ErrorCode = 99
)

// Enum value maps for BlurredImage_ErrorCode.
var (
	BlurredImage_ErrorCode_name = map[int32]string{
		0:  "InvalidContent",
		1:  "InvalidPercentage",
		2:  "InternalError",
		99: "Unexpected",
	}
	BlurredImage_ErrorCode_value = map[string]int32{
		"InvalidContent":    0,
		"InvalidPercentage": 1,
		"InternalError":     2,
		"Unexpected":        99,
	}
)

func (x BlurredImage_ErrorCode) Enum() *BlurredImage_ErrorCode {
	p := new(BlurredImage_ErrorCode)
	*p = x
	return p
}

func (x BlurredImage_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlurredImage_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_blur_proto_enumTypes[1].Descriptor()
}

func (BlurredImage_ErrorCode) Type() protoreflect.EnumType {
	return &file_blur_proto_enumTypes[1]
}

func (x BlurredImage_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlurredImage_ErrorCode.Descriptor instead.
func (BlurredImage_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{3, 0}
}

type BlurImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*ImageForBlur `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *BlurImagesRequest) Reset() {
	*x = BlurImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blur_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlurImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlurImagesRequest) ProtoMessage() {}

func (x *BlurImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blur_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlurImagesRequest.ProtoReflect.Descriptor instead.
func (*BlurImagesRequest) Descriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{0}
}

func (x *BlurImagesRequest) GetImages() []*ImageForBlur {
	if x != nil {
		return x.Images
	}
	return nil
}

type BlurImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlurredImages []*BlurredImage `protobuf:"bytes,1,rep,name=blurred_images,json=blurredImages,proto3" json:"blurred_images,omitempty"`
	Error         *BlurError      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BlurImagesResponse) Reset() {
	*x = BlurImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blur_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlurImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlurImagesResponse) ProtoMessage() {}

func (x *BlurImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blur_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlurImagesResponse.ProtoReflect.Descriptor instead.
func (*BlurImagesResponse) Descriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{1}
}

func (x *BlurImagesResponse) GetBlurredImages() []*BlurredImage {
	if x != nil {
		return x.BlurredImages
	}
	return nil
}

func (x *BlurImagesResponse) GetError() *BlurError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ImageForBlur struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte              `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Coords  []*Detection_Coords `protobuf:"bytes,2,rep,name=coords,proto3" json:"coords,omitempty"`
	// between 0 and 100
	Percentage uint32 `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *ImageForBlur) Reset() {
	*x = ImageForBlur{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blur_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageForBlur) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageForBlur) ProtoMessage() {}

func (x *ImageForBlur) ProtoReflect() protoreflect.Message {
	mi := &file_blur_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageForBlur.ProtoReflect.Descriptor instead.
func (*ImageForBlur) Descriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{2}
}

func (x *ImageForBlur) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ImageForBlur) GetCoords() []*Detection_Coords {
	if x != nil {
		return x.Coords
	}
	return nil
}

func (x *ImageForBlur) GetPercentage() uint32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

type BlurredImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte              `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Error   *BlurredImage_Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BlurredImage) Reset() {
	*x = BlurredImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blur_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlurredImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlurredImage) ProtoMessage() {}

func (x *BlurredImage) ProtoReflect() protoreflect.Message {
	mi := &file_blur_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlurredImage.ProtoReflect.Descriptor instead.
func (*BlurredImage) Descriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{3}
}

func (x *BlurredImage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *BlurredImage) GetError() *BlurredImage_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type BlurError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    BlurErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=blur.BlurErrorCode" json:"code,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BlurError) Reset() {
	*x = BlurError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blur_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlurError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlurError) ProtoMessage() {}

func (x *BlurError) ProtoReflect() protoreflect.Message {
	mi := &file_blur_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlurError.ProtoReflect.Descriptor instead.
func (*BlurError) Descriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{4}
}

func (x *BlurError) GetCode() BlurErrorCode {
	if x != nil {
		return x.Code
	}
	return BlurErrorCode_BadRequest
}

func (x *BlurError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BlurredImage_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    BlurredImage_ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=blur.BlurredImage_ErrorCode" json:"code,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BlurredImage_Error) Reset() {
	*x = BlurredImage_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blur_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlurredImage_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlurredImage_Error) ProtoMessage() {}

func (x *BlurredImage_Error) ProtoReflect() protoreflect.Message {
	mi := &file_blur_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlurredImage_Error.ProtoReflect.Descriptor instead.
func (*BlurredImage_Error) Descriptor() ([]byte, []int) {
	return file_blur_proto_rawDescGZIP(), []int{3, 0}
}

func (x *BlurredImage_Error) GetCode() BlurredImage_ErrorCode {
	if x != nil {
		return x.Code
	}
	return BlurredImage_InvalidContent
}

func (x *BlurredImage_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_blur_proto protoreflect.FileDescriptor

var file_blur_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c,
	0x75, 0x72, 0x1a, 0x0f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x42, 0x6c, 0x75, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x75, 0x72, 0x52, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x12, 0x42, 0x6c, 0x75, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x62, 0x6c,
	0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42, 0x6c, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x62, 0x6c, 0x75, 0x72, 0x72, 0x65, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42, 0x6c, 0x75, 0x72,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7d, 0x0a, 0x0c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x42, 0x6c, 0x75, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x0c,
	0x42, 0x6c, 0x75, 0x72, 0x72, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42, 0x6c, 0x75,
	0x72, 0x72, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x53, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x30, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42, 0x6c, 0x75, 0x72, 0x72, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x10, 0x63, 0x22, 0x4e, 0x0a, 0x09, 0x42, 0x6c, 0x75, 0x72, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42, 0x6c, 0x75, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x72, 0x0a, 0x0d, 0x42, 0x6c, 0x75, 0x72, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x61, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x61, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x78, 0x63,
	0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x63, 0x32, 0x50, 0x0a, 0x0b, 0x42, 0x6c,
	0x75, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x42, 0x6c, 0x75,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42,
	0x6c, 0x75, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x42, 0x6c, 0x75, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blur_proto_rawDescOnce sync.Once
	file_blur_proto_rawDescData = file_blur_proto_rawDesc
)

func file_blur_proto_rawDescGZIP() []byte {
	file_blur_proto_rawDescOnce.Do(func() {
		file_blur_proto_rawDescData = protoimpl.X.CompressGZIP(file_blur_proto_rawDescData)
	})
	return file_blur_proto_rawDescData
}

var file_blur_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_blur_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_blur_proto_goTypes = []any{
	(BlurErrorCode)(0),          // 0: blur.BlurErrorCode
	(BlurredImage_ErrorCode)(0), // 1: blur.BlurredImage.ErrorCode
	(*BlurImagesRequest)(nil),   // 2: blur.BlurImagesRequest
	(*BlurImagesResponse)(nil),  // 3: blur.BlurImagesResponse
	(*ImageForBlur)(nil),        // 4: blur.ImageForBlur
	(*BlurredImage)(nil),        // 5: blur.BlurredImage
	(*BlurError)(nil),           // 6: blur.BlurError
	(*BlurredImage_Error)(nil),  // 7: blur.BlurredImage.Error
	(*Detection_Coords)(nil),    // 8: detection.Detection.Coords
}
var file_blur_proto_depIdxs = []int32{
	4, // 0: blur.BlurImagesRequest.images:type_name -> blur.ImageForBlur
	5, // 1: blur.BlurImagesResponse.blurred_images:type_name -> blur.BlurredImage
	6, // 2: blur.BlurImagesResponse.error:type_name -> blur.BlurError
	8, // 3: blur.ImageForBlur.coords:type_name -> detection.Detection.Coords
	7, // 4: blur.BlurredImage.error:type_name -> blur.BlurredImage.Error
	0, // 5: blur.BlurError.code:type_name -> blur.BlurErrorCode
	1, // 6: blur.BlurredImage.Error.code:type_name -> blur.BlurredImage.ErrorCode
	2, // 7: blur.BlurService.BlurImages:input_type -> blur.BlurImagesRequest
	3, // 8: blur.BlurService.BlurImages:output_type -> blur.BlurImagesResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_blur_proto_init() }
func file_blur_proto_init() {
	if File_blur_proto != nil {
		return
	}
	file_detection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blur_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BlurImagesRequest); i {
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
		file_blur_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BlurImagesResponse); i {
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
		file_blur_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ImageForBlur); i {
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
		file_blur_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BlurredImage); i {
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
		file_blur_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BlurError); i {
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
		file_blur_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BlurredImage_Error); i {
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
			RawDescriptor: file_blur_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blur_proto_goTypes,
		DependencyIndexes: file_blur_proto_depIdxs,
		EnumInfos:         file_blur_proto_enumTypes,
		MessageInfos:      file_blur_proto_msgTypes,
	}.Build()
	File_blur_proto = out.File
	file_blur_proto_rawDesc = nil
	file_blur_proto_goTypes = nil
	file_blur_proto_depIdxs = nil
}
