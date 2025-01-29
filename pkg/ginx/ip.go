package ginx

import "context"

func IpFromContext(ctx context.Context) string {
	c := FromContext(ctx)
	return c.ClientIP()
}
