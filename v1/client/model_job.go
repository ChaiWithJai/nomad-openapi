/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Job struct for Job
type Job struct {
	Affinities *[]Affinity `json:"Affinities,omitempty"`
	AllAtOnce *bool `json:"AllAtOnce,omitempty"`
	Constraints *[]Constraint `json:"Constraints,omitempty"`
	ConsulNamespace *string `json:"ConsulNamespace,omitempty"`
	ConsulToken *string `json:"ConsulToken,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	Datacenters *[]string `json:"Datacenters,omitempty"`
	Dispatched *bool `json:"Dispatched,omitempty"`
	ID *string `json:"ID,omitempty"`
	JobModifyIndex *int32 `json:"JobModifyIndex,omitempty"`
	Meta *Object `json:"Meta,omitempty"`
	Migrate *MigrateStrategy `json:"Migrate,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Multiregion *Multiregion `json:"Multiregion,omitempty"`
	Name *string `json:"Name,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	NomadTokenID *string `json:"NomadTokenID,omitempty"`
	ParameterizedJob *ParameterizedJobConfig `json:"ParameterizedJob,omitempty"`
	ParentID *string `json:"ParentID,omitempty"`
	Payload *string `json:"Payload,omitempty"`
	Periodic *PeriodicConfig `json:"Periodic,omitempty"`
	Priority *int32 `json:"Priority,omitempty"`
	Region *string `json:"Region,omitempty"`
	Reschedule *ReschedulePolicy `json:"Reschedule,omitempty"`
	Spreads *[]Spread `json:"Spreads,omitempty"`
	Stable *bool `json:"Stable,omitempty"`
	Status *string `json:"Status,omitempty"`
	StatusDescription *string `json:"StatusDescription,omitempty"`
	Stop *bool `json:"Stop,omitempty"`
	SubmitTime *int64 `json:"SubmitTime,omitempty"`
	TaskGroups *[]TaskGroup `json:"TaskGroups,omitempty"`
	Type *string `json:"Type,omitempty"`
	Update *UpdateStrategy `json:"Update,omitempty"`
	VaultNamespace *string `json:"VaultNamespace,omitempty"`
	VaultToken *string `json:"VaultToken,omitempty"`
	Version *int32 `json:"Version,omitempty"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob() *Job {
	this := Job{}
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetAffinities returns the Affinities field value if set, zero value otherwise.
func (o *Job) GetAffinities() []Affinity {
	if o == nil || o.Affinities == nil {
		var ret []Affinity
		return ret
	}
	return *o.Affinities
}

// GetAffinitiesOk returns a tuple with the Affinities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetAffinitiesOk() (*[]Affinity, bool) {
	if o == nil || o.Affinities == nil {
		return nil, false
	}
	return o.Affinities, true
}

// HasAffinities returns a boolean if a field has been set.
func (o *Job) HasAffinities() bool {
	if o != nil && o.Affinities != nil {
		return true
	}

	return false
}

// SetAffinities gets a reference to the given []Affinity and assigns it to the Affinities field.
func (o *Job) SetAffinities(v []Affinity) {
	o.Affinities = &v
}

// GetAllAtOnce returns the AllAtOnce field value if set, zero value otherwise.
func (o *Job) GetAllAtOnce() bool {
	if o == nil || o.AllAtOnce == nil {
		var ret bool
		return ret
	}
	return *o.AllAtOnce
}

// GetAllAtOnceOk returns a tuple with the AllAtOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetAllAtOnceOk() (*bool, bool) {
	if o == nil || o.AllAtOnce == nil {
		return nil, false
	}
	return o.AllAtOnce, true
}

// HasAllAtOnce returns a boolean if a field has been set.
func (o *Job) HasAllAtOnce() bool {
	if o != nil && o.AllAtOnce != nil {
		return true
	}

	return false
}

// SetAllAtOnce gets a reference to the given bool and assigns it to the AllAtOnce field.
func (o *Job) SetAllAtOnce(v bool) {
	o.AllAtOnce = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *Job) GetConstraints() []Constraint {
	if o == nil || o.Constraints == nil {
		var ret []Constraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetConstraintsOk() (*[]Constraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *Job) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *Job) SetConstraints(v []Constraint) {
	o.Constraints = &v
}

// GetConsulNamespace returns the ConsulNamespace field value if set, zero value otherwise.
func (o *Job) GetConsulNamespace() string {
	if o == nil || o.ConsulNamespace == nil {
		var ret string
		return ret
	}
	return *o.ConsulNamespace
}

// GetConsulNamespaceOk returns a tuple with the ConsulNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetConsulNamespaceOk() (*string, bool) {
	if o == nil || o.ConsulNamespace == nil {
		return nil, false
	}
	return o.ConsulNamespace, true
}

// HasConsulNamespace returns a boolean if a field has been set.
func (o *Job) HasConsulNamespace() bool {
	if o != nil && o.ConsulNamespace != nil {
		return true
	}

	return false
}

// SetConsulNamespace gets a reference to the given string and assigns it to the ConsulNamespace field.
func (o *Job) SetConsulNamespace(v string) {
	o.ConsulNamespace = &v
}

// GetConsulToken returns the ConsulToken field value if set, zero value otherwise.
func (o *Job) GetConsulToken() string {
	if o == nil || o.ConsulToken == nil {
		var ret string
		return ret
	}
	return *o.ConsulToken
}

// GetConsulTokenOk returns a tuple with the ConsulToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetConsulTokenOk() (*string, bool) {
	if o == nil || o.ConsulToken == nil {
		return nil, false
	}
	return o.ConsulToken, true
}

// HasConsulToken returns a boolean if a field has been set.
func (o *Job) HasConsulToken() bool {
	if o != nil && o.ConsulToken != nil {
		return true
	}

	return false
}

// SetConsulToken gets a reference to the given string and assigns it to the ConsulToken field.
func (o *Job) SetConsulToken(v string) {
	o.ConsulToken = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *Job) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *Job) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *Job) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetDatacenters returns the Datacenters field value if set, zero value otherwise.
func (o *Job) GetDatacenters() []string {
	if o == nil || o.Datacenters == nil {
		var ret []string
		return ret
	}
	return *o.Datacenters
}

// GetDatacentersOk returns a tuple with the Datacenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetDatacentersOk() (*[]string, bool) {
	if o == nil || o.Datacenters == nil {
		return nil, false
	}
	return o.Datacenters, true
}

// HasDatacenters returns a boolean if a field has been set.
func (o *Job) HasDatacenters() bool {
	if o != nil && o.Datacenters != nil {
		return true
	}

	return false
}

// SetDatacenters gets a reference to the given []string and assigns it to the Datacenters field.
func (o *Job) SetDatacenters(v []string) {
	o.Datacenters = &v
}

// GetDispatched returns the Dispatched field value if set, zero value otherwise.
func (o *Job) GetDispatched() bool {
	if o == nil || o.Dispatched == nil {
		var ret bool
		return ret
	}
	return *o.Dispatched
}

// GetDispatchedOk returns a tuple with the Dispatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetDispatchedOk() (*bool, bool) {
	if o == nil || o.Dispatched == nil {
		return nil, false
	}
	return o.Dispatched, true
}

// HasDispatched returns a boolean if a field has been set.
func (o *Job) HasDispatched() bool {
	if o != nil && o.Dispatched != nil {
		return true
	}

	return false
}

// SetDispatched gets a reference to the given bool and assigns it to the Dispatched field.
func (o *Job) SetDispatched(v bool) {
	o.Dispatched = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *Job) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *Job) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *Job) SetID(v string) {
	o.ID = &v
}

// GetJobModifyIndex returns the JobModifyIndex field value if set, zero value otherwise.
func (o *Job) GetJobModifyIndex() int32 {
	if o == nil || o.JobModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.JobModifyIndex
}

// GetJobModifyIndexOk returns a tuple with the JobModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetJobModifyIndexOk() (*int32, bool) {
	if o == nil || o.JobModifyIndex == nil {
		return nil, false
	}
	return o.JobModifyIndex, true
}

// HasJobModifyIndex returns a boolean if a field has been set.
func (o *Job) HasJobModifyIndex() bool {
	if o != nil && o.JobModifyIndex != nil {
		return true
	}

	return false
}

// SetJobModifyIndex gets a reference to the given int32 and assigns it to the JobModifyIndex field.
func (o *Job) SetJobModifyIndex(v int32) {
	o.JobModifyIndex = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Job) GetMeta() Object {
	if o == nil || o.Meta == nil {
		var ret Object
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetMetaOk() (*Object, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Job) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Object and assigns it to the Meta field.
func (o *Job) SetMeta(v Object) {
	o.Meta = &v
}

// GetMigrate returns the Migrate field value if set, zero value otherwise.
func (o *Job) GetMigrate() MigrateStrategy {
	if o == nil || o.Migrate == nil {
		var ret MigrateStrategy
		return ret
	}
	return *o.Migrate
}

// GetMigrateOk returns a tuple with the Migrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetMigrateOk() (*MigrateStrategy, bool) {
	if o == nil || o.Migrate == nil {
		return nil, false
	}
	return o.Migrate, true
}

// HasMigrate returns a boolean if a field has been set.
func (o *Job) HasMigrate() bool {
	if o != nil && o.Migrate != nil {
		return true
	}

	return false
}

// SetMigrate gets a reference to the given MigrateStrategy and assigns it to the Migrate field.
func (o *Job) SetMigrate(v MigrateStrategy) {
	o.Migrate = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *Job) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *Job) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *Job) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetMultiregion returns the Multiregion field value if set, zero value otherwise.
func (o *Job) GetMultiregion() Multiregion {
	if o == nil || o.Multiregion == nil {
		var ret Multiregion
		return ret
	}
	return *o.Multiregion
}

// GetMultiregionOk returns a tuple with the Multiregion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetMultiregionOk() (*Multiregion, bool) {
	if o == nil || o.Multiregion == nil {
		return nil, false
	}
	return o.Multiregion, true
}

// HasMultiregion returns a boolean if a field has been set.
func (o *Job) HasMultiregion() bool {
	if o != nil && o.Multiregion != nil {
		return true
	}

	return false
}

// SetMultiregion gets a reference to the given Multiregion and assigns it to the Multiregion field.
func (o *Job) SetMultiregion(v Multiregion) {
	o.Multiregion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Job) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Job) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Job) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Job) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Job) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Job) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNomadTokenID returns the NomadTokenID field value if set, zero value otherwise.
func (o *Job) GetNomadTokenID() string {
	if o == nil || o.NomadTokenID == nil {
		var ret string
		return ret
	}
	return *o.NomadTokenID
}

// GetNomadTokenIDOk returns a tuple with the NomadTokenID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetNomadTokenIDOk() (*string, bool) {
	if o == nil || o.NomadTokenID == nil {
		return nil, false
	}
	return o.NomadTokenID, true
}

// HasNomadTokenID returns a boolean if a field has been set.
func (o *Job) HasNomadTokenID() bool {
	if o != nil && o.NomadTokenID != nil {
		return true
	}

	return false
}

// SetNomadTokenID gets a reference to the given string and assigns it to the NomadTokenID field.
func (o *Job) SetNomadTokenID(v string) {
	o.NomadTokenID = &v
}

// GetParameterizedJob returns the ParameterizedJob field value if set, zero value otherwise.
func (o *Job) GetParameterizedJob() ParameterizedJobConfig {
	if o == nil || o.ParameterizedJob == nil {
		var ret ParameterizedJobConfig
		return ret
	}
	return *o.ParameterizedJob
}

// GetParameterizedJobOk returns a tuple with the ParameterizedJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetParameterizedJobOk() (*ParameterizedJobConfig, bool) {
	if o == nil || o.ParameterizedJob == nil {
		return nil, false
	}
	return o.ParameterizedJob, true
}

// HasParameterizedJob returns a boolean if a field has been set.
func (o *Job) HasParameterizedJob() bool {
	if o != nil && o.ParameterizedJob != nil {
		return true
	}

	return false
}

// SetParameterizedJob gets a reference to the given ParameterizedJobConfig and assigns it to the ParameterizedJob field.
func (o *Job) SetParameterizedJob(v ParameterizedJobConfig) {
	o.ParameterizedJob = &v
}

// GetParentID returns the ParentID field value if set, zero value otherwise.
func (o *Job) GetParentID() string {
	if o == nil || o.ParentID == nil {
		var ret string
		return ret
	}
	return *o.ParentID
}

// GetParentIDOk returns a tuple with the ParentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetParentIDOk() (*string, bool) {
	if o == nil || o.ParentID == nil {
		return nil, false
	}
	return o.ParentID, true
}

// HasParentID returns a boolean if a field has been set.
func (o *Job) HasParentID() bool {
	if o != nil && o.ParentID != nil {
		return true
	}

	return false
}

// SetParentID gets a reference to the given string and assigns it to the ParentID field.
func (o *Job) SetParentID(v string) {
	o.ParentID = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Job) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Job) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *Job) SetPayload(v string) {
	o.Payload = &v
}

// GetPeriodic returns the Periodic field value if set, zero value otherwise.
func (o *Job) GetPeriodic() PeriodicConfig {
	if o == nil || o.Periodic == nil {
		var ret PeriodicConfig
		return ret
	}
	return *o.Periodic
}

// GetPeriodicOk returns a tuple with the Periodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetPeriodicOk() (*PeriodicConfig, bool) {
	if o == nil || o.Periodic == nil {
		return nil, false
	}
	return o.Periodic, true
}

// HasPeriodic returns a boolean if a field has been set.
func (o *Job) HasPeriodic() bool {
	if o != nil && o.Periodic != nil {
		return true
	}

	return false
}

// SetPeriodic gets a reference to the given PeriodicConfig and assigns it to the Periodic field.
func (o *Job) SetPeriodic(v PeriodicConfig) {
	o.Periodic = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Job) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Job) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *Job) SetPriority(v int32) {
	o.Priority = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Job) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Job) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Job) SetRegion(v string) {
	o.Region = &v
}

// GetReschedule returns the Reschedule field value if set, zero value otherwise.
func (o *Job) GetReschedule() ReschedulePolicy {
	if o == nil || o.Reschedule == nil {
		var ret ReschedulePolicy
		return ret
	}
	return *o.Reschedule
}

// GetRescheduleOk returns a tuple with the Reschedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetRescheduleOk() (*ReschedulePolicy, bool) {
	if o == nil || o.Reschedule == nil {
		return nil, false
	}
	return o.Reschedule, true
}

// HasReschedule returns a boolean if a field has been set.
func (o *Job) HasReschedule() bool {
	if o != nil && o.Reschedule != nil {
		return true
	}

	return false
}

// SetReschedule gets a reference to the given ReschedulePolicy and assigns it to the Reschedule field.
func (o *Job) SetReschedule(v ReschedulePolicy) {
	o.Reschedule = &v
}

// GetSpreads returns the Spreads field value if set, zero value otherwise.
func (o *Job) GetSpreads() []Spread {
	if o == nil || o.Spreads == nil {
		var ret []Spread
		return ret
	}
	return *o.Spreads
}

// GetSpreadsOk returns a tuple with the Spreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetSpreadsOk() (*[]Spread, bool) {
	if o == nil || o.Spreads == nil {
		return nil, false
	}
	return o.Spreads, true
}

// HasSpreads returns a boolean if a field has been set.
func (o *Job) HasSpreads() bool {
	if o != nil && o.Spreads != nil {
		return true
	}

	return false
}

// SetSpreads gets a reference to the given []Spread and assigns it to the Spreads field.
func (o *Job) SetSpreads(v []Spread) {
	o.Spreads = &v
}

// GetStable returns the Stable field value if set, zero value otherwise.
func (o *Job) GetStable() bool {
	if o == nil || o.Stable == nil {
		var ret bool
		return ret
	}
	return *o.Stable
}

// GetStableOk returns a tuple with the Stable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetStableOk() (*bool, bool) {
	if o == nil || o.Stable == nil {
		return nil, false
	}
	return o.Stable, true
}

// HasStable returns a boolean if a field has been set.
func (o *Job) HasStable() bool {
	if o != nil && o.Stable != nil {
		return true
	}

	return false
}

// SetStable gets a reference to the given bool and assigns it to the Stable field.
func (o *Job) SetStable(v bool) {
	o.Stable = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Job) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Job) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Job) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *Job) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *Job) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *Job) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetStop returns the Stop field value if set, zero value otherwise.
func (o *Job) GetStop() bool {
	if o == nil || o.Stop == nil {
		var ret bool
		return ret
	}
	return *o.Stop
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetStopOk() (*bool, bool) {
	if o == nil || o.Stop == nil {
		return nil, false
	}
	return o.Stop, true
}

// HasStop returns a boolean if a field has been set.
func (o *Job) HasStop() bool {
	if o != nil && o.Stop != nil {
		return true
	}

	return false
}

// SetStop gets a reference to the given bool and assigns it to the Stop field.
func (o *Job) SetStop(v bool) {
	o.Stop = &v
}

// GetSubmitTime returns the SubmitTime field value if set, zero value otherwise.
func (o *Job) GetSubmitTime() int64 {
	if o == nil || o.SubmitTime == nil {
		var ret int64
		return ret
	}
	return *o.SubmitTime
}

// GetSubmitTimeOk returns a tuple with the SubmitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetSubmitTimeOk() (*int64, bool) {
	if o == nil || o.SubmitTime == nil {
		return nil, false
	}
	return o.SubmitTime, true
}

// HasSubmitTime returns a boolean if a field has been set.
func (o *Job) HasSubmitTime() bool {
	if o != nil && o.SubmitTime != nil {
		return true
	}

	return false
}

// SetSubmitTime gets a reference to the given int64 and assigns it to the SubmitTime field.
func (o *Job) SetSubmitTime(v int64) {
	o.SubmitTime = &v
}

// GetTaskGroups returns the TaskGroups field value if set, zero value otherwise.
func (o *Job) GetTaskGroups() []TaskGroup {
	if o == nil || o.TaskGroups == nil {
		var ret []TaskGroup
		return ret
	}
	return *o.TaskGroups
}

// GetTaskGroupsOk returns a tuple with the TaskGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTaskGroupsOk() (*[]TaskGroup, bool) {
	if o == nil || o.TaskGroups == nil {
		return nil, false
	}
	return o.TaskGroups, true
}

// HasTaskGroups returns a boolean if a field has been set.
func (o *Job) HasTaskGroups() bool {
	if o != nil && o.TaskGroups != nil {
		return true
	}

	return false
}

// SetTaskGroups gets a reference to the given []TaskGroup and assigns it to the TaskGroups field.
func (o *Job) SetTaskGroups(v []TaskGroup) {
	o.TaskGroups = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Job) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Job) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Job) SetType(v string) {
	o.Type = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *Job) GetUpdate() UpdateStrategy {
	if o == nil || o.Update == nil {
		var ret UpdateStrategy
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetUpdateOk() (*UpdateStrategy, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *Job) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given UpdateStrategy and assigns it to the Update field.
func (o *Job) SetUpdate(v UpdateStrategy) {
	o.Update = &v
}

// GetVaultNamespace returns the VaultNamespace field value if set, zero value otherwise.
func (o *Job) GetVaultNamespace() string {
	if o == nil || o.VaultNamespace == nil {
		var ret string
		return ret
	}
	return *o.VaultNamespace
}

// GetVaultNamespaceOk returns a tuple with the VaultNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetVaultNamespaceOk() (*string, bool) {
	if o == nil || o.VaultNamespace == nil {
		return nil, false
	}
	return o.VaultNamespace, true
}

// HasVaultNamespace returns a boolean if a field has been set.
func (o *Job) HasVaultNamespace() bool {
	if o != nil && o.VaultNamespace != nil {
		return true
	}

	return false
}

// SetVaultNamespace gets a reference to the given string and assigns it to the VaultNamespace field.
func (o *Job) SetVaultNamespace(v string) {
	o.VaultNamespace = &v
}

// GetVaultToken returns the VaultToken field value if set, zero value otherwise.
func (o *Job) GetVaultToken() string {
	if o == nil || o.VaultToken == nil {
		var ret string
		return ret
	}
	return *o.VaultToken
}

// GetVaultTokenOk returns a tuple with the VaultToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetVaultTokenOk() (*string, bool) {
	if o == nil || o.VaultToken == nil {
		return nil, false
	}
	return o.VaultToken, true
}

// HasVaultToken returns a boolean if a field has been set.
func (o *Job) HasVaultToken() bool {
	if o != nil && o.VaultToken != nil {
		return true
	}

	return false
}

// SetVaultToken gets a reference to the given string and assigns it to the VaultToken field.
func (o *Job) SetVaultToken(v string) {
	o.VaultToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Job) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Job) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Job) SetVersion(v int32) {
	o.Version = &v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Affinities != nil {
		toSerialize["Affinities"] = o.Affinities
	}
	if o.AllAtOnce != nil {
		toSerialize["AllAtOnce"] = o.AllAtOnce
	}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.ConsulNamespace != nil {
		toSerialize["ConsulNamespace"] = o.ConsulNamespace
	}
	if o.ConsulToken != nil {
		toSerialize["ConsulToken"] = o.ConsulToken
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.Datacenters != nil {
		toSerialize["Datacenters"] = o.Datacenters
	}
	if o.Dispatched != nil {
		toSerialize["Dispatched"] = o.Dispatched
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.JobModifyIndex != nil {
		toSerialize["JobModifyIndex"] = o.JobModifyIndex
	}
	if o.Meta != nil {
		toSerialize["Meta"] = o.Meta
	}
	if o.Migrate != nil {
		toSerialize["Migrate"] = o.Migrate
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Multiregion != nil {
		toSerialize["Multiregion"] = o.Multiregion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.NomadTokenID != nil {
		toSerialize["NomadTokenID"] = o.NomadTokenID
	}
	if o.ParameterizedJob != nil {
		toSerialize["ParameterizedJob"] = o.ParameterizedJob
	}
	if o.ParentID != nil {
		toSerialize["ParentID"] = o.ParentID
	}
	if o.Payload != nil {
		toSerialize["Payload"] = o.Payload
	}
	if o.Periodic != nil {
		toSerialize["Periodic"] = o.Periodic
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}
	if o.Reschedule != nil {
		toSerialize["Reschedule"] = o.Reschedule
	}
	if o.Spreads != nil {
		toSerialize["Spreads"] = o.Spreads
	}
	if o.Stable != nil {
		toSerialize["Stable"] = o.Stable
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusDescription != nil {
		toSerialize["StatusDescription"] = o.StatusDescription
	}
	if o.Stop != nil {
		toSerialize["Stop"] = o.Stop
	}
	if o.SubmitTime != nil {
		toSerialize["SubmitTime"] = o.SubmitTime
	}
	if o.TaskGroups != nil {
		toSerialize["TaskGroups"] = o.TaskGroups
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Update != nil {
		toSerialize["Update"] = o.Update
	}
	if o.VaultNamespace != nil {
		toSerialize["VaultNamespace"] = o.VaultNamespace
	}
	if o.VaultToken != nil {
		toSerialize["VaultToken"] = o.VaultToken
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


