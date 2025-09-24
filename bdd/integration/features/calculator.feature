Feature: Calculator

  Scenario: Add two numbers
    Given the first number is 3
    And the second number is 5
    When the numbers are added
    Then the result should be 8
