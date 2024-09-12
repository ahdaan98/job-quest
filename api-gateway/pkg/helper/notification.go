package helper

import (
	"fmt"
	"log"

	"github.com/ahdaan67/JobQuest/pkg/config"
	"gopkg.in/gomail.v2"
)

func SendNotification(email, jobTitle string, cfg config.Config) error {
    // Check if email and password are set
    if cfg.Email == "" || cfg.Password == "" {
        return fmt.Errorf("email configuration is missing")
    }
    
    // Check if recipient email is valid
    if email == "" {
        return fmt.Errorf("recipient email is empty")
    }
    
    // Create a new message
    m := gomail.NewMessage()
    m.SetHeader("From", cfg.Email)
    m.SetHeader("To", email)
    m.SetHeader("Subject", "Job Application Accepted")
    
    // HTML email body
    body := fmt.Sprintf(`
        <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    margin: 20px;
                    color: #333;
                }
                .header {
                    background-color: #f8f9fa;
                    padding: 10px;
                    text-align: center;
                    border-bottom: 2px solid #e9ecef;
                }
                .content {
                    padding: 20px;
                    border: 1px solid #e9ecef;
                    border-radius: 5px;
                    background-color: #ffffff;
                }
                .footer {
                    text-align: center;
                    margin-top: 20px;
                    font-size: 0.9em;
                    color: #6c757d;
                }
            </style>
        </head>
        <body>
            <div class="header">
                <h2>Job Application Status Update</h2>
            </div>
            <div class="content">
                <p>The application for the job titled <strong>"%s"</strong> has been accepted.</p>
                <p>Thank you for your attention.</p>
                <p>Best regards,<br>The Team</p>
            </div>
            <div class="footer">
                <p>&copy; 2024 Company Name. All rights reserved.</p>
            </div>
        </body>
        </html>`, jobTitle)

    m.SetBody("text/html", body)

    // Create a new dialer
    d := gomail.NewDialer("smtp.gmail.com", 587, cfg.Email, cfg.Password)

    // Send the email
    if err := d.DialAndSend(m); err != nil {
        log.Printf("Error sending notification: %v", err)
        return err
    }

    return nil
}