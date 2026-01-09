package handler

import (
	"OpenFabControl/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"math/rand"
	"time"
	"os"
)


func sendConfirmationEmail(recipientEmail string, verificationLink string) error {
	// SMTP Credentials (from .env)
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	// Construct SMTP authentication
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Email message
	from := smtpUsername
	to := []string{recipientEmail}
	subject := "Account Confirmation"
	body := fmt.Sprintf(`
Dear User,

Your account has been created. To activate your account, please click the following link:

%s

Thank you!
`, verificationLink)

	message := strings.Join([]string{
		"From: " + from + "\r\n",
		"To: " + strings.Join(to, ", ") + "\r\n",
		"Subject: " + subject + "\r\n",
		"\r\n",
		body + "\r\n"}, "")

	// Connect to the SMTP server
	addr := smtpHost + ":" + smtpPort
	conn, err := smtp.Dial(addr)
	if err != nil {
		log.Printf("Error connecting to SMTP server: %v", err)
		return err
	}
	defer conn.Close()

	// Authenticate
	if err := conn.Auth(auth); err != nil {
		log.Printf("Error authenticating: %v", err)
		return err
	}

	// Set the sender and recipient
	if err := conn.Mail(from); err != nil {
		log.Printf("Error setting sender: %v", err)
		return err
	}
	for _, recipient := range to {
		if err := conn.Rcpt(recipient); err != nil {
			log.Printf("Error setting recipient: %v", err)
			return err
		}
	}

	// Send the email body
	w, err := conn.Data()
	if err != nil {
		log.Printf("Error creating data writer: %v", err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Printf("Error writing message: %v", err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Printf("Error closing writer: %v", err)
		return err
	}

	// Close the connection
	err = conn.Quit()
	if err != nil {
		log.Printf("Error quitting connection: %v", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}



func Create_user(w http.ResponseWriter, r* http.Request) {
	if reject_all_methode_exept(r, w, http.MethodPost) != nil {
		return
	}

	var payload struct {
		ACCESS_KEY	string `json:"access_key"`
		EMAIL		string `json:"email"`
	}

	if extract_payload_data(r, w, &payload) != nil {
		return
	}

	if payload.ACCESS_KEY == "" {
		http.Error(w, "invalid payload: access key cannot be empty", http.StatusBadRequest)
		return
	}
	if payload.EMAIL == "" {
		http.Error(w, "invalid payload: email cannot be empty", http.StatusBadRequest)
		return
	}

	// check if email already registered
	var existing_email string
	err := database.Self.QueryRow(`SELECT email FROM users WHERE email = $1`, payload.EMAIL).Scan(&existing_email)
	if err == nil {
		http.Error(w, "Email already registered", http.StatusBadRequest)
		return
	} else if err != sql.ErrNoRows {
		log.Printf("db query error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// generate the verif code
	seeder := 0
	p_seeder := &seeder
	rand.Seed(int64(*p_seeder)) // use the address of the int as seed (yeah I know what you think, answer: why not)
 	randomNum := rand.Intn(999999)
	verif_code := fmt.Sprintf("%v-%v", time.Now().UnixMilli(), randomNum)

	// insert user in the DB
	query := `INSERT INTO users (access_key, email, password, verification_code) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING`
	_, err = database.Self.Exec(query, payload.ACCESS_KEY, payload.EMAIL, "UNDEFINED", verif_code)
	if err != nil {
		log.Printf("db insert error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// send email to create password
	// if err = sendConfirmationEmail(payload.EMAIL, os.Getenv("DOMAIN_NAME") + "/confirm-email?code=" + verif_code); err != nil {
	// 	http.Error(w, fmt.Sprintf("Error sending the email: %v", err), http.StatusInternalServerError)
	// }

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	// do NOT leave the confirm link here
	json.NewEncoder(w).Encode(map[string]any{"msg": "user created, email sent", "link_sent": os.Getenv("DOMAIN_NAME") + "/confirm-email?code=" + verif_code})
}
