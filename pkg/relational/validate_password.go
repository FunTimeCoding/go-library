package relational

import (
	"log"
	"strings"
)

func (d *Database) validatePassword(p string) {
	if p == "" {
		log.Panicf("password cannot be empty")
	}

	if len(p) < 8 {
		log.Panicf("password must be at least 8 characters long")
	}

	if len(p) > 128 {
		log.Panicf("password cannot be longer than 128 characters")
	}

	if strings.Contains(p, "'") {
		log.Panicf("password cannot contain single quotes")
	}

	if strings.ToLower(p) == "password" ||
		strings.ToLower(p) == "123456" ||
		strings.ToLower(p) == "admin" {
		log.Panicf("password is too weak")
	}
}
