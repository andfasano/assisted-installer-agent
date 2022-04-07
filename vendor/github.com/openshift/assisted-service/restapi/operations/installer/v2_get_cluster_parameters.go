// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewV2GetClusterParams creates a new V2GetClusterParams object
// with the default values initialized.
func NewV2GetClusterParams() V2GetClusterParams {

	var (
		// initialize parameters with default values

		getUnregisteredClustersDefault = bool(false)
	)

	return V2GetClusterParams{
		GetUnregisteredClusters: &getUnregisteredClustersDefault,
	}
}

// V2GetClusterParams contains all the bound params for the v2 get cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2GetCluster
type V2GetClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The cluster to be retrieved.
	  Required: true
	  In: path
	*/
	ClusterID strfmt.UUID
	/*The software version of the discovery agent that is retrieving the cluster details.
	  In: header
	*/
	DiscoveryAgentVersion *string
	/*If true, do not include hosts.
	  In: query
	*/
	ExcludeHosts *bool
	/*Whether to return clusters that have been unregistered.
	  In: header
	  Default: false
	*/
	GetUnregisteredClusters *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2GetClusterParams() beforehand.
func (o *V2GetClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rClusterID, rhkClusterID, _ := route.Params.GetOK("cluster_id")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindDiscoveryAgentVersion(r.Header[http.CanonicalHeaderKey("discovery_agent_version")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qExcludeHosts, qhkExcludeHosts, _ := qs.GetOK("exclude-hosts")
	if err := o.bindExcludeHosts(qExcludeHosts, qhkExcludeHosts, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindGetUnregisteredClusters(r.Header[http.CanonicalHeaderKey("get_unregistered_clusters")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterID binds and validates parameter ClusterID from path.
func (o *V2GetClusterParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cluster_id", "path", "strfmt.UUID", raw)
	}
	o.ClusterID = *(value.(*strfmt.UUID))

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

// validateClusterID carries on validations for parameter ClusterID
func (o *V2GetClusterParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cluster_id", "path", "uuid", o.ClusterID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindDiscoveryAgentVersion binds and validates parameter DiscoveryAgentVersion from header.
func (o *V2GetClusterParams) bindDiscoveryAgentVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.DiscoveryAgentVersion = &raw

	return nil
}

// bindExcludeHosts binds and validates parameter ExcludeHosts from query.
func (o *V2GetClusterParams) bindExcludeHosts(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("exclude-hosts", "query", "bool", raw)
	}
	o.ExcludeHosts = &value

	return nil
}

// bindGetUnregisteredClusters binds and validates parameter GetUnregisteredClusters from header.
func (o *V2GetClusterParams) bindGetUnregisteredClusters(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewV2GetClusterParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("get_unregistered_clusters", "header", "bool", raw)
	}
	o.GetUnregisteredClusters = &value

	return nil
}
