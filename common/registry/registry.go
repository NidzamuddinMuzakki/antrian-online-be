package registry

import (
	"sync"

	"antrian-golang/common/cache"

	"antrian-golang/common/middleware/gin/panic_recovery"

	"github.com/go-playground/validator/v10"
)

type IRegistry interface {
	GetValidator() *validator.Validate
	GetCache() cache.Cacher
	GetPanicRecoveryMiddleware() panic_recovery.IMiddlewarePanicRecovery
}

type registry struct {
	mu                      *sync.Mutex
	validator               *validator.Validate
	cache                   cache.Cacher
	panicRecoveryMiddleware panic_recovery.IMiddlewarePanicRecovery
}

func WithPanicRecoveryMiddleware(panicRecoveryMiddleware panic_recovery.IMiddlewarePanicRecovery) Option {
	return func(s *registry) {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.panicRecoveryMiddleware = panicRecoveryMiddleware
	}
}

func WithValidator(validator *validator.Validate) Option {
	return func(s *registry) {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.validator = validator
	}
}

func WithCache(cache cache.Cacher) Option {
	return func(s *registry) {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.cache = cache
	}
}

type Option func(r *registry)

func NewRegistry(
	options ...Option,
) IRegistry {
	registry := &registry{mu: &sync.Mutex{}}

	for _, option := range options {
		option(registry)
	}

	return registry
}

func (r *registry) GetPanicRecoveryMiddleware() panic_recovery.IMiddlewarePanicRecovery {
	return r.panicRecoveryMiddleware
}

func (r *registry) GetValidator() *validator.Validate {
	return r.validator
}

func (r *registry) GetCache() cache.Cacher {
	return r.cache
}
