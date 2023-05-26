@positive
Feature: List the subnet of Project X
    Scenario: Ensure the project name is valid
      Given: I have a google_project
      Then its "name" is "acceleration-hub"
      And check its lifecycle_state is "Active"

    Scenario: Ensure the project name is valid
      Given: I have a google_project defined
      Then it has the network
      And List and Validate the network name
