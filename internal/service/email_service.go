package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

const (
	defaultTimeout     = 10 * time.Second
	linkExpiryMinutes = 15
)

type EmailConfig struct {
	Domain      string
	APIKey      string
	SenderEmail string
	APIURL      string
}

// EmailService handles all email-related operations
type EmailService struct {
	mg     mailgun.Mailgun
	config EmailConfig
}

// NewEmailService creates a new instance of EmailService
func NewEmailService() (*EmailService, error) {
	config := EmailConfig{
		Domain:      os.Getenv("MG_DOMAIN"),
		APIKey:      os.Getenv("MG_API_KEY"),
		SenderEmail: os.Getenv("MG_SENDER"),
		APIURL:      os.Getenv("APP_URL"),
	}

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid email service configuration: %w", err)
	}

	return &EmailService{
		mg:     mailgun.NewMailgun(config.Domain, config.APIKey),
		config: config,
	}, nil
}

func validateConfig(config EmailConfig) error {
	if config.Domain == "" {
		return fmt.Errorf("mailgun domain is required")
	}
	if config.APIKey == "" {
		return fmt.Errorf("mailgun API key is required")
	}
	if config.SenderEmail == "" {
		return fmt.Errorf("sender email is required")
	}
	if config.APIURL == "" {
		return fmt.Errorf("API URL is required")
	}
	return nil
}

type EmailData struct {
	Recipient    string
	Token        string
	Subject      string
	PlainText    string
	HTMLTemplate string
}

func (s *EmailService) sendEmail(ctx context.Context, data EmailData) error {
	message := s.mg.NewMessage(
		s.config.SenderEmail,
		data.Subject,
		data.PlainText,
		data.Recipient,
	)
	message.SetHTML(data.HTMLTemplate)

	// Set timeout for email sending
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, _, err := s.mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// SendEmailVerification sends an email verification link to the specified recipient
func (s *EmailService) SendEmailVerification(ctx context.Context, recipient, token string) error {
	verificationLink := fmt.Sprintf("https://%s/v1/verify-email?token=%s&email=%s", s.config.APIURL, token, recipient)

	htmlBody := fmt.Sprintf(`
        <div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
            <h2>Email Verification</h2>
            <p>Thank you for registering an account with Conflux! Click the link below to verify your email:</p>
            <p style="margin: 20px 0;">
                <a href="%s" style="background-color: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Verify Email</a>
            </p>
            <p>The verification link will expire in <b>%d minutes</b>.</p>
            <p>If you did not register for an account at Conflux, please ignore this message.</p>
        </div>
    `, verificationLink, linkExpiryMinutes)

	emailData := EmailData{
		Recipient:    recipient,
		Token:        token,
		Subject:      "Email Verification",
		PlainText:    fmt.Sprintf("Click the following link to verify your email: %s", verificationLink),
		HTMLTemplate: htmlBody,
	}

	return s.sendEmail(ctx, emailData)
}

// SendPasswordResetEmail sends a password reset link to the specified recipient
func (s *EmailService) SendPasswordResetEmail(ctx context.Context, recipient, token string) error {
	resetLink := fmt.Sprintf("http://%s/reset-password?token=%s&email=%s", s.config.APIURL, token, recipient)

	htmlBody := fmt.Sprintf(`
        <div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
            <h2>Password Reset</h2>
            <p>We received a request to reset your password. Click the link below to set a new password:</p>
            <p style="margin: 20px 0;">
                <a href="%s" style="background-color: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Reset Password</a>
            </p>
            <p>This link will expire in <b>%d minutes</b>.</p>
            <p>If you did not request a password reset for your Conflux account, please ignore this message.</p>
        </div>
    `, resetLink, linkExpiryMinutes)

	emailData := EmailData{
		Recipient:    recipient,
		Token:        token,
		Subject:      "Password Reset",
		PlainText:    fmt.Sprintf("Click the following link to reset your password: %s", resetLink),
		HTMLTemplate: htmlBody,
	}

	return s.sendEmail(ctx, emailData)
}
