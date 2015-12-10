// Code generated by go-bindata.
// sources:
// data/about.ui
// data/properties.ui
// data/viewer.ui
// DO NOT EDIT!

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataAboutUi = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xc1\x6e\xd4\x40\x0c\x86\xef\x7d\x0a\xcb\x57\x94\x4c\xd2\x0a\x09\xa1\x24\x55\xb5\x5d\x22\xb4\xa5\x20\x58\x10\xb7\x68\x32\x71\x13\x93\x64\x26\xcc\x4c\xda\xe6\x91\x78\x0d\x9e\x8c\xd9\x76\x57\x70\x40\x84\x5b\xec\xf8\xff\xec\xb1\xff\xec\xf2\x71\x1c\xe0\x9e\xac\x63\xa3\x73\x4c\xe3\x04\x81\xb4\x32\x0d\xeb\x36\xc7\xcf\xfb\x37\xd1\x2b\xbc\x2c\xce\x32\xd6\x9e\xec\x9d\x54\x54\x9c\x01\x64\x96\xbe\xcf\x6c\xc9\xc1\xc0\x75\x8e\xad\xef\x5f\xe0\x6f\xc6\x45\x9c\x9e\xa3\x78\xaa\x33\xf5\x37\x52\x1e\xd4\x20\x9d\xcb\xb1\xf4\xfd\x55\x6d\x66\x7f\xcd\x72\x30\x2d\x02\x37\x39\xca\x43\x22\x6a\x9e\x33\x07\x4d\x50\x4d\xd6\x4c\x64\xfd\x02\x5a\x8e\x94\xe3\x3d\x3b\xae\x07\xc2\x62\x6f\x67\xca\xc4\xe9\xef\xdf\x8b\x03\xc6\x44\xac\x8c\x8e\x0e\x31\x16\x8f\xbd\x6a\xa2\x30\xe0\x9a\x2e\x84\xad\x95\xe3\x51\xf5\x75\xb7\xb9\x86\x2f\x4c\x0f\x64\xd7\x84\xc7\x67\x63\x91\xc4\xe9\x5a\xad\x32\xe3\x48\xda\x3b\x2c\xae\xc0\xf1\x38\x0d\x04\x4f\x9d\x42\x9e\x15\x58\x92\x0d\x59\xb8\x33\x16\xca\xdb\xf7\xef\xb6\x6b\xb4\x07\xaa\x1d\xfb\x30\x6d\xe7\xfd\xe4\x5e\x0b\xd1\xb2\xef\xe6\x3a\x0e\x34\x61\x7b\x43\x2e\x5c\xcc\x89\xff\x5d\x80\x32\xd3\x62\xb9\xed\x3c\x16\x9b\xd3\x27\xfc\xfc\x01\xe7\x49\xfa\x12\x3e\x2e\x52\xc3\xee\x88\x5c\x3d\x01\x2b\xd2\x8e\x22\xbf\x4c\x61\xb8\x72\xbf\xab\x6e\xde\x6e\xb6\xb7\x9f\xb6\x55\xf9\xe1\xa6\xba\xa8\x92\x35\x80\x9c\x7d\x67\x6c\xd8\xd2\x3f\xda\x66\xe2\xd9\x5b\xc1\x9a\xe2\x0f\x6f\xfe\x0a\x00\x00\xff\xff\xa1\x46\xe8\x98\xcf\x02\x00\x00"

func dataAboutUiBytes() ([]byte, error) {
	return bindataRead(
		_dataAboutUi,
		"data/about.ui",
	)
}

func dataAboutUi() (*asset, error) {
	bytes, err := dataAboutUiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/about.ui", size: 719, mode: os.FileMode(436), modTime: time.Unix(1449785247, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPropertiesUi = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x5d\x6f\x9b\x3c\x14\x80\xef\xfb\x2b\x90\x6f\x5f\x51\xfa\xf1\x76\xdb\x05\xa1\xda\xd6\xae\xaa\x56\x75\xd5\x96\x6d\xda\x55\x65\xe0\x94\x78\x31\x36\xb3\x4d\x93\xfc\xfb\x99\x24\x6d\xb2\xc6\x40\xa8\xa7\x92\x21\x2e\x53\x9f\xc7\xe1\x7c\xf0\x40\x21\xfe\xe9\x34\xa5\xce\x3d\x08\x49\x38\x1b\xa0\xc3\xfd\x03\xe4\x00\x8b\x78\x4c\x58\x32\x40\x5f\x87\x1f\xdc\x37\xe8\x34\xd8\xf3\x09\x53\x20\xee\x70\x04\xc1\x9e\xe3\xf8\x02\x7e\xe5\x44\x80\x74\x28\x09\x07\x28\x51\xe3\xff\xd0\x6a\x8f\xe3\xfd\xc3\x23\xe4\xcd\xe3\x78\xf8\x13\x22\xe5\x44\x14\x4b\x39\x40\x17\x6a\x7c\x46\x30\xe5\x09\x72\x48\x3c\x40\x99\xe0\x19\x08\x45\x40\xba\xf1\xe2\xcf\x05\xa3\xa9\xe5\xc2\xcc\x61\x38\x85\x01\xba\x27\x92\x84\x14\x50\x30\x14\x39\xf8\xde\xc3\xaa\x39\x58\x11\x55\x84\xde\x3c\xee\x5d\x07\xe4\x12\xdc\x11\xe0\x18\x84\x1b\x62\x81\x82\xc3\x3a\x20\x86\x3b\x9c\x53\xa5\x21\x92\x8c\x14\x0a\x4e\x0e\x0e\x36\x90\x68\x44\x68\xec\xcc\x6b\xc6\x30\x75\xe7\x1f\x75\x22\x21\x9f\x2e\x73\x34\xd5\xe6\xdd\xda\x6a\xc3\x2a\x98\x80\x11\x4f\x79\x02\x0c\x78\x2e\xb7\x87\x42\x2e\x8a\x4a\x4c\x48\xac\x46\x28\xd8\x48\x6c\x95\xdc\xea\xb3\x29\x95\x2f\x91\xe0\x94\x42\xfc\x9d\xb0\x98\x4f\xd0\x7a\xf0\x33\x32\x33\x66\x27\xe7\x5f\xa1\x5b\xe6\x66\x9c\x92\x68\x86\x82\x8b\xe1\xc7\xdb\x9b\x4f\x57\x97\xef\x7f\xdc\x5e\x9f\x7f\x3b\xff\x5c\xba\xd5\x46\x02\xe6\x24\x2e\x04\x89\xd1\xd3\xb0\x67\x1e\xbe\x09\x4c\xb1\x48\x08\x43\xc1\xd1\xff\x4d\x28\xc1\x27\xae\xcc\x70\xa4\xcf\x4f\x3d\xac\x47\x4d\xd0\x88\xd3\x3c\x65\x2b\xba\xe6\x8b\x8d\x65\x32\x97\xea\x0a\x87\x40\x0d\xb5\xb2\xaa\x97\x09\x9e\x62\x4a\x12\x66\x38\x4b\xeb\xc0\xd9\x12\x34\x8e\x74\x15\x48\x17\xa9\x5d\xe7\x69\x08\xa2\x16\x96\x6a\x46\xc1\xbc\x56\x54\xb4\x28\xd8\x83\x46\x48\xea\x2e\xf6\xf6\x4a\xf6\xf2\x4a\x37\xf3\xbd\x45\x07\x8c\x6b\xba\xbb\x63\xdd\xde\xed\xb2\x53\x3c\xbb\xc5\x4a\xe1\xa8\xec\x6c\xaf\x2c\x0d\xdc\xa9\x6d\x69\xbd\x5a\x76\x5c\xbe\x57\x32\x69\xcd\x27\xf0\xe9\x75\x85\xcd\xbb\xf6\x12\x83\x29\x81\xea\x03\xc2\xd6\x83\xdd\xb8\x09\x73\x55\xeb\x2b\x0c\x16\xb2\xb8\x16\x35\xc6\x05\xce\xb6\x39\xe2\x5d\x1c\xb9\x1a\x09\xbc\xcc\xc8\x75\x59\x7a\xc3\xe2\x66\xaa\xd3\xce\x6b\x5c\xd2\xdd\x77\xde\xf2\x0e\xb8\x57\x5e\x29\xde\xaa\xf2\xec\x26\xae\x57\x5e\x0d\x68\xab\xbc\xcb\x14\x27\xdd\x56\x5e\xe5\xbf\x0c\xc6\xca\xec\xbc\xf2\x48\xd1\xb4\x5e\x79\x15\x78\xab\xca\xb3\x9b\xb8\x5e\x79\x35\xa0\xad\xf2\xde\x52\xe5\x0c\x61\xaa\x3a\x6d\xbd\xe3\xee\x59\x0f\x53\xd5\x3b\xaf\x02\x6f\xd5\x79\x76\xf3\xd6\x3b\xaf\x06\xb4\x75\xde\x19\x56\xdd\xbe\xcb\xab\x7c\xb4\x6b\x2c\xcc\xce\xfb\x2e\xd6\x3d\xeb\x85\x57\x81\xb7\x2a\x3c\xbb\x81\xeb\x85\x57\x03\x5a\xbf\xbf\x80\xc9\xc6\xab\xd0\x0d\xf4\x5f\x16\xde\x49\xf7\x84\xc7\x74\xcf\x7a\xe1\x55\xe0\xad\x0a\xcf\x6e\xe0\x7a\xe1\xd5\x80\xb6\xc2\xbb\x22\x6c\xdc\x69\xe1\xbd\xea\x9e\xf0\xa8\xee\x59\x2f\xbc\x0a\xbc\x55\xe1\xd9\x0d\x5c\x2f\xbc\x1a\xd0\xfa\x65\xad\xc0\x4c\x46\x82\x64\xdd\x7e\x90\xf7\xba\x7b\xda\x53\x8f\x9d\xeb\xe5\x57\x81\xb7\x2a\x3f\xbb\xb1\xfb\xeb\xf2\x33\x27\x69\x08\xde\x0c\x7c\x12\xf4\x67\xc0\xda\xe2\x6a\xc1\xf7\xd6\x7e\xfd\xfb\x3b\x00\x00\xff\xff\x97\x84\x70\x1d\x31\x2c\x00\x00"

func dataPropertiesUiBytes() ([]byte, error) {
	return bindataRead(
		_dataPropertiesUi,
		"data/properties.ui",
	)
}

func dataPropertiesUi() (*asset, error) {
	bytes, err := dataPropertiesUiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/properties.ui", size: 11313, mode: os.FileMode(436), modTime: time.Unix(1449785257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataViewerUi = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4d\x6f\xdb\x38\x10\xbd\xe7\x57\x08\xbc\x2e\x14\xc7\x09\xb0\xc8\xc1\x56\xb0\x49\xb0\xbb\x45\x91\x22\x68\xd2\xa6\x57\x4a\x1a\x4b\x53\x53\x1c\x95\xa2\xec\xb8\xbf\xbe\x14\xe5\xf8\x93\x8e\x65\xd9\x49\x8b\x20\x37\x93\x9e\x47\xce\xbc\x79\xa4\x66\xd8\xbb\x78\xcc\x84\x37\x02\x55\x20\xc9\x3e\xeb\x1e\x9f\x30\x0f\x64\x44\x31\xca\xa4\xcf\xbe\xdc\xff\xeb\x9f\xb3\x8b\xe0\xa8\x87\x52\x83\x1a\xf0\x08\x82\x23\xcf\xeb\x29\xf8\x51\xa2\x82\xc2\x13\x18\xf6\x59\xa2\x87\x7f\xb1\xf9\x1a\x67\xc7\xdd\x53\xd6\xb1\x76\x14\x7e\x87\x48\x7b\x91\xe0\x45\xd1\x67\xff\xe9\xe1\x03\xca\x98\xc6\xcc\xc3\xb8\xcf\x46\x08\x63\x50\xfe\xb8\x9e\xaa\xec\x0d\x22\x57\x94\x83\xd2\x13\x4f\xf2\x0c\x2a\x9b\x02\x43\x01\x2c\xb8\x57\x25\xf4\x3a\x4f\xff\xba\x8d\x63\x18\xf0\x52\x68\x3f\x05\x4c\x52\xcd\x82\xbf\x4f\x4e\x9a\x42\xc6\x18\xeb\x94\x05\xe7\x0e\x44\x94\xa2\x88\x3d\x3d\xc9\x8d\xb9\x46\x2d\x20\xe4\x6a\xea\xad\x2b\xc2\xff\x81\xc7\xa0\x2e\x8d\x8d\x0d\x32\xb5\xc3\x99\xfd\x8e\x11\xba\x00\x45\x4a\x63\x3f\x12\x54\x80\x1f\x96\x5a\x93\x6c\x0e\xb5\xee\xb3\xe0\xdb\xc7\xab\x6b\xef\xab\x65\xdf\x89\xb2\x11\xcf\xc7\xae\x28\x2f\xe9\x91\x2d\x5a\xb4\x88\xcb\x05\x22\x85\x20\x35\xd7\x58\x45\x95\x9a\xd1\x4f\x32\x43\xd1\x14\x9e\x52\x46\x09\x48\xa0\xb2\xd8\xb2\x6f\xa1\x27\x02\x96\xe7\xaa\xc0\xab\xf8\xa6\x6b\x09\x94\x43\x88\x6b\x1d\x2f\x98\x74\x1c\xc8\x75\xc2\x36\x90\x56\xa7\xcb\xea\x22\x57\x30\x42\xeb\xe7\x0a\xac\x25\x93\x75\x54\x98\x48\x2e\xa6\xb0\x48\x60\x54\x45\xe0\xa5\x5c\xc6\x02\x54\x9f\xdd\x4e\xf7\xbc\xa2\x0c\xa3\xd5\xc8\x36\xd3\xb2\x46\x0d\x66\x3c\x99\x89\xcf\xb5\x8c\x8b\xa4\xcd\x44\xb9\xc9\xfa\x50\x6d\xe2\x60\x67\x2f\x86\x5c\x60\x8c\x48\xfa\xd5\x4f\x16\x24\xe4\x3f\xe5\xc5\x2f\x26\x59\x48\x86\xc2\xe7\x97\xeb\x75\x6a\xc7\x5d\x24\xb8\x55\xe1\x04\x38\x8d\xdb\xc8\x4a\xc2\xa3\x7e\x4d\x49\x7d\x32\xfb\xbd\xcb\x69\xb3\x9c\xaa\x7c\xfc\x76\x29\xad\x1b\xae\x19\x35\xba\xf3\x17\x74\xa6\x8c\x00\x28\x3b\xc8\x27\x60\x8b\xc2\x3e\xdb\x9d\x9c\x1a\x6b\x70\x8b\x3f\xa7\xad\xbd\xee\xf2\x4d\x7a\x6a\x7f\xd2\x36\xeb\x28\x83\x18\xb9\x9f\x0b\x3e\x11\x58\x18\x39\xa5\xe5\x60\x20\xa0\x81\xac\x5e\x5b\x21\x37\x20\xcb\x45\x95\x64\x66\x3c\x2b\x52\x0e\x5f\x2d\x94\xa6\x04\xca\x29\xa7\x51\x55\x60\xed\x02\x9c\x81\xac\x83\xd3\xd1\x1e\xb5\xc2\x9b\x50\x99\x99\x97\xbe\x25\xe4\x45\x85\x55\xf9\xc0\xa3\xa1\x69\x2f\xb6\xe4\xc8\x18\xf9\x55\xd5\xcd\x02\x90\xb1\xdb\x15\x33\xbb\xba\xd4\x41\x74\x5b\x00\x57\x51\xfa\x47\x2b\x77\xea\xe2\xbb\x76\xad\x76\xcd\x05\xa9\xfd\x81\x69\x24\xdf\x92\x76\x97\xf7\x5f\xfa\x73\x39\x07\xeb\xdc\xdf\x45\x8a\x84\x80\xf8\x61\xb1\xb7\x6e\x45\x7e\x93\xe3\x53\x67\xda\x9e\x9c\xa8\xfa\x4e\xfb\xe8\x48\x7d\x8b\xb4\x6f\xfd\x2a\x6d\x64\x68\xf1\x8f\x75\x7f\x6f\xa7\x87\x68\xfe\x8d\x9a\x1d\xab\x16\x8f\x10\x5b\x52\xb1\xdc\x29\xef\xdd\xfd\x2f\x75\xc8\xc6\x65\x8d\x91\xbb\x3f\x5e\x05\x66\x5c\x25\x68\x30\xdd\xb5\x17\x8e\xa6\x69\xbe\xa1\x18\xc4\xe5\x8b\xdd\x89\xda\xb6\x2e\xb7\xf5\x24\x42\xd1\xb6\x60\xbc\x4b\x69\x3c\x5f\x65\xf9\x4a\x3b\x50\x9d\x73\x07\x39\x57\x5c\x93\x7a\x09\x1e\xea\x44\xf9\x85\xe6\xca\xf0\xd1\x3d\xdd\x11\x67\x2e\x9c\x16\x28\x4d\x39\x0b\xce\x76\x04\x85\x64\xb4\x90\xed\x80\xdb\xe9\x79\xe7\x50\x35\xe9\x6b\xa8\xf6\x9f\x90\x4a\xdd\x56\xb0\xd5\x83\x9e\x5d\xe0\x1a\xb9\xa0\x64\x47\xc5\x36\xb8\x03\x7b\x9d\x85\xd7\xdb\x5f\x01\x00\x00\xff\xff\xad\xb2\x6b\x6a\xf1\x15\x00\x00"

func dataViewerUiBytes() ([]byte, error) {
	return bindataRead(
		_dataViewerUi,
		"data/viewer.ui",
	)
}

func dataViewerUi() (*asset, error) {
	bytes, err := dataViewerUiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/viewer.ui", size: 5617, mode: os.FileMode(436), modTime: time.Unix(1449785266, 0)}
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
	"data/about.ui": dataAboutUi,
	"data/properties.ui": dataPropertiesUi,
	"data/viewer.ui": dataViewerUi,
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
	"data": &bintree{nil, map[string]*bintree{
		"about.ui": &bintree{dataAboutUi, map[string]*bintree{}},
		"properties.ui": &bintree{dataPropertiesUi, map[string]*bintree{}},
		"viewer.ui": &bintree{dataViewerUi, map[string]*bintree{}},
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

