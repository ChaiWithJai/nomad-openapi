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

// ScalingRequest struct for ScalingRequest
type ScalingRequest struct {
	Count *int64 `json:"Count,omitempty"`
	Error *bool `json:"Error,omitempty"`
	Message *string `json:"Message,omitempty"`
	Meta *map[string]interface{} `json:"Meta,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	PolicyOverride *bool `json:"PolicyOverride,omitempty"`
	Region *string `json:"Region,omitempty"`
	SecretID *string `json:"SecretID,omitempty"`
	Target *Object `json:"Target,omitempty"`
}

// NewScalingRequest instantiates a new ScalingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalingRequest() *ScalingRequest {
	this := ScalingRequest{}
	return &this
}

// NewScalingRequestWithDefaults instantiates a new ScalingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalingRequestWithDefaults() *ScalingRequest {
	this := ScalingRequest{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ScalingRequest) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ScalingRequest) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ScalingRequest) SetCount(v int64) {
	o.Count = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ScalingRequest) GetError() bool {
	if o == nil || o.Error == nil {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetErrorOk() (*bool, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ScalingRequest) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *ScalingRequest) SetError(v bool) {
	o.Error = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ScalingRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ScalingRequest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ScalingRequest) SetMessage(v string) {
	o.Message = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScalingRequest) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScalingRequest) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ScalingRequest) SetMeta(v map[string]interface{}) {
	o.Meta = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ScalingRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ScalingRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ScalingRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPolicyOverride returns the PolicyOverride field value if set, zero value otherwise.
func (o *ScalingRequest) GetPolicyOverride() bool {
	if o == nil || o.PolicyOverride == nil {
		var ret bool
		return ret
	}
	return *o.PolicyOverride
}

// GetPolicyOverrideOk returns a tuple with the PolicyOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetPolicyOverrideOk() (*bool, bool) {
	if o == nil || o.PolicyOverride == nil {
		return nil, false
	}
	return o.PolicyOverride, true
}

// HasPolicyOverride returns a boolean if a field has been set.
func (o *ScalingRequest) HasPolicyOverride() bool {
	if o != nil && o.PolicyOverride != nil {
		return true
	}

	return false
}

// SetPolicyOverride gets a reference to the given bool and assigns it to the PolicyOverride field.
func (o *ScalingRequest) SetPolicyOverride(v bool) {
	o.PolicyOverride = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ScalingRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ScalingRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ScalingRequest) SetRegion(v string) {
	o.Region = &v
}

// GetSecretID returns the SecretID field value if set, zero value otherwise.
func (o *ScalingRequest) GetSecretID() string {
	if o == nil || o.SecretID == nil {
		var ret string
		return ret
	}
	return *o.SecretID
}

// GetSecretIDOk returns a tuple with the SecretID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetSecretIDOk() (*string, bool) {
	if o == nil || o.SecretID == nil {
		return nil, false
	}
	return o.SecretID, true
}

// HasSecretID returns a boolean if a field has been set.
func (o *ScalingRequest) HasSecretID() bool {
	if o != nil && o.SecretID != nil {
		return true
	}

	return false
}

// SetSecretID gets a reference to the given string and assigns it to the SecretID field.
func (o *ScalingRequest) SetSecretID(v string) {
	o.SecretID = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ScalingRequest) GetTarget() Object {
	if o == nil || o.Target == nil {
		var ret Object
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingRequest) GetTargetOk() (*Object, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ScalingRequest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given Object and assigns it to the Target field.
func (o *ScalingRequest) SetTarget(v Object) {
	o.Target = &v
}

func (o ScalingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Error != nil {
		toSerialize["Error"] = o.Error
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Meta != nil {
		toSerialize["Meta"] = o.Meta
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.PolicyOverride != nil {
		toSerialize["PolicyOverride"] = o.PolicyOverride
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}
	if o.SecretID != nil {
		toSerialize["SecretID"] = o.SecretID
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableScalingRequest struct {
	value *ScalingRequest
	isSet bool
}

func (v NullableScalingRequest) Get() *ScalingRequest {
	return v.value
}

func (v *NullableScalingRequest) Set(val *ScalingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScalingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScalingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalingRequest(val *ScalingRequest) *NullableScalingRequest {
	return &NullableScalingRequest{value: val, isSet: true}
}

func (v NullableScalingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


