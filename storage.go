package gofile

import "fbnoi.com/goutil/collection"

var disk = map[string]Driver{}

type Driver interface {
	FileExists(location string) bool
	DirectoryExists(location string) bool
	Has(location string) bool

	Write(location string, content []byte, flag int) (int, error)
	WriteString(location, content string, flag int) error
	List(location string) collection.Collection[*FileInfo]
	Read(location string) ([]byte, error)
	ReadString(location string) (string, error)

	Delete(location string) error
	Create(location string) error
	CreateDirectory(location string) error
}

func Disk(name string) Driver {
	if disk, ok := disk[name]; ok {
		return disk
	}
	panic("")
}
