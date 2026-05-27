# Ejemplo de uso - terraform-s3-module

Este directorio contiene un ejemplo funcional de cómo usar el módulo `s3_module`.

## ¿Qué despliega este ejemplo?

- Un bucket S3 con nombre único generado con sufijo aleatorio
- Versionado habilitado
- Acceso público bloqueado

## Uso

```hcl
module "s3" {
  source      = "github.com/ServandoSoto/terraform-s3-module?ref=v1.0.0"
  bucket_name = "mi-bucket-ejemplo"
  environment = "dev"
}
```

## Outputs disponibles

| Output | Descripción |
|--------|-------------|
| bucket_name | Nombre del bucket S3 creado |
| bucket_arn | ARN del bucket S3 creado |

## Requisitos

- Terraform >= 1.0.0
- AWS Provider ~> 5.0
- Random Provider ~> 3.0
- Credenciales AWS configuradas
