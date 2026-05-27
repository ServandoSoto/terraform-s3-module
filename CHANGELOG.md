# CHANGELOG - terraform-s3-module

Todos los cambios notables de este módulo serán documentados en este archivo.
Formato basado en [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [v0.1.0] - 2026-05-26
### Added
- Commit inicial del módulo s3
- Recurso `aws_s3_bucket` con nombre único usando `random_id`
- Recurso `aws_s3_bucket_versioning` con versionado habilitado
- Recurso `aws_s3_bucket_public_access_block` con acceso público bloqueado
- Archivo `variables.tf` con variables: bucket_name, environment
- Archivo `outputs.tf` con outputs: bucket_name, bucket_arn

---

## [v0.2.0] - 2026-05-26
### Added
- Documentación automática generada con terraform-docs
- Archivo `README.md` con descripción de inputs, outputs y recursos del módulo

---

## [v0.3.0] - 2026-05-26
### Added
- Carpeta `s3_module_test/` con test de infraestructura usando Terratest
- Archivo `s3_test.go` que valida outputs bucket_name y bucket_arn
- Archivo `.gitignore` para excluir archivos terraform generados localmente

---

## [v0.3.1] - 2026-05-26
### Fixed
- Corrección de ruta `TerraformDir` en `s3_test.go` de `../s3_module` a `..`

---

## [v1.0.0] - 2026-05-27
### Changed
- Versión estable y funcional del módulo
- Arquitectura modular completa desacoplada del código monolítico original
