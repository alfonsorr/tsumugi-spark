// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: strategies.proto

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

type AnomalyDetectionStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Strategy:
	//
	//	*AnomalyDetectionStrategy_AbsoluteChangeStrategy
	//	*AnomalyDetectionStrategy_BatchNormalStrategy
	//	*AnomalyDetectionStrategy_OnlineNormalStrategy
	//	*AnomalyDetectionStrategy_RelativeRateOfChangeStrategy
	//	*AnomalyDetectionStrategy_SimpleThresholdsStrategy
	Strategy isAnomalyDetectionStrategy_Strategy `protobuf_oneof:"strategy"`
}

func (x *AnomalyDetectionStrategy) Reset() {
	*x = AnomalyDetectionStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyDetectionStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyDetectionStrategy) ProtoMessage() {}

func (x *AnomalyDetectionStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyDetectionStrategy.ProtoReflect.Descriptor instead.
func (*AnomalyDetectionStrategy) Descriptor() ([]byte, []int) {
	return file_strategies_proto_rawDescGZIP(), []int{0}
}

func (m *AnomalyDetectionStrategy) GetStrategy() isAnomalyDetectionStrategy_Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (x *AnomalyDetectionStrategy) GetAbsoluteChangeStrategy() *AbsoluteChangeStrategy {
	if x, ok := x.GetStrategy().(*AnomalyDetectionStrategy_AbsoluteChangeStrategy); ok {
		return x.AbsoluteChangeStrategy
	}
	return nil
}

func (x *AnomalyDetectionStrategy) GetBatchNormalStrategy() *BatchNormalStrategy {
	if x, ok := x.GetStrategy().(*AnomalyDetectionStrategy_BatchNormalStrategy); ok {
		return x.BatchNormalStrategy
	}
	return nil
}

func (x *AnomalyDetectionStrategy) GetOnlineNormalStrategy() *OnlineNormalStrategy {
	if x, ok := x.GetStrategy().(*AnomalyDetectionStrategy_OnlineNormalStrategy); ok {
		return x.OnlineNormalStrategy
	}
	return nil
}

func (x *AnomalyDetectionStrategy) GetRelativeRateOfChangeStrategy() *RelativeRateOfChangeStrategy {
	if x, ok := x.GetStrategy().(*AnomalyDetectionStrategy_RelativeRateOfChangeStrategy); ok {
		return x.RelativeRateOfChangeStrategy
	}
	return nil
}

func (x *AnomalyDetectionStrategy) GetSimpleThresholdsStrategy() *SimpleThresholdStrategy {
	if x, ok := x.GetStrategy().(*AnomalyDetectionStrategy_SimpleThresholdsStrategy); ok {
		return x.SimpleThresholdsStrategy
	}
	return nil
}

type isAnomalyDetectionStrategy_Strategy interface {
	isAnomalyDetectionStrategy_Strategy()
}

type AnomalyDetectionStrategy_AbsoluteChangeStrategy struct {
	AbsoluteChangeStrategy *AbsoluteChangeStrategy `protobuf:"bytes,1,opt,name=absolute_change_strategy,json=absoluteChangeStrategy,proto3,oneof"`
}

type AnomalyDetectionStrategy_BatchNormalStrategy struct {
	BatchNormalStrategy *BatchNormalStrategy `protobuf:"bytes,2,opt,name=batch_normal_strategy,json=batchNormalStrategy,proto3,oneof"`
}

type AnomalyDetectionStrategy_OnlineNormalStrategy struct {
	OnlineNormalStrategy *OnlineNormalStrategy `protobuf:"bytes,3,opt,name=online_normal_strategy,json=onlineNormalStrategy,proto3,oneof"`
}

type AnomalyDetectionStrategy_RelativeRateOfChangeStrategy struct {
	RelativeRateOfChangeStrategy *RelativeRateOfChangeStrategy `protobuf:"bytes,4,opt,name=relative_rate_of_change_strategy,json=relativeRateOfChangeStrategy,proto3,oneof"`
}

type AnomalyDetectionStrategy_SimpleThresholdsStrategy struct {
	SimpleThresholdsStrategy *SimpleThresholdStrategy `protobuf:"bytes,5,opt,name=simple_thresholds_strategy,json=simpleThresholdsStrategy,proto3,oneof"`
}

func (*AnomalyDetectionStrategy_AbsoluteChangeStrategy) isAnomalyDetectionStrategy_Strategy() {}

func (*AnomalyDetectionStrategy_BatchNormalStrategy) isAnomalyDetectionStrategy_Strategy() {}

func (*AnomalyDetectionStrategy_OnlineNormalStrategy) isAnomalyDetectionStrategy_Strategy() {}

func (*AnomalyDetectionStrategy_RelativeRateOfChangeStrategy) isAnomalyDetectionStrategy_Strategy() {}

func (*AnomalyDetectionStrategy_SimpleThresholdsStrategy) isAnomalyDetectionStrategy_Strategy() {}

type AbsoluteChangeStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxRateDecrease *float64 `protobuf:"fixed64,1,opt,name=max_rate_decrease,json=maxRateDecrease,proto3,oneof" json:"max_rate_decrease,omitempty"`
	MaxRateIncrease *float64 `protobuf:"fixed64,2,opt,name=max_rate_increase,json=maxRateIncrease,proto3,oneof" json:"max_rate_increase,omitempty"`
	Order           *int32   `protobuf:"varint,3,opt,name=order,proto3,oneof" json:"order,omitempty"`
}

func (x *AbsoluteChangeStrategy) Reset() {
	*x = AbsoluteChangeStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbsoluteChangeStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsoluteChangeStrategy) ProtoMessage() {}

func (x *AbsoluteChangeStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsoluteChangeStrategy.ProtoReflect.Descriptor instead.
func (*AbsoluteChangeStrategy) Descriptor() ([]byte, []int) {
	return file_strategies_proto_rawDescGZIP(), []int{1}
}

func (x *AbsoluteChangeStrategy) GetMaxRateDecrease() float64 {
	if x != nil && x.MaxRateDecrease != nil {
		return *x.MaxRateDecrease
	}
	return 0
}

func (x *AbsoluteChangeStrategy) GetMaxRateIncrease() float64 {
	if x != nil && x.MaxRateIncrease != nil {
		return *x.MaxRateIncrease
	}
	return 0
}

func (x *AbsoluteChangeStrategy) GetOrder() int32 {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return 0
}

type BatchNormalStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowerDeviationFactor *float64 `protobuf:"fixed64,1,opt,name=lower_deviation_factor,json=lowerDeviationFactor,proto3,oneof" json:"lower_deviation_factor,omitempty"`
	UpperDeviationFactor *float64 `protobuf:"fixed64,2,opt,name=upper_deviation_factor,json=upperDeviationFactor,proto3,oneof" json:"upper_deviation_factor,omitempty"`
	IncludeInterval      *bool    `protobuf:"varint,3,opt,name=include_interval,json=includeInterval,proto3,oneof" json:"include_interval,omitempty"`
}

func (x *BatchNormalStrategy) Reset() {
	*x = BatchNormalStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchNormalStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchNormalStrategy) ProtoMessage() {}

func (x *BatchNormalStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchNormalStrategy.ProtoReflect.Descriptor instead.
func (*BatchNormalStrategy) Descriptor() ([]byte, []int) {
	return file_strategies_proto_rawDescGZIP(), []int{2}
}

func (x *BatchNormalStrategy) GetLowerDeviationFactor() float64 {
	if x != nil && x.LowerDeviationFactor != nil {
		return *x.LowerDeviationFactor
	}
	return 0
}

func (x *BatchNormalStrategy) GetUpperDeviationFactor() float64 {
	if x != nil && x.UpperDeviationFactor != nil {
		return *x.UpperDeviationFactor
	}
	return 0
}

func (x *BatchNormalStrategy) GetIncludeInterval() bool {
	if x != nil && x.IncludeInterval != nil {
		return *x.IncludeInterval
	}
	return false
}

type OnlineNormalStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowerDeviationFactor  *float64 `protobuf:"fixed64,1,opt,name=lower_deviation_factor,json=lowerDeviationFactor,proto3,oneof" json:"lower_deviation_factor,omitempty"`
	UpperDeviationFactor  *float64 `protobuf:"fixed64,2,opt,name=upper_deviation_factor,json=upperDeviationFactor,proto3,oneof" json:"upper_deviation_factor,omitempty"`
	IgnoreStartPercentage *float64 `protobuf:"fixed64,3,opt,name=ignore_start_percentage,json=ignoreStartPercentage,proto3,oneof" json:"ignore_start_percentage,omitempty"`
	IgnoreAnomalies       *bool    `protobuf:"varint,4,opt,name=ignore_anomalies,json=ignoreAnomalies,proto3,oneof" json:"ignore_anomalies,omitempty"`
}

func (x *OnlineNormalStrategy) Reset() {
	*x = OnlineNormalStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineNormalStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineNormalStrategy) ProtoMessage() {}

func (x *OnlineNormalStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineNormalStrategy.ProtoReflect.Descriptor instead.
func (*OnlineNormalStrategy) Descriptor() ([]byte, []int) {
	return file_strategies_proto_rawDescGZIP(), []int{3}
}

func (x *OnlineNormalStrategy) GetLowerDeviationFactor() float64 {
	if x != nil && x.LowerDeviationFactor != nil {
		return *x.LowerDeviationFactor
	}
	return 0
}

func (x *OnlineNormalStrategy) GetUpperDeviationFactor() float64 {
	if x != nil && x.UpperDeviationFactor != nil {
		return *x.UpperDeviationFactor
	}
	return 0
}

func (x *OnlineNormalStrategy) GetIgnoreStartPercentage() float64 {
	if x != nil && x.IgnoreStartPercentage != nil {
		return *x.IgnoreStartPercentage
	}
	return 0
}

func (x *OnlineNormalStrategy) GetIgnoreAnomalies() bool {
	if x != nil && x.IgnoreAnomalies != nil {
		return *x.IgnoreAnomalies
	}
	return false
}

type RelativeRateOfChangeStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxRateDecrease *float64 `protobuf:"fixed64,1,opt,name=max_rate_decrease,json=maxRateDecrease,proto3,oneof" json:"max_rate_decrease,omitempty"`
	MaxRateIncrease *float64 `protobuf:"fixed64,2,opt,name=max_rate_increase,json=maxRateIncrease,proto3,oneof" json:"max_rate_increase,omitempty"`
	Order           *int32   `protobuf:"varint,3,opt,name=order,proto3,oneof" json:"order,omitempty"`
}

func (x *RelativeRateOfChangeStrategy) Reset() {
	*x = RelativeRateOfChangeStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelativeRateOfChangeStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelativeRateOfChangeStrategy) ProtoMessage() {}

func (x *RelativeRateOfChangeStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelativeRateOfChangeStrategy.ProtoReflect.Descriptor instead.
func (*RelativeRateOfChangeStrategy) Descriptor() ([]byte, []int) {
	return file_strategies_proto_rawDescGZIP(), []int{4}
}

func (x *RelativeRateOfChangeStrategy) GetMaxRateDecrease() float64 {
	if x != nil && x.MaxRateDecrease != nil {
		return *x.MaxRateDecrease
	}
	return 0
}

func (x *RelativeRateOfChangeStrategy) GetMaxRateIncrease() float64 {
	if x != nil && x.MaxRateIncrease != nil {
		return *x.MaxRateIncrease
	}
	return 0
}

func (x *RelativeRateOfChangeStrategy) GetOrder() int32 {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return 0
}

type SimpleThresholdStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowerBound *float64 `protobuf:"fixed64,1,opt,name=lower_bound,json=lowerBound,proto3,oneof" json:"lower_bound,omitempty"`
	UpperBound float64  `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
}

func (x *SimpleThresholdStrategy) Reset() {
	*x = SimpleThresholdStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleThresholdStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleThresholdStrategy) ProtoMessage() {}

func (x *SimpleThresholdStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_strategies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleThresholdStrategy.ProtoReflect.Descriptor instead.
func (*SimpleThresholdStrategy) Descriptor() ([]byte, []int) {
	return file_strategies_proto_rawDescGZIP(), []int{5}
}

func (x *SimpleThresholdStrategy) GetLowerBound() float64 {
	if x != nil && x.LowerBound != nil {
		return *x.LowerBound
	}
	return 0
}

func (x *SimpleThresholdStrategy) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

var File_strategies_proto protoreflect.FileDescriptor

var file_strategies_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e,
	0x6b, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04, 0x0a, 0x18, 0x41, 0x6e, 0x6f,
	0x6d, 0x61, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x68, 0x0a, 0x18, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73,
	0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x16, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x5f, 0x0a, 0x15, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x13, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x62, 0x0a, 0x16, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x14,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x12, 0x7c, 0x0a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x48, 0x00, 0x52, 0x1c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x12, 0x6d, 0x0a, 0x1a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x69,
	0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x18, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0xcb, 0x01,
	0x0a, 0x16, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x2f, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x86, 0x02, 0x0a, 0x13,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x12, 0x39, 0x0a, 0x16, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x14, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x39,
	0x0a, 0x16, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01,
	0x52, 0x14, 0x75, 0x70, 0x70, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x22, 0xe0, 0x02, 0x0a, 0x14, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x39, 0x0a,
	0x16, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52,
	0x14, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x16, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x14, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x17, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x02, 0x52, 0x15, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x2e, 0x0a, 0x10, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x6e, 0x6f, 0x6d, 0x61,
	0x6c, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x0f, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x69, 0x65, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x19, 0x0a, 0x17, 0x5f,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x69, 0x65, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x2f, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x17, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0xaf, 0x01,
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x69, 0x6e, 0x63, 0x68,
	0x65, 0x6e, 0x6b, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0f, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0d, 0x74,
	0x73, 0x75, 0x6d, 0x75, 0x67, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa0, 0x01, 0x01, 0xa2,
	0x02, 0x03, 0x43, 0x53, 0x50, 0xaa, 0x02, 0x14, 0x43, 0x6f, 0x6d, 0x2e, 0x53, 0x73, 0x69, 0x6e,
	0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x14, 0x43,
	0x6f, 0x6d, 0x5c, 0x53, 0x73, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x5c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x20, 0x43, 0x6f, 0x6d, 0x5c, 0x53, 0x73, 0x69, 0x6e, 0x63, 0x68,
	0x65, 0x6e, 0x6b, 0x6f, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x43, 0x6f, 0x6d, 0x3a, 0x3a, 0x53, 0x73,
	0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x6b, 0x6f, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strategies_proto_rawDescOnce sync.Once
	file_strategies_proto_rawDescData = file_strategies_proto_rawDesc
)

func file_strategies_proto_rawDescGZIP() []byte {
	file_strategies_proto_rawDescOnce.Do(func() {
		file_strategies_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategies_proto_rawDescData)
	})
	return file_strategies_proto_rawDescData
}

var file_strategies_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_strategies_proto_goTypes = []any{
	(*AnomalyDetectionStrategy)(nil),     // 0: com.ssinchenko.proto.AnomalyDetectionStrategy
	(*AbsoluteChangeStrategy)(nil),       // 1: com.ssinchenko.proto.AbsoluteChangeStrategy
	(*BatchNormalStrategy)(nil),          // 2: com.ssinchenko.proto.BatchNormalStrategy
	(*OnlineNormalStrategy)(nil),         // 3: com.ssinchenko.proto.OnlineNormalStrategy
	(*RelativeRateOfChangeStrategy)(nil), // 4: com.ssinchenko.proto.RelativeRateOfChangeStrategy
	(*SimpleThresholdStrategy)(nil),      // 5: com.ssinchenko.proto.SimpleThresholdStrategy
}
var file_strategies_proto_depIdxs = []int32{
	1, // 0: com.ssinchenko.proto.AnomalyDetectionStrategy.absolute_change_strategy:type_name -> com.ssinchenko.proto.AbsoluteChangeStrategy
	2, // 1: com.ssinchenko.proto.AnomalyDetectionStrategy.batch_normal_strategy:type_name -> com.ssinchenko.proto.BatchNormalStrategy
	3, // 2: com.ssinchenko.proto.AnomalyDetectionStrategy.online_normal_strategy:type_name -> com.ssinchenko.proto.OnlineNormalStrategy
	4, // 3: com.ssinchenko.proto.AnomalyDetectionStrategy.relative_rate_of_change_strategy:type_name -> com.ssinchenko.proto.RelativeRateOfChangeStrategy
	5, // 4: com.ssinchenko.proto.AnomalyDetectionStrategy.simple_thresholds_strategy:type_name -> com.ssinchenko.proto.SimpleThresholdStrategy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_strategies_proto_init() }
func file_strategies_proto_init() {
	if File_strategies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strategies_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AnomalyDetectionStrategy); i {
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
		file_strategies_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AbsoluteChangeStrategy); i {
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
		file_strategies_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BatchNormalStrategy); i {
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
		file_strategies_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*OnlineNormalStrategy); i {
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
		file_strategies_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RelativeRateOfChangeStrategy); i {
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
		file_strategies_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SimpleThresholdStrategy); i {
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
	file_strategies_proto_msgTypes[0].OneofWrappers = []any{
		(*AnomalyDetectionStrategy_AbsoluteChangeStrategy)(nil),
		(*AnomalyDetectionStrategy_BatchNormalStrategy)(nil),
		(*AnomalyDetectionStrategy_OnlineNormalStrategy)(nil),
		(*AnomalyDetectionStrategy_RelativeRateOfChangeStrategy)(nil),
		(*AnomalyDetectionStrategy_SimpleThresholdsStrategy)(nil),
	}
	file_strategies_proto_msgTypes[1].OneofWrappers = []any{}
	file_strategies_proto_msgTypes[2].OneofWrappers = []any{}
	file_strategies_proto_msgTypes[3].OneofWrappers = []any{}
	file_strategies_proto_msgTypes[4].OneofWrappers = []any{}
	file_strategies_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_strategies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_strategies_proto_goTypes,
		DependencyIndexes: file_strategies_proto_depIdxs,
		MessageInfos:      file_strategies_proto_msgTypes,
	}.Build()
	File_strategies_proto = out.File
	file_strategies_proto_rawDesc = nil
	file_strategies_proto_goTypes = nil
	file_strategies_proto_depIdxs = nil
}
