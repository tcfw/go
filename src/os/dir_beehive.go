package os

import (
	"errors"
	"sync"
)

// Auxiliary information if the File describes a directory
type dirInfo struct {
	mu   sync.Mutex
	buf  *[]byte // buffer for directory I/O
	nbuf int     // length of buf; return value from Getdirentries
	bufp int     // location of next record in buf.
}

func (f *File) readdir(n int, mode readdirMode) (names []string, dirents []DirEntry, infos []FileInfo, err error) {
	return nil, nil, nil, errors.New("not implemented")
}
