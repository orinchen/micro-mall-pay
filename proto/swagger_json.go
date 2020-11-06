package proto

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _proto_micro_mall_pay_proto_pay_business_pay_business_swagger_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4f\x6f\xdb\x36\x14\xbf\xfb\x53\x10\xda\x8e\x41\xdd\x65\xc3\x0e\xb9\xa9\x8a\x96\x18\x49\x6c\x43\xb6\x81\x16\x43\x20\xd0\xd2\xb3\xc3\x42\x22\x15\x92\x4a\x63\x0c\x01\x76\xec\x65\xdb\xa5\x1b\x06\xec\xb0\xdb\xae\xbb\x0c\xeb\xb0\x7c\x9d\x26\xeb\xc7\x18\x28\xcb\xb6\xfe\x50\xfe\x23\xbb\x5b\x0a\xc4\x40\x80\x98\x8f\x7c\x7c\xff\x7e\x8f\x3f\xd2\xdf\x34\x10\x32\xc4\x2b\x3c\x1e\x03\x37\x0e\x90\xb1\xff\xe4\xa9\xb1\xa7\xc6\x08\x1d\x31\xe3\x00\x29\x39\x42\x86\x24\x32\x00\x25\x8f\x38\x93\xac\x19\x12\x8f\x33\x37\xc4\x41\xe0\x46\x78\xe2\x4e\x07\xd5\x7f\xc3\x58\x10\x0a\x42\xe4\xbe\x3c\x49\xe4\x89\x5a\x84\x8c\x2b\xe0\x82\x30\xaa\x94\xa5\xff\x22\xca\x24\x12\x20\x8d\x06\x42\x37\xc9\xe6\x1e\xa3\x22\x0e\x41\x18\x07\xe8\xeb\xe9\x2a\x1c\x45\x01\xf1\xb0\x24\x8c\x36\x5f\x0a\x46\xd5\xdc\xf3\x64\x6e\xc4\x99\x1f\x7b\x6b\xce\xc5\xf2\x42\x2c\xbc\x6a\x5e\x7d\xd6\x94\x1c\xfb\xd0\xc4\x9e\xc7\x62\x2a\xe7\x22\x84\x8c\x31\x64\xbf\xaa\x28\xc5\x61\x88\xf9\x44\x19\xfe\xfe\xfb\xb7\x77\x3f\xfc\xf4\xfe\x8f\xdf\xee\x5f\xbf\x4d\xfd\x4a\xa6\xb0\x08\x78\xb2\x6f\xcb\x57\xd3\xba\x78\xf2\x2c\x8d\x41\x0f\xf8\x15\xf1\xc0\x3d\x02\x69\xa6\x7b\x65\xd6\x71\x10\x11\xa3\x02\x44\x6e\x47\x84\x8c\xfd\xa7\x4f\x0b\x43\x08\x19\x3e\x08\x8f\x93\x48\xa6\x51\x34\x91\x88\x3d\x0f\x84\x18\xc5\x01\x9a\x69\x7a\x92\x51\x3f\xb5\xde\xbb\x80\x10\x97\x94\x21\x64\x7c\xca\x61\xa4\xf4\x7c\xd2\xf4\x61\x44\x28\x51\x7a\xf3\x09\x5c\x18\xed\xa4\xea\x8d\x9c\x92\x9b\xcc\xb7\x9b\xec\xbe\x86\x0f\x23\x1c\x07\x72\xb5\x0f\x14\xc5\x14\xae\x23\xf0\x24\xf8\x08\x38\x67\x7c\xee\xca\xb6\x9e\xf0\x98\x4a\x12\x82\xad\x94\x2e\xb1\xbb\xa1\xf1\xc0\x88\x30\xc7\x21\x48\xe0\x8b\xf2\x9a\x7e\x0a\xee\x50\x1c\x26\xf0\x60\xaf\x28\xf0\xa2\xc1\x24\x71\xf1\x32\x06\x3e\x29\x8a\x38\x5c\xc6\x84\x83\xaa\x96\x11\x0e\x04\x14\xc4\x72\x12\x25\x6a\x85\xe4\x84\x8e\x8d\xaa\x30\x57\x18\x93\x16\xb5\x9b\x68\xf9\x50\x36\x15\xa4\x40\xe3\xb0\x10\xaa\x64\xbc\x0b\x5c\x21\x71\xaf\x38\x6e\xb1\x30\xc2\xb4\x68\x03\x42\x46\x6f\x22\x24\x84\xf9\x7c\x9d\xef\x15\x8b\x68\x56\x5d\x33\xfd\x1b\x46\xc8\x63\x84\x3e\x88\xf0\x58\xed\x17\xe5\x10\x0c\x7a\x87\x6b\xfb\xaf\x14\x68\xab\x39\xb3\xc6\x90\x78\x5c\xac\x63\x4d\x8f\x5a\xe8\x39\x6f\x14\x42\x69\x44\x4c\x54\x77\xc5\xbb\xd7\xbf\xdc\xdd\xfe\x5d\xa7\x2b\x5a\x1c\xb0\x84\x8f\xaf\x31\xe6\xec\x7e\xec\x8d\xe9\xa7\x02\x6c\x43\xe6\x97\xc0\x34\xc5\x99\x4e\x92\x81\x99\xe4\x71\x11\x65\x3b\x4c\xda\x65\x0c\x42\xae\xe3\xfb\x4e\x90\xd4\xc8\x44\x31\xc3\x3f\x22\x3c\xc9\x72\x8f\xa5\x30\xfb\xe7\xf6\xd7\x77\x7f\x7d\x7b\xff\xe6\xcf\xbb\xef\x7e\xbc\x7f\xf3\xfb\xbb\xdb\x9f\x37\x03\x5b\x5f\xed\xd8\xc5\x93\x8f\x08\x67\x33\x93\x1f\x21\x96\x7e\x1e\x38\xc4\x16\xf9\xfa\x3f\xd0\x35\xbf\x46\x64\xcc\x5b\x90\xfe\xac\x9d\x69\x13\xb0\xa9\xe4\x39\xfc\xcd\x0e\x71\x36\x7c\x09\xde\xe2\x40\x52\x77\x8d\x08\xb8\x24\x05\x9c\xe4\x99\x56\x01\x41\xab\xa3\x95\x5a\xd1\x57\x8b\xb5\x65\xb0\xa0\x29\x1b\xeb\xb6\x18\xa1\xd5\x8a\x87\x38\xc0\xd4\x2b\xa9\xad\xa2\x9d\x37\xda\x2e\x56\xe5\x8a\x26\x9e\x05\x52\x54\xa6\x43\x65\x9e\xa8\x61\x88\x05\x6e\x38\xaf\x9c\x2a\x3e\xa8\x33\x74\x1e\x97\x5a\x56\xe6\xe9\x5a\x96\xa8\x69\x8d\x99\x93\x33\xbd\x25\x61\xc8\xe8\xbc\xb3\x6d\x51\x85\x1e\xf3\x6b\x54\x88\x03\xd2\x52\x0b\xb5\x05\x12\x8a\xf1\xee\x8a\x43\x7b\xf0\x6e\xe1\xef\xf4\xb2\xb5\xae\x79\x7b\x8f\x70\x5d\x95\x91\x5d\x94\xa0\x2a\xe6\x3a\x7e\xe7\x40\xb0\x34\x6b\xba\x32\xaf\x1f\x02\xcd\xdb\xc6\x16\xfe\x83\x3a\x49\xdc\x80\x88\x22\xe7\x98\x6b\xc2\x9c\xe3\xfc\x49\x6c\x10\x09\x61\x91\x78\x6d\x52\x89\xd3\xe3\x4b\x4f\x26\x56\xfa\x3f\x83\x7f\xad\x3e\xd8\x1b\x58\x96\xdd\xeb\x65\x7b\xa1\xed\x38\x1d\x27\xdf\x1c\x6d\xc7\x6d\x77\xfa\xae\xfd\xbc\xd5\xeb\x97\x24\xa5\xd1\x33\xdb\xb1\x8e\xcd\x76\x5f\xbf\x66\x2e\x2d\x49\x7a\xc7\x9d\xae\x7e\x4d\x22\x29\x8f\x9e\x0c\x2a\xa6\x9f\x0c\xf4\x83\xe6\x59\x67\x30\xb3\xab\xdd\x19\x1c\x1d\x97\x9c\x79\x66\x9e\x9a\x6d\xcb\xae\x98\x32\xb7\x7d\xf9\x34\xd3\xb2\x92\x7d\x4e\x3b\xd6\x89\x76\xf9\x6c\xc2\xf2\x00\xcd\x66\xf5\xfa\x66\xdf\x2e\x29\x4b\xcc\x5d\xaa\x28\x37\x43\xaf\xe4\xd0\xb6\x5a\x67\xe6\xa9\xdb\x35\x9d\x9e\xed\xda\x4e\x2e\xef\x7d\xc7\x6c\xf7\x4c\xab\xdf\xea\xb4\xdd\xaf\xcc\xd6\xa9\x7d\xa8\x73\xb2\xb4\x6d\xdf\x31\x0f\x6d\xb7\x6b\xbe\x70\x9d\x41\x5b\x2f\xd0\x54\xdd\x42\x68\x3f\xef\xb6\x1c\x7b\xe9\x71\x3c\x5b\x5f\x89\x89\x84\xbe\x1e\x31\xe6\x8b\x43\x90\x98\x04\x5b\x51\xc3\xb0\xf0\xb0\x8c\xd6\x3d\xa5\x38\xf8\xb1\x97\xde\x50\x76\xd4\xec\x66\xbc\x7c\x6b\xc2\xcb\x62\xe9\x26\x37\x57\x97\xb2\x3a\xbe\xa9\x1b\x91\x0b\xd7\x11\xe1\xeb\xb7\xf2\xcc\x72\xca\x24\x19\x4d\xdc\x98\x07\x75\x56\xe7\x2f\x7f\x1b\x2f\x0f\x81\x7b\x17\xb8\x5e\x4e\xb1\x94\xd8\xbb\xa8\x67\x74\xa1\x12\xd1\x06\x57\xb1\x6c\x2d\xd7\x28\x97\x1d\x70\xb5\xf2\x0f\x2c\x6b\x7b\xfe\xc1\xe8\xd4\x7f\x7b\x58\xe7\xb1\xa7\x3f\xad\x73\xef\x38\x6e\x4c\xfc\x15\xf1\xca\x59\x36\x62\x3c\xc4\x49\x83\x23\x54\x7e\xf9\x85\xde\x67\x16\xb9\x24\xaa\x93\x85\x04\xf1\xd7\x3b\x26\x5f\xa5\x77\x9d\x07\x4a\x3d\xa7\xad\x6e\x65\x3a\x56\x79\xce\x99\x64\xc3\x78\x64\xd2\xad\x7a\xaf\x9a\x5f\xb7\xf5\x5d\xe1\x20\x5e\x95\xbf\x8a\xa2\x1a\x4e\x24\xac\xf2\x30\xf7\xd6\xb5\x0d\x8d\x2e\x28\x58\xdb\xbf\x65\xe5\x49\xa8\x84\x71\xe1\xd7\xb9\x3c\x68\x3e\xdf\xaf\xea\xf8\x42\xe0\x71\xad\x93\x6a\xda\xb6\x4b\xcf\xaa\x3b\x6b\x30\x99\x92\xd2\xb7\x94\x7c\x9e\x1a\xea\xef\xa6\xf1\x6f\x00\x00\x00\xff\xff\xce\xa0\x8b\x53\xf6\x1f\x00\x00")

func proto_micro_mall_pay_proto_pay_business_pay_business_swagger_json() ([]byte, error) {
	return bindata_read(
		_proto_micro_mall_pay_proto_pay_business_pay_business_swagger_json,
		"proto/micro_mall_pay_proto/pay_business/pay_business.swagger.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"proto/micro_mall_pay_proto/pay_business/pay_business.swagger.json": proto_micro_mall_pay_proto_pay_business_pay_business_swagger_json,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"proto": &_bintree_t{nil, map[string]*_bintree_t{
		"micro_mall_pay_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"pay_business": &_bintree_t{nil, map[string]*_bintree_t{
				"pay_business.swagger.json": &_bintree_t{proto_micro_mall_pay_proto_pay_business_pay_business_swagger_json, map[string]*_bintree_t{}},
			}},
		}},
	}},
}}
