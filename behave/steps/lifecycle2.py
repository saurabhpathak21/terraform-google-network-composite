from behave import given, when, then

@given('I have a variable "{variable_name}" with value "{variable_value}"')
def step_impl(context, variable_name, variable_value):
    context.variables[variable_name] = variable_value

@when('I print the value of the variable')
def step_impl(context):
    context.output = context.variables['name']

@then('the output should be "{expected_output}"')
def step_impl(context, expected_output):
    assert context.output == expected_output
