// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/runit/log/run.tmpl
// templates/runit/run.tmpl
// templates/systemd/default/master.target.tmpl
// templates/systemd/default/program.service.tmpl
// templates/systemd-user/default/program.service.tmpl
package main

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

var _templatesRunitLogRunTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8c\xcf\x0a\x82\x40\x10\x87\xef\xf3\x14\xbf\x2c\xbc\x6d\x46\x10\x9d\x3a\x7b\x11\x7c\x84\xa8\x75\x50\x49\xdd\xc5\x59\xfb\x83\xce\xbb\xc7\x42\x1e\xba\x7d\xcc\x37\xbf\x6f\xbb\xc9\xee\xed\x90\x49\x43\xc2\x01\x86\x89\x8a\x32\xbf\xcc\x33\xf6\x9d\xab\xa1\x9a\x45\xbc\x79\xbf\xa2\x1f\x9d\x65\x91\x6b\xf8\x78\x86\xaa\x89\xb7\x61\xea\xa1\x4a\x14\x58\x02\x4c\x85\x64\x57\x94\x79\x82\x65\x41\xff\xa8\xda\x11\xc6\xc3\xf4\x38\x9e\x4f\x87\x55\xa5\x29\x6c\xe3\x5e\x03\xe2\x7c\x12\x1e\xa1\xfa\x73\xc4\x6f\xb6\xb0\x8d\x8f\xad\xe9\xef\x41\x9e\x9d\xab\xd7\x3c\x7d\x03\x00\x00\xff\xff\x0a\xaa\x1d\x31\xba\x00\x00\x00")

func templatesRunitLogRunTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRunitLogRunTmpl,
		"templates/runit/log/run.tmpl",
	)
}

func templatesRunitLogRunTmpl() (*asset, error) {
	bytes, err := templatesRunitLogRunTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/runit/log/run.tmpl", size: 186, mode: os.FileMode(420), modTime: time.Unix(1584456288, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRunitRunTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x4d\x0a\x02\x31\x0c\x46\xf7\x73\x8a\x88\xe0\xae\x16\x3d\x80\x57\x19\x6a\x1a\x9c\xa2\x4d\x4a\x7f\xd4\x61\xc8\xdd\xa5\x95\xd9\xbd\x17\xf2\xf8\x8e\x07\x7b\x0f\x6c\xcb\x32\xa1\x87\x6d\x83\xf3\x47\xf2\x33\xf0\x63\xf6\x21\x13\x56\xc9\x2b\xa8\x4e\xf4\x25\x84\xeb\xed\x74\xf9\x13\x2e\xa9\x54\x30\x6d\x04\xad\x50\x06\x55\x30\x34\xf4\x25\xe8\x6a\x10\x06\x55\xdb\xdd\xa5\x04\xaa\xa6\x63\xca\x82\x54\xca\x5c\xd7\x44\xfb\x8d\x5b\xec\x9f\xc4\xef\x51\xa3\xc4\xe8\xd8\xf7\xcd\x5f\x00\x00\x00\xff\xff\xb9\xe8\x41\x7e\x9a\x00\x00\x00")

func templatesRunitRunTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRunitRunTmpl,
		"templates/runit/run.tmpl",
	)
}

func templatesRunitRunTmpl() (*asset, error) {
	bytes, err := templatesRunitRunTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/runit/run.tmpl", size: 154, mode: os.FileMode(420), modTime: time.Unix(1584458764, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSystemdDefaultMasterTargetTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x0e\xcd\xcb\x2c\x89\xe5\x0a\x4f\xcc\x2b\x29\xb6\xad\xae\x56\x28\x4a\xcc\x4b\x4f\x55\x50\x29\x4b\xcc\x29\x4d\x55\xb0\xb2\x55\xd0\x2b\x28\xca\x4f\x4e\x2d\x2e\x4e\x2d\x56\xa8\xad\xad\xae\x86\xc9\xd4\xd6\x2a\x54\x57\x2b\xa4\xe6\xa5\x28\xd4\xd6\x72\x71\x45\x7b\xe6\x15\x97\x24\xe6\xe4\x40\xcc\x49\x4d\x71\xaa\xb4\xcd\x2d\xcd\x29\xc9\xd4\x2d\x2d\x4e\x2d\xd2\x2b\x49\x2c\x4a\x4f\x2d\xe1\x02\x04\x00\x00\xff\xff\x4f\xc2\xa4\x0c\x6a\x00\x00\x00")

func templatesSystemdDefaultMasterTargetTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSystemdDefaultMasterTargetTmpl,
		"templates/systemd/default/master.target.tmpl",
	)
}

func templatesSystemdDefaultMasterTargetTmpl() (*asset, error) {
	bytes, err := templatesSystemdDefaultMasterTargetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/systemd/default/master.target.tmpl", size: 106, mode: os.FileMode(420), modTime: time.Unix(1584336786, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSystemdDefaultProgramServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xc1\x6a\xe3\x3c\x10\x80\xef\x7e\x0a\x11\xfa\xd3\xcb\x9f\x98\x85\x3d\x2d\xe8\xd6\x74\x09\xdb\x6d\x4a\xdd\xd0\x43\x29\x41\x95\xc6\xee\x50\x79\x64\x46\xe3\x34\xc6\xf8\xdd\x17\xa9\x9b\xa4\xb4\xbb\x7b\xb2\xf4\xcd\x37\xa3\xd1\x58\x0f\x1b\x42\x79\x2c\xc6\x71\xae\xb0\x56\x0b\x07\xd1\x32\x76\x82\x81\xd4\x34\x15\x17\xa7\xad\x1e\xc7\x4f\xe1\x71\x54\x40\x2e\xad\x6e\x0c\xcb\xba\xce\x8e\xe9\x3a\x35\x4d\x0b\x31\xdc\x80\x14\x95\x84\xee\xfe\x19\x68\x43\x04\xe0\xc0\xe9\x01\x62\x51\x3c\x54\xc0\x3b\xb4\xf0\x58\x6c\x22\x70\x4e\xeb\x23\x70\xaa\xf4\x9d\x43\xdf\x65\xd2\xa4\x55\x42\xf7\x81\x5f\x90\x9a\x0b\x64\xb0\x12\x78\xc8\xd1\xd7\x37\xb8\x75\x07\x9a\xcc\x25\xed\x90\x03\xb5\x40\xa2\x6f\xd6\xb7\x77\xd9\xec\x02\xcb\xa7\x60\xf5\xae\xd7\x79\xb6\x38\x58\x88\x71\x4b\xa6\x85\x0f\xf6\x25\x7a\xd0\xf3\x12\xc4\x96\x0e\x6a\xd3\x7b\x29\x0f\xc9\x7f\x37\xe3\x10\x6d\xa0\x1a\x9b\xf2\x74\x50\x1e\x33\x1b\x6a\x40\x9d\xbd\xc0\xf0\xbf\x3a\xdb\x19\xdf\x83\xfa\xa6\xd5\x02\x68\xf7\xb1\xc9\xd9\x38\x66\x4f\x4d\x53\xea\xf6\xb7\x3c\x4d\xb3\x77\x83\x5f\xee\xc1\x56\x62\x58\x74\xf9\x84\x54\x3e\x99\xf8\xac\xe6\xde\xaa\x73\xd8\x83\x55\x73\xa3\x66\xff\xbe\xe7\x4c\x25\x68\x43\xdb\x9a\x5c\xf0\xbc\xb8\x85\x98\xeb\x19\xff\x6a\x86\x78\xd8\x56\x60\xf5\x97\xaf\xb1\xa8\xc4\x90\x33\xec\x56\xd4\xf5\xa2\xa9\xf7\xfe\x88\xd6\xbd\x24\x16\x87\xe8\x43\x73\xa4\x4b\xe6\xc0\x47\x98\x3f\x2b\x07\x24\x58\x23\xb0\xfe\x8f\x8a\x1f\xe8\xfd\xcf\xe0\x40\xb7\xb8\x07\x57\xdc\x61\x0b\xa1\x97\xf4\x6e\xd2\x99\xa9\x3b\x79\x43\xe9\xba\xc7\x87\x4a\x68\xf3\x6f\xba\x46\x0b\x59\x3a\x80\xd3\x68\x0e\xaa\xc7\x16\x65\x1b\x3a\xa0\x6d\x8d\x1e\x62\x8a\x5d\x25\x76\xbd\xbe\x5c\x5d\x2d\x73\xf6\x9f\x9c\x53\xa5\x5f\x01\x00\x00\xff\xff\xcd\x06\x2d\xbd\x25\x03\x00\x00")

func templatesSystemdDefaultProgramServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSystemdDefaultProgramServiceTmpl,
		"templates/systemd/default/program.service.tmpl",
	)
}

func templatesSystemdDefaultProgramServiceTmpl() (*asset, error) {
	bytes, err := templatesSystemdDefaultProgramServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/systemd/default/program.service.tmpl", size: 805, mode: os.FileMode(420), modTime: time.Unix(1584458757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSystemdUserDefaultProgramServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x41\x6a\xf3\x30\x10\x85\xf7\x3a\x85\x2f\x10\xfb\x04\xde\xfc\x24\x3f\x74\x95\x12\xa7\x64\x11\x42\x18\xa4\x89\x3b\xad\x34\x12\xa3\x71\x52\x63\x7c\xf7\x62\x85\x12\x9a\x45\x97\x7a\xdf\x37\x8f\xa7\xe3\x1b\x93\x9e\xcc\x34\xad\x2a\xba\x54\xb5\xc3\x6c\x85\x92\x52\xe4\x6a\x9e\xcd\xfa\xf1\x6c\xa7\xe9\x19\x4f\x53\x85\xec\x16\xaf\x32\xc7\x0e\xe5\x4a\x16\x4f\x66\x3f\x26\x6c\x33\x85\xe4\xd1\x6c\xf8\x4a\x12\x39\x20\x6b\xfb\xba\xdd\xed\x4b\x49\x8a\xa2\xcb\xd1\x2f\xd8\x15\x04\x29\x55\xf3\xbc\x2a\x96\x44\x8b\x39\x9f\x19\x02\x3e\xd9\xff\xc9\x63\x5b\xa4\xf7\x58\x60\x53\xdb\xc8\x17\xea\x9b\x3c\x66\xc5\xe0\x9a\x21\xa3\x34\x7f\x17\xd6\xf8\x28\x34\x9b\x2f\xb4\x9d\x82\x68\x59\x61\x63\x08\x70\xff\xd8\x0e\x73\x89\xc1\xdf\x60\xcc\xe6\x10\xe5\x93\xb8\x5f\x93\xa0\xd5\x28\x63\xd1\x6f\xf7\xf0\xec\x7e\xd2\xe5\xb0\x53\x60\x07\xe2\xb6\x83\xa6\x41\xdb\x8f\x38\x08\x83\x37\xe6\xf8\xc2\x59\xc1\xfb\x93\x39\x00\x2b\xba\x7f\x63\x1b\x06\xaf\xb4\x5a\x26\xd7\x0a\xd2\xa3\x9a\xef\x00\x00\x00\xff\xff\x46\xb5\xdd\x93\x95\x01\x00\x00")

func templatesSystemdUserDefaultProgramServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSystemdUserDefaultProgramServiceTmpl,
		"templates/systemd-user/default/program.service.tmpl",
	)
}

func templatesSystemdUserDefaultProgramServiceTmpl() (*asset, error) {
	bytes, err := templatesSystemdUserDefaultProgramServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/systemd-user/default/program.service.tmpl", size: 405, mode: os.FileMode(420), modTime: time.Unix(1584458760, 0)}
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
	"templates/runit/log/run.tmpl":                        templatesRunitLogRunTmpl,
	"templates/runit/run.tmpl":                            templatesRunitRunTmpl,
	"templates/systemd/default/master.target.tmpl":        templatesSystemdDefaultMasterTargetTmpl,
	"templates/systemd/default/program.service.tmpl":      templatesSystemdDefaultProgramServiceTmpl,
	"templates/systemd-user/default/program.service.tmpl": templatesSystemdUserDefaultProgramServiceTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"runit": &bintree{nil, map[string]*bintree{
			"log": &bintree{nil, map[string]*bintree{
				"run.tmpl": &bintree{templatesRunitLogRunTmpl, map[string]*bintree{}},
			}},
			"run.tmpl": &bintree{templatesRunitRunTmpl, map[string]*bintree{}},
		}},
		"systemd": &bintree{nil, map[string]*bintree{
			"default": &bintree{nil, map[string]*bintree{
				"master.target.tmpl":   &bintree{templatesSystemdDefaultMasterTargetTmpl, map[string]*bintree{}},
				"program.service.tmpl": &bintree{templatesSystemdDefaultProgramServiceTmpl, map[string]*bintree{}},
			}},
		}},
		"systemd-user": &bintree{nil, map[string]*bintree{
			"default": &bintree{nil, map[string]*bintree{
				"program.service.tmpl": &bintree{templatesSystemdUserDefaultProgramServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
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
