package crypto

type MailioCrypto interface {
	PublicKeyToMailioAddress(publicKey string) (*string, *error)
	Handshake()
	VerifyHandshake()
}
