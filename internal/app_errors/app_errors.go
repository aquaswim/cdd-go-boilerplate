package appErrors

var (
	FrameworkError  = appErrorType.New("Framework Error").WithProperty(HttpCodeProperty, 500).WithProperty(CodeProperty, "framework_error")
	InternalError   = appErrorType.New("Internal Error").WithProperty(HttpCodeProperty, 500).WithProperty(CodeProperty, "internal_error")
	ValidationError = appErrorType.New("Validation Error").WithProperty(HttpCodeProperty, 400).WithProperty(CodeProperty, "validation_error")
	NotFoundError   = appErrorType.New("Not Found").WithProperty(HttpCodeProperty, 404).WithProperty(CodeProperty, "not_found")
)
