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
// DhcpOptionsSet Information about the DHCP options set.
type DhcpOptionsSet struct {
	// If `true`, the DHCP options set is a default one. If `false`, it is not.
	Default bool `json:"Default,omitempty"`
	// The name of the DHCP options set.
	DhcpOptionsName string `json:"DhcpOptionsName,omitempty"`
	// The ID of the DHCP options set.
	DhcpOptionsSetId string `json:"DhcpOptionsSetId,omitempty"`
	// The domain name.
	DomainName string `json:"DomainName,omitempty"`
	// One or more IP addresses for the domain name servers.
	DomainNameServers []string `json:"DomainNameServers,omitempty"`
	// One or more IP addresses for the NTP servers.
	NtpServers []string `json:"NtpServers,omitempty"`
	// One or more tags associated with the DHCP options set.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
