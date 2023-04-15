Feature: role feature

Scenario: Ensure my specific role is editor
  Given I have google_project_iam_binding defined
  Then it must contain role
  And its value must be roles/editor
