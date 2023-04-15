Feature: project feature 
  Scenario: Ensure project module should include  hub or spoke
    Given I have google_project defined
    Then it must contain name
    And its value must be acceleration-hub


  Scenario Outline: Ensure that specific tags are defined
    Given I have resource that supports labels defined
    When it has labels
    Then it must contain labels
    Then it must contain "<labels>"
    And its value must match the "<value>" regex

    Examples:
      | labels        | value              |
      | Name        | .+                 |
      | workstream  | .+                 |
      | role        | .+                 |
      | environment | ^(prod\|uat\|dev)$ |


  Scenario: Ensure my specific role is editor
    Given I have google_project_iam_binding defined
    Then it must contain role
    And its value must be roles/editor
  
  Scenario: Ensure my VPC network is created
    Given I have google_compute_network defined
    Then it must contain name
    And its value must be hub-acceleration-xpn-001

  Scenario: Ensure my router region is europe-west2
    Given I have google_compute_router defined
    Then it must contain region
    And its value must be europe-west2

  Scenario: Ensure my vpn gateway is created
    Given I have google_compute_ha_vpn_gateway defined
    Then it must contain name
    And its value must be hub-vpn
