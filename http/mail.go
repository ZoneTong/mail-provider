package http

import (
	"encoding/json"
	"github.com/ZoneTong/mail-provider/config" // "github.com/toolkits/smtp"
	"github.com/farmerx/mail"
	"github.com/toolkits/web/param"
	"log"
	"net/http"
	"strings"
)

func configProcRoutes() {

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		// tos = strings.Replace(tos, ",", ";", -1)
		log.Println(tos, subject, content)

		// s := smtp.New(cfg.Smtp.Addr, cfg.Smtp.Username, cfg.Smtp.Password)
		// err := s.SendMail(cfg.Smtp.From, tos, subject, content)

		server, err := json.Marshal(cfg.Smtp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		email := mail.NewEMail(string(server))
		email.Auth = mail.NTLMAuth(email.Host, email.Username, email.Password, mail.NTLMVersion1)
		email.To = strings.Split(tos, ",")
		email.Subject = subject
		email.Text = content
		err = email.Send()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "success", http.StatusOK)
		}
	})

}
