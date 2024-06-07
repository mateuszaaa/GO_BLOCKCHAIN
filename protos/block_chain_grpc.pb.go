// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: protos/block_chain.proto

package block_chain

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BlockchainService_AddBlock_FullMethodName           = "/main.BlockchainService/AddBlock"
	BlockchainService_GetBlockchain_FullMethodName      = "/main.BlockchainService/GetBlockchain"
	BlockchainService_SubscribeNewBlocks_FullMethodName = "/main.BlockchainService/SubscribeNewBlocks"
	BlockchainService_GetNodes_FullMethodName           = "/main.BlockchainService/GetNodes"
	BlockchainService_Welcome_FullMethodName            = "/main.BlockchainService/Welcome"
	BlockchainService_GetBlocks_FullMethodName          = "/main.BlockchainService/GetBlocks"
	BlockchainService_GetTransactionPool_FullMethodName = "/main.BlockchainService/GetTransactionPool"
	BlockchainService_NewTransaction_FullMethodName     = "/main.BlockchainService/NewTransaction"
	BlockchainService_NewBlock_FullMethodName           = "/main.BlockchainService/NewBlock"
	BlockchainService_Ping_FullMethodName               = "/main.BlockchainService/Ping"
)

// BlockchainServiceClient is the client API for BlockchainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainServiceClient interface {
	AddBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockchain(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockchainResponse, error)
	SubscribeNewBlocks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BlockchainService_SubscribeNewBlocksClient, error)
	GetNodes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodesResponse, error)
	Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error)
	GetBlocks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlocksResponse, error)
	GetTransactionPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TransactionPoolResponse, error)
	NewTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*BlockResponse, error)
	NewBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PongResponse, error)
}

type blockchainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainServiceClient(cc grpc.ClientConnInterface) BlockchainServiceClient {
	return &blockchainServiceClient{cc}
}

func (c *blockchainServiceClient) AddBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, BlockchainService_AddBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetBlockchain(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockchainResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockchainResponse)
	err := c.cc.Invoke(ctx, BlockchainService_GetBlockchain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) SubscribeNewBlocks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BlockchainService_SubscribeNewBlocksClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[0], BlockchainService_SubscribeNewBlocks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceSubscribeNewBlocksClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockchainService_SubscribeNewBlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type blockchainServiceSubscribeNewBlocksClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceSubscribeNewBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) GetNodes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodesResponse)
	err := c.cc.Invoke(ctx, BlockchainService_GetNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WelcomeResponse)
	err := c.cc.Invoke(ctx, BlockchainService_Welcome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetBlocks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlocksResponse)
	err := c.cc.Invoke(ctx, BlockchainService_GetBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetTransactionPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TransactionPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionPoolResponse)
	err := c.cc.Invoke(ctx, BlockchainService_GetTransactionPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) NewTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, BlockchainService_NewTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) NewBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, BlockchainService_NewBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PongResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, BlockchainService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServiceServer is the server API for BlockchainService service.
// All implementations must embed UnimplementedBlockchainServiceServer
// for forward compatibility
type BlockchainServiceServer interface {
	AddBlock(context.Context, *BlockRequest) (*BlockResponse, error)
	GetBlockchain(context.Context, *Empty) (*BlockchainResponse, error)
	SubscribeNewBlocks(*Empty, BlockchainService_SubscribeNewBlocksServer) error
	GetNodes(context.Context, *Empty) (*NodesResponse, error)
	Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error)
	GetBlocks(context.Context, *Empty) (*BlocksResponse, error)
	GetTransactionPool(context.Context, *Empty) (*TransactionPoolResponse, error)
	NewTransaction(context.Context, *Transaction) (*BlockResponse, error)
	NewBlock(context.Context, *BlockRequest) (*BlockResponse, error)
	Ping(context.Context, *Empty) (*PongResponse, error)
	mustEmbedUnimplementedBlockchainServiceServer()
}

// UnimplementedBlockchainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainServiceServer struct {
}

func (UnimplementedBlockchainServiceServer) AddBlock(context.Context, *BlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
}
func (UnimplementedBlockchainServiceServer) GetBlockchain(context.Context, *Empty) (*BlockchainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockchain not implemented")
}
func (UnimplementedBlockchainServiceServer) SubscribeNewBlocks(*Empty, BlockchainService_SubscribeNewBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeNewBlocks not implemented")
}
func (UnimplementedBlockchainServiceServer) GetNodes(context.Context, *Empty) (*NodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodes not implemented")
}
func (UnimplementedBlockchainServiceServer) Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}
func (UnimplementedBlockchainServiceServer) GetBlocks(context.Context, *Empty) (*BlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedBlockchainServiceServer) GetTransactionPool(context.Context, *Empty) (*TransactionPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionPool not implemented")
}
func (UnimplementedBlockchainServiceServer) NewTransaction(context.Context, *Transaction) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTransaction not implemented")
}
func (UnimplementedBlockchainServiceServer) NewBlock(context.Context, *BlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBlock not implemented")
}
func (UnimplementedBlockchainServiceServer) Ping(context.Context, *Empty) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBlockchainServiceServer) mustEmbedUnimplementedBlockchainServiceServer() {}

// UnsafeBlockchainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServiceServer will
// result in compilation errors.
type UnsafeBlockchainServiceServer interface {
	mustEmbedUnimplementedBlockchainServiceServer()
}

func RegisterBlockchainServiceServer(s grpc.ServiceRegistrar, srv BlockchainServiceServer) {
	s.RegisterService(&BlockchainService_ServiceDesc, srv)
}

func _BlockchainService_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_AddBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).AddBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_GetBlockchain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetBlockchain(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_SubscribeNewBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockchainServiceServer).SubscribeNewBlocks(m, &blockchainServiceSubscribeNewBlocksServer{ServerStream: stream})
}

type BlockchainService_SubscribeNewBlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type blockchainServiceSubscribeNewBlocksServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceSubscribeNewBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockchainService_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_GetNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetNodes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_Welcome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).Welcome(ctx, req.(*WelcomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_GetBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetBlocks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetTransactionPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetTransactionPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_GetTransactionPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetTransactionPool(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_NewTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).NewTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_NewTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).NewTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_NewBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).NewBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_NewBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).NewBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockchainService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockchainService_ServiceDesc is the grpc.ServiceDesc for BlockchainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockchainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.BlockchainService",
	HandlerType: (*BlockchainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlock",
			Handler:    _BlockchainService_AddBlock_Handler,
		},
		{
			MethodName: "GetBlockchain",
			Handler:    _BlockchainService_GetBlockchain_Handler,
		},
		{
			MethodName: "GetNodes",
			Handler:    _BlockchainService_GetNodes_Handler,
		},
		{
			MethodName: "Welcome",
			Handler:    _BlockchainService_Welcome_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _BlockchainService_GetBlocks_Handler,
		},
		{
			MethodName: "GetTransactionPool",
			Handler:    _BlockchainService_GetTransactionPool_Handler,
		},
		{
			MethodName: "NewTransaction",
			Handler:    _BlockchainService_NewTransaction_Handler,
		},
		{
			MethodName: "NewBlock",
			Handler:    _BlockchainService_NewBlock_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _BlockchainService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeNewBlocks",
			Handler:       _BlockchainService_SubscribeNewBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/block_chain.proto",
}
