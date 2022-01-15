// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RetrieveTransactionURL generates an URL for the retrieve transaction operation
type RetrieveTransactionURL struct {
	TransactionID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveTransactionURL) WithBasePath(bp string) *RetrieveTransactionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveTransactionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RetrieveTransactionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/transaction/{transaction_id}"

	transactionID := o.TransactionID
	if transactionID != "" {
		_path = strings.Replace(_path, "{transaction_id}", transactionID, -1)
	} else {
		return nil, errors.New("transactionId is required on RetrieveTransactionURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RetrieveTransactionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RetrieveTransactionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RetrieveTransactionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RetrieveTransactionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RetrieveTransactionURL")
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
func (o *RetrieveTransactionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
