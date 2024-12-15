package utilities

import (
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type Result struct {
	QRCode string
	URL    string
	SECRET string
}

func SetupTwoFactorAuth(userId uuid.UUID) (Result, error) {
	// Generate a new TOTP key for the user
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ExpenseTracker",
		AccountName: userId.String(),
	})
	if err != nil {
		return Result{}, err
	}

	qrCodeData, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		return Result{}, err
	}

	return Result{
		QRCode: string(qrCodeData),
		URL:    key.URL(),
		SECRET: key.Secret(),
	}, nil
}

func VerifyTwoFactorAuth(secret string, code string) bool {
	// Verify the TOTP code for the user
	res := totp.Validate(code, secret)
	return res
}
