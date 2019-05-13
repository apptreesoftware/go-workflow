package core

import (
	"context"
	"github.com/twitchtv/twirp"
	"net/http"
	"time"
)

func (h WorkflowHistory) GetStartTime() time.Time {
	return time.Unix(h.Start, 0)
}

func (h WorkflowHistory) GetEndTime() time.Time {
	return time.Unix(h.End, 0)
}

func CopyTwirpClientAuthorization(ctx context.Context) context.Context  {
	if auth, ok := ctx.Value("authorization").(string); ok {
		header := make(http.Header)
		header.Set("Authorization", auth)
		outCtx, _ := twirp.WithHTTPRequestHeaders(ctx, header)
		return outCtx
	}
	return ctx
}