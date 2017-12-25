package call911

// Option is a functional option.
type Option func(*Options)

// Options are worker options.
type Options struct {
	Verbose *bool
	Console *bool
}

func (o Options) Merge(opts Options) Options {
	if opts.Verbose != nil {
		o.Verbose = opts.Verbose
	}

	if opts.Console != nil {
		o.Console = opts.Console
	}

	return o
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
		o.Verbose = &verbose
	}
}

func WithConsole(console bool) Option {
	return func(o *Options) {
		o.Console = &console
	}
}
