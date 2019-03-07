package core

import "context"

type PackageFetcher interface {
	GetStepPackage(context.Context, *StepPackageRequest) (*StepPackageResponse, error)
}