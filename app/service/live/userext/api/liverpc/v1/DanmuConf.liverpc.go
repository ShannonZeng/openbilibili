// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/DanmuConf.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ===================
// DanmuConf Interface
// ===================

type DanmuConfRPCClient interface {
	// * 获取个人弹幕配置
	//
	GetAll(ctx context.Context, req *DanmuConfGetAllReq, opts ...liverpc.CallOption) (resp *DanmuConfGetAllResp, err error)

	// * 添加用户可佩戴的的弹幕配置特权
	//
	AddByType(ctx context.Context, req *DanmuConfAddByTypeReq, opts ...liverpc.CallOption) (resp *DanmuConfAddByTypeResp, err error)

	// * 获取所有弹幕特权
	//
	GetAllPrivilege(ctx context.Context, req *DanmuConfGetAllPrivilegeReq, opts ...liverpc.CallOption) (resp *DanmuConfGetAllPrivilegeResp, err error)
}

// =========================
// DanmuConf Live Rpc Client
// =========================

type danmuConfRPCClient struct {
	client *liverpc.Client
}

// NewDanmuConfRPCClient creates a client that implements the DanmuConfRPCClient interface.
func NewDanmuConfRPCClient(client *liverpc.Client) DanmuConfRPCClient {
	return &danmuConfRPCClient{
		client: client,
	}
}

func (c *danmuConfRPCClient) GetAll(ctx context.Context, in *DanmuConfGetAllReq, opts ...liverpc.CallOption) (*DanmuConfGetAllResp, error) {
	out := new(DanmuConfGetAllResp)
	err := doRPCRequest(ctx, c.client, 1, "DanmuConf.getAll", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *danmuConfRPCClient) AddByType(ctx context.Context, in *DanmuConfAddByTypeReq, opts ...liverpc.CallOption) (*DanmuConfAddByTypeResp, error) {
	out := new(DanmuConfAddByTypeResp)
	err := doRPCRequest(ctx, c.client, 1, "DanmuConf.addByType", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *danmuConfRPCClient) GetAllPrivilege(ctx context.Context, in *DanmuConfGetAllPrivilegeReq, opts ...liverpc.CallOption) (*DanmuConfGetAllPrivilegeResp, error) {
	out := new(DanmuConfGetAllPrivilegeResp)
	err := doRPCRequest(ctx, c.client, 1, "DanmuConf.getAllPrivilege", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}