module "s3" {
  source      = "github.com/ServandoSoto/terraform-s3-module?ref=v1.0.0"
  bucket_name = "mi-bucket-ejemplo"
  environment = "dev"
}

output "bucket_name" {
  value = module.s3.bucket_name
}

output "bucket_arn" {
  value = module.s3.bucket_arn
}
