package test

import (
    "fmt"
    "testing"

    "github.com/gruntwork-io/terratest/modules/gcp"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestProjectCreation(t *testing.T) {
    expectedProjectName := "acceleration-test"
    expectedRegion := "europe-west2"
    expectedNetworkName := "test"
    expectedFirewallRules := []interface{}{
        map[string]interface{}{
            "name":      "allow-http",
            "direction": "ingress",
            "allow": []interface{}{
                map[string]interface{}{
                    "protocol": "tcp",
                    "ports":    []interface{}{"80"},
                },
            },
            "priority": 1000,
        },
    }

    terraformOptions := &terraform.Options{
        TerraformDir: ".",
        Vars: map[string]interface{}{
            "type":                  "test",
            "organization_id":       "32792173",
            "folder_id":             "312900238",
            "billing_account":       "97320230",
            "default_network_tier":  "Premium",
            "firewall_rules":        expectedFirewallRules,
        },
    }

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    projectId := terraform.Output(t, terraformOptions, "project_id")
    networkUrl := fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/networks/%s", projectId, expectedNetworkName)
    firewallUrl := fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/firewalls/allow-http", projectId)

    // Check the project was created with the expected name and region
    actualProjectName := gcp.GetProjectName(t, projectId)
    actualRegion := gcp.GetProjectRegion(t, projectId)
    assert.Equal(t, expectedProjectName, actualProjectName)
    assert.Equal(t, expectedRegion, actualRegion)

    // Check the network was created with the expected name
    actualNetworkName := gcp.GetNetworkName(t, networkUrl)
    assert.Equal(t, expectedNetworkName, actualNetworkName)

    // Check the firewall rule was created with the expected properties
    actualFirewallRules := gcp.GetFirewallRules(t, projectId)
    assert.Len(t, actualFirewallRules, 1)
    actualFirewallRule := actualFirewallRules[0]
    assert.Equal(t, expectedFirewallRules[0].(map[string]interface{})["name"], actualFirewallRule.Name)
    assert.Equal(t, expectedFirewallRules[0].(map[string]interface{})["direction"], actualFirewallRule.Direction)
    assert.Equal(t, expectedFirewallRules[0].(map[string]interface{})["priority"], actualFirewallRule.Priority)
    assert.Equal(t, expectedFirewallRules[0].(map[string]interface{})["allow"], actualFirewallRule.Allow)
    assert.Equal(t, expectedFirewallRules[0].(map[string]interface{})["deny"], actualFirewallRule.Deny)
    assert.Equal(t, expectedFirewallRules[0].(map[string]interface{})["log_config"], actualFirewallRule.LogConfig)

    // Check the firewall rule applies to the network
    actualFirewallRuleNetworks := gcp.GetFirewallRuleNetworks(t, firewallUrl)
    assert.Contains(t, actualFirewallRuleNetworks, expectedNetworkName)
}
