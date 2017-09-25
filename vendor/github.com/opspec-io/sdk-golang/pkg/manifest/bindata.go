// Code generated by go-bindata.
// sources:
// github.com/opspec-io/spec/spec/pkg-manifest.schema.json
// DO NOT EDIT!

package manifest

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

var _githubComOpspecIoSpecSpecPkgManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x7b\x73\xdb\x36\x12\xff\xdf\x9f\x02\xa3\x66\xa6\xf2\xc5\x92\x9c\x47\x73\x57\x77\x3a\x1e\x4f\x9a\xce\xe5\xa6\x4d\x32\x97\x26\x37\x53\x5b\xe9\x40\xe4\x4a\x42\x4d\x12\x2c\x00\xca\x71\xef\xf2\xdd\x6f\x00\x50\x7c\x80\x00\x45\x52\x64\xad\xbc\xfe\x89\x4d\x60\x17\xd8\x1f\x16\x8b\xc5\x2e\x00\xff\xf7\x08\xa1\xd1\x3d\xee\xad\x21\xc4\xa3\x33\x34\x5a\x0b\x11\x9f\xcd\x66\xbf\x73\x1a\x4d\xf4\xd7\x29\x65\xab\x99\xcf\xf0\x52\x4c\x4e\x1f\xcf\xf4\xb7\xaf\x46\x27\x92\x4e\x10\x11\x80\xa4\x8a\xb1\x77\x8d\x57\xf0\x33\x8e\xc8\x12\xb8\xd0\xa5\x3e\x70\x8f\x91\x58\x10\x1a\xc9\x3a\xdb\x42\xb4\xa4\x0c\x61\x94\x92\xe8\xaa\x31\xa3\x31\x30\x41\x80\x8f\xce\x90\xec\x12\x42\xa3\x08\x87\x90\xfd\x56\x65\xf7\x02\x87\x80\xe8\x12\x89\x35\x20\x1a\x2b\x36\xaa\x9a\xb8\x8d\x55\x97\xb8\x60\x24\x5a\x8d\xd4\xe7\x0f\xba\xd4\x60\xe1\xe2\xfc\x43\xfe\x6b\xdb\x06\x48\x14\x27\x82\x17\x79\xdf\x63\xb0\x94\xb5\xbf\x9a\xf9\xb0\x24\x11\x91\x5c\xf9\x2c\xc6\x0c\x87\x65\x52\x9a\x88\xce\xb4\x2c\x89\x76\xd3\x71\xcf\xe8\xec\x06\x18\xaf\x47\xe2\xad\xae\x61\x43\xc1\xd1\x06\x84\x6f\x81\xa5\xcd\x1c\xa5\x4d\x8d\x18\xfc\x91\x10\x06\xfe\xe8\x0c\x5d\x16\xc6\xf6\x08\xa1\xb9\x2a\xc7\xbe\xaf\xe8\x71\xf0\xaa\xa8\x07\x4b\x1c\x70\x48\x35\x29\x6b\x22\xd7\x0f\xcc\x18\xbe\x7d\xa5\xc0\x28\x48\x90\xa9\x64\xa1\xf8\xc4\x21\xde\x85\xac\x82\x14\x9e\x20\x80\x49\x31\x71\x64\x1d\x6b\xba\xf8\x1d\x3c\x91\x7f\xb7\xe8\x6b\xde\xa7\xd2\x27\x77\xe5\x1a\x8d\xcc\x8a\x6d\xba\xb6\xfd\xf7\xe1\xc4\x64\xb5\xc4\x49\x20\xea\xd8\xe8\xee\xd5\x72\x21\xfc\x35\x78\x0c\xac\x6c\x0c\xf8\x9e\x6b\xad\x50\x4c\x11\xe1\x88\x6b\xc2\x13\x57\xeb\x0b\x4a\x03\xc0\x51\x7d\xfb\x1e\x8d\xb8\x60\x98\x44\xa2\x8a\x96\x53\xed\x54\x17\x9e\x16\x28\xcb\x4d\x1c\x39\x9a\xab\x55\xbc\x23\x93\x3c\x23\xdd\xa9\xaf\xaa\x52\x45\xe9\x51\xae\x20\xe9\xef\xf3\xd2\x74\xac\x08\xe1\xd4\xea\x62\xa5\xce\xaa\x1a\xe2\xf7\xcf\x05\x84\x26\xca\xe6\x20\xff\xeb\xf5\xcb\x17\xe8\xb5\x32\xfa\xe8\x72\x4b\x83\xae\xe1\xf6\x86\x32\x7f\x3e\x96\xcb\x05\x3f\x9b\xcd\x04\xa5\x01\x9f\x12\x10\x4b\xb5\x5c\xac\x45\x18\xa4\x6b\xc6\x0d\x23\xab\xb5\x98\x14\x16\x94\xc9\x06\x07\xc4\xc7\xb2\x85\xc9\xe9\xe9\x57\x1c\x3c\xf5\xe3\x37\xd3\x07\xa7\xc7\x25\xed\xc9\x64\x22\x91\x80\x15\xb0\x72\x61\x48\x22\x12\x26\x72\xf2\x9f\x1e\x59\x86\x57\x96\xb7\x17\x30\xa5\x19\x4a\xc0\x07\x7d\x0a\x48\xda\x4a\x47\x7a\x16\xed\x41\x26\xda\x93\xe9\xb7\x86\x64\x34\x82\x97\xcb\x92\xee\xcb\x7f\x0d\xe7\xb3\x84\xc5\x3d\x9d\x4f\xea\x59\xda\x60\xe9\xde\x5a\xd9\x7c\x94\x7f\x9b\x5b\x87\x25\xb7\x0e\xad\xd5\xcf\x20\x1d\x68\xa8\x2a\xd3\xac\x25\x2c\x45\x61\x93\x88\xfc\x91\x40\x6b\x41\x0b\x64\x43\x09\xf9\xc8\x31\xd5\x2a\xab\x50\x3b\xfb\x5e\xf6\x27\x09\x73\x3b\x20\x59\xa1\xcb\xfd\xf8\x81\x30\xf0\x04\x65\xfd\xba\x20\x3e\x61\x87\xe7\x80\x54\x3c\x6c\x55\x13\x6d\x70\x90\xc0\x77\x52\x62\xbc\xe0\x34\x48\x04\xa0\x18\x8b\x35\x5a\x32\x1a\x22\x46\xa9\x90\x78\xc4\xd7\x2b\x44\x19\x62\x10\x60\x41\x36\x69\x0d\x69\x30\x59\xcc\x40\x80\xaf\x6b\x4b\x4f\xc4\x27\x0c\x91\x08\xdd\xac\x89\xb7\x4e\x1d\x56\xe9\x97\x48\xef\xd8\xe9\x94\x34\x11\xac\xbd\x4f\xe4\x67\x43\xdb\xd9\x2f\x3a\x20\xa7\x45\xaa\x94\xd5\x65\x59\x92\x00\xdc\x13\x20\x2f\x75\xcd\x80\x1f\x49\x00\xbd\x2a\xbf\x6c\xf2\x8b\xf6\xdf\xb5\xf6\xcb\x51\xf8\x24\x14\x5f\xa9\x93\x55\xf3\xe3\x24\x08\x9e\x32\xf0\xcb\x5e\xba\x43\x5b\x0d\x94\x24\x1d\x44\x82\xe0\x80\xa3\x84\x83\x8f\xfc\x44\x8e\x02\xc2\x89\x58\xcb\xef\x9e\x5a\xce\xd0\x0d\x11\x7a\x1c\x39\x4d\x98\x07\xe9\xec\x20\x21\x5e\x81\xd4\x88\x62\xfc\x04\xd5\xcd\x89\x84\x03\x33\x62\x29\xc8\xbd\xea\x6b\xd5\xa2\x01\x16\xe0\xbf\x36\xb4\xa3\xb8\xf0\xc7\x98\x73\xb9\x6c\xf7\xc0\xb5\x32\x62\xf6\xc1\xc8\xe4\xb0\xf5\x62\x3b\x48\xed\xd7\xf0\x28\x09\x17\xc0\x76\xed\xbb\xaa\xb5\xba\xc7\x08\x82\x40\x79\xc6\xcd\xbd\x42\x49\x30\xd0\x8e\xe4\xe1\x43\x87\x9b\xa4\x37\xaa\xa5\x22\xbb\x63\xed\x18\xf1\x2a\x60\xc5\x79\x6d\x77\x9c\xa3\xdb\x96\xc0\x48\x82\xa1\x80\x71\xf9\x8f\x77\x00\x0c\x44\x49\xd8\x06\x17\x59\x7f\x28\x58\x5c\x5b\xf4\xe6\xb0\x6c\x29\x34\x10\xbb\xa5\x5f\x52\x16\x62\x73\xf5\x69\xba\xbf\xcc\x26\xb0\x6d\x87\x6d\x03\xf2\xdf\xda\xf6\x70\x65\x79\x75\x17\xd1\x02\x94\xe5\x75\x71\x30\x16\xd3\x4a\x79\x3a\x7c\x97\xd5\xfd\xe8\x96\xa5\x51\x32\x6f\xb9\xe3\x0c\xf1\xfb\x34\x50\xd0\x26\x92\x23\x49\x86\xd2\x12\x87\x92\x98\x43\x6e\x84\x6b\x5a\x0b\xa1\x49\x06\x12\xe2\x71\x17\x21\x92\x40\x90\x38\x80\x76\x76\x2c\xa7\x1a\x2a\xee\xd4\x41\x94\x88\x56\xe6\x5c\x9d\x0c\x11\x15\x43\x29\xd3\x37\x8d\xc2\x15\x35\x76\xb5\x28\xd6\xd6\x6e\x34\x16\x4c\x11\x0c\x25\x9a\x4b\xc7\xfe\xaa\x45\xa6\x95\xb3\x6c\x71\x9b\xdc\x9b\xbf\x62\xb9\xcb\x11\x7e\xa1\xcd\x6b\x9f\x1b\xc0\x54\xa3\x0f\x6e\x0b\xe8\x5e\xf4\xfa\xd8\x70\xa5\xeb\xd4\x9d\xe6\x60\x6a\x55\xee\xb0\xb6\x75\xe5\x41\x28\x6f\xec\xb4\xa2\xed\xda\x0e\x54\x6b\xed\x93\x87\x79\xe5\x52\xcf\x5d\x4b\x78\x4e\x38\xd4\xc2\xe1\xf2\x82\xbb\x66\x64\x3a\x8a\x5a\x24\x1c\x4a\x54\x97\x2d\xee\x24\x6a\x41\xf1\x1a\x4b\xb9\xa5\x19\x4a\x40\x73\x1d\xed\xee\xba\x57\x8d\xa1\xdd\x75\x77\xda\xdd\x5a\x1c\xe2\xc1\x87\xfa\x89\x03\x09\x63\xce\xa2\x1a\x3b\xd3\x68\x19\xae\xc9\x2c\x39\x00\xc3\x42\x00\xeb\x38\x4b\x2a\xc4\x43\xc1\xf7\xf7\x43\x85\xaf\x41\x63\x0d\xf3\x6f\xc3\x83\xf8\x8f\x6e\x09\xd3\x36\x0b\x79\x2f\xd9\xd6\x9d\xbb\x4f\x1f\x62\x88\x7c\x88\xbc\x96\x60\x17\xe9\x86\x02\xf9\xa3\xc8\x4a\xd7\xb9\x98\xed\xd3\xcf\x9f\x64\xb0\xc8\x8c\x9d\x8c\xa2\x24\x08\xaa\x5e\x6e\x6a\x80\x4a\x9f\xe7\xbb\xad\xc6\x67\x11\x90\xad\xba\xac\xbb\x81\xf9\x2c\x02\xb2\x1d\x80\xf9\x3c\x82\x08\x1d\x80\xf9\xd8\x82\x46\x35\x22\xee\x11\x1b\xd1\x5c\xdd\xb1\x91\x62\xb9\x2b\x36\xf2\x52\xd5\xe9\x35\x36\x92\xd6\x3b\xd8\xd8\x88\xcd\x7a\xef\x1f\x1b\xd1\x5c\xef\x36\x36\x52\x3b\x93\x0e\x2b\x36\x52\x1e\x84\x72\x6c\x84\x53\xef\x1a\x6a\xf4\xba\x58\xbe\x53\x4b\x8d\xd1\x7a\xad\x68\x6b\xf5\xdd\xa5\xd7\xba\xd9\x3b\xd2\xeb\xf6\x0a\xa9\xbb\xfb\x49\x9c\x8f\x48\x91\xb7\x2b\x8b\x02\x6f\x57\x20\xad\x5a\xeb\x4b\x5e\x3d\xfd\xec\xb8\x61\x51\x01\xec\x8b\x1b\xd7\x15\x98\x4f\x72\xab\xd4\x38\x38\x67\xcf\xab\x9b\x66\x39\x89\x81\x71\x50\x47\xd1\x4a\x58\x68\xea\x41\xd0\x30\xc3\x4b\x6d\x53\xfd\x3e\x16\x30\x11\xa4\x74\x52\xa8\xc1\xe8\x66\x64\x48\xcb\xd6\xaf\x4c\xd3\x47\x66\x16\xd6\x36\x68\x2d\xce\x0e\xe4\x52\x1a\x65\xf3\xba\xf5\xaa\x06\x35\x69\xcb\xd9\x44\x9d\x2f\x9b\xc8\x19\xb6\x0b\xbc\x0b\xa4\x49\xd2\x23\x69\x0c\x96\xc0\x20\xf2\x00\x61\x8e\xd4\xc4\x04\x1f\x2d\x6e\xd1\xe5\x8a\x88\x75\xb2\x98\x7a\x34\x9c\x69\x82\x99\x4f\xa4\xb8\x8b\x44\x72\x9a\x65\x74\x39\xde\x3b\x28\x04\x03\xd8\x16\x3c\x98\x3e\x78\x94\xb3\xe8\x17\x60\x13\x90\x7e\x70\x86\x10\x13\x4b\xe8\xa2\xd6\xee\x48\x92\xa1\xb4\xf2\x61\xaf\xa0\x69\xe9\xfa\x41\x6a\x4d\xb9\x30\x0e\xfc\x35\x00\x6b\x4b\x35\x14\x5e\x8f\x7a\xc5\x2b\x93\xb1\x1f\xc8\x48\xbc\x79\xdc\x0e\x2e\x49\x31\x14\x54\x8f\x7b\x85\x4a\xc9\xd6\x1b\x4c\x4f\x5a\xc3\xf4\x64\x28\x98\xbe\xe9\x1b\xa6\x27\x3d\xc1\x94\x30\xd2\x0e\xa5\x84\x91\xa1\x40\x7a\xd2\x2b\x48\x52\xb2\x7e\x30\xe2\x10\x6e\x1a\x1c\x2b\xbc\x40\x1c\x42\x1c\x09\xe2\xa1\xf4\xda\xb5\xb9\x4c\x6a\x46\x12\x23\x8d\xdd\xd9\x6c\x96\x7f\x9a\xf5\x2a\x7d\xda\xe7\x7a\x00\x8e\x6c\x25\xc6\xa9\xc3\x9f\x20\x5a\x89\x75\xcb\x4c\xbe\x26\x1a\xc8\x8f\x76\x65\x76\x77\x24\xf1\x1f\xd8\x25\xdc\x76\xf6\x90\x24\x74\x25\x5f\x77\x1d\x53\x38\x29\x0b\xb0\x8d\xc6\xd9\x8f\x2f\x7c\x6c\x71\xdc\x9a\xcd\xdf\xe7\x17\xb7\xef\xb0\x13\x4e\x8f\x0e\x74\x38\x6d\x30\x10\x38\x66\x76\xbc\xc6\xd8\xe5\x1b\xd9\x11\x83\x15\xbc\xef\x25\x84\xaf\xdb\xa9\x09\x75\x16\xca\x5b\x87\x3a\x15\x6d\xb7\x50\xa7\x16\xff\x60\x43\xf8\x03\x45\x4c\x35\x60\x77\x1a\xc2\xaf\x9d\x54\x07\x16\x95\x2d\x0d\x82\x71\x6f\xcd\xd4\x68\x03\xf1\x57\x5d\xf2\x4d\xb5\x87\x96\x46\x97\x93\xdf\xa6\x78\xf2\xe7\xc5\xe4\xd7\xd3\xc9\xb7\xf3\xfb\x1d\x2f\x72\xd4\x3c\xfc\xf1\x2a\x7f\x10\xc7\x31\xe4\x0d\xb9\x65\x37\xb6\x7b\xe0\x95\x5f\x7e\xed\x81\x59\xf1\x30\x75\x0f\xec\x8a\xf9\xc7\x1e\xd8\x15\xd3\x3e\x7d\xb0\x2b\x98\xd6\x26\x2e\x69\x77\x23\x6f\x1e\xe8\xb1\x19\x7a\xb3\x8e\xcb\xa8\xe7\x13\xc7\xb3\xd5\xae\x6a\x79\x83\x5b\x93\xee\x67\x6d\x0a\xc8\x36\xe0\xd3\xec\x6a\x44\x03\x46\x75\xe9\xf2\x56\x8c\xea\xfc\x35\xab\xdd\xe2\xde\xca\xbe\x0e\x7b\xab\x9a\x75\x16\xcb\xfd\x96\x87\x83\x00\xad\x18\x8e\xd7\x99\x4d\xfb\x0e\x71\x00\xb4\x75\x58\x20\x9a\xde\x90\x6b\x12\x83\x4f\xf4\x63\x67\xf2\xb7\xd9\x53\x1c\x04\xbf\x29\xb2\x3d\x47\xd1\xa3\x91\xc0\x24\x02\x26\x39\x76\x46\x3e\xde\x87\x5a\x5a\xfd\x20\x80\x60\x1f\x1e\x1c\x18\xc1\x26\x07\xeb\x58\x95\x05\xb6\x8d\x5a\xb9\x46\xe7\x8c\x5f\xc6\xa6\x8d\x43\xe4\x85\xe6\xa1\x6c\x9b\xee\x3c\xa5\x61\x88\x23\x1f\xb1\x24\x92\xbb\x73\x8c\xb2\xb6\xbe\x43\x74\x03\x8c\x11\x1f\x38\xc2\xd1\x2d\xe2\x20\x10\x16\xca\x4f\xd1\x81\xf0\x00\x36\x60\x09\xf0\xba\xbd\x7b\xe4\xf6\xf0\x6b\xc6\xa3\xee\x4a\x75\x79\x74\x2a\xa3\xad\x05\x26\xcc\xea\xfe\xd4\x1c\xe6\xb5\x01\xb5\x7d\xc4\x84\x00\x47\x24\x52\x30\xe4\xc3\x52\x21\xde\x75\xbe\x39\xad\xf6\x6e\x7c\xa9\x3d\x86\xf9\xd9\xf1\xb9\xf4\x1f\xae\xae\x66\x05\x17\xe2\x9e\x95\xca\xe9\x4b\x6c\xff\xd9\x48\x6c\x22\x8d\x6f\x48\x10\xa0\x05\xa0\x05\x4d\x22\x1f\x09\x8a\x38\x0e\xb3\x97\x17\xb6\x17\xef\xab\x11\x99\x0a\x84\xea\x90\xa2\xb5\xd2\x07\x3b\x6d\xd3\x1e\xfa\x84\x21\x06\x4b\xfd\x0c\x40\xa9\x57\xbb\x3b\x65\x3f\x65\x9a\x76\xcb\xf2\x75\x5e\xf9\x66\xd6\xaa\x08\xd3\xcc\xe7\xb5\x90\x8e\x20\xda\xbc\xc5\xbd\xe8\xe5\xb3\x68\x43\x18\x8d\x42\x88\x04\xda\x60\x46\xf0\x22\xe8\x55\x43\x2f\xdf\x7d\x7f\x07\x8a\x48\x22\xc4\x3d\x1a\xab\x6c\x1b\xba\x99\x69\xc5\x8c\x70\x78\xa7\xda\xa8\x35\x2a\xf5\x32\xb6\x7a\x99\xee\xd9\xd4\x1b\x27\xee\xbe\x75\x35\x6d\x69\xbf\xef\x5a\x5d\xa5\x93\xdf\x87\xb2\xfe\x48\xfa\x55\xce\x2f\xe6\xb3\xb6\x87\x72\xd8\x66\x56\xad\xfd\xb4\xac\xa9\xf2\x47\xba\xa8\x67\x8d\xff\x84\x74\xdc\xc1\x8c\xa2\x66\x45\x95\xe7\x14\xb6\xa7\x02\x04\xcd\x9e\xaf\xb1\xc2\xba\x87\x2d\xb0\xa8\x83\xf5\x81\x9e\x06\xcd\xe5\x64\x1d\x06\xca\x1a\x8e\x29\x81\x66\x7c\x9d\xf7\x37\xd6\x96\x17\x7e\x90\x65\x3c\xd4\xab\xc9\x25\x2b\x83\x3c\x1c\xc9\xd9\x9a\x1d\xa2\x50\xe9\x28\xf5\xaa\x14\x15\x6b\xbd\x9f\xd5\x35\xf9\x7e\xef\x48\xc5\x94\xd9\x03\x6e\xe6\x4e\x5a\xd6\x4b\x8d\x47\xf6\xb8\x55\xde\x5d\x41\xd5\x87\x35\xe5\x35\x71\x40\xa7\x62\x37\x33\xa0\x97\xca\x4e\x8e\x27\xfa\xff\xe3\xf3\xb1\xf0\xe2\xff\x25\x7e\x7c\x7c\xde\x50\xed\xff\x49\xb9\x40\x52\xe0\x31\x3f\x96\x3d\x5e\x10\x65\x09\xed\x8a\xbf\x23\x91\x87\xca\x01\xfa\x4a\xe7\xba\x68\x6a\x67\x35\xd3\xb1\x9e\x4e\x6b\x5e\x53\xec\xcf\xdc\xc1\xc2\xac\x52\x65\xb7\xb6\xd5\x8e\xf4\x70\x2d\xf6\x7d\x06\x9c\xa3\x10\xc7\x31\xa8\x35\x08\x6f\x8b\x6c\x47\x9b\x50\x13\x93\x3e\x24\xaa\xc2\x7f\xc6\xcc\x5d\x6c\x9f\xa0\xbe\x9b\xba\x57\x7d\x37\x96\xc2\x07\xc6\x50\xcc\x60\x49\xde\x97\xa1\xd4\xce\xdd\x81\x42\xf9\x32\x69\x72\xcd\xe0\xaf\x86\x92\x26\xe2\x23\x83\xf2\x86\xb2\xeb\x1f\x2a\xaf\x89\xda\x04\xfd\x0f\x65\xd7\x52\x0a\xbf\xf0\xa2\xa9\x58\xa3\x71\x39\x52\x52\x38\xed\xa0\xdc\x80\xdd\x67\x1a\x8e\x5c\x92\x96\x53\x2c\xce\x75\x37\xf5\x80\x0a\xdf\xe6\x7d\xe4\x66\xec\x69\x97\xdc\x6f\x3f\x32\xda\x6a\x71\x73\x28\x76\xc6\xcc\xd2\xa2\xce\xc1\x32\x1a\x9b\x51\xb2\xba\x6b\xdb\x75\x11\xb4\xf8\xda\xcc\x40\xee\x62\xb7\x8b\x25\xda\xc7\xa9\x44\xb2\x3f\x9d\x17\xd6\x3c\x79\x9c\x30\x32\xc9\x9c\xa0\x2f\x4e\xa6\xa5\xf5\xea\x9f\xb7\xc8\x4a\xcc\xc4\xad\x14\x15\x07\xe4\x4f\xe0\xe8\xf9\x8b\x57\x6f\x7e\xf9\xed\xc5\xc5\xcf\xcf\xb4\x3b\xf7\xf6\xe2\xa7\x37\xcf\xe4\x26\x2b\x3d\x42\xfe\x75\x5e\xe1\x4c\x17\x7e\x3d\x45\xcf\x97\xdb\x7a\x1c\xc9\x7d\xe0\x09\x22\x02\xfd\xfc\xe6\xf5\x2f\xea\x31\x36\xce\x93\x10\xfc\xb4\xc6\xf7\xe8\xde\x38\x67\x51\x63\x54\xf6\x75\x4c\x6a\xf3\x98\x59\xb5\x8e\x9b\xe8\xfe\x37\xbd\x7b\xc6\x56\x5a\xb6\x56\xf7\xc0\xd1\xbe\x3c\x6d\x17\x03\x33\x9e\x77\xbd\xc9\xae\xfe\xdd\x96\xac\xa8\x66\x52\xe4\xd3\xe1\xe5\x9b\x5f\xb2\xf9\x51\x98\x14\x7a\x3a\x14\x0a\xf5\xa4\x28\xd5\xae\x99\x1a\xaa\x82\x9c\x19\x05\x82\x2f\x53\xc3\xe4\x78\x58\xd1\x1b\xa7\x7f\xd3\xf0\x88\x44\xfd\xca\x20\x57\x49\x73\x65\xa8\x53\xec\x96\xec\x77\x5f\x90\xf8\x6b\xbc\x30\x1a\xef\xe1\x7e\x95\xf2\xa4\x36\x27\xac\x54\xa1\xb3\x2b\xb6\xe5\xe2\x72\xc8\xf6\x3e\x00\xe8\xb9\xae\x7f\x35\x04\x31\xeb\x60\x77\x28\x0b\xe9\x62\x6b\xde\x3e\x2f\xee\x0c\xa3\xe6\x71\xb0\x20\xa6\xdd\xeb\x0e\xa1\x65\xa9\xb6\x41\x69\xa9\xe6\x3e\x11\xa1\x36\x9b\x3c\x89\x63\xca\x84\xfc\xf1\xde\x58\xad\x0c\x88\x27\x0b\x2e\x88\x50\x17\xac\x10\xbf\x8d\x04\x7e\x8f\x6e\xd6\xc0\x20\xab\xb1\x8d\xea\x33\x88\x03\x9c\x46\x09\xc5\x1a\x74\x86\x07\xd1\xa5\x5e\x72\xb0\x40\x2c\x89\x4a\x97\xef\x6a\xff\x60\x5a\xfa\x07\xc3\xec\x2a\xa2\x8a\x5c\xb2\x6c\x1e\x4e\x4f\xa7\xa7\xd5\x63\xf5\xe3\xed\xd1\x8e\xf2\x01\x7a\x1e\x83\x37\xd3\x34\xd3\xb5\x08\x83\x63\x67\xff\xcc\x73\x6d\xb2\xe8\xdd\x38\x8d\xb6\x5d\x5d\x4d\x2d\x3f\x8e\xcf\xcf\xc6\x57\x57\x2a\x22\x77\x31\xf9\x15\x4f\xfe\x9c\xcc\xef\x8f\xcf\xcf\xae\xae\xa6\xa5\x4f\xc7\x7f\x3b\x3e\x3e\x57\xdf\xef\x17\xbe\x5f\x5d\x4d\xae\xae\xa6\xf3\xfb\xc7\xe7\xf7\x0a\x7f\x38\xed\xe8\xc3\xd1\xd1\xff\x03\x00\x00\xff\xff\x03\xf6\xb7\x9d\x9c\x6f\x00\x00")

func githubComOpspecIoSpecSpecPkgManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpspecIoSpecSpecPkgManifestSchemaJson,
		"github.com/opspec-io/spec/spec/pkg-manifest.schema.json",
	)
}

func githubComOpspecIoSpecSpecPkgManifestSchemaJson() (*asset, error) {
	bytes, err := githubComOpspecIoSpecSpecPkgManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opspec-io/spec/spec/pkg-manifest.schema.json", size: 28572, mode: os.FileMode(420), modTime: time.Unix(1506371717, 0)}
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
	"github.com/opspec-io/spec/spec/pkg-manifest.schema.json": githubComOpspecIoSpecSpecPkgManifestSchemaJson,
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
	"github.com": &bintree{nil, map[string]*bintree{
		"opspec-io": &bintree{nil, map[string]*bintree{
			"spec": &bintree{nil, map[string]*bintree{
				"spec": &bintree{nil, map[string]*bintree{
					"pkg-manifest.schema.json": &bintree{githubComOpspecIoSpecSpecPkgManifestSchemaJson, map[string]*bintree{}},
				}},
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
