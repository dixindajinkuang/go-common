package gls

import (
	"context"

	"github.com/jtolds/gls"
)

var (
	mgr           = gls.NewContextManager()
	glsContextKey = "gls_context_key"
)

func GlsSetContext(ctx context.Context, cb func()) {
	mgr.SetValues(gls.Values{glsContextKey: ctx}, cb)
}

func GlsContext() (ctx context.Context) {
	glsCtx, ok := mgr.GetValue(glsContextKey)
	if ok {
		ctx = glsCtx.(context.Context)
	} else {
		ctx = context.Background()
	}
	return
}
