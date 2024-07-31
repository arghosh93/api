package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ClusterNetwork = map[string]string{
	"":                 "ClusterNetwork was used by OpenShift SDN. DEPRECATED: OpenShift SDN is no longer supported and this object is no longer used in any way by OpenShift.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata":         "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"network":          "Network is a CIDR string specifying the global overlay network's L3 space",
	"hostsubnetlength": "HostSubnetLength is the number of bits of network to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pods",
	"serviceNetwork":   "ServiceNetwork is the CIDR range that Service IP addresses are allocated from",
	"pluginName":       "PluginName is the name of the network plugin being used",
	"clusterNetworks":  "ClusterNetworks is a list of ClusterNetwork objects that defines the global overlay network's L3 space by specifying a set of CIDR and netmasks that the SDN can allocate addresses from.",
	"vxlanPort":        "VXLANPort sets the VXLAN destination port used by the cluster. It is set by the master configuration file on startup and cannot be edited manually. Valid values for VXLANPort are integers 1-65535 inclusive and if unset defaults to 4789. Changing VXLANPort allows users to resolve issues between openshift SDN and other software trying to use the same VXLAN destination port.",
	"mtu":              "MTU is the MTU for the overlay network. This should be 50 less than the MTU of the network connecting the nodes. It is normally autodetected by the cluster network operator.",
}

func (ClusterNetwork) SwaggerDoc() map[string]string {
	return map_ClusterNetwork
}

var map_ClusterNetworkEntry = map[string]string{
	"":                 "ClusterNetworkEntry defines an individual cluster network. The CIDRs cannot overlap with other cluster network CIDRs, CIDRs reserved for external ips, CIDRs reserved for service networks, and CIDRs reserved for ingress ips.",
	"CIDR":             "CIDR defines the total range of a cluster networks address space.",
	"hostSubnetLength": "HostSubnetLength is the number of bits of the accompanying CIDR address to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pods.",
}

func (ClusterNetworkEntry) SwaggerDoc() map[string]string {
	return map_ClusterNetworkEntry
}

var map_ClusterNetworkList = map[string]string{
	"":         "ClusterNetworkList is a collection of ClusterNetworks\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is the list of cluster networks",
}

func (ClusterNetworkList) SwaggerDoc() map[string]string {
	return map_ClusterNetworkList
}

var map_EgressNetworkPolicy = map[string]string{
	"":         "EgressNetworkPolicy was used by OpenShift SDN. DEPRECATED: OpenShift SDN is no longer supported and this object is no longer used in any way by OpenShift.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the specification of the current egress network policy",
}

func (EgressNetworkPolicy) SwaggerDoc() map[string]string {
	return map_EgressNetworkPolicy
}

var map_EgressNetworkPolicyList = map[string]string{
	"":         "EgressNetworkPolicyList is a collection of EgressNetworkPolicy\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of policies",
}

func (EgressNetworkPolicyList) SwaggerDoc() map[string]string {
	return map_EgressNetworkPolicyList
}

var map_EgressNetworkPolicyPeer = map[string]string{
	"":             "EgressNetworkPolicyPeer specifies a target to apply egress network policy to",
	"cidrSelector": "CIDRSelector is the CIDR range to allow/deny traffic to. If this is set, dnsName must be unset Ideally we would have liked to use the cidr openapi format for this property. But openshift-sdn only supports v4 while specifying the cidr format allows both v4 and v6 cidrs We are therefore using a regex pattern to validate instead.",
	"dnsName":      "DNSName is the domain name to allow/deny traffic to. If this is set, cidrSelector must be unset",
}

func (EgressNetworkPolicyPeer) SwaggerDoc() map[string]string {
	return map_EgressNetworkPolicyPeer
}

var map_EgressNetworkPolicyRule = map[string]string{
	"":     "EgressNetworkPolicyRule contains a single egress network policy rule",
	"type": "type marks this as an \"Allow\" or \"Deny\" rule",
	"to":   "to is the target that traffic is allowed/denied to",
}

func (EgressNetworkPolicyRule) SwaggerDoc() map[string]string {
	return map_EgressNetworkPolicyRule
}

var map_EgressNetworkPolicySpec = map[string]string{
	"":       "EgressNetworkPolicySpec provides a list of policies on outgoing network traffic",
	"egress": "egress contains the list of egress policy rules",
}

func (EgressNetworkPolicySpec) SwaggerDoc() map[string]string {
	return map_EgressNetworkPolicySpec
}

var map_HostSubnet = map[string]string{
	"":            "HostSubnet was used by OpenShift SDN. DEPRECATED: OpenShift SDN is no longer supported and this object is no longer used in any way by OpenShift.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata":    "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"host":        "Host is the name of the node. (This is the same as the object's name, but both fields must be set.)",
	"hostIP":      "HostIP is the IP address to be used as a VTEP by other nodes in the overlay network",
	"subnet":      "Subnet is the CIDR range of the overlay network assigned to the node for its pods",
	"egressIPs":   "EgressIPs is the list of automatic egress IP addresses currently hosted by this node. If EgressCIDRs is empty, this can be set by hand; if EgressCIDRs is set then the master will overwrite the value here with its own allocation of egress IPs.",
	"egressCIDRs": "EgressCIDRs is the list of CIDR ranges available for automatically assigning egress IPs to this node from. If this field is set then EgressIPs should be treated as read-only.",
}

func (HostSubnet) SwaggerDoc() map[string]string {
	return map_HostSubnet
}

var map_HostSubnetList = map[string]string{
	"":         "HostSubnetList is a collection of HostSubnets\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is the list of host subnets",
}

func (HostSubnetList) SwaggerDoc() map[string]string {
	return map_HostSubnetList
}

var map_NetNamespace = map[string]string{
	"":          "NetNamespace was used by OpenShift SDN. DEPRECATED: OpenShift SDN is no longer supported and this object is no longer used in any way by OpenShift.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata":  "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"netname":   "NetName is the name of the network namespace. (This is the same as the object's name, but both fields must be set.)",
	"netid":     "NetID is the network identifier of the network namespace assigned to each overlay network packet. This can be manipulated with the \"oc adm pod-network\" commands.",
	"egressIPs": "EgressIPs is a list of reserved IPs that will be used as the source for external traffic coming from pods in this namespace. (If empty, external traffic will be masqueraded to Node IPs.)",
}

func (NetNamespace) SwaggerDoc() map[string]string {
	return map_NetNamespace
}

var map_NetNamespaceList = map[string]string{
	"":         "NetNamespaceList is a collection of NetNamespaces\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is the list of net namespaces",
}

func (NetNamespaceList) SwaggerDoc() map[string]string {
	return map_NetNamespaceList
}

// AUTO-GENERATED FUNCTIONS END HERE
