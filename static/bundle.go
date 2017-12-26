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
        
          "dockerfile.tml",
        
          "makefile.tml",
        
          "sql-api-backend.tml",
        
          "sql-api-json.tml",
        
          "sql-api-migration.tml",
        
          "sql-api-readme.tml",
        
          "sql-api-test.tml",
        
          "sql-api.tml",
        
          "sql-simple-readme.tml",
        
          "sql-simple.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "dockerfile.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x51\x6f\xd3\x30\x14\x85\xdf\xf3\x2b\x8e\x54\x34\x81\xb4\x38\x19\x85\x3e\x01\x22\x23\x01\x55\xd0\x64\x24\x61\x88\xa7\xc8\x49\x6e\x1c\xab\x89\xdd\xda\x4e\xb5\x6a\xda\x7f\x47\xed\x10\x0c\x54\xfa\x7a\xbf\xcf\xf7\xf8\xc8\xfe\x98\x67\x2b\x48\xd5\x0d\xd3\xdd\x22\xb0\xdb\x89\x1b\x1a\xfc\x90\x85\xec\xca\x5b\x45\xcb\xb4\x8c\x96\x69\x92\xe3\x53\xf6\x79\x59\x3e\x17\xd2\xd5\x53\xcd\x1a\x3d\x06\x42\xaf\xa5\x7b\x81\x37\xce\x48\xa5\xef\xba\xf7\x62\xe4\x72\x38\x90\x77\x9e\x37\x43\xd9\x4b\x0b\x69\xc1\xe1\xc8\x3a\xc8\x91\x0b\xba\x44\xab\x95\x43\xab\xe1\x0e\xb4\xd3\x06\x5c\xed\xb1\x31\xba\x9d\x1a\x27\xb5\xf2\x66\xb0\x7b\xeb\x68\xc4\x66\x20\x6e\xe9\x12\x7c\x44\x4d\x42\x48\x25\x7e\x8d\x18\x96\xca\x3a\xe2\x2d\x06\xcd\x5b\x58\x6a\x0c\x39\x0b\x6f\x06\xd7\x1b\x3d\x89\x1e\xae\x27\xf8\x3e\xa9\x9d\xdf\xc9\x81\xd0\x0d\x5c\x1c\xa3\x5a\xdd\xac\xc9\xc0\x4c\x0a\xcc\x4b\xd2\x5b\xac\x7e\x14\x5f\xbf\x54\x79\x96\x95\xd5\x4d\x54\x14\xdf\xb3\x3c\xc6\xd5\xcb\xf9\xab\xd7\x8b\x27\x38\x8e\xca\xe8\x3a\x2a\x92\x63\x8d\xaa\xad\x9f\xa0\x6f\x45\x92\x3f\x8e\x27\x4b\xe6\x1f\xf0\x67\xe5\x6f\xa3\xda\x70\x6b\x8f\xda\xfd\x3d\x4b\xd2\xdb\x94\x8f\xf4\xf0\x50\x1d\x4e\x94\x49\x51\x56\xf1\xf5\x5f\x21\xa7\xa5\x13\xa1\xa7\xc5\x28\x8e\x73\x1c\x5e\x31\x64\xe1\x19\xed\x26\xcb\x4b\xcc\xe7\xe1\xe2\xdc\xc5\x8c\xdc\x91\xc1\xb8\xb7\xdb\x73\x9b\xfe\x53\xd8\xfb\xb0\x8a\x11\xd4\x52\x05\xb6\x87\xdf\xe0\x59\xa7\xcd\xfa\xad\x33\x13\xa1\xd6\xda\xd9\xed\x80\x8b\x0b\x08\xfd\xf8\x51\xfc\x1d\x58\xc0\x18\xfb\x19\x00\x00\xff\xff\xd8\x1a\x9c\x3b\x95\x02\x00\x00"),
          path: "dockerfile.tml",
          root: "dockerfile.tml",
        },
      
        "makefile.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x4a\x2a\xcd\xcc\x49\xb1\x52\x48\xc9\x4f\xce\x4e\x2d\x52\x00\xf3\x14\x74\x4b\x14\xaa\xab\x4b\xf2\x7d\xf2\xcb\x53\x8b\x14\xf4\x02\x12\x93\xb3\x13\xd3\x53\xfd\x12\x73\x53\x6b\x6b\x15\x74\xd3\x14\xf4\xf4\x4b\x52\x8b\x4b\xf4\x20\x5a\xd2\x32\x73\x52\xb9\xb8\x40\x02\x56\x10\xdd\x5c\x9c\x50\xb3\x8a\x4a\xf3\x14\x74\x75\x8b\x72\x71\x9a\x05\x08\x00\x00\xff\xff\xd1\x0a\xfe\x5d\x7c\x00\x00\x00"),
          path: "makefile.tml",
          root: "makefile.tml",
        },
      
        "sql-api-backend.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x41\x8b\xa3\x40\x10\x85\xef\xfe\x8a\x3a\x26\x20\xfa\x17\x76\x8d\x6c\xd8\xcb\x26\xb0\xec\x69\xd9\x43\xdb\xbe\x68\x6f\xda\xee\xa6\xbb\xcc\x44\x82\xff\x7d\x30\x86\x44\x06\x1d\x48\xe6\x24\x56\xd7\xab\xf7\xbe\xa2\xd2\x94\x2e\x97\xe4\x37\xfb\x56\x72\xb2\x2b\xfe\x43\x72\xf2\x4b\x34\xe8\xfb\x3c\xcb\x84\x3c\xc2\x94\x54\xe2\xa0\x0c\x02\x09\x2a\x6e\x95\xb7\x5a\xc9\x9a\x3c\x9c\x47\x80\xe1\x40\x5c\x83\x2a\x75\x52\xa6\x8a\xd2\x94\x1a\x70\x6d\xcb\x40\x38\x3b\x1b\x50\x52\xd1\x5d\x1b\xf2\x8c\x54\xe3\x34\x1a\x18\x16\xac\xac\xa1\x83\xf5\x13\x29\x71\xe7\xb0\x14\x27\x19\x06\x7f\xbb\xeb\xa3\xcf\x7a\x1f\xd1\x95\x61\xf8\x83\x90\xb8\x44\x44\x1b\xdb\x1a\x5e\x49\x3e\x93\xb4\x86\x71\xe6\x64\x33\x7e\xd7\xb4\x52\x86\x63\x82\xf7\xd6\xaf\x23\xa2\x1c\x1a\x8c\xb9\xd6\x98\x5c\x5b\x68\x25\x7f\xe6\x14\xd8\x2b\x53\xad\x47\xd5\x30\xde\x43\x2c\x89\xa0\xd1\x4c\xd2\xee\x85\x3c\x8a\x6a\xa0\x5a\x20\xb8\x4d\xa5\x88\x68\x8b\xd9\xcc\x33\x41\x68\xf5\x84\x43\x4c\x0f\xdc\x3f\xae\x5c\x4c\xfe\xc1\xe5\x65\x94\x91\xe4\xbb\xd6\x59\xb7\xf3\x25\xfc\xbc\x9b\x1d\x9e\xee\x56\xd7\xbf\xac\x9b\x00\xfe\xfd\xf7\x22\xe2\x16\x9c\x75\x3f\x14\x74\x39\x6f\x7c\x44\x77\xb7\x3d\x09\xdd\x62\x72\x3a\xfd\x17\x56\x3b\x32\x3f\x0f\x1b\x93\x13\xd5\x35\x44\x4c\x1e\xc1\x59\x13\xb0\x87\xdf\xdf\x8a\xaf\xec\x62\x7a\xe2\xfd\x7b\x00\x00\x00\xff\xff\x63\x4a\x7e\xe2\xf8\x03\x00\x00"),
          path: "sql-api-backend.tml",
          root: "sql-api-backend.tml",
        },
      
        "sql-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xcf\x4a\xc4\x30\x10\xc6\xcf\x0d\xe4\x1d\xc6\x3d\x48\x0b\x4b\xf6\xae\xf4\x05\x44\x76\xc5\xc5\x93\x08\x3b\xcd\x4e\xd7\xd6\x76\x22\x49\xea\x1f\x42\xde\x5d\x92\xae\x78\x11\xb4\x97\x04\x86\xcc\xef\xf7\xf1\x45\x8a\xcd\x06\x7a\x67\x18\xda\xee\xc3\x4f\x96\x1c\x28\xa5\xa4\x78\x43\x0b\xa5\x14\x45\x08\x6a\xef\xed\xa4\xbd\xda\x35\x3d\x69\xaf\xb6\x38\x52\x3e\x62\xbc\xd9\xef\xb6\x50\xc3\x21\x04\x18\xf1\xf5\x1e\xf9\x68\xc6\x3c\x3b\xaf\xc0\xaa\x71\x86\x57\xb0\xea\xf3\x15\xe3\x41\x8a\x4a\x8a\xac\xbc\x35\x78\xfc\x93\x6d\xc9\x4f\x96\x1d\x20\x30\xbd\x43\xc7\xce\x23\x6b\x02\xd3\x02\xc2\xcf\xf2\x1d\xea\x17\x3c\x51\x8c\xea\x57\x60\x8c\x4a\x8a\x76\x62\xfd\x2f\x67\xa9\x0d\x7b\x62\x0f\xce\xdb\x8e\x4f\x15\x94\x0b\x44\x6b\x20\x6b\x8d\xad\x20\x48\x51\xa4\x06\x69\xa0\x71\x49\xd2\x54\x4e\xd1\xb5\x09\x03\x57\x75\xfe\x17\xf5\xc0\x23\x5a\xf7\x8c\x43\xf9\xf8\xd4\x7c\x7a\xfa\x4e\x58\xad\xe1\x32\xf1\xab\xeb\xfc\xfc\xa2\x06\xee\x86\x6c\x2e\xe6\xde\x96\x88\xc3\x9c\x5d\x8a\x62\xce\x70\x26\x24\xfe\x3a\x71\xa5\xc8\xf3\xaf\x00\x00\x00\xff\xff\x8a\xed\x5b\xef\x2e\x02\x00\x00"),
          path: "sql-api-json.tml",
          root: "sql-api-json.tml",
        },
      
        "sql-api-migration.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcf\x4b\xfb\x40\x10\xc5\xef\xfd\x2b\x1e\x3d\x7d\xbf\x50\xd3\x7b\xc0\x43\xd5\x0a\x42\xb5\xa0\xf1\x24\x22\xd3\xec\xb4\x59\xe9\xfe\x70\x77\x6a\x94\x90\xff\x5d\x76\x25\xc9\x41\xcd\x69\xb2\xf3\x3e\x6f\xde\x7b\xa7\x80\x7f\x33\x00\x58\x2e\xd1\x75\xc5\x83\x84\x53\x2d\xc5\x76\xf7\xca\xb5\x14\x77\x64\xb8\xef\x2b\xda\x1d\x19\x8a\xf7\xda\x72\x84\x34\x0c\xc9\x2f\x46\x1f\x02\x89\x76\x16\x86\x85\xce\x14\x09\xa1\x6d\x74\xdd\x60\x30\xd4\x11\xa7\xc8\x0a\xe2\x70\x60\xcb\x81\x84\x33\x4f\xde\x07\xe7\x83\x4e\xff\x97\xf7\xeb\x55\xb5\x46\xb5\xba\xd8\xac\x11\xdf\x8e\x88\x42\xc2\x86\xad\x60\xef\xc2\xc0\x69\x7b\x18\x5d\x29\x46\x57\x27\x56\x65\x7d\x0e\x53\x0c\xcb\x6a\x7b\xb5\x2d\xf1\xe8\x55\xf2\x6e\xb5\x34\xd8\x6b\x3e\x2a\x58\x32\x0c\xf9\xf4\x1c\x41\x56\x41\x5b\xc5\x1f\xdf\xd0\x1f\xad\x6f\xc7\x76\xe7\x18\xe7\x2e\x13\xe9\xab\xb4\xe1\x28\x64\x3c\xab\x12\x12\x4e\xbc\xc0\xb4\x4b\x89\x92\x4b\x89\x79\xd7\x89\xdb\xb8\x96\x03\x7e\xbd\xf2\x92\xd3\xcf\x17\x23\x7b\x93\x82\x71\x2c\xf1\xf4\x9c\xc7\xe9\x74\x3f\x89\xae\x53\xa5\xac\xc9\xd3\x0f\x4d\x3f\xfb\x3f\xfb\x0a\x00\x00\xff\xff\xde\xf6\xc3\xf5\xda\x01\x00\x00"),
          path: "sql-api-migration.tml",
          root: "sql-api-migration.tml",
        },
      
        "sql-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x51\x6f\xda\x30\x10\xc7\xdf\x91\xf8\x0e\x37\xf1\x12\xa6\xcc\xbc\x57\xda\x03\x90\x0d\x55\xaa\x36\x5a\xd6\x27\x54\x29\xc6\x39\x12\x0f\xc7\xce\xec\xcb\x96\x0a\xe5\xbb\x4f\x4e\xd2\x02\x5b\x58\x8b\xd6\xbc\xd8\x77\xca\xfd\xff\xbf\xbb\xf3\x7e\xcf\x56\x64\x4b\x41\xec\xeb\xe6\x3b\x0a\x62\x5f\x78\x8e\x75\x0d\xab\xdb\x9b\x68\x06\xd3\xe5\xf5\x70\xf0\xf1\xe5\x6f\x38\x18\x0e\xd6\xef\xd6\x0b\x03\x77\x58\x18\x4b\x30\xe7\x36\x79\x08\x32\xa2\xc2\x5d\x4d\x26\xa9\xb1\x4d\x5a\x70\x9b\x30\x61\xf2\xc9\x86\x27\x29\x4e\xf6\x7b\xb6\xe4\x62\xc7\x53\x5c\x72\xca\xea\x7a\xfc\x8f\x8a\x36\xfc\xbb\xc4\x3b\xbf\xd8\x03\x48\x07\x1c\x78\x49\xe6\x43\x8a\x1a\x2d\x27\x4c\x60\x7e\x77\x1f\x81\xcc\x0b\x85\x39\x6a\xe2\x24\x8d\x86\xad\xb1\x40\x19\x42\xdc\x2b\xd9\xe9\xc6\x20\x35\x14\x2d\x45\xf3\xe7\x72\x97\xb2\x16\x27\x66\x9e\xe7\x5b\x86\xb0\x35\x4a\x99\x5f\x52\xa7\x90\x23\x65\x26\x01\xac\xa4\x23\xd7\x38\x88\xd2\x91\xc9\xc1\x14\x9e\x44\x1a\xed\xae\x7c\xd5\x68\x04\x9f\x2a\x14\xfe\x1a\xc7\x71\x6a\x86\x03\x1f\x06\x82\x2a\x10\x46\x13\x56\xc4\xe6\xed\x19\xc2\xb6\x82\x6d\xa9\x45\xf0\xde\xfd\x50\x6c\x75\x7b\x13\x82\xbf\x44\xb3\x31\xa0\xb5\xc6\x76\x47\x23\x74\x0e\xc8\x3d\x11\x49\xdd\xb4\x7c\x18\x8c\x1f\x18\x77\x50\xa0\x25\x2e\xb5\xaf\x20\xd3\x4c\xeb\x09\x73\x6e\x4a\x4d\x47\x9c\x4d\xdc\x07\x3a\x86\x40\x6a\x0a\x3b\xa8\x67\x1c\x2f\x61\x91\x13\x1e\x6b\x34\x89\xfe\x6e\x51\x61\x0e\x87\x8d\x74\x0f\xa0\xae\xd9\x99\xc5\xff\xd9\xfe\x68\x04\x0b\xa4\xd9\xe3\x67\x89\x2a\x39\xf2\x3c\x24\xfb\x7d\x77\xf8\x08\x8e\xac\xd4\x69\x08\x3f\xb9\x2a\xb1\x8b\xc6\x00\xc1\x05\x38\x21\xf4\x0c\x60\x81\x74\x4a\xd2\x8f\x50\x94\x1b\x25\xc5\x75\xf4\xec\x7c\x99\x71\xaf\x2f\x4c\x95\x3a\xf5\x9e\x2a\x75\x66\x7d\xeb\x87\xff\xf4\xbb\x2f\x92\xd3\x45\xb7\x89\x57\x75\xfb\x26\x9b\x8f\x50\xe1\x09\x40\x9b\x78\xe5\xb8\x8f\xe5\x7e\x07\x00\x00\xff\xff\xd7\xc5\x8d\x8b\x2b\x05\x00\x00"),
          path: "sql-api-readme.tml",
          root: "sql-api-readme.tml",
        },
      
        "sql-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x51\x6f\xe2\x38\x10\x7e\x2e\x12\xff\x61\x2e\xba\x93\xc2\x1e\xf2\xee\xde\x63\x4f\x3c\x94\xa6\x57\xed\x69\x55\xb8\x86\xee\x3e\x22\x13\x4f\xa8\x77\x4d\x4c\x6d\x07\x5a\x45\xfc\xf7\x93\xed\xc0\x96\x16\xda\xa4\xa5\x3a\xe9\x36\x3c\x81\x99\x6f\xc6\xdf\xe7\x99\xcf\x51\x8a\x02\x7e\x9d\xa7\x1c\x05\x83\xe3\x1e\xb8\x2f\xfd\xbb\x0b\x3a\x43\x20\xb1\x51\x79\x62\x20\x18\xe6\x13\xc1\x93\x4f\x51\x00\xab\x55\xbb\xe5\x00\x86\x4e\x6d\xf8\x14\xcd\x88\x4e\x37\x09\x02\x7d\x23\x02\x08\xbe\x69\x99\xf9\xd8\x76\x6b\x41\x15\x84\xed\x16\x00\x40\x22\xb3\x94\x4f\xa1\x07\xfa\x46\x90\x53\xf7\xa3\xf0\xff\xd8\x4f\xd4\xb7\x45\x8f\x41\x6a\x72\x8e\x06\xb3\x45\x18\x14\x05\x39\xbb\xf8\x62\x97\x57\xab\x71\xfc\xcf\xe7\xf1\xe8\x2c\x1e\x8d\xa3\x7e\xd0\xe9\xfe\xc0\x5d\x69\x54\x15\x50\x57\xf1\xd9\xe5\x16\x2e\xea\x7f\x1a\x56\xc0\x9d\x44\xd1\x43\xdc\x50\x2a\x53\x01\x39\x1c\x5c\x8e\x1e\x20\x23\xc5\x17\x95\x76\xeb\x03\x1f\xf1\x1c\x52\xad\x97\x52\xb1\x2a\xd5\x4f\xe2\xf8\xeb\xe0\x32\xda\xe4\x70\xa7\x61\xbf\x18\xd4\xe6\x54\x0a\xe8\x41\x50\x14\x42\x2e\x51\xad\x4f\x9a\x0c\x26\xdf\x30\x31\xc4\xe6\x22\x65\x42\x1b\x1d\xb4\x5b\x1d\x8b\x7e\xff\x1e\x46\xa8\xcd\x39\x9a\xa2\xd8\x01\x59\xad\x60\x41\x05\x67\xd4\xa0\x06\x73\x8d\xa0\xd0\x28\x8e\x0b\x2a\x40\xa6\x40\x61\x0f\xc8\xe5\x55\x98\x48\xc5\x20\x55\x72\x06\xd4\x36\x08\x9b\x90\x76\x2b\xcd\xb3\xe4\x99\x92\xa1\x81\x77\x76\x8f\x3c\x9b\x92\x51\xa7\x68\xb7\x8e\x70\x81\x99\xd1\xb6\x3b\x67\xb6\x7e\xa2\xc9\x05\x2e\xc3\x4e\xbb\x75\xc4\x53\x58\x87\x7e\x41\x35\x91\x1a\x43\x87\x58\x43\xb6\x11\x49\xae\x8d\x9c\x91\xd8\xd0\xe4\x7b\xc4\xf5\x5c\xd0\xbb\x50\x6a\x12\x1b\x26\x73\xd3\xb1\xf9\x36\x92\xd2\x39\xb7\xf5\xfc\xb6\x2d\xb6\xd4\xb8\x0b\x3e\x73\xd7\xb5\xfc\x05\x2e\xa3\x7e\xe8\xa7\x60\xfd\x4f\xc7\x09\x7b\x94\x98\xdb\x2e\x24\x34\x4b\x50\xd8\x44\x89\xcc\x0c\xde\x1a\xf2\x95\x9b\xeb\x11\x9f\xa1\xcc\x4d\xb8\x5e\xeb\xd3\xe4\xfb\x54\xc9\x3c\x63\x61\xa7\x0b\x1f\x3f\xc0\x3b\x30\x7c\x86\x24\xc6\x44\x66\xcc\x6e\x8b\x61\x8a\xaa\xcc\x16\xfa\xfc\x28\x70\xd6\x05\x54\xca\xcf\xf8\xad\xc9\x15\x6a\xf2\x59\x52\xb6\x53\xd6\x52\xdb\xbf\xe3\xc1\x45\xb8\x89\x7e\x2e\xb2\x54\xd8\x16\xf9\xa5\x07\x19\x17\xe0\xb4\xb5\x52\x68\xf2\x17\xe5\x02\x59\x18\xc4\x79\x92\xa0\xd6\x69\x2e\xc4\x1d\x08\x49\x19\x32\xb0\x60\x48\xa5\xda\xd7\x20\x65\x73\x1c\xc3\x6f\xbf\xdf\x90\xc0\xd1\x28\xd5\x2f\x93\xdb\xb1\x78\x65\xf2\xc0\xeb\xe4\xa5\xa3\x73\x4e\x22\x14\x68\x30\x74\x07\x63\xd5\x23\x6b\x13\xf4\x81\x25\xd1\xe3\x9e\x0b\x3e\x55\x48\xef\x07\x77\xfe\xac\x23\x03\x65\x76\xa3\xeb\x09\x78\x62\xab\x3c\x33\x12\xd8\xa4\xa6\x10\x75\xd3\x93\x52\x8b\xb1\x6f\x18\xcf\xf0\x1c\xcd\x6e\x2d\x6a\x1e\x79\x69\x09\xc8\x40\x1b\xa9\xaa\xed\xcb\x99\x42\x6d\xda\xaf\xa8\xe4\x14\x58\x6d\x3b\xde\x89\x10\x2f\x31\x3d\x21\x5e\x69\x7b\xfb\xeb\xfe\xc7\xce\xd7\xf8\x5e\xe3\x7b\xff\x4b\xdf\xf3\x08\xdd\xbd\x4f\xd4\x0f\x62\xff\x6e\xa0\x18\x2a\xcf\x37\xa0\x3a\x09\xba\x10\xcc\x9d\x3e\x63\xee\xa4\x7c\xb1\x1b\x5a\xa3\x28\x0b\xbf\xa9\x17\xd6\xac\xb3\xd6\x84\xa7\x20\x30\x0b\x4b\x64\x07\x7a\x3d\xf8\x50\x87\x9d\x11\x48\xb5\x81\x8f\x75\x4d\xb8\x06\xb1\x97\x96\xb8\xef\xf3\x7b\xe2\x7d\x9b\x3f\xf0\xf9\xc4\x2e\x72\x99\x55\x7e\xb6\x5d\x72\x73\xbd\xcb\xe4\x9f\x2c\xda\x98\x7c\x63\xf2\x8d\xc9\x1f\xd8\xe4\x2b\x8c\xfc\xd5\x9c\x3d\x1e\xf9\xdc\x2f\xbe\xd1\xc0\xfb\x92\xcd\xc0\x37\x03\xdf\x0c\xfc\xe1\x9f\xea\x2c\x91\x3f\x9a\xa6\x79\xbe\x69\x9c\x50\x9b\xfe\x80\xde\x76\xbf\xec\x68\x97\xd2\xb7\x36\xed\xf2\x03\x5c\xfe\xae\xd7\x3f\xde\x64\xdf\xb0\x83\xea\x17\xa8\x7a\x69\xf8\x29\x7b\xf4\x3e\x60\x26\x0f\xf1\x0a\xf4\xc9\x9a\xcd\xad\xd1\xdc\x1a\x87\x32\x80\x9f\xe0\x32\xd8\xa6\xb8\xff\x72\xac\xc5\xd5\x0d\xfa\x5b\xb2\xad\x5f\xe0\x1e\xdf\x2a\x6f\x73\x3d\xdd\x5e\x35\xba\xa9\x5b\x04\x23\x61\x8a\x06\x98\xd3\xb0\xf6\xde\x9e\xa3\x7c\x90\x22\xab\x76\xeb\xdf\x00\x00\x00\xff\xff\xd8\xb4\x82\x33\xe5\x1c\x00\x00"),
          path: "sql-api-test.tml",
          root: "sql-api-test.tml",
        },
      
        "sql-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4b\x8f\xdb\x38\x12\x3e\xc7\x80\xff\x43\xad\xb0\x18\xc8\x03\x8d\x82\xbd\xf6\x6c\x1f\xa6\x1f\x19\x04\xd8\xe9\x64\x93\xcc\xee\x61\x30\x08\x68\xb1\x64\x33\xd1\xc3\x21\xa9\x6e\x1b\x86\xff\xfb\xa2\x48\xea\x69\xcb\xb6\xdc\x4e\x36\x9d\x49\x03\x89\xdb\x12\x59\x8f\xaf\x8a\xf5\x15\x45\xf5\xf3\xe7\x80\x52\xe6\x52\x41\x18\x86\xe3\xd1\x3d\x93\xe0\x8f\x47\x00\x00\xb7\x52\xde\xe5\xfa\x45\x5e\x64\x1c\x2e\xdd\xa0\xf0\x0e\x1f\x7c\x4f\x62\x94\x4b\x0e\x59\xae\x21\xa6\xdb\xde\x64\x3c\x9a\x8c\x47\xe3\xd1\xf3\xe7\x90\xb2\xc5\x0b\x81\x09\x57\xc0\x31\x16\x19\x2a\x60\xa0\x57\x0b\x84\x38\x97\xc0\xe8\x36\xe8\x39\xd3\x80\xcb\x45\xae\xcc\x5d\x3b\xdc\x9f\x40\x8a\x7a\x9e\xf3\x70\x3c\x32\xe3\x6b\x41\x29\x5b\xfc\xa1\xb4\x14\xd9\xec\x4f\x91\x69\x94\x31\x8b\x70\xbd\x71\xfa\xdc\x18\x89\xba\x90\x99\x02\x3d\x37\x33\x41\x68\x85\x49\x0c\x2c\xe3\xb0\x90\xf9\xbd\xe0\x46\x95\xd5\x00\x3a\x87\x94\xe9\x68\x6e\x46\xab\x4f\x49\xf8\x8e\x4d\x13\x34\x92\xa0\xd2\x10\x8e\x47\x71\x91\x45\xe0\xa7\xb5\x29\x93\xda\x58\x7f\xb7\x55\x81\x05\x6a\x02\x6b\x0b\xa2\xb5\x0b\xd2\x00\x32\x91\x8c\x47\xa5\xd5\xeb\x75\xf8\x56\xcb\x22\xd2\xe1\xab\xe9\x07\x8c\x74\x78\xc7\x52\xdc\x6c\xba\xc8\x65\xb5\x39\xf0\x30\x17\xd1\xbc\x82\xad\xf6\xc4\x69\xb0\xd0\xe6\x31\xb0\x24\x31\x2a\x98\xd6\x52\x4c\x0b\x4d\x72\x94\xca\x23\xc1\x34\x72\x78\x10\xda\xba\x6d\x75\x70\x50\xc6\x8c\x42\x22\x30\x52\x1c\x09\x8e\x1c\xa6\x2b\x0b\x4d\x79\xaf\x8c\xc9\x7e\xb3\x1b\xc6\xae\xc7\xa3\x67\xc7\x42\x75\x10\x95\xeb\x3c\x53\x45\x8a\x72\x1f\x2e\x2c\x8a\x70\xa1\x55\x0d\x03\x67\x9a\xb9\x7b\x0f\x22\x49\x60\x8a\x10\x59\x39\xdc\xe8\x12\x99\xce\x8d\x93\x33\x71\x2f\xb2\x19\x88\x74\x91\x60\x8a\x99\xa6\x2f\xe7\x00\xa5\xb2\xba\x8d\x8a\xbb\xdc\x83\xc9\xc4\x62\x72\x10\x92\x9b\xab\xc6\xf2\xaa\xad\xb5\xfe\xba\x84\x87\x9b\x2b\xb8\x7e\xf3\xfb\x0d\xe4\x0b\x94\x4c\x8b\x3c\x53\x46\x64\xa1\x8c\x8b\x9f\x12\x72\x8e\x3c\x2a\x32\x8e\x32\x11\x19\x02\x9f\x1e\x70\xea\xe6\xca\x69\x33\xe9\x1d\xe5\x09\x58\x1f\xe8\x9b\x5a\x9a\xb5\x74\x73\x45\x5f\xf8\x12\x7e\xa4\x6f\x6f\xff\xfd\x2f\xfa\x9a\xa2\x96\x22\x52\xe5\x67\xf8\x9b\xfd\xa4\x5b\x9a\xd6\x1e\x69\x36\x8b\xf0\x25\xa7\x10\xe8\x55\x8d\xc0\x1d\x3e\x54\xab\x9b\x41\x86\x0f\x20\x32\xa5\x59\x16\x21\x85\xb9\xd7\xce\x72\xf9\x52\xc1\xb2\x2a\xac\xa5\x01\xa4\x5d\x2b\x82\xda\xf4\x00\x74\x4a\x85\xd0\xcc\x50\xd6\xa4\xdf\xc4\xcc\xc2\x37\x81\x1f\x7b\xd5\xb9\xe5\xce\x97\x70\x71\x69\x64\x91\xde\x94\x24\x93\xc8\x30\x0c\x4d\x7d\x6c\x54\x84\x1f\x0e\x89\xa2\x1f\xb5\xbc\x20\x09\xf5\x05\xbe\xbc\x00\xde\xbc\x10\xe5\xc9\x85\x85\xb0\x71\xd1\xf9\x77\x01\x69\xe3\xa2\x19\x74\x51\x01\x4d\xda\xd6\xf4\x9f\x9b\xbe\x71\x43\x37\x35\xf2\x37\x98\xa0\x46\x2a\x24\x98\xd2\xda\x32\xc5\x26\xcd\xef\xd1\x64\x8d\x23\x81\x58\xe6\xa9\x2d\x28\x53\x97\x59\xf4\xc5\xa5\x20\x87\x45\x31\x4d\x44\xf4\xf2\x26\x34\x12\xdf\x98\x39\xaa\x1a\x28\x14\x25\x69\x5a\x28\x0d\x73\x76\x8f\xc0\xdc\x78\x10\x1c\xee\x59\x52\x60\x40\x05\x4f\xa2\x52\xc8\x01\x85\x9e\xa3\xa4\x65\xc8\x60\xaa\xf2\x0c\x72\x09\x1f\xe8\x53\xb3\x99\x91\x4e\xbf\xda\x05\x8d\x59\x23\x31\x5e\xb3\xe8\x23\x9b\xe1\x66\x13\xf6\x40\xee\x52\xba\x2e\xf8\x7c\xba\x27\xd2\x13\x87\x8b\x1f\xe9\x25\x95\x14\x8d\x4b\x1d\x5e\xdb\xcf\xa0\xf2\xd7\x65\x9b\x5b\xcf\x25\x1b\xa4\x94\x1d\x65\xf6\xdd\xe1\xc3\x3b\xc9\x22\xf4\xbd\xfe\x24\xb6\xaa\x88\x5f\x4d\xf4\x31\x46\x09\x29\x9f\x86\xa5\x8c\xdb\x54\x68\xbf\xfc\xf2\x32\x8b\xf3\x63\x84\x05\xe5\x84\xff\x0a\x3d\xf7\xbd\xd2\x64\xaf\xb6\xbe\x3d\xc4\x5a\x99\x86\xb7\x19\xf7\x27\x93\x2a\x91\x45\x0c\x42\x39\xc7\x6f\x97\x0b\x21\x91\x13\x26\x15\xf5\xd1\x0f\x4a\x49\x2e\xc7\xa9\x0e\x6f\x09\x88\xd8\xf7\xdc\x0c\x98\x33\x45\xc1\xa5\x69\xa5\x7f\x06\xa2\x3e\xef\xca\xf9\x2f\x98\x48\xd0\x30\x1f\xb7\xf9\x69\x13\x71\xb8\x5b\xbe\x67\x12\xdf\x0b\x8c\xce\x28\x4f\xba\xb7\x4d\xe8\x3c\x43\x53\x56\xbb\xf5\xbe\xb4\xd4\x2d\x64\x94\xb2\x5a\x39\x15\x30\xce\x6f\x12\xcc\x97\x0e\x78\xca\x2b\x5b\x58\x02\x70\xd6\xbd\x17\xbc\x69\xde\xcf\x66\xde\xdf\x2e\xa9\x57\x68\xa2\x78\x36\x4c\x0e\x38\x7c\x04\x64\x07\x30\x71\xae\x5f\x5e\x02\xa7\x22\x68\x5a\xc8\x37\xf9\x83\x6a\x7a\xd3\xc0\xae\xd1\x62\xd6\xf7\x37\x47\x40\x7c\x60\x0d\x58\xc0\xf9\x19\x61\xe8\x96\xef\x56\x37\x77\x2d\x91\x75\x0a\x25\xe3\xbc\x59\x25\xab\x6e\x63\x77\x95\x6c\x32\x9a\x9e\x63\x87\xf9\x0f\x16\xb0\xff\x63\x71\x7d\x54\x21\xb5\xb8\xed\x2e\xa4\x98\x60\x3a\x04\x83\xc7\x56\x5a\x6b\xcb\x99\x2a\x6d\x29\xac\x3f\xb1\xc8\xbd\xf0\xf5\x93\x2a\xb7\x91\xcd\xf2\xc3\xe5\x76\x8f\x6f\xa7\xd6\xdc\x67\xcf\x9e\x3d\xeb\x2d\x06\xeb\x35\x81\xe3\xcf\x99\x7a\x41\x69\xe7\x82\x02\x9e\xdd\x80\x78\x13\xd8\x6c\xb6\x2a\x54\x5d\x9c\xdf\xb2\xfb\x56\x69\x26\xeb\x7b\x2b\xf1\x4e\xc8\xda\xb7\x1b\x0d\x58\x3f\x88\x7d\x0b\xa8\x02\xb7\x5f\x66\x0f\x8c\x87\x26\x90\x5b\x2e\x36\x47\x0c\xde\x11\x85\xce\xa4\x2a\x2f\xfb\x6a\x75\x2b\x44\x47\xd4\xec\xeb\x4e\x7a\x6d\xb5\xb3\x87\x3d\xef\xf7\x78\x52\x66\x0a\x26\x0a\x9b\xe9\x60\xea\x4e\xa6\x4d\x3e\x94\xbb\x7c\x7f\xbd\x6e\x3c\x7d\x70\x8b\xf9\x97\x88\xb6\x00\x60\xa5\x82\x47\x35\xd2\x03\xef\x83\xf9\xd8\x6c\xb6\x29\xb0\x3f\xc1\x9c\xca\x61\x39\xf6\xf8\x8c\x1a\xb4\x0c\x3f\x15\x28\x57\x5e\x6d\xeb\xb0\x2e\xe0\x73\xa4\xc3\xa3\xcc\xaf\xc3\x9f\xf1\xcd\x66\x1f\x95\xff\x8a\xfa\x97\x24\xb9\x5a\xbd\x92\x1c\x65\x67\xeb\xa3\xa5\x40\xe2\xd1\x24\x31\x49\x85\x99\x56\x76\xfb\x53\x91\x7a\x45\xe8\xb9\x9d\x9e\xb9\xdf\xa6\x2b\x30\xd2\x0d\xed\xaa\xa3\xf9\xb1\x65\xcc\x6e\x9a\xb4\x8a\xca\xad\x6d\xa9\xac\xdc\x7c\x80\xff\xc7\x9f\x03\x38\x74\xfb\xe1\x95\x0a\xde\x07\xcd\x74\xb6\x16\x91\x29\x4e\x59\xa5\x33\x80\x9f\xfe\x41\xff\x26\x2d\x6c\x49\x82\xcd\x81\x36\xc0\x15\x9a\xca\xc0\x29\x5d\xeb\xd2\xdc\x4c\x12\x78\xf5\xbe\x5f\x25\xa2\xbb\xe1\x3f\xbc\xaf\xd3\xab\x05\x3e\xd5\xde\xa8\x46\x7a\x58\xd0\x03\x58\xb0\x19\x52\xc3\x19\x10\xfa\x8b\x3c\x53\xf8\x1a\xe5\x6b\x77\xf1\x84\x9c\x30\xa2\xda\x89\x31\xb8\xbb\xb2\xde\x9c\xa9\xbb\x2a\x85\x3d\x89\xae\xa9\x2a\x1b\xa7\x55\xb2\x9e\x4a\xbb\xb5\x45\xca\x44\x62\x57\xe0\x76\x7b\x74\xcf\x24\xe8\x5c\xb3\x84\x22\x59\x5f\x92\x42\x63\xaa\x60\x50\x32\x1c\xea\xb8\xdc\xc3\xcd\x6e\xcb\x45\x8a\x02\x6b\x43\xd0\x21\x47\x1b\x4b\x97\xa0\x4d\x96\xec\x56\x17\xca\xea\xad\x8c\x9e\xfc\xbc\xc5\xbb\x8f\xe6\xd3\x56\x9d\x2f\x0b\x53\xff\x93\x46\x77\xb8\x61\x69\xe0\x4c\xd1\x6d\x5a\x7e\xec\x8e\xba\x9b\x0b\x0e\xed\x9d\x3b\x6c\x68\xef\xb2\x7b\xa6\xf6\x53\x77\x9c\x4b\x78\x1f\x98\xb8\x52\x28\x25\xcb\xa8\xb8\x98\x74\xea\x58\x45\x89\x36\x74\x27\xb7\xd3\xff\x8b\x4b\xbb\xad\x28\x1f\x9f\x93\xb6\xbd\x1d\xd4\xc0\xa8\xbb\xb3\x81\x72\xbf\x6e\x4e\x0f\xc8\xcd\xbe\xa8\x9f\x18\x70\x92\xeb\x59\xe4\x4e\xc8\x85\xbd\xeb\xbd\x2f\xb0\x36\x2e\x97\xc0\x16\x0b\xcc\xb8\x2f\xdd\x6a\x34\x6d\xf1\xd6\x43\x97\xed\x06\x79\xc1\x32\x11\xf5\x54\xe3\x12\x0c\xc3\xa3\xd5\x01\xca\x6e\xcc\x3a\xa7\x22\xde\xde\x96\x4c\xb6\x4b\x46\xb7\x43\xbb\x5a\xd9\xb3\xba\x46\x13\xb1\xff\x71\x74\x6c\x86\x7f\xc4\x95\xe9\x2a\x2c\xc1\x1b\x61\xcd\x63\xc3\xbf\x50\x67\xe1\x00\xdc\xdd\x5d\x10\x4c\x65\x2f\x61\x2c\x86\xd6\x01\x15\xf8\x83\xba\x87\x33\x74\x0e\xe7\x6b\x1b\xba\xcb\xf5\x23\xae\x9c\x8f\xdf\x44\x33\xd1\xe7\xce\xd9\xba\x8c\x01\x91\x77\xe7\xbb\x3b\x7a\x91\xe1\x7c\x70\x52\xd3\xd1\xdd\x86\xff\x8a\xba\xd9\x5f\xfc\x40\x66\x98\x64\x2f\x21\x3b\xd3\x7e\x7c\x50\xff\xd0\xd3\x44\x74\x94\x0f\x79\xf2\x33\x30\xb0\x70\x5a\x8b\x31\x30\x11\xf6\x75\x21\x3b\xfb\x90\x53\x13\xed\xeb\xa3\x32\x9b\x66\x5d\x06\x3b\x9e\xba\xaa\x03\xc5\xe6\x7e\xf8\x2f\x46\x57\x47\x1e\xb5\x7e\x4b\xcc\x74\xd2\xa9\xec\xd3\x60\xa9\x2f\x76\x32\xfb\x35\xb0\x15\xf4\xd2\x15\x40\x9b\xb0\xd6\x6b\xf8\xfb\xc2\x36\xaa\x14\x19\xfa\xe5\x6a\x45\x32\xeb\xa9\xe5\x99\x8a\xb7\x3d\x51\xb3\x19\x4d\x9b\xa1\x7e\xc7\x66\x95\x20\x4f\x7d\x4a\xca\xa7\xd4\x9b\x56\xad\x3d\x92\x1e\xbd\xf5\xda\xc8\x0e\xff\x43\xb5\x60\xb3\x39\xe6\xb4\xfa\x3b\x5f\x7e\x56\xbe\xfc\xf6\xd8\xf2\x10\x5f\xfe\xbe\xe0\x4c\x23\x14\xea\x3b\x5b\x1e\x64\x4b\x8b\xd5\x51\x84\xf9\xe5\xcf\xd8\xad\x71\x67\x22\xcd\x52\xd8\xb7\xf2\x36\x53\x2c\x32\xa1\xe6\x41\x19\xb6\xb6\xc4\x53\x5e\x60\xe9\xbc\x5c\x74\xa6\x97\x9b\x4e\xa2\xa9\xc1\x14\xd5\xcb\x9e\xee\x4c\xbf\x4b\x9e\x5b\x74\xe6\xd6\x41\xe7\x5c\xff\x8b\x10\x5a\x61\xcb\xd5\xc0\x03\xd9\xf2\xc8\xd2\x1e\xd0\x0f\x0a\xf7\x1e\x97\xbe\xc8\xf3\xe6\x81\x6c\x75\x1c\xfd\x98\x27\xb0\x9f\xf1\x54\x7e\x3b\x41\x48\xe3\xf7\x04\xf9\x7a\x13\xa4\xdd\x2e\x1c\xe0\x0d\xd7\x33\x9c\x76\xe8\x76\x3c\x5e\xad\xa0\xec\x7d\x47\xf0\x76\x89\x51\xf3\x4f\x51\x88\xce\x4d\xee\xba\xbf\x5f\x48\x12\xc2\x90\xfa\x01\x5c\x62\x54\x98\x5b\x79\x0c\x0c\xa2\x42\xe9\x3c\xad\xc7\xb3\x19\x13\x99\xd2\x66\xa8\xf1\xe3\xe8\xde\x80\x4c\xd8\xdd\x19\xc4\x4b\x23\xdf\x2f\x5f\xd9\x0f\xdc\xeb\xf0\x93\x72\x8f\xfc\x38\xde\x27\xc5\x67\x62\x7d\x2b\xea\x49\x10\xba\x8d\x23\xd6\x7f\x7e\x71\xe6\xd3\xc1\x23\xde\x3f\x8e\x97\xbe\x2d\x77\x56\x87\x5a\x9e\xe3\x15\xe3\xcf\xed\xd7\x57\xf2\x0e\xf1\xab\xd2\xbd\xd2\xe1\xc3\x65\xa4\xbf\x00\x98\x05\xba\x2b\x1d\xbb\x6b\x71\x02\xd3\x3c\xaf\x22\xa3\x30\x41\xf7\xb7\x36\xf6\x27\x62\x0a\xe1\x9f\x3f\x45\x7a\x19\xde\xe4\x19\xfa\x93\x8b\x9d\x98\x68\x59\x60\x83\x4a\x31\x66\x45\xa2\x77\x0f\x8d\x59\xa2\xb0\x42\x66\x33\x1e\xfd\x2f\x00\x00\xff\xff\x7d\xb7\xa5\xf3\x11\x38\x00\x00"),
          path: "sql-api.tml",
          root: "sql-api.tml",
        },
      
        "sql-simple-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xc1\x6a\xf3\x30\x10\x84\xef\x06\xbf\xc3\xfe\xe4\x92\xfc\xa4\xf6\x3d\xd0\x43\x93\x94\x52\xc8\x21\x69\x7a\x0b\x85\x6c\xe4\xb5\x24\x6a\x6b\x1d\x69\x45\x05\x21\xef\x5e\xe4\x94\x5e\x0a\xad\x2e\xa3\x59\xf8\x66\x66\xbf\xdb\xac\x97\xf0\xb0\x7d\x2e\x8b\xfb\xbf\x5f\x59\x94\xc5\xe1\xdf\xe1\x89\xe1\x85\x06\xf6\x02\x2b\xf4\xcd\xdb\xd4\x88\x0c\x61\x51\xd7\x9a\xfd\x78\x56\xe8\x9b\x4a\x71\x5f\x9f\xb0\xd1\x54\x5f\x2e\xd5\x16\xd5\x3b\x6a\xda\xa2\x98\xeb\x75\xf6\x0b\x71\xb3\x3f\x91\xdc\xfc\xbd\x15\x6c\x00\x04\x8c\xc2\x77\x9a\x1c\x79\x14\x6a\xe0\x84\x81\xc0\xf6\x43\x47\x3d\x39\x41\xb1\xec\xa0\x65\x0f\x94\x48\x45\xb1\x4e\x43\x38\x77\x70\x8e\xe4\x2d\x05\x10\xe3\x39\x6a\x03\x08\xcd\x09\xac\x0b\x82\x4e\x51\x95\x5b\x5e\x0d\x41\xcb\x5d\xc7\x1f\x99\xe9\x49\x0c\x37\x40\xc9\x06\x09\x63\x9e\x8a\x41\xb8\x07\x1e\x72\xaf\x65\x17\x16\x99\x9a\x4c\xe0\x31\x91\xca\xdf\xe3\xf1\xa8\xb9\x2c\xb2\x9d\x2a\x49\xa0\xd8\x09\x25\xa9\x56\x37\x9d\x43\x9b\xa0\x8d\x4e\x4d\xff\x87\x73\x57\xed\x77\x9b\x79\x5e\x56\xad\x97\x33\x20\xef\xd9\x7f\xc9\x18\x94\xf3\x3e\x03\x00\x00\xff\xff\x7f\xa3\x9c\x12\xa3\x01\x00\x00"),
          path: "sql-simple-readme.tml",
          root: "sql-simple-readme.tml",
        },
      
        "sql-simple.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\xcb\x6e\xeb\x36\x10\x5d\x47\x80\xfe\x61\xaa\x45\x21\x05\x2a\xbd\x77\xeb\x4d\x62\x07\x08\x90\xb8\xc8\xa3\xe8\x9a\x26\x47\x36\x01\x89\x74\x48\x2a\x56\x10\xf8\xdf\x8b\x21\x29\x5b\x35\x72\x91\xbb\xbc\x5e\x98\x22\xe7\x75\xce\x99\x21\x67\x33\x40\x6b\x8d\x75\xc0\x18\xcb\xb3\x77\x6e\xa1\xcc\x33\x00\x80\x95\xb5\x6b\xe3\xef\x4c\xaf\x25\x2c\x92\x13\x5b\xe3\xa1\x2c\x2c\x0a\x63\x25\x68\xe3\xa1\x21\x73\x51\xe5\x59\x95\x67\x79\x36\x9b\xc1\xcb\xd3\xc3\xf2\x06\x24\x36\x4a\xa3\x03\x0e\xce\xdb\x5e\xf8\xde\x22\x1c\x76\x4a\xec\x60\x6f\xcd\xbb\x92\x08\xcb\x1b\xb8\x7d\xfe\x67\x09\x66\x8f\x96\x7b\x65\xb4\x0b\xe1\xbd\x53\x7a\x0b\xee\xad\x05\xee\xc0\xef\x10\x7a\x2d\xd1\xb6\x4a\x23\xc8\x0d\xcb\x33\xff\xb1\xc7\x54\x23\x66\xfe\xcc\xb3\xab\x6b\xf7\xd6\xb2\x97\xa7\x07\xc2\x2d\x4c\x4b\x16\xa5\xb7\xb4\x73\x03\xe5\x62\xcb\x1b\xda\x74\xe8\xad\x12\x6e\x5c\xd9\x63\x5c\xc9\xe4\xf9\xa6\x0d\x15\x5e\xe9\xe3\x5e\xa2\xf6\xca\x7f\xe4\xd9\x31\xb1\x5a\xe3\x01\x2c\xfa\xde\x6a\xe2\xa4\xf1\x00\x4a\x3b\xcf\xb5\x40\x30\x4d\xc4\xc3\xf2\xac\xe9\xb5\x20\xd7\x32\xa6\x8b\x30\x6a\xe8\x2e\x2b\xd6\x67\x5c\x35\xf8\x8e\x94\x0f\x11\x2e\x96\x7f\x54\xdb\x28\x49\x05\xd7\x21\xf5\x67\x6c\x88\x1c\x60\xbe\x08\x71\x54\xa3\xa3\x2c\x14\xce\x18\xab\xa2\x43\x44\x08\xbf\x4f\x83\xae\x5e\x9e\x1e\xe6\x20\x87\x3a\x6e\xe9\xe7\x86\x39\x85\x9e\x0f\x84\x69\xe7\x51\x82\xc9\x61\xc2\x3c\x87\x6e\x72\x18\x9c\xe6\x27\xa1\xd6\xbc\xc3\x4f\xfa\x4b\xe1\xc7\xe4\x7a\x3c\x2b\xb7\x1a\x50\x8c\x4d\x27\xed\x48\x23\xe2\x96\xc6\x81\xb7\xad\x39\xc4\x4e\xe3\x80\xa2\x0f\x26\xd3\x00\x07\xd1\x3b\x6f\xba\xb3\x3f\xdf\x72\xd2\x3c\xb8\x86\x62\xa3\xe0\x65\x27\x37\x49\xa8\x2a\x94\x2b\x85\x1f\x40\x18\xed\x71\xf0\xec\x36\xae\x35\x34\x43\xc8\x55\x8e\xc3\x52\xa7\x0e\x54\x71\xb4\xd3\x02\x49\xb6\x8e\xa4\x1e\xdb\xb6\xc6\xc3\xab\xe5\x02\xcb\x22\x76\x9a\x8a\x14\x49\x73\x89\x0d\x5a\xe8\xe4\x86\x8d\xde\xab\x4e\xf9\x72\xdc\xdc\xeb\xc6\xfc\x3f\xac\x1e\x4d\xff\x2a\xbf\x8b\x69\x3b\xb6\xd2\xb2\xac\xaa\x70\x87\x28\xa9\x6a\x40\xb9\x84\x7c\x35\xec\x95\x45\x49\xa4\xaa\x11\x1d\xfd\xd0\x5a\xc2\xd8\x74\x9e\xad\x08\x79\x53\x16\x29\x02\x76\xdc\x01\xc6\xb0\x11\x66\xe0\xf4\x23\x90\x63\xfc\x1d\x57\x2d\x4a\xf0\x26\xf5\x02\xcf\xb7\xf3\x02\x77\x59\x84\x16\x14\x75\x48\x2a\x4c\x7b\x69\x0e\x62\x16\x35\xa1\x8c\xe9\x23\xbd\x11\x4a\x9a\x54\xb4\xf6\x34\x30\x27\xe6\x23\xb1\x81\x1a\x1b\x3b\x45\x1f\x6e\xa8\xfe\x0c\xb6\xdf\x16\xa0\x55\x3b\x95\xe2\x17\x22\x96\xf0\x2f\x16\x20\x69\xbc\xc2\x1b\xfa\x4c\x13\x3e\x81\x3b\x11\x60\xf2\xc6\x9e\xed\xc7\x9f\xd0\xe9\x9b\x81\xfb\x7b\xa4\x37\x12\x96\xdf\xf2\x3c\x0d\x5f\xaa\xa9\x55\x9b\x6e\x71\xb8\x65\x5f\xcd\xe3\xe5\x25\xab\x60\x63\xcc\xa9\x33\x0e\x5b\x0c\xef\xf3\xe9\x99\xe1\x0e\xe1\xaf\x3f\x84\x1f\xd8\xd2\x68\x2c\xab\xf9\x97\x9a\x78\xdb\xe3\xd9\x20\xb1\xe1\x7d\xeb\xbf\x76\x6d\x78\xeb\x70\xfa\xe4\xfc\x17\x00\x00\xff\xff\x52\x80\x92\x6b\xce\x06\x00\x00"),
          path: "sql-simple.tml",
          root: "sql-simple.tml",
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
