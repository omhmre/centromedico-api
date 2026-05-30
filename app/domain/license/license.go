package license

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Clave Pública Maestra Ed25519 para verificar las firmas de las licencias
const PublicKeyHex = "fa32146add1015048389a30a315e1fffbe2565a3105e2d9cab4d2f1901d7b70c"

type LicensePayload struct {
	ClientName   string    `json:"client_name"`
	Rif          string    `json:"rif"`
	ValidUntil   time.Time `json:"valid_until"`
	IsPremium    bool      `json:"is_premium"`
	AllowedUsers int       `json:"allowed_users"`
	IssuedAt     time.Time `json:"issued_at"`
}

// CleanRif sanitiza un RIF removiendo guiones, espacios y pasándolo a minúsculas
func CleanRif(rif string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	return strings.ToLower(reg.ReplaceAllString(rif, ""))
}

// VerifyLicenseKey verifica la autenticidad e integridad de una clave de licencia
func VerifyLicenseKey(licenseKey string, expectedRif string) (*LicensePayload, error) {
	parts := strings.Split(licenseKey, ".")
	if len(parts) != 2 {
		return nil, errors.New("formato de licencia inválido (debe contener payload y firma)")
	}

	payloadB64 := parts[0]
	signatureB64 := parts[1]

	// Decodificar el payload JSON
	payloadBytes, err := base64.StdEncoding.DecodeString(payloadB64)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar payload de la licencia: %w", err)
	}

	// Decodificar la firma
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureB64)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar firma de la licencia: %w", err)
	}

	// Decodificar la clave pública hex
	pubKeyBytes, err := hex.DecodeString(PublicKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar clave pública maestra: %w", err)
	}

	publicKey := ed25519.PublicKey(pubKeyBytes)

	// Verificar la firma criptográfica Ed25519
	if !ed25519.Verify(publicKey, payloadBytes, signatureBytes) {
		return nil, errors.New("firma de la licencia inválida (la clave de licencia ha sido alterada o es falsa)")
	}

	// Parsear el JSON del payload
	var payload LicensePayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("error al interpretar payload de la licencia: %w", err)
	}

	// Verificar expiración
	if time.Now().After(payload.ValidUntil) {
		return nil, fmt.Errorf("la licencia ha expirado el %s", payload.ValidUntil.Format("2006-01-02"))
	}

	// Verificar RIF si se provee uno esperado
	if expectedRif != "" {
		if CleanRif(payload.Rif) != CleanRif(expectedRif) {
			return nil, fmt.Errorf("la licencia pertenece al RIF %s, pero el sistema está registrado para el RIF %s", payload.Rif, expectedRif)
		}
	}

	return &payload, nil
}
