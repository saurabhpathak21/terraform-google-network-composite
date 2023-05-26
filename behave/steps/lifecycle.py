from behave import *
use_step_matcher("parse")
import requests
from google.cloud import asset_v1
import sys

scope='project/acceleration-hub'
project_query='name:cloudbilling.googleapis.com/project/acceleration-hub/billinInfo'

@given (u' I have a google_project)
def step_imp(context):
   pass
@given (u' its "name" is "acceleration-hub")
def step_imp(context):
   # Create a client
   client = asset_ v1.AssetServiceClient()
   #Intialize request arguments()
   request = asset_v1.SearchAllResourceRequest(scope=scope,query=project_query)
   #Make a request
   page_result = client.search_all_resources(request=request)
   context_actual_state = ""
   #Handle the response
   for response in page_result:
     print(response.state)
     context.actual_state = response.state
@then('u check its lifecycle_state is "Active")
def step_imp(context)
 assert context.actual_state = response.state