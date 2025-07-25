// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: tick/tick.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Tick struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbol        string                 `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Ts            int64                  `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	StartTime     int64                  `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	TradeId       int64                  `protobuf:"varint,4,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	Quantity      string                 `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         string                 `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tick) Reset() {
	*x = Tick{}
	mi := &file_tick_tick_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tick) ProtoMessage() {}

func (x *Tick) ProtoReflect() protoreflect.Message {
	mi := &file_tick_tick_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tick.ProtoReflect.Descriptor instead.
func (*Tick) Descriptor() ([]byte, []int) {
	return file_tick_tick_proto_rawDescGZIP(), []int{0}
}

func (x *Tick) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Tick) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *Tick) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Tick) GetTradeId() int64 {
	if x != nil {
		return x.TradeId
	}
	return 0
}

func (x *Tick) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Tick) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type StreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbols       []string               `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Interval      string                 `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	mi := &file_tick_tick_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tick_tick_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_tick_tick_proto_rawDescGZIP(), []int{1}
}

func (x *StreamRequest) GetSymbols() []string {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *StreamRequest) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

var File_tick_tick_proto protoreflect.FileDescriptor

const file_tick_tick_proto_rawDesc = "" +
	"\n" +
	"\x0ftick/tick.proto\x12\x04tick\"\x9a\x01\n" +
	"\x04Tick\x12\x16\n" +
	"\x06symbol\x18\x01 \x01(\tR\x06symbol\x12\x0e\n" +
	"\x02ts\x18\x02 \x01(\x03R\x02ts\x12\x1d\n" +
	"\n" +
	"start_time\x18\x03 \x01(\x03R\tstartTime\x12\x19\n" +
	"\btrade_id\x18\x04 \x01(\x03R\atradeId\x12\x1a\n" +
	"\bquantity\x18\x05 \x01(\tR\bquantity\x12\x14\n" +
	"\x05price\x18\x06 \x01(\tR\x05price\"E\n" +
	"\rStreamRequest\x12\x18\n" +
	"\asymbols\x18\x01 \x03(\tR\asymbols\x12\x1a\n" +
	"\binterval\x18\x02 \x01(\tR\binterval2C\n" +
	"\x10StreamingService\x12/\n" +
	"\n" +
	"StreamTick\x12\x13.tick.StreamRequest\x1a\n" +
	".tick.Tick0\x01BHZFgithub.com/acom21/chart-streaming-service/service/stream/proto/tick;pbb\x06proto3"

var (
	file_tick_tick_proto_rawDescOnce sync.Once
	file_tick_tick_proto_rawDescData []byte
)

func file_tick_tick_proto_rawDescGZIP() []byte {
	file_tick_tick_proto_rawDescOnce.Do(func() {
		file_tick_tick_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tick_tick_proto_rawDesc), len(file_tick_tick_proto_rawDesc)))
	})
	return file_tick_tick_proto_rawDescData
}

var file_tick_tick_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tick_tick_proto_goTypes = []any{
	(*Tick)(nil),          // 0: tick.Tick
	(*StreamRequest)(nil), // 1: tick.StreamRequest
}
var file_tick_tick_proto_depIdxs = []int32{
	1, // 0: tick.StreamingService.StreamTick:input_type -> tick.StreamRequest
	0, // 1: tick.StreamingService.StreamTick:output_type -> tick.Tick
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tick_tick_proto_init() }
func file_tick_tick_proto_init() {
	if File_tick_tick_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tick_tick_proto_rawDesc), len(file_tick_tick_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tick_tick_proto_goTypes,
		DependencyIndexes: file_tick_tick_proto_depIdxs,
		MessageInfos:      file_tick_tick_proto_msgTypes,
	}.Build()
	File_tick_tick_proto = out.File
	file_tick_tick_proto_goTypes = nil
	file_tick_tick_proto_depIdxs = nil
}
