package gofile

import (
	"os"
	"time"
)

type FileInfo struct {
	location   string
	name       string
	size       int64
	mode       os.FileMode
	modTime    time.Time
	isDir      bool
	visibility string
}

// FileInfo interface
func (f *FileInfo) Location() string   { return f.location }
func (f *FileInfo) Name() string       { return f.name }
func (f *FileInfo) Size() int64        { return f.size }
func (f *FileInfo) Mode() os.FileMode  { return f.mode }
func (f *FileInfo) ModTime() time.Time { return f.modTime }
func (f *FileInfo) IsDir() bool        { return f.isDir }
func (f *FileInfo) Visibility() string { return f.visibility }
