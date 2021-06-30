package networking

import "context"

type RequestContext struct {
}

const RequestContextKey = "requestContext"

func getReqCtx(ctx context.Context) RequestContext {
	value := ctx.Value(RequestContextKey)

	c, _ := value.(RequestContext)

	return c
}
