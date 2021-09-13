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

// TaskArtifact struct for TaskArtifact
type TaskArtifact struct {
	GetterHeaders *Object `json:"GetterHeaders,omitempty"`
	GetterMode *string `json:"GetterMode,omitempty"`
	GetterOptions *Object `json:"GetterOptions,omitempty"`
	GetterSource *string `json:"GetterSource,omitempty"`
	RelativeDest *string `json:"RelativeDest,omitempty"`
}

// NewTaskArtifact instantiates a new TaskArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskArtifact() *TaskArtifact {
	this := TaskArtifact{}
	return &this
}

// NewTaskArtifactWithDefaults instantiates a new TaskArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskArtifactWithDefaults() *TaskArtifact {
	this := TaskArtifact{}
	return &this
}

// GetGetterHeaders returns the GetterHeaders field value if set, zero value otherwise.
func (o *TaskArtifact) GetGetterHeaders() Object {
	if o == nil || o.GetterHeaders == nil {
		var ret Object
		return ret
	}
	return *o.GetterHeaders
}

// GetGetterHeadersOk returns a tuple with the GetterHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskArtifact) GetGetterHeadersOk() (*Object, bool) {
	if o == nil || o.GetterHeaders == nil {
		return nil, false
	}
	return o.GetterHeaders, true
}

// HasGetterHeaders returns a boolean if a field has been set.
func (o *TaskArtifact) HasGetterHeaders() bool {
	if o != nil && o.GetterHeaders != nil {
		return true
	}

	return false
}

// SetGetterHeaders gets a reference to the given Object and assigns it to the GetterHeaders field.
func (o *TaskArtifact) SetGetterHeaders(v Object) {
	o.GetterHeaders = &v
}

// GetGetterMode returns the GetterMode field value if set, zero value otherwise.
func (o *TaskArtifact) GetGetterMode() string {
	if o == nil || o.GetterMode == nil {
		var ret string
		return ret
	}
	return *o.GetterMode
}

// GetGetterModeOk returns a tuple with the GetterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskArtifact) GetGetterModeOk() (*string, bool) {
	if o == nil || o.GetterMode == nil {
		return nil, false
	}
	return o.GetterMode, true
}

// HasGetterMode returns a boolean if a field has been set.
func (o *TaskArtifact) HasGetterMode() bool {
	if o != nil && o.GetterMode != nil {
		return true
	}

	return false
}

// SetGetterMode gets a reference to the given string and assigns it to the GetterMode field.
func (o *TaskArtifact) SetGetterMode(v string) {
	o.GetterMode = &v
}

// GetGetterOptions returns the GetterOptions field value if set, zero value otherwise.
func (o *TaskArtifact) GetGetterOptions() Object {
	if o == nil || o.GetterOptions == nil {
		var ret Object
		return ret
	}
	return *o.GetterOptions
}

// GetGetterOptionsOk returns a tuple with the GetterOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskArtifact) GetGetterOptionsOk() (*Object, bool) {
	if o == nil || o.GetterOptions == nil {
		return nil, false
	}
	return o.GetterOptions, true
}

// HasGetterOptions returns a boolean if a field has been set.
func (o *TaskArtifact) HasGetterOptions() bool {
	if o != nil && o.GetterOptions != nil {
		return true
	}

	return false
}

// SetGetterOptions gets a reference to the given Object and assigns it to the GetterOptions field.
func (o *TaskArtifact) SetGetterOptions(v Object) {
	o.GetterOptions = &v
}

// GetGetterSource returns the GetterSource field value if set, zero value otherwise.
func (o *TaskArtifact) GetGetterSource() string {
	if o == nil || o.GetterSource == nil {
		var ret string
		return ret
	}
	return *o.GetterSource
}

// GetGetterSourceOk returns a tuple with the GetterSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskArtifact) GetGetterSourceOk() (*string, bool) {
	if o == nil || o.GetterSource == nil {
		return nil, false
	}
	return o.GetterSource, true
}

// HasGetterSource returns a boolean if a field has been set.
func (o *TaskArtifact) HasGetterSource() bool {
	if o != nil && o.GetterSource != nil {
		return true
	}

	return false
}

// SetGetterSource gets a reference to the given string and assigns it to the GetterSource field.
func (o *TaskArtifact) SetGetterSource(v string) {
	o.GetterSource = &v
}

// GetRelativeDest returns the RelativeDest field value if set, zero value otherwise.
func (o *TaskArtifact) GetRelativeDest() string {
	if o == nil || o.RelativeDest == nil {
		var ret string
		return ret
	}
	return *o.RelativeDest
}

// GetRelativeDestOk returns a tuple with the RelativeDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskArtifact) GetRelativeDestOk() (*string, bool) {
	if o == nil || o.RelativeDest == nil {
		return nil, false
	}
	return o.RelativeDest, true
}

// HasRelativeDest returns a boolean if a field has been set.
func (o *TaskArtifact) HasRelativeDest() bool {
	if o != nil && o.RelativeDest != nil {
		return true
	}

	return false
}

// SetRelativeDest gets a reference to the given string and assigns it to the RelativeDest field.
func (o *TaskArtifact) SetRelativeDest(v string) {
	o.RelativeDest = &v
}

func (o TaskArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GetterHeaders != nil {
		toSerialize["GetterHeaders"] = o.GetterHeaders
	}
	if o.GetterMode != nil {
		toSerialize["GetterMode"] = o.GetterMode
	}
	if o.GetterOptions != nil {
		toSerialize["GetterOptions"] = o.GetterOptions
	}
	if o.GetterSource != nil {
		toSerialize["GetterSource"] = o.GetterSource
	}
	if o.RelativeDest != nil {
		toSerialize["RelativeDest"] = o.RelativeDest
	}
	return json.Marshal(toSerialize)
}

type NullableTaskArtifact struct {
	value *TaskArtifact
	isSet bool
}

func (v NullableTaskArtifact) Get() *TaskArtifact {
	return v.value
}

func (v *NullableTaskArtifact) Set(val *TaskArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskArtifact(val *TaskArtifact) *NullableTaskArtifact {
	return &NullableTaskArtifact{value: val, isSet: true}
}

func (v NullableTaskArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


