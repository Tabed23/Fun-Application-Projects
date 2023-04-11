# Make a APIGW
resource "aws_api_gateway_rest_api" "apigw" {
  name = var.api_gw_name
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}


resource "aws_api_gateway_deployment" "deployment" {
  depends_on = [
    aws_api_gateway_integration.integration,
  ]
  rest_api_id = aws_api_gateway_rest_api.apigw.id
  stage_name  = var.tag
}


resource "aws_api_gateway_authorizer" "auth" {
  name            = "Authorization"
  rest_api_id     = aws_api_gateway_rest_api.apigw.id
  type            = "COGNITO_USER_POOLS"
  identity_source = "method.request.header.Authorization"
  provider_arns   = [aws_cognito_user_pool.user_pool.arn]
}



resource "aws_api_gateway_resource" "resource" {
  rest_api_id = aws_api_gateway_rest_api.apigw.id
  parent_id   = aws_api_gateway_rest_api.apigw.root_resource_id
  path_part   = "any"
}

resource "aws_api_gateway_method" "get_method" {
  rest_api_id          = aws_api_gateway_rest_api.apigw.id
  resource_id          = aws_api_gateway_resource.resource.id
  http_method          = "GET"
  authorization        = "COGNITO_USER_POOLS"
  authorizer_id        = aws_api_gateway_authorizer.auth.id
  authorization_scopes = ["email"]
  request_parameters = {
    "method.request.path.proxy" = true
  }
}



resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = aws_api_gateway_rest_api.apigw.id
  resource_id             = aws_api_gateway_resource.resource.id
  http_method             = aws_api_gateway_method.get_method.http_method
  content_handling        = "CONVERT_TO_TEXT"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.function.invoke_arn
  passthrough_behavior    = "WHEN_NO_MATCH"
}




