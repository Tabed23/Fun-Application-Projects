
output "Congnito-pool-id" {
  value = aws_cognito_user_pool.user_pool.id
}

output "invoke_url" {
  value = aws_api_gateway_deployment.deployment.invoke_url
}
output "lambda_function_name" {
  value = aws_lambda_function.function.function_name
}



