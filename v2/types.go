package v2

// This file contains the user-facing types used for the Open Service Broker
// client.

// Service is an available service listed in a broker's catalog.
type Service struct {
	// ID is a globally unique ID that identifies the service.
	ID string `json:"id"`
	// Name is the service's display name.
	Name string `json:"name"`
	// Description is a brief description of the service, suitable for
	// printing by a CLI.
	Description string `json:"description"`
	// A list of 'tags' describing different classification referents or
	// attributes of the service.  CF-specific.
	Tags []string `json:"tags,omitempty"`
	// A list of permissions the user must give instances of this service.
	// CF-specific.  Current valid values are:
	//
	// - syslog_drain
	// - route_forwarding
	// - volume_mount
	//
	// See the doc on service permissions for more information.
	// TODO: write doc
	Requires []string `json:"requires,omitempty"`
	// Bindable represents whether a service is bindable.  May be overriden on
	// a per-plan basis by the Plan.Bindable field.
	Bindable bool `json:"bindable"`
	// PlanUpdatable represents whether instances of this service may be
	// updated to a different plan.  The serialized form 'plan_updateable' is
	// a mistake that has become written into the API for backward
	// compatibility reasons and is intentional.  Optional; defaults to false.
	PlanUpdatable *bool `json:"plan_updateable,omitempty"`
	// Plans is the list of the Plans for a service.  Plans represent
	// different tiers.
	Plans []Plan `json:"plans"`
	// DashboardClient holds information about the OAuth SSO for the service's
	// dashboard.  Optional.
	DashboardClient *DashboardClient `json:"dashboard_client,omitempty"`
	// Metadata is a blob of information about the plan, meant to be user-
	// facing content and display instructions.  Metadata may contain
	// platform-conventional values.  Optional.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// DashboardClient contains information about the OAuth SSO
// flow for a Service's dashboard.
type DashboardClient struct {
	// ID is the ID to use for the dashboard SSO OAuth client for this
	// service.
	ID string `json:"id"`
	// Secret is a secret for the dashboard SSO OAuth client.
	Secret string `json:"secret"`
	// RedirectURI is the redirect URI that should be used to obtain an OAuth
	// token.
	RedirectURI string `json:"redirect_uri"`
}

// Plan is a plan (or tier) within a service offering.
type Plan struct {
	// ID is a globally unique ID that identifies the plan.
	ID string `json:"id"`
	// Name is the plan's display name.
	Name string `json:"name"`
	// Description is a brief description of the plan, suitable for
	// printing by a CLI.
	Description string `json:"description"`
	// Free indicates whether the plan is available without charge.  Optional;
	// defaults to true.
	Free *bool `json:"free,omitempty"`
	// Bindable indicates whether the plan is bindable and overrides the value
	// of the Service.Bindable field if set.  Optional, defaults to unset.
	Bindable *bool `json:"bindable,omitempty"`
	// Metadata is a blob of information about the plan, meant to be user-
	// facing content and display instructions.  Metadata may contain
	// platform-conventional values.  Optional.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// CatalogResponse is sent as the response to catalog requests.
type CatalogResponse struct {
	Services []Service `json:"services"`
}

// ProvisionRequest encompasses the request and body parameters
type ProvisionRequest struct {
	// InstanceID is the ID of the new instance to provision.  The Open
	// Service Broker API specification recommends using a GUID for this
	// field.
	InstanceID string `json:"instance_id"`
	// AcceptsIncomple indicates whether the client can accept asynchronous
	// provisioning.  If the broker does not support asynchronous provisioning
	// of a service, it will reject a request with this field set to true.
	AcceptsIncomplete bool `json:"accepts_incomplete"`
	// ServiceID is the ID of the service to provision a new instance of.
	ServiceID string `json:"service_id"`
	// PlanID is the ID of the plan to use for the new instance.
	PlanID string `json:"plan_id"`
	// OrganizationGUID is the platform GUID for the organization under which
	// the service is to be provisioned.  CF-specific.
	OrganizationGUID string `json:"organization_guid"`
	// SpaceGUID is the identifier for the project space within the platform
	// organization.  CF-specific.
	SpaceGUID string `json:"space_guid"`
	// Parameters is a set of configuration options for the service instance.
	// Optional.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// ProvisionResponse is sent in response to a provision call
type ProvisionResponse struct {
	// DashboardURL is the URL of a web-based management user interface for
	// the service instance.
	DashboardURL *string `json:"dashboard_url,omitempty"`
	// OperationKey is an extra identifier supplied by the broker to identify
	// asynchronous operations.
	OperationKey *OperationKey `json:"operationKey,omitempty"`
}

// OperationKeys may be returned by the broker in order to provide extra
// identifiers for asynchronous operations.
type OperationKey string

// UpdateInstanceRequest is the user-facing object that represents a request
// to update an instance's plan or parameters.
type UpdateInstanceRequest struct {
	// InstanceID is the ID of the instance to update.
	InstanceID string `json:"instance_id"`
	// AcceptsIncomple indicates whether the client can accept asynchronous
	// provisioning.  If the broker does not support asynchronous provisioning
	// of a service, it will reject a request with this field set to true.
	AcceptsIncomplete bool `json:"accepts_incomplete"`
	// ServiceID is the ID of the service the instance is provisioned from.
	ServiceID string `json:"service_id"`
	// PlanID is the ID the plan to update the instance to.  The service must
	// support plan updates.  If unspecified, indicates that the client does
	// not wish to update the plan of the instance.
	PlanID *string `json:"plan_id,omitempty"`
	// Parameters is a set of configuration options for the instance.  If
	// unset, indicates that the client does not wish to update the parameters
	// for an instance.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The OSB API also has a field called `previous_values` that contains:
	// OrgID
	// SpaceID
	// ServiceID
	// PlanID
	//
	// ...but those fields seem to be a relic of some API design mistakes in
	// the past.  As such, this client library does not currently support
	// them.  I will happily change this if someone can present a specific
	// example of a broker that requires these fields to be sent.
}

// UpdateInstanceResponse represents a broker's response to an update instance
// request.
type UpdateInstanceResponse struct {
	// OperationKey is an extra identifier supplied by the broker to identify
	// asynchronous operations.
	OperationKey *OperationKey `json:"operationKey,omitempty"`
}

// Deprovision request represents a request to deprovision an instance of a
// service.
type DeprovisionRequest struct {
	// InstanceID is the ID of the instance to deprovision.
	InstanceID string `json:"instance_id"`
	// AcceptsIncomple indicates whether the client can accept asynchronous
	// provisioning.  If the broker does not support asynchronous provisioning
	// of a service, it will reject a request with this field set to true.
	AcceptsIncomplete bool   `json:"accepts_incomplete"`
	ServiceID         string `json:service_id"`
	PlanID            string `json:plan_id"`
}

// DeprovisionResponse represents a broker's response to a deprovision request.
type DeprovisionResponse struct {
	// OperationKey is an extra identifier supplied by the broker to identify
	// asynchronous operations.
	OperationKey *OperationKey `json:"operationKey,omitempty"`
}

// LastOperationRequest represents a request to a broker to give the state of
// the action it is completing asynchronously.
type LastOperationRequest struct {
	// InstanceID is the instance of the service to query the last operation
	// for.
	InstanceID string `json:"instance_id"`
	// ServiceID is the ID of the service the instance is provisioned from.
	// Optional, but recommended.
	ServiceID *string `json:"service_id,omitempty"`
	// PlanID is the ID of the plan the instance is provisioned from.
	// Optional, but recommended.
	PlanID *string `json:"plan_id,omitempty"`
	// OperationKey is the operation key provided by the broker in the
	// response to the initial request.  Optional, but must be sent if
	// supplied in the response to the original request.
	OperationKey *OperationKey `json:"operation,omitempty"`
}

// LastOperationResponse represents the broker response with the state of a
// discrete action that the broker is completing asynchronously.
type LastOperationResponse struct {
	// State is the state of the queried operation.
	State LastOperationState `json:"state"`
	// Description is a message from the broker describing the current state
	// of the operation.
	Description *string `json:"description,omitempty"`
}

// LastOperationState is a typedef representing the state of an ongoing
// operation for an instance.
type LastOperationState string

// Defines the possible states of an asynchronous request to a broker.
const (
	StateInProgress LastOperationState = "in progress"
	StateSucceeded  LastOperationState = "succeeded"
	StateFailed     LastOperationState = "failed"
)

// BindRequest represents a request to create a new binding to an instance of
// a service.
type BindRequest struct {
	// InstanceID is the ID of the instance to bind to.
	InstanceID string `json:"instance_id"`
	// BindingID is the ID of the new binding to create.  The Open Service
	// Broker API specification recommends using a GUID for this field.
	BindingID string `json:"binding_id"`

	// ServiceID is the ID of the service the instance was provisioned from.
	ServiceID string `json:"service_id"`
	// PlanID is the ID of the plan the instance was provisioned from.
	PlanID string `json:"plan_id"`
	// Deprecated; use bind_resource.app_guid to send this value instead.
	AppGUID *string `json:"app_guid,omitempty"`
	// BindResource holds extra information about a binding.  Optional, but
	// it's complicated. TODO: clarify
	BindResource *BindResource `json:"bind_resource,omitempty"`
	// Parameters is configuration parameters for the binding.  Optional.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// BindResource contains data for platform resources associated with a
// binding.
type BindResource struct {
	AppGUID *string `json:"appGuid,omitempty"`
	Route   *string `json:"route,omitempty"`
}

// BindResponse represents a broker's response to an UnbindRequest.
type BindResponse struct {
	// Credentials is a free-form hash of credentials that can be used by
	// applications or users to access the service.
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	// SyslogDrainURl is a URL to which logs must be streamed.  CF-specific.
	// May only be supplied by a service that declares a requirement for the
	// 'syslog_drain' permission.
	SyslogDrainURL *string `json:"syslog_drain_url,omitempty"`
	// RouteServiceURL is a URL to which the platform must proxy requests to
	// the application the binding is for.  CF-specific.  May only be supplied
	// by a service that declares a requirement for the 'route_service'
	// permission.
	RouteServiceURL *string `json:"route_service_url,omitempty"`
	// VolumeMounts is an array of configuration string for mounting volumes.
	// CF-specific.  May only be supplied by a service that declares a
	// requirement for the 'volume_mount' permission.
	VolumeMounts []interface{} `json:"volume_mounts,omitempty"`
}

// UnbindRequest represents a request to unbind a particular binding.
type UnbindRequest struct {
	// BindingID is the ID of the binding to delete.
	BindingID string `json:"binding_id"`
}

// UnbindResponse represents a broker's response to an UnbindRequest.
type UnbindResponse struct {
	// Currently, unbind responses have no fields.
}
