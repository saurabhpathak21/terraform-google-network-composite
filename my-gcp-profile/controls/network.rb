


# copyright: 2018, The Authors

title "network-composite-module integration test"

gcp_project_id = input("gcp_project_id")
network        = input("network")

# you add controls here
control "network-composite-module" do                        # A unique ID for this control
  impact 0.7                                # The criticality, if this control fails.
  title "Test the Network Module"             # A human-readable title
  desc "Test cases for the network composite module"

  describe google_project(project: gcp_project_id) do
    it { should exist }
    its('project_id') { should cmp gcp_project_id }
    its('lifecycle_state') { should cmp 'ACTIVE' }
  end
  
  describe google_compute_network(project: gcp_project_id , region: 'europe-west2', name: network) do
    it { should exist }
    its('routing_config.routing_mode') { should eq 'GLOBAL' }
  end
  
  describe google_compute_subnetwork(project: gcp_project_id, region: 'europe-west2', name: '') do
    it { should exist }
    its('ip_cidr_range') { should cmp '' }
  end
    
  describe google_compute_firewall(project: gcp_project_id, name: '' ) do
    it { should exist }
    its('direction') { should eq 'INGRESS' }
  end

  describe google_compute_router(project: gcp_project_id, region: 'europe-west2', router: 'nonexistent', name: 'abc') do
    it { should exist }
  end
    
  describe google_compute_vpn_tunnel(project: gcp_project_id, region: 'europe-west2', name: 'gcp-vpn-tunnel') do
    it { should exist }
  end
    
end


