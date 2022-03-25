// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: books/v1/books.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BooksServiceClient is the client API for BooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksServiceClient interface {
	BooksByTags(ctx context.Context, in *BooksByTagsRequest, opts ...grpc.CallOption) (*BooksByTagsResponse, error)
	BooksByTitleYear(ctx context.Context, in *BooksByTitleYearRequest, opts ...grpc.CallOption) (*BooksByTitleYearResponse, error)
	CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*CreateAuthorResponse, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*GetAuthorResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error)
	UpdateBookISBN(ctx context.Context, in *UpdateBookISBNRequest, opts ...grpc.CallOption) (*UpdateBookISBNResponse, error)
}

type booksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksServiceClient(cc grpc.ClientConnInterface) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) BooksByTags(ctx context.Context, in *BooksByTagsRequest, opts ...grpc.CallOption) (*BooksByTagsResponse, error) {
	out := new(BooksByTagsResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/BooksByTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) BooksByTitleYear(ctx context.Context, in *BooksByTitleYearRequest, opts ...grpc.CallOption) (*BooksByTitleYearResponse, error) {
	out := new(BooksByTitleYearResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/BooksByTitleYear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*CreateAuthorResponse, error) {
	out := new(CreateAuthorResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*GetAuthorResponse, error) {
	out := new(GetAuthorResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) UpdateBookISBN(ctx context.Context, in *UpdateBookISBNRequest, opts ...grpc.CallOption) (*UpdateBookISBNResponse, error) {
	out := new(UpdateBookISBNResponse)
	err := c.cc.Invoke(ctx, "/books.v1.BooksService/UpdateBookISBN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServiceServer is the server API for BooksService service.
// All implementations must embed UnimplementedBooksServiceServer
// for forward compatibility
type BooksServiceServer interface {
	BooksByTags(context.Context, *BooksByTagsRequest) (*BooksByTagsResponse, error)
	BooksByTitleYear(context.Context, *BooksByTitleYearRequest) (*BooksByTitleYearResponse, error)
	CreateAuthor(context.Context, *CreateAuthorRequest) (*CreateAuthorResponse, error)
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
	GetAuthor(context.Context, *GetAuthorRequest) (*GetAuthorResponse, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error)
	UpdateBookISBN(context.Context, *UpdateBookISBNRequest) (*UpdateBookISBNResponse, error)
	mustEmbedUnimplementedBooksServiceServer()
}

// UnimplementedBooksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServiceServer struct {
}

func (UnimplementedBooksServiceServer) BooksByTags(context.Context, *BooksByTagsRequest) (*BooksByTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BooksByTags not implemented")
}
func (UnimplementedBooksServiceServer) BooksByTitleYear(context.Context, *BooksByTitleYearRequest) (*BooksByTitleYearResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BooksByTitleYear not implemented")
}
func (UnimplementedBooksServiceServer) CreateAuthor(context.Context, *CreateAuthorRequest) (*CreateAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (UnimplementedBooksServiceServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBooksServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBooksServiceServer) GetAuthor(context.Context, *GetAuthorRequest) (*GetAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedBooksServiceServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBooksServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBooksServiceServer) UpdateBookISBN(context.Context, *UpdateBookISBNRequest) (*UpdateBookISBNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookISBN not implemented")
}
func (UnimplementedBooksServiceServer) mustEmbedUnimplementedBooksServiceServer() {}

// UnsafeBooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServiceServer will
// result in compilation errors.
type UnsafeBooksServiceServer interface {
	mustEmbedUnimplementedBooksServiceServer()
}

func RegisterBooksServiceServer(s grpc.ServiceRegistrar, srv BooksServiceServer) {
	s.RegisterService(&BooksService_ServiceDesc, srv)
}

func _BooksService_BooksByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BooksByTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).BooksByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/BooksByTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).BooksByTags(ctx, req.(*BooksByTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_BooksByTitleYear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BooksByTitleYearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).BooksByTitleYear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/BooksByTitleYear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).BooksByTitleYear(ctx, req.(*BooksByTitleYearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).CreateAuthor(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_UpdateBookISBN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookISBNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).UpdateBookISBN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.v1.BooksService/UpdateBookISBN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).UpdateBookISBN(ctx, req.(*UpdateBookISBNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BooksService_ServiceDesc is the grpc.ServiceDesc for BooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.v1.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BooksByTags",
			Handler:    _BooksService_BooksByTags_Handler,
		},
		{
			MethodName: "BooksByTitleYear",
			Handler:    _BooksService_BooksByTitleYear_Handler,
		},
		{
			MethodName: "CreateAuthor",
			Handler:    _BooksService_CreateAuthor_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BooksService_CreateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BooksService_DeleteBook_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _BooksService_GetAuthor_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BooksService_GetBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BooksService_UpdateBook_Handler,
		},
		{
			MethodName: "UpdateBookISBN",
			Handler:    _BooksService_UpdateBookISBN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books/v1/books.proto",
}
