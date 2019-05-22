package core


import (
	"context"
	"fmt"
	"golang.org/x/xerrors"
)

type MultiRegistryPackageStore struct {
	registries []PackageRegistry
}

func NewMultiRegistryPackageStore(registries []PackageRegistry) PackageRegistry {
	return &MultiRegistryPackageStore{registries: registries}
}

func (p MultiRegistryPackageStore) Ping(context.Context, *EmptyMessage) (*BasicResponse, error) {
	return &BasicResponse{Success: true}, nil
}

func (p MultiRegistryPackageStore) GetPackage(ctx context.Context, req *StepPackageRequest) (*RegisteredPackage, error) {
	for _, v := range p.registries {
		if pkg, err := v.GetPackage(ctx, req); err == nil {
			return pkg, nil
		}
	}
	return nil, NewPackageNotFoundErr(req.Id, req.Version, req.Os, req.Arch)
}

func (p MultiRegistryPackageStore) GetPackageById(ctx context.Context, req *StepPackageIdRequest) (*RegisteredPackage, error) {
	for _, v := range p.registries {
		if pkg, err := v.GetPackageById(ctx, req); err == nil {
			return pkg, nil
		}
	}
	return nil, xerrors.Errorf("Unable to find package with ID %s", req.Id)
}

func (p MultiRegistryPackageStore) GetStep(ctx context.Context, req *GetSingleStepRequest) (*GetSingleStepResponse, error) {
	for _, v := range p.registries {
		if stp, err := v.GetStep(ctx, req); err == nil {
			return stp, nil
		}
	}
	return nil, xerrors.Errorf("Unable to find step with ID %s", req.StepId)
}

func (p MultiRegistryPackageStore) SearchStepsForLibrary(ctx context.Context, req *SearchStepsRequest) (*SearchStepsResponse, error) {
	foundSteps := map[string]bool{}
	var steps []*LibraryStep

	for _, v := range p.registries {
		if stepResp, err := v.SearchStepsForLibrary(ctx, req); err == nil {
			for _, v := range stepResp.Steps {
				key := fmt.Sprintf("%s_%s", v.Name, v.PackageName)
				if foundSteps[key] == true {
					continue
				}
				steps = append(steps, v)
			}
		}
	}
	return &SearchStepsResponse{
		Steps: steps,
	}, nil
}

func (p MultiRegistryPackageStore) GetAllStepsForLibrary(ctx context.Context, req *GetAllStepsRequest) (*GetAllStepsResponse, error) {
	foundSteps := map[string]bool{}
	var steps []*LibraryStep

	for _, v := range p.registries {
		if stepResp, err := v.GetAllStepsForLibrary(ctx, req); err == nil {
			for _, v := range stepResp.Steps {
				key := fmt.Sprintf("%s_%s", v.Name, v.PackageName)
				if foundSteps[key] == true {
					continue
				}
				steps = append(steps, v)
			}
		}
	}
	return &GetAllStepsResponse{
		Steps: steps,
	}, nil
}

func (p MultiRegistryPackageStore) GetPackageSteps(ctx context.Context, req *GetStepsForPackageRequest) (*GetStepsForPackageResponse, error) {
	for _, v := range p.registries {
		if stp, err := v.GetPackageSteps(ctx, req); err == nil {
			return stp, nil
		}
	}
	return &GetStepsForPackageResponse{Steps: []*LibraryStep{}}, nil
}

func (p MultiRegistryPackageStore) GetAllPackageNamesForLibrary(ctx context.Context, req *GetAllPackagesInfoRequest) (*GetAllPackagesInfoResponse, error) {
	foundPackages := map[string]bool{}
	var packages []*LibraryPackage

	for _, v := range p.registries {
		if pkgResponse, err := v.GetAllPackageNamesForLibrary(ctx, req); err == nil {
			for _, v := range pkgResponse.Packages {
				key := fmt.Sprintf("%s_%s", v.PackageName, v.Version)
				if foundPackages[key] == true {
					continue
				}
				packages = append(packages, v)
			}
		}
	}
	return &GetAllPackagesInfoResponse{
		Packages: packages,
	}, nil
}

func (p MultiRegistryPackageStore) GetPackageInfoForLibrary(ctx context.Context, req *GetPackageInfoRequest) (*GetPackageInfoResponse, error) {
	for _, v := range p.registries {
		if stp, err := v.GetPackageInfoForLibrary(ctx, req); err == nil {
			return stp, nil
		}
	}
	return nil, xerrors.Errorf("Package not found: %s", req.PackageName)
}
