// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/launchd/launchd.plist.tmpl
// templates/runit/log/run.tmpl
// templates/runit/run.tmpl
// templates/systemd/default/control.target.tmpl
// templates/systemd/default/program.service.tmpl
// templates/systemd-user/default/program.service.tmpl
// templates/upstart/default/control.conf.tmpl
// templates/upstart/default/process-type.conf.tmpl
// templates/upstart/default/program.conf.tmpl
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

var _templatesLaunchdLaunchdPlistTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xcd\x6e\xda\x40\x10\xbe\xf3\x14\x53\x2b\xc7\xe2\x85\x5b\x15\x19\x47\x34\x50\x29\x2a\x0a\x56\x80\x56\x3d\xa1\x8d\x3d\x75\x56\xd8\xbb\xee\xec\x1a\x6a\x59\xfb\xee\xd5\x1a\x5c\x9b\x10\xa2\x9e\x7a\x1b\x8d\xbe\xbf\x1d\x7d\x76\x70\xf7\x3b\xcf\x60\x8f\xa4\x85\x92\x13\x6f\xec\x8f\x3c\x40\x19\xab\x44\xc8\x74\xe2\x6d\xd6\x5f\x86\x9f\xbc\xbb\x70\x10\x7c\x98\x2d\xef\xd7\x3f\xa2\x39\x14\x99\xd0\x06\xa2\xcd\xe7\xc5\xc3\x3d\x78\x43\xc6\xa6\x45\x91\x21\x63\xb3\xf5\x0c\xa2\xc5\xc3\x6a\x0d\x63\x7f\xc4\xd8\xfc\xd1\x03\xef\xc5\x98\xe2\x96\xb1\xc3\xe1\xe0\x73\x87\xf2\x63\x95\x3b\xa0\x66\x11\xa9\x02\xc9\x54\x0b\xa1\xcd\x70\xec\x8f\xfc\xc4\x24\x5e\x38\x08\x8e\xea\x67\x71\xc2\x41\x90\x88\xd8\x84\x03\x00\x80\x60\x87\x55\xb8\xe0\xcf\x98\x05\xcc\x8d\xc7\xa5\x36\x24\x64\x1a\xd6\x35\x38\x1f\xb0\x76\xe8\xc6\x82\x54\x8c\x5a\x6f\x4d\x55\x60\xbb\x93\x65\x0e\xd6\x06\xec\xc4\xe8\x34\xe7\x72\x2f\x48\xc9\x1c\xa5\xf9\xc6\x49\xf0\xe7\x0c\x75\xdf\xa2\x8b\xf0\x97\x12\x2d\x9f\xd6\x3d\xc8\xeb\x24\x85\x22\x73\xe1\xd5\x91\x57\xef\x51\xf5\x05\xb1\xae\x87\x40\x5c\xa6\x08\x37\x3b\xac\x3e\xc2\xcd\x9e\x67\x25\xc2\xed\x04\x7c\x94\x7b\xb0\xf6\x5c\xbe\xae\x1b\x5c\x23\x73\xcd\xe6\x24\xd1\x73\xaa\x6b\x40\x99\xb4\x62\x01\x7b\x75\xf7\x88\x54\x4a\x3c\x9f\x52\x5a\xba\x3b\x9d\xdd\x87\x13\xf1\xea\x22\x6a\xac\xf2\x9c\xcb\xa4\x89\x79\x9a\xb7\x9c\x52\x7d\x96\xb7\x33\x17\x3f\x01\x7f\x75\x2c\xef\xc6\x9d\xd8\x03\x6b\x5d\xda\xf6\xa0\x2e\x64\xa6\xf1\xb4\x6d\xb1\xc7\x7d\x33\x5c\x7d\x4e\x2f\x63\xf3\x9e\xaf\x88\xc5\x34\x13\x7b\xec\x3f\xc4\x50\x89\xac\x07\x7a\x2a\xe5\xd4\x2c\x14\x4f\xde\x03\xad\x0c\x97\x09\xa7\x64\x59\x9a\x88\x9b\x97\x2b\xdd\xcc\x54\x0a\xd6\xb2\xae\xa6\xec\x9f\x1b\xeb\xb8\x6f\xb4\xb6\xf5\x9d\x13\x29\xfa\xaf\xce\x1b\x8d\xf4\xc8\x73\xbc\x62\x58\x6a\xa4\xb7\x3f\xb4\xef\x8a\x76\x42\xa6\x33\x41\x18\x1b\x45\xd5\x15\x81\xc3\x11\xb6\x4d\x5a\xdc\x99\x5a\x5b\xcd\x80\x35\x3f\x8c\x70\xf0\x27\x00\x00\xff\xff\x17\x6e\x65\x70\xc7\x04\x00\x00")

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

	info := bindataFileInfo{name: "templates/launchd/launchd.plist.tmpl", size: 1223, mode: os.FileMode(420), modTime: time.Unix(1584464125, 0)}
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

var _templatesSystemdDefaultControlTargetTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x0e\xcd\xcb\x2c\x89\xe5\x0a\x4f\xcc\x2b\x29\xb6\xad\xae\x56\x28\x4a\xcc\x4b\x4f\x55\x50\x29\x4b\xcc\x29\x4d\x55\xb0\xb2\x55\xd0\x2b\x28\xca\x4f\x4e\x2d\x2e\x4e\x2d\x56\xa8\xad\xad\xae\x86\xc9\xd4\xd6\x2a\x54\x57\x2b\xa4\xe6\xa5\x28\xd4\xd6\x72\x71\x45\x7b\xe6\x15\x97\x24\xe6\xe4\x40\xcc\x49\x4d\x71\xaa\xb4\xcd\x2d\xcd\x29\xc9\xd4\x2d\x2d\x4e\x2d\xd2\x2b\x49\x2c\x4a\x4f\x2d\xe1\x02\x04\x00\x00\xff\xff\x4f\xc2\xa4\x0c\x6a\x00\x00\x00")

func templatesSystemdDefaultControlTargetTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSystemdDefaultControlTargetTmpl,
		"templates/systemd/default/control.target.tmpl",
	)
}

func templatesSystemdDefaultControlTargetTmpl() (*asset, error) {
	bytes, err := templatesSystemdDefaultControlTargetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/systemd/default/control.target.tmpl", size: 106, mode: os.FileMode(420), modTime: time.Unix(1584336786, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSystemdDefaultProgramServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xc1\x6a\xe3\x3c\x10\x80\xef\x7a\x0a\x11\xfa\xd3\xcb\x9f\x98\x85\x3d\x2d\xe8\xd6\x74\x09\xdb\x6d\x4a\xdd\xd0\x43\x29\x41\x95\xc6\xee\x50\x79\x64\xa4\x71\x1a\x63\xfc\xee\x8b\x94\x8d\x0d\xd9\xdd\x93\xa5\x6f\xbe\x91\x67\x46\x7a\xd9\x11\xf2\xab\x18\x86\xa5\xc4\x4a\xae\x2c\x44\x13\xb0\x65\xf4\x24\xc7\x51\xdc\xcc\x5b\x35\x0c\x97\xe1\x61\x90\x40\x36\x79\x0f\x3a\xf0\xb6\xca\x8a\x6e\x5b\x39\x8e\x2b\xd6\xa1\x06\x16\x25\xfb\xf6\xf9\x1d\x68\x47\x04\x60\xc1\xaa\x1e\xa2\x10\x2f\x25\x84\x03\x1a\x78\x15\xbb\x08\x21\xa7\x75\x11\x42\x3a\xe9\x7b\xf0\x5d\x9b\x49\x9d\x56\x09\x3d\xfb\xf0\x81\x54\xdf\x60\x00\xc3\x3e\xf4\x39\xfa\x79\x82\x7b\x7b\xa6\xc9\x5c\xd3\x01\x83\xa7\x06\x88\xd5\xc3\xf6\xf1\x29\x9b\xad\x0f\xfc\x47\xb0\x3c\x85\xe2\x45\xe0\x16\x1d\xa8\x65\x01\x6c\x0a\x0b\x95\xee\x1c\x17\xe7\x9e\xfe\x6d\xc6\x3e\x1a\x4f\x15\xd6\xc5\xdc\x7f\x1e\x68\xd0\x54\x83\xbc\xfa\x80\xfe\x7f\x79\x75\xd0\xae\x03\xf9\x4d\xc9\x15\xd0\xe1\xb2\x9e\xc5\x30\x64\x4f\x8e\x63\x2a\xec\xb7\x3c\x8e\x0b\x31\xcf\x78\x7d\x04\x53\xb2\x0e\xac\x8a\x37\xa4\xe2\x4d\xc7\x77\xb9\x74\x46\x5e\xc3\x11\x8c\x5c\x6a\xb9\x98\x5a\x5a\xc8\xb4\x34\xbe\x69\x74\xce\xbd\x16\x8f\x10\x73\xaa\x76\x9f\xba\x8f\xe7\x6d\x09\x46\x7d\xf9\x1a\x45\xc9\x9a\xac\x0e\x76\x43\x6d\xc7\x8a\x3a\xe7\x26\xb4\xed\x38\xb1\xd8\x47\xe7\xeb\x89\xae\x43\xf0\x61\x82\xf9\xb3\xb1\x40\x8c\x15\x42\x50\xff\x91\xf8\x81\xce\xfd\xf4\x16\x54\x83\x47\xb0\xe2\x09\x1b\xf0\x1d\xa7\xd7\x90\xfe\x99\xaa\xe3\x13\x4a\x9d\x4d\xaf\x8f\xd0\xa4\xae\xc5\x3d\x1a\xc8\xd2\x19\xcc\x53\x38\xab\x0e\x1b\xe4\xbd\x6f\x81\xf6\x15\x3a\xc8\x17\x79\x97\xd8\xfd\xf6\x76\x73\xb7\xce\xd9\x7f\x73\xe6\x93\x7e\x05\x00\x00\xff\xff\x88\x3d\x27\x0c\xfa\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/systemd/default/program.service.tmpl", size: 762, mode: os.FileMode(420), modTime: time.Unix(1584464031, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSystemdUserDefaultProgramServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x41\x6a\xc3\x30\x10\x45\xf7\x3a\x85\x2e\x10\xfb\x04\xda\x94\xa4\xd0\x55\x4a\x9c\x92\x45\x30\x45\x48\x13\x77\x5a\x69\x64\x46\xe3\xa4\x46\xf8\xee\xc5\x36\x6d\x68\x96\xf3\xdf\xbc\xcf\x3f\xbf\x11\x4a\xab\x4a\xd9\x68\xbc\xe8\xca\x43\x76\x8c\xbd\x60\x22\x3d\x4d\x6a\x7b\x3f\x4d\x29\x8f\xb8\x14\x0d\xe4\xe7\x3f\xad\xce\x0d\xf0\x15\x1d\xb4\xea\x38\xf6\x60\x32\xc6\x3e\x80\xda\xd1\x15\x39\x51\x04\x12\xf3\xba\x3f\x1c\x97\x92\x3e\xb1\xcc\xd2\x3f\xd8\xac\x28\x3f\x80\x67\x0c\x60\x36\x33\xfa\x48\x11\xf4\x34\xd5\x95\x4b\x74\xc1\xae\xce\x63\x16\x88\xbe\x1e\x32\x70\xfd\xe7\x56\x70\x77\xd5\xee\x1b\x5c\x23\x96\x65\xe9\x76\x29\x46\xbb\xce\x3d\x40\x5e\x62\x1b\x6e\x76\xcc\xea\x94\xf8\x0b\xa9\xdb\x22\x83\x93\xc4\xe3\xf2\x7e\x5b\xc3\x77\xff\x9b\xce\x62\x23\x96\xbc\x65\xbf\x1f\xa4\x1f\xc4\x7c\xa6\x81\xc9\x06\xa5\xce\x2f\x94\xc5\x86\xd0\xaa\x93\x25\x01\xff\x34\x9a\x38\x04\xc1\xcd\xbc\xae\x12\xcb\x1d\x88\xfa\x09\x00\x00\xff\xff\xd3\x11\x0c\x91\x6b\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/systemd-user/default/program.service.tmpl", size: 363, mode: os.FileMode(420), modTime: time.Unix(1584464041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUpstartDefaultControlConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xae\x56\xd0\x4b\x2c\x28\x50\xa8\xad\xe5\xe2\x2a\x2e\x49\x2c\x2a\x51\xc8\xcf\x53\xd0\x00\xb3\x52\x53\x14\xf2\x52\x4b\xca\xf3\x8b\xb2\x33\xf3\xd2\x35\xb9\x8a\x4b\xf2\x0b\x40\x92\x45\xa5\x79\x39\xa9\x65\xa9\x39\x0a\xd1\x06\x66\xb1\x5c\x80\x00\x00\x00\xff\xff\x88\x35\x19\x59\x4c\x00\x00\x00")

func templatesUpstartDefaultControlConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesUpstartDefaultControlConfTmpl,
		"templates/upstart/default/control.conf.tmpl",
	)
}

func templatesUpstartDefaultControlConfTmpl() (*asset, error) {
	bytes, err := templatesUpstartDefaultControlConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/upstart/default/control.conf.tmpl", size: 76, mode: os.FileMode(420), modTime: time.Unix(1584463645, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUpstartDefaultProcessTypeConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcb\x31\x0a\x42\x41\x0c\x84\xe1\x3e\xa7\x98\x52\x0b\xc5\xca\xcb\x88\xc8\x63\x0d\xb2\xb0\x24\x21\x89\x82\x2c\xb9\xbb\xac\x95\x76\x3f\x7c\x33\x77\x8e\xe6\xdd\xb2\xab\x60\x4e\x1c\x37\x33\x54\x1d\x56\x9a\x6b\xe3\x88\x5b\xbe\x8d\x51\x45\x14\xb9\x79\x42\x05\xdf\xe8\xf2\xf8\x79\x50\xa4\xda\xb2\x9d\x3f\x65\xf0\x8b\x07\x2e\xa7\xf3\x15\xea\x58\x62\xff\xeb\x3d\x7d\x02\x00\x00\xff\xff\xc1\x84\x51\xbe\x78\x00\x00\x00")

func templatesUpstartDefaultProcessTypeConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesUpstartDefaultProcessTypeConfTmpl,
		"templates/upstart/default/process-type.conf.tmpl",
	)
}

func templatesUpstartDefaultProcessTypeConfTmpl() (*asset, error) {
	bytes, err := templatesUpstartDefaultProcessTypeConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/upstart/default/process-type.conf.tmpl", size: 120, mode: os.FileMode(420), modTime: time.Unix(1584463645, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUpstartDefaultProgramConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xc1\x6e\x13\x31\x10\xbd\xfb\x2b\xde\xa1\x6a\x1b\xa1\xcd\x02\x37\x0e\xed\xa9\x97\x9e\x88\x80\x5b\xa9\xda\x95\x3d\xd9\x58\x24\x1e\x6b\x66\x4c\x89\xda\xfd\x77\xe4\xcd\x96\xb6\x90\x8a\x20\x71\xb3\xdf\xcc\x7b\xf3\xe6\x69\x02\xa9\x97\x98\x2d\x72\xc2\xfd\x3d\xe6\x5d\xce\x18\x86\xa6\x3e\xb3\xb0\x27\xd5\x1b\xdb\x66\x7a\xc4\x52\xd9\x60\x18\x9c\x53\xeb\xc4\xc0\x09\xa7\xe3\x2b\xa6\xfe\x2f\x6c\xb0\x60\x4f\xeb\xcc\xa9\x71\x9e\x84\x38\xe7\x43\x85\xfe\x68\x9d\x39\xa7\x64\x25\x86\x11\x2b\x4a\x52\x7d\x2a\x59\x3f\x41\xbd\x70\xa9\x8d\xce\xaf\x42\x14\x8c\xd8\x1d\xcb\xb7\x98\xfa\x9b\x10\x85\xbc\xb1\x6c\x6b\x5d\x48\x73\x77\x97\x9c\xdb\x05\xe3\x00\xfa\x91\x59\x0c\x8b\x8f\x9f\xbe\x9c\x8d\x8e\xea\x6f\x18\x9e\x55\x3e\xef\x70\xdd\xa1\x57\x68\x04\x2d\x99\x6f\x03\x2d\xbb\xb2\xb6\xf6\xd1\x67\xf5\x7f\x8d\xe3\x63\xcc\x5f\xad\xbf\xe0\xeb\x56\x3d\xa7\x65\xec\xdb\xa7\x4d\x5f\x08\xec\x6b\x70\x80\x0f\xaf\xef\x57\x5d\x93\x1f\xeb\x9e\x37\x9b\x2e\x85\x2a\xfa\xd5\x01\xc0\xf9\x39\xda\xef\x9d\xb4\x6b\x7e\x2e\xd8\xee\xbb\x06\xb5\xc0\xc5\xe6\x6b\xee\x27\xee\xfb\x7f\x23\x93\x48\x25\x3b\x4a\x01\x53\xd2\x2e\xb3\x5a\xb3\x3b\xad\x5f\xe1\x2f\x2e\x2f\xce\x6e\xd5\x3a\x2b\x7a\xf8\x81\xe2\x01\xd4\x0b\x65\x34\x1c\x71\x72\x7a\xf5\xb6\xf9\x70\xfd\x66\x76\x74\x82\x07\xac\xa8\x0b\x68\xd2\xbb\xdb\x1a\x84\x5f\x31\x8e\x16\x97\x17\x98\xac\x4b\x49\xbf\x5b\x3f\x68\xde\x3c\xc7\xb0\x6f\x13\xce\x4f\x8b\xc8\x06\xcd\xf2\xff\x8e\xf9\x19\x00\x00\xff\xff\x33\xff\x89\xa3\xbd\x03\x00\x00")

func templatesUpstartDefaultProgramConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesUpstartDefaultProgramConfTmpl,
		"templates/upstart/default/program.conf.tmpl",
	)
}

func templatesUpstartDefaultProgramConfTmpl() (*asset, error) {
	bytes, err := templatesUpstartDefaultProgramConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/upstart/default/program.conf.tmpl", size: 957, mode: os.FileMode(420), modTime: time.Unix(1584464146, 0)}
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
	"templates/systemd/default/control.target.tmpl":       templatesSystemdDefaultControlTargetTmpl,
	"templates/systemd/default/program.service.tmpl":      templatesSystemdDefaultProgramServiceTmpl,
	"templates/systemd-user/default/program.service.tmpl": templatesSystemdUserDefaultProgramServiceTmpl,
	"templates/upstart/default/control.conf.tmpl":         templatesUpstartDefaultControlConfTmpl,
	"templates/upstart/default/process-type.conf.tmpl":    templatesUpstartDefaultProcessTypeConfTmpl,
	"templates/upstart/default/program.conf.tmpl":         templatesUpstartDefaultProgramConfTmpl,
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
				"control.target.tmpl":  &bintree{templatesSystemdDefaultControlTargetTmpl, map[string]*bintree{}},
				"program.service.tmpl": &bintree{templatesSystemdDefaultProgramServiceTmpl, map[string]*bintree{}},
			}},
		}},
		"systemd-user": &bintree{nil, map[string]*bintree{
			"default": &bintree{nil, map[string]*bintree{
				"program.service.tmpl": &bintree{templatesSystemdUserDefaultProgramServiceTmpl, map[string]*bintree{}},
			}},
		}},
		"upstart": &bintree{nil, map[string]*bintree{
			"default": &bintree{nil, map[string]*bintree{
				"control.conf.tmpl":      &bintree{templatesUpstartDefaultControlConfTmpl, map[string]*bintree{}},
				"process-type.conf.tmpl": &bintree{templatesUpstartDefaultProcessTypeConfTmpl, map[string]*bintree{}},
				"program.conf.tmpl":      &bintree{templatesUpstartDefaultProgramConfTmpl, map[string]*bintree{}},
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
