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
		return "", NewPackageNotFoundErr(pkg.Name, pkg.Version, os, arch)
	}
	path := ""
	switch arch {
	case `amd64`:
		path = binary.Amd64
	case `i386`:
		path = binary.I386
	}
	if path == "" {
		return "", NewPackageNotFoundErr(pkg.Name, pkg.Version, os, arch)
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


type PackageNotFoundErr struct {
	Id      string
	Version string
	Os      string
	Arch    string
}

func NewPackageNotFoundErr(id string, version string, os string, arch string) *PackageNotFoundErr {
	return &PackageNotFoundErr{Id: id, Version: version, Os: os, Arch: arch}
}

func (e PackageNotFoundErr) Error() string {
	return fmt.Sprintf("Package not found %s@%s for OS: %s, ARCH: %s", e.Id, e.Version, e.Os, e.Arch)
}

