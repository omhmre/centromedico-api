package license

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"testing"
	"time"
)

func TestVerifyLicenseKey(t *testing.T) {
	// 1. Decodificar la clave pública hex para obtener la clave pública real
	pubKeyBytes, err := hex.DecodeString(PublicKeyHex)
	if err != nil {
		t.Fatalf("Failed to decode Master Public Key: %v", err)
	}
	_ = pubKeyBytes // Evitar advertencia de no usado

	// Como no tenemos la clave privada maestra real cargada en tiempo de ejecución (se genera fuera de la app),
	// podemos simular el proceso de validación generando un par de llaves temporal para la prueba.
	tempPub, tempPriv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("Failed to generate temp keys: %v", err)
	}
	_ = tempPub // Evitar advertencia de no usado

	// Payload de prueba
	payload := LicensePayload{
		ClientName:   "Centro Medico Test",
		Rif:          "J-12345678-9",
		ValidUntil:   time.Now().Add(24 * time.Hour),
		IsPremium:    true,
		AllowedUsers: 5,
		IssuedAt:     time.Now(),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	// Firmar con la clave privada temporal
	signature := ed25519.Sign(tempPriv, payloadBytes)

	// Codificar en Base64
	payloadB64 := base64.StdEncoding.EncodeToString(payloadBytes)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)
	licenseKey := payloadB64 + "." + signatureB64

	t.Run("Sanitización de RIF", func(t *testing.T) {
		rif1 := "J-12345678-9"
		rif2 := "j123456789"
		if CleanRif(rif1) != CleanRif(rif2) {
			t.Errorf("CleanRif failed: %s should equal %s", CleanRif(rif1), CleanRif(rif2))
		}
	})

	t.Run("Verificación con firma falsa", func(t *testing.T) {
		// Intentar verificar con la firma falsa o corrupta usando la PublicKeyHex real
		_, err := VerifyLicenseKey(licenseKey, "J-12345678-9")
		if err == nil {
			t.Error("Expected verification to fail with custom keys since PublicKeyHex is fixed, but it succeeded")
		} else {
			t.Logf("Expected failure occurred: %v", err)
		}
	})
}
