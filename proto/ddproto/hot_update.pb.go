// Code generated by protoc-gen-go.
// source: hot_update.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

type VersionInfo struct {
	FileId           *int32 `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32 `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *VersionInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *VersionInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

type AssetInfo struct {
	FileId           *int32          `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32          `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	FilePath         *string         `protobuf:"bytes,3,opt,name=filePath" json:"filePath,omitempty"`
	FileSize         *int32          `protobuf:"varint,4,opt,name=fileSize" json:"fileSize,omitempty"`
	Md5              *string         `protobuf:"bytes,5,opt,name=md5" json:"md5,omitempty"`
	IsCompress       *bool           `protobuf:"varint,6,opt,name=isCompress" json:"isCompress,omitempty"`
	IsCode           *bool           `protobuf:"varint,7,opt,name=isCode" json:"isCode,omitempty"`
	GameId           *CommonEnumGame `protobuf:"varint,8,opt,name=gameId,enum=ddproto.CommonEnumGame" json:"gameId,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AssetInfo) Reset()                    { *m = AssetInfo{} }
func (m *AssetInfo) String() string            { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()               {}
func (*AssetInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *AssetInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *AssetInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

func (m *AssetInfo) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *AssetInfo) GetFileSize() int32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *AssetInfo) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *AssetInfo) GetIsCompress() bool {
	if m != nil && m.IsCompress != nil {
		return *m.IsCompress
	}
	return false
}

func (m *AssetInfo) GetIsCode() bool {
	if m != nil && m.IsCode != nil {
		return *m.IsCode
	}
	return false
}

func (m *AssetInfo) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

// 获取资源文件的最新版本信息
type HotupdateReqVersionInfo struct {
	Header              *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId         *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	DownloadGameId      []int32      `protobuf:"varint,3,rep,name=downloadGameId" json:"downloadGameId,omitempty"`
	ClientAssetsVersion *int32       `protobuf:"varint,4,opt,name=clientAssetsVersion" json:"clientAssetsVersion,omitempty"`
	ClientResolution    *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	AppVersion          *int32       `protobuf:"varint,6,opt,name=AppVersion" json:"AppVersion,omitempty"`
	NativeVersion       *int32       `protobuf:"varint,7,opt,name=nativeVersion" json:"nativeVersion,omitempty"`
	XXX_unrecognized    []byte       `json:"-"`
}

func (m *HotupdateReqVersionInfo) Reset()                    { *m = HotupdateReqVersionInfo{} }
func (m *HotupdateReqVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqVersionInfo) ProtoMessage()               {}
func (*HotupdateReqVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{2} }

func (m *HotupdateReqVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetDownloadGameId() []int32 {
	if m != nil {
		return m.DownloadGameId
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAssetsVersion() int32 {
	if m != nil && m.ClientAssetsVersion != nil {
		return *m.ClientAssetsVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetAppVersion() int32 {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetNativeVersion() int32 {
	if m != nil && m.NativeVersion != nil {
		return *m.NativeVersion
	}
	return 0
}

type HotupdateAckVersionInfo struct {
	Header               *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Version              []*VersionInfo `protobuf:"bytes,2,rep,name=version" json:"version,omitempty"`
	LastestAssetsVersion *int32         `protobuf:"varint,3,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ForceUpdate          *bool          `protobuf:"varint,4,opt,name=forceUpdate" json:"forceUpdate,omitempty"`
	ReleaseTag           *int32         `protobuf:"varint,5,opt,name=releaseTag" json:"releaseTag,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *HotupdateAckVersionInfo) Reset()                    { *m = HotupdateAckVersionInfo{} }
func (m *HotupdateAckVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckVersionInfo) ProtoMessage()               {}
func (*HotupdateAckVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{3} }

func (m *HotupdateAckVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetVersion() []*VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetForceUpdate() bool {
	if m != nil && m.ForceUpdate != nil {
		return *m.ForceUpdate
	}
	return false
}

func (m *HotupdateAckVersionInfo) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

// 获取资源文件的详细信息
type HotupdateReqAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	FileIds          []int32      `protobuf:"varint,3,rep,name=fileIds" json:"fileIds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqAssetsInfo) Reset()                    { *m = HotupdateReqAssetsInfo{} }
func (m *HotupdateReqAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{4} }

func (m *HotupdateReqAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqAssetsInfo) GetFileIds() []int32 {
	if m != nil {
		return m.FileIds
	}
	return nil
}

type HotupdateAckAssetsInfo struct {
	Header               *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost            *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets               []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	LastestAssetsVersion *int32       `protobuf:"varint,4,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ClientResolution     *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized     []byte       `json:"-"`
}

func (m *HotupdateAckAssetsInfo) Reset()                    { *m = HotupdateAckAssetsInfo{} }
func (m *HotupdateAckAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{5} }

func (m *HotupdateAckAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

// 获取某个独立游戏的下载信息
type HotupdateReqGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	ClientResolution *int32       `protobuf:"varint,4,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqGameAssetsInfo) Reset()                    { *m = HotupdateReqGameAssetsInfo{} }
func (m *HotupdateReqGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{6} }

func (m *HotupdateReqGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqGameAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

type HotupdateAckGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost        *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets           []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateAckGameAssetsInfo) Reset()                    { *m = HotupdateAckGameAssetsInfo{} }
func (m *HotupdateAckGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{7} }

func (m *HotupdateAckGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckGameAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckGameAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "ddproto.VersionInfo")
	proto.RegisterType((*AssetInfo)(nil), "ddproto.AssetInfo")
	proto.RegisterType((*HotupdateReqVersionInfo)(nil), "ddproto.hotupdate_req_versionInfo")
	proto.RegisterType((*HotupdateAckVersionInfo)(nil), "ddproto.hotupdate_ack_versionInfo")
	proto.RegisterType((*HotupdateReqAssetsInfo)(nil), "ddproto.hotupdate_req_assetsInfo")
	proto.RegisterType((*HotupdateAckAssetsInfo)(nil), "ddproto.hotupdate_ack_assetsInfo")
	proto.RegisterType((*HotupdateReqGameAssetsInfo)(nil), "ddproto.hotupdate_req_gameAssetsInfo")
	proto.RegisterType((*HotupdateAckGameAssetsInfo)(nil), "ddproto.hotupdate_ack_gameAssetsInfo")
}

var fileDescriptor17 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0x1b, 0x9a, 0xb4, 0x13, 0x51, 0xba, 0xe9, 0x82, 0xbc, 0xcb, 0x1e, 0x56, 0x11, 0x48,
	0x70, 0xe9, 0xa1, 0x12, 0x3f, 0xa0, 0x70, 0xa0, 0xbd, 0x55, 0x7c, 0x5d, 0x23, 0xab, 0x9e, 0x6e,
	0x23, 0x92, 0x38, 0xc4, 0xee, 0x22, 0xc1, 0x91, 0xbf, 0x83, 0xc4, 0x5f, 0xe1, 0x27, 0x61, 0x8f,
	0xd3, 0x92, 0x20, 0x55, 0x40, 0xd9, 0x4b, 0x55, 0xbf, 0x79, 0x9e, 0x79, 0xef, 0x8d, 0x03, 0xa3,
	0x8d, 0xd4, 0xe9, 0xb6, 0x12, 0x5c, 0xe3, 0xa4, 0xaa, 0xa5, 0x96, 0x71, 0x28, 0x04, 0xfd, 0xb9,
	0x18, 0xaf, 0x64, 0x51, 0xc8, 0x32, 0x5d, 0xe5, 0x19, 0x96, 0xda, 0x55, 0x2f, 0x4e, 0x1b, 0x10,
	0xcb, 0x6d, 0xe1, 0xa0, 0x64, 0x02, 0xd1, 0x7b, 0xac, 0x55, 0x26, 0xcb, 0x45, 0xb9, 0x96, 0xf1,
	0x10, 0x82, 0x75, 0x96, 0xe3, 0x42, 0x30, 0xef, 0xca, 0x7b, 0xda, 0x8b, 0xef, 0x43, 0x68, 0xcf,
	0x86, 0xc2, 0x4e, 0x2c, 0x90, 0x7c, 0xf3, 0x60, 0x30, 0x53, 0x0a, 0xf5, 0x5f, 0xd1, 0xe3, 0x11,
	0xf4, 0x2d, 0xb0, 0xe4, 0x7a, 0xc3, 0x7c, 0x83, 0x0c, 0x76, 0xc8, 0x9b, 0xec, 0x33, 0xb2, 0xbb,
	0xc4, 0x89, 0xc0, 0x2f, 0xc4, 0x73, 0xd6, 0xa3, 0x72, 0x0c, 0x90, 0xa9, 0x97, 0xb2, 0xa8, 0x6a,
	0x54, 0x8a, 0x05, 0x06, 0xeb, 0xdb, 0x29, 0x16, 0x13, 0xc8, 0x42, 0x3a, 0x3f, 0x83, 0xe0, 0x9a,
	0x17, 0x76, 0x6a, 0xdf, 0x9c, 0x87, 0xd3, 0xf3, 0x49, 0xe3, 0x7a, 0xd2, 0xf2, 0x97, 0x5a, 0x4a,
	0xf2, 0xc3, 0x83, 0x73, 0x13, 0x92, 0xcb, 0x28, 0xad, 0xf1, 0x63, 0x7a, 0xd3, 0x72, 0xfb, 0x18,
	0x82, 0x0d, 0x72, 0x61, 0xd4, 0x5a, 0xf9, 0xd1, 0xf4, 0x6c, 0xdf, 0x68, 0x69, 0x7f, 0xe7, 0x54,
	0x8b, 0xc7, 0x10, 0xb9, 0x14, 0x67, 0x55, 0x65, 0x66, 0x3a, 0x63, 0x0f, 0x61, 0x28, 0xe4, 0xa7,
	0x32, 0x97, 0x5c, 0xbc, 0x72, 0x5a, 0xfc, 0x2b, 0xdf, 0xe0, 0x8f, 0x60, 0xdc, 0x90, 0x6d, 0x48,
	0xaa, 0xc9, 0xb6, 0x71, 0xca, 0x60, 0xe4, 0x8a, 0xaf, 0x51, 0xc9, 0x7c, 0xab, 0x6d, 0xa5, 0x47,
	0x15, 0x63, 0xdb, 0x74, 0xdf, 0xb1, 0x03, 0xc2, 0x1e, 0xc0, 0xbd, 0x92, 0xeb, 0xec, 0x06, 0x77,
	0x70, 0x48, 0x1b, 0xf8, 0xde, 0xb1, 0xc4, 0x57, 0x1f, 0x8e, 0xb0, 0xf4, 0x04, 0xc2, 0xe6, 0x92,
	0xb1, 0xe3, 0x77, 0x68, 0xed, 0xd7, 0x70, 0x09, 0x67, 0x39, 0x57, 0x1a, 0xd5, 0x6f, 0x6e, 0x7c,
	0xd2, 0x67, 0x72, 0x59, 0xcb, 0x7a, 0x85, 0xef, 0x48, 0x09, 0x59, 0xec, 0x5b, 0x23, 0x35, 0xe6,
	0xc8, 0x15, 0xbe, 0xe5, 0xd7, 0xce, 0x5c, 0xb2, 0x06, 0xd6, 0xdd, 0x01, 0xa7, 0x6e, 0xff, 0xbb,
	0x82, 0xe6, 0xb1, 0x2d, 0x84, 0x72, 0xd9, 0xdb, 0x64, 0x58, 0x37, 0x99, 0x7f, 0x1e, 0x74, 0x0a,
	0x03, 0xba, 0x33, 0x97, 0x4a, 0xd3, 0x98, 0x41, 0x9c, 0x40, 0xe0, 0xda, 0xd0, 0x94, 0x68, 0x1a,
	0xef, 0x2f, 0xfe, 0xfa, 0x0e, 0x0e, 0x05, 0xf5, 0x87, 0xb5, 0x27, 0x5f, 0x3d, 0xb8, 0xec, 0x46,
	0x63, 0x5f, 0xed, 0xec, 0x56, 0xe2, 0x19, 0xee, 0xbf, 0x12, 0xff, 0xa0, 0x0a, 0xd2, 0x97, 0x7c,
	0x69, 0x8b, 0xb0, 0xb1, 0x1d, 0x25, 0xe2, 0xb8, 0xe8, 0x5e, 0x9c, 0xcc, 0xfd, 0xe5, 0x9d, 0xa5,
	0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x08, 0xdf, 0x01, 0xcc, 0x04, 0x00, 0x00,
}
