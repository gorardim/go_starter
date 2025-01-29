package trc

type options struct {
	sign bool
}
type Option func(*options)

func WithSign() Option {
	return func(o *options) {
		o.sign = true
	}
}
