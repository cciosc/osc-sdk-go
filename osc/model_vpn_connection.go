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
// VpnConnection Information about a VPN connection.
type VpnConnection struct {
	// The configuration to apply to the client gateway to establish the VPN connection, in XML format.
	ClientGatewayConfiguration string `json:"ClientGatewayConfiguration,omitempty"`
	// The ID of the client gateway used on the client end of the connection.
	ClientGatewayId string `json:"ClientGatewayId,omitempty"`
	// The type of VPN connection (always `ipsec.1`).
	ConnectionType string `json:"ConnectionType,omitempty"`
	// Information about one or more static routes associated with the VPN connection, if any.
	Routes []RouteLight `json:"Routes,omitempty"`
	// The state of the VPN connection (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State string `json:"State,omitempty"`
	// If `false`, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If `true`, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute).
	StaticRoutesOnly bool `json:"StaticRoutesOnly,omitempty"`
	// One or more tags associated with the VPN connection.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The ID of the virtual gateway used on the 3DS OUTSCALE end of the connection.
	VirtualGatewayId string `json:"VirtualGatewayId,omitempty"`
	// The ID of the VPN connection.
	VpnConnectionId string `json:"VpnConnectionId,omitempty"`
}
