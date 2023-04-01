
describe google_project_iam_binding.project do
  its('role') { should cmp 'roles/editor' }
  its('members') { should include 'user:saurabh.pathaks21@gmail.com' }
end

describe google_compute_network(project: module.project.project_id, name: "#{local.network_name}-acceleration-xpn-001") do
  it { should exist }
  its('routing_config.routing_mode') { should eq 'REGIONAL' }
end

describe google_compute_subnetwork(project: module.project.project_id, region: "region", name: var.subnets[0].name) do
  it { should exist }
  its('ip_cidr_range') { should cmp "#{var.subnets[0].ip_cidr_range}" }
end

describe google_compute_firewall(project: module.project.project_id, name: local.rules[0].name ) do
  it { should exist }
  its('direction') { should eq 'INGRESS' }
  its('allowed') { should include get_regex(local.rules[0].allow[0].ip_protocol) } unless local.rules[0].allow[0].ip_protocol == null
end

describe google_dns_managed_zone(project: module.project.project_id, name: "#{local.network_name}-acceleration-private", visibility_config: { private_visibility_config: { networks: ["#{var.type}-network"] } }) do
  it { should exist }
  its('name') { should cmp "#{local.network_name}-acceleration" }
end

describe google_cloud_router(project: module.project.project_id, region: "europe-west2", name: "#{local.network_name}-router") do
  it { should exist }
end

describe google_compute_vpn_gateway(project: module.project.project_id, region: "europe-west2", name: "#{local.network_name}-gateway") do
  it { should exist }
end

