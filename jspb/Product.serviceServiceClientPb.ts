/**
 * @fileoverview gRPC-Web generated client stub for productservice
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.1
// 	protoc              v3.17.3
// source: product.service


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as product_pb from './product_pb';


export class ProductServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreate = new grpcWeb.MethodDescriptor(
    '/productservice.ProductService/Create',
    grpcWeb.MethodType.UNARY,
    product_pb.Product,
    product_pb.Product,
    (request: product_pb.Product) => {
      return request.serializeBinary();
    },
    product_pb.Product.deserializeBinary
  );

  create(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null): Promise<product_pb.Product>;

  create(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: product_pb.Product) => void): grpcWeb.ClientReadableStream<product_pb.Product>;

  create(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: product_pb.Product) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/productservice.ProductService/Create',
        request,
        metadata || {},
        this.methodDescriptorCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/productservice.ProductService/Create',
    request,
    metadata || {},
    this.methodDescriptorCreate);
  }

  methodDescriptorList = new grpcWeb.MethodDescriptor(
    '/productservice.ProductService/List',
    grpcWeb.MethodType.UNARY,
    product_pb.Product,
    product_pb.ProductList,
    (request: product_pb.Product) => {
      return request.serializeBinary();
    },
    product_pb.ProductList.deserializeBinary
  );

  list(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null): Promise<product_pb.ProductList>;

  list(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: product_pb.ProductList) => void): grpcWeb.ClientReadableStream<product_pb.ProductList>;

  list(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: product_pb.ProductList) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/productservice.ProductService/List',
        request,
        metadata || {},
        this.methodDescriptorList,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/productservice.ProductService/List',
    request,
    metadata || {},
    this.methodDescriptorList);
  }

  methodDescriptorGet = new grpcWeb.MethodDescriptor(
    '/productservice.ProductService/Get',
    grpcWeb.MethodType.UNARY,
    product_pb.Product,
    product_pb.Product,
    (request: product_pb.Product) => {
      return request.serializeBinary();
    },
    product_pb.Product.deserializeBinary
  );

  get(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null): Promise<product_pb.Product>;

  get(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: product_pb.Product) => void): grpcWeb.ClientReadableStream<product_pb.Product>;

  get(
    request: product_pb.Product,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: product_pb.Product) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/productservice.ProductService/Get',
        request,
        metadata || {},
        this.methodDescriptorGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/productservice.ProductService/Get',
    request,
    metadata || {},
    this.methodDescriptorGet);
  }

  methodDescriptorSubscribe = new grpcWeb.MethodDescriptor(
    '/productservice.ProductService/Subscribe',
    grpcWeb.MethodType.SERVER_STREAMING,
    product_pb.Product,
    product_pb.Product,
    (request: product_pb.Product) => {
      return request.serializeBinary();
    },
    product_pb.Product.deserializeBinary
  );

  subscribe(
    request: product_pb.Product,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<product_pb.Product> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/productservice.ProductService/Subscribe',
      request,
      metadata || {},
      this.methodDescriptorSubscribe);
  }

}

