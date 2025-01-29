package nsqx

import "context"

type contextKey struct{}

func NewContext(ctx context.Context) (context.Context, func()) {
	ctx = context.WithValue(ctx, contextKey{}, &consumerCancelRef{})
	return ctx, func() {
		ref, ok := ctx.Value(contextKey{}).(*consumerCancelRef)
		if !ok || ref == nil {
			return
		}
		for _, cancel := range ref.list {
			cancel()
		}
	}
}

type consumerCancelRef struct {
	list []func()
}
