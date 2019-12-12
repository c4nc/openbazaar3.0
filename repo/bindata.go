// Code generated by go-bindata.
// sources:
// sample-openbazaar.conf
// DO NOT EDIT!

package repo

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

var _sampleOpenbazaarConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xdf\x73\x13\xb9\x93\x7f\xf7\x5f\xd1\x0f\xfb\xad\xbd\xab\x32\xe3\xd8\x90\x10\xf0\xf9\xaa\x0c\x09\x4b\x8a\x90\xb8\xe2\x40\x76\xf3\x26\xcf\xf4\x78\x74\xd1\x48\x83\xa4\xb1\xe3\xbd\xba\xfd\xdb\xaf\xba\x25\xcd\xd8\x01\xf6\xe1\x0a\x6a\x0f\x1e\xc0\x1a\xa9\xbb\xd5\x3f\x3f\xdd\x9a\xc2\xb3\x1f\xfa\x67\x30\x85\x33\xe1\x05\x38\xf4\x5e\xea\xb5\x1b\xfc\x70\x06\x83\x29\xdc\x56\x08\x85\xb4\x98\x7b\x63\x77\xe0\x0d\x38\x6f\x2c\x42\xc1\x8c\xdb\xbc\x02\xe1\xc0\x57\x08\xa6\x41\xbd\x12\x7f\x0a\x61\xf9\xdb\x4a\x38\x1c\x82\x6c\x4a\x07\x35\x7a\x41\x4b\x43\x10\xba\x18\x4c\xa1\x69\x57\x4a\xe6\xbc\x2b\x4b\x0c\xb0\x14\xad\xf2\x20\x1d\xfc\x35\xca\xf6\x48\x19\x0d\x8b\xeb\xe5\xc5\xef\x70\xbd\x44\x37\x84\x5f\x2e\xaf\xdf\xce\x2f\xe7\x8b\xc5\xd9\xfc\x76\x3e\xba\x6e\x50\xbf\xe9\xf6\xdd\x49\x5d\x98\xad\x1b\x0e\xa6\xf0\xd7\xe8\x52\xae\xac\xb0\xbb\xd1\xbc\x69\x94\xcc\x85\x97\x46\xc3\xb2\x6d\x1a\x63\xfd\x93\x63\x1f\x45\x0e\xd7\x4b\x96\x0d\x7e\xa9\x4c\x8d\xa3\x03\xf6\x83\x29\x2c\x94\xd0\xaf\x32\x80\x73\xbd\x91\xd6\xe8\x1a\xb5\x87\x8d\xb0\x52\xac\x14\x3a\x10\x16\x01\x1f\x1b\xa1\x0b\x2c\xc0\x19\xd2\xc5\x0e\x6a\xb1\x83\x15\x42\xeb\xb0\xc8\x00\xae\xae\x6f\xcf\x5f\x27\xf9\x06\x53\xc0\xef\x12\xf2\xbb\x46\xe6\x42\xa9\x1d\xfc\xeb\xf3\xfc\xe6\x62\xfe\xe6\xf2\xfc\x5f\x43\x58\xb5\x3e\x92\x6d\x9d\x27\xba\x22\xcf\xd1\x39\x2c\x60\x2b\x7d\x35\x98\xc2\x2f\x69\x33\x54\x68\x31\x03\x98\x2b\x67\x86\xf0\x17\xe9\xb3\x93\xcd\x9b\x43\xf5\xed\xe9\x8c\xcc\x40\xe6\x28\xa4\x9d\x1d\xe8\x7f\x30\xf8\xf1\x3e\x35\x85\x2b\xf4\x5b\x63\x1f\x7e\xae\xdf\x7e\x72\x08\x1e\x9d\xd7\xe8\xe9\x7a\xf1\xbf\xb3\x31\x7f\xd3\x72\x83\xd6\x09\x05\x0b\xd5\xae\xd9\xf4\x0b\x25\x76\xf0\x6f\x9f\x16\x7a\xf1\xef\x20\x5a\x6f\x6a\xe1\xa3\x25\x48\x1b\xc1\xc5\x95\x74\x1e\x35\x90\x13\x81\x59\x79\x21\x35\x89\x4e\x5f\xf0\xd1\xa3\xd5\x42\xc1\xc5\x02\x44\x51\x58\x74\x0e\x4a\x6b\x6a\x70\xc1\xe7\xb0\x80\x02\x37\x32\x47\x97\xc1\x6d\x25\x1d\x98\x86\x5d\xb2\x90\x2e\x18\x5f\xb2\x90\xda\xb4\x8d\x6e\x82\x8c\x7f\x98\x96\xdd\xc8\x35\x98\xcb\x72\x07\x46\x23\x18\x0b\x35\x05\x9f\xdb\x0a\x5b\x27\x46\xe8\xc8\xb4\x51\x36\xa3\xa1\x34\x16\xa4\xce\x4d\x2d\xf5\x1a\x74\x50\xf5\x60\x0a\xb9\xd1\x1a\x73\xe2\xca\x32\xa0\xc3\x3d\x02\xe4\xa8\xe4\x58\x52\x83\x80\x8d\x50\xb2\x80\xba\x55\x5e\xd2\x0e\x22\x58\x0b\x96\x8f\xf9\xd2\xda\x6c\x24\x9b\x17\xa3\xa3\x8c\xff\x8e\x7c\xde\x8c\x5e\x1c\x1d\x8d\x9f\xee\x38\x19\xbd\x7e\xfd\xdd\x8f\x87\xc7\x5f\x1d\x1d\x1d\x8f\x38\x38\xbe\x4d\x21\x7d\x8f\xf9\x62\x2d\x3c\x6e\xc5\xae\xd3\x35\x0b\xdb\x28\x7c\x44\x07\x2b\xe3\x2b\x36\xca\xc5\xe2\xdd\xb2\xdb\x39\x5f\x5c\xb0\x9d\x0f\x53\xd5\x60\xca\x1f\xcc\x06\x2d\x7f\x71\xa2\xee\xd4\xc2\x5a\xda\xe3\xe0\xaa\xa8\xa1\xef\xeb\x27\x32\xeb\xaf\x38\x9e\xbc\xe4\x4b\x8e\x93\x1a\x26\x74\x83\x37\xc6\x78\xe7\x45\xb3\x67\x00\x0a\x7e\x36\x82\x37\xf0\x5f\x46\x6a\x96\x26\x1a\x2f\x83\x6b\x0d\xce\x0b\xeb\xc3\xaa\x29\x10\xb6\x52\x29\xa8\xc5\x03\x0e\xa6\x60\x5a\xbf\x36\x64\xec\x3d\x13\x13\x1d\xda\xbc\x62\x56\x56\x34\xd0\x20\x5a\xc7\x2a\x68\x29\x32\x2a\xac\x69\x4f\x21\x5d\xce\xb7\x37\xbe\x42\x52\x47\xd8\xf6\x44\x80\xc1\xb4\x27\xd4\x5f\xee\x31\xe3\xbf\x9d\x85\x47\xcd\xa4\x19\x8d\x27\x67\xcf\x3f\x18\x73\xf7\xe6\x7c\x99\x4f\xfc\x52\x6f\x3e\xdf\x60\xfd\xc1\xb9\xb3\x8f\xf2\xc3\xe5\x3d\x7e\x28\x3f\xdd\x54\xdb\xdf\xc5\xf6\xfe\x4e\x48\xf3\xc5\x2d\x9e\x6f\xc6\x5b\xd2\xc9\x92\xab\x0a\x89\x57\x1a\xbb\x15\xb6\x00\x87\x76\xc3\x22\xef\xa9\xc6\x62\x8e\x72\x83\x50\xa3\x73\x62\x8d\x0e\xb6\x15\x39\x7d\x59\x2a\xa9\x31\x83\x05\xa2\xbd\x38\x63\x2f\xe2\xa8\x91\x58\x70\x46\x0c\xea\x5a\x21\x65\x9c\x74\xb7\xc6\x9a\x52\xaa\xc0\x92\x2f\xcf\x8a\x75\x61\x6b\xa8\x71\x89\xcb\x60\xca\x89\x36\x28\x4d\x96\x21\x17\xe7\x42\x6b\xe3\x93\xce\x83\xbe\xa5\x63\x22\x29\xbe\xf6\x6f\xe0\x49\xd0\x2f\x2d\xda\x1d\x05\xfc\x60\xda\x39\x63\x6f\xce\xc2\x6c\xb5\x32\xa2\xe8\x6f\xc7\x29\x84\xb8\x66\x83\xa9\xd3\x65\xa0\x37\xfb\xbf\xaa\xf8\x07\xe7\xd8\x29\xdc\x1a\xfb\x73\x73\xf8\xec\x87\xfe\x19\x4c\xe1\x7b\x7f\xee\xe6\x37\x57\x17\x57\xbf\xc1\xb3\x67\x70\x36\xbf\xfa\xed\xfc\x06\xee\xaf\xaf\xce\xe9\x67\xfc\x32\x98\xc2\x1e\x6c\x68\x39\xe9\xa6\x7c\x41\x21\x03\x17\x67\x9c\x78\x05\x39\x0f\x3a\x17\xd2\xec\x45\x09\x3b\xd3\x1e\xfa\x08\xee\x11\xa2\x94\x1f\x6b\x21\x6e\x38\x7b\xe7\x98\xfc\x33\x57\x28\xec\x90\xce\x5b\xb0\x78\x58\x5a\x22\xbc\x68\xd0\xd6\x42\xa3\xf6\x8a\x10\x47\xd3\x84\x18\xa1\x13\x31\x90\x49\x2a\xf2\xb3\x8d\x74\x72\xa5\x90\xbe\x86\xf8\x36\x4f\x12\x4c\x14\x94\x1c\x55\x6a\x8f\xba\xa0\x74\xe2\x0d\xa7\x0a\xb2\xb2\x37\x50\x0b\x47\x65\x84\xe5\xe9\x45\x61\x01\x03\x2e\xb9\x3a\xff\x7c\x7e\x13\xf3\xd4\x9e\xae\x28\x72\x4c\xeb\xa1\x75\x44\xf3\xd6\xd8\x0c\xae\x8c\x4f\xf7\x25\x31\x06\x53\x28\xa5\x75\x3e\x9c\xcd\x98\x61\x42\x3a\xb9\xd1\xa5\x5c\xb7\x16\x8b\x94\xba\x0a\x3a\x85\x1b\xb4\x3b\x20\x8a\x0a\xc3\xb1\xb6\x49\xb7\xa0\xd8\xca\x73\x59\xa0\xf6\x5c\xbf\xf9\x33\x16\x7f\x2b\x53\xb8\xc6\xc7\x4f\xcb\x5b\x28\x50\xa1\xc7\x70\x4f\x06\xb9\x1d\xf8\x8d\x41\x1b\x6e\x48\x49\x33\x83\x33\xda\xcc\xba\xaa\xf0\xc9\xee\x10\xd3\xa5\xb1\xf9\xbe\xc5\x93\x52\x69\x63\x59\xa2\x45\xed\x7b\x5b\x65\x5c\xf4\xf9\x9c\x32\xb4\x49\xef\xb8\xae\x53\x7c\x0d\xc1\xd8\x02\x2d\xff\x0b\xb9\x91\xda\xb1\xc8\x95\xd8\x90\x17\x6e\xb0\xa0\xc4\x44\x2b\x85\x01\x67\xb2\x1f\x1f\x3c\x31\xde\xeb\x2e\x5d\x05\x3d\x08\x0d\x58\xaf\xb0\x20\x84\x49\xdf\x0b\x81\xb5\xd1\x94\x5d\x1f\x77\xa1\x14\x77\x58\x84\x33\xed\x37\x6a\x15\x95\xb0\x54\x80\x89\x44\xe7\x95\x8c\x94\x98\x21\x07\x1c\x7d\xc3\xc7\x5c\xb5\x4e\x6e\x50\xed\x98\x1e\xa5\xe0\x2e\x5a\xd8\x77\x6d\x02\x7c\xc6\x06\x20\x75\xd6\x0a\x16\x36\x7f\xd8\x13\x9e\x10\x74\xe3\x7b\xd9\x0e\x4a\x67\x65\x6c\xbb\xae\x82\xf4\xc4\x74\x7e\x75\xd6\x33\x19\x4c\x7b\x36\x94\xe7\x2d\x96\xdc\x0f\xb5\x42\xed\x31\x91\x8e\x50\x3f\x34\x56\x6e\x84\xc7\x0c\xae\xbf\x55\xa3\x63\x55\x1a\x4c\xa1\x16\x05\xf6\x4a\x38\xbc\x0c\xb4\x5a\x51\xd0\x7b\xa1\x1e\x62\x58\x8a\x50\x35\x6c\xab\x35\xad\xec\x2b\x65\x85\x95\xe4\x2e\x8b\x22\x8d\x60\x7d\x92\x2b\x28\xe3\x07\xe7\xe8\x29\x5c\x91\x20\xcb\x58\x04\xe0\x19\x63\xa6\xd2\x28\x65\xb6\x24\x59\x80\xb9\x3f\xaf\x31\xd5\x6d\xbd\x22\xf0\x52\x82\x45\xd7\x18\x1d\xc1\xf0\x56\x48\xcf\xe9\x98\xe1\x41\x2d\x58\x6f\x17\x8b\xab\x25\x57\x60\xd9\xa1\x70\xe9\x40\x80\xb7\xa2\x40\x53\x96\x04\x72\xd0\x6f\x11\x43\x6e\x14\x79\xde\x5a\x91\xef\x88\x38\xfd\xe6\xda\xdd\x55\x6d\xd7\x20\x16\xa4\x5f\xd9\x68\xf7\xa5\x35\xb6\xad\x67\x8c\xed\x08\x4e\x7e\xba\xb9\x4c\x71\x5e\x06\x9f\xad\x84\x5e\x23\x58\xe1\x89\xf5\x47\xca\x6d\x94\xd8\x8c\xad\x53\x4d\x78\x23\x3d\x05\xf5\x7c\x83\x56\xac\x71\x0f\x52\xa6\xc3\x74\xb6\xb1\x66\x23\x0b\xb4\xb3\xca\xfb\xc6\xbd\x1e\x8d\xbc\xcc\x1f\xd0\xee\xf5\x6b\x99\xb1\xeb\x91\x68\x24\xfb\x7d\xe8\x2d\xb8\x24\xed\x25\x20\x8b\x4a\x50\x3a\x2c\x5b\xcd\x6e\x28\x94\xf4\x3b\x62\x43\xf1\xd0\xc1\x66\xf6\x2f\xba\x6c\xf8\x15\x22\x52\xea\x75\xb8\x72\xe9\x8c\x56\xbb\xe0\x51\xf3\xa6\x41\x5d\x80\x80\xdc\xd4\xdc\xd1\xc6\x1b\xb5\x0e\x2d\x88\x35\xad\x24\xc4\xd5\xf7\xfd\x7d\x9a\xcc\x06\xd3\x56\xc4\xa3\xb3\xf8\xef\x4f\x71\x54\x32\xcc\x3f\xe2\xa7\xa9\x91\xdb\x4a\x57\x91\x72\x50\xb3\x59\x96\xcb\xcb\x54\x86\x49\xb4\x3e\x2f\xf4\xbe\x59\xc9\x75\x45\xb5\xdd\x62\x50\x4c\x81\x14\xd8\xb2\xaf\xd5\x29\x01\xb0\x47\x32\x38\x24\x92\x02\x2c\xd6\xc6\x23\xd4\x22\xaf\xa4\x46\x10\x0e\x4a\x21\x55\x6b\x31\xb9\x25\x31\xa7\xec\x43\x25\x8d\x74\x40\xa5\x86\x1a\x4c\x6f\xf6\xc1\x0a\xd9\x3f\x37\xda\x5b\xa3\x42\x3d\x24\xaf\x18\x52\xd2\x54\x2d\x23\x84\xc2\x0a\xd9\x09\xb0\x15\x4a\x85\xd4\xeb\x9c\x0a\xbe\x71\xdb\x73\xdb\xa5\xca\xa6\x31\xc0\x14\xa1\x9c\xe9\x9a\x5b\x76\x0f\xe1\x2b\x8e\xde\xae\x81\xcb\x91\x0b\x4c\x01\x0f\xb8\x03\x02\xeb\x64\x20\x8a\x28\x16\x86\xbe\xca\x52\xe6\x94\x5f\x03\x53\x5a\xa1\x6d\xb3\x11\xd1\x1a\x79\x33\x72\x4e\x65\xb4\x1a\xbe\x3f\xe0\xee\xeb\xcf\x0f\xb8\x4b\xd9\xa4\xf7\x87\x88\xd8\x61\x25\x9c\xcc\x41\xb4\xbe\x82\xdc\x22\x41\x0a\x29\x94\x63\x19\x92\xe1\xa2\x39\x92\x75\x5b\xc7\xe0\xbe\x25\xbc\xef\xe3\xe4\x89\x11\x0f\x11\x14\xbe\x87\x4b\xa4\x18\xbe\x29\x69\x87\x2a\xcc\xe1\x19\xc6\x6a\xd6\x78\xcc\x49\xf8\xce\xa4\xc1\xca\x19\x5c\xf8\x5f\x5d\x50\x21\x39\xc9\xbe\x8f\xf4\x6c\x18\x67\x1c\x12\x25\xd4\x45\xe5\x56\x83\x32\xb9\x50\x95\x71\x3e\x30\xa2\x0f\x3e\xf6\x41\x8d\x35\x6b\x2b\xea\xd8\x7e\x84\x59\x53\x32\xf2\x7c\x71\xc1\x33\x3b\xf1\x40\x9d\x4b\xba\x54\xd2\x45\x23\x9c\xdb\x1a\x5b\xc0\x0a\xc9\xa9\x12\x88\xa3\xcf\x15\x3e\x02\xea\xdc\x10\x4e\x58\xbe\x9f\x4f\x8e\x4f\xa0\x12\xae\x02\x53\xc6\x11\x8a\xc8\x3d\x15\xea\x44\xa2\x8f\x82\x22\x3a\x66\xd4\x46\xf4\x95\xc8\x68\x5b\x51\x0f\x27\x3d\x38\xe9\x1d\xf7\x7a\x5c\x9f\x83\xfb\x30\x76\x64\xc7\xc9\xe0\x8e\x2a\x01\x2b\x9f\x44\x17\x9a\xe5\xb5\xf8\xa5\x45\xe7\x7b\xe7\x24\xba\xe9\x78\xab\x9f\x91\x84\x1c\x73\x1d\xbf\x94\xff\x59\xf6\xd4\x55\xe6\xa6\x6e\x84\x0d\x6e\xdd\x7d\x0c\xa0\x8c\xe7\x71\x83\xa9\x68\x24\xe5\x43\x2d\x6a\x9c\x09\x25\x73\xe4\xa5\x44\x75\x76\x8c\xa7\xa7\x2f\x4e\x5f\x9d\x16\x62\x72\x7a\xf4\xe2\xe5\xf8\x78\x5c\x1c\xe1\xf1\x49\x79\x5a\xe4\x27\x93\x57\x93\x97\x2f\x9f\x9f\x1c\x3d\x2f\x8e\x8a\x13\x21\x56\xab\xa2\x38\x99\x88\xf1\x18\xcb\x97\x93\x71\x31\x3e\x7e\x31\x29\x4e\x39\x0f\x53\x73\x0f\x42\xf1\x20\xca\x53\x93\x4c\xa1\xd4\xfb\x2f\x37\x22\x42\xb3\x57\xe4\xc6\x3c\x48\xf6\x6e\xc2\xd5\x4f\x7c\xf5\x96\x11\x79\x63\x65\x2d\xec\x2e\x6c\x17\xb1\x92\xf9\x68\x12\xfa\x7f\xe7\x25\xec\x01\xf1\x57\x37\x34\xeb\xc7\x15\xc1\x63\x19\x90\x1d\x98\x90\x3c\x09\xee\x90\xf0\x2a\x81\xb8\xde\x7f\x83\x23\x10\x8d\x90\xad\x03\xd7\x8d\x50\x6d\xec\x8d\xa4\x8b\xa6\xa5\x8a\xdd\x7a\x2a\xab\xec\xb6\x22\xb8\xa9\x8c\x05\xc7\x1a\x02\x71\xc1\x11\xea\x9a\x0c\xa7\x28\x19\xc6\x54\x1f\xc6\xd6\xe1\x3a\xc4\xbf\x33\x75\x48\x65\xbb\xa7\xe1\xdf\x79\x80\x74\xc1\x9e\x41\x87\xb3\x3f\x7e\xbf\x7a\xb8\xaf\xdf\xfd\x79\xff\xdb\xbb\xfa\xfe\xfd\x55\x75\xff\xfe\xaa\xee\xd7\xee\xab\x7c\x72\x53\xdf\xd7\xef\x1e\xee\xd7\x09\x43\x93\xcf\x7a\x24\x5c\x9f\xa6\x14\xf9\x5e\x43\x85\x6e\x08\x4d\x98\xf7\xd6\x9d\xf7\x50\x5a\xc2\x42\x36\xb3\xc9\x69\xf6\xe2\x38\x3b\x79\x99\x8d\x5f\x1e\xef\xaf\x3f\x9f\x64\x93\xe7\xaf\xb2\xf1\xd1\xab\x6c\x7c\xcc\xa9\xf7\xed\xf5\xcd\x92\xc7\xbf\x5c\x6d\x0a\x58\xed\xd2\x90\x9d\x1a\xac\x34\x78\xe4\x81\x88\x3f\x48\x7d\xde\x40\x29\x94\x23\xbe\xda\xe4\xc6\xba\x90\xca\x2f\x0e\xd3\x5c\xa8\x1a\xdd\xc4\x83\x7d\x2d\x8e\x69\x04\x81\xaa\x58\xeb\x09\xa8\xa4\xa9\xd8\x30\x0e\x9e\x24\x03\xfe\x30\xff\x24\xab\x48\x1f\xc7\x1f\x09\xaf\x70\xc2\x89\x4c\x38\x4c\x51\x17\x8d\x91\xda\x3b\x52\x5d\x5e\xa5\x1d\xa1\x1b\x91\xe5\x6e\x30\x4d\x83\x94\x5f\x5d\x84\xcd\x01\xf2\xfb\x80\x61\x98\x3c\x21\x96\x28\x76\x89\x9e\x0a\xe3\x3a\x40\x11\x72\xe6\x38\x00\x8a\xbd\x38\x0f\x82\xb2\xc1\x34\x5c\x22\x8a\xff\x93\xf0\xf3\x1d\x57\xcd\x7f\x06\x99\x9c\x07\x24\x72\x68\xff\x50\xc6\x43\x8c\x93\xc1\xa4\xde\x43\x8f\x3c\x8a\xa4\xb6\x34\x1a\xab\x48\xdb\x79\x44\x16\xb2\x21\x55\x2a\x4a\x7d\x69\x78\x15\x7b\x2d\x2c\x20\x6f\xad\x45\x9d\x4b\xe4\x01\x04\x8a\xbc\x4a\xed\x6d\x36\x88\x7e\x1a\xc8\xcd\xde\xbc\x7d\xff\x74\xe5\xf6\xed\x60\x7a\xb8\x74\xf9\xf5\xd2\xfd\xf9\x57\x4b\xe7\xb7\xef\x7f\x8a\xe1\xc2\xc0\x72\xae\x0b\x78\x17\x07\x96\xcb\x80\xc0\xfe\x39\x53\x76\x50\x90\x44\x7b\x26\x74\xf1\xec\x70\x96\x1a\x5b\xdf\xaf\x43\xd7\x94\x25\xda\x38\xf4\x0c\x6f\x74\xfb\x07\x65\x8e\xdd\x3c\xb9\x1f\x49\x3f\x1d\x99\xae\x10\x44\x9a\x31\xb5\xae\xea\x87\x98\x61\xa0\x84\x91\x6a\x1a\xd6\xee\xcd\xa3\x7d\x65\x1c\x7e\x87\x94\x45\x6f\x25\x6e\x82\x93\x1e\x4e\x7d\x7d\x85\x3b\x7e\xd7\xa8\x29\x51\xe7\x0f\x14\xe1\x3c\x05\x8e\x60\x2b\xf5\x77\x8d\xd9\xa2\x0d\xed\x48\xac\x28\x19\xdc\x74\xc0\x59\xba\xa4\x1c\x57\x99\x56\x71\x05\xe8\x1e\xe1\x56\x18\xd0\x07\x83\xea\x95\x79\x0c\x63\x60\x01\xca\x78\x6a\x0b\x03\xe5\x30\x13\x32\xdc\xb7\x09\x17\x31\x06\x9f\xa5\x55\xa9\xd7\x8c\xdf\xd6\xc6\x14\x50\xa0\x50\x74\x30\x3e\x7d\x06\x47\xdd\x1b\xec\x76\x83\xf0\x6f\x18\x2f\x8c\x91\xc5\xde\x18\x2f\x48\xc3\x61\x14\xf2\x57\x82\xa5\x4d\x6b\x1b\x52\x29\xb7\xc5\xf1\xfd\x93\xc5\x60\xbe\x4f\x53\x79\x4f\x8a\x02\x3b\x50\x4a\x2c\x13\x68\x40\x4a\xaa\x44\x5b\xda\x34\xac\x72\xa9\x3a\x39\x5d\xd2\xd2\xff\x9f\xb1\xf4\x19\xae\xda\xf5\x4f\x89\x32\xa6\x0c\xca\xac\xd7\xe4\x3c\x0a\x37\xa8\x08\x0d\x7f\xe6\x07\x21\xfe\x19\xac\xf4\xdf\x05\x6d\xa4\x4e\xa9\x34\x43\x82\x17\x32\xc7\x21\x6c\x85\x25\xa7\x1b\x02\x5a\x6b\xec\x10\x72\x2b\x19\x2e\xfd\xcf\x60\x4a\x34\xf9\xfc\x8c\x8e\xfc\xcd\xb3\xbb\x32\xeb\xae\x13\x52\x66\xfd\xd5\x83\xed\x48\x99\x75\xf7\x4a\xc6\x0f\x95\xe9\xe9\x24\x3e\x10\x92\x8f\xbc\xbf\xbd\x5d\x74\xef\x1f\x11\x03\xbb\x0c\xc2\x99\xb8\xbc\x97\x31\x78\x30\xd2\x67\x7c\x7e\x00\xe9\x9f\x30\x23\x7e\xea\x1e\x5c\x9e\xd0\x91\x3a\xcc\x32\x68\x2b\x79\x12\x8f\xbb\xba\xf7\x6b\xe1\x19\x22\xbc\x1e\x8d\xba\x7e\xe4\xf5\x7f\xc4\xa3\x24\xfd\x7f\x8e\x58\x93\xa3\x86\xd6\xc2\x78\x3c\xf6\xbc\x19\x63\x54\xde\x38\x3b\x39\x3a\xe1\xd0\xb9\xb3\xd2\x23\xbc\x5d\x7c\xea\xb8\xc7\xac\xd5\xbf\x06\x71\x33\x40\x59\xa3\x69\xd3\xe9\x91\xaf\x9b\xbd\x37\xff\x8c\xd6\xff\x37\x00\x00\xff\xff\x6a\xfc\x89\x91\xaa\x21\x00\x00")

func sampleOpenbazaarConfBytes() ([]byte, error) {
	return bindataRead(
		_sampleOpenbazaarConf,
		"sample-openbazaar.conf",
	)
}

func sampleOpenbazaarConf() (*asset, error) {
	bytes, err := sampleOpenbazaarConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sample-openbazaar.conf", size: 8618, mode: os.FileMode(420), modTime: time.Unix(1576131894, 0)}
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
	"sample-openbazaar.conf": sampleOpenbazaarConf,
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
	"sample-openbazaar.conf": &bintree{sampleOpenbazaarConf, map[string]*bintree{}},
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
