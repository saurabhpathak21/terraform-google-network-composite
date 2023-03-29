<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13 |
| <a name="requirement_google"></a> [google](#requirement\_google) | ~> 4.5 |
| <a name="requirement_google-beta"></a> [google-beta](#requirement\_google-beta) | ~> 4.5 |
| <a name="requirement_null"></a> [null](#requirement\_null) | ~> 3.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | ~> 4.5 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_cloud_router"></a> [cloud\_router](#module\_cloud\_router) | terraform-google-modules/cloud-router/google | 5.0.0 |
| <a name="module_dns-peering-zone"></a> [dns-peering-zone](#module\_dns-peering-zone) | terraform-google-modules/cloud-dns/google | 3.0.0 |
| <a name="module_dns-private-zone"></a> [dns-private-zone](#module\_dns-private-zone) | terraform-google-modules/cloud-dns/google | 4.2.1 |
| <a name="module_firewall_rules"></a> [firewall\_rules](#module\_firewall\_rules) | terraform-google-modules/network/google//modules/firewall-rules | 6.0.1 |
| <a name="module_project"></a> [project](#module\_project) | terraform-google-modules/project-factory/google | 14.2.0 |
| <a name="module_routes"></a> [routes](#module\_routes) | terraform-google-modules/network/google//modules/routes | 6.0.1 |
| <a name="module_subnets"></a> [subnets](#module\_subnets) | terraform-google-modules/network/google//modules/subnets | 6.0.1 |
| <a name="module_vpc"></a> [vpc](#module\_vpc) | terraform-google-modules/network/google//modules/vpc | 6.0.1 |
| <a name="module_vpn"></a> [vpn](#module\_vpn) | terraform-google-modules/vpn/google//modules/vpn_ha | ~> 1.3.0 |

## Resources

| Name | Type |
|------|------|
| [google_project_iam_binding.project](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_binding) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_auto_create_subnetworks"></a> [auto\_create\_subnetworks](#input\_auto\_create\_subnetworks) | When set to true, the network is created in 'auto subnet mode' and it will create a subnet for each region automatically across the 10.128.0.0/9 address range. When set to false, the network is created in 'custom subnet mode' so the user can explicitly connect subnetwork resources. | `bool` | `false` | no |
| <a name="input_billing_account"></a> [billing\_account](#input\_billing\_account) | The ID of the billing account to associate this project with | `any` | n/a | yes |
| <a name="input_default_network_tier"></a> [default\_network\_tier](#input\_default\_network\_tier) | Default Network Service Tier for resources created in this project. If unset, the value will not be modified. See https://cloud.google.com/network-tiers/docs/using-network-service-tiers and https://cloud.google.com/network-tiers. | `string` | `""` | no |
| <a name="input_delete_default_internet_gateway_routes"></a> [delete\_default\_internet\_gateway\_routes](#input\_delete\_default\_internet\_gateway\_routes) | If set, ensure that all routes within the network specified whose names begin with 'default-route' and with a next hop of 'default-internet-gateway' are deleted | `bool` | `true` | no |
| <a name="input_description"></a> [description](#input\_description) | An optional description of this resource. The resource must be recreated to modify this field. | `string` | `""` | no |
| <a name="input_domain"></a> [domain](#input\_domain) | Zone domain. | `string` | `"foo.local."` | no |
| <a name="input_enable_peering"></a> [enable\_peering](#input\_enable\_peering) | enable the dns peering and provide the target project id | `bool` | `false` | no |
| <a name="input_enable_shared_vpc_host_project"></a> [enable\_shared\_vpc\_host\_project](#input\_enable\_shared\_vpc\_host\_project) | Makes this project a Shared VPC host if 'true' (default 'false') | `bool` | `true` | no |
| <a name="input_firewall_rules"></a> [firewall\_rules](#input\_firewall\_rules) | List of firewall rules | `any` | `[]` | no |
| <a name="input_folder_id"></a> [folder\_id](#input\_folder\_id) | The folder id for the associated project | `string` | `""` | no |
| <a name="input_labels"></a> [labels](#input\_labels) | A set of key/value label pairs to assign to this ManagedZone | `map(any)` | <pre>{<br>  "owner": "newacceleration",<br>  "version": "1.0"<br>}</pre> | no |
| <a name="input_module_depends_on"></a> [module\_depends\_on](#input\_module\_depends\_on) | List of modules or resources this module depends on. | `list(any)` | `[]` | no |
| <a name="input_mtu"></a> [mtu](#input\_mtu) | The network MTU (If set to 0, meaning MTU is unset - defaults to '1460'). Recommended values: 1460 (default for historic reasons), 1500 (Internet default), or 8896 (for Jumbo packets). Allowed are all values in the range 1300 to 8896, inclusively. | `number` | `0` | no |
| <a name="input_name"></a> [name](#input\_name) | DNS zone name. | `string` | `"foo-local"` | no |
| <a name="input_network_self_links"></a> [network\_self\_links](#input\_network\_self\_links) | Self link of the network that will be allowed to query the zone. | `list` | `[]` | no |
| <a name="input_organization_id"></a> [organization\_id](#input\_organization\_id) | The organization id for the associated services | `any` | n/a | yes |
| <a name="input_router_name"></a> [router\_name](#input\_router\_name) | Name of router, leave blank to create one. | `string` | `"vpn-router"` | no |
| <a name="input_routes"></a> [routes](#input\_routes) | List of routes being created in this VPC | `list(map(string))` | `[]` | no |
| <a name="input_routing_mode"></a> [routing\_mode](#input\_routing\_mode) | The network routing mode (default 'GLOBAL') | `string` | `"GLOBAL"` | no |
| <a name="input_secondary_ranges"></a> [secondary\_ranges](#input\_secondary\_ranges) | Secondary ranges that will be used in some of the subnets | `map(list(object({ range_name = string, ip_cidr_range = string })))` | `{}` | no |
| <a name="input_subnets"></a> [subnets](#input\_subnets) | The list of subnets being created | `list(map(string))` | `[]` | no |
| <a name="input_target_network_self_link"></a> [target\_network\_self\_link](#input\_target\_network\_self\_link) | Self link of the network that the zone will peer to. | `string` | `""` | no |
| <a name="input_type"></a> [type](#input\_type) | The type of the project if it is hub or spoke. | `string` | `"hub"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_name_servers"></a> [name\_servers](#output\_name\_servers) | Zone name servers. |
| <a name="output_network"></a> [network](#output\_network) | The created network |
| <a name="output_network_id"></a> [network\_id](#output\_network\_id) | The ID of the VPC being created |
| <a name="output_network_name"></a> [network\_name](#output\_network\_name) | The name of the VPC being created |
| <a name="output_network_self_link"></a> [network\_self\_link](#output\_network\_self\_link) | The URI of the VPC being created |
| <a name="output_project_id"></a> [project\_id](#output\_project\_id) | VPC project id |
| <a name="output_project_info"></a> [project\_info](#output\_project\_info) | The ID of the created project |
| <a name="output_route_names"></a> [route\_names](#output\_route\_names) | The route names associated with this VPC |
| <a name="output_router_name"></a> [router\_name](#output\_router\_name) | router name |
| <a name="output_subnets"></a> [subnets](#output\_subnets) | A map with keys of form subnet\_region/subnet\_name and values being the outputs of the google\_compute\_subnetwork resources used to create corresponding subnets. |
| <a name="output_subnets_flow_logs"></a> [subnets\_flow\_logs](#output\_subnets\_flow\_logs) | Whether the subnets will have VPC flow logs enabled |
| <a name="output_subnets_ids"></a> [subnets\_ids](#output\_subnets\_ids) | The IDs of the subnets being created |
| <a name="output_subnets_ips"></a> [subnets\_ips](#output\_subnets\_ips) | The IPs and CIDRs of the subnets being created |
| <a name="output_subnets_names"></a> [subnets\_names](#output\_subnets\_names) | The names of the subnets being created |
| <a name="output_subnets_private_access"></a> [subnets\_private\_access](#output\_subnets\_private\_access) | Whether the subnets will have access to Google API's without a public IP |
| <a name="output_subnets_regions"></a> [subnets\_regions](#output\_subnets\_regions) | The region where the subnets will be created |
| <a name="output_subnets_secondary_ranges"></a> [subnets\_secondary\_ranges](#output\_subnets\_secondary\_ranges) | The secondary ranges associated with these subnets |
| <a name="output_subnets_self_links"></a> [subnets\_self\_links](#output\_subnets\_self\_links) | The self-links of subnets being created |
<!-- END_TF_DOCS -->