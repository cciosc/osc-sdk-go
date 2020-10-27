/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// FiltersVirtualGateway One or more filters.
type FiltersVirtualGateway struct {
	// The types of the virtual gateways (only `ipsec.1` is supported).
	ConnectionTypes []string `json:"ConnectionTypes,omitempty"`
	// The IDs of the Nets the virtual gateways are attached to.
	LinkNetIds []string `json:"LinkNetIds,omitempty"`
	// The current states of the attachments between the virtual gateways and the Nets (`attaching` \\| `attached` \\| `detaching` \\| `detached`).
	LinkStates []string `json:"LinkStates,omitempty"`
	// The states of the virtual gateways (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the virtual gateways.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the virtual gateways.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the virtual gateways, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags []string `json:"Tags,omitempty"`
	// The IDs of the virtual gateways.
	VirtualGatewayIds []string `json:"VirtualGatewayIds,omitempty"`
}
