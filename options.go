package call911

// Option is a functional option.
type Option func(*Options)

// Options are worker options.
type Options struct {
	Verbose      *bool
	Console      *bool
	ErrorHandler ErrorHandler
}

// Merge merges two options.
func (o Options) Merge(opts Options) Options {
	if opts.Verbose != nil {
		o.Verbose = opts.Verbose
	}

	if opts.Console != nil {
		o.Console = opts.Console
	}

	if opts.ErrorHandler != nil {
		o.ErrorHandler = opts.ErrorHandler
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

// WithVerbose sets the verbose option.
func WithVerbose(verbose bool) Option {
	return func(o *Options) {
		o.Verbose = &verbose
	}
}

// WithConsole sets the console option.
func WithConsole(console bool) Option {
	return func(o *Options) {
		o.Console = &console
	}
}

// WithErrorHandler sets the error handler option.
func WithErrorHandler(errorHandler ErrorHandler) Option {
	return func(o *Options) {
		o.ErrorHandler = errorHandler
	}
}
