// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: detection.proto

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

type DetectionObject int32

const (
	DetectionObject_Person DetectionObject = 0
)

// Enum value maps for DetectionObject.
var (
	DetectionObject_name = map[int32]string{
		0: "Person",
	}
	DetectionObject_value = map[string]int32{
		"Person": 0,
	}
)

func (x DetectionObject) Enum() *DetectionObject {
	p := new(DetectionObject)
	*p = x
	return p
}

func (x DetectionObject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DetectionObject) Descriptor() protoreflect.EnumDescriptor {
	return file_detection_proto_enumTypes[0].Descriptor()
}

func (DetectionObject) Type() protoreflect.EnumType {
	return &file_detection_proto_enumTypes[0]
}

func (x DetectionObject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DetectionObject.Descriptor instead.
func (DetectionObject) EnumDescriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{0}
}

type DetectErrorCode int32

const (
	DetectErrorCode_BadRequest          DetectErrorCode = 0
	DetectErrorCode_MaxRequestsExceeded DetectErrorCode = 1
	DetectErrorCode_MaxImagesExceeded   DetectErrorCode = 2
	DetectErrorCode_InternalError       DetectErrorCode = 10
	DetectErrorCode_Unexpected          DetectErrorCode = 99
)

// Enum value maps for DetectErrorCode.
var (
	DetectErrorCode_name = map[int32]string{
		0:  "BadRequest",
		1:  "MaxRequestsExceeded",
		2:  "MaxImagesExceeded",
		10: "InternalError",
		99: "Unexpected",
	}
	DetectErrorCode_value = map[string]int32{
		"BadRequest":          0,
		"MaxRequestsExceeded": 1,
		"MaxImagesExceeded":   2,
		"InternalError":       10,
		"Unexpected":          99,
	}
)

func (x DetectErrorCode) Enum() *DetectErrorCode {
	p := new(DetectErrorCode)
	*p = x
	return p
}

func (x DetectErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DetectErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_detection_proto_enumTypes[1].Descriptor()
}

func (DetectErrorCode) Type() protoreflect.EnumType {
	return &file_detection_proto_enumTypes[1]
}

func (x DetectErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DetectErrorCode.Descriptor instead.
func (DetectErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{1}
}

type ImageDetections_ErrorCode int32

const (
	ImageDetections_InvalidImage    ImageDetections_ErrorCode = 0
	ImageDetections_MaxSizeExceeded ImageDetections_ErrorCode = 1
	ImageDetections_InternalError   ImageDetections_ErrorCode = 10
	ImageDetections_Unexpected      ImageDetections_ErrorCode = 99
)

// Enum value maps for ImageDetections_ErrorCode.
var (
	ImageDetections_ErrorCode_name = map[int32]string{
		0:  "InvalidImage",
		1:  "MaxSizeExceeded",
		10: "InternalError",
		99: "Unexpected",
	}
	ImageDetections_ErrorCode_value = map[string]int32{
		"InvalidImage":    0,
		"MaxSizeExceeded": 1,
		"InternalError":   10,
		"Unexpected":      99,
	}
)

func (x ImageDetections_ErrorCode) Enum() *ImageDetections_ErrorCode {
	p := new(ImageDetections_ErrorCode)
	*p = x
	return p
}

func (x ImageDetections_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageDetections_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_detection_proto_enumTypes[2].Descriptor()
}

func (ImageDetections_ErrorCode) Type() protoreflect.EnumType {
	return &file_detection_proto_enumTypes[2]
}

func (x ImageDetections_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageDetections_ErrorCode.Descriptor instead.
func (ImageDetections_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{3, 0}
}

type DetectImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *DetectImagesRequest) Reset() {
	*x = DetectImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectImagesRequest) ProtoMessage() {}

func (x *DetectImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectImagesRequest.ProtoReflect.Descriptor instead.
func (*DetectImagesRequest) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{0}
}

func (x *DetectImagesRequest) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type DetectImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detections []*ImageDetections `protobuf:"bytes,1,rep,name=detections,proto3" json:"detections,omitempty"`
	Error      *DetectError       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DetectImagesResponse) Reset() {
	*x = DetectImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectImagesResponse) ProtoMessage() {}

func (x *DetectImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectImagesResponse.ProtoReflect.Descriptor instead.
func (*DetectImagesResponse) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{1}
}

func (x *DetectImagesResponse) GetDetections() []*ImageDetections {
	if x != nil {
		return x.Detections
	}
	return nil
}

func (x *DetectImagesResponse) GetError() *DetectError {
	if x != nil {
		return x.Error
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{2}
}

func (x *Image) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type ImageDetections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detections []*Detection           `protobuf:"bytes,1,rep,name=detections,proto3" json:"detections,omitempty"`
	Error      *ImageDetections_Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ImageDetections) Reset() {
	*x = ImageDetections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageDetections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDetections) ProtoMessage() {}

func (x *ImageDetections) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDetections.ProtoReflect.Descriptor instead.
func (*ImageDetections) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{3}
}

func (x *ImageDetections) GetDetections() []*Detection {
	if x != nil {
		return x.Detections
	}
	return nil
}

func (x *ImageDetections) GetError() *ImageDetections_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Detection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object DetectionObject   `protobuf:"varint,1,opt,name=object,proto3,enum=detection.DetectionObject" json:"object,omitempty"`
	Score  float32           `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Coords *Detection_Coords `protobuf:"bytes,3,opt,name=coords,proto3" json:"coords,omitempty"`
}

func (x *Detection) Reset() {
	*x = Detection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detection) ProtoMessage() {}

func (x *Detection) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detection.ProtoReflect.Descriptor instead.
func (*Detection) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{4}
}

func (x *Detection) GetObject() DetectionObject {
	if x != nil {
		return x.Object
	}
	return DetectionObject_Person
}

func (x *Detection) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Detection) GetCoords() *Detection_Coords {
	if x != nil {
		return x.Coords
	}
	return nil
}

type DetectError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    DetectErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=detection.DetectErrorCode" json:"code,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DetectError) Reset() {
	*x = DetectError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectError) ProtoMessage() {}

func (x *DetectError) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectError.ProtoReflect.Descriptor instead.
func (*DetectError) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{5}
}

func (x *DetectError) GetCode() DetectErrorCode {
	if x != nil {
		return x.Code
	}
	return DetectErrorCode_BadRequest
}

func (x *DetectError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ImageDetections_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ImageDetections_ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=detection.ImageDetections_ErrorCode" json:"code,omitempty"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ImageDetections_Error) Reset() {
	*x = ImageDetections_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageDetections_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDetections_Error) ProtoMessage() {}

func (x *ImageDetections_Error) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDetections_Error.ProtoReflect.Descriptor instead.
func (*ImageDetections_Error) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ImageDetections_Error) GetCode() ImageDetections_ErrorCode {
	if x != nil {
		return x.Code
	}
	return ImageDetections_InvalidImage
}

func (x *ImageDetections_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Detection box coordinates in pixels
type Detection_Coords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X      float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y      float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Width  float32 `protobuf:"fixed32,3,opt,name=width,proto3" json:"width,omitempty"`
	Height float32 `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Detection_Coords) Reset() {
	*x = Detection_Coords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detection_Coords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detection_Coords) ProtoMessage() {}

func (x *Detection_Coords) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detection_Coords.ProtoReflect.Descriptor instead.
func (*Detection_Coords) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Detection_Coords) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Detection_Coords) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Detection_Coords) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Detection_Coords) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_detection_proto protoreflect.FileDescriptor

var file_detection_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x13,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x80, 0x01,
	0x0a, 0x14, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x21, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x5b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x55, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x45, 0x78, 0x63, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x63, 0x22, 0xde, 0x01, 0x0a, 0x09, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x06,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x52, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x57, 0x0a, 0x0b, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x1d, 0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x10, 0x00, 0x2a, 0x74, 0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x4d, 0x61, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x78, 0x63, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x63, 0x32, 0x65, 0x0a, 0x10, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_detection_proto_rawDescOnce sync.Once
	file_detection_proto_rawDescData = file_detection_proto_rawDesc
)

func file_detection_proto_rawDescGZIP() []byte {
	file_detection_proto_rawDescOnce.Do(func() {
		file_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_detection_proto_rawDescData)
	})
	return file_detection_proto_rawDescData
}

var file_detection_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_detection_proto_goTypes = []any{
	(DetectionObject)(0),           // 0: detection.DetectionObject
	(DetectErrorCode)(0),           // 1: detection.DetectErrorCode
	(ImageDetections_ErrorCode)(0), // 2: detection.ImageDetections.ErrorCode
	(*DetectImagesRequest)(nil),    // 3: detection.DetectImagesRequest
	(*DetectImagesResponse)(nil),   // 4: detection.DetectImagesResponse
	(*Image)(nil),                  // 5: detection.Image
	(*ImageDetections)(nil),        // 6: detection.ImageDetections
	(*Detection)(nil),              // 7: detection.Detection
	(*DetectError)(nil),            // 8: detection.DetectError
	(*ImageDetections_Error)(nil),  // 9: detection.ImageDetections.Error
	(*Detection_Coords)(nil),       // 10: detection.Detection.Coords
}
var file_detection_proto_depIdxs = []int32{
	5,  // 0: detection.DetectImagesRequest.images:type_name -> detection.Image
	6,  // 1: detection.DetectImagesResponse.detections:type_name -> detection.ImageDetections
	8,  // 2: detection.DetectImagesResponse.error:type_name -> detection.DetectError
	7,  // 3: detection.ImageDetections.detections:type_name -> detection.Detection
	9,  // 4: detection.ImageDetections.error:type_name -> detection.ImageDetections.Error
	0,  // 5: detection.Detection.object:type_name -> detection.DetectionObject
	10, // 6: detection.Detection.coords:type_name -> detection.Detection.Coords
	1,  // 7: detection.DetectError.code:type_name -> detection.DetectErrorCode
	2,  // 8: detection.ImageDetections.Error.code:type_name -> detection.ImageDetections.ErrorCode
	3,  // 9: detection.DetectionService.DetectImages:input_type -> detection.DetectImagesRequest
	4,  // 10: detection.DetectionService.DetectImages:output_type -> detection.DetectImagesResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_detection_proto_init() }
func file_detection_proto_init() {
	if File_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detection_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DetectImagesRequest); i {
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
		file_detection_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DetectImagesResponse); i {
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
		file_detection_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Image); i {
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
		file_detection_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ImageDetections); i {
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
		file_detection_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Detection); i {
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
		file_detection_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DetectError); i {
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
		file_detection_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ImageDetections_Error); i {
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
		file_detection_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Detection_Coords); i {
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
			RawDescriptor: file_detection_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_detection_proto_goTypes,
		DependencyIndexes: file_detection_proto_depIdxs,
		EnumInfos:         file_detection_proto_enumTypes,
		MessageInfos:      file_detection_proto_msgTypes,
	}.Build()
	File_detection_proto = out.File
	file_detection_proto_rawDesc = nil
	file_detection_proto_goTypes = nil
	file_detection_proto_depIdxs = nil
}
