resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "aws_s3_bucket" "storage" {
  bucket = "${var.bucket_name}-${random_id.bucket_suffix.hex}"

  tags = {
    Name        = "${var.bucket_name}"
    Environment = var.environment
  }
}

resource "aws_s3_bucket_versioning" "storage" {
  bucket = aws_s3_bucket.storage.id

  versioning_configuration {
    status = "Enabled"
  }
}
resource "aws_s3_bucket_public_access_block" "storage" {
  bucket                  = aws_s3_bucket.storage.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}
