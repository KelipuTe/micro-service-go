package registry

import "context"

// S6ServiceInstance 代表的是一个实例
type S6ServiceInstance struct {
	ServiceName string
	Address     string
}

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeAdd
	EventTypeDelete
	EventTypeUpdate
	// EventTypeErr
)

type Event struct {
	Type     EventType
	Instance S6ServiceInstance
}

type I9Registry interface {
	Register(i9ctx context.Context, s6si S6ServiceInstance) error
	// Unregister(ctx context.Context, serviceName string) error
	Unregister(i9ctx context.Context, s6si S6ServiceInstance) error
	ListService(i9ctx context.Context, serviceName string) ([]S6ServiceInstance, error)
	// 可以考虑利用 ctx 来 close 掉返回的 channel
	// Subscribe(ctx context.Context, serviceName string) (<- chan Event, error)

	Subscribe(serviceName string) (<-chan Event, error)
	// 可有可无，不定义的话，具体的实现也可以额外的添加
	Close() error
}
