Feature: project feature 
  Scenario: Ensure project module should include  hub or spoke
    Given I have google_project defined
    Then it must contain name
    And its value must be acceleration-hub
