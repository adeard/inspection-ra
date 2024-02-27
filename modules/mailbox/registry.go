package mailbox

import "gorm.io/gorm"

func MailboxRegistry(db *gorm.DB) Service {
	mailboxRepository := NewRepository(db)
	mailboxService := NewService(mailboxRepository)

	return mailboxService
}
