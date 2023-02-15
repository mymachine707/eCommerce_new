// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: protos/order.proto

package eCommerce

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

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Clients    *Clients `protobuf:"bytes,2,opt,name=clients,proto3" json:"clients,omitempty"`
	TotalPrice int32    `protobuf:"varint,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status     string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetClients() *Clients {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *Order) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId   string       `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Orderitems []*OrderItem `protobuf:"bytes,2,rep,name=orderitems,proto3" json:"orderitems,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateOrderRequest) GetOrderitems() []*OrderItem {
	if x != nil {
		return x.Orderitems
	}
	return nil
}

type GetOrderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetOrderListRequest) Reset() {
	*x = GetOrderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderListRequest) ProtoMessage() {}

func (x *GetOrderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderListRequest.ProtoReflect.Descriptor instead.
func (*GetOrderListRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetOrderListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetOrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrderListResponse) Reset() {
	*x = GetOrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderListResponse) ProtoMessage() {}

func (x *GetOrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderListResponse.ProtoReflect.Descriptor instead.
func (*GetOrderListResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderListResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type GetOrderByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderByIdRequest) Reset() {
	*x = GetOrderByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdRequest) ProtoMessage() {}

func (x *GetOrderByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdRequest.ProtoReflect.Descriptor instead.
func (*GetOrderByIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order      *Order                        `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Items      []*GetOrderByIdResponse_Items `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	OrderItems []*OrderItem                  `protobuf:"bytes,3,rep,name=OrderItems,proto3" json:"OrderItems,omitempty"`
}

func (x *GetOrderByIdResponse) Reset() {
	*x = GetOrderByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdResponse) ProtoMessage() {}

func (x *GetOrderByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdResponse.ProtoReflect.Descriptor instead.
func (*GetOrderByIdResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrderByIdResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *GetOrderByIdResponse) GetItems() []*GetOrderByIdResponse_Items {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetOrderByIdResponse) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order      *Order       `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Orderitems []*OrderItem `protobuf:"bytes,2,rep,name=orderitems,proto3" json:"orderitems,omitempty"`
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderInfo) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *OrderInfo) GetOrderitems() []*OrderItem {
	if x != nil {
		return x.Orderitems
	}
	return nil
}

type GetOrderByIdResponse_Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product  *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Quantity int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *GetOrderByIdResponse_Items) Reset() {
	*x = GetOrderByIdResponse_Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdResponse_Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdResponse_Items) ProtoMessage() {}

func (x *GetOrderByIdResponse_Items) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdResponse_Items.ProtoReflect.Descriptor instead.
func (*GetOrderByIdResponse_Items) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetOrderByIdResponse_Items) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *GetOrderByIdResponse_Items) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_protos_order_proto protoreflect.FileDescriptor

var file_protos_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x46, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x07, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5d, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x43, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x36, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xdc, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x47, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x55, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order_proto_rawDescOnce sync.Once
	file_protos_order_proto_rawDescData = file_protos_order_proto_rawDesc
)

func file_protos_order_proto_rawDescGZIP() []byte {
	file_protos_order_proto_rawDescOnce.Do(func() {
		file_protos_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order_proto_rawDescData)
	})
	return file_protos_order_proto_rawDescData
}

var file_protos_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_order_proto_goTypes = []interface{}{
	(*OrderItem)(nil),                  // 0: OrderItem
	(*Order)(nil),                      // 1: Order
	(*CreateOrderRequest)(nil),         // 2: CreateOrderRequest
	(*GetOrderListRequest)(nil),        // 3: GetOrderListRequest
	(*GetOrderListResponse)(nil),       // 4: GetOrderListResponse
	(*GetOrderByIdRequest)(nil),        // 5: GetOrderByIdRequest
	(*GetOrderByIdResponse)(nil),       // 6: GetOrderByIdResponse
	(*OrderInfo)(nil),                  // 7: OrderInfo
	(*GetOrderByIdResponse_Items)(nil), // 8: GetOrderByIdResponse.Items
	(*Clients)(nil),                    // 9: Clients
	(*Product)(nil),                    // 10: Product
}
var file_protos_order_proto_depIdxs = []int32{
	9,  // 0: Order.clients:type_name -> Clients
	0,  // 1: CreateOrderRequest.orderitems:type_name -> OrderItem
	1,  // 2: GetOrderListResponse.orders:type_name -> Order
	1,  // 3: GetOrderByIdResponse.order:type_name -> Order
	8,  // 4: GetOrderByIdResponse.items:type_name -> GetOrderByIdResponse.Items
	0,  // 5: GetOrderByIdResponse.OrderItems:type_name -> OrderItem
	1,  // 6: OrderInfo.order:type_name -> Order
	0,  // 7: OrderInfo.orderitems:type_name -> OrderItem
	10, // 8: GetOrderByIdResponse.Items.product:type_name -> Product
	2,  // 9: OrderService.CreateOrder:input_type -> CreateOrderRequest
	3,  // 10: OrderService.GetOrderList:input_type -> GetOrderListRequest
	5,  // 11: OrderService.GetOrderById:input_type -> GetOrderByIdRequest
	1,  // 12: OrderService.CreateOrder:output_type -> Order
	4,  // 13: OrderService.GetOrderList:output_type -> GetOrderListResponse
	6,  // 14: OrderService.GetOrderById:output_type -> GetOrderByIdResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_protos_order_proto_init() }
func file_protos_order_proto_init() {
	if File_protos_order_proto != nil {
		return
	}
	file_protos_client_proto_init()
	file_protos_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_protos_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_protos_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_protos_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderListRequest); i {
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
		file_protos_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderListResponse); i {
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
		file_protos_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdRequest); i {
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
		file_protos_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdResponse); i {
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
		file_protos_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfo); i {
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
		file_protos_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdResponse_Items); i {
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
			RawDescriptor: file_protos_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_order_proto_goTypes,
		DependencyIndexes: file_protos_order_proto_depIdxs,
		MessageInfos:      file_protos_order_proto_msgTypes,
	}.Build()
	File_protos_order_proto = out.File
	file_protos_order_proto_rawDesc = nil
	file_protos_order_proto_goTypes = nil
	file_protos_order_proto_depIdxs = nil
}
