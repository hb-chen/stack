package service

import (
	"context"
	"time"

	"github.com/stack-labs/stack-rpc/auth"
	"github.com/stack-labs/stack-rpc/broker"
	"github.com/stack-labs/stack-rpc/broker/http"
	"github.com/stack-labs/stack-rpc/client"
	clientM "github.com/stack-labs/stack-rpc/client/mucp"
	"github.com/stack-labs/stack-rpc/client/selector"
	selectorR "github.com/stack-labs/stack-rpc/client/selector/registry"
	"github.com/stack-labs/stack-rpc/config"
	"github.com/stack-labs/stack-rpc/debug/profile"
	"github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/registry"
	"github.com/stack-labs/stack-rpc/registry/mdns"
	"github.com/stack-labs/stack-rpc/server"
	"github.com/stack-labs/stack-rpc/server/mucp"
	"github.com/stack-labs/stack-rpc/transport"
	transportH "github.com/stack-labs/stack-rpc/transport/http"
)

type Option func(o *Options)

type Options struct {
	// maybe put them in metadata is better
	Id   string
	Name string
	RPC  string

	Broker    broker.Broker
	Client    client.Client
	Server    server.Server
	Registry  registry.Registry
	Transport transport.Transport
	Selector  selector.Selector
	Config    config.Config
	Logger    logger.Logger
	Auth      auth.Auth
	Profile   profile.Profile

	// Before and After funcs
	BeforeInit  []func() error
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	ClientWrapper     []client.Wrapper
	CallWrapper       []client.CallWrapper
	HandlerWrapper    []server.HandlerWrapper
	SubscriberWrapper []server.SubscriberWrapper
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}

func Logger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

func Broker(b broker.Broker) Option {
	return func(o *Options) {
		o.Broker = b
	}
}

func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

func Config(c config.Config) Option {
	return func(o *Options) {
		o.Config = c
	}
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service.
// Can be used for extra option values.
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// HandleSignal toggles automatic installation of the signal handler that
// traps TERM, INT, and QUIT.  Users of this feature to disable the signal
// handler, should control liveness of the service through the context.
func HandleSignal(b bool) Option {
	return func(o *Options) {
		o.Signal = b
	}
}

func Server(s server.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

// Registry sets the registry for the service
// and the underlying components
func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// Selector sets the selector for the service client
func Selector(s selector.Selector) Option {
	return func(o *Options) {
		o.Selector = s
	}
}

// Transport sets the transport for the service
// and the underlying components
func Transport(t transport.Transport) Option {
	return func(o *Options) {
		o.Transport = t
	}
}

// Convenience options

// Address sets the address of the server
func Address(addr string) Option {
	return func(o *Options) {
		// todo no second level component inited here
		o.Server.Init(server.Address(addr))
	}
}

// Unique server id
func Id(id string) Option {
	return func(o *Options) {
		o.Id = id
	}
}

// Name of the service
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Version of the service
func Version(v string) Option {
	return func(o *Options) {
		// todo no second level component inited here
		o.Server.Init(server.Version(v))
	}
}

// Metadata associated with the service
func Metadata(md map[string]string) Option {
	return func(o *Options) {
		// todo no second level component inited here
		o.Server.Init(server.Metadata(md))
	}
}

// Profile to be used for debug profile
func Profile(p profile.Profile) Option {
	return func(o *Options) {
		o.Profile = p
	}
}

// RegisterTTL specifies the TTL to use when registering the service
func RegisterTTL(t time.Duration) Option {
	return func(o *Options) {
		// todo no second level component inited here
		o.Server.Init(server.RegisterTTL(t))
	}
}

// RegisterInterval specifies the interval on which to re-register
func RegisterInterval(t time.Duration) Option {
	return func(o *Options) {
		// todo no second level component inited here
		o.Server.Init(server.RegisterInterval(t))
	}
}

// WrapClient is a convenience method for wrapping a Client with
// some middleware component. A list of wrappers can be provided.
// Wrappers are applied in reverse order so the last is executed first.
func WrapClient(w ...client.Wrapper) Option {
	return func(o *Options) {
		o.ClientWrapper = w
	}
}

// WrapCall is a convenience method for wrapping a Client CallFunc
func WrapCall(w ...client.CallWrapper) Option {
	return func(o *Options) {
		o.CallWrapper = w
	}
}

// WrapHandler adds a handler Wrapper to a list of options passed into the server
func WrapHandler(w ...server.HandlerWrapper) Option {
	return func(o *Options) {
		o.HandlerWrapper = w
	}
}

// WrapSubscriber adds a subscriber Wrapper to a list of options passed into the server
func WrapSubscriber(w ...server.SubscriberWrapper) Option {
	return func(o *Options) {
		o.SubscriberWrapper = w
	}
}

func BeforeInit(fn func() error) Option {
	return func(o *Options) {
		o.BeforeInit = append(o.BeforeInit, fn)
	}
}

func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}

func NewOptions(opts ...Option) Options {
	opt := Options{
		Broker:    http.NewBroker(),
		Client:    clientM.NewClient(),
		Server:    mucp.NewServer(),
		Registry:  mdns.NewRegistry(),
		Transport: transportH.NewTransport(),
		Selector:  selectorR.NewSelector(),
		Logger:    logger.DefaultLogger,
		Config:    config.DefaultConfig,
		Auth:      auth.NoopAuth,
		Context:   context.Background(),
		Signal:    true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}
