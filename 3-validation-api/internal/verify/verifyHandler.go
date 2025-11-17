package verify

import (
	"PurpleHW/3-validation-api/configs"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type verifyHandler struct {
	*configs.Config
}
type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &verifyHandler{
		deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *verifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler reached")
		for range 1 {
			e := &email.Email{
				To:          []string{""},
				From:        "",
				Subject:     "Чето там",
				Text:        []byte("Alimrantus"),
				Attachments: []*email.Attachment{},
			}

			err := e.Send("smtp.mail.ru:587", smtp.PlainAuth("", "", "", "smtp.mail.ru"))
			if err != nil {
				http.Error(w, "Не удалось отправить письмо", http.StatusInternalServerError)
				w.Write([]byte("Not ok"))
				log.Print("Email doesn't sent")
				return
			}

		}
		w.Header().Set("Content-Type", "text; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		fmt.Println("Email sent")
	}
}
func (h *verifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
