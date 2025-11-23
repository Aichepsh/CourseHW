package verify

import (
	"PurpleHW/3-validation-api/configs"
	"PurpleHW/3-validation-api/internal/pkg/request"
	"PurpleHW/3-validation-api/internal/pkg/resp"
	"PurpleHW/3-validation-api/internal/verify/payload"
	"fmt"
	"log"
	"net/http"
)

type verifyHandler struct {
	*configs.Config
	Codes map[string]string
}
type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &verifyHandler{
		Config: deps.Config,
		Codes:  make(map[string]string),
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *verifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Send reached")
		body, err := request.HandleBody[payload.SendRequest](w, r)
		if err != nil {
			log.Printf("Error HandleBody: %v", err)
			return
		}
		recipient := body.Email
		code, err := h.SendCode(recipient)
		if err != nil {
			log.Printf("failed to send verification email to %s: %v", recipient, err)
			http.Error(w, "Internal server error while sending email", http.StatusInternalServerError)
			return
		}
		responseUser := payload.SendResponse{
			Email: recipient,
			Code:  code,
		}
		h.Codes[responseUser.Email] = responseUser.Code
		resp.WriteJSON(w, responseUser, http.StatusOK)
		fmt.Println("Email sent, code: " + code)
	}
}
func (h *verifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Verify reached")
		body, err := request.HandleBody[payload.VerifyRequest](w, r)
		if err != nil {
			log.Printf("Error HandleBody: %v", err)
			return
		}
		hash := r.PathValue("hash")
		fmt.Println("Hash: ", hash)
		fmt.Println("Code: ", h.Codes[body.Email])
		if h.CheckCode(body.Email, hash) {
			resp.WriteJSON(w, "You confirm email", http.StatusOK)
			fmt.Println("Email verified")
			return
		}
		fmt.Println("Code doesnt confirmed")

	}
}
