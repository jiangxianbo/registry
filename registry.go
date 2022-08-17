package registry

import "context"

//@定义服务注册总接口 Registry,定义方法
//-- Name(): 插件名，例如etcd
//-- Init(opts ...Option): 初始化，里面的用选项模式
//-- Registry(): 服务注册
//-- Unregistry(): 服务反注册，例如服务端停了，注册列表销毁
//-- GetService: 服务发现（ip port[]string）

type Registry interface {
	// Name 插件名
	Name() string
	// Init 初始化
	Init(ctx context.Context, opts ...Option) (err error)
	// Register 服务注册
	Register(ctx context.Context, service *Service) (err error)
	// UnRegister 服务反注册
	UnRegister(ctx context.Context, service *Service) (err error)
	// GetService 服务发现
	GetService(ctx context.Context, name string) (service *Service, err error)
}
