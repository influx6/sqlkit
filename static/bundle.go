package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "sql-api-json.tml",
        
          "sql-api-migration.tml",
        
          "sql-api-readme.tml",
        
          "sql-api-test.tml",
        
          "sql-api.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "sql-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x41\x6b\xc2\x40\x10\x85\xcf\x09\xe4\x3f\xbc\x7a\x28\x06\x24\xde\x5b\x3c\x94\x42\x0f\x3d\x68\x41\x3c\x95\x82\xeb\x3a\xd1\xd8\x64\x57\x26\xa3\x52\x96\xfd\xef\x65\x37\x4a\x3d\xb4\xa5\xa1\x97\x84\x9d\x19\xde\xf7\xf1\x8e\x8a\xe1\x5c\x6d\x4f\xc4\x28\xe6\xc2\x07\x2d\xc5\x6c\xb5\x23\x2d\xc5\x54\x35\x14\x3f\xde\x3f\xcf\x67\x53\x4c\xb0\x74\x0e\x8d\xda\xc7\xd7\xf9\x18\x83\x55\x6b\xcd\x00\x83\x5d\xfc\x79\xbf\xcc\xd2\x2c\xfd\x5b\xea\x23\x93\x12\xfa\x2e\xbb\xdb\x3c\x68\xa9\xac\xf9\x0f\x61\xb1\x5f\xff\x40\xe8\x36\xbf\x11\xc6\x63\xd4\x56\xad\xc3\xf9\x93\x65\x30\xc9\x81\x4d\x0b\x05\x43\x27\x54\xa6\x15\x65\x34\xc1\x96\x50\x70\xee\xa2\xf0\xa2\xf4\xbb\xda\x90\xf7\xc5\xd7\xec\x4a\xcb\x7b\x94\x6c\x1b\xc8\x96\xb0\x67\x7b\xac\xd6\x84\x00\x85\xb6\x46\xc8\x48\x91\xa5\xe5\xc1\xe8\x6b\xf0\xf0\xbc\x42\x2b\x5c\x99\x4d\x8e\x61\x0f\xda\x08\xc4\x6c\x39\x87\xcb\xd2\x24\x54\x46\x35\x35\x7d\x74\x43\x11\x49\x55\x86\x18\xdc\x4d\xa2\x6b\xb1\x30\x8d\xe2\x76\xab\xea\xe1\xeb\xdb\xea\x43\xe8\x62\x98\x8f\x70\x1b\xf2\xf3\xfb\x78\x7e\x33\x81\xa9\xea\x48\x4e\xba\xf2\xfa\x80\x5d\xe7\x9e\xa5\x49\xe7\x70\x4e\x08\xf9\xa3\x90\x9b\xa5\x71\xfe\x19\x00\x00\xff\xff\x20\x59\xff\x66\xbf\x02\x00\x00"),
          path: "sql-api-json.tml",
          root: "sql-api-json.tml",
        },
      
        "sql-api-migration.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcf\x4b\xfb\x40\x10\xc5\xef\xfd\x2b\x1e\x3d\x7d\xbf\x50\xd3\x7b\xc0\x43\xd5\x0a\x42\xb5\xa0\xf1\x24\x22\xd3\xec\xb4\x59\xe9\xfe\x70\x77\x6a\x94\x90\xff\x5d\x76\x25\xc9\x41\xcd\x69\xb2\xf3\x3e\x6f\xde\x7b\xa7\x80\x7f\x33\x00\x58\x2e\xd1\x75\xc5\x83\x84\x53\x2d\xc5\x76\xf7\xca\xb5\x14\x77\x64\xb8\xef\x2b\xda\x1d\x19\x8a\xf7\xda\x72\x84\x34\x0c\xc9\x2f\x46\x1f\x02\x89\x76\x16\x86\x85\xce\x14\x09\xa1\x6d\x74\xdd\x60\x30\xd4\x11\xa7\xc8\x0a\xe2\x70\x60\xcb\x81\x84\x33\x4f\xde\x07\xe7\x83\x4e\xff\x97\xf7\xeb\x55\xb5\x46\xb5\xba\xd8\xac\x11\xdf\x8e\x88\x42\xc2\x86\xad\x60\xef\xc2\xc0\x69\x7b\x18\x5d\x29\x46\x57\x27\x56\x65\x7d\x0e\x53\x0c\xcb\x6a\x7b\xb5\x2d\xf1\xe8\x55\xf2\x6e\xb5\x34\xd8\x6b\x3e\x2a\x58\x32\x0c\xf9\xf4\x1c\x41\x56\x41\x5b\xc5\x1f\xdf\xd0\x1f\xad\x6f\xc7\x76\xe7\x18\xe7\x2e\x13\xe9\xab\xb4\xe1\x28\x64\x3c\xab\x12\x12\x4e\xbc\xc0\xb4\x4b\x89\x92\x4b\x89\x79\xd7\x89\xdb\xb8\x96\x03\x7e\xbd\xf2\x92\xd3\xcf\x17\x23\x7b\x93\x82\x71\x2c\xf1\xf4\x9c\xc7\xe9\x74\x3f\x89\xae\x53\xa5\xac\xc9\xd3\x0f\x4d\x3f\xfb\x3f\xfb\x0a\x00\x00\xff\xff\xde\xf6\xc3\xf5\xda\x01\x00\x00"),
          path: "sql-api-migration.tml",
          root: "sql-api-migration.tml",
        },
      
        "sql-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xd1\x6a\xdb\x30\x14\x86\xef\x03\x79\x87\x1f\x72\x93\x8c\x4c\x0f\x30\xd8\x45\x12\x6f\xa3\x50\xb6\x74\x59\xaf\xc6\xc0\xaa\x72\x62\x6b\x91\x25\x4f\x3a\xde\x5c\x8a\xdf\x7d\xc8\x76\x97\xb8\xb8\xb4\x65\xf5\x8d\xa4\x83\x75\xbe\x4f\xff\xb9\xbb\x13\x3b\xf6\x95\x62\xf1\xe5\xe6\x27\x29\x16\x9f\x65\x41\x4d\x83\xdd\xd5\x65\xb2\xc6\x6a\x7b\x31\x9d\xbc\x7f\xfa\x9b\x4e\xa6\x93\x27\x3b\x41\x07\x48\xc8\x8a\xdd\xdb\x8c\x2c\x79\xc9\xb4\xc7\xe6\xeb\x75\x02\x5d\x94\x86\x0a\xb2\x2c\x59\x3b\x8b\x83\xf3\xe0\x9c\x90\x8e\xb6\xec\xfb\xa6\xd0\x16\xa5\x54\x47\x99\x75\x7f\x6e\x8f\x99\xd8\x4a\xce\x9b\x26\x15\xd1\xe7\x5b\x4e\x38\x38\x63\xdc\x1f\x6d\x33\x14\xc4\xb9\xdb\x83\x6a\x1d\x38\xb4\x04\x55\x05\x76\x05\x5c\x19\x4d\xb4\xb3\xe1\x5d\xbc\x35\x9b\xe1\x43\x4d\x2a\x6e\xd3\x34\xcd\xdc\x74\x12\x8f\x73\xc5\x35\x94\xb3\x4c\x35\x8b\x4d\xb7\x2e\x71\xa8\x71\xa8\xac\x9a\xbf\x09\xbf\x8c\xd8\x5d\x5d\x2e\x11\x37\xc9\x7a\x01\xf2\xde\xf9\x7e\x69\x1b\x3d\x26\x14\xee\x8d\xb4\x6d\x9f\x7c\x0a\x26\x06\x26\x03\x4a\xf2\x2c\xb5\x8d\x37\xd8\xb5\x69\xdd\x6b\x6e\x3c\x49\xa6\x33\xd1\xae\x30\xae\x4a\x86\x0a\x9c\xe2\xdc\x76\xb9\x35\x8d\x78\x64\x6a\x0f\xdd\x67\x33\x7c\x22\x5e\xdf\x7e\xd4\x64\xf6\x67\xcc\x53\x71\x9c\x7b\xa4\x5b\x04\xf6\xda\x66\x4b\xfc\x96\xa6\xa2\xfe\xb4\x00\xe6\x2f\xd0\x59\xa2\x8f\xf4\x81\xd0\xd0\x64\x5c\xa1\xac\x6e\x8c\x56\x17\xc9\x3f\xf2\xcb\xc0\xa3\x5c\xac\x8c\x19\xb2\x57\xc6\x8c\xe1\x17\x98\x7f\xff\xf1\x9f\xbc\xeb\x72\x3f\x1c\x74\x57\x78\xd6\x6b\x5f\x65\xf2\x09\x19\x1a\x08\x74\x85\x67\xc6\x7d\xde\xee\x6f\x00\x00\x00\xff\xff\x66\x01\x04\xbe\x6e\x04\x00\x00"),
          path: "sql-api-readme.tml",
          root: "sql-api-readme.tml",
        },
      
        "sql-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x4f\xe3\xc6\x17\x7d\x8f\x94\xff\xe1\xfe\xac\x5f\x25\x7b\xeb\x0e\xcb\x6b\xaa\x3c\x10\xdc\x45\x5b\x55\x49\x8a\x43\xf7\x11\x0d\xf6\x75\x98\xdd\xb1\x07\x66\xae\x13\x90\xe5\xff\xbd\x1a\x8f\xf3\xc1\x47\xd8\x24\x2c\x95\x4a\x83\x04\x02\xfb\x7e\x9e\x39\xe7\x60\xa7\xaa\xe0\xff\x37\x99\x40\x99\x42\xaf\x0f\xcd\x2f\x83\xfb\x21\xcf\x11\x58\x4c\xba\x4c\x08\xbc\x71\x79\x25\x45\xf2\x39\xf2\xa0\xae\xbb\x9d\x26\x81\xf8\xd4\x86\x4f\x91\x26\x7c\xba\x2c\xe0\x99\x5b\xe9\x81\xf7\xd5\xa8\xc2\xc5\x76\x3b\x33\xae\xc1\xef\x76\x00\x00\x70\x86\x05\x19\xe8\x43\x8e\xa4\x45\x62\xd8\x10\xe7\x7e\x52\x1a\x52\x39\x8b\x89\x27\xdf\x22\x61\x6e\x24\xbf\xf7\x95\x61\x31\xa5\xaa\xa4\x20\xb0\x25\x6c\x6e\xa2\x8a\x4c\x4c\xa1\x0f\xe6\x56\xb2\xd3\xe6\x8f\xca\xdd\xb1\x5f\xd1\xc0\x0e\xdc\x03\x65\xd8\x19\x12\x16\x33\xdf\xab\xaa\x76\x7c\x36\xe6\xc9\x37\x3e\xc5\xba\xbe\x8c\xff\xfc\xe3\x32\x1a\x78\x41\xb8\xca\xbc\x30\xa8\xb7\xca\xbb\x88\x7f\x3b\x7f\x90\x19\x0d\x3e\x8f\xb7\xca\x3c\x89\xa2\xc7\x99\x63\xa5\x69\xab\xdc\xf1\xe8\x7c\xf2\x28\x37\xd2\x62\xb6\xe5\xcc\x2e\xf4\xc9\xbe\x63\x6e\xcc\x5c\xe9\x74\xbb\x09\x4e\xe2\xf8\xcb\xe8\x3c\x5a\x56\xa9\x17\x47\x42\x68\xe8\x54\x49\xe8\x83\x57\x55\x52\xcd\x51\x2f\x18\xc3\x46\x57\x5f\x31\x21\x66\x0f\xa5\xf9\x51\xd7\x97\x36\xda\xeb\x76\x9a\x03\x3d\x3a\x82\x09\x1a\x3a\x43\x5a\x75\x5d\x4b\xa9\x6b\x98\x71\x29\x52\x4e\x68\x80\xae\x11\xb4\xe5\x0b\xce\xb8\x04\x95\x01\x87\x0d\x49\x4d\x5d\x8d\x89\xd2\x29\x64\x5a\xe5\xc0\x2d\x59\xd2\x2b\xd6\xed\x64\x65\x91\x7c\xa7\xa5\x4f\xf0\xc1\xce\x28\x8a\x29\x9b\x04\x2d\xb7\xf8\x8d\xb0\x2c\x77\x65\x2c\x5b\xdb\x9d\xc3\x96\xcb\x61\x43\xc7\x21\xce\xa3\x81\xef\x18\xba\xb8\xb3\xc6\x5c\xba\xb3\x35\x12\x55\x10\xde\x11\xfb\x22\xe8\x7a\x22\x72\x54\x25\xf9\x8b\x6b\x43\x9c\xff\xc5\x65\x89\x03\x3e\xf5\x83\x10\x8e\x3f\xc2\x07\x20\x91\x23\x8b\x31\x51\x45\xba\x2c\x85\x12\xf3\x10\x50\x6b\x5b\x50\x2a\x9e\xfe\x1e\x8f\x86\x9f\x94\xf6\xbf\x8f\xff\xa9\x46\x4e\x68\xe3\x03\x57\x4b\x64\x4d\xa1\xff\xf5\xa1\x10\x12\xd6\xb4\x64\x57\x34\xec\x13\x17\x12\x53\xdf\x8b\xcb\x24\x41\x63\xb2\x52\xca\xfb\xa6\x25\xa6\x60\xab\x40\xa6\xf4\xa6\x83\x68\x0f\xa1\x07\x3f\xfd\x7c\xcb\xbc\x66\xe0\x60\xc9\x9d\x55\x0b\x4b\xc3\x57\xb6\xf0\x96\xd0\xa4\x98\xa1\xb6\xe7\xc5\x22\x94\x48\xe8\x27\x74\x17\x36\x80\xb1\x85\x7b\x2d\x63\xdb\xd5\x7b\xfd\x26\xde\x21\xb3\x8a\x0f\x7e\xdd\x1d\x18\x9e\xda\xa1\x17\xdc\x7b\x61\x6c\x51\x90\x82\xf4\x6a\x2f\x68\x76\x6d\xc2\x56\xe8\x5c\x3a\xd6\xb8\x85\xcf\x90\x9e\x47\x67\x5f\x5a\xb4\xf2\xc4\x14\x0c\x29\xbd\xdd\x8c\x8d\x40\xf7\x04\xe2\x15\xfd\x1a\x4c\xea\x87\x1e\x74\x22\xe5\x3e\x36\x24\xe5\x2b\x8d\x68\x73\xdf\x83\x17\x1d\xbc\xe8\xdd\x7a\x91\x4b\x32\xe1\xc2\x94\x7a\x4b\x57\x3a\x91\xd2\xad\xee\x71\x93\x78\xa1\x57\x55\xcd\x73\x26\x6b\xf8\x58\xd7\x5e\x08\xbf\x1c\xdb\xef\x1f\x62\x55\x56\xbf\xed\x28\xff\x80\x51\xed\xd8\x6d\x0d\x2e\x91\x81\xc4\xc2\x6f\x93\x03\xe8\xf7\xe1\xe3\xee\xcb\x92\x44\x6e\x08\x8e\x77\xb5\xca\x9d\xf7\xdc\xb7\xd1\xba\x27\x6f\x88\x77\xe2\x78\xe4\xc9\x89\xbd\x28\x54\xb1\xf5\x93\xe1\x5c\xd0\xf5\x73\x86\xfc\x62\xd3\x83\x21\x1f\x0c\xf9\x1d\x1a\xf2\x16\xa2\xbb\xb8\x49\x9f\x8a\xae\x74\x17\xdf\x48\x72\xae\xe5\x41\x72\x07\xc9\xbd\x43\xc9\xb9\xaa\x47\x47\x93\x51\x34\xea\x41\xab\x2e\xa3\x72\xa4\x6b\xcb\xf4\xe7\x21\x69\x15\xf1\x14\xc2\xfd\x11\x72\x12\x7e\x73\x8c\x76\x6f\x73\x78\x6b\xdd\xff\x09\xc9\x49\xed\xc9\x5b\x6b\xae\x7e\xc4\x47\x67\x2f\xf6\x3c\xb8\xf5\xbf\xda\xad\xff\x7b\x0e\xfc\x70\xe3\xcd\xff\xa3\xf6\x58\xbd\x11\xdc\xdb\x2f\xbf\x7b\x9b\x85\x8d\xfc\x1d\x00\x00\xff\xff\x08\x8f\xe4\xc4\xf2\x19\x00\x00"),
          path: "sql-api-test.tml",
          root: "sql-api-test.tml",
        },
      
        "sql-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xdb\x38\x12\x7f\x4e\x80\xfc\x0f\xb3\xc2\x61\x21\x2d\x74\x0a\xee\x35\x41\x1f\x36\x4d\x7b\x28\x70\x9b\xf6\xda\xee\xdd\xc3\x62\xb1\xa0\xc5\x91\xcd\x56\x1f\x2e\x49\xc7\x36\x0c\xff\xef\x87\x21\xa9\x0f\x2b\x96\x2d\xdb\x4a\xdb\xa4\xd7\x87\x3a\xb4\xc8\xe1\xcc\x6f\x3e\x39\xa2\x2f\x2f\x21\x63\xd3\xd7\x02\x53\xae\x80\x63\x22\x72\x54\xc0\x40\x2f\xa7\x08\x49\x21\x81\xd1\x63\xd0\x13\xa6\x01\x17\xd3\x42\x99\xa7\x76\xba\x1f\x40\x86\x7a\x52\xf0\xe8\xe2\xdc\xcc\xaf\x09\x65\x6c\xfa\x87\xd2\x52\xe4\xe3\x3f\x45\xae\x51\x26\x2c\xc6\xd5\xfa\xe2\xfc\xe2\xfc\xf2\xd2\xad\x06\x89\x7a\x26\x73\x05\x7a\x62\x56\x82\xd0\x0a\xd3\x04\x58\xce\x61\x2a\x8b\x7b\xc1\xcd\x56\x76\x07\xd0\x05\x64\x4c\xc7\x13\x33\x5b\x7d\x49\xa3\x8f\x6c\x94\xa2\xa1\x04\xd5\x0e\xd1\xc5\x79\x32\xcb\x63\xf0\xb3\x9a\x95\xa0\xc1\xec\x56\xa6\x60\x75\x71\x0e\x00\x8e\x1d\xc8\x2e\xce\x4b\x3e\x57\xab\xe8\x83\x96\xb3\x58\x47\x6f\x47\x9f\x30\xd6\xd1\x1d\xcb\x70\xbd\x6e\x63\x95\xd7\x0c\xc0\x7c\x22\xe2\x49\x05\x54\xcd\xbb\x23\x6e\xc1\x2c\x12\x60\x69\x6a\xb6\x60\x5a\x4b\x31\x9a\x69\xa2\xa3\x54\x11\x0b\xa6\x91\xc3\x5c\x68\x2b\xa8\xdd\x83\x83\x32\x6c\xcc\x24\x02\xa3\x8d\x63\xc1\x91\xc3\x68\x69\xc1\x28\x9f\x95\x5a\xd8\xcd\x76\x83\xd9\xd5\xc5\xf9\xd9\x1e\x70\xf6\x82\xf1\xb2\xc8\xd5\x2c\x43\xb9\x0b\x0e\x16\xc7\x38\xd5\xaa\x96\x9e\x33\xcd\xdc\xb3\xb9\x48\x53\x18\x21\xc4\x96\x0e\x37\x7b\x89\x5c\x17\x46\xb6\xb1\xb8\x17\xf9\x18\x44\x36\x4d\x31\xc3\x5c\xd3\x60\x08\x2c\x2a\xae\x37\xc1\x70\x5f\xfb\xdb\xb1\x08\x00\xa5\x2c\xe4\x5e\x48\x6e\x6f\x1a\x7e\x54\x73\x6b\xe5\x75\x96\x0d\xb7\x37\xf0\xf2\xfd\xef\xb7\x50\x4c\x51\x32\x2d\x8a\x5c\x19\x92\x33\x65\x44\xfc\x92\x92\x70\x24\xd1\x2c\xe7\x28\x53\x91\x23\xf0\xd1\x1e\xa1\x6e\x6f\xdc\x6e\xc6\xa0\xe3\x22\x05\x2b\x03\x8d\xd4\xc2\x38\xcd\xed\x0d\x0d\xf8\x02\x7e\xa1\xd1\x87\x7f\xff\x8b\x86\x19\x6a\x29\x62\x55\x7e\x46\xbf\xd9\x4f\x7a\xa4\xc9\xc9\x68\x67\xe3\x6d\x6f\x38\xa9\x40\x2f\x6b\x04\xee\x70\x5e\xb9\x31\x83\x1c\xe7\x20\x72\xa5\x59\x1e\x23\xa9\xb9\x93\xcf\xd2\x4f\xef\x70\xee\xdb\x2d\x2c\xa7\x21\x64\x6d\x2e\xc2\x9a\xf5\x10\x74\x06\x51\x14\x99\x15\xca\xb2\xf4\x9b\x18\x5b\xf8\x02\xf8\xa5\x73\x3b\xe7\xe0\x7c\x01\x57\x2f\x0c\x2d\xda\x37\x23\xca\x44\x32\x8a\xa2\x80\xc4\x69\xc4\x80\x9f\xf7\x91\xa2\x7f\x6a\x71\x45\x14\xea\x2f\xf8\xe2\x0a\x78\xf3\x8b\xb8\x48\xaf\x2c\x84\x8d\x2f\x9d\x7c\x57\x90\x35\xbe\x34\x93\xae\x2a\xa0\x69\xb7\x15\xfd\xe7\x96\xaf\xdd\xd4\x86\x3b\xde\x62\x8a\x1a\x29\x7e\x60\x46\xbe\x65\x62\x4c\x56\xdc\xa3\xb1\x1a\x89\x71\x21\x39\x24\xb2\xc8\x6c\x1c\x19\x39\xcb\xa2\x81\x33\x41\x0e\xd3\xd9\x28\x15\xf1\x9b\xdb\xc8\x50\x7c\x6f\xd6\xa8\x6a\xa2\x50\x64\xa4\xd9\x4c\x69\x98\xb0\x7b\x04\xe6\xe6\x83\xe0\x70\xcf\xd2\x19\x86\x14\xe7\x24\x2a\x85\x1c\x50\xe8\x09\x4a\x72\x43\x06\x23\x55\xe4\x50\x48\xf8\x44\x9f\x9a\x8d\x0d\x75\xfa\xd3\x3a\x34\xe6\x0d\xc3\x78\xc7\xe2\xcf\x6c\x8c\xeb\x75\xd4\x01\xb9\x33\xe9\x3a\xb2\xf3\xd1\x0e\x4d\x07\x0e\x17\x3f\xd6\x0b\x0a\x29\x1a\x17\x3a\x7a\x69\x3f\xc3\x4a\x5e\x67\x6d\xce\x9f\xcb\xf8\x9f\x91\x75\x94\xd6\x77\x87\xf3\x8f\x92\xc5\xe8\x7b\xdd\x46\x6c\xb7\xf2\x02\x67\x5d\x98\xa0\x84\x8c\x8f\xa2\x92\xc6\xab\x4c\x68\xbf\x1c\xbc\xc9\x93\xa2\x0f\xb1\xb0\x5c\xf0\x5f\xa1\x27\xbe\x57\xb2\xec\xd5\xdc\x6f\x4e\xb1\x5c\x66\xd1\xab\x9c\xfb\x41\x50\x19\xb2\x48\x2a\xf1\xdf\xa8\x57\x8b\xa9\x90\xc8\x09\x94\x00\x1a\x16\x8c\x52\x92\xcc\x49\xa6\xa3\x57\x84\x44\xe2\x7b\x0e\x2b\x98\x30\x45\xda\xa5\x65\xa5\x80\x06\xa3\x2e\xf1\xca\xf5\xaf\x99\x48\xd1\x64\x3c\x6e\x0d\xd4\x5a\x62\x4b\x2e\x97\x74\xf6\x4b\x57\x4e\x34\x5e\xe0\x85\x66\xff\xb8\x48\xb7\x4f\x32\xda\xf4\x42\x12\xcb\xf2\x63\x01\x29\x79\x77\xbe\x8d\x52\x56\xce\x54\x61\xe5\x90\x20\xf2\x7c\xe1\x74\x41\xa6\x66\x63\x4d\x08\x8e\xd3\xbf\x04\x6f\xb2\x7a\x6d\xd6\xfd\xf4\x02\x72\x91\x36\x71\x1d\x18\xa5\x5e\xc2\xf7\x86\xf2\x58\x94\xf6\x58\xb6\xc5\x8c\x3f\x92\x24\xed\x00\x9d\x8b\xb4\x0e\x85\x2f\x25\xb2\x56\x28\x64\x9c\x37\xe3\x60\x55\x4f\x6c\x8f\x83\xcd\x9c\xa5\x27\xd8\xca\xed\x7b\x43\xd4\x37\x0c\x9f\x27\x85\x4a\x8b\xdb\xf6\x50\x89\x29\x66\x87\x60\x70\x6a\x2c\xb5\xbc\x0c\x14\x4b\x4b\x62\xdd\xb1\x94\xc4\x8b\xde\x3d\xad\x80\x1a\x5b\x33\xef\x1b\x50\x77\x88\x38\x44\x54\x3d\x3b\x3b\x3b\xdb\x15\x51\x13\xb3\x3c\x84\xe2\x33\x21\xd2\xa8\xa2\x7d\x62\x2c\x88\xfc\x9d\xa7\x94\xe0\x9a\x16\x36\x90\x7d\x10\xa3\x3f\xb0\xfb\x8d\x08\x5d\x9d\xf7\x7c\xbb\x73\x54\x1e\x6c\x82\xce\x30\x7d\x1c\xfe\x5d\xce\x77\x7a\xe0\x23\x60\x9c\xe2\x0e\x52\xc5\xa6\x44\x6d\xa5\x6c\x28\xa6\x47\x20\x7f\xd9\xc3\xcc\x0e\x16\x67\x83\xcb\x66\x0c\x6f\x9b\xcd\x90\x3e\xf6\x0c\xbc\x6c\x6f\x56\x36\x78\xe5\xda\x38\x46\xe5\x01\xab\x15\x0d\x8a\x99\x76\x47\x7e\x17\x10\x7f\x8d\xe9\xa0\x04\x56\x31\xe0\x51\x9e\xf1\xc0\xfb\x64\x3e\xd6\xeb\xa0\xb3\x1c\x6a\xbb\x9a\xdb\x73\x88\xf2\xe7\xd1\x7d\xea\xcb\x0c\xe5\xd2\xab\x79\xfe\x9a\x25\xd1\x60\x9e\xd4\x16\x22\xe8\xae\x85\xfe\x89\xfa\xd7\x34\xa5\x47\x52\xe0\x3d\x2a\x60\x66\x64\x4b\x93\xe6\x71\x90\xe5\xbc\x71\x72\x57\xa9\x68\x1f\xd9\xf7\x9f\xcc\xf4\x72\x8a\x4f\xb5\xf6\xb1\x38\x6d\xaf\x7d\x0a\xc9\x51\x56\x1d\x09\x33\x1a\x2d\xab\xf1\x94\x8d\x91\xf2\x59\x08\x12\xd5\xb4\xc8\x15\xbe\x43\xf9\xce\x7d\x19\x00\xf8\x7f\xfc\x79\x00\x88\xa1\x25\x65\x2c\x30\x38\xba\x7a\xb2\xd2\x0c\x54\x3d\x95\xc4\x9e\x46\x55\x54\x5a\xfa\x10\x09\x6b\x5f\x7a\xad\x1d\x2e\x84\xbf\xff\x23\xdc\x12\x11\x56\x2b\x82\xc4\x27\xb9\x5e\x93\x1d\x3a\x8c\xc1\x73\x9d\x45\x0f\x20\x80\xf5\xba\x26\x79\xcf\x24\x48\xa1\x31\x53\x70\x90\xdd\x34\xd9\x32\xcb\x43\xd0\x85\x66\x69\xd8\x8a\xdd\x56\x99\xce\x42\x9b\x41\xdc\x98\x75\x65\xdd\xd6\xac\x1f\x98\x74\x70\xfd\xa0\x02\x3b\xb9\x8c\xaa\x14\xd6\x8c\x4c\xdd\xcd\x42\xf7\x22\x82\x22\x17\x1f\x3d\xaa\x76\xdb\x1a\xae\xf1\xec\x28\xa4\x92\x42\xc2\x5f\xa1\x81\x9f\x10\x97\x2c\xa7\x20\x60\x74\xd9\x82\x87\xb4\x7c\xe8\x89\xaa\xcd\x5a\x9d\x96\x4d\xcd\x51\xb6\xaa\x69\xbf\x9d\x15\xee\xa1\x19\xd9\xd2\x2d\x4f\xce\xa6\x53\x4f\x82\x76\xa9\x67\x00\xcd\xd0\x1e\x9e\xc5\xf1\x04\xd5\xed\x74\xd0\xad\xfa\x33\x2b\xac\xbe\x5e\x00\x9b\x4e\x31\xe7\xbe\x74\xce\x64\xaa\xd6\x0e\xbd\xbb\x5d\xe4\xa6\xdf\x55\xe5\xec\x6a\x05\x98\x2a\x1c\xdc\xcd\xf7\xfb\xf7\xcd\xf2\x60\x0f\x0f\x81\xf2\xa5\x2f\x8b\xb9\x32\xef\x04\x16\xd1\xfb\x62\xae\x5a\xa7\xf9\xa6\xc1\xd3\xcc\xe8\x0e\x17\xda\x0f\xb6\x59\x5a\x25\xe8\x29\xa6\xbe\x69\xee\x66\x47\xbb\xec\x43\xcc\x72\xff\x67\xd9\xcb\xe6\xfb\xd8\x3d\x4a\xb9\xd5\xe2\xca\xaf\xcc\xb0\x83\x36\xfd\x23\xcb\xf4\xae\x6c\x8a\x0b\x77\x4c\xb3\x6e\x70\x05\x95\x76\x6c\x9b\xdf\x0f\x3a\x16\xad\xb7\x9b\x37\x74\x1c\xef\xaa\x65\x5b\x7d\xa2\xc3\xc2\x2d\x8a\xfb\xdc\xa3\x75\x54\xb3\xcc\x5d\x6f\x64\x9f\x81\xf2\xc2\x37\x08\x3c\xc3\xa7\x84\x1e\xa1\x21\xe7\x26\x32\x98\x65\x75\xd1\x7e\xb3\xb4\xaf\xb2\x1b\x85\xfb\xee\x97\x38\xa6\xcb\x01\x9f\x71\x69\x2a\x79\x5b\x54\x1b\x62\xcd\xb7\xea\x3f\x50\x35\xef\x00\xdc\x5e\xd1\x13\x4c\x65\xfd\x6e\x38\xae\xde\x00\x81\x7f\x50\xb1\x3e\x40\xa1\x3e\x5c\x95\xde\x6e\x70\x7e\xc6\xa5\x13\xef\x39\xd5\xee\x5d\x52\x3d\x4a\x69\x7f\x80\x31\xac\xd6\xdb\x0e\x00\xe4\xf8\x5b\x72\x74\x33\x2f\x37\x04\xba\xde\xe8\xb8\x9c\xd4\x49\xf9\x6e\xcb\xea\xa1\x90\x3d\xb6\x86\x3e\x3b\x13\x09\x95\xe9\xb6\x11\x6d\x6a\xe7\xae\xde\x73\x79\x2b\x24\xb8\x86\x9f\x6c\xfb\xb9\xd1\xe4\x4e\x9d\x62\x0b\x69\x5c\xdd\xf7\xde\xe6\xe9\x72\xef\xfd\x12\x96\xa6\xc5\xdc\x3a\xc8\xd9\xd9\x96\x37\x8e\xfd\x6b\xf9\xef\x3b\x95\x9e\x5e\xc3\x0f\x65\x27\x1b\xfa\x6a\xf7\xc6\xfa\xe7\xd7\xea\xae\x40\xb3\x51\xf6\x83\xe5\xd4\x9e\xb7\x28\x9e\x53\x0e\x3d\xea\xc2\xc5\x93\xca\xa7\xdf\xe4\xd2\xc5\xc9\x5e\x7d\x70\x63\xad\x47\x22\xee\xb8\xd4\xf1\x63\xb6\xbc\x8e\xd5\xd0\x83\x23\xd0\xa9\xbd\xae\x63\xfb\x5c\xcf\x3d\x3f\x3e\x86\xae\x1e\xe6\x4a\x68\x35\xb0\x86\xd3\xeb\x6a\x05\x7f\x9b\xda\x53\x2b\x05\x40\xfa\xe3\x66\x49\xd3\x6a\x5f\x2e\xdf\xa3\x7a\x1b\x9e\x6c\x16\x6a\x36\xa6\x65\x63\xd4\x1f\xd9\xb8\x22\xe4\xa9\x2f\x69\xf9\x0a\xb3\xcb\x8a\x6a\xdf\xdf\x6c\x8f\x95\x7d\xaf\xba\xed\xd5\xd1\xf5\xda\xe8\x44\x6d\x34\xa2\x4c\x93\xf0\xd4\xde\xeb\x29\x3d\xa8\x1e\xfd\xa7\x83\x7a\x4f\x1d\x7d\xa7\xae\x9e\x53\xbf\x76\x51\x08\xde\x6a\x65\x34\x18\xfd\x87\x2a\x9b\xf5\xba\xcf\x1d\xba\x3e\xd0\x3d\x9d\xee\xd1\x57\xf0\xd8\xcd\xbe\xd2\xe5\x25\xfc\x3e\xe5\x4c\x23\xcc\xd4\xff\x8b\xdd\xbd\xc5\xae\xc5\xaa\x57\xbd\xfb\xf5\xef\xc6\x59\xe6\x06\xaa\x79\x4b\x62\xcf\xe6\x9e\x71\x22\x72\xa1\x26\x61\xc9\xc6\x26\xc5\x53\x6f\xa1\xb6\x6a\xc3\x41\x6f\x8f\x50\x42\x35\x71\xaa\xba\xb8\xf3\x78\xd7\xe7\xcc\x3e\x2f\x76\x5c\x92\x73\xbc\x51\xde\xdf\xb5\xec\xd8\x9b\x45\xd0\xaa\xe5\x0f\xaf\x04\x0e\xae\x02\x1e\x54\x00\xce\xc7\x1b\x25\x00\x89\x77\x5c\x76\x3a\xc0\x3e\x67\x36\x0c\x1f\x75\xc5\xa9\xbc\xf8\x43\x8c\x9e\x60\xcb\x3b\x04\xfc\x9a\xf7\xa1\x5c\x42\x3a\xfd\xa2\xc6\xa1\xf2\x6c\xc0\xb8\xf3\x2a\xf9\xab\x05\xc6\xcd\x1f\x1f\x52\xf6\x30\xa6\xed\x7e\xc8\x96\xa6\xc5\xdc\x66\x46\x5c\x60\x3c\x33\x8f\x8a\x04\x18\xc4\x33\xa5\x8b\xac\x9e\xcf\xc6\x4c\xe4\x4a\x9b\xa9\xb6\xea\xea\x9b\x8a\x88\x85\xed\x89\x28\x59\xd8\xa2\xb5\xfc\xed\x56\xe8\x7e\x17\x15\x94\x1d\x95\xd3\xd2\x0c\x6d\x3c\x50\x92\xb1\xa4\x9e\x46\xfe\xb0\x8a\xc4\xfa\x87\x78\x8f\x50\x13\xee\xf5\x9c\x3a\x58\x25\x0b\xdf\xc6\x2b\xbb\x93\x5a\x0c\x11\x83\xbe\x0b\x19\xf7\x58\xd3\xdb\x92\xb7\x92\xdb\xbe\x21\xa2\xcb\xa1\xff\x17\x00\x00\xff\xff\x1c\x60\xfb\xa0\xd9\x3c\x00\x00"),
          path: "sql-api.tml",
          root: "sql-api.tml",
        },
      
    
  }
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
  return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
  reader, size, err := FindFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

  gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
  }

  return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
  body, err := ReadFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error){
  body, err := ReadFileByte(path, doGzip)
  return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
  body, err := ReadFileByte(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
