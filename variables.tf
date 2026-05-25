variable "bucket_name" {
  description = "Nombre del bucket s3"
  type        = string
}

variable "environment" {
  description = "Ambiente de despliegue (dev, staging, prod)"
  type        = string
  default     = "dev"
}