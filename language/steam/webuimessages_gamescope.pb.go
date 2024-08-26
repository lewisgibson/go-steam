// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: webuimessages_gamescope.proto

package steam

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

type CMsgDisplayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Make                  *string `protobuf:"bytes,1,opt,name=make" json:"make,omitempty"`
	Model                 *string `protobuf:"bytes,2,opt,name=model" json:"model,omitempty"`
	ConnectorName         *string `protobuf:"bytes,3,opt,name=connector_name,json=connectorName" json:"connector_name,omitempty"`
	SupportedRefreshRates []int32 `protobuf:"varint,4,rep,name=supported_refresh_rates,json=supportedRefreshRates" json:"supported_refresh_rates,omitempty"`
	SupportedFrameRates   []int32 `protobuf:"varint,5,rep,name=supported_frame_rates,json=supportedFrameRates" json:"supported_frame_rates,omitempty"`
	IsExternal            *bool   `protobuf:"varint,6,opt,name=is_external,json=isExternal" json:"is_external,omitempty"`
	IsHdrCapable          *bool   `protobuf:"varint,7,opt,name=is_hdr_capable,json=isHdrCapable" json:"is_hdr_capable,omitempty"`
	IsVrrCapable          *bool   `protobuf:"varint,8,opt,name=is_vrr_capable,json=isVrrCapable" json:"is_vrr_capable,omitempty"`
}

func (x *CMsgDisplayInfo) Reset() {
	*x = CMsgDisplayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgDisplayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgDisplayInfo) ProtoMessage() {}

func (x *CMsgDisplayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgDisplayInfo.ProtoReflect.Descriptor instead.
func (*CMsgDisplayInfo) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgDisplayInfo) GetMake() string {
	if x != nil && x.Make != nil {
		return *x.Make
	}
	return ""
}

func (x *CMsgDisplayInfo) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *CMsgDisplayInfo) GetConnectorName() string {
	if x != nil && x.ConnectorName != nil {
		return *x.ConnectorName
	}
	return ""
}

func (x *CMsgDisplayInfo) GetSupportedRefreshRates() []int32 {
	if x != nil {
		return x.SupportedRefreshRates
	}
	return nil
}

func (x *CMsgDisplayInfo) GetSupportedFrameRates() []int32 {
	if x != nil {
		return x.SupportedFrameRates
	}
	return nil
}

func (x *CMsgDisplayInfo) GetIsExternal() bool {
	if x != nil && x.IsExternal != nil {
		return *x.IsExternal
	}
	return false
}

func (x *CMsgDisplayInfo) GetIsHdrCapable() bool {
	if x != nil && x.IsHdrCapable != nil {
		return *x.IsHdrCapable
	}
	return false
}

func (x *CMsgDisplayInfo) GetIsVrrCapable() bool {
	if x != nil && x.IsVrrCapable != nil {
		return *x.IsVrrCapable
	}
	return false
}

type CMsgGamescopeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsServiceAvailable               *bool            `protobuf:"varint,1,opt,name=is_service_available,json=isServiceAvailable" json:"is_service_available,omitempty"`
	IsReshadeSupported               *bool            `protobuf:"varint,2,opt,name=is_reshade_supported,json=isReshadeSupported" json:"is_reshade_supported,omitempty"`
	IsAppHdrEnabled                  *bool            `protobuf:"varint,3,opt,name=is_app_hdr_enabled,json=isAppHdrEnabled" json:"is_app_hdr_enabled,omitempty"`
	IsAppRefreshRateSupported        *bool            `protobuf:"varint,4,opt,name=is_app_refresh_rate_supported,json=isAppRefreshRateSupported" json:"is_app_refresh_rate_supported,omitempty"`
	ActiveDisplayInfo                *CMsgDisplayInfo `protobuf:"bytes,5,opt,name=active_display_info,json=activeDisplayInfo" json:"active_display_info,omitempty"`
	IsAppRefreshRateCapable          *bool            `protobuf:"varint,6,opt,name=is_app_refresh_rate_capable,json=isAppRefreshRateCapable" json:"is_app_refresh_rate_capable,omitempty"`
	IsRefreshRateSwitchingSupported  *bool            `protobuf:"varint,7,opt,name=is_refresh_rate_switching_supported,json=isRefreshRateSwitchingSupported" json:"is_refresh_rate_switching_supported,omitempty"`
	IsRefreshRateSwitchingRestricted *bool            `protobuf:"varint,8,opt,name=is_refresh_rate_switching_restricted,json=isRefreshRateSwitchingRestricted" json:"is_refresh_rate_switching_restricted,omitempty"`
	IsHdrVisualizationSupported      *bool            `protobuf:"varint,9,opt,name=is_hdr_visualization_supported,json=isHdrVisualizationSupported" json:"is_hdr_visualization_supported,omitempty"`
	IsMuraCorrectionSupported        *bool            `protobuf:"varint,10,opt,name=is_mura_correction_supported,json=isMuraCorrectionSupported" json:"is_mura_correction_supported,omitempty"`
}

func (x *CMsgGamescopeState) Reset() {
	*x = CMsgGamescopeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgGamescopeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgGamescopeState) ProtoMessage() {}

func (x *CMsgGamescopeState) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgGamescopeState.ProtoReflect.Descriptor instead.
func (*CMsgGamescopeState) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgGamescopeState) GetIsServiceAvailable() bool {
	if x != nil && x.IsServiceAvailable != nil {
		return *x.IsServiceAvailable
	}
	return false
}

func (x *CMsgGamescopeState) GetIsReshadeSupported() bool {
	if x != nil && x.IsReshadeSupported != nil {
		return *x.IsReshadeSupported
	}
	return false
}

func (x *CMsgGamescopeState) GetIsAppHdrEnabled() bool {
	if x != nil && x.IsAppHdrEnabled != nil {
		return *x.IsAppHdrEnabled
	}
	return false
}

func (x *CMsgGamescopeState) GetIsAppRefreshRateSupported() bool {
	if x != nil && x.IsAppRefreshRateSupported != nil {
		return *x.IsAppRefreshRateSupported
	}
	return false
}

func (x *CMsgGamescopeState) GetActiveDisplayInfo() *CMsgDisplayInfo {
	if x != nil {
		return x.ActiveDisplayInfo
	}
	return nil
}

func (x *CMsgGamescopeState) GetIsAppRefreshRateCapable() bool {
	if x != nil && x.IsAppRefreshRateCapable != nil {
		return *x.IsAppRefreshRateCapable
	}
	return false
}

func (x *CMsgGamescopeState) GetIsRefreshRateSwitchingSupported() bool {
	if x != nil && x.IsRefreshRateSwitchingSupported != nil {
		return *x.IsRefreshRateSwitchingSupported
	}
	return false
}

func (x *CMsgGamescopeState) GetIsRefreshRateSwitchingRestricted() bool {
	if x != nil && x.IsRefreshRateSwitchingRestricted != nil {
		return *x.IsRefreshRateSwitchingRestricted
	}
	return false
}

func (x *CMsgGamescopeState) GetIsHdrVisualizationSupported() bool {
	if x != nil && x.IsHdrVisualizationSupported != nil {
		return *x.IsHdrVisualizationSupported
	}
	return false
}

func (x *CMsgGamescopeState) GetIsMuraCorrectionSupported() bool {
	if x != nil && x.IsMuraCorrectionSupported != nil {
		return *x.IsMuraCorrectionSupported
	}
	return false
}

type CGamescope_GetState_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CGamescope_GetState_Request) Reset() {
	*x = CGamescope_GetState_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGamescope_GetState_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGamescope_GetState_Request) ProtoMessage() {}

func (x *CGamescope_GetState_Request) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGamescope_GetState_Request.ProtoReflect.Descriptor instead.
func (*CGamescope_GetState_Request) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{2}
}

type CGamescope_GetState_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *CMsgGamescopeState `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
}

func (x *CGamescope_GetState_Response) Reset() {
	*x = CGamescope_GetState_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGamescope_GetState_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGamescope_GetState_Response) ProtoMessage() {}

func (x *CGamescope_GetState_Response) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGamescope_GetState_Response.ProtoReflect.Descriptor instead.
func (*CGamescope_GetState_Response) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{3}
}

func (x *CGamescope_GetState_Response) GetState() *CMsgGamescopeState {
	if x != nil {
		return x.State
	}
	return nil
}

type CGamescope_StateChanged_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CGamescope_StateChanged_Notification) Reset() {
	*x = CGamescope_StateChanged_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGamescope_StateChanged_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGamescope_StateChanged_Notification) ProtoMessage() {}

func (x *CGamescope_StateChanged_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGamescope_StateChanged_Notification.ProtoReflect.Descriptor instead.
func (*CGamescope_StateChanged_Notification) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{4}
}

type CGamescope_SetBlurParams_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode           *EGamescopeBlurMode `protobuf:"varint,1,opt,name=mode,enum=EGamescopeBlurMode,def=0" json:"mode,omitempty"`
	Radius         *int32              `protobuf:"varint,2,opt,name=radius" json:"radius,omitempty"`
	FadeDurationMs *int32              `protobuf:"varint,3,opt,name=fade_duration_ms,json=fadeDurationMs" json:"fade_duration_ms,omitempty"`
}

// Default values for CGamescope_SetBlurParams_Request fields.
const (
	Default_CGamescope_SetBlurParams_Request_Mode = EGamescopeBlurMode_k_EGamescopeBlurMode_Disabled
)

func (x *CGamescope_SetBlurParams_Request) Reset() {
	*x = CGamescope_SetBlurParams_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGamescope_SetBlurParams_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGamescope_SetBlurParams_Request) ProtoMessage() {}

func (x *CGamescope_SetBlurParams_Request) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGamescope_SetBlurParams_Request.ProtoReflect.Descriptor instead.
func (*CGamescope_SetBlurParams_Request) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{5}
}

func (x *CGamescope_SetBlurParams_Request) GetMode() EGamescopeBlurMode {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return Default_CGamescope_SetBlurParams_Request_Mode
}

func (x *CGamescope_SetBlurParams_Request) GetRadius() int32 {
	if x != nil && x.Radius != nil {
		return *x.Radius
	}
	return 0
}

func (x *CGamescope_SetBlurParams_Request) GetFadeDurationMs() int32 {
	if x != nil && x.FadeDurationMs != nil {
		return *x.FadeDurationMs
	}
	return 0
}

type CGamescope_SetBlurParams_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CGamescope_SetBlurParams_Response) Reset() {
	*x = CGamescope_SetBlurParams_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webuimessages_gamescope_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGamescope_SetBlurParams_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGamescope_SetBlurParams_Response) ProtoMessage() {}

func (x *CGamescope_SetBlurParams_Response) ProtoReflect() protoreflect.Message {
	mi := &file_webuimessages_gamescope_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGamescope_SetBlurParams_Response.ProtoReflect.Descriptor instead.
func (*CGamescope_SetBlurParams_Response) Descriptor() ([]byte, []int) {
	return file_webuimessages_gamescope_proto_rawDescGZIP(), []int{6}
}

var File_webuimessages_gamescope_proto protoreflect.FileDescriptor

var file_webuimessages_gamescope_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x77, 0x65, 0x62, 0x75, 0x69, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x77, 0x65, 0x62, 0x75, 0x69, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbb, 0x02, 0x0a, 0x0f, 0x43, 0x4d, 0x73, 0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x15, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x13, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x68, 0x64, 0x72, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x48, 0x64,
	0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x76,
	0x72, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x73, 0x56, 0x72, 0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x8b,
	0x05, 0x0a, 0x12, 0x43, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x68, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x52, 0x65, 0x73, 0x68, 0x61, 0x64, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f,
	0x61, 0x70, 0x70, 0x5f, 0x68, 0x64, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x41, 0x70, 0x70, 0x48, 0x64, 0x72, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x1d, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x70,
	0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69,
	0x73, 0x41, 0x70, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x1b, 0x69, 0x73,
	0x5f, 0x61, 0x70, 0x70, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x17, 0x69, 0x73, 0x41, 0x70, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x23, 0x69, 0x73, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x69, 0x73, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x52, 0x61, 0x74, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x24, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x20, 0x69, 0x73, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x61, 0x74, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x1e, 0x69, 0x73, 0x5f, 0x68, 0x64, 0x72,
	0x5f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b,
	0x69, 0x73, 0x48, 0x64, 0x72, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x69,
	0x73, 0x5f, 0x6d, 0x75, 0x72, 0x61, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x19, 0x69, 0x73, 0x4d, 0x75, 0x72, 0x61, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x22, 0x1d, 0x0a, 0x1b,
	0x43, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x1c, 0x43,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x4d, 0x73,
	0x67, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x24, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xac,
	0x01, 0x0a, 0x20, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x53, 0x65,
	0x74, 0x42, 0x6c, 0x75, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x45, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x6c,
	0x75, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x1d, 0x6b, 0x5f, 0x45, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x42, 0x6c, 0x75, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x64,
	0x69, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66,
	0x61, 0x64, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x22, 0x23, 0x0a,
	0x21, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x53, 0x65, 0x74, 0x42,
	0x6c, 0x75, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x81, 0x02, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x43,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x43, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12,
	0x25, 0x2e, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x57, 0x65, 0x62, 0x55, 0x49, 0x4e, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x42,
	0x6c, 0x75, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x2e, 0x43, 0x47, 0x61, 0x6d,
	0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x75, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x43,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x75,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x04, 0x80, 0x97, 0x22, 0x01, 0x42, 0x05, 0x48, 0x01, 0x80, 0x01, 0x01,
}

var (
	file_webuimessages_gamescope_proto_rawDescOnce sync.Once
	file_webuimessages_gamescope_proto_rawDescData = file_webuimessages_gamescope_proto_rawDesc
)

func file_webuimessages_gamescope_proto_rawDescGZIP() []byte {
	file_webuimessages_gamescope_proto_rawDescOnce.Do(func() {
		file_webuimessages_gamescope_proto_rawDescData = protoimpl.X.CompressGZIP(file_webuimessages_gamescope_proto_rawDescData)
	})
	return file_webuimessages_gamescope_proto_rawDescData
}

var file_webuimessages_gamescope_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_webuimessages_gamescope_proto_goTypes = []any{
	(*CMsgDisplayInfo)(nil),                      // 0: CMsgDisplayInfo
	(*CMsgGamescopeState)(nil),                   // 1: CMsgGamescopeState
	(*CGamescope_GetState_Request)(nil),          // 2: CGamescope_GetState_Request
	(*CGamescope_GetState_Response)(nil),         // 3: CGamescope_GetState_Response
	(*CGamescope_StateChanged_Notification)(nil), // 4: CGamescope_StateChanged_Notification
	(*CGamescope_SetBlurParams_Request)(nil),     // 5: CGamescope_SetBlurParams_Request
	(*CGamescope_SetBlurParams_Response)(nil),    // 6: CGamescope_SetBlurParams_Response
	(EGamescopeBlurMode)(0),                      // 7: EGamescopeBlurMode
	(*WebUINoResponse)(nil),                      // 8: WebUINoResponse
}
var file_webuimessages_gamescope_proto_depIdxs = []int32{
	0, // 0: CMsgGamescopeState.active_display_info:type_name -> CMsgDisplayInfo
	1, // 1: CGamescope_GetState_Response.state:type_name -> CMsgGamescopeState
	7, // 2: CGamescope_SetBlurParams_Request.mode:type_name -> EGamescopeBlurMode
	2, // 3: Gamescope.GetState:input_type -> CGamescope_GetState_Request
	4, // 4: Gamescope.NotifyStateChanged:input_type -> CGamescope_StateChanged_Notification
	5, // 5: Gamescope.SetBlurParams:input_type -> CGamescope_SetBlurParams_Request
	3, // 6: Gamescope.GetState:output_type -> CGamescope_GetState_Response
	8, // 7: Gamescope.NotifyStateChanged:output_type -> WebUINoResponse
	6, // 8: Gamescope.SetBlurParams:output_type -> CGamescope_SetBlurParams_Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_webuimessages_gamescope_proto_init() }
func file_webuimessages_gamescope_proto_init() {
	if File_webuimessages_gamescope_proto != nil {
		return
	}
	file_enums_proto_init()
	file_steammessages_base_proto_init()
	file_webuimessages_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_webuimessages_gamescope_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgDisplayInfo); i {
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
		file_webuimessages_gamescope_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgGamescopeState); i {
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
		file_webuimessages_gamescope_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CGamescope_GetState_Request); i {
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
		file_webuimessages_gamescope_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CGamescope_GetState_Response); i {
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
		file_webuimessages_gamescope_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CGamescope_StateChanged_Notification); i {
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
		file_webuimessages_gamescope_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CGamescope_SetBlurParams_Request); i {
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
		file_webuimessages_gamescope_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CGamescope_SetBlurParams_Response); i {
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
			RawDescriptor: file_webuimessages_gamescope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webuimessages_gamescope_proto_goTypes,
		DependencyIndexes: file_webuimessages_gamescope_proto_depIdxs,
		MessageInfos:      file_webuimessages_gamescope_proto_msgTypes,
	}.Build()
	File_webuimessages_gamescope_proto = out.File
	file_webuimessages_gamescope_proto_rawDesc = nil
	file_webuimessages_gamescope_proto_goTypes = nil
	file_webuimessages_gamescope_proto_depIdxs = nil
}
