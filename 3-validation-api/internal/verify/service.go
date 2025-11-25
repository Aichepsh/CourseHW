package verify

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func RandomDigits(n int) string {
	digits := "0123456789"
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		result[i] = digits[num.Int64()]
	}
	return string(result)
}
func (h *verifyHandler) SendCode(recipient string) (string, error) {
	code := RandomDigits(6)
	verifyURL := "http://localhost:8081/verify/" + code
	e := &email.Email{
		To:          []string{recipient},
		From:        h.Config.Email,
		Subject:     "Some subject",
		Text:        []byte(code),
		HTML:        []byte(fmt.Sprintf(`<a href="%s">Подтвердить email</a>`, verifyURL)),
		Attachments: []*email.Attachment{},
	}
	err := e.Send("smtp.mail.ru:587", smtp.PlainAuth("", h.Config.Email, h.Config.Password, h.Config.Address))
	if err != nil {
		return "", fmt.Errorf("failed to send verification email: %w", err)
	}
	return code, nil
}
func (h *verifyHandler) CheckCode(clientCode string) bool {
	h.Storage.Load()
	clientEmail, ok := h.Storage.CodeToEmail[clientCode]
	if !ok {
		log.Println("No code for email: ", clientEmail)
		return false
	}
	if h.Storage.EmailToCode[clientEmail] != clientCode {
		log.Println("Invalid code for email: ", clientEmail)
		return false
	}
	delete(h.Storage.EmailToCode, clientEmail)
	delete(h.Storage.CodeToEmail, clientCode)
	h.Storage.Save()
	return true
}
