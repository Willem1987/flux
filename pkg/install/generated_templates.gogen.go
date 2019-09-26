// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6872,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xdd\x6f\x1b\x37\x12\x7f\xf7\x5f\x31\x50\x0e\x48\x0c\x48\x2b\xbb\x6e\x8b\xc3\xf6\x5c\x5c\x9a\x0f\xd7\x97\xc6\x31\xe2\xf8\x0e\x79\xaa\x29\xee\x48\x4b\x88\x4b\xee\x71\xb8\x52\x05\xa3\xff\xfb\x61\xc8\xfd\xe0\xca\xb2\x5d\xe4\xed\xfa\xd0\xd8\xdc\xe1\x7c\x7f\xfc\x38\x9e\xcd\x66\x47\xa2\x56\xff\x46\x47\xca\x9a\x1c\x44\x5d\xd3\x7c\x73\x7a\xb4\x56\xa6\xc8\xe1\x2d\xd6\xda\xee\x2a\x34\xfe\xa8\x42\x2f\x0a\xe1\x45\x7e\x04\x60\x44\x85\x39\x2c\x75\xf3\xc7\xfd\x3d\xa8\x25\x64\x57\xa2\x42\xaa\x85\x44\xf8\xf3\xcf\xf6\x7b\xf8\x35\x87\xfb\xfb\xf1\xd7\xfb\x7b\x40\x53\x30\x19\xd5\x28\x99\x99\xc3\x5a\x2b\x29\x28\x87\xd3\x23\x00\x42\x8d\xd2\x5b\xc7\x5f\x00\x2a\xe1\x65\xf9\x9b\x58\xa0\xa6\x78\x90\xca\x66\x6a\xef\x84\xc7\xd5\x2e\x7e\xf4\xbb\x1a\x73\xf8\x8c\xd2\xa1\xf0\x78\x04\xe0\xb1\xaa\xb5\xf0\xd8\x32\x4b\x2c\xe0\xff\x84\x31\xd6\x0b\xaf\xac\xe9\x99\x03\xd4\xce\x56\xe8\x4b\x6c\x28\x53\x76\x5e\x5b\xe7\x73\x98\x9c\x9d\x9c\x9d\x4e\xe0\x05\x78\xd4\x3a\xa1\x00\x6f\x81\xa4\x13\x35\xc2\xbc\x42\xef\x94\x24\x36\xae\xb6\xca\xf8\x97\x04\x7c\x39\x6b\x19\xeb\x91\x0d\x7b\x56\x00\x74\xbe\x08\x3f\xa3\xdb\x28\x89\xaf\xa5\xb4\x8d\xf1\x57\x63\x42\x80\x8d\xd5\x4d\x85\x3d\xab\x59\xcb\x6a\xa5\xfc\x6c\x8d\xbb\x5e\x00\xb1\x17\xfc\x20\xb0\x3b\x19\xf8\xcd\xf8\x4a\x11\x02\x9c\x50\x15\xb8\x14\x8d\xf6\x1f\x6d\x81\x39\x9c\x7c\x7f\x72\x02\x2f\x60\x5b\xa2\x81\x8a\xb5\xc1\x02\x1c\x8a\x62\x66\x8d\xde\x4d\x61\x8b\xb0\xb5\xe6\xa5\x87\x05\x82\x58\x68\x64\x7f\xc8\xb2\xb2\xc5\x51\xcb\xf0\x05\x7c\x29\x15\x81\x22\x10\xe0\xab\x7a\x49\xd0\x10\x16\xb0\xb4\x0e\x56\x68\xd0\x09\xaf\xcc\x0a\x6e\x6e\x7e\x85\x35\xee\x28\x83\x4b\x03\x1f\xfe\x4e\xf0\xf3\x39\x9c\x66\xa7\x27\xd3\x9e\x4b\x27\x3b\x9a\x40\x20\x1c\xa6\x7a\x90\x65\x55\x0c\x62\x01\x02\x08\x6b\xc1\x49\xd1\x3a\x0a\xb6\xd8\xb3\x91\xc2\xc0\xd6\x29\xcf\x8a\x66\x87\xfd\xb7\x42\xd3\x3b\x03\xab\xda\xef\xde\x2a\x97\x3a\xb1\xc2\x42\x35\x55\x0e\x1f\xb1\xb2\x6e\x97\xda\x89\xb0\xb4\x5a\xdb\x2d\x5b\xd4\x8a\x56\x14\x4c\x6d\x88\xcf\x04\xc8\x86\xbc\xad\x14\x7b\x60\x6d\xec\xd6\xfc\x5e\x5a\xf2\xd4\xb3\x58\x2a\x8d\x53\xd8\x96\x4a\x96\xb0\xb3\x0d\x6c\x95\xd6\xd1\x28\x6f\xa1\xb0\x5c\x67\x7c\xcc\x97\xf8\x07\x07\x76\x6b\x58\xed\x9e\x81\xc3\xda\x82\x13\xbe\x44\x07\xbe\x14\xa6\x15\xbc\x52\xbe\x6c\x16\x60\xf9\x10\x41\xab\x35\x66\xf0\xd5\x36\x2f\xb5\x06\xa1\xc9\x76\x22\xc6\xce\x06\xe5\x41\x19\x6f\xc3\x1d\x69\x8d\x17\xca\xa0\x9b\xc2\x02\xb5\xdd\x66\x70\x83\x83\x57\x4b\xef\x6b\xca\xe7\xf3\xc2\x4a\xca\x38\xb1\x64\xc1\xa5\x83\x66\xce\xa5\x47\x7e\xbe\x6a\x54\x81\x34\x6f\x08\x67\xb5\x53\x1b\xe1\x31\xa4\x1e\x1b\x92\x95\xbe\xd2\x3d\xa7\x2e\x16\x44\xe5\x4c\x5a\xb3\x54\xab\xfe\x13\x40\x3c\xf8\x28\xea\x3c\x39\x4c\x0b\x69\x96\x5c\xfb\xd6\xb8\x64\xeb\x66\x81\xf3\xc8\x64\x48\xbf\x67\x63\xb2\x55\x54\xf2\x49\x29\x36\x08\x02\x0a\xb5\x5c\xa2\xe3\xa6\xd9\x71\x68\xab\x6a\x68\x8c\x21\x04\x91\x5d\x1a\x04\x6e\x2e\x1b\x55\x60\xe7\xf6\xa5\x5a\x55\xa2\x1e\x14\x51\xbe\x04\x61\x00\x8d\x77\xbb\x60\xc3\x5d\x24\xba\x9b\x82\x30\x05\x34\x46\xda\x8a\xbb\x75\xb8\x1f\xad\xfd\x18\xc2\x29\x4c\xd1\x73\x41\xb3\x09\x1c\x14\x52\x1b\xcf\x07\x11\x60\x37\x7c\x43\x04\x92\x6b\xcf\x46\x20\x74\x02\x6f\x41\x55\xdc\x27\xe1\xe2\xfa\x22\x34\x01\x78\xc5\x66\x91\x5a\x19\x65\x06\xe1\x6c\xdc\x06\x9d\x5a\x2a\x19\x1a\x36\xd4\x8d\xab\x2d\x21\x1d\xff\x05\x47\xf6\x5c\x62\xfb\x88\x5e\x64\x07\xb1\xbc\xbf\xe0\x38\x10\x6e\x35\x94\xe9\x23\x1e\x5b\xd5\x2b\xee\x1f\x94\xb8\x66\xdc\x82\x5f\x3c\xd2\x84\x1f\xde\x3b\xd0\x84\x3b\x77\xf6\x95\xf8\xa0\xff\x27\x13\xa2\xf5\xba\xc3\xd0\x27\x8d\x85\x49\x1e\x2b\x71\x02\xaa\x12\x2b\x8c\xd9\xcf\x17\x32\x78\xaf\x4c\x11\x6c\xae\xb8\xad\x38\x94\x43\xd6\xc6\x96\xa2\x51\x10\x72\xf3\x08\x57\x39\x08\x8c\x13\x40\xf8\xbe\xee\xcb\x66\x91\x15\x56\xae\xd1\x65\xd2\x56\x73\x37\x8f\x3d\x20\xfc\x33\xf7\xa2\x77\x5d\x17\x47\x9e\xf7\x8c\x05\x58\xaa\x17\x2b\x60\x4d\xb3\x9e\x26\x88\xc9\xa1\x65\xa8\x6c\xca\x2d\x3f\xcd\x4e\xbf\xcf\xbe\x1b\xd3\x5e\x37\x5a\x5f\x5b\xad\xe4\x2e\x87\xcb\xe5\x95\xf5\xd7\x0e\x29\xb5\xc2\x21\xd9\xc6\x49\xa4\xb4\x8f\x3b\xfc\x6f\x83\xe4\x47\x67\x00\xb2\x6e\x72\xf8\xe1\xa4\x1a\x1d\x56\xa1\xd5\xe7\xf0\xe3\xf7\x1f\xd5\x00\x13\xac\x4b\x2f\xcf\x86\xc8\x5c\x07\xc8\x70\x76\x72\xc6\x93\x53\x99\xa5\x75\x55\x48\x59\xa1\x7b\x6a\xad\x36\x68\x90\xe8\xda\xd9\x05\xa6\x1a\xb0\x4b\x2f\xc6\x53\x3b\x8a\x8a\x0c\xc7\xc7\xc2\x97\x39\xcc\x45\xad\xa2\xa7\x37\x3f\xce\x55\x81\xc6\x2b\xbf\xcb\xea\x66\x91\xd0\x2a\xa3\xbc\x12\xfa\x2d\x6a\xb1\xbb\xe1\xfa\x2c\x28\x87\x1f\x12\x02\xaf\x2a\xb4\x8d\x3f\xf0\x8d\x87\xac\xfa\xff\x50\x35\x29\xda\x51\x60\x0e\xc3\x23\x88\x63\xee\x3a\x6a\x86\x5e\x06\xcd\x8a\x39\x51\xc9\x38\xcf\x46\xe4\x09\xda\xb6\xfd\x66\xc5\x21\x03\x65\x62\xce\xbd\xa4\x78\x87\xa8\x9c\x8f\xda\x64\xe7\xb3\x4f\x46\xef\x72\xf0\xae\x41\xe6\xc6\x18\x28\x74\xa8\x45\xdb\xd8\xb9\xa4\x6a\x74\x4b\xeb\x24\x32\xd3\x08\x7a\x18\xf3\x3c\xa6\x78\x8a\x4b\xc6\xba\x6f\x84\x6b\x75\x8f\x64\xdf\xa6\x7e\x52\xa3\x97\x46\xea\x26\x74\x4e\x86\x6e\x71\xc0\x75\x5d\x35\x62\x83\x67\xa0\x4c\x07\x66\x7e\xe2\xab\x7b\x30\xa3\xef\xae\x50\xa0\xd4\xc2\x31\x64\x5b\xd8\x4d\xd2\x00\x9e\x80\x01\xb1\x3d\xa6\xc6\x3b\x6b\xfd\x3c\x23\x2a\x1f\x35\x40\x98\x91\xd4\xc9\x30\xa2\x26\x51\xf2\xb4\x23\x49\x38\xa0\xd9\x28\x67\x4d\x18\x08\x71\xd6\x4e\x3e\xdc\xfe\xf2\xee\xcd\xa7\xab\xf7\x97\x17\x93\x38\x02\xa6\xec\x0f\xbb\x41\xe7\xc6\xf3\x3a\x61\x13\x46\xdc\x62\x17\xa7\xa9\xd7\x87\x6c\x7c\x30\x68\x1f\xda\x38\x24\x27\x13\x3f\x6a\x28\xcf\x3c\x7e\x78\x74\xd2\xb8\x45\x27\x50\xa4\xd5\x2e\xc4\x24\x61\xb1\x0f\x68\xd2\xa0\x07\x34\xd3\x41\x6f\x61\x40\x68\x8f\xce\x30\xb4\x7e\xa0\xf1\xd2\xd9\x8a\xd3\xa2\x43\x2c\x53\x10\xc4\xe9\xd6\x4e\x55\x76\x83\xb6\x72\x4d\x0f\x83\x8d\x66\x93\x1f\xf0\xcb\xe0\xee\x91\x5f\x36\x42\x37\xf8\xc0\x27\xcf\x25\xf1\x7e\x0e\x74\x33\xf7\x89\x0c\xe0\x91\x3f\x1e\xf5\x4f\x0c\xfb\x47\xf2\x92\xa9\x22\xba\x19\xd1\x8d\xfb\xc3\x73\x95\xb7\x15\x0c\x4a\x2c\x50\x53\xd7\x7a\x07\xbf\x7e\xf9\x72\x0d\x0b\x41\x4a\x82\x68\x7c\x09\xd2\x61\xe8\xa4\x42\xc7\xa9\x3e\xbc\x07\x98\xe1\x46\x89\x60\xf8\xdd\xc5\xe5\x97\xdf\x5f\xdf\x7e\xf9\xf5\xf6\xe6\xdd\xe7\xbb\x60\x6e\x7f\xf4\xe1\xdd\xd7\xbb\x51\xc2\x6f\x84\x53\xfc\x9a\xa3\x0e\x20\x27\x0c\x23\x7c\xd9\x8b\xdf\x7b\x67\xab\x71\x0c\x23\xd9\x67\x5c\xe6\x23\xcb\x47\x58\x91\x1b\x1b\x9b\x30\x38\x80\x7d\x9e\x8f\xfc\x11\x5d\x10\xdf\xa8\x58\xf0\x24\x96\x42\x96\x58\x70\x6a\xa5\xb9\xdd\xc3\x6a\xf6\x14\x73\x9f\x26\x5c\xac\x6b\x71\x73\x72\xa1\x7d\x63\x87\x8b\xd3\x20\x84\xdf\x86\xad\x8f\x7d\x89\x94\xe6\xc2\x80\x5e\xfd\xd6\xb2\x96\x0d\xfb\x29\x54\x5c\x58\x08\x84\x44\x84\xd2\x6e\xc3\xfb\xd7\x1a\x83\x32\x84\x4c\xf9\x71\xee\xcc\x66\xbd\x01\xe1\xf1\xc3\xc2\xcf\xfb\xa3\xac\x05\x7d\x19\x6d\x64\x26\x75\x43\x1e\x5d\xc6\x0d\x5c\xa7\x2e\xb9\xa5\xd8\x6b\x06\x57\xbc\x89\xa4\x97\xd7\x23\xa3\xb8\xed\x10\xfa\xf0\xbe\x1e\x67\xf6\xa0\x43\x47\xcf\xd9\xe5\x1d\x53\x86\x17\x6f\x32\x82\x52\x8d\x5b\xea\xf3\xa3\x11\xca\x54\x04\x55\x43\x61\x03\x10\xbc\xa7\xb0\x88\xe5\xb4\x08\x83\x2d\x60\xbc\xf0\xf0\x7f\xd5\xbd\xa6\x8f\x53\x5d\xba\xe6\x12\xcb\x90\x13\x38\x79\xff\x8f\x14\xe1\x61\x10\x07\xdc\xac\x50\xee\xfc\xc1\xd8\x4b\xd5\xfa\x9c\x20\xcc\x21\x78\xb7\x9f\x7f\x8b\x0b\x0a\x61\x56\xf1\xdb\x85\xf2\xe1\xd1\x4c\xca\x5b\xb7\xeb\xdb\xf5\x7b\x46\xc6\x09\xbb\xa7\x6a\x8e\xd3\x26\xb1\xbd\x2d\x99\x83\xe5\x94\xd6\x42\x87\x9d\xff\xf6\x2a\xad\xcc\xe3\x7c\xf8\xfd\xc3\xbb\xaf\xc7\xff\x8c\x4f\xf7\x00\xab\x1b\x42\x37\x1f\x94\xcd\xd2\x42\x67\xff\x70\x39\x35\x4e\x9f\xdf\xdf\x43\x76\xa1\x3c\x1b\x1b\x56\x71\x63\x8a\x85\x13\x46\x96\x1d\xd1\x2f\xe1\xb7\xb8\x94\x53\xcb\x70\xc4\xfd\x8b\x0e\xdd\x64\x0c\xc7\xf7\x6e\x42\xa6\xd0\xbf\xac\x32\xc9\x85\xc9\x74\xd2\xee\xf6\x34\x61\x7a\xfd\xe9\xa6\xe6\x90\x13\x4f\xc6\x57\x57\x25\x8c\x5a\x32\x26\xe7\x1a\x22\x55\xa0\x8b\xe1\xd8\x7b\xd9\x84\x9d\x84\x25\x84\xc6\x14\xe8\xf6\x62\xec\x50\x0b\xaf\x36\x18\x20\x27\x75\x19\xb8\x1a\xc5\x79\xaf\x26\x7b\xe3\xa8\x59\x14\xca\x9d\x4e\xe3\xbf\xdf\xf5\x8b\xca\xc1\x39\x61\x11\x79\xc8\x39\x61\xbb\xd7\x79\xb5\xa3\x3a\xc0\xe0\x96\xd0\x1d\xba\xcf\xc1\xed\x23\xc7\x34\x70\xf8\xfe\xbb\x4a\xa8\x83\x0a\x20\x7f\xe8\x38\x74\x54\xc3\xaa\xf5\x60\x38\x90\x5b\xc9\xd6\xb2\x43\xd1\x84\xf5\x1d\xfb\x89\x27\xb6\xf2\x7b\x0f\xf0\xd4\x57\xed\xec\x6b\x27\xdb\xf9\x13\xa3\xae\xbb\xd1\xf2\xe2\x5b\xe7\xff\x58\xe3\x0e\x54\xf1\x73\x4f\xf6\x04\x9c\x49\xb4\x62\x16\xc2\x37\x0e\x47\x5b\x80\x03\xb2\xc2\xe7\xdd\xac\xa7\xa7\x51\xbb\xea\xba\x35\x28\x0f\xa5\xa0\x30\x8a\xad\xd1\x3b\x10\x52\x22\xc5\x8e\x5e\x62\x5c\xa4\xbd\xea\x76\x36\x77\x4b\xa1\x09\xef\x8e\x0f\x48\xeb\xee\x8f\x1d\x4c\xde\x35\xd2\x47\x41\xdb\xf0\x0e\x67\x6c\xd6\x78\xa0\x9d\x91\xb0\xb0\x76\xbd\x46\xac\x39\x5d\x7b\x19\x93\x95\xf2\x93\x29\x54\x28\xd8\x53\xdc\x89\x40\x84\xc7\x71\x9b\xc1\x4d\x4d\xde\xa1\xa8\xfa\x54\xde\xd7\x86\x59\xcf\xc8\x0b\x8f\xe7\xdc\x19\x1e\x0d\xb8\xc1\x3f\x7c\x17\xf5\x64\x54\x09\x03\x93\x4e\xc6\xa4\x1b\x24\x09\x93\x57\x98\xad\xb2\x29\xfc\x07\x19\x12\xbe\xd1\xb6\x29\x8e\xb3\xb0\xd9\xf1\x76\xcd\x0f\x0b\x82\x5a\x38\xaf\x64\xa3\x85\xeb\xbc\xd8\x72\xd9\x9f\x81\xad\xd4\xf3\x2d\x71\x03\x94\xcc\x2b\xdb\x32\xdf\x6c\x6b\xdd\x9a\xfa\x57\xe2\xde\xb5\x20\xe8\x5c\x2c\xe4\xe9\x77\x67\x0f\xff\x9f\x1a\x7c\x83\x6e\x73\x60\x21\xcf\x78\x78\x00\x00\x9c\xaa\x3f\xa5\x93\x48\xac\xb9\x8b\xc7\x58\x11\xfa\x64\xcb\xff\x32\xf9\x43\x41\xb2\xf1\x67\x13\xc3\xe6\x2a\x60\xd2\x6c\x54\x92\x5a\x91\x47\x33\x6b\x55\x38\xcf\xcf\x4e\xce\x4e\x8f\xda\x32\x7e\x5d\x14\x2a\xee\x03\x78\xce\xbc\x66\x9c\x39\xea\x97\xc3\xf7\x01\x6a\xdc\xdf\x83\x0b\x53\xeb\x99\xdb\xb3\xf0\xe7\x96\x51\xe9\x0f\x3f\x75\x02\x3e\xd5\x2d\xfb\xb7\x57\x37\x1d\x46\xa0\x69\x8b\xdd\x1b\xd7\x22\x06\x30\x85\xf5\x04\x36\x10\x43\x25\x76\x61\x8f\xa2\x37\xc3\x36\xcd\x90\xb6\x76\xdd\xd4\xa0\x88\x1a\x24\xb0\x06\xc8\x56\x08\x1f\x9a\x05\x3a\x83\x1e\x89\xb9\x37\x35\x0d\xcb\xb2\xc2\x50\xb7\xaa\x99\x5c\x59\x83\x93\xf4\xcb\x9b\xa0\x40\xba\x2e\x8b\xc2\x69\xbc\x41\xeb\x30\x78\xd0\x6f\xf4\xa5\x7f\x1e\x4c\x4e\x27\x47\xff\x0b\x00\x00\xff\xff\x19\x66\xdf\x16\xd8\x1a\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}