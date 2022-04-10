package gofile

import (
	"io/ioutil"
	"os"

	"fbnoi.com/goutil/collection"
	"github.com/pkg/errors"
)

type LocalDriver struct{}

func (ld *LocalDriver) FileExists(location string) bool {
	info, err := os.Stat(location)
	if err == nil {
		return !info.IsDir()
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(location)
}

func (ld *LocalDriver) DirectoryExists(location string) bool {
	info, err := os.Stat(location)
	if err == nil {
		return info.IsDir()
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(location)
}

func (ld *LocalDriver) Has(location string) bool {
	_, err := os.Stat(location)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(location)
}

func (ld *LocalDriver) Write(location string, content []byte, flag int) (int, error) {
	fh, err := os.OpenFile(location, flag, 0644)

	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer fh.Close()

	n, err := fh.Write(content)

	if err != nil {
		return n, errors.WithStack(err)
	}
	return n, nil
}

func (ld *LocalDriver) WriteString(location, content string, flag int) (int, error) {
	fh, err := os.OpenFile(location, flag, 0644)

	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer fh.Close()

	n, err := fh.WriteString(content)

	if err != nil {
		return n, errors.WithStack(err)
	}
	return n, nil
}

func (ld *LocalDriver) List(location string) (collection.Collection[*FileInfo], error) {
	if dir, err := os.ReadDir(location); err == nil {
		cols := &collection.ArrayList[*FileInfo]{}
		for _, entry := range dir {
			if info, err := entry.Info(); err != nil {
				return nil, errors.WithStack(err)
			} else {
				cols.Add(ld.fromLocalFile(info))
			}
		}
		return cols, nil
	} else {
		return nil, errors.WithStack(err)
	}
}

func (ld *LocalDriver) Read(location string) ([]byte, error) {
	if fh, err := os.Open(location); err == nil {
		defer fh.Close()

		if content, err := ioutil.ReadAll(fh); err == nil {
			return content, nil
		}

		return nil, errors.WithStack(err)
	} else {
		return nil, errors.WithStack(err)
	}
}

func (ld *LocalDriver) ReadString(location string) (string, error) {
	if fh, err := os.Open(location); err == nil {
		defer fh.Close()

		if content, err := ioutil.ReadAll(fh); err == nil {
			return string(content), nil
		}

		return "", errors.WithStack(err)
	} else {
		return "", errors.WithStack(err)
	}
}

func (ld *LocalDriver) Delete(location string) error {
	if err := os.Remove(location); err != nil {

		return errors.WithStack(err)
	}

	return nil
}

func (ld *LocalDriver) Create(location string) error {
	if ld.FileExists(location) {
		return nil
	}
	f, err := os.Create(location)
	if err != nil {
		return errors.WithStack(err)
	}
	f.Close()
	return nil
}

func (ld *LocalDriver) CreateDirectory(location string) error {
	return os.Mkdir(location, os.ModeDir)
}

func (ld *LocalDriver) fromLocalFile(file os.FileInfo) *FileInfo {
	return &FileInfo{
		location: file.Name(),
		name:     file.Name(),
		size:     file.Size(),
		mode:     file.Mode(),
		modTime:  file.ModTime(),
		isDir:    file.IsDir(),
	}
}
