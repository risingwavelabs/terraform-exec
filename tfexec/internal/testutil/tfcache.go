package testutil

import (
	"context"
	"sync"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/build"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
)

const (
	Latest011   = "0.11.15"
	Latest012   = "0.12.31"
	Latest013   = "0.13.7"
	Latest014   = "0.14.11"
	Latest015   = "0.15.5"
	Latest_v1   = "1.0.0"
	Latest_v1_1 = "1.1.0-alpha20211006"
)

const appendUserAgent = "tfexec-testutil"

type TFCache struct {
	sync.Mutex

	dir   string
	execs map[string]string
}

func NewTFCache(dir string) *TFCache {
	return &TFCache{
		dir:   dir,
		execs: map[string]string{},
	}
}

func (tf *TFCache) GitRef(t *testing.T, ref string) string {
	t.Helper()

	key := "gitref:" + ref

	return tf.find(t, key, func(ctx context.Context) (string, error) {
		gr := &build.GitRevision{
			Product: product.Terraform,
			Ref:     ref,
		}
		// gr.SetLogger(TestLogger())

		return gr.Build(ctx)
	})
}

func (tf *TFCache) Version(t *testing.T, v string) string {
	t.Helper()

	key := "v:" + v

	return tf.find(t, key, func(ctx context.Context) (string, error) {
		ev := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(v)),
		}
		// ev.SetLogger(TestLogger())

		return ev.Install(ctx)
	})
}
