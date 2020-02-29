// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.server.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.server.proto
*/
package api

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathServerPing = "/server.service.v1.Server/Ping"
var PathServerSayHello = "/server.service.v1.Server/SayHello"
var PathServerSayHelloURL = "/kratos-demo/say_hello"

// ServerBMServer is the server API for Server service.
type ServerBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	SayHello(ctx context.Context, req *HelloReq) (resp *google_protobuf1.Empty, err error)

	SayHelloURL(ctx context.Context, req *HelloReq) (resp *HelloResp, err error)
}

var ServerSvc ServerBMServer

func serverPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ServerSvc.Ping(c, p)
	c.JSON(resp, err)
}

func serverSayHello(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ServerSvc.SayHello(c, p)
	c.JSON(resp, err)
}

func serverSayHelloURL(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ServerSvc.SayHelloURL(c, p)
	c.JSON(resp, err)
}

// RegisterServerBMServer Register the blademaster route
func RegisterServerBMServer(e *bm.Engine, server ServerBMServer) {
	ServerSvc = server
	e.GET("/server.service.v1.Server/Ping", serverPing)
	e.GET("/server.service.v1.Server/SayHello", serverSayHello)
	e.GET("/kratos-demo/say_hello", serverSayHelloURL)
}