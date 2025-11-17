package verify

import (
	"PurpleHW/3-validation-api/configs"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type verifyHandler struct {
	*configs.Config
	Code string
}
type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &verifyHandler{
		Config: deps.Config,
		Code:   "",
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func RandomDigits(n int) string {
	digits := "0123456789"
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		result[i] = digits[num.Int64()]
	}
	return string(result)
}
func (h *verifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Send reached")
		h.Code = RandomDigits(6)
		verifyURL := "http://localhost:8080/verify/" + h.Code
		e := &email.Email{
			To:          []string{h.Config.Address},
			From:        h.Config.Email,
			Subject:     "Some subject",
			Text:        []byte(h.Code),
			HTML:        []byte(fmt.Sprintf(`<a href="%s">Нажми сюда чтобы подтвердить</a>`, verifyURL)),
			Attachments: []*email.Attachment{},
		}
		err := e.Send("smtp.mail.ru:587", smtp.PlainAuth("", h.Config.Email, h.Config.Password, "smtp.mail.ru"))
		if err != nil {
			http.Error(w, "Не удалось отправить письмо", http.StatusInternalServerError)
			log.Print("Email doesn't sent", err)
			return
		}

		w.Header().Set("Content-Type", "text; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		fmt.Println("Email sent, code: " + h.Code)
	}
}
func (h *verifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Verify reached")
		hash := r.PathValue("hash")
		fmt.Println("Hash: ", hash)
		fmt.Println("Code: ", h.Code)
		if hash != h.Code {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid verification code"))
			log.Print("Invalid verification code")
			return
		}
		h.Code = ""
		w.Header().Set("Content-Type", "text; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("You confirm email"))
		fmt.Println("Email verified")

	}
}
