package core

import (
	"context"
	"fmt"
)

type PackageFetcher interface {
	GetStepPackage(context.Context, *StepPackageRequest) (*StepPackageResponse, error)
}

func (pkg *Package) GetOSArchBinaryPath(os, arch string) (string, error) {
	binary := pkg.getOSBinaries(os)
	if binary == nil {
		return "", NewExecutableNotFoundError(os, arch)
	}
	path := ""
	switch arch {
	case `amd64`:
		path = binary.Amd64
	case `i386`:
		path = binary.I386
	}
	if path == "" {
		return "", NewExecutableNotFoundError(os, arch)
	}
	return path, nil
}

func (pkg *Package) getOSBinaries(os string) *Binary {
	var binary *Binary
	switch os {
	case `darwin`:
		binary = pkg.Executables.Darwin
	case `linux`:
		binary = pkg.Executables.Linux
	case `windows`:
		binary = pkg.Executables.Windows
	}
	return binary
}

type ExecutableNotFound struct {
	os, arch string
}

func NewExecutableNotFoundError(os, arch string) error {
	return ExecutableNotFound{
		os:   os,
		arch: arch,
	}
}

func (e ExecutableNotFound) Error() string {
	return fmt.Sprintf("Executable not found OS: %s, ARCH: %s", e.os, e.arch)
}