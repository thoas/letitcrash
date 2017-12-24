package call911

// Option is a functional option.
type Option func(*Options)

// Options are worker options.
type Options struct {
	Verbose bool
}

func newOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func WithVerbose(verbose bool) Option {
	return func(o *Options) {
		o.Verbose = verbose
	}
}
