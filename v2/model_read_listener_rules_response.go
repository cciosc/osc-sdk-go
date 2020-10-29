/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadListenerRulesResponse struct for ReadListenerRulesResponse
type ReadListenerRulesResponse struct {
	// The list of the rules to describe.
	ListenerRules *[]ListenerRule `json:"ListenerRules,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadListenerRulesResponse instantiates a new ReadListenerRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadListenerRulesResponse() *ReadListenerRulesResponse {
	this := ReadListenerRulesResponse{}
	return &this
}

// NewReadListenerRulesResponseWithDefaults instantiates a new ReadListenerRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadListenerRulesResponseWithDefaults() *ReadListenerRulesResponse {
	this := ReadListenerRulesResponse{}
	return &this
}

// GetListenerRules returns the ListenerRules field value if set, zero value otherwise.
func (o *ReadListenerRulesResponse) GetListenerRules() []ListenerRule {
	if o == nil || o.ListenerRules == nil {
		var ret []ListenerRule
		return ret
	}
	return *o.ListenerRules
}

// GetListenerRulesOk returns a tuple with the ListenerRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadListenerRulesResponse) GetListenerRulesOk() (*[]ListenerRule, bool) {
	if o == nil || o.ListenerRules == nil {
		return nil, false
	}
	return o.ListenerRules, true
}

// HasListenerRules returns a boolean if a field has been set.
func (o *ReadListenerRulesResponse) HasListenerRules() bool {
	if o != nil && o.ListenerRules != nil {
		return true
	}

	return false
}

// SetListenerRules gets a reference to the given []ListenerRule and assigns it to the ListenerRules field.
func (o *ReadListenerRulesResponse) SetListenerRules(v []ListenerRule) {
	o.ListenerRules = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadListenerRulesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadListenerRulesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadListenerRulesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadListenerRulesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadListenerRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListenerRules != nil {
		toSerialize["ListenerRules"] = o.ListenerRules
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadListenerRulesResponse struct {
	value *ReadListenerRulesResponse
	isSet bool
}

func (v NullableReadListenerRulesResponse) Get() *ReadListenerRulesResponse {
	return v.value
}

func (v *NullableReadListenerRulesResponse) Set(val *ReadListenerRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadListenerRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadListenerRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadListenerRulesResponse(val *ReadListenerRulesResponse) *NullableReadListenerRulesResponse {
	return &NullableReadListenerRulesResponse{value: val, isSet: true}
}

func (v NullableReadListenerRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadListenerRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

