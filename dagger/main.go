// A generated module for Wheelbuilder functions
//
// This module contains a function BuildWheel to build Python wheel files for a given package
// and python version.
// It uses the `build` package to build the wheel.
// You can specify extra packages to install using the `extraBuildPackages` argument.

package main

import "fmt"

type Wheelbuilder struct{}

var pythonPackages = map[string][]string{
	"3.10": {"python-3.10", "py3.10-pip", "build-base"},
	"3.11": {"python-3.11", "py3.11-pip", "build-base"},
	"3.12": {"python3", "py3-pip", "build-base"},
}

// Builds a python wheel using the specified python version and source directory
func (m *Wheelbuilder) BuildWheel(
	src *Directory,
	// +default="3.12"
	pythonVersion string,
	// +optional
	// +default=[]
	extraBuildPackages []string,
) (*Directory, error) {
	if pythonVersion == "" {
		pythonVersion = "3.12"
	}
	packages, ok := pythonPackages[pythonVersion]
	if !ok {
		return nil, fmt.Errorf("unsupported python version %s", pythonVersion)
	}

	packages = append(packages, extraBuildPackages...)

	return dag.Wolfi().Container(WolfiContainerOpts{
		Packages: packages,
	}).WithExec([]string{"pip", "install", "build"}).
		WithDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"python3", "-m", "build"}).
		Directory("dist/"), nil
}
