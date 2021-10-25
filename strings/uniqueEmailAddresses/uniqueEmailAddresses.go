package strings

import (
	"strings"
)

type Emails map[string]map[string]bool

type UniqueEmails struct {
	emails                          Emails
	numberOfReceivingEmailAddresses int
}

func (ue *UniqueEmails) processEmail(emailAddress string) {
	splittedEmail := strings.Split(emailAddress, "@")
	domainName := splittedEmail[1]
	localName := cleanLocalName(splittedEmail[0])

	if _, domainNameExistsInEmailMap := ue.emails[domainName]; !domainNameExistsInEmailMap {
		ue.emails[domainName] = map[string]bool{}
	}

	if _, localNameExistsInDomainMap := ue.emails[domainName][localName]; !localNameExistsInDomainMap {
		ue.emails[domainName][localName] = true
		ue.numberOfReceivingEmailAddresses++
	}
}

func cleanLocalName(localName string) string {
	localNameNoDots := strings.ReplaceAll(localName, ".", "")
	return strings.Split(localNameNoDots, "+")[0]
}

func numUniqueEmails(emails []string) int {
	if len(emails) == 1 {
		return 1
	}

	uniqueEmails := UniqueEmails{emails: make(Emails)}

	for i := 0; i < len(emails); i++ {
		emailAddress := emails[i]
		uniqueEmails.processEmail(emailAddress)
	}

	return uniqueEmails.numberOfReceivingEmailAddresses
}
