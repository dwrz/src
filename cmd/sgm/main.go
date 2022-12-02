package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"code.dwrz.net/src/cmd/sgm/config"
	"code.dwrz.net/src/pkg/gmail"
	"code.dwrz.net/src/pkg/log"
)

func main() {
	var l = log.New(os.Stderr)

	// Get the configuration.
	cfg, err := config.New()
	if err != nil {
		l.Error.Fatalf("failed to get config: %v", err)
	}

	// Compose email.
	var str strings.Builder

	fmt.Fprintf(&str, "From: %s\r\n", cfg.From)
	fmt.Fprintf(&str, "To: %s\r\n", cfg.To)
	fmt.Fprintf(&str, "Subject: %s\r\n", cfg.Subject)

	str.WriteString("MIME-Version: 1.0\n")
	str.WriteString("Content-Type: multipart/mixed; boundary=\"=-=-=\"\n")
	str.WriteString("\n")
	str.WriteString("--=-=-=\n")
	str.WriteString("Content-Type: text/plain; format=flowed\n")
	str.WriteString("\n")
	str.WriteString("\n")
	str.WriteString(cfg.Text)
	str.WriteString("\n")
	str.WriteString("--=-=-=\n")

	// Add the attachment, if provided.
	if cfg.Attachment != "" {
		data, err := os.ReadFile(cfg.Attachment)
		if err != nil {
			l.Error.Fatalf("failed to open file: %v", err)
		}

		fmt.Fprintf(
			&str,
			"Content-Type: %s\n", http.DetectContentType(data),
		)
		fmt.Fprintf(
			&str,
			"Content-Disposition: attachment; filename=%s\n",
			*&cfg.Attachment,
		)
		str.WriteString("Content-Transfer-Encoding: base64\n")
		fmt.Fprintf(&str, "Content-Description: %s\n", cfg.Attachment)
		str.WriteString("\n")
		str.WriteString(base64.StdEncoding.EncodeToString([]byte(data)))
		str.WriteString("--=-=-=\n")
	}

	// Send the email.
	if err := smtp.SendMail(
		gmail.Address,
		&gmail.Auth{
			Username: cfg.User,
			Password: cfg.Pass,
		},
		cfg.From,
		strings.Split(cfg.To, ","),
		[]byte(str.String()),
	); err != nil {
		l.Error.Fatalf("failed to send email: %v", err)
	}
}
