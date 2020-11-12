package main

import "os"

func GetFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY, 0444)
}
