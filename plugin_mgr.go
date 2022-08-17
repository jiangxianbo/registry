package registry

import (
	"context"
	"fmt"
	"sync"
)

//@插件管理
//-- 可以用一个大map管理，key字符串，value是Regitry接口对象
//-- 用户自定义去调用，自定义插件
//-- 实现注册中心的初始化，供系统后使用

type PluginMgr struct {
	// map维护所有插件
	plugin map[string]Registry
	lock   sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugin: make(map[string]Registry),
	}
)

// RegisterPlugin 插件注册
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

// 注册插件
func (p PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 查看里面有没有
	_, ok := p.plugin[plugin.Name()]
	if ok {
		err = fmt.Errorf("registry plugin exist")
		return
	}
	p.plugin[plugin.Name()] = plugin
	return
}

// InitRegistry 进行初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

func (p PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 查看服务是否存在
	plugin, ok := p.plugin[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exist", name)
		return
	}
	registry = plugin
	// 进行组件初始化
	err = plugin.Init(ctx, opts...)
	return
}
