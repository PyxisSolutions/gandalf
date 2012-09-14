// Package fs provides types and methods for interacting with the filesystem,
// as an abstraction layer.
//
// It provides an implementation that uses the operating system filesystem, and
// an interface that should be implemented if you want to provide your own
// filesystem.
package fs

import (
	"io"
	"os"
)

// File represents a file in the filesystem.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
	WriteString(s string) (ret int, err error)
}

// Fs is the filesystem interface.
//
// Any simulated or real filesystem should implement this interface.
type Fs interface {
	// Create creates a file in the filesystem, returning the file and an
	// error, if any happens.
	Create(name string) (File, error)

	// Mkdir creates a directory in the filesystem, return an error if any
	// happens.
	Mkdir(name string, perm os.FileMode) error

	// MkdirAll creates a directory path and all parents that does not exist
	// yet.
	MkdirAll(path string, perm os.FileMode) error

	// Open opens a file, returning it or an error, if any happens.
	Open(name string) (File, error)

	// OpenFile opens a file using the given flags and the given mode.
	OpenFile(name string, flag int, perm os.FileMode) (File, error)

	// Remove removes a file identified by name, returning an error, if any
	// happens.
	Remove(name string) error

	// RemoveAll removes a directory path and all any children it contains. It
	// does not fail if the path does not exist (return nil).
	RemoveAll(path string) error

	// Stat returns a FileInfo describing the named file, or an error, if any
	// happens.
	Stat(name string) (os.FileInfo, error)
}

// OsFs is a Fs implementation that uses functions provided by the os package.
//
// For details in any method, check the documentation of the os package
// (http://golang.org/pkg/os/).
type OsFs struct{}

func (fs OsFs) Create(name string) (File, error) {
	return os.Create(name)
}

func (fs OsFs) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func (fs OsFs) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs OsFs) Open(name string) (File, error) {
	return os.Open(name)
}

func (fs OsFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

func (fs OsFs) Remove(name string) error {
	return os.Remove(name)
}

func (fs OsFs) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (fs OsFs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
