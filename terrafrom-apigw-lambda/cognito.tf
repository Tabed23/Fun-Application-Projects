# Congnito
resource "aws_cognito_user_pool" "user_pool" {

  name                     = var.cognito_pool_name
  username_attributes      = ["email"]
  auto_verified_attributes = ["email"]

}

resource "aws_cognito_user_pool_domain" "user_pool_domain" {
  domain       = "authing-dev"
  user_pool_id = aws_cognito_user_pool.user_pool.id
}



resource "aws_cognito_user_pool_client" "user_pool_client" {
  depends_on = [
    aws_cognito_user_pool.user_pool,
  ]

  name                                 = var.cognito_user_pool_client_name
  user_pool_id                         = aws_cognito_user_pool.user_pool.id
  generate_secret                      = true
  allowed_oauth_scopes                 = ["openid", "email", "profile", "phone", "aws.cognito.signin.user.admin"]
  allowed_oauth_flows                  = ["code", "implicit"]
  supported_identity_providers         = ["COGNITO"]
  allowed_oauth_flows_user_pool_client = true
  callback_urls                        = ["https://example.com/callback"]
  default_redirect_uri                 = "https://example.com/callback"

}
