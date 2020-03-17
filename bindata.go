// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/launchd/launchd.plist.tmpl
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

var _templatesLaunchdLaunchdPlistTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x93\x41\x6f\xda\x4e\x10\xc5\xef\x7c\x8a\xf9\x5b\x1c\xff\x78\xe1\x56\x45\xc6\x11\x0d\x54\x8a\x8a\x82\x95\x40\xab\x9e\xd0\xc6\x3b\x75\x56\xd8\xbb\xee\xec\x1a\x6a\x59\xfb\xdd\xab\x35\xb8\x98\x24\x54\x3d\xf5\x36\x1a\xbd\x79\xbf\x37\xab\xd9\xe8\xf6\x67\x91\xc3\x1e\xc9\x48\xad\xa6\xc1\x24\x1c\x07\x80\x2a\xd5\x42\xaa\x6c\x1a\x6c\xd6\x9f\x46\x1f\x82\xdb\x78\x10\xfd\x37\x5f\xdd\xad\xbf\x25\x0b\x28\x73\x69\x2c\x24\x9b\x8f\xcb\xfb\x3b\x08\x46\x8c\xcd\xca\x32\x47\xc6\xe6\xeb\x39\x24\xcb\xfb\xa7\x35\x4c\xc2\x31\x63\x8b\x87\x00\x82\x17\x6b\xcb\x1b\xc6\x0e\x87\x43\xc8\xbd\x2a\x4c\x75\xe1\x85\x86\x25\xa4\x4b\x24\x5b\x2f\xa5\xb1\xa3\x49\x38\x0e\x85\x15\x41\x3c\x88\x8e\xee\x17\x71\xe2\x41\x24\x64\x6a\xe3\x01\x00\x40\xb4\xc3\x3a\x5e\xf2\x67\xcc\x23\xe6\xcb\x63\xd3\x58\x92\x2a\x8b\x9b\x06\x3c\x07\x9c\x1b\xf9\xb2\x24\x9d\xa2\x31\x5b\x5b\x97\xd8\xf5\x54\x55\x80\x73\x11\x3b\x4d\x9c\x3d\x17\x6a\x2f\x49\xab\x02\x95\xfd\xc2\x49\xf2\xe7\x1c\x4d\x1f\x71\x8e\xf0\x7b\x24\x59\x3d\xae\x7b\x92\xd7\x49\x4a\x4d\xf6\x0d\xab\x69\x46\x40\x5c\x65\x08\xc3\x1d\xd6\xff\xc3\x70\xcf\xf3\x0a\xe1\x66\x0a\x21\xaa\x3d\x38\x77\x89\x68\x9a\x56\xd7\xda\x5c\x03\x9d\x2c\x7a\xa4\xa6\x01\x54\xa2\x33\x8b\xd8\xab\xe7\x4b\x48\x67\xc4\x8b\x19\x65\x95\x5f\xf7\x62\x4d\x4e\xc4\xeb\x37\x51\x53\x5d\x14\x5c\x89\x36\xe6\xa9\xde\x72\xca\xcc\x45\xde\x33\x5c\x7e\x07\xfc\x71\x9e\x0a\x86\xfe\xa5\x02\x70\xce\xa7\xed\xde\xc5\x87\xcc\x0d\x9e\xba\x9d\xf6\xd8\x6f\x8b\xab\xeb\xf4\x32\xb6\xfb\x7c\x46\x2c\x67\xb9\xdc\x63\x7f\x11\x4b\x15\xb2\x9e\xe8\xb1\x52\x33\xbb\xd4\x5c\xfc\x49\xf4\x64\xb9\x12\x9c\xc4\xaa\xb2\x09\xb7\x2f\x57\x4e\x2c\xd7\x19\x38\xc7\xce\xd7\xc6\xfe\xfa\xf0\xfc\xec\x3b\xc7\xd7\x71\x17\x44\x9a\xfe\x29\x79\x63\x90\x1e\x78\x81\x57\x80\x95\x41\x7a\xff\xbf\x7c\xd5\xb4\x93\x2a\x9b\x4b\xc2\xd4\x6a\xaa\xaf\x18\x1c\x8e\xb2\xad\xe8\x74\x17\x6e\xdd\x69\x46\xac\xfd\xf7\xf1\xe0\x57\x00\x00\x00\xff\xff\x2d\x1b\xa9\x26\x8e\x04\x00\x00")

func templatesLaunchdLaunchdPlistTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesLaunchdLaunchdPlistTmpl,
		"templates/launchd/launchd.plist.tmpl",
	)
}

func templatesLaunchdLaunchdPlistTmpl() (*asset, error) {
	bytes, err := templatesLaunchdLaunchdPlistTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/launchd/launchd.plist.tmpl", size: 1166, mode: os.FileMode(420), modTime: time.Unix(1584462008, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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

var _templatesSystemdDefaultProgramServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x51\x6b\xdb\x3e\x10\xc0\xdf\xf5\x29\x44\xe8\x9f\xbe\xfc\x13\x33\xd8\xd3\xc0\x6f\x4d\x47\x58\xd7\x94\xba\xa1\x0f\xa5\x04\x55\x3a\xbb\x47\xe5\x93\x39\x9d\xd3\x18\xe3\xef\x3e\xa4\x2e\xc9\xc8\x36\xd8\x93\xa5\xdf\xfd\xee\x74\x27\xeb\x69\x43\x28\xcf\x6a\x1c\xe7\x1a\x6b\xbd\x70\x10\x2d\x63\x27\x18\x48\x4f\x93\xba\x3a\x6d\xcb\x71\x3c\x0f\x8f\xa3\x06\x72\xc9\xbb\x33\x2c\xeb\x3a\x2b\xa6\xeb\xf4\x34\x2d\xc4\x70\x03\xa2\x2a\x09\xdd\xe3\x2b\xd0\x86\x08\xc0\x81\x2b\x07\x88\x4a\x3d\x55\xc0\x3b\xb4\xf0\xac\x36\x11\x38\xa7\xf5\x11\x38\x55\xfa\xca\xa1\xef\x32\x69\xd2\x2a\xa1\xc7\xc0\x6f\x48\xcd\x15\x32\x58\x09\x3c\xe4\xe8\xfb\x07\xdc\xba\x03\x4d\xe6\x92\x76\xc8\x81\x5a\x20\x29\xef\xd6\xf7\x0f\xd9\xec\x02\xcb\x6f\xc1\xea\x97\x5e\xe7\xd9\xe2\x60\x21\xc6\xad\x0c\x1d\xa4\xfe\x13\xa3\xbe\x3d\x4b\xbc\x46\x0f\xe5\xbc\x00\xb1\x85\x83\xda\xf4\x5e\x8a\x43\x9d\xbf\x9b\x71\x88\x36\x50\x8d\x4d\x71\x3a\x33\x5f\x38\x1b\x6a\x40\x5f\xbc\xc1\xf0\xbf\xbe\xd8\x19\xdf\x83\xfe\x52\xea\x05\xd0\xee\xbc\xdf\xd9\x38\x66\x4f\x4f\x53\x6a\xfc\xa7\x3c\x4d\x33\x75\xfa\x07\xcb\x3d\xd8\x4a\x0c\x4b\x59\xbc\x20\x15\x2f\x26\xbe\xea\xb9\xb7\xfa\x12\xf6\x60\xf5\xdc\xe8\xd9\x3f\x8f\x3c\xd3\x69\x6d\x43\xdb\x9a\x5c\xfb\x52\xdd\x43\xcc\xa5\x8d\x7f\x37\x43\x3c\x6c\x2b\xb0\xe5\xa7\xcf\x51\x55\x62\xc8\x19\x76\x2b\xea\x7a\x29\xa9\xf7\xfe\x88\xd6\xbd\x24\x16\x87\xe8\x43\x73\xa4\x4b\xe6\xc0\x47\x98\x3f\x2b\x07\x24\x58\x23\x70\xf9\x1f\xa9\x6f\xe8\xfd\xf7\xe0\xa0\x6c\x71\x0f\x4e\x3d\x60\x0b\xa1\x97\xf4\x9a\xd2\x99\xa9\x3b\xf9\x40\x69\xf2\xe3\xeb\x25\xb4\x69\x12\x75\x8b\x16\xb2\x74\x00\xa7\x5b\x3a\xa8\x1e\x5b\x94\x6d\xe8\x80\xb6\x35\x7a\x88\x29\x76\x93\xd8\xed\xfa\x7a\x75\xb3\xcc\xd9\x7f\x72\x4e\x95\x7e\x04\x00\x00\xff\xff\xb7\xf7\x3a\x4f\x3a\x03\x00\x00")

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

	info := bindataFileInfo{name: "templates/systemd/default/program.service.tmpl", size: 826, mode: os.FileMode(420), modTime: time.Unix(1584460118, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSystemdUserDefaultProgramServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x41\x6a\xf3\x30\x10\x85\xf7\x3a\x85\x2e\x10\xfb\x04\xde\xfc\x24\x3f\x74\x95\x12\xa7\x64\x11\x42\x10\xd2\xc4\x9d\x56\x1a\x89\xd1\x38\xa9\x31\xbe\x7b\x91\x42\x09\xcd\xaa\x4b\x7d\xef\x7b\x8f\xd1\xf1\x8d\x50\x4e\x6a\x9e\x57\x1a\x2f\xba\x71\x90\x2d\x63\x12\x8c\xa4\x97\x45\xad\x1f\xcf\x6e\x9e\x9f\xe3\x79\xd6\x40\xae\x78\x5a\x1d\x7b\xe0\x2b\x5a\x38\xa9\xfd\x94\xa0\xcb\x18\x92\x07\xb5\xa1\x2b\x72\xa4\x00\x24\xdd\xeb\x76\xb7\xaf\x23\x29\xb2\x94\xd2\xaf\xb0\xaf\x91\x49\x49\x2f\xcb\xaa\x5a\x1c\x2d\xe4\x7c\x96\x29\x81\x5e\x96\xa6\x30\x1a\xc3\x53\xf1\x3f\x7a\xe8\xaa\xff\x1e\x43\xf1\xda\xc6\x46\xba\xe0\xd0\xe6\x29\x0b\x04\xd7\x8e\x19\xb8\xfd\xf3\x76\x03\x8f\x6d\xb5\xf9\x02\xdb\x8b\x61\xa9\xb7\xd9\x18\x82\xb9\x7f\x77\x07\xb9\x62\xe3\x6f\x66\xca\xea\x10\xf9\x13\x69\x58\x23\x83\x95\xc8\x53\xd5\x6f\x77\x78\x76\x3f\xb4\x14\x7b\x31\xe4\x0c\xbb\xed\x28\x69\x94\xee\x23\x8e\x4c\xc6\x2b\x75\x7c\xa1\x2c\xc6\xfb\x93\x3a\x18\x12\x70\xff\xa6\x2e\x8c\x5e\x70\x55\xae\x6f\xc4\xf0\x00\xa2\xbe\x03\x00\x00\xff\xff\x88\xfb\xad\x15\xab\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/systemd-user/default/program.service.tmpl", size: 427, mode: os.FileMode(420), modTime: time.Unix(1584460126, 0)}
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
	"templates/launchd/launchd.plist.tmpl":                templatesLaunchdLaunchdPlistTmpl,
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
		"launchd": &bintree{nil, map[string]*bintree{
			"launchd.plist.tmpl": &bintree{templatesLaunchdLaunchdPlistTmpl, map[string]*bintree{}},
		}},
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
