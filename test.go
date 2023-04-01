package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

const (
	expectedProjectId    = "my-project-id"
	expectedNetwork      = "my-network"
	expectedSubnetName   = "my-subnet"
	expectedSubnetCidr   = "10.0.0.0/24"
	expectedRouterName   = "my-router"
	expectedRouterRegion = "us-central1"
	expectedVpnName      = "my-vpn"
	terraformCodeDir = "../"
)

var terraformOptions = &terraform.Options{
	TerraformDir: terraformCodeDir,
	Vars: map[string]interface{}{
		"type":          "acceleration",
		"organization_id": "my-org-id",
		"folder_id":        "my-folder-id",
		"billing_account":  "my-billing-account",
		"default_network_tier": "PREMIUM",
		"firewall_rules": []map[string]interface{}{
			{
				"name":       "ssh",
				"direction":  "INGRESS",
				"ranges":     []string{"0.0.0.0/0"},
				"target_tags":[]string{"web"},
				"allow": []map[string]string{
					{
						"protocol": "tcp",
						"ports":    "22",
					},
				},
			},
		},
		"subnets": []map[string]interface{}{
			{
				"name":  expectedSubnetName,
				"cidr":  expectedSubnetCidr,
				"region": "us-central1",
			},
		},
		"secondary_ranges": []map[string]interface{}{
			{
				"subnet_name":    expectedSubnetName,
				"range_name":     "my-secondary-range",
				"ip_cidr_range":  "172.16.0.0/24",
			},
		},
		"routes": []map[string]interface{}{
			{
				"name":                      "my-route",
				"network":                   expectedNetwork,
				"dest_range":                "10.0.1.0/24",
				"next_hop_network_interface": expectedSubnetName,
			},
		},
		"router_name": expectedRouterName,
		"target_network_self_link": "my-target-network",
	},
}

func TestTerratest(t *testing.T) {

	// set the environment variables
	os.Setenv("GOOGLE_PROJECT", expectedProjectId)

	// deploy the infrastructure code
	terraform.InitAndApply(t, terraformOptions)

	// test that the project was created
	project := gcp.GetProject(t, expectedProjectId)
	if project == nil {
		t.Fatalf("Project %q was not created", expectedProjectId)
	}

	// test that the network was created
	network := gcp.GetNetwork(t, expectedProjectId, expectedNetwork)
	if network == nil {
		t.Fatalf("Network %q was not created", expectedNetwork)
	}

	// test that the subnet was created
	subnet := gcp.GetSubnetwork(t, expectedProjectId, expectedSubnetName)
	if subnet == nil {
		t.Fatalf("Subnet %q was not created", expectedSubnetName)
	}

	// test that the IP range matches
	if subnet.IpCidrRange != expectedSubnetCidr {
		t.Fatalf("Subnet IP range %q does not match expected %q", subnet.IpCidrRange, expectedSubnetCidr)
	}

	// test that the router was created
	router := gcp.GetRouter(t, expectedProjectId, expectedRouterRegion, expectedRouterName)
	if router == nil {
		t.Fatalf("Router %q was not created", expectedRouterName)
	}

	// test that the VPN was created
	vpn := gcp.GetVPN(t, fmt.Sprintf("projects/%s/regions/%s/vpnGateways/%s", expectedProjectId, expectedRouterRegion, expectedVpnName))
	if vpn == nil {
		t.Fatalf("VPN %q was not created", expectedVpnName)
	}

	// wait for the VPN to be fully created
	maxRetries := 5
	timeBetweenRetries := 10 * time.Second
	retryableErrors := []string{
		"operation is not ready",
	}
	err := gcp.WaitForVPNConnectionStatusE(t, vpn.Id, maxRetries, timeBetweenRetries, retryableErrors)
	if err != nil {
		t.Fatalf("Failed waiting for VPN connection status: %s", err)
	}
	// clean up infrastructure resources after the test case ends
	defer terraform.Destroy(t, terraformOptions)
}