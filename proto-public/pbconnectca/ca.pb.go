// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.8
// source: proto-public/pbconnectca/ca.proto

package pbconnectca

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type WatchRootsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// active_root_id is the ID of a root in Roots that is the active CA root.
	// Other roots are still valid if they're in the Roots list but are in the
	// process of being rotated out.
	ActiveRootId string `protobuf:"bytes,1,opt,name=active_root_id,json=activeRootId,proto3" json:"active_root_id,omitempty"`
	// trust_domain is the identification root for this Consul cluster. All
	// certificates signed by the cluster's CA must have their identifying URI
	// in this domain.
	//
	// This does not include the protocol (currently spiffe://) since we may
	// implement other protocols in future with equivalent semantics. It should
	// be compared against the "authority" section of a URI (i.e. host:port).
	TrustDomain string `protobuf:"bytes,2,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	// roots is a list of root CA certs to trust.
	Roots []*CARoot `protobuf:"bytes,3,rep,name=roots,proto3" json:"roots,omitempty"`
}

func (x *WatchRootsResponse) Reset() {
	*x = WatchRootsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRootsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRootsResponse) ProtoMessage() {}

func (x *WatchRootsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRootsResponse.ProtoReflect.Descriptor instead.
func (*WatchRootsResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_pbconnectca_ca_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRootsResponse) GetActiveRootId() string {
	if x != nil {
		return x.ActiveRootId
	}
	return ""
}

func (x *WatchRootsResponse) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

func (x *WatchRootsResponse) GetRoots() []*CARoot {
	if x != nil {
		return x.Roots
	}
	return nil
}

type CARoot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a globally unique ID (UUID) representing this CA root.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is a human-friendly name for this CA root. This value is opaque to
	// Consul and is not used for anything internally.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// serial_number is the x509 serial number of the certificate.
	SerialNumber uint64 `protobuf:"varint,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// signing_key_id is the connect.HexString encoded id of the public key that
	// corresponds to the private key used to sign leaf certificates in the
	// local datacenter.
	//
	// The value comes from x509.Certificate.SubjectKeyId of the local leaf
	// signing cert.
	//
	// See https://www.rfc-editor.org/rfc/rfc3280#section-4.2.1.1 for more detail.
	SigningKeyId string `protobuf:"bytes,4,opt,name=signing_key_id,json=signingKeyId,proto3" json:"signing_key_id,omitempty"`
	// root_cert is the PEM-encoded public certificate.
	RootCert string `protobuf:"bytes,5,opt,name=root_cert,json=rootCert,proto3" json:"root_cert,omitempty"`
	// intermediate_certs is a list of PEM-encoded intermediate certs to
	// attach to any leaf certs signed by this CA.
	IntermediateCerts []string `protobuf:"bytes,6,rep,name=intermediate_certs,json=intermediateCerts,proto3" json:"intermediate_certs,omitempty"`
	// active is true if this is the current active CA. This must only
	// be true for exactly one CA.
	Active bool `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	// rotated_out_at is the time at which this CA was removed from the state.
	// This will only be set on roots that have been rotated out from being the
	// active root.
	RotatedOutAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=rotated_out_at,json=rotatedOutAt,proto3" json:"rotated_out_at,omitempty"`
}

func (x *CARoot) Reset() {
	*x = CARoot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CARoot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CARoot) ProtoMessage() {}

func (x *CARoot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CARoot.ProtoReflect.Descriptor instead.
func (*CARoot) Descriptor() ([]byte, []int) {
	return file_proto_public_pbconnectca_ca_proto_rawDescGZIP(), []int{1}
}

func (x *CARoot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CARoot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CARoot) GetSerialNumber() uint64 {
	if x != nil {
		return x.SerialNumber
	}
	return 0
}

func (x *CARoot) GetSigningKeyId() string {
	if x != nil {
		return x.SigningKeyId
	}
	return ""
}

func (x *CARoot) GetRootCert() string {
	if x != nil {
		return x.RootCert
	}
	return ""
}

func (x *CARoot) GetIntermediateCerts() []string {
	if x != nil {
		return x.IntermediateCerts
	}
	return nil
}

func (x *CARoot) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *CARoot) GetRotatedOutAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RotatedOutAt
	}
	return nil
}

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// csr is the PEM-encoded Certificate Signing Request (CSR).
	//
	// The CSR's SAN must include a SPIFFE ID that identifies a service or agent
	// to which the ACL token provided in the `x-consul-token` metadata has write
	// access.
	Csr string `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_pbconnectca_ca_proto_rawDescGZIP(), []int{2}
}

func (x *SignRequest) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cert_pem is the PEM-encoded leaf certificate.
	CertPem string `protobuf:"bytes,2,opt,name=cert_pem,json=certPem,proto3" json:"cert_pem,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_pbconnectca_ca_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_pbconnectca_ca_proto_rawDescGZIP(), []int{3}
}

func (x *SignResponse) GetCertPem() string {
	if x != nil {
		return x.CertPem
	}
	return ""
}

var File_proto_public_pbconnectca_ca_proto protoreflect.FileDescriptor

var file_proto_public_pbconnectca_ca_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70,
	0x62, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x63, 0x61, 0x2f, 0x63, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x63, 0x61, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a,
	0x12, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x05,
	0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x63, 0x61, 0x2e, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x05,
	0x72, 0x6f, 0x6f, 0x74, 0x73, 0x22, 0x9d, 0x02, 0x0a, 0x06, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x12,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x75, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x4f, 0x75, 0x74, 0x41, 0x74, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x63, 0x73, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x70,
	0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x50, 0x65,
	0x6d, 0x32, 0x96, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x41, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x6f, 0x6f, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x63, 0x61, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f,
	0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x39, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x63, 0x61, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x63, 0x61, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_public_pbconnectca_ca_proto_rawDescOnce sync.Once
	file_proto_public_pbconnectca_ca_proto_rawDescData = file_proto_public_pbconnectca_ca_proto_rawDesc
)

func file_proto_public_pbconnectca_ca_proto_rawDescGZIP() []byte {
	file_proto_public_pbconnectca_ca_proto_rawDescOnce.Do(func() {
		file_proto_public_pbconnectca_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_public_pbconnectca_ca_proto_rawDescData)
	})
	return file_proto_public_pbconnectca_ca_proto_rawDescData
}

var file_proto_public_pbconnectca_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_public_pbconnectca_ca_proto_goTypes = []interface{}{
	(*WatchRootsResponse)(nil),    // 0: connectca.WatchRootsResponse
	(*CARoot)(nil),                // 1: connectca.CARoot
	(*SignRequest)(nil),           // 2: connectca.SignRequest
	(*SignResponse)(nil),          // 3: connectca.SignResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_proto_public_pbconnectca_ca_proto_depIdxs = []int32{
	1, // 0: connectca.WatchRootsResponse.roots:type_name -> connectca.CARoot
	4, // 1: connectca.CARoot.rotated_out_at:type_name -> google.protobuf.Timestamp
	5, // 2: connectca.ConnectCAService.WatchRoots:input_type -> google.protobuf.Empty
	2, // 3: connectca.ConnectCAService.Sign:input_type -> connectca.SignRequest
	0, // 4: connectca.ConnectCAService.WatchRoots:output_type -> connectca.WatchRootsResponse
	3, // 5: connectca.ConnectCAService.Sign:output_type -> connectca.SignResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_public_pbconnectca_ca_proto_init() }
func file_proto_public_pbconnectca_ca_proto_init() {
	if File_proto_public_pbconnectca_ca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_public_pbconnectca_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRootsResponse); i {
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
		file_proto_public_pbconnectca_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CARoot); i {
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
		file_proto_public_pbconnectca_ca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_proto_public_pbconnectca_ca_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
			RawDescriptor: file_proto_public_pbconnectca_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_public_pbconnectca_ca_proto_goTypes,
		DependencyIndexes: file_proto_public_pbconnectca_ca_proto_depIdxs,
		MessageInfos:      file_proto_public_pbconnectca_ca_proto_msgTypes,
	}.Build()
	File_proto_public_pbconnectca_ca_proto = out.File
	file_proto_public_pbconnectca_ca_proto_rawDesc = nil
	file_proto_public_pbconnectca_ca_proto_goTypes = nil
	file_proto_public_pbconnectca_ca_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectCAServiceClient is the client API for ConnectCAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectCAServiceClient interface {
	// WatchRoots provides a stream on which you can receive the list of active
	// Connect CA roots. Current roots are sent immediately at the start of the
	// stream, and new lists will be sent whenever the roots are rotated.
	WatchRoots(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ConnectCAService_WatchRootsClient, error)
	// Sign a leaf certificate for the service or agent identified by the SPIFFE
	// ID in the given CSR's SAN.
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type connectCAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectCAServiceClient(cc grpc.ClientConnInterface) ConnectCAServiceClient {
	return &connectCAServiceClient{cc}
}

func (c *connectCAServiceClient) WatchRoots(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ConnectCAService_WatchRootsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConnectCAService_serviceDesc.Streams[0], "/connectca.ConnectCAService/WatchRoots", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectCAServiceWatchRootsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConnectCAService_WatchRootsClient interface {
	Recv() (*WatchRootsResponse, error)
	grpc.ClientStream
}

type connectCAServiceWatchRootsClient struct {
	grpc.ClientStream
}

func (x *connectCAServiceWatchRootsClient) Recv() (*WatchRootsResponse, error) {
	m := new(WatchRootsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connectCAServiceClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/connectca.ConnectCAService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectCAServiceServer is the server API for ConnectCAService service.
type ConnectCAServiceServer interface {
	// WatchRoots provides a stream on which you can receive the list of active
	// Connect CA roots. Current roots are sent immediately at the start of the
	// stream, and new lists will be sent whenever the roots are rotated.
	WatchRoots(*emptypb.Empty, ConnectCAService_WatchRootsServer) error
	// Sign a leaf certificate for the service or agent identified by the SPIFFE
	// ID in the given CSR's SAN.
	Sign(context.Context, *SignRequest) (*SignResponse, error)
}

// UnimplementedConnectCAServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConnectCAServiceServer struct {
}

func (*UnimplementedConnectCAServiceServer) WatchRoots(*emptypb.Empty, ConnectCAService_WatchRootsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchRoots not implemented")
}
func (*UnimplementedConnectCAServiceServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

func RegisterConnectCAServiceServer(s *grpc.Server, srv ConnectCAServiceServer) {
	s.RegisterService(&_ConnectCAService_serviceDesc, srv)
}

func _ConnectCAService_WatchRoots_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConnectCAServiceServer).WatchRoots(m, &connectCAServiceWatchRootsServer{stream})
}

type ConnectCAService_WatchRootsServer interface {
	Send(*WatchRootsResponse) error
	grpc.ServerStream
}

type connectCAServiceWatchRootsServer struct {
	grpc.ServerStream
}

func (x *connectCAServiceWatchRootsServer) Send(m *WatchRootsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ConnectCAService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectCAServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectca.ConnectCAService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectCAServiceServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectCAService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connectca.ConnectCAService",
	HandlerType: (*ConnectCAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _ConnectCAService_Sign_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchRoots",
			Handler:       _ConnectCAService_WatchRoots_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto-public/pbconnectca/ca.proto",
}