package env

import (
	"os"
	"strconv"
)

type ImapConnection struct {
	ImapAdress string
	Username   string
	Password   string
	//emailAdressFrom string
	//smtpAdress      string
	//smtpPort        string // Added smtpPort to store the SMTP port
}

type SmtpConnection struct {
	SmtpAdress string
	SmtpPort   string // Added smtpPort to store the SMTP port
	Username   string
	Password   string
	//emailAdressFrom string
	//smtpAdress      string
}

type ValkeyConnection struct {
	Host string
	Port int
}

// var (
// 	// Ctx context.Context

// 	GmailImapConnectionCredentials = imapConnection{
// 		imapAdress: os.Getenv("GmailImapServerAdress"), // Get IMAP address from environment variable
// 		username:   os.Getenv("GmailServerUsername"),   // Get IMAP username from environment variable
// 		password:   os.Getenv("GmailServerPassword"),   // Get IMAP password from environment variable
// 	}

// 	GmailSmtpConnectionCredentials = smtpConnection{
// 		smtpAdress: os.Getenv("GmailSMTPServerAdress"), // Get SMTP address from environment variable
// 		smtpPort:   os.Getenv("GmailSMTPServerPort"),   // Get SMTP port from environment variable
// 		username:   os.Getenv("GmailServerUsername"),   // Assuming the same username as IMAP
// 		password:   os.Getenv("GmailServerPassword"),   // Assuming the same password as IMAP
// 	}
// )

func GetGmailImapConnectionCredentials() ImapConnection {
	return ImapConnection{
		ImapAdress: os.Getenv("GmailImapServerAdress"), // Get IMAP address from environment variable
		Username:   os.Getenv("GmailServerUsername"),   // Get IMAP username from environment variable
		Password:   os.Getenv("GmailServerPassword"),   // Get IMAP password from environment variable
	}
}

func GetGmailSmtpConnectionCredentials() SmtpConnection {
	return SmtpConnection{
		SmtpAdress: os.Getenv("GmailSMTPServerAdress"), // Get SMTP address from environment variable
		SmtpPort:   os.Getenv("GmailSMTPServerPort"),   // Get SMTP port from environment variable
		Username:   os.Getenv("GmailServerUsername"),   // Assuming the same username as IMAP
		Password:   os.Getenv("GmailServerPassword"),   // Assuming the same password as IMAP
	}
}

func GetValkeyConnectionCredentials() ValkeyConnection {

	portStr := os.Getenv("ValkeyPort")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 6379 // default port
	}

	host := os.Getenv("ValkeyHost")
	if host == "" {
		// host = "localhost" // default host
		host = "valkey-service.valkey-queue.svc.cluster.local" // default host in kubernetes
	}

	return ValkeyConnection{
		Host: host,
		Port: port,
	}

}

// GetWhiteListForProducer returns the whitelist for the producer
// Returns a one email as string from environment variable "ProducerWhiteListEmail"
func GetWhiteListForProducer() string {
	return os.Getenv("ProducerWhiteListEmail")
}
