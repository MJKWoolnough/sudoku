package main

import (
	"net/http"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
}

func (fi FileInfo) Name() string {
	return fi.Name
}

func (fi FileInfo) Size() int64 {
	return fi.Size
}

func (FileInfo) Mode() FileMode {
	return os.FileMode
}

func (fi FileInfo) ModTime() time.Time {
	return fi.ModTime
}

func (FileInfo) IsDir() bool {
	return false
}

func (FileInfo) Sys() interface{} {
	return nil
}

type File struct {
	*strings.Reader
	FI os.FileInfo
}

func (File) Close() error {
	return nil
}

func (f *File) Readdir(_ int) ([]os.FileInfo, error) {
	return make([]os.FileInfo, 0), nil
}

func (f *File) Stat() (os.FileInfo, error) {
	return f.FI, nil
}

type FileSystem map[string]*File

func (fs FileSystem) Open(name string) (http.File, error) {
	f, ok := fs[name]
	if !ok {
		return os.ErrNotExist
	}
	return f, nil
}

var fs = make(FileSystem)
