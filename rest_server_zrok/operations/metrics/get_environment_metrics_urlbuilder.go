// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetEnvironmentMetricsURL generates an URL for the get environment metrics operation
type GetEnvironmentMetricsURL struct {
	EnvID string

	Duration *float64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetEnvironmentMetricsURL) WithBasePath(bp string) *GetEnvironmentMetricsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetEnvironmentMetricsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetEnvironmentMetricsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/metrics/environment/{envId}"

	envID := o.EnvID
	if envID != "" {
		_path = strings.Replace(_path, "{envId}", envID, -1)
	} else {
		return nil, errors.New("envId is required on GetEnvironmentMetricsURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var durationQ string
	if o.Duration != nil {
		durationQ = swag.FormatFloat64(*o.Duration)
	}
	if durationQ != "" {
		qs.Set("duration", durationQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetEnvironmentMetricsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetEnvironmentMetricsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetEnvironmentMetricsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetEnvironmentMetricsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetEnvironmentMetricsURL")
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
func (o *GetEnvironmentMetricsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
