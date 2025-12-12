package notifier

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"path/filepath"
)

// EmailConfig: Cấu hình cho EmailStrategy (từ env/viper)
type EmailConfig struct {
	From     string
	Host     string
	Port     int
	Username string
	Password string
	AppName  string
}

// EmailStrategy: Implement NotificationStrategy với template
type EmailStrategy struct {
	From      string
	Host      string
	Port      int
	Username  string
	Password  string
	AppName   string
	Templates *template.Template // Loaded templates
}

func NewEmailStrategy(config EmailConfig) (*EmailStrategy, error) {
	tmpl := template.New("email")

	// Parse base template
	basePath := "internal/notifier/templates/base.html"
	tmpl, err := tmpl.ParseFiles(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base template: %w", err)
	}

	// Parse tất cả template con (otp.html, welcome.html,...)
	templateDir := "internal/notifier/templates/"
	files, err := filepath.Glob(templateDir + "[!base]*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to glob templates: %w", err)
	}
	tmpl, err = tmpl.ParseFiles(files...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return &EmailStrategy{
		From:      config.From,
		Host:      config.Host,
		Port:      config.Port,
		Username:  config.Username,
		Password:  config.Password,
		AppName:   config.AppName,
		Templates: tmpl,
	}, nil
}

func (s *EmailStrategy) Name() string {
	return "email"
}

func (s *EmailStrategy) Send(ctx context.Context, msg NotificationMessage) error {
	// Chuẩn bị data cho template
	data := struct {
		Title     string
		AppName   string
		UserName  string
		OTP       string
		LoginURL  string
		ResetLink string
	}{
		AppName: s.AppName,
	}

	// Xử lý theo type
	var tplName, title string
	switch msg.Type {
	case MessageTypeOTP:
		tplName = "otp.html"
		title = "Mã xác minh của bạn"
		data.Title = title
		data.UserName = msg.Data["user_name"].(string)
		data.OTP = msg.Data["otp"].(string)

	case MessageTypeWelcome:
		tplName = "welcome.html"
		title = "Chào mừng bạn!"
		data.Title = title
		data.UserName = msg.Data["user_name"].(string)
		data.LoginURL = "https://yourapp.com/login" // Có thể từ config hoặc data

	default:
		return fmt.Errorf("email strategy không hỗ trợ type: %s", msg.Type)
	}

	// Render template
	var body bytes.Buffer
	if err := s.Templates.ExecuteTemplate(&body, tplName, data); err != nil {
		return fmt.Errorf("failed to render template %s: %w", tplName, err)
	}

	// // Gửi email HTML
	// m := mail.NewMessage()
	// m.SetHeader("From", s.From)
	// m.SetHeader("To", msg.To)
	// m.SetHeader("Subject", title)
	// m.SetBody("text/html", body.String())

	// d := mail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Dev only, prod dùng cert proper

	// if err := d.DialAndSend(m); err != nil {
	// 	return fmt.Errorf("failed to send email: %w", err)
	// }
	return nil
}
