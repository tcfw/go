package os

import "errors"

func statNolog(name string) (FileInfo, error) {
	return nil, errors.New("not implemented")
}

func lstatNolog(name string) (FileInfo, error) {
	return nil, errors.New("not implemented")
}
