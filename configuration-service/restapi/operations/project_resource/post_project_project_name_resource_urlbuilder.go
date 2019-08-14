// Code generated by go-swagger; DO NOT EDIT.

package project_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PostProjectProjectNameResourceURL generates an URL for the post project project name resource operation
type PostProjectProjectNameResourceURL struct {
	ProjectName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostProjectProjectNameResourceURL) WithBasePath(bp string) *PostProjectProjectNameResourceURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostProjectProjectNameResourceURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostProjectProjectNameResourceURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/project/{projectName}/resource"

	projectName := o.ProjectName
	if projectName != "" {
		_path = strings.Replace(_path, "{projectName}", projectName, -1)
	} else {
		return nil, errors.New("projectName is required on PostProjectProjectNameResourceURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostProjectProjectNameResourceURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostProjectProjectNameResourceURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostProjectProjectNameResourceURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostProjectProjectNameResourceURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostProjectProjectNameResourceURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PostProjectProjectNameResourceURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
