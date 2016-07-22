// Code generated by go-bindata.
// sources:
// templates/init/django/.dockerignore
// templates/init/django/Dockerfile
// templates/init/django/docker-compose.yml
// templates/init/rails/.dockerignore
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/.dockerignore
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/.dockerignore
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/init/unknown/.dockerignore
// templates/init/unknown/Dockerfile
// templates/init/unknown/docker-compose.yml
// templates/templates.go
// DO NOT EDIT!

package templates

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

var _initDjangoDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\x01\x13\x99\xe9\x79\xf9\x45\xa9\x5c\x80\x00\x00\x00\xff\xff\x57\x31\x5f\xce\x1d\x00\x00\x00")

func initDjangoDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerignore,
		"init/django/.dockerignore",
	)
}

func initDjangoDockerignore() (*asset, error) {
	bytes, err := initDjangoDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/.dockerignore", size: 29, mode: os.FileMode(420), modTime: time.Unix(1468853813, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xcd\xae\xd3\x30\x10\x85\xf7\x79\x8a\x91\x60\xdb\x66\xd1\x27\x40\x25\x2c\x40\xb4\x51\x28\x48\x5d\x21\xe3\x4c\x52\x17\xc7\x63\xfc\x03\x8d\x50\xdf\x9d\xb1\xd3\xd0\xe6\xde\xbb\xb8\xbb\xcc\xf1\xcc\x99\x6f\x4e\x3e\x34\xfb\xcf\x20\xc9\xfc\xa6\x4b\xd9\x9e\x85\xe9\xa9\x28\xde\x80\x43\xab\x85\x44\xc0\x8b\x18\xac\x46\x61\x2d\x08\xd3\xce\xa5\x75\x74\x46\x19\x20\x10\x04\xa1\x34\x39\x08\x27\x04\x35\x88\x1e\x93\x36\x52\x74\x70\xeb\x61\xaf\xba\xd9\x7f\xac\xb6\x07\x50\x1e\x84\xf6\x04\xd1\x63\x0b\x3f\x46\xe8\xa3\x51\x92\x9c\x01\x65\xf2\xfc\x02\x02\xde\x93\xfc\x89\xae\x53\x1a\x8b\x6a\xf7\x0d\xde\xd5\xf5\x03\x4c\x96\x66\xdf\x25\x54\xa2\x17\x06\x70\xb0\x61\x84\x2f\xd5\xb6\xa9\x0e\xdf\x3f\x55\x47\x68\xa3\x53\xa6\x87\x41\x18\xa6\x5c\xdb\x91\xd7\x0d\x5c\xb4\x1e\xfe\x28\xad\xf9\x60\x1f\x75\x48\x28\x69\xd8\x39\x72\x79\xc7\x83\x41\x47\x39\x19\x49\x3c\x4b\x46\x8f\x99\x39\xf1\x79\x30\x88\x2d\xdf\xd4\x71\x10\x56\x59\x36\xf1\x41\x68\x5d\x6c\xf7\xf5\x91\x8d\x7f\x45\xe5\x70\x40\x13\xfc\x3a\x5c\x02\x94\xcc\x5f\x3e\x55\x8b\xe6\xeb\x2e\xcd\x6e\xe6\x61\x58\xad\xa2\xed\x9d\x68\x31\xc9\x2f\x3c\xbb\x67\xce\xaf\xa0\x93\xa4\x35\x67\xc4\x16\x41\xc9\x89\xef\xed\x5f\x8e\xf6\x5a\x4e\xd2\x04\xb7\x90\xe6\xae\x5b\xda\xd7\xb9\x65\xae\xa7\xf7\x7b\xac\xf9\xf9\x7f\x39\x81\x8f\xe1\x44\x66\xb3\xc8\xfe\x81\x83\x4f\x35\xa4\x8c\x8d\xf7\x0b\x12\x3c\xff\x91\x00\xd4\xe5\xef\xf4\xcb\xf3\x9e\x75\xf6\x2f\xfe\x05\x00\x00\xff\xff\x22\xcb\xe6\x65\xb5\x02\x00\x00")

func initDjangoDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerfile,
		"init/django/Dockerfile",
	)
}

func initDjangoDockerfile() (*asset, error) {
	bytes, err := initDjangoDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/Dockerfile", size: 693, mode: os.FileMode(420), modTime: time.Unix(1468853813, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8d\xc1\xaa\xc2\x30\x10\x45\xf7\xfd\x8a\xfc\xc0\xcb\x8b\x36\x82\x04\xba\x92\xae\xdc\xa9\x1b\x57\x92\xb4\x43\x09\xa6\x99\x92\x4c\x6b\xfd\x7b\x13\x68\xbb\x10\xdc\xcd\xbd\xf7\x70\xe6\x05\x46\x15\x8c\x99\xd1\xba\x56\x31\x9e\x4e\xf0\x93\x0d\xe8\x7b\xf0\x94\x17\xc6\xfe\xd8\xb5\x3e\x5d\xea\xdb\xe3\x5c\xdf\x53\xe1\xb4\x01\x17\xd7\xa9\x41\x3f\xe1\xcc\x07\x0c\xc4\xa5\x2c\xf9\x10\x90\xb0\x41\x57\x91\x8b\xbf\x91\xf9\x5d\x51\x18\x21\xdb\xac\x7f\x6e\xb2\x56\x93\x36\x3a\xe6\x3e\xd3\x5b\x7f\x14\x4a\x0a\x21\x96\x94\x1c\x39\xee\x8a\x15\xcf\x98\xed\x75\x07\x6a\xf9\xf5\x3f\x60\xa4\x2e\x40\xfc\x16\x1d\x64\xb9\x2f\x3e\x01\x00\x00\xff\xff\x25\x21\x30\xfe\xf3\x00\x00\x00")

func initDjangoDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerComposeYml,
		"init/django/docker-compose.yml",
	)
}

func initDjangoDockerComposeYml() (*asset, error) {
	bytes, err := initDjangoDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/docker-compose.yml", size: 243, mode: os.FileMode(420), modTime: time.Unix(1468853813, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x97\x7e\x4a\x92\xbe\x96\x5e\x71\x61\x4e\x66\x49\xaa\x31\x2a\x4f\x37\x2b\xbf\xb4\x28\x2f\x31\x87\x4b\x3f\x27\x3f\x5d\x5f\x8b\x4b\xbf\x24\xb7\x80\x0b\x10\x00\x00\xff\xff\xa0\x04\x95\x56\x4e\x00\x00\x00")

func initRailsDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerignore,
		"init/rails/.dockerignore",
	)
}

func initRailsDockerignore() (*asset, error) {
	bytes, err := initRailsDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 78, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xbb\x4e\xc4\x30\x10\xec\xfd\x15\x2b\xd1\x5f\x7a\x5a\x24\xa8\xe0\x50\x24\x0a\xba\xf3\x39\xeb\x60\xe2\xd8\x96\x1f\x27\xf2\xf7\x6c\xd6\x09\x21\xa7\xa4\xca\x3c\xec\x99\xf1\x73\x7b\x7e\x05\xe5\xdd\xcd\xff\x34\x51\x1a\x9b\x84\x78\x20\x1c\x26\xf0\xce\x4e\x90\xbf\x10\xb4\xb1\x98\xc0\x21\x76\xd8\x81\xf6\x11\xae\xc5\x75\x16\xc1\xb8\x94\xa5\xb5\xe4\x2f\x4e\xf9\x71\x44\x97\xd9\x7f\x43\xd7\xf9\xd8\x28\xa9\x08\x58\xe3\xc8\xa9\x61\xf2\x05\x2e\xcb\xc1\x20\xd5\x20\x7b\xbc\xcc\x64\x84\x1e\xc7\x24\x9e\xce\xef\x9f\xf0\x82\xe3\x9c\x05\xfc\x35\x32\x84\x66\x61\x76\xf2\xc9\x7a\x35\xec\x64\x66\xa8\x06\xbb\x76\xe9\xec\xfa\xcf\x88\xf6\xe3\xed\xbe\xff\x3a\xf8\xbb\xa4\x7c\x3c\x58\xa6\x84\x39\x3d\x86\x88\xb4\x33\xfc\x15\x6a\xe5\x80\x4b\x61\x0e\x5a\x71\x55\xe9\x51\xb5\xe9\xb7\x2d\x15\x57\x2d\x94\xab\x35\x6a\xd3\x2a\xae\xda\x8c\x6b\x60\xd5\x36\xcc\xe5\x23\x85\x1c\x14\x5a\x47\xcc\xfd\x23\xd2\x10\xaf\xf9\x9f\x4e\xd7\x6b\x4f\x7c\x9b\xf8\x0d\x00\x00\xff\xff\x08\xae\x24\x4e\xf0\x01\x00\x00")

func initRailsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerfile,
		"init/rails/Dockerfile",
	)
}

func initRailsDockerfile() (*asset, error) {
	bytes, err := initRailsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initRailsDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerComposeYml,
		"init/rails/docker-compose.yml",
	)
}

func initRailsDockerComposeYml() (*asset, error) {
	bytes, err := initRailsDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initRubyDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerignore,
		"init/ruby/.dockerignore",
	)
}

func initRubyDockerignore() (*asset, error) {
	bytes, err := initRubyDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xcd\x4e\xc0\x20\x10\x84\xef\x3c\xc5\x26\xde\xcb\x43\x98\xe8\x49\x6b\x9a\x78\xf0\x56\x0a\x4b\x6d\x0a\xbb\x84\x9f\x46\xde\x5e\x8a\x35\xb1\x72\x62\x67\xbf\x61\x86\xa7\x69\x7c\x01\xcd\x74\xf0\x97\x8c\x65\xa9\x42\x3c\xb4\x31\x54\x60\x72\x15\xf2\x27\x82\xdd\x1c\x26\x20\x44\x83\x06\x2c\x47\x58\x0a\x19\x87\xb0\x51\xca\xca\xb9\xc6\x17\xd2\xec\x3d\x52\xee\xfc\x81\x64\x38\x4a\xad\x74\x1b\xdc\x46\x8d\xb4\x50\xb9\xc0\x7c\x19\x83\xd2\xbb\x5a\x71\x3e\xc5\x08\x2b\xfa\x24\x1e\xc7\xb7\x0f\x78\x46\x7f\x66\x41\x3f\x52\x85\x20\x2f\xe5\xb6\x1e\x1c\xeb\xfd\xb6\xee\x4a\xab\xd1\xa9\x5b\x7a\xa7\xfe\x2a\x62\x7a\x7f\xfd\xdf\xff\xf7\xc3\x67\xf7\x88\x29\x03\xdb\x7e\x6f\xde\x9f\xe0\xa1\xbf\x23\xbe\x03\x00\x00\xff\xff\x8b\xae\xa0\xae\x2a\x01\x00\x00")

func initRubyDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerfile,
		"init/ruby/Dockerfile",
	)
}

func initRubyDockerfile() (*asset, error) {
	bytes, err := initRubyDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xf0\x6e\xe9\x98\x87\x21\xf8\x30\x5a\x03\xca\xa2\x29\x69\x3a\xed\xdb\x2f\x85\x6d\xb7\xdd\xfa\xf5\xff\xf2\x9d\x38\x0f\x0d\xc0\x9c\x37\x5a\x06\x70\xf6\x0c\xbc\xef\xd3\x61\x80\x61\x65\x68\x71\xd9\x14\x16\x0e\x4f\x94\xce\xa6\xc8\x09\x5d\xd9\x09\xce\x4d\x57\x28\x9c\x05\x92\x4e\xa2\x39\x7e\x0f\x5b\x6b\xd0\x34\x23\xa5\x1a\x06\xe8\x6c\x38\x5e\x7c\xb9\xc8\xa2\xae\xef\xef\x2e\x0a\x2b\x07\xa6\x51\x29\xfd\x57\xae\x32\xaa\x64\x34\xa1\xfe\xfe\x62\x0f\x3f\xf4\xde\xfb\x0f\x99\x5b\xf1\xd6\xbc\x03\x00\x00\xff\xff\xae\x01\x4e\xf5\xc8\x00\x00\x00")

func initRubyDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerComposeYml,
		"init/ruby/docker-compose.yml",
	)
}

func initRubyDockerComposeYml() (*asset, error) {
	bytes, err := initRubyDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initSinatraDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerignore,
		"init/sinatra/.dockerignore",
	)
}

func initSinatraDockerignore() (*asset, error) {
	bytes, err := initSinatraDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xc1\x6e\x84\x20\x10\x86\xef\x3c\xc5\x24\xbd\xcb\x43\x34\x69\x4f\xad\x8d\x49\x0f\xbd\x49\x61\x70\x89\x30\x43\x00\xcd\xfa\xf6\x8b\xac\x9b\xac\xeb\x49\xbe\xf9\x86\xff\xe7\x63\xe8\xbf\x40\x33\xad\x7c\x95\xd9\x91\x2a\x49\x09\xf1\x56\x49\xdc\x80\xc9\x6f\x50\x2e\x08\xd6\x79\xcc\x40\x88\x06\x0d\x58\x4e\xf0\xbf\x90\xf1\x08\x8e\x72\x51\xde\x57\x7f\x21\xcd\x21\x20\x95\xe6\xaf\x48\x86\x93\xd4\x4a\xd7\x83\x77\x54\x4d\x0b\x1b\x2f\x30\x1e\x8b\x51\xe9\x59\x4d\x38\xee\x30\xc1\x84\x21\x8b\xf7\xfe\xe7\x0f\x3e\x31\xec\x59\xd0\x3e\xa9\x62\x94\x07\x39\x8d\x3b\xcf\x7a\x3e\x8d\x1b\xa9\x35\x9a\x75\x4a\x6f\xd6\x33\x11\xc3\xef\xf7\x6b\xff\xc7\x83\xf7\xee\x09\x73\x01\xb6\xed\xbf\xee\xde\x83\xbb\x76\x8f\xb8\x05\x00\x00\xff\xff\x9c\x51\x49\xbe\x2d\x01\x00\x00")

func initSinatraDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerfile,
		"init/sinatra/Dockerfile",
	)
}

func initSinatraDockerfile() (*asset, error) {
	bytes, err := initSinatraDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initSinatraDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerComposeYml,
		"init/sinatra/docker-compose.yml",
	)
}

func initSinatraDockerComposeYml() (*asset, error) {
	bytes, err := initSinatraDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\xe1\x02\x04\x00\x00\xff\xff\x9c\x10\x28\x7b\x0a\x00\x00\x00")

func initUnknownDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerignore,
		"init/unknown/.dockerignore",
	)
}

func initUnknownDockerignore() (*asset, error) {
	bytes, err := initUnknownDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4d\x2a\xcd\x2b\x29\xb5\x32\x34\xd3\x33\x30\xe1\xe2\x72\xf6\x0f\x88\x54\xd0\x53\xd0\x4f\x2c\x28\xe0\x02\x04\x00\x00\xff\xff\xf1\xa3\x65\xfc\x1f\x00\x00\x00")

func initUnknownDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerfile,
		"init/unknown/Dockerfile",
	)
}

func initUnknownDockerfile() (*asset, error) {
	bytes, err := initUnknownDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 31, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\xcc\xcc\xb3\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\xa6\xe1\xc1\x85\x11\x00\x00\x00")

func initUnknownDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerComposeYml,
		"init/unknown/docker-compose.yml",
	)
}

func initUnknownDockerComposeYml() (*asset, error) {
	bytes, err := initUnknownDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 17, mode: os.FileMode(420), modTime: time.Unix(1466784924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9b\x5b\x6f\x23\xc7\x72\xc7\x9f\xc5\x4f\x31\x47\xc0\x39\x90\x82\x3d\xd2\xdc\x2f\x02\xfc\x72\x76\x1d\xc0\x0f\xb1\x01\x5f\x1e\x82\x6c\x60\xcc\xa5\x47\x61\x2c\x91\x1b\x92\xb2\x57\x5e\xf8\xbb\xa7\x7f\x55\x35\x22\x57\x1c\x52\x12\x25\xc1\x9b\xcb\x02\xb3\x22\x7b\xba\xab\xab\xba\xab\xfe\x75\xe9\xe6\xf9\x79\xf0\x76\xde\xb9\xe0\xd2\xcd\xdc\xa2\x5e\xb9\x2e\x68\x6e\x83\xcb\xf9\xdf\x9b\xe9\xac\xab\x57\xf5\xd9\xc4\x77\x58\xce\x6f\x16\xad\x5b\x5e\xf0\x79\xe5\xae\x3f\x5c\xf9\x7e\xcb\xf3\xe9\x6c\xba\x3a\xef\xfe\xb3\x9e\x5d\xce\xcf\xcf\xba\x79\xfb\x8b\x5b\x4c\x2f\x67\xf3\x85\xdb\xdd\xed\x9d\xf4\xea\xa7\x57\x7b\xfa\x28\xa5\xbf\xb7\xf3\xeb\x0f\xf3\xa5\x3b\xbb\xbd\xbe\x1a\xe9\xbb\xa8\xa7\x57\xcb\x07\x67\xd5\x5e\x7b\x27\xd5\x2e\x8f\x9b\xf3\xa6\xb9\x7d\x78\x4a\x3a\xed\x9f\x91\x1e\x8f\x9a\x70\x39\x9d\xd5\xab\x45\xfd\xe0\x9c\x43\xbf\xbd\xd3\x0e\x9d\x1e\x35\xf3\xcd\xec\x97\xd9\xfc\xb7\xd9\x83\x33\x0f\xfd\xf6\xce\x3c\x74\x7a\x68\xe6\xbb\x4f\x67\x97\x73\xde\xbc\xfb\x2e\xf8\xf6\xbb\x1f\x83\xaf\xdf\x7d\xf3\xe3\x5f\x26\x93\x0f\x75\xfb\x4b\x7d\xe9\xd6\xfd\x27\x93\xa9\x27\xb4\x58\x05\x27\x93\xa3\xe3\xe6\xd6\xb7\x1c\xfb\x0f\x50\x5f\xb8\xe5\xf2\xfc\xf2\xf7\xe9\x07\x1a\xfa\xeb\x15\x7f\xa6\x73\xfd\xff\x7c\x3a\xbf\x59\x4d\xaf\xf8\x32\x97\x01\x1f\xea\xd5\x7f\x9c\xc3\x39\x1f\x68\x58\xae\x16\xd3\xd9\xa5\xbc\x5b\x4d\xaf\xdd\xf1\xe4\x74\x32\xe9\x6f\x66\x6d\x60\x16\xf1\xbd\xab\xbb\x13\x3e\x04\xff\xf6\xef\x4c\xfb\x26\x98\xd5\xd7\x2e\xd0\x61\xa7\xc1\xc9\xd0\xea\x16\x8b\xf9\xe2\x34\xf8\x34\x39\xba\xfc\x5d\xbe\x05\x17\x5f\x05\x70\x75\xf6\xad\xfb\x0d\x22\x6e\x71\x22\x6c\xf3\xfd\x1f\x37\x7d\xef\xbf\x43\xf6\xf4\x74\x72\x34\xed\x65\xc0\x5f\xbe\x0a\x66\xd3\x2b\x48\x1c\x2d\xdc\xea\x66\x31\xe3\xeb\x9b\xc0\x8b\x74\xf6\x35\xd4\xfb\x93\x63\x08\x05\x7f\xfd\xaf\x8b\xe0\xaf\xbf\x1e\x2b\x27\x32\x97\xa7\xf1\xc7\x64\x72\xf4\x6b\xbd\x08\x9a\x9b\x3e\xd0\x79\x74\x92\xc9\xd1\xcf\xca\xce\x57\xc1\x74\x7e\xf6\x76\xfe\xe1\xf6\xe4\x6f\xbe\xcf\x1b\xcf\x9b\x1f\xd5\x5e\x7d\x3d\x70\x7a\xf6\xf6\xca\xef\xd3\x89\x17\xff\x85\xf8\x81\x8c\xd2\xdf\x41\xc8\x77\x54\xbe\xad\xd1\xb3\x75\xf6\x0f\x58\x3f\x39\x7d\x43\x8f\x89\x7f\xb7\xba\xfd\xe0\x82\x7a\xb9\x74\x2b\x96\xfc\xa6\x5d\x41\x45\xe4\xb3\xfd\xf0\xd3\xcc\xfa\x79\x10\xcc\x97\x67\xff\xec\xb7\xf5\x1b\xff\xe5\x6e\x9c\x6d\xe1\xd0\xbe\x41\x41\xf6\xd0\xff\xd3\x6d\x9c\x1c\x2d\xa7\xbf\xcb\xf7\xe9\x6c\x95\xa7\x93\xa3\x6b\x20\x32\xb8\x23\xfa\x2f\xfe\xab\x34\xfe\xe8\x35\x24\x40\x4d\xce\xf8\xc4\x3c\xa2\x2a\x27\xfd\xf4\xfe\x5c\xa7\xc1\xb7\x7e\x8a\x93\x53\x9b\x81\x39\x4d\xca\x7e\x7a\xc6\xec\x7e\xf0\xee\xb1\x3f\x78\x76\xfc\x58\xe1\xe6\xf3\xa1\x30\xba\x77\x28\xbc\xfa\xa1\x1b\x9c\x7f\x4e\x00\xd1\x1e\x22\x80\x70\x9e\xc6\x9d\xa0\x5b\x14\x4c\xfa\xdd\x44\xbe\x59\xbe\x9b\x2e\x3c\x89\x66\x3e\xbf\xda\x1c\x5d\x5f\x2d\x1f\x90\xfc\x76\xa9\x82\x7b\x7c\xa9\x5b\xf7\xe9\x8f\x8d\xd1\xa6\x12\x68\xf9\xcf\x40\xcd\x3b\xf1\x20\xef\x36\x30\xcb\x2b\xb9\x6a\xc5\xc9\xf1\xfb\x8f\x51\xff\xfe\x63\xd9\xbc\xff\x18\x96\xfe\x09\xed\xa9\xde\x7f\xcc\x9d\x6f\xb7\xb6\xde\xf7\xe9\xe2\xf7\x1f\x53\xdf\xaf\xf5\xed\xad\xff\x1e\xf3\xd9\x3f\xb5\xff\xec\xc2\x8d\xf7\x9d\xbe\x73\xc9\x46\x1b\xfd\x5b\x4f\x2b\xf2\xf3\xf9\xf6\xca\xd3\x77\xfe\x29\xfc\xd3\xfb\x27\xcd\x3c\x1d\xff\x37\xf3\x7d\xca\x70\x83\x0f\x9b\x9b\x27\x2b\xde\x7f\x4c\xfc\xf8\xac\x57\x1e\xa2\x6e\xb3\xdf\xf1\x80\x47\xe3\x12\x9b\xbd\x8c\xe1\xd0\x60\x55\x1b\x38\xe6\x0d\x70\xc7\xca\xbd\xf1\xaf\x8e\x77\xba\xf8\x63\xff\xfa\xf4\x4e\xdd\xc7\x29\xc0\xc4\x3f\x89\xa5\x6e\x32\x21\xa6\x7a\x87\x87\x7b\x65\x78\x08\x77\xee\xe0\x42\x0c\xde\x53\xbb\xa7\x3c\x9f\x30\xab\x8b\x60\x9f\x14\x01\xe6\x73\x11\xc4\xd5\x9b\x00\x3b\xb8\xd8\x34\x93\x93\x34\x0e\x4f\xa5\x1d\xed\xbe\x50\xed\xff\x69\x36\xfd\x78\x12\xa5\x79\x15\xa6\x71\x5a\xfa\x61\xe1\xa9\x07\xb6\x9a\xd9\xff\x26\xb2\x7e\x12\x01\x2f\x02\x93\x13\xd6\x2e\xe4\xff\x3f\xee\x36\xa0\x7e\xb3\x57\x73\x71\x46\x07\xe9\x6d\xe9\x75\xaa\x8a\x54\x2f\x6b\xff\xae\xf3\xfa\x97\xf8\x77\x91\x7f\x4a\xaf\x77\x7d\xa1\x7a\x58\xd6\xda\x2f\x47\x97\x3d\xdd\x3c\xf7\x7f\xfd\xf7\xd8\xbf\x4f\x7d\x5b\x9c\xa9\x0e\xf3\xb9\x49\xbd\x1e\xf2\xce\xcf\x93\xfa\x27\x43\xe7\x23\xd5\xf9\xd4\xf7\xc9\xbc\xde\x47\x7e\x5c\xeb\x9f\xdc\xb7\xf5\xe8\xbe\xff\x5b\xfa\x7e\x19\xf4\x3d\x5f\x95\xff\xdc\x44\xca\x4f\xe7\xdb\x1c\xf3\x79\xfe\x1a\x3f\x77\x53\xea\xdf\xd6\x8f\xeb\x23\xfd\x8b\xcd\xe4\x7e\x5c\xea\xfb\x24\x3c\x9e\x87\x7e\xb0\x2d\x3f\xbe\xad\x74\x9e\x3a\x57\x9b\xeb\xfc\xf7\xca\xa9\x8c\xd8\x1a\xf6\x05\xbf\xcc\x81\x8d\xa5\x7e\xde\xba\xd1\xf7\xa9\xa7\xd5\x86\xba\x9e\x91\xef\x53\x23\xa7\xa7\x93\x23\x63\xa7\x6b\x0c\x9f\xd8\x5d\xed\xfb\x17\x3c\xa9\xf6\x89\x2a\x9d\x9f\xf5\x0c\x7d\x5b\x1d\x29\x6f\x49\xa5\xe3\x58\x3f\xda\x93\x4c\xf7\x25\xf2\x34\x2a\xf6\x20\xd7\x75\x82\x4e\x81\xfc\x8d\xce\x07\x9e\x34\xb5\xf2\x5f\xf4\xca\x4b\xe3\xfb\x86\x85\xae\x1d\xe3\x4b\x64\xcf\x95\x2e\x7b\xc4\x1a\x87\x7e\x7c\xd2\x2b\x4f\x0e\x19\x12\xdd\xa3\xca\xcf\x51\x19\xf6\xe4\xec\x77\x6c\xfb\x11\xeb\xd3\x19\x3f\xb4\x95\x95\xea\x48\xe6\xbf\x47\xb5\xae\x47\x5e\xab\x8e\x84\x9d\xf6\xed\xa0\x91\xe9\x7e\xb2\xd7\x55\x6e\xba\xd2\xab\x8e\x64\xac\x81\xed\x7f\x98\xab\x6c\x4d\xa8\xb2\xc1\x77\xdc\x2b\x0d\x64\x62\x4f\x42\xa7\x63\xe1\x3d\x63\x2f\xd0\x99\x81\xff\x44\xf7\xb3\x44\x07\x23\xdb\x9b\x5c\x71\x12\x1d\x45\x5f\x3b\xe3\x8d\x36\xf4\x92\xf5\xe9\x9d\xee\x75\xdd\x29\xbe\xa2\xd3\xd8\x0b\xfb\x86\xbe\xf2\x2e\xf7\xed\x5d\xa9\xfb\x54\xc4\x6a\x03\xe8\x6b\x99\xe8\x5c\xf0\xc1\x3b\xf6\x37\xf5\x4f\xd2\xaa\x5e\xb1\xbe\x65\xaf\xfa\xc8\x7b\xf4\x93\xb1\xd8\x14\xfb\x8b\xbe\x20\x4f\xc7\xbe\x46\xaa\x17\x19\x3c\x57\xba\xe7\xf4\x87\x7e\x6e\x76\x93\xb7\xba\xbe\xac\x29\xf2\x60\x23\xec\x3b\x3e\xc1\x65\xba\x7e\xd8\x5c\x64\x7b\x94\xd4\x2a\x2b\x7b\x57\xa5\x6a\x1b\xf8\x04\x6c\x82\xf5\x63\xcf\xb0\x25\xf4\x29\x76\x6a\xf7\x60\x82\x33\x7d\xce\x6c\x5d\xd8\x23\xd7\xa9\x1d\xc2\x0b\xbe\x05\x1b\x62\x7f\x90\x15\xfb\xcb\x0b\xd3\x79\xf4\x30\x54\x3d\xa9\x4d\x97\xe5\x1d\xeb\x9d\xab\x3c\x8c\x45\x7f\x5c\xaf\x74\xe1\xa9\x74\xaa\xa7\x59\xad\x76\x8b\x3f\x44\x67\x1b\x3f\xb6\x32\x9b\x17\x7d\xc3\x5e\x6b\xdd\xcb\xa6\x52\x3d\xa5\xbd\x2e\x0c\x9f\x1a\xb5\x83\xde\xfc\x25\xeb\xc3\xda\x97\x91\xee\x85\x8b\xd4\x86\xd1\xc3\x06\x3b\x2d\x55\x07\xe4\x3d\xfb\xd9\x2b\xcf\xf0\x0e\x1e\xb2\xc6\xa2\xd3\xd8\x7b\xac\xf2\x82\x95\xac\x3f\xb8\xc9\xde\xb1\xf6\xc8\xd2\xa7\xea\xe7\xfb\x44\xf1\x04\x1d\x42\x26\xd6\x89\x39\xc2\x6c\xdb\x57\xc7\xb1\x8e\x91\x35\x47\xd7\x33\xb3\xb7\xfd\xbe\x1a\x8c\x7f\xbe\xa7\x86\xca\x96\x9f\x5e\xbf\xda\xef\xa4\xe9\x71\x88\x8b\xde\x60\xfd\x35\x1c\xf4\x26\xfb\xe6\x9d\xf3\x2a\xf9\x82\xdc\xf3\x5b\xcd\x5f\xff\xf5\xfa\xea\x20\x27\x8d\x13\x40\x29\x5b\x1c\x80\x37\x84\x36\x5e\x3b\xe9\xd4\x9c\x74\xdf\xa9\x93\x06\x04\x70\x56\x28\x18\xb4\x01\x95\x72\x30\xac\x5a\x01\x5f\x1c\x7d\xab\x00\x1b\x35\x1a\x3c\xd2\x0e\x40\xe2\xf8\xe0\x01\x20\x05\xc4\x68\x07\xc8\xf3\x46\xe7\xc0\xd8\x00\x9b\xdc\x9c\x30\x3c\x40\x0b\x20\x69\xcc\x70\x0a\x33\x5e\x94\x5f\x9c\x60\x66\x81\x46\xa5\x0e\x09\x3e\x00\x3d\x40\x05\xa3\xc1\xf8\x7b\x03\x10\xc0\x1a\x07\xc5\x3c\xd2\x96\x5a\xb0\x90\xab\x41\x01\xca\x18\x8c\x00\x1a\x7d\x6b\x05\x7b\x82\x0b\x01\xfe\x5e\x1d\x03\x46\x8f\x3c\x12\x54\x17\x0a\x1e\xc8\x0b\x18\x25\x06\x0a\x80\x23\x8e\x33\x6c\x15\xac\x6a\x0b\x62\x00\x11\xe4\xaa\xcc\x39\x31\x46\xd6\x28\xd2\x35\x6d\x0c\x0c\xe8\x07\x0f\x99\x39\x1f\x82\x9c\xce\xc0\x08\x10\x62\x1f\x9b\x58\x65\x1d\x9c\x3a\x00\xcd\xda\x24\x16\x6c\x01\x6e\xf4\x8d\x58\x7b\xff\x2e\xac\x95\x0e\x8e\x53\x64\x6f\x15\xcc\x9c\xd3\xfd\x65\x2d\x09\x6a\xaa\x52\x81\x14\xc0\x41\x06\x71\xc4\x95\x3a\x0a\xe4\xc3\xa9\x01\x68\x80\x3c\x0e\x02\xbd\x00\x8c\x49\x0c\xf2\x54\x81\x34\x36\xc7\x10\x46\x23\x20\x95\x29\x1f\xe8\x19\xeb\x0e\xc0\x3d\x22\xa1\x58\x6b\xfa\xf3\xa1\x6a\x4d\x6b\x0b\xb0\xb6\xeb\x42\xfb\x81\x6b\x4d\xea\x10\xf8\xda\x12\xea\x35\x40\x6c\x4c\xa4\x21\xd5\x48\xff\x6c\x30\xfb\x9e\x92\xe7\x8b\x24\xc9\xd8\x21\x81\x52\x5c\x2b\x66\x48\x02\x6c\x41\x14\x4e\x75\xb3\xcf\x58\x32\x0d\xad\x24\xd6\xe0\x18\x9d\x6c\x6b\x0d\xd8\xd1\xe9\xa2\x56\x27\xcf\xdc\x15\x38\xe4\xd4\xa6\x04\xdb\x9c\xda\x54\xe6\x34\xf0\x22\xd8\x01\x77\xe8\xcf\xdc\xe0\x2a\x78\x04\x5f\x12\x10\x14\x3a\x2f\xb6\x0e\x8e\x10\x14\x8a\xbd\x44\x16\x70\x5a\x40\x4d\x00\x2f\xc1\xe9\x10\xd4\x34\xfa\x8e\x71\xb1\x05\x33\x92\xb4\x1b\x56\xde\xb7\xb3\xda\x12\x83\x2a\x53\x9c\x80\xa7\x1d\x76\xb6\xb5\x09\x87\x99\xd8\x16\x99\xb5\x75\x8d\x94\xc8\xb7\xed\x6a\x6b\xfc\x63\x4d\x6a\x17\xff\x2f\x6a\x4d\xa3\x22\x98\x1d\x15\xe5\x53\xcd\x28\x8d\xc2\x38\x8e\xb3\x57\x30\xa3\x83\x33\x76\x32\x02\x32\x3d\x9c\x2f\xca\x82\xe3\x1a\x82\x01\xd7\xaa\x93\xc6\x49\xa0\xbc\x9d\x55\x84\x30\x0c\xa2\x6c\x14\xb2\x2e\xb5\x22\x05\x0d\xbe\xe3\x70\x70\xc8\x18\x93\x38\x9a\x46\x9d\x87\x8b\xd5\xd9\x62\x34\x91\x29\xba\x64\x87\x85\xf2\xd0\x99\xa1\xe1\x24\x88\xc4\x71\x36\x18\x23\xd1\xb2\x33\x67\x2f\xd9\x4f\xa2\xce\xbc\xb0\x68\x58\x8c\xbc\x53\xe3\x47\x9e\xc4\xaa\x05\x64\x94\x55\xa3\x59\x60\x51\x5a\xd6\xec\xe9\xc4\xb9\x06\x02\xad\x05\x0a\xe2\x94\x7b\x95\x17\xc3\xa4\x4a\x80\x33\x24\xe8\x81\x97\x22\x57\xbe\x89\xee\x89\xc4\xa3\x48\x03\x9e\xd6\x32\x73\x9c\x20\x81\x14\x55\x84\xca\x9c\x3e\x01\x85\xb3\xec\x53\x82\x91\x4a\x41\x82\xac\x01\x20\x19\xb2\x6d\x82\x23\xc0\xa9\xb7\xec\x50\xb2\x48\xcb\xe4\xc9\x9c\x9c\x65\x17\xb4\x01\x30\xac\x17\x32\x4b\xa6\xd5\xe8\xbc\xc8\x45\x86\xcb\x5f\xd6\xa5\x69\x35\x10\x01\xc0\x4a\xcb\x50\xc9\x34\xd9\x1b\x82\x23\xe4\x26\xd3\xa8\x7a\x5d\x07\x82\x39\xe6\x91\xac\x3d\xb3\xac\x2f\x53\x50\x2c\x2d\xcb\xc9\x6d\x1d\x98\x1f\x70\x65\xff\xe9\xe3\x6c\x1e\xfa\xa0\x07\x64\xe0\xe8\x11\x7d\xa9\x82\x88\xfe\xd4\x0a\xba\xe8\x1c\xeb\xc7\x5c\x04\x1a\x00\xa8\xec\x17\xba\x62\xc1\x1f\x7b\x5e\xb4\xba\xe7\x99\x65\xdf\xa5\x55\x37\x12\x02\xd2\x5c\xe9\xb0\x4f\x89\x65\xa8\xe8\x26\xc1\x0a\x7a\x20\x95\x16\xab\x6c\x90\xe9\xa3\xa3\xac\x7b\x66\x55\x19\xf4\xa2\xb1\xcf\x00\x69\x6c\xba\x9d\xd9\x67\xd9\xbf\xc6\x32\xce\x5a\x03\x2a\x00\x5c\x02\xdc\x5a\x83\xce\xdc\xf4\x99\x35\x27\x00\x63\xad\x99\x3b\x4e\xb4\xba\x83\x8e\x11\x7c\x49\x16\x5a\x58\x55\xa9\xd3\xf7\x38\x23\xf4\x0c\xc7\xc3\x9e\x57\xa6\x1f\xe8\x00\x74\x71\x14\xc8\x8f\xde\xb2\x26\x61\xb7\x0d\xf0\xe8\x06\xfc\xb0\x9f\x43\xa0\xbb\x0e\xb8\x76\x01\xfc\xe1\xc9\xde\x3d\x22\xf7\xc1\x7d\x5f\xaa\x77\x6f\xe8\x01\xb8\xfe\x5a\x89\xde\x36\xef\x06\xe9\x69\x95\x3f\x15\xd3\x93\xa2\xac\xca\x24\x7a\x05\x4c\x7f\x66\x9a\x37\x84\x1e\x69\xa7\x69\x83\xb3\x30\x07\xa4\x96\xba\x94\x85\x4c\x68\x28\x21\x0b\xc8\xc7\x7b\xea\xa3\xd4\xdb\x40\x00\xd0\x96\x9a\x05\xe9\x46\x66\xdf\x49\x1b\x48\x61\x24\x0d\xac\xd7\x88\xcf\xdf\x76\xa8\x99\x58\x5d\x45\x52\x9f\x4c\x11\x23\xcc\xd6\xe7\x0f\xb1\xb5\x81\xc6\x3c\xa4\x63\x43\x9f\xd4\xfa\xc5\xf6\x77\xa0\x09\x0a\x50\xc7\xa3\x9d\xcf\x20\xb0\xa0\xae\xd5\x01\x79\xb0\x52\xf1\x30\x16\x02\x81\x20\xd4\x42\xa2\x01\x1d\x52\x4d\xd5\x4a\xab\x91\x45\x86\xb4\x95\x7d\x1e\x1e\x52\x1c\x64\x90\xfa\xb3\xa1\xab\x58\x9d\xa5\x7f\xf1\x48\xe8\x85\xc7\x80\x36\xf5\x18\x10\x16\xe4\x7a\x38\xf4\x7a\x6e\x86\x33\x4a\xea\xbe\x95\x3e\x26\xbf\x19\x25\x74\x80\xcd\xbe\x6e\x76\xb3\x5b\x1e\xb3\xe0\x28\x89\xff\x6c\x0b\xbe\x69\x6e\xff\x47\xe5\x36\x3b\x15\xba\xd2\xba\x0c\xf9\x4e\x61\x87\x01\xbb\x14\xfa\x9e\xcc\x07\xea\xf2\x3d\x2a\x1b\x6a\xbc\x75\xf1\x65\x44\x81\xef\x8d\x7e\xb4\xee\x8e\xf3\xfe\xb2\x6a\x3b\xc2\xbf\x29\x6c\x12\xfe\xd9\x59\xc4\x9d\xfc\x07\x27\x11\x72\x94\xdc\x9b\x96\x3a\x0b\x9e\x87\x63\xbf\x54\xe1\x90\xa0\x0e\xc8\x8e\xed\xe8\x0d\x2d\x25\x00\x96\xca\x53\xa9\x5a\x4d\x10\x54\xd5\x1a\xb0\xf7\x76\x94\x81\x5b\x91\xa0\xae\x33\x57\x63\x41\x7d\x65\xc7\x4f\xd0\xca\xac\x9a\x45\xe0\x86\x1b\xe3\x88\x25\x8f\xf5\xd8\x80\xe0\x3e\x37\xc8\x27\xb1\xe0\x28\xa6\xb0\x23\x2a\xf8\xa5\x1a\x2a\x95\xc1\xc2\x02\x3c\x3b\x1e\x4f\xad\x54\x2f\xc1\x66\xa4\xee\x85\xc0\xaf\x30\x77\x42\x30\x4c\x12\x83\x1b\x24\x01\x19\x92\x8b\xd8\xdc\x0c\x47\x01\x04\xc4\x72\x8c\x59\xa8\x9b\x81\x7f\x8e\x78\xc8\xf8\xa9\x04\x12\x90\x62\x95\x3c\x54\xd0\x70\x83\x52\x5d\x8c\x35\x78\x26\xe8\xc5\x7a\xa9\xd6\xc9\x91\x59\xa4\xd6\x9b\x94\x5a\x89\xa8\x8d\x7f\xaa\xae\x9d\x1d\x23\x52\xd9\x95\x23\xd3\xd0\x68\x86\x2a\xaf\xcc\x6b\x41\x28\x6b\x38\x04\xc7\x89\x25\x33\x72\x7c\xd8\x68\xc0\x1e\x6f\x24\x0e\x04\xca\x61\xaf\x7b\x02\x7d\xa9\x44\xe6\xeb\xca\x27\xfc\x95\xe6\x22\x49\x78\x9c\x1d\xb7\x31\x9e\x35\x71\x56\xe9\xa5\xcd\x59\x05\xa5\xb6\x64\x52\xae\x21\x34\x8a\x46\x72\xfc\xe2\x14\xed\xd8\x3b\xda\x98\xab\xb7\x23\x2e\x41\x25\x02\xea\x44\xf7\x95\xcf\x72\x04\x56\x69\xb8\xd2\x59\x52\xc8\xb1\xad\x1c\x33\xda\x55\x09\x8e\x48\xa5\x7a\x92\x68\xd2\x11\x26\xdb\x48\x87\x8e\xcb\x11\xd0\xe0\xe6\xeb\xdd\x41\xf5\x67\xd6\xf2\x5c\x9c\xbb\x17\x52\x7f\x7e\x73\x6f\x1f\xc4\x3d\x29\xa0\x1e\x63\xf9\xe5\xe1\x6d\x24\x9c\x8e\xab\x27\x97\x48\x5e\xcd\x19\xbf\xc4\xa1\x49\xab\x39\x37\x66\x5c\xda\xcd\x06\x8a\xfd\x40\x8b\xb3\xe2\xb7\x44\xaf\xa1\x9d\x82\x56\x0a\x6f\x98\x3b\x91\x23\x39\x1d\x63\x30\xbd\xd0\x1c\x32\x27\xa3\xa8\x30\xb0\x24\xc5\xfc\x4e\x55\x19\x53\xc2\xac\x80\x3d\x4c\x10\x88\xc1\x34\x29\xc8\x8b\xda\xc6\x7a\x9a\x0a\x14\x60\x9e\xe4\x8a\x98\x3a\x79\x3c\xc5\x46\x99\xc3\x4e\xc4\xe5\xf4\xbe\xd5\xda\x00\x7c\x92\xd7\x0a\x24\xda\x49\x39\x90\x87\xa3\x07\x0e\xa8\x4d\x44\x76\xd2\x4a\x0e\x4d\x06\x21\x27\xbf\x4e\xa3\x69\xea\x11\xa1\x1d\x54\xd0\x97\xcf\x52\x4b\xe9\x2c\x8f\x2e\xd7\xb5\x04\x39\xec\x71\x2a\xa3\x9c\x7c\x3b\xe5\x15\x93\x07\xea\xc9\x34\xa8\xd7\x60\x9a\xd4\x4d\x22\xe3\x19\x77\x40\x4d\x08\xd9\xa4\xc0\xda\x1a\x94\x95\x4a\xab\xb6\x1a\x06\x0f\xb5\x10\x39\xe4\x68\xf4\xe0\x42\x4e\xec\x23\x5d\xd3\xde\x78\xa2\x3f\x81\x10\x45\x59\x81\x4f\xbb\x39\xd1\xdb\xe9\x3d\x70\x13\xda\xad\x00\x4e\x60\x81\x0f\x6e\x7b\x84\x76\xf2\x0e\xbf\xd4\x9e\xa8\x51\x35\xed\x38\x84\xd4\x76\xf0\x21\x39\x79\x66\xeb\xf4\x50\xb0\xf4\xec\xe0\x7f\x84\xd2\x3d\x38\x79\x54\xe8\x3f\x42\xe6\xe9\xe0\xf2\xca\x81\xff\x2e\x61\x06\xa8\x09\x9f\x1c\x47\xbd\x30\xd4\xfc\xa0\xd7\x9b\xff\xaf\x85\xfe\x23\x62\x1f\xa6\xcc\x23\x84\xd6\xba\x3c\x7a\x11\x7d\x5b\x93\x47\x68\x3c\x56\x91\x77\xcb\xf1\xa2\x7a\xbc\x43\x90\x2f\x25\x19\xf8\x6c\x15\x9e\x9f\x0f\x44\xf6\x2e\xdd\xc8\x07\xf2\x7b\xf9\x40\xaa\x27\xfd\x43\x3e\x00\x38\xe3\x0c\x71\x3c\xc4\xb7\xc4\xb3\xa8\x3d\x60\xdc\xd8\x67\x29\x48\x87\x7a\x05\x08\x87\x9a\x1a\x20\xb7\x83\x93\xcc\x2d\xa6\xb3\x62\x6a\x6d\x71\xa8\x94\xb9\xac\x20\x0f\x1f\xf0\xea\xec\xda\x20\x4e\x87\x38\x91\x92\x0f\xc5\x65\x89\xad\x5b\x2d\x23\xc9\xf5\xbe\xe1\x34\xcf\x9c\x1e\x0e\x21\xb3\xd3\xbe\xce\xae\xd2\xb6\x76\x6d\x50\x0e\x43\x9c\xae\x13\x4e\x43\xae\x7b\xd9\x75\xc6\xd2\x0a\xe5\x72\x8d\xac\x5c\x17\x71\xe5\x16\x40\xa8\x74\x19\xc7\x75\x26\x72\x0d\x9c\x99\xe4\x38\x85\xc6\xc3\xb1\xe5\x2d\xb1\x15\xff\xe1\x4f\xae\x36\xf5\x5a\xd6\x2b\xed\x76\x40\x67\x37\x00\x3a\xbb\x66\x56\x14\x9a\xc3\xe0\xa8\xe5\xf4\xd0\x82\x8c\xc4\xcc\x9b\xf5\x62\x2e\x29\x94\x57\xda\x0f\x07\x4c\xcc\x4e\x9f\x76\xb8\x52\x67\xb7\x16\x7a\x8b\xd7\x9d\xe5\x07\x72\x4d\x32\xd2\xbc\x85\xf1\xe2\xdc\x9d\x1d\x5e\xd8\xf5\xb1\xd0\xae\xb9\xb1\xa7\x89\x5d\xa7\x94\xc3\x04\xa7\x6d\xf0\x44\x7f\xf2\x38\xc9\x03\x33\x95\x41\x62\xfa\x5c\xf7\x85\x03\x18\x72\x00\x67\xf9\x80\xdc\x24\xe9\xd7\xd7\xe9\x68\x83\x67\xf6\x88\x72\x25\x4e\x59\x72\x8f\x4e\xd7\x98\x77\x72\x2d\x33\x59\x1f\x56\xa0\xab\x94\x12\xc7\xae\x57\x11\xcc\xb0\x2e\x83\xee\xc8\x15\xcd\xf1\xdc\x60\xcb\x78\x5e\x00\x08\x3f\xcf\x10\xb6\x7f\x67\xf3\x00\x06\x3e\x25\x4f\xd8\xc5\xfe\xab\xe0\xdf\x48\xb6\x90\x84\xd1\x97\xe4\xc2\xff\xbf\xfc\xfe\xbf\xb5\xfc\xbe\x63\x9b\x5f\xc0\x5a\xc7\xc2\xf0\xdd\x3f\x7b\x7b\xc0\x76\x9f\x1e\x8c\xef\x17\xec\x55\xec\xf8\x8b\x2e\xc5\xff\xa4\xbf\xfb\x7b\xb9\x9f\xe3\xec\xf9\xb9\x0d\xba\x17\xda\xcd\xc8\x31\x1f\x42\xec\x13\xdb\x4d\x44\xb1\xdb\x71\xdd\x1c\x61\xf9\x30\xbd\x1c\x21\xb4\xd6\xc9\xd1\x5f\x57\x6e\xab\xe3\x08\x8d\xc7\xaa\xe2\x6e\x39\x5e\x54\x0d\x77\x08\x32\x68\xe0\xd3\xc3\xe9\xb4\x8c\xb3\x22\x7c\x15\x05\x3c\x38\x9c\x26\x0c\x24\x64\x1b\xee\xcb\x70\x9f\x41\x7e\x0d\xd3\xad\xdd\x08\x6a\x49\xcd\xa3\xb1\xbb\x1b\x84\xd1\xf2\xeb\x1b\xab\x53\xa1\x9e\xb8\x20\x68\x49\x3d\xc8\xee\x64\x08\x6c\x87\x7a\xc9\x95\xbf\xa9\xa9\x73\x6c\xf7\x7a\x76\xa9\x34\x61\x2a\xbf\x9a\xc8\xed\xce\x48\xd4\x3f\x4e\xa5\x0f\x0f\x8b\xb6\xc8\x6c\xab\xf3\xbe\xb0\x68\x6b\xf8\x41\x9a\xfc\x5a\x61\xd1\x98\x04\x43\x58\xf4\xe4\xa8\xe8\x35\x95\xf8\x99\x51\x11\x85\x88\xd4\xa2\x1e\x1e\x7e\x40\xf1\xdf\x01\x00\x00\xff\xff\x23\x49\x69\x80\x00\x40\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 24576, mode: os.FileMode(420), modTime: time.Unix(1468603699, 0)}
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
	"init/django/.dockerignore": initDjangoDockerignore,
	"init/django/Dockerfile": initDjangoDockerfile,
	"init/django/docker-compose.yml": initDjangoDockerComposeYml,
	"init/rails/.dockerignore": initRailsDockerignore,
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/.dockerignore": initRubyDockerignore,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/.dockerignore": initSinatraDockerignore,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"init/unknown/.dockerignore": initUnknownDockerignore,
	"init/unknown/Dockerfile": initUnknownDockerfile,
	"init/unknown/docker-compose.yml": initUnknownDockerComposeYml,
	"templates.go": templatesGo,
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
	"init": &bintree{nil, map[string]*bintree{
		"django": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initDjangoDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initDjangoDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initDjangoDockerComposeYml, map[string]*bintree{}},
		}},
		"rails": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRailsDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRubyDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initSinatraDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{}},
		}},
		"unknown": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initUnknownDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initUnknownDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initUnknownDockerComposeYml, map[string]*bintree{}},
		}},
	}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{}},
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
