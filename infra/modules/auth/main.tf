resource "aws_cognito_user_pool" "tunetrail" {
  name                     = "user_pool"
  alias_attributes         = ["email"] # emailをユーザー名として使用する
  auto_verified_attributes = ["email"]
  mfa_configuration        = "OFF"
  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }
  schema {
    attribute_data_type = "String"
    name                = "email"
    required            = true
  }
  username_configuration {
    case_sensitive = false # ユーザー名は大文字小文字を区別しない
  }
}

resource "aws_cognito_user_pool_client" "client" {
  name            = "client"
  user_pool_id    = aws_cognito_user_pool.tunetrail.id
  generate_secret = true // サーバーサイドで認証するため、シークレットを生成する
  explicit_auth_flows = [
    "ADMIN_NO_SRP_AUTH", // AdminInitiateAuth APIを使用して認証する
  ]
  allowed_oauth_flows_user_pool_client = false // OAuthフローを使用しない
}
