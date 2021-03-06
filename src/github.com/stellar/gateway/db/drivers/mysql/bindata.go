// Code generated by go-bindata.
// sources:
// migrations_gateway/01_init.sql
// migrations_gateway/02_payment_id.sql
// migrations_gateway/03_transaction_id.sql
// migrations_compliance/01_init.sql
// migrations_compliance/02_auth_data.sql
// DO NOT EDIT!

package mysql

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

var _migrations_gateway01_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xc1\xaf\x9a\x40\x10\xc6\xef\xfc\x15\x73\x84\xb4\x26\x6a\xaa\x69\x62\x3c\xa0\x6c\x5b\x52\x44\x8b\xcb\xc1\x13\xac\x30\xa5\x9b\xca\x2e\x59\x06\x6b\xff\xfb\x06\x9b\xd6\xc7\x9a\xe7\x7b\xef\xc8\xcc\x6f\x86\xfd\xbe\x2f\x33\x1a\xc1\xbb\x5a\x56\x46\x10\x42\xda\x38\xeb\x84\xf9\x9c\x01\xf7\x57\x11\x83\x3c\xc1\x02\xe5\x19\xcb\x9d\xf8\x5d\xa3\xa2\x1c\x5c\x07\x20\x97\x65\x0e\x52\x91\x3b\x99\x78\x10\x6f\x39\xc4\x69\x14\x81\x9f\xf2\x6d\x16\xc6\xeb\x84\x6d\x58\xcc\xdf\xf7\x9c\x6e\xd0\x08\x92\x5a\x65\xfd\xc4\x59\x98\xe2\x87\x30\xee\x74\x36\xbb\x8d\x5d\xb9\xc6\xe8\x02\xdb\x16\xcb\x4c\x50\x0e\xa5\x20\x24\x59\xa3\xc5\x88\x4a\xaa\x2a\x23\xfd\x13\xd5\xa3\x5d\x2d\x09\xea\xda\x07\xc4\x2e\x09\x37\x7e\x72\x80\xaf\xec\x00\x6e\x2f\xc5\xeb\xab\x69\x1c\x7e\x4b\xd9\xb5\x68\x3d\xdb\x1d\x7e\x7b\x8e\x07\x2c\xfe\x1c\xc6\x6c\x19\x2a\xa5\x83\x15\x04\xec\x93\x9f\x46\x1c\xd6\x5f\xfc\x64\xcf\xf8\xb2\xa3\xef\x1f\x17\x8e\x65\xe4\x1e\x15\x71\x23\x54\x2b\x8a\x7e\xd3\x1b\x8d\xa4\xdb\xe4\xc0\xca\xf9\x87\x17\xd4\x4f\xc6\x36\xa0\x3b\x53\xe0\x0d\x98\xcd\x6d\xa0\x3b\xd6\x92\xe8\x61\x16\x6d\x57\x14\x88\xa5\xcd\xfc\x33\xe2\x3f\x77\xc2\xb2\x42\x93\xc3\x51\x56\xbd\xca\xe9\xd8\xbb\x67\x50\x9d\xf1\xa4\x1b\xcc\x2e\xa5\xc9\x81\xf0\x42\xc3\x7f\x19\x6c\xbb\x13\xfd\xed\x0e\x32\xb5\x37\xdd\xe7\xfa\xda\xa4\x9e\x5e\x40\xa0\x7f\x29\x27\x48\xb6\xbb\xe7\x2e\x60\x31\xe8\xda\xb1\x2e\x9c\x3f\x01\x00\x00\xff\xff\x83\xe1\xb3\xac\x4f\x03\x00\x00")

func migrations_gateway01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_gateway01_initSql,
		"migrations_gateway/01_init.sql",
	)
}

func migrations_gateway01_initSql() (*asset, error) {
	bytes, err := migrations_gateway01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_gateway/01_init.sql", size: 847, mode: os.FileMode(420), modTime: time.Unix(1461004393, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations_gateway02_payment_idSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x08\x4e\xcd\x2b\x09\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\x4b\x50\x70\x74\x71\x51\x48\x28\x48\xac\xcc\x4d\xcd\x2b\x89\xcf\x4c\x49\x50\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd5\x54\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x81\x70\x1c\xdd\x40\xa6\x25\x64\xa6\x24\xe8\x80\xf5\x86\xfa\x79\x06\x86\xba\x2a\x68\x20\x9b\xa1\xa9\x60\xcd\xc5\x85\xec\x0a\x97\xfc\xf2\x3c\x02\xee\x70\x09\xf2\x0f\x40\x71\x88\x35\x17\x20\x00\x00\xff\xff\x8e\xff\xcd\xef\xc8\x00\x00\x00")

func migrations_gateway02_payment_idSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_gateway02_payment_idSql,
		"migrations_gateway/02_payment_id.sql",
	)
}

func migrations_gateway02_payment_idSql() (*asset, error) {
	bytes, err := migrations_gateway02_payment_idSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_gateway/02_payment_id.sql", size: 200, mode: os.FileMode(420), modTime: time.Unix(1516147186, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations_gateway03_transaction_idSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x08\x4a\x4d\x4e\xcd\x2c\x4b\x4d\x09\x48\xac\xcc\x4d\xcd\x2b\x49\x50\x70\x74\x71\x51\x48\x28\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\x8b\xcf\x4c\x49\x50\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x33\xd1\x54\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\xf7\xd3\x77\x54\xb7\xe6\xe2\x42\x36\xdd\x25\xbf\x3c\x8f\x80\xf9\x2e\x41\xfe\x01\x18\x16\x58\x73\x01\x02\x00\x00\xff\xff\xa5\x9e\xfe\x52\xa4\x00\x00\x00")

func migrations_gateway03_transaction_idSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_gateway03_transaction_idSql,
		"migrations_gateway/03_transaction_id.sql",
	)
}

func migrations_gateway03_transaction_idSql() (*asset, error) {
	bytes, err := migrations_gateway03_transaction_idSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_gateway/03_transaction_id.sql", size: 164, mode: os.FileMode(420), modTime: time.Unix(1516212470, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations_compliance01_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\x4d\x73\xaa\x30\x14\x86\xf7\xfc\x8a\xb3\xc4\xb9\xba\xf0\xce\xd5\xb9\x33\x8e\x0b\x94\xd8\x32\x45\xb4\x34\x2c\x5c\x85\x54\x42\xcd\x54\x12\x27\x86\x6a\xfb\xeb\x3b\xd0\x96\x2f\xbf\xea\xb4\x3b\x38\x3c\x27\xbc\xe7\x49\x26\x9d\x0e\xfc\x49\xf8\x93\xa2\x9a\x41\xb0\x31\xc6\x3e\xb2\x30\x02\x6c\x8d\x5c\x04\xa1\x95\xea\x95\x54\xfc\x8d\x45\x58\x51\xb1\xa5\x4b\xcd\xa5\x08\xc1\x34\x00\x42\x1e\x85\xc0\x85\x36\xbb\xdd\x16\x78\x33\x0c\x5e\xe0\xba\x60\x05\x78\x46\x1c\x6f\xec\xa3\x29\xf2\x70\x3b\xe3\x74\xd9\x49\xb2\x9e\xe5\x8a\x2a\xb3\xff\xaf\x6c\xca\xa9\x84\x25\x32\x84\x17\xaa\x8e\x7f\xae\x2e\xb2\x8f\x54\x08\x9a\xed\x75\x1d\xa1\x45\x56\x42\x75\x08\x11\xd5\x4c\xf3\x84\xd5\xa1\x88\x6a\x7a\xa4\x79\xee\x3b\x53\xcb\x5f\xc0\x1d\x5a\x80\x99\x4d\xd6\x32\x5a\x80\xbc\x1b\xc7\x43\x43\x47\x08\x69\x8f\xc0\x46\x13\x2b\x70\x31\x8c\x6f\x2d\xff\x01\xe1\x61\xaa\xe3\xff\x03\xa3\xe9\x6b\xbd\x96\x3b\x16\x4d\x9c\x2b\x1d\x09\x9a\xb0\x72\xfa\xbf\xbd\x5e\x63\xfc\x48\x26\x94\x8b\x73\xc4\x26\x7d\x5c\xf3\x25\x79\x66\xaf\x9f\x86\x7b\xfd\x06\x41\x3f\xb2\x9d\x96\x73\x28\x21\xab\x06\x9e\x73\x1f\xa0\xbc\x58\xc4\x30\xbf\x9e\x0e\x88\x6a\x0c\xb3\xfa\xf6\x33\xa1\xc1\x96\xa9\x2b\x95\xc6\x9c\x5c\xb2\x1a\x73\x72\x59\x6c\xcc\xc9\x65\xb7\xe9\x96\xa9\xfc\x70\x9f\x5e\xe7\x17\xf4\xd7\xa2\x90\xe2\x9f\x66\x23\x63\xbb\xcc\xf3\x6d\xeb\xd5\x5b\xc0\x96\x3b\x61\xd8\xfe\x6c\x7e\xfe\x16\x18\xd4\x99\xe2\xe4\x1f\xad\xe7\x1b\x38\x30\xde\x03\x00\x00\xff\xff\xb0\xd9\x8a\xda\x6d\x04\x00\x00")

func migrations_compliance01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_compliance01_initSql,
		"migrations_compliance/01_init.sql",
	)
}

func migrations_compliance01_initSql() (*asset, error) {
	bytes, err := migrations_compliance01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_compliance/01_init.sql", size: 1133, mode: os.FileMode(420), modTime: time.Unix(1461004393, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations_compliance02_auth_dataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x4e\xc3\x30\x10\x86\x77\x3f\xc5\x8d\x8e\xa0\x43\x91\x2a\x21\x55\x1d\xdc\xf8\x00\x8b\xd4\x29\xc6\x1e\x3a\x25\x16\x09\x24\x43\x1d\x30\x67\xe0\xf1\x91\xb3\xd0\x0a\xa9\xe3\xfd\xf7\x9d\xee\xd7\xb7\x58\xc0\xd5\x71\x7c\x8b\x9e\x7a\x70\xef\xac\x34\x28\x2c\x82\x15\xdb\x0a\xa1\x15\x89\x06\xe9\xc9\xb7\xc0\x19\x40\x3b\x76\x2d\x8c\x81\xf8\x72\x59\x80\xae\x2d\x68\x57\x55\x20\x9c\xad\x1b\xa5\x4b\x83\x3b\xd4\xf6\x3a\x73\xb1\xff\x48\xfd\x27\x35\x99\xff\xf2\xf1\x65\xf0\x91\xdf\xac\x56\x7f\x47\x33\xd5\x4d\x47\x3f\x86\x4b\x84\x4f\x34\x34\xdd\xfc\x9f\xfa\x1f\x3a\x5b\xee\x8d\xda\x09\x73\x80\x47\x3c\x00\xcf\xcd\x8a\x9c\x3a\xad\x9e\x1c\xce\xe1\x59\x0b\x7e\x3a\x15\xac\x00\xd4\xf7\x4a\xe3\x46\x85\x30\xc9\x2d\x48\xbc\x13\xae\xb2\x50\x3e\x08\xf3\x8c\x76\x93\xe8\xf5\x76\xcd\xd8\xa9\x1b\x39\x7d\x07\x26\x4d\xbd\xff\xe7\x66\xcd\x7e\x03\x00\x00\xff\xff\xdd\xf4\xfc\xaa\x44\x01\x00\x00")

func migrations_compliance02_auth_dataSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_compliance02_auth_dataSql,
		"migrations_compliance/02_auth_data.sql",
	)
}

func migrations_compliance02_auth_dataSql() (*asset, error) {
	bytes, err := migrations_compliance02_auth_dataSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_compliance/02_auth_data.sql", size: 324, mode: os.FileMode(420), modTime: time.Unix(1528452152, 0)}
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
	"migrations_gateway/01_init.sql": migrations_gateway01_initSql,
	"migrations_gateway/02_payment_id.sql": migrations_gateway02_payment_idSql,
	"migrations_gateway/03_transaction_id.sql": migrations_gateway03_transaction_idSql,
	"migrations_compliance/01_init.sql": migrations_compliance01_initSql,
	"migrations_compliance/02_auth_data.sql": migrations_compliance02_auth_dataSql,
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
	"migrations_compliance": &bintree{nil, map[string]*bintree{
		"01_init.sql": &bintree{migrations_compliance01_initSql, map[string]*bintree{}},
		"02_auth_data.sql": &bintree{migrations_compliance02_auth_dataSql, map[string]*bintree{}},
	}},
	"migrations_gateway": &bintree{nil, map[string]*bintree{
		"01_init.sql": &bintree{migrations_gateway01_initSql, map[string]*bintree{}},
		"02_payment_id.sql": &bintree{migrations_gateway02_payment_idSql, map[string]*bintree{}},
		"03_transaction_id.sql": &bintree{migrations_gateway03_transaction_idSql, map[string]*bintree{}},
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

