// Paquete test — nombre estándar para tests en Go
package test

import (
	"testing" // Librería estándar de Go para tests

	// Terratest: librería que permite desplegar infraestructura Terraform real y validarla
	"github.com/gruntwork-io/terratest/modules/terraform"
	// Testify: librería para hacer assertions (verificar que algo cumple una condición)
	"github.com/stretchr/testify/assert"
)

// TestS3Module es la función principal del test — Go la reconoce como test por empezar con "Test"
func TestS3Module(t *testing.T) {

	// terraformOptions define la configuración del despliegue:
	// TerraformDir apunta al módulo que vamos a testear
	// Vars son los valores que le pasamos a las variables del módulo
	terraformOptions := &terraform.Options{
		TerraformDir: "../s3_module",
		Vars: map[string]interface{}{
			"bucket_name": "mi-bucket-prueba2",
			"environment": "dev",
		},
	}

	// defer garantiza que terraform destroy se ejecute al final del test
	// aunque el test falle — evita dejar recursos huérfanos en AWS
	defer terraform.Destroy(t, terraformOptions)

	// InitAndApply ejecuta terraform init + terraform apply en el módulo
	// despliega la infraestructura real en AWS
	terraform.InitAndApply(t, terraformOptions)

	// terraform.Output captura el valor del output del módulo
	// assert.NotEmpty verifica que no esté vacío — si lo está, el test falla
	bucketName := terraform.Output(t, terraformOptions, "bucket_name")
	assert.NotEmpty(t, bucketName)

	bucketArn := terraform.Output(t, terraformOptions, "bucket_arn")
	assert.NotEmpty(t, bucketArn)
}
