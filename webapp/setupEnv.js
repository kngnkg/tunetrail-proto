module.exports = () => {
  process.env.NEXT_PUBLIC_API_ROOT = "http://api.example.com"
  process.env.NEXT_PUBLIC_AUTH_API_ROOT = "http://example.com/api"
  process.env.TUNETRAIL_AWS_REGION = "test-region"
  process.env.COGNITO_USER_POOL_ID = "test-pool-id"
  process.env.ALLOWED_DOMAIN = "example.com"
  process.env.API_ROOT = "http://api.example.com"
}
