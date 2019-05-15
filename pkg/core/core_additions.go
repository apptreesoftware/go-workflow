package core

import (
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func ReadPackageYaml(dir string) (Package, error) {
	yamlFile := filepath.Join(dir, "package.yaml")
	b, err := ioutil.ReadFile(yamlFile)
	if err != nil && os.IsNotExist(err) {
		return Package{}, fmt.Errorf("Could not find a package.yaml in %s", dir)
	}
	pkg := Package{}
	err = yaml.Unmarshal(b, &pkg)
	return pkg, err
}
