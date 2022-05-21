// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/echo.yaml
//go:build !no_stage
// +build !no_stage

package deploy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _echoYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x4d\x6f\xdb\x48\x0c\x86\xef\xfe\x15\x04\x82\xbd\xad\x57\x76\x16\x01\x0a\xdd\xd2\xa6\x5f\x68\x8d\x1a\x71\x52\xa0\xa7\x80\x1e\xd1\xd6\xc0\x9c\x8f\xce\x70\x5c\x2b\xbf\xbe\x90\x15\xd9\x23\xdb\x87\x02\x6d\x79\x13\x5f\xf2\x91\x5e\x8a\x33\xe3\xf1\x78\x84\x5e\x7f\xa5\x10\xb5\xb3\x25\xa0\xf7\xb1\xd8\x4e\x47\x1b\x6d\xab\x12\xee\xc8\xb3\x6b\x0c\x59\x19\x19\x12\xac\x50\xb0\x1c\x01\x58\x34\x54\x02\xcc\xbe\xdd\xce\xe7\x2f\x8f\xd1\xa3\xa2\x12\x2a\x5a\x61\x62\x19\x01\x30\x2e\x89\x63\x5b\x0d\x2d\xf4\x50\x1e\x3d\xa9\x36\x1b\x89\x49\x89\x0b\x5d\x85\x41\x51\xf5\xe7\xac\xe5\xa5\xa9\x7f\x45\x20\xcf\x5a\x61\x2c\x61\xda\xf6\x4a\x40\xa1\x75\xd3\x95\x06\xc7\xac\xed\xfa\xd1\x57\x28\xd4\x77\x1b\xdc\x2d\x52\x58\x53\x09\xd7\x37\xff\x1c\x73\x8f\x16\xb7\xa8\x19\x97\x9c\x29\xd2\x78\x2a\xe1\x3e\xc7\x8c\x00\x84\x8c\xe7\x03\x31\xb7\xdf\x06\x0f\xbe\x75\x68\xb1\x7d\xee\x6d\xb6\x71\x05\xda\x6a\x79\xe3\xac\xa0\xb6\x14\xb2\xae\x2b\xf8\x68\xb5\x80\x3a\x48\x80\x81\x80\x76\xa8\x84\x1b\x60\xbd\x21\x08\xb4\x4e\x8c\x21\xab\xf9\x17\x68\xa7\xc8\xcb\x11\xd3\x82\xc6\xe7\x28\xfe\x81\x4d\x84\x90\x2c\x88\x03\xe5\x8c\x67\x12\xed\xec\x7f\x27\x7d\x6f\x51\xd5\xfb\x2f\x3c\x36\x83\x49\x51\xfa\x16\x82\x98\x94\xa2\x18\x57\x89\xb9\x81\x25\xad\x5c\x20\x90\x9a\xc0\xd2\x4e\xc0\x59\x82\x28\x18\x24\xf6\x60\x75\xe6\x74\x7c\xba\x32\x5d\x68\x83\xeb\x43\xb6\x6c\x87\x1d\xe5\x20\x06\x8a\x2e\x05\x45\x31\xf7\x19\xe8\x7b\xa2\x28\x83\x1c\x80\xf2\xa9\x84\xe9\x64\x62\x06\x59\x43\xc6\x85\x66\x2f\xcc\x74\xa6\xb0\x36\xfa\x77\x08\xac\xb7\x64\x29\xc6\x79\x70\x4b\xca\x31\xa2\xfc\xc2\xa9\x0d\xc9\x90\xed\x5d\x90\x12\x5e\x4d\xb2\x64\x3b\x6d\x8d\x7c\x47\x8c\xcd\x82\x94\xb3\x55\x2c\xe1\x26\x47\x69\x43\x2e\xc9\x45\xed\xe5\x6f\x3c\xd4\x81\x62\xed\xb8\xea\xce\x44\x1f\x2b\xd4\x9c\x02\x65\xea\xff\x99\xea\x29\x68\x57\x1d\xb0\xd3\x49\x36\x6e\xac\xf4\x45\x5b\xb5\x88\x7f\x7f\x66\x0a\xa5\x2e\xa1\x78\x8a\x82\x92\x62\x51\x13\xb2\xd4\xcf\x7f\xdc\xf7\xf5\x5f\xf7\x4d\x76\x7b\x74\xd6\xef\xe9\xdd\xeb\xa7\x0f\x5f\x16\x0f\x19\x60\x8b\x9c\xe8\x5d\x70\xe6\x64\x6f\x9c\x5d\xe9\xf5\x0c\xfd\x27\x6a\xee\x69\x35\x14\xfb\x8b\x72\xb8\xf4\x5d\x6c\xa8\x39\x7f\x4d\x3b\xb1\x98\x7f\xcd\xe1\x24\xcd\xf7\xb3\x1c\x0e\xf3\xe2\x91\xda\x3a\x4e\x86\x66\x2e\xd9\x21\xa9\x2b\x66\xa7\x90\xdb\x19\x67\x18\xd3\xd6\xce\xbb\xdf\x49\xa2\x8a\xd3\x9a\x8e\xf8\x8b\xb0\xda\xc5\x8e\x75\x69\x59\x52\x0c\x45\xac\x31\x50\xf1\xec\x2c\x69\xbb\x72\xc5\x6d\xd4\x58\x2c\x6a\xb4\xeb\x1a\xfb\x13\x16\x68\x7f\x9d\xcc\x1d\x6b\xd5\x94\x70\xbb\xbf\xc6\x7e\x06\x00\x00\xff\xff\x2e\xff\x7e\x7c\xa8\x06\x00\x00")

func echoYamlBytes() ([]byte, error) {
	return bindataRead(
		_echoYaml,
		"echo.yaml",
	)
}

func echoYaml() (*asset, error) {
	bytes, err := echoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "echo.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"echo.yaml": echoYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"echo.yaml": &bintree{echoYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
