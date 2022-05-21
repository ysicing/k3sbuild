//go:generate go run pkg/codegen/cleanup/main.go
//go:generate rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go
//go:generate go fmt pkg/deploy/zz_generated_bindata.go
//go:generate go fmt pkg/static/zz_generated_bindata.go

package main

import (
	"path/filepath"

	"github.com/ergoapi/util/file"
	"github.com/sirupsen/logrus"
	"github.com/ysicing/k3sbuild/pkg/static"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	datadir, _ := file.MkTmpdir("/tmp")
	logrus.Debug("tmp datadir ", datadir)
	dataDir := filepath.Join(datadir, "static")
	if err := static.Stage(dataDir); err != nil {
		logrus.Error(err)
		return
	}
}
