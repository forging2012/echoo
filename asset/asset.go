// Code generated by go-bindata.
// sources:
// static/favicon.ico
// view/echo.html
// view/logfile.tpl
// view/terminal.tpl
// view/ws.html
// DO NOT EDIT!

package asset

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

var _staticFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\xd9\x53\x1c\xd7\xb9\x6f\x49\x0f\xb7\xca\x7a\xb8\xae\xfb\x70\x9f\xf5\xa8\xff\xc0\xf2\x53\x92\x37\x29\x95\xb2\x2b\xa9\x72\xec\x38\x89\x12\xc9\x8e\x14\x39\x8b\x1d\x10\x96\x84\x56\x40\x80\x24\x0c\x42\xc3\x8c\x40\x02\x84\xd8\xc4\x2e\x66\x80\x61\xd1\x2c\x1a\x60\xd8\x61\x30\xab\x40\xac\x83\x60\x06\xb1\xcc\xd6\xdd\xa7\x4f\x77\x0f\xfc\x52\xa7\x99\x41\x03\xb6\x12\xc7\xe9\xaa\xaf\x0e\x9c\x9a\x73\x7e\xdf\xf2\xfb\x96\xc3\x71\x07\xb8\x03\xdc\x91\x23\x6c\x3d\xc2\x55\xbe\xcb\x71\xff\xcf\x71\xdc\x51\x8e\xe3\x8e\x70\x1c\xf7\x53\x6e\x67\x5f\xfb\xde\xe5\xb8\xff\x3b\xbc\x23\x3f\xf4\x13\x55\x99\xc9\x3b\xbc\x2c\x1d\x13\x64\x9a\xc2\xcb\x92\x45\x50\xa8\x57\x54\x65\xec\x8a\xa2\x09\xdb\xb3\x88\xaa\x9c\x22\xaa\xf2\x31\x76\x86\x9d\xfd\xb1\x1f\x51\x65\x26\x87\x78\x2a\x1d\x17\x64\x6a\x22\xaa\xbc\x2e\x6d\xa9\xdb\x74\x3b\x0c\xba\x15\x86\xb6\x6e\x87\x21\x6d\xa9\x9a\x44\xff\xa7\xdb\xe1\x6d\x12\x56\xd6\x45\x55\x36\x89\x8a\x7c\x5c\x54\xe5\x43\xff\xa9\x1e\xec\xf7\x82\x4c\x8f\xf2\xb2\xa4\xe3\xa9\xe4\x5f\x0f\xf8\x31\x39\x3b\x83\xde\xe1\x21\x34\xdb\x2c\x78\x6a\x6e\x84\xd9\x66\xc1\xd0\xd8\x08\x5c\x93\xe3\xe8\x72\x0d\xc0\xd2\xd9\x8e\x16\xbb\x0d\x56\x67\x07\x5e\xad\x79\x41\xc2\x0a\x44\x85\xfa\x45\x45\xd6\x89\xaa\x7c\xf4\x87\xe8\xc0\x7e\x43\xb7\xc3\x07\x79\x59\x7a\x8f\xa7\x92\x4b\x90\x29\x26\xe6\x66\x91\x91\x77\x1f\x09\x37\xae\xe3\x4e\xc1\x03\x64\x95\x16\x23\x2d\xd7\x80\x8c\xc2\x87\x88\x4f\x49\xc2\x37\x45\x8f\x90\x5b\x5b\x83\xec\x92\x62\xa4\x1a\xf4\xb8\x9c\x95\x81\x06\xbb\x05\x82\x22\x83\x97\xa9\x16\x1f\x41\xa1\x2e\x41\xa1\xef\x91\xb0\x72\xf0\x6d\x7a\x6c\x08\x3c\xb7\xb8\xb6\x76\x30\x44\xa5\xd3\x82\x42\x3d\xbc\x2c\xb1\x73\x58\xf0\x7a\xf0\x4d\xc1\x43\xa4\xe6\xde\x47\x76\xc9\x63\x98\xfa\x7b\x51\xdd\xe1\x40\x49\x5b\x0b\x92\x0c\x39\xa8\xef\xe9\x86\x65\x7c\x02\xf6\x89\x49\xe4\x9b\x8c\xc8\x2e\x2d\x41\xcf\xe8\xb7\xda\x59\xaa\x2a\x1a\x37\x98\x1d\xec\x4e\xa2\xca\xa7\x01\x1c\x04\xf0\x1d\xfc\x90\x2c\x71\x8c\x63\x41\x89\x78\x42\x54\x02\x13\xa6\x03\x5b\xab\x5a\x9a\x91\x55\x5c\x84\x3b\x85\xf9\x28\x7d\xd6\x8a\x0a\x87\x1d\x8f\x9a\x1a\x90\x9c\x73\x0f\xc6\xee\x1e\x58\x27\x26\xf1\x6c\x74\x02\x55\xcf\xdb\x91\x57\xf1\x04\x4b\xab\x5e\xcd\x6e\xa2\x46\xb0\x77\xf0\xd9\xca\x74\x38\x46\x94\xbd\x3e\x08\xca\x12\xc3\x3f\x1a\x94\x88\x8b\xe1\x05\x25\x02\x3e\x82\xcf\xce\x0d\xbe\x98\xc4\xad\x87\x79\x28\x6e\x6d\x46\x7e\x43\x3d\x0a\x1a\x8d\x48\xcc\xb8\x83\x84\xd4\x9b\xc8\xad\xad\x85\x75\x6c\x02\xed\x53\x33\x30\xf7\x0f\xa0\xa6\xb9\x09\x41\x49\xd4\xec\xd6\xce\xbf\x89\x81\x26\xa2\x2a\xbb\x62\xf9\xc0\x53\x89\x23\x84\x1c\x0c\x4a\x44\xa7\xd9\xcc\x6c\x97\xc8\x8e\xd0\x9d\xf3\x1b\x7c\x08\xc5\xf5\x75\x48\x2f\x78\x88\x8c\x92\x22\xc4\xa7\xdf\xc6\x67\x57\xee\xe2\x83\x93\x67\xf1\xd1\xd9\x44\xe4\x1b\x9b\x60\x74\x76\xa3\xd4\x64\xc4\xfc\xca\x2b\xed\x0e\x86\xc5\xf0\xb5\xdc\x54\xbf\x23\xba\x68\x1c\x82\x44\xe4\x42\x12\x39\x11\x92\x88\x2f\xea\xf7\x10\x8d\xd8\x1f\xc1\x67\xf2\x3a\xe0\x87\xb5\xb7\x17\x19\x05\x45\x38\x75\x4d\x8f\xaf\xb2\x1b\xf0\xc9\x17\x89\x38\x99\xf0\x08\xe7\xae\x65\xc2\x64\xb3\x60\x69\x6d\x55\xc3\x8c\xf2\xee\x6d\x42\x54\xd9\x47\x54\xe5\x84\x14\x56\x38\xb2\xa5\x1c\x16\x64\x6a\x8c\xda\xbb\x1b\xfb\x18\x09\x52\x82\x4d\x22\xc2\x27\x11\xcc\xad\xf9\x50\x65\xed\x41\x66\xf9\x53\xb4\xbe\x58\xc0\x35\x7d\x09\x5a\xbb\xfa\x22\xf7\x2a\xd1\x5a\xf4\x2f\xf1\x23\xbf\x35\x02\x38\x2c\xaa\xf2\xfb\x82\x42\xd7\xa2\xb6\x46\xf5\x8f\xc6\x8b\xed\x05\x28\xc1\x06\x11\xe0\x93\x44\xac\x06\x7c\x98\x5c\x9c\xc7\xbd\x27\xd5\x28\xb5\xf7\xe2\xa1\xb1\x11\xce\xe1\x61\x8c\xce\xcc\xc0\xb3\xb9\x01\xa2\xc8\xa0\x61\x15\x74\x4b\xd5\xf8\xb7\x5b\x1f\xbf\x8b\xbf\x26\x85\x95\xf7\xa5\xb0\x92\xc2\x30\x34\x9f\xcb\x3b\x9c\x8b\xb5\x3d\x24\x33\xfb\x25\xbc\x0e\x06\xe0\x18\x72\xa1\xd2\xd1\x85\x8a\xce\x5e\x94\x77\x0d\xc1\x3c\x36\x0f\xe7\x82\x07\x4d\x2f\x96\xf0\xc0\x31\x84\x72\x8b\x03\xb6\x81\x41\xf4\x8e\x8d\xe1\xe5\xab\xa5\x1d\xfe\x28\x6f\xf8\xbf\x5f\x07\x86\x2d\x85\x15\x8b\x20\xef\x70\x6e\x97\xf3\x91\x98\xb3\xbf\x03\x12\xc1\xa6\xc0\xa3\xb1\x77\x10\x39\xb6\x51\xd4\x4f\xf9\xe1\xda\x10\x31\x1a\xa0\x18\xf0\x51\xf4\x6d\x48\xa8\x1b\xe7\x51\xe0\x14\xd0\x32\x17\x44\xb7\xdb\x83\x3e\xf7\x32\xca\xed\x1d\x98\x59\x5e\x7e\x93\xff\xdf\xc3\x09\x86\x2d\x85\x15\xaf\x66\xbf\x44\xde\xe4\x48\xc4\x5f\x51\xdd\x97\x5e\xaf\xa2\xb4\xdd\x85\x3c\x7b\x08\xfa\x86\x4d\x58\x67\x05\xf4\xac\x09\x30\x2f\x05\xd1\xe0\x16\x61\x68\xf6\x23\xcf\x1c\x84\x69\x52\x44\xdf\x5a\x08\x43\x9e\x75\xd4\x76\x0d\x60\x6a\x61\x61\xa7\x0e\x47\xee\xfa\x1e\x6e\x68\x7d\x8c\x8f\xb5\x3f\xe2\xb3\x37\xbd\x8d\x62\x65\x73\x1d\xe5\x1d\x83\xc8\x73\x04\xa0\x33\xfb\x50\x37\x2d\xc2\xbc\x14\x42\xf3\xfc\x3a\x5a\x3d\x04\x95\xc3\x21\x54\x74\x07\x61\x5b\x22\x30\x58\xfb\x71\x25\xaf\x1a\xba\xb2\x6a\xac\x87\x82\x6f\xee\x52\xf6\xd4\x80\x3d\x12\xcd\xf9\x58\x7c\x21\xa2\x33\x1f\xc9\xbd\x9a\xae\x01\x34\xba\x79\x3c\x70\xce\xa0\x7a\xc6\x87\xfa\x15\x09\x66\x8f\x84\x36\xaf\x84\x42\x47\x1f\x4c\x2f\x5e\xc1\xe9\xf1\xe3\x6a\x5e\x15\xae\x16\x76\xe0\x42\xd6\x63\x4c\xbb\xdd\x7b\xe2\xbf\xdf\xbf\xbb\xf8\xfb\xec\xe7\x23\x39\x18\xd4\x72\x92\x62\x7a\x79\x19\xe5\x1d\xfd\xb0\xaf\x8a\xf8\x24\xee\x22\xf4\xb6\x7e\x34\x78\x65\xb4\x7a\x29\xda\x5f\x4b\xb8\x74\xcf\x80\x62\xa7\x0b\xa6\x89\x05\xc4\x65\x94\xe1\x6a\x7e\x07\x12\x73\xcd\xa8\x6a\xb3\x82\x8f\xf4\x21\x9e\xd9\x12\xb5\x6d\x9f\x0f\xa2\xf1\xe7\xf7\xe5\x7d\x28\xc2\xff\x8e\x91\x71\x54\xf4\x4f\xc3\xb1\x26\xe3\x96\xc9\x86\xf2\x71\x0f\x4c\x2b\x14\x8d\x1e\x19\x56\xaf\x08\xe3\xf4\x2a\xba\x56\x79\x3c\xea\x18\x44\xdc\xdd\x2a\x5c\x34\xb4\x22\xfe\x9e\x09\x09\xb7\x73\xb0\xe8\x5d\xd1\x7a\x61\x6c\x1f\xd8\xe7\x7f\x2f\xdb\x8b\xad\x79\xd1\x5a\xa0\xf9\x40\xa6\x78\x3e\x38\x8c\xf2\xa1\x05\x34\x2c\xcb\xa8\x74\x53\x54\x2e\x4a\xa8\x5e\x24\xda\x5a\xb3\x44\x60\x5c\x96\xd1\xb1\x21\x23\xb5\xe6\x19\xfe\x9e\x51\x85\xb8\xbb\xf5\x38\x9b\x64\xc0\x1f\xe3\x2f\xa1\xb9\xdd\xa1\xd9\x1e\x88\xf2\x7b\x6f\x1e\x78\xb5\xfc\x8b\xd4\xea\x68\x1d\x88\xf6\xbe\x00\xb3\x5f\x91\x51\xd7\xd2\x82\xcc\x06\x07\x6a\xe6\x25\x94\x4d\x89\x28\x79\x29\xa2\x62\x81\xa2\x6c\x9e\xa0\x7c\x8e\x47\xd9\xb4\x1f\x95\x93\x5e\x5c\x2e\xac\xc1\xb9\xeb\x59\xb8\x94\x69\xc0\xb5\xec\x7b\xb8\x96\x99\x81\x4b\xa9\x37\xb1\xe8\xf5\xec\xf6\xb2\xd8\xd8\x47\xf2\x2f\x65\x3f\x7e\xd4\x07\xac\xee\xb1\x9a\x7b\x5b\xaf\xc3\xaf\xcf\xfd\x05\x5f\xdf\x2b\xc4\x85\xfb\x4f\x70\xbd\xa4\x01\x77\x4c\x0e\xa4\x55\xb7\x22\xb5\xb4\x0e\x19\xa5\xd5\x48\xb9\x9f\x8f\x14\xbd\x01\xa9\x39\x3a\x5c\x4c\x4f\x47\x5c\x52\x12\xfe\x7c\xf1\x22\x8e\xff\xf6\x0f\x78\x5c\xf7\x14\x3e\x51\xd8\xd3\x53\xa3\xf5\x27\xb6\xfe\xee\x72\x4e\x22\xda\xea\x97\x08\xc6\xe6\x67\xf1\xd5\x95\x44\x7c\xfc\xa7\xcf\xf1\xd7\xc4\x0b\x48\x48\x4e\xc2\xb9\x84\x78\x5c\xcd\xb8\x05\x7d\xf1\x63\xdc\x32\xe8\x11\x97\x9c\x8c\x1b\xd9\x77\x11\x7f\xe3\x06\x3e\xfc\xf4\x37\x38\x73\xfe\x3c\x7e\x75\xfa\xac\xa6\xd7\xad\x5a\x1b\x2e\xe7\x16\x61\x76\x65\x25\x76\x56\xdd\xad\xbf\x24\xac\x1c\x16\x14\x6a\x64\x76\x33\xcc\x00\x11\x35\x61\xfa\xb2\x9e\xd3\xd2\xd5\x85\x2f\x93\x53\xf0\xb3\x0f\x3e\xc4\xcf\x3f\xfe\x04\x67\xbe\x4e\xc0\xa7\x67\xcf\xe0\x7c\x72\x12\x2e\xa4\xa5\x21\x2e\x39\x05\x27\x4e\x7e\x86\x5f\x9e\x3a\x85\x53\x5f\x7d\x89\x8f\xbf\xf8\x1b\xd2\xf2\x8b\x90\xa8\x2f\xc6\xb3\x89\x45\x14\xf6\xbc\x44\x7a\x85\x09\xe3\xf3\x73\xcc\x5e\xad\x27\xb0\x3e\x15\xed\x3f\x21\x4a\xd8\x0c\x70\x22\x28\x11\x1f\xc3\xf5\x33\x6c\x22\x62\x53\x14\xb0\x21\x8a\xa8\xb3\xd9\x71\xa3\xa8\x1a\x71\x77\xf3\xf1\xbb\xf8\xcb\x38\x73\x39\x09\x1f\x9d\x39\x87\x5f\x9c\x3c\x8b\xcf\x2f\x5c\xc1\x3f\x6e\xa6\xe1\xf7\x89\xa9\x88\xcf\xba\x8f\xc2\x67\x9d\xd0\x35\x3a\x60\x1c\x5d\x44\xc1\xf3\x31\x98\x46\x16\x91\xd7\x39\x85\x5b\xa5\xb5\x98\x5b\x59\xde\xad\x41\x62\x4c\xff\x65\xf3\xc7\x5a\x30\x74\x30\x48\x89\x8e\x71\x34\x20\x89\x3b\xbe\x17\x05\xcd\xff\x6d\xdd\x5d\x38\x9f\x5b\x89\xcb\x65\x76\xa4\xd7\x3a\x90\xd9\xd8\x8d\xf3\xb9\x4f\x70\x3a\xad\x10\xba\xe6\x1e\xe4\x5b\xfb\x70\xbb\xa1\x1b\x55\xc3\xb3\x68\x7a\xe1\x41\xd9\xc0\x12\x2a\x86\xdc\x28\xec\x18\x47\x66\xbd\x1d\x45\x9d\x63\xc8\xa9\x69\x84\x4f\xe0\x63\x73\x5f\x17\x3b\x07\xf2\x32\xd5\xe6\x6d\x3f\x11\x5d\xcc\x07\x1a\xf7\x34\x5f\x10\x8c\xbb\xdd\x48\x2b\xae\xc0\xa5\x52\x1b\xb2\xdb\x5c\x48\xad\x73\xc2\xf0\x6c\x08\x5f\xe6\x99\x71\xbd\xb2\x1d\x99\x2d\x43\x28\xea\x9d\x45\xdd\xd0\x34\xcc\x93\xcb\x68\x78\xe1\x85\x7d\x6e\x0d\xf9\x76\x17\x12\x1f\x94\x23\xc7\xdc\x8e\x26\x67\xcf\x6e\x1f\x12\x15\xd9\x25\x6f\xa9\x47\x69\x58\xdd\x33\x03\xfa\x88\xc8\xe4\x58\x88\x12\x8f\x36\x73\x44\xe2\xb0\x26\xf0\x30\xb6\x77\xe0\xda\xe3\xa7\x88\xcf\xad\xc6\xf9\x9c\x32\xe4\xd9\x46\x90\x5a\xed\x44\xc2\x63\x2b\x6e\x1a\xfb\xa0\xb7\x8d\x21\xa3\xaa\x01\xb5\x7d\x63\xe8\x5c\xda\x44\xef\xab\x4d\xe8\x9b\x7b\x50\x60\xeb\x87\xa1\xbe\x05\xaf\x7d\x9b\x5a\xcc\x79\x99\x7a\x04\x85\xb2\x77\xd4\xf7\xce\xfe\xd2\xb6\x7a\x50\x90\xe9\x69\x9e\x4a\x1e\x16\x03\x1f\x9b\x39\x44\x01\xee\x4d\x1f\x2a\xad\xed\xb8\x5d\x5a\x87\x04\x43\x31\x92\xeb\xfb\x90\x6e\xea\xc7\x37\x4d\x83\xc8\xb2\x8e\x43\xef\x98\x84\xbe\xb5\x07\xcf\xe7\x3c\xe8\x74\x6f\xa2\xc0\x31\x81\x0a\x4b\x3b\x2a\x2d\x36\x4c\xb9\x17\x21\x87\x55\x6d\xfe\x66\x77\x8b\x8a\xfc\xd6\x37\x00\xdb\x27\x5b\x0a\xd3\x41\x7b\x7f\x30\x1f\x30\x3e\x30\x1e\xbc\xe6\x79\x74\x8e\x8c\x42\x57\x55\x87\xfc\x36\x27\x8a\x9c\x13\xa8\x19\x71\xa3\xa0\x6b\x12\x69\xc6\x6e\xdc\xb7\x7e\x0b\xf3\xf8\x22\xca\x3a\x47\xf0\xc4\xda\x81\xc9\xb9\x59\xf8\xf8\xe0\x0e\xdf\x15\xd9\x15\x94\xc4\xf7\x54\xe0\xad\xd8\x7b\xde\x02\x12\xd1\xf8\x10\x92\x88\x4e\x90\xa9\x3f\x18\xa9\x07\x1b\xa2\x80\x85\xcd\x0d\x8c\xce\xcd\xc0\x3e\x34\x8c\xc6\x9e\x01\x18\x9d\xbd\x78\xd2\x66\x45\x53\x57\x37\x7a\xa7\x5e\x62\x66\x65\x45\xeb\xbb\x24\xac\x32\x5c\x3f\x51\x65\x9d\xa4\x2a\x47\xb7\xb0\xfd\x6f\x71\x63\x3f\x96\x17\x3c\x95\xb4\xf7\x67\x80\x88\x26\xbf\x28\xac\xfb\x89\xb8\xcd\xfa\x11\x1f\xa9\x97\xbc\x56\x33\x44\x6d\xe6\x93\x54\x05\x12\x9b\xfb\xc2\xea\x36\x51\x95\x75\x41\xa1\x26\x41\xa1\xc7\x05\x55\x3e\x24\xfc\x17\xef\x60\x36\x9f\x07\x89\xf8\x4e\x80\x88\xc7\x78\x2a\xb1\x5a\x6d\x11\x64\xea\x15\x63\xfa\x3a\x9b\x37\x25\x55\xf1\x52\x55\xb1\xa8\x5b\x5b\x29\xca\x56\xf8\x18\xaf\xd0\x1f\xf4\xfe\xc6\x8d\x03\xcc\x33\xd7\xc3\x1c\xf7\x93\x00\xc7\xfd\xef\x22\xc7\xfd\x0f\x13\x07\xc7\x1d\xba\xc1\x71\x07\x98\xfc\x58\xdd\xa3\xe7\xd9\x5d\xd1\x7b\x19\x06\xc3\x62\x98\x0c\xfb\x9f\x01\x00\x00\xff\xff\xd1\xc0\x86\xcf\xbe\x10\x00\x00")

func staticFaviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_staticFaviconIco,
		"static/favicon.ico",
	)
}

func staticFaviconIco() (*asset, error) {
	bytes, err := staticFaviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/favicon.ico", size: 4286, mode: os.FileMode(420), modTime: time.Unix(1526217292, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewEchoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\xcd\x6e\xe3\x36\x10\xbe\xfb\x29\xa6\xda\x3d\xb4\xc0\x5a\xb4\xbd\x0d\x10\x78\x29\x01\x6d\xd1\x22\x45\xd3\x22\x08\x92\x5e\x0b\x5a\x1c\x9b\x44\x28\x52\x25\x47\x8e\x5d\xc1\x40\xdf\xa1\x6f\xd8\x27\x29\xa8\x1f\xc7\xf1\xae\x5d\x64\x81\x5c\x16\x3e\x48\xe4\x7c\xfa\xe6\xe3\x7c\x23\x8f\xf8\x57\xd2\x15\xb4\xad\x10\x14\x95\x26\x1f\xf1\x78\x01\x23\xec\x2a\x4b\xd0\x26\xf9\x68\xc4\x15\x0a\x99\x8f\x00\x78\x89\x24\xa0\x50\xc2\x07\xa4\x2c\xa9\x69\x39\xbe\x4c\xda\x00\x69\x32\x98\xff\x58\x28\xe7\x38\xeb\x16\x7b\xbc\x15\x25\x66\x89\xc4\x50\x78\x5d\x91\x76\x36\x81\xc2\x59\x42\x4b\x59\x92\x1c\xc3\xd6\x1a\x1f\x2b\xe7\xe9\x00\xf3\xa8\x25\xa9\x4c\xe2\x5a\x17\x38\x6e\x17\xef\x40\x5b\x4d\x5a\x98\x71\x28\x84\xc1\x6c\x1a\x55\x02\x70\xa3\xed\x03\x78\x34\x59\x12\x68\x6b\x30\x28\x44\x4a\x40\x79\x5c\x66\x89\x22\xaa\xc2\x9c\xb1\x52\x6c\x0a\x69\xd3\x85\x73\x14\xc8\x8b\x2a\x2e\x0a\x57\xb2\xfd\x06\x7b\x9f\xbe\x4f\xbf\x65\x45\x08\x4f\x7b\x69\xa9\x6d\x5a\x84\xd0\xc9\x7d\xb5\x34\x63\x52\x58\xe2\xf3\x64\x6d\x0a\x88\xfe\x64\x09\xe1\x86\xd8\x10\x01\x58\x38\xb9\x85\xa6\xbd\x05\x28\x85\x5f\x69\x3b\x87\x8b\x6a\x03\xb3\x49\xb5\xf9\xd0\xee\xef\x46\xed\x65\xe9\x1c\xa1\xff\x08\x3b\x8b\xe0\x49\xfc\x7d\xe8\x23\x95\x90\x52\xdb\xd5\x1c\x2e\x63\x64\xd8\x8d\x79\xc7\xc2\xe8\x95\x9d\x43\x81\x96\xd0\x0f\x91\x85\xf3\x12\xfd\x98\x5c\x35\x87\x69\xb5\x81\xe0\x8c\x96\xf0\x46\x4a\x39\x20\x96\xce\xd2\x38\xe8\xbf\x70\x0e\xd3\xd9\x20\x0b\xa0\x70\xc6\xf9\x39\xbc\xb9\xb8\xb8\x78\xa6\x34\xed\x6d\x3f\xd0\xba\xe9\x4c\x9f\xc3\xe5\x64\xf2\x44\x30\x9c\x61\x02\xa2\x26\x37\x70\x00\x70\xd6\x56\x2c\x1f\x71\xd6\xb5\xed\x88\xc7\x3a\xb5\xb5\x94\x7a\x0d\x85\x11\x21\x64\x49\x9f\xa6\xaf\x24\x57\xd3\xa1\x79\xd5\x34\xef\x94\xf0\xca\x63\xde\x34\x69\x89\xa4\x9c\xdc\xed\xa0\x69\xd2\xda\x9b\xee\xa6\xf2\x8e\x5c\xe1\xcc\x6e\xc7\x59\xc4\xf5\x8f\xa8\x59\x7e\x8b\x7f\xd6\x18\x88\x33\x35\x1b\x76\x0f\xd2\x92\x58\x18\x1c\x7b\x0c\x95\xb3\x41\xaf\xb1\xcf\x1f\xdf\xa0\x18\x79\x06\x83\x0e\xdc\x95\x18\xe5\x1e\x1a\xc1\xfe\x69\x11\x97\x32\xbf\x72\x31\x27\xc9\xe3\xfd\xa6\x49\x95\x0b\x14\x75\x1e\x06\x39\x3b\xa4\xf8\x04\xdf\xaf\xed\xa9\x4f\x30\x0e\x25\x79\x21\xe7\x4f\xb5\x31\x70\x7f\x7b\x7d\x82\x75\x59\x1b\x73\xef\xcd\x8b\x69\x6f\x04\xa9\x13\x94\xf5\x67\xd0\xfd\xe0\xdc\x83\xc6\x13\x84\x45\x1b\xfc\x0c\xce\xb6\xdb\xae\xd1\xae\x4e\x6a\x2d\x0e\x31\x2f\xcf\x60\x74\x7c\x6b\x7e\xbe\x39\xc1\xee\xb1\x74\x84\xdf\x49\xe9\xcf\x51\x73\xd6\xf6\x5c\xff\x52\x30\xa9\xd7\x07\x9d\x7d\x85\x42\xa2\x0f\x5d\x67\xc3\x2b\x35\xf6\x30\x6a\x4e\x1c\xb5\x85\xe4\xbf\xe0\x96\x33\x52\x1f\x07\x7e\x17\xa6\xc6\xe3\xd0\x51\xed\xd8\x51\x8e\xa6\xf1\xc2\xae\x10\xde\x3e\xe0\xf6\x1d\xbc\x5d\x47\x0a\x98\x67\x90\xaa\xf6\xbc\xbb\xdd\xb9\xb2\x37\x4d\x7c\xec\xb8\xa4\xfb\x58\x4b\x76\xde\xcb\xa6\x41\x2b\xf7\x49\xce\x1a\x70\x7f\x7b\x0d\x37\xc2\x8b\x12\xe9\x0b\xf4\x81\xd3\xf0\x3f\xfd\xbf\xce\xd4\xde\xb4\x75\x08\x07\xe6\x7c\x52\xe2\x69\x83\xce\x5b\x74\x2c\xf6\xd8\xa6\x36\x7e\x28\xf7\xac\x6f\xdf\x3b\xb9\x7d\x72\xab\x69\xf4\x12\xd2\xf8\x70\x37\x4a\xf6\x77\x68\x02\xf6\x19\x38\x96\xf9\xd7\xbf\x39\xf0\xdd\x30\x69\x27\xfd\x37\x9c\x61\x39\x50\x74\x62\x3a\x6c\x37\xdc\xfb\x7c\x00\xd7\x28\xbc\x85\xd2\x79\x04\xb1\x70\x35\x0d\x0a\x45\xff\x85\x92\x00\x09\xbf\x8a\xdf\x6f\x7f\x2c\x8c\xb0\x0f\x49\x8e\x71\xf2\xc1\xbf\x7f\xff\x03\x77\x0a\xa1\x74\x81\x00\x45\xd8\xc2\xa3\xd8\x02\x39\xf0\xb5\x85\xab\xbb\xbb\x9b\xbd\x1a\xc2\x40\x29\x67\x62\x38\xed\x20\x60\x7f\x72\xce\xba\xda\xc4\x19\x1c\xbf\x29\xff\x0b\x00\x00\xff\xff\x30\xd1\xc4\xb7\x63\x0a\x00\x00")

func viewEchoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewEchoHtml,
		"view/echo.html",
	)
}

func viewEchoHtml() (*asset, error) {
	bytes, err := viewEchoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/echo.html", size: 2659, mode: os.FileMode(420), modTime: time.Unix(1526212617, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewLogfileTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4d\x6a\xc3\x30\x10\x85\xf7\x3a\xc5\x2c\xb2\x68\x69\xeb\x03\x18\xba\xe8\xcf\xa2\x85\xb4\x84\x40\x0e\xa0\x46\x93\xd8\x44\x96\x52\x79\x5c\x08\xd2\xdc\xbd\x8c\x64\x8a\x29\x94\x7a\x61\xad\x86\xf7\x1e\xbc\x4f\xd2\x24\x95\xfe\x3f\xd0\x93\x0e\x04\x33\x92\x29\x25\x95\xe0\x59\x13\x42\x8c\x15\xb5\x1d\x32\x2b\x91\x62\xac\x3a\xa4\xc6\x1b\x66\x99\x87\x60\xcb\x70\x0e\x9e\xfc\xde\xdb\x31\xb6\xc5\xcf\x01\x7b\x82\xbb\x45\x8e\x4a\xf0\xe2\x7b\xaa\xa5\xa8\xf1\x3d\x49\x09\xbc\x65\x8c\x7a\x4a\xa4\x12\xec\xb6\xeb\x2c\x1d\x06\x6b\x77\x21\xe3\xc0\x46\x53\x53\xff\xd0\xaa\x04\x4f\xde\x9f\x5a\xcc\xd2\x3e\x8f\xa3\xea\x08\x1d\xad\xd1\x1d\xc7\xfc\x7e\xaa\x94\x8c\x6d\xd1\x11\xbc\x6e\xb2\x1f\xb0\xf3\x84\x0f\xc6\x84\x6c\x0a\x26\x6a\x83\xa1\x5f\xea\xde\x31\x06\xed\x8e\x08\xab\x13\x5e\x6e\x61\xf5\xa5\xed\x80\x50\xdf\x43\xd5\x94\x1e\x66\xf9\x11\x71\x99\x85\xa8\x24\x98\x55\x8c\xe8\x0c\x73\x2a\x2f\x02\x67\x1d\x74\x87\xb4\x00\xd9\xdf\x48\x43\xb0\x1b\xa9\x99\x05\xf5\xe8\xcd\x65\xa1\x47\x2a\x50\xed\x01\xaa\x0f\x6f\x2e\xa5\x7c\x1c\xa5\xd1\xf6\x28\xda\xd5\xbb\x87\x30\x2e\xa5\x98\xd7\x53\x9a\x5f\x7b\x7e\x33\xa3\x12\xd0\x99\xb9\x17\xf8\x0e\x00\x00\xff\xff\x29\xcc\x76\x42\x9c\x03\x00\x00")

func viewLogfileTplBytes() ([]byte, error) {
	return bindataRead(
		_viewLogfileTpl,
		"view/logfile.tpl",
	)
}

func viewLogfileTpl() (*asset, error) {
	bytes, err := viewLogfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/logfile.tpl", size: 924, mode: os.FileMode(420), modTime: time.Unix(1526212563, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewTerminalTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4d\x6a\xc3\x30\x10\x85\xf7\x3a\xc5\x2c\xb2\x68\x69\xeb\x03\x18\xba\xe8\xcf\xa2\x85\xb4\x84\x40\x0e\xa0\x46\x93\xd8\x44\x96\x52\x79\x5c\x08\xd2\xdc\xbd\x8c\x64\x8a\x29\x94\x7a\x61\xad\x86\xf7\x1e\xbc\x4f\xd2\x24\x95\xfe\x3f\xd0\x93\x0e\x04\x33\x92\x29\x25\x95\xe0\x59\x13\x42\x8c\x15\xb5\x1d\x32\x2b\x91\x62\xac\x3a\xa4\xc6\x1b\x66\x99\x87\x60\xcb\x70\x0e\x9e\xfc\xde\xdb\x31\xb6\xc5\xcf\x01\x7b\x82\xbb\x45\x8e\x4a\xf0\xe2\x7b\xaa\xa5\xa8\xf1\x3d\x49\x09\xbc\x65\x8c\x7a\x4a\xa4\x12\xec\xb6\xeb\x2c\x1d\x06\x6b\x77\x21\xe3\xc0\x46\x53\x53\xff\xd0\xaa\x04\x4f\xde\x9f\x5a\xcc\xd2\x3e\x8f\xa3\xea\x08\x1d\xad\xd1\x1d\xc7\xfc\x7e\xaa\x94\x8c\x6d\xd1\x11\xbc\x6e\xb2\x1f\xb0\xf3\x84\x0f\xc6\x84\x6c\x0a\x26\x6a\x83\xa1\x5f\xea\xde\x31\x06\xed\x8e\x08\xab\x13\x5e\x6e\x61\xf5\xa5\xed\x80\x50\xdf\x43\xd5\x94\x1e\x66\xf9\x11\x71\x99\x85\xa8\x24\x98\x55\x8c\xe8\x0c\x73\x2a\x2f\x02\x67\x1d\x74\x87\xb4\x00\xd9\xdf\x48\x43\xb0\x1b\xa9\x99\x05\xf5\xe8\xcd\x65\xa1\x47\x2a\x50\xed\x01\xaa\x0f\x6f\x2e\xa5\x7c\x1c\xa5\xd1\xf6\x28\xda\xd5\xbb\x87\x30\x2e\xa5\x98\xd7\x53\x9a\x5f\x7b\x7e\x33\xa3\x12\xd0\x99\xb9\x17\xf8\x0e\x00\x00\xff\xff\x29\xcc\x76\x42\x9c\x03\x00\x00")

func viewTerminalTplBytes() ([]byte, error) {
	return bindataRead(
		_viewTerminalTpl,
		"view/terminal.tpl",
	)
}

func viewTerminalTpl() (*asset, error) {
	bytes, err := viewTerminalTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/terminal.tpl", size: 924, mode: os.FileMode(420), modTime: time.Unix(1526212604, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewWsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x5f\x6f\xe3\xb8\x11\x7f\xf7\xa7\x98\xea\xee\xc1\x69\x36\x62\x92\xeb\x02\x0b\x5b\x72\xd1\xbb\x5e\xb1\xd7\x66\x77\x83\xbd\x6c\xef\x61\x77\x5b\xd0\xe2\xd8\x62\x43\x91\x2a\x49\xf9\x4f\x0d\x03\xfd\x0e\xfd\x86\xfd\x24\x05\x49\x49\xa6\xed\x38\xee\x16\xb8\x87\x1e\xf2\x60\x6a\xe6\x37\xff\x47\x33\x54\xb2\x5f\x31\x55\xd8\x75\x8d\x50\xda\x4a\x4c\x06\x99\xfb\x01\x41\xe5\x3c\x4f\x50\x26\x93\xc1\x20\x2b\x91\xb2\xc9\x00\x20\xab\xd0\x52\x28\x4a\xaa\x0d\xda\x3c\x69\xec\xec\xea\x55\xe2\x19\x96\x5b\x81\x93\xef\x8b\x52\xa9\x8c\x84\x87\x1e\x2f\x69\x85\x79\xc2\xd0\x14\x9a\xd7\x96\x2b\x99\x40\xa1\xa4\x45\x69\xf3\x24\x39\x84\x2d\x38\x2e\x6b\xa5\x6d\x84\x59\x72\x66\xcb\x9c\xe1\x82\x17\x78\xe5\x1f\x5e\x00\x97\xdc\x72\x2a\xae\x4c\x41\x05\xe6\x37\xce\x4b\x80\x4c\x70\xf9\x08\x1a\x45\x9e\x18\xbb\x16\x68\x4a\x44\x9b\x40\xa9\x71\x96\x27\xa5\xb5\xb5\x19\x11\x52\xd1\x55\xc1\x64\x3a\x55\xca\x1a\xab\x69\xed\x1e\x0a\x55\x91\x9e\x40\xbe\x49\xbf\x49\x7f\x43\x0a\x63\x76\xb4\xb4\xe2\x32\x2d\x8c\x09\xee\xfe\x6c\x66\xae\x6c\x89\x15\xee\x1b\xf3\x26\xc0\xd5\x27\x4f\x2c\xae\x2c\xe9\x38\x00\x53\xc5\xd6\xb0\xf1\x47\x80\x8a\xea\x39\x97\x23\x78\x59\xaf\xe0\xf6\xba\x5e\x8d\x3d\x7d\x3b\xf0\x3f\x33\xa5\x2c\xea\x23\xec\xad\x03\x5f\xbb\xbf\x71\xcb\xa9\x29\x63\x5c\xce\x47\xf0\xca\x71\x3a\xaa\xb3\x7b\x45\x05\x9f\xcb\x11\x14\x28\x2d\xea\x8e\x33\x55\x9a\xa1\xbe\xb2\xaa\x1e\xc1\x4d\xbd\x02\xa3\x04\x67\xf0\x15\x63\xac\x43\xcc\x94\xb4\x57\x86\xff\x03\x47\x70\x73\xdb\xb9\x05\x50\x28\xa1\xf4\x08\xbe\x7a\xf9\xf2\xe5\x9e\xa7\x69\x5b\xf6\xc8\xd7\x55\x28\xfa\x08\x5e\x5d\x5f\xef\x14\x74\x31\x5c\x03\x6d\xac\xea\x74\x00\x64\xc4\x67\x6c\x32\xc8\x48\x68\xdb\x41\xe6\xf2\xe4\x73\xc9\xf8\x02\x0a\x41\x8d\xc9\x93\xd6\x4c\x68\x1c\x80\xac\xbc\x09\xdd\x0b\x3f\xe1\xd4\xa8\xe2\x11\x6d\x46\xca\x9b\x8e\x5b\x6b\x04\xce\xf2\x84\xcb\x99\x4a\x26\x9b\x4d\x5a\xa1\x2d\x15\xdb\x6e\x61\xb3\x49\x1b\x2d\xc2\xa1\xd6\xca\xaa\x42\x89\xed\x36\x23\xb5\xc6\x5e\xf5\xed\xe4\x81\xea\x39\x5a\x68\xb4\xc8\x48\x79\x3b\xe9\x95\x3a\x55\xb4\xe6\x1f\xde\xdf\x1d\xcb\xbc\xc7\xbf\x37\x68\x6c\x10\x08\xd4\xc8\x7f\x4b\xa7\x02\xaf\x34\x9a\x5a\x49\xc3\x17\xd8\xb6\x84\x7b\x15\x1d\x67\x0f\x06\x01\x1c\x6a\x85\xac\x87\x3a\xb0\xde\x3d\xb8\x47\x36\x79\xad\x9c\x4d\xcb\x0e\xe8\x3e\xfc\x52\x19\xeb\xc3\x77\x07\xe7\x71\x0c\xcb\x48\xac\xec\x09\xcd\x6f\x7c\xce\x4e\xe9\x0e\x19\xdd\x4b\xee\x17\xea\xff\x43\x23\x04\x7c\x78\x7f\x77\xca\xc2\xac\x11\xe2\x83\x16\xde\x44\x7b\xfe\x62\x1b\xf7\xd4\x96\xa7\xf4\x37\xad\xee\xe6\x7f\xd0\xfb\x5d\x68\xc7\x3b\x94\xf3\xd3\x06\x8a\x18\xe4\x4d\xed\x51\xbe\xdc\xa8\xe0\xee\x4d\xfb\xe1\xfe\x94\x41\x8d\x95\xb2\xf8\x3b\xc6\xb4\xb7\xb6\x7b\x7c\xce\x54\x46\x7c\xbb\xb5\x4d\x4e\x18\x5f\x44\x4d\xfd\x1a\x29\x43\x6d\xa2\xb7\xe0\x67\xe8\xe9\x6e\x5d\x9d\x08\xdd\x43\x26\x7f\xc2\x75\x46\x6c\x79\xcc\xf8\x33\x15\x0d\x1e\xb2\x0e\x72\x49\x0e\x6c\x64\xd6\xcf\x62\xff\x8e\x84\x08\x93\x58\x7a\xb3\xd1\x54\xce\x11\xbe\x7e\xc4\xf5\x0b\xf8\x7a\xe1\x2c\xc0\x28\x87\x34\x80\xb7\xdb\xe7\x9d\x65\x93\xcd\xc6\x89\x1e\x66\x3d\xe2\x7a\x95\xc7\xfc\x7d\xb7\x9d\x23\x28\x59\x64\x2e\x23\xb6\x1b\x8e\x67\x4b\xf7\xe1\xfd\x1d\xdc\x53\x4d\x2b\xb4\xbf\xec\x0a\xd6\x2e\xc8\xff\xb2\x80\x8d\x16\x3e\x27\xe6\xff\xa3\x86\xdf\x2a\xb6\xde\xdf\x40\x3e\x62\x27\xdf\xc7\xbb\xd9\xf0\x19\xa4\x8e\x14\xb6\x5a\x7f\x42\x61\xb0\xb7\x9b\x61\x35\x19\xbe\x55\xa0\xc3\x96\xf2\x77\x91\x8b\x8c\x60\xb5\x53\xb3\x73\xb3\x5d\x6d\xfe\x18\xae\x22\xad\x4f\x00\x77\x48\xb5\x84\x4a\x69\x04\x3a\x55\x8d\xed\xd4\xd3\xf6\x3e\x95\x80\xf5\xbb\x33\x4f\xfe\x3a\x15\x54\x3e\x26\x13\xf4\x6b\xfa\xdf\xff\xfc\x17\x3c\x94\x08\x95\x32\x16\x90\x9a\x35\x2c\xe9\x1a\xac\x02\xdd\x48\x78\xfd\xf0\x70\xdf\x7b\x66\xd1\xd8\x34\x23\xb4\xcb\xc8\x9e\x03\x59\xb8\x94\x76\x4e\x2f\xa8\x06\xb7\xe3\x7f\xaf\x2a\xc8\x81\xa9\xa2\xa9\x50\xda\x74\x8e\xf6\x7b\x81\xee\xf8\xed\xfa\x07\x36\x0c\xd7\x80\x8b\x17\x7d\x09\xdc\x3a\xfc\x42\x91\xb0\xe3\xce\x08\xb5\x3b\x31\x12\x6b\xf7\xd6\x19\xb9\x6e\xd3\x45\x82\xcd\x59\xa1\x66\x5f\x60\x6f\xb5\x9c\x11\xdd\x5f\x4c\x91\x92\xdd\xc6\x38\xa3\x21\xda\x34\x91\xb8\x6b\xaa\x33\x82\xbe\x75\x23\x11\x42\xc2\x5c\x31\xbb\xe2\x84\x81\x7c\x46\x4f\x37\xb6\x23\x55\x61\x0e\x9c\x91\x6b\x87\xc5\xc5\xb8\x6b\xe8\x59\x23\x0b\xf7\x91\x03\x1a\x25\x43\x3d\xac\xcc\xfc\xa2\xbf\xcc\x42\xd7\x5d\x29\x97\x12\xf5\x03\xae\x2c\xe4\x50\x99\x79\xea\xe8\xe3\xc3\x86\x3a\x42\x39\xfa\xf8\xb8\x87\x8e\x70\x81\x33\x7e\xa2\x6d\x8e\xa0\x2d\x6b\x7c\xd0\x29\x47\xb8\x26\xc6\x1c\x36\xc7\x11\x7a\x0f\x30\x7e\xba\x1f\x8e\x84\x76\xdc\xf1\x61\x0b\x1c\x61\x1d\xbd\xcf\x79\x5c\xe4\x80\x7c\xfd\xf0\xe6\x0e\x72\xb0\x58\xd5\x82\x5a\x1c\xfa\xe4\x05\xcc\xc5\x71\x81\x4f\xcb\xf4\xf3\xbd\x93\xda\x1e\xd5\xb9\xc7\x33\x6a\x69\x5c\x6a\x37\x4c\x1e\x71\x6d\x20\x87\x77\xd3\xbf\x61\xe1\x9b\xe7\xdd\x52\xde\x6b\x55\xa3\xb6\xeb\xb7\xb4\x42\x13\xa4\x5e\x44\xf3\xfd\x11\xd7\x90\x43\x92\xc4\x34\x8d\xa6\x11\xd6\x93\xa3\xa0\x67\x4a\xc3\xd0\x8f\x2c\xc8\xe1\x7a\x0c\x1c\x32\x6f\x30\x15\x21\xed\xc0\x2f\x2f\x63\x87\x3a\xdd\x0e\xf3\x91\x7f\x1e\x1f\x1b\xb8\xcc\x21\x71\xcb\x2b\x81\xcb\xbd\x0d\x95\xb8\x15\x95\xc0\xa5\x57\x70\x09\x89\xdf\x51\x27\x41\x2e\xa4\x8f\x8f\xb8\xfe\xfc\x0c\xd4\x2d\xb5\xa4\xa7\x6d\xa3\x0e\xb1\x8d\x96\xad\x3f\xe3\x41\xcc\xce\x48\x3c\xac\x9f\x98\xdc\x4d\xc3\x19\xe4\x40\xfc\x37\xf8\x6f\x47\x9f\xc8\x27\xf2\xf1\x2f\x9f\xc8\xe7\x5f\x7f\x22\xc3\x70\xb8\x20\x29\xae\xb0\x18\x0a\x55\x50\x57\xbc\xd4\x2d\x99\x8b\x8f\x37\x9f\x5b\x2d\x7c\x06\xc3\xa0\x25\xcf\xa1\x91\x0c\x67\x5c\x22\x8b\x73\x58\x28\x69\x94\xc0\x14\xb5\x56\x7a\x98\xfc\xa4\x95\x9c\x83\x1f\x9d\x3d\xc4\x96\x5a\x2d\x21\x62\x0d\xf6\x83\x74\xae\x2e\x5d\x57\x48\x5c\xba\xcf\xcd\x1f\xfd\xe7\xe6\x30\x59\x9a\x11\x21\x2e\x7f\x3b\xef\xdc\x72\xbb\x84\x84\x2c\xe3\x09\xb3\x34\xa9\x92\xaa\x46\x09\xf9\xae\x0b\x87\xb8\xb0\x4f\xf9\x29\xd4\x7c\x98\x7c\xa7\xa4\xc4\x80\xf3\x72\x69\x9a\x3a\x7d\x1d\x76\x69\x52\x83\x92\xf9\xc8\x7b\xf2\x76\xdf\x5e\x85\xc6\xd0\x39\x3e\x67\xb2\x1d\x78\x7f\xfc\xf1\xdd\xdb\xb4\xa6\xda\xa0\x03\xa4\xbe\xbf\x2f\x9e\xf6\xeb\x8d\x99\x83\xc6\x82\xe3\x02\x59\xe4\xd1\x81\xe9\x42\x28\xf3\xac\xe1\x53\xb1\x7a\xc1\x03\xc5\x87\x8d\xd4\xde\x90\x32\x12\xee\x50\x83\x8c\xf8\xff\x82\xfd\x27\x00\x00\xff\xff\x5c\x03\xf1\x21\x15\x13\x00\x00")

func viewWsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewWsHtml,
		"view/ws.html",
	)
}

func viewWsHtml() (*asset, error) {
	bytes, err := viewWsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/ws.html", size: 4885, mode: os.FileMode(420), modTime: time.Unix(1526282867, 0)}
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
	"static/favicon.ico": staticFaviconIco,
	"view/echo.html": viewEchoHtml,
	"view/logfile.tpl": viewLogfileTpl,
	"view/terminal.tpl": viewTerminalTpl,
	"view/ws.html": viewWsHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"favicon.ico": &bintree{staticFaviconIco, map[string]*bintree{}},
	}},
	"view": &bintree{nil, map[string]*bintree{
		"echo.html": &bintree{viewEchoHtml, map[string]*bintree{}},
		"logfile.tpl": &bintree{viewLogfileTpl, map[string]*bintree{}},
		"terminal.tpl": &bintree{viewTerminalTpl, map[string]*bintree{}},
		"ws.html": &bintree{viewWsHtml, map[string]*bintree{}},
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

