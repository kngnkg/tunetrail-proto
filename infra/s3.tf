resource "aws_s3_bucket" "schema" {
  bucket = var.db_schema_bucket_name
}

resource "aws_s3_bucket_versioning" "schema_versioning" {
  bucket = aws_s3_bucket.schema.id

  versioning_configuration {
    mfa_delete = "Disabled"
    status     = "Enabled"
  }
}
