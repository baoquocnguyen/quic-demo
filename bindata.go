// Code generated by go-bindata.
// sources:
// html/index.html
// html/rtt.html
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

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x55\xc1\x6e\xdc\x36\x10\x3d\x4b\x5f\x31\xd8\x5e\x5a\xc0\x12\x63\x7b\x93\x38\x2a\x2d\x20\x75\xdb\xb4\x40\x83\xb4\x58\x07\xe8\xa1\x40\xc1\x25\x47\x12\x61\x8a\xc3\x92\x94\xed\x6d\xd0\x7f\x2f\x28\x6a\x37\xb1\xb3\x69\x8a\x5e\x7b\x92\xf4\xc8\x79\xf3\x38\x8f\x33\xe2\x43\x1c\x4d\x5b\xf2\x01\x85\x6a\x4b\x00\x1e\x75\x34\xd8\xfe\xf2\xf6\xc7\x2b\xf8\x16\x47\xe2\x2c\x03\x69\x29\xc4\x5d\x7a\x2b\x8a\x2d\xa9\x1d\xbc\x2b\x8b\xa2\xe8\xc8\xc6\xaa\x13\xa3\x36\xbb\x06\x82\xb0\xa1\x0a\xe8\x75\xf7\x75\x5a\x1b\x85\xef\xb5\x6d\xe0\x14\xc7\xf4\xfd\x57\x59\x16\x85\x24\x85\x39\x52\xe9\xe0\x8c\xd8\x35\xb0\x35\x24\x6f\x1e\x04\x9c\x1d\x02\x00\x00\xea\x3f\x26\x2d\xab\x10\x45\x9c\x02\x28\x7d\x0b\xef\x66\x18\xc0\x09\xa5\xb4\xed\x97\x04\x19\x9b\xf5\x04\xfd\x27\x36\x70\x5a\x3f\x7d\x8f\x47\xbc\x8f\x95\x30\xba\xb7\x0d\x48\xb4\x11\x7d\x5e\x79\x90\x03\xad\xd8\x1a\x54\x87\x04\x5b\x21\x6f\x7a\x4f\x93\x55\x95\x24\x43\xbe\x81\x2f\x54\xd7\x3d\x51\x17\x7b\xd6\x3d\x7a\x2e\x9f\x3f\x3b\x57\x47\x18\x95\x0e\x9f\xa5\xec\xce\x14\x2a\x7c\x4c\x29\x5e\xac\xd7\xeb\xb3\x4f\x8b\xac\x73\x41\x2a\x4b\x07\xee\x43\x45\x2d\x59\xfc\x07\x31\xfb\xd0\x1d\x86\xcf\xc4\x26\x7f\x29\xa2\xcf\x8e\x65\x7b\xaa\x48\xee\x03\x8b\x00\x38\x5b\xee\x05\x67\xf9\x12\xf1\x74\x3d\xda\xb2\xe0\xc3\xe9\xfb\x7b\x04\x3f\x8b\x1e\x39\x1b\x4e\xdb\x24\x89\x27\x1f\xa5\x11\x21\x5c\xae\x3e\xb0\x77\xd5\xce\x99\x3f\x5a\xdd\x1b\xb3\x08\x4f\xa4\xbf\x6f\xae\x5f\x5e\xbf\xdd\x2c\x11\x00\x73\x22\x1d\x60\xbf\xb5\xf9\x2a\x53\x31\xa5\x6f\x3f\xc1\x7a\x30\xe7\x5f\xd0\x1e\xf6\x36\x5f\x3e\xe4\x5d\x5e\xca\x82\xbb\xd4\x19\xd7\x83\x0e\x70\x87\x5b\x27\x7a\x04\x85\x23\x61\x00\x61\x01\xef\x1d\x7a\x3d\xa2\x8d\xc2\x00\x17\x30\x78\xec\x2e\x57\x43\x8c\x2e\x34\x8c\xf5\x3a\x0e\xd3\xb6\x96\x34\x32\x33\x49\x11\x2a\x69\x30\xed\x45\x36\xeb\xec\x69\xd5\xbe\x22\xd0\xa3\xcb\xb0\x26\xcb\x99\x68\x81\x3a\x88\x03\x82\xa5\x5b\x3c\x42\x8a\xb6\xbe\xd3\x37\xda\xa1\xd2\xa2\x26\xdf\xb3\xf4\xc5\xd2\x79\x56\xd9\x15\xe7\x29\x92\x24\x33\x73\x39\x4f\x8e\x02\x2a\xd8\xee\xe0\x15\x51\x6f\xb0\x2e\x0b\xce\xd2\x99\x96\x83\x11\x44\xbf\x03\x1d\x81\xa6\x78\x02\x3b\x9a\xc0\x22\x2a\x88\x04\x53\x40\x10\x20\x27\xef\xd1\x46\xb8\x1a\x3c\x8d\x08\xb7\xe8\x83\x26\x7b\x02\xc2\x2a\x30\x62\xb2\x72\x48\xc1\x77\x3a\x0e\x20\x20\x38\x94\x5a\x18\x90\x34\x8e\x69\x43\xe7\x69\x9c\x0f\x23\xc9\x06\x32\x58\xc3\xf7\xe4\xe1\xcd\x06\x7e\x3d\x49\x69\x9b\xb2\x28\x78\x1a\x1d\x49\x49\xc1\x5e\x3a\x67\xb4\x14\xa9\x10\x81\x65\xb5\xbf\x2d\x79\x6b\xe1\x1c\xbb\x22\x1b\xd1\xc6\xc0\x5e\x0b\xf9\x66\xf3\x68\x07\x40\x55\x4d\x01\x7d\xa5\x44\x14\x95\xd2\xfe\x92\xc5\xd1\x31\x99\x17\xab\x8a\xbc\xce\xb7\xbc\xea\xc8\x4b\xac\x66\x0b\xc8\x5e\xa6\x67\xbd\x37\xa6\xd6\xd4\xac\xd7\xe7\x27\x1f\x81\x17\x4f\x9e\x3c\x3d\x8a\x3e\x3b\x8a\x3e\x3f\x8a\x5e\x1c\x45\x5f\x40\xb5\x34\xc2\xac\x09\xf6\x46\x3f\xde\x9b\x6a\xc5\x72\xb1\xca\xa2\x48\x75\x34\xda\x4e\xf7\x47\x0a\x29\xff\x77\x05\xc9\x57\x3a\xb5\xed\x70\xd6\x7e\x77\x68\xca\xc0\xd9\x70\x96\xf1\xc9\x2c\xd3\xc2\xe8\xf6\xd0\x53\xcc\xc7\xb8\x6a\xaf\x07\x04\xec\x3a\x94\x31\xb5\x5e\x1e\xe2\xd1\x6b\x17\x52\x07\x71\x66\x74\x9e\x07\x89\xa1\x2c\x78\x9e\x9c\xa9\xce\x7c\xf0\x2c\x3d\xbf\x99\xb4\x89\xa9\xc1\x7e\x4a\x4d\x0e\x57\x8b\xc2\xb9\x43\x5e\x0b\x1f\xd1\xc2\x06\x71\x14\xd6\x9e\x40\x40\xfc\x6f\x73\x42\xc7\x1f\xa6\x6d\x12\x34\xf7\xef\x5e\x04\x67\x79\x26\x73\x96\x7f\xf7\x7f\x07\x00\x00\xff\xff\x4b\x59\x4e\x86\xf6\x07\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 2038, mode: os.FileMode(436), modTime: time.Unix(1463129733, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlRttHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\xcd\x4e\xdd\x3c\x10\xdd\xf3\x14\xf3\xdd\xcd\x57\x44\x73\x9d\xf0\x53\xd1\xd4\x64\x43\x17\x45\xea\x82\x42\xfa\x00\xbe\xf1\xe4\xda\xc5\x3f\xa9\x67\x22\x88\xaa\xbe\x7b\x95\x9f\x5b\x01\x02\x75\xc3\x5d\xc5\x3a\x47\x3e\x67\xce\xcc\xc4\xd2\xb0\x77\xd5\x81\x34\xa8\x74\x75\x00\x20\xd9\xb2\xc3\xea\xdb\xf7\xab\x4b\xf8\x8c\x3e\x4a\x31\x03\x23\x45\x3c\xcc\x27\x80\x4d\xd4\x03\xfc\x9a\x8e\x00\x6d\x0c\x9c\xb5\xca\x5b\x37\x94\x40\x2a\x50\x46\x98\x6c\xfb\x69\xa1\xbd\x4a\x5b\x1b\x4a\x28\xd0\xcf\xd0\xef\x83\xe9\xb3\xd6\xe8\xe3\x5f\x91\x7b\xab\xd9\x94\x50\xe4\x79\xf7\xb0\xbb\x69\xd0\x6e\x0d\x3f\x03\x19\x1f\x38\x53\xce\x6e\x43\x09\x0d\x06\xc6\xb4\x63\x36\x31\x69\x4c\x25\x14\xdd\x03\x50\x74\x56\xc3\x36\xa9\x61\xe7\x09\x20\xc5\x92\x40\x8a\x39\xae\x1c\x63\x4c\xd1\xcc\x71\x55\x1b\x04\x6c\x5b\x6c\x18\x62\x0b\x37\xb1\x0f\x3a\xab\x93\xed\xb2\xda\x7a\x84\x77\x9c\x06\x48\xe8\xa2\xd2\x36\x6c\xff\x3b\x94\xc2\x1c\x57\x63\x0e\xd9\xcd\x1d\x99\x3a\xd6\xa2\xe2\x3e\x21\x81\x82\xa6\x27\x8e\x1e\x9a\x34\x74\x1c\x81\x58\x35\x77\xc0\x46\x31\x58\x02\xb5\x71\x08\x1c\x01\x89\xd5\xc6\x59\x32\xd0\xc4\x10\xb0\x61\x1b\x03\x81\x0d\x90\x67\x05\xdc\xd4\x35\xad\xa1\xbe\xbc\x3e\xaa\xbf\xde\x42\x4f\xbd\x72\x6e\x80\x80\xa8\x09\x4e\xb2\xd3\x85\xbf\x6a\x81\xa2\x47\x36\x36\x6c\x41\x47\xa4\xf0\x3f\xc3\x7d\x4c\x77\xef\xe1\x47\x4f\x0c\x4f\xca\x86\xf2\x70\x6a\x43\x37\x57\x3e\x9a\x2f\xf3\x94\x9c\xaa\xa5\x8b\x92\x4d\x75\x6b\x7d\xef\x14\xa3\x1e\x5d\xa4\x60\xf3\x98\xcc\xc1\xd3\x73\xac\xc8\x5f\x42\xcf\x5e\x44\x0b\x78\x04\x49\xb1\x73\x7e\x52\x82\xae\x96\xe0\x47\x5f\xea\xfa\xfa\x58\x0a\xd6\x8f\x49\x69\xdb\xa4\x3c\x42\xe3\x14\xd1\xc5\x6a\x5c\xa4\x15\x50\x6a\x2e\x56\x86\xb9\xa3\x52\x88\x9f\xbd\x6d\xd6\x8d\x43\x3f\x6e\xc8\xda\xc6\xf2\x3c\xcf\x73\xc1\x48\xbc\xaa\xa4\x98\xaf\x57\x6f\x23\x5b\xec\x47\xf6\x64\x3f\xb2\xa7\xaf\xcb\xbe\x3a\x8b\x71\xb7\xdf\xc6\xfd\x6c\x72\xcf\x46\xea\xcd\x93\x7d\xd8\xa3\xf6\xf9\x1e\xb5\x3f\xfe\x43\x7b\x37\x16\x29\x96\x1f\x56\x8a\xf9\xdd\x92\x62\x7e\xbc\xff\x04\x00\x00\xff\xff\x9f\xe4\x14\x00\xc4\x05\x00\x00")

func htmlRttHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlRttHtml,
		"html/rtt.html",
	)
}

func htmlRttHtml() (*asset, error) {
	bytes, err := htmlRttHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/rtt.html", size: 1476, mode: os.FileMode(420), modTime: time.Unix(1463130242, 0)}
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
	"html/index.html": htmlIndexHtml,
	"html/rtt.html": htmlRttHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
		"rtt.html": &bintree{htmlRttHtml, map[string]*bintree{}},
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

