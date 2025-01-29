package helper

import (
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

func ComparePassword(hashedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

func StringISOToDateTime(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}

func StringToUint(value string) (uint, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(result), nil
}

// HELPER FOR SESSION AND PROPOSAL USECASE

type SessionDates struct {
	RegistrationStart time.Time
	RegistrationEnd   time.Time
	SessionStart      time.Time
	SessionEnd        time.Time
}

func ValidateDates(dates SessionDates) error {
	if dates.RegistrationStart.Before(time.Now()) {
		return errors.New("registration start date should be after today")
	}

	if dates.RegistrationStart.After(dates.RegistrationEnd) {
		return errors.New("registration start date should be before the registration end date")
	}

	if dates.SessionStart.Before(dates.RegistrationEnd) {
		return errors.New("session start date should be after the registration end date")
	}

	if dates.SessionStart.After(dates.SessionEnd) {
		return errors.New("session start date should be before the session end date")
	}

	return nil
}

func ParseDatesFromRequest(
	RegistrationStarts, RegistrationEnds, SessionStarts, SessionEnds string,
) (SessionDates, error) {
	registrationStart, err := StringISOToDateTime(RegistrationStarts)
	if err != nil {
		return SessionDates{}, err
	}

	registrationEnd, err := StringISOToDateTime(RegistrationEnds)
	if err != nil {
		return SessionDates{}, err
	}

	sessionStart, err := StringISOToDateTime(SessionStarts)
	if err != nil {
		return SessionDates{}, err
	}

	sessionEnd, err := StringISOToDateTime(SessionEnds)
	if err != nil {
		return SessionDates{}, err
	}

	return SessionDates{
		RegistrationStart: registrationStart,
		RegistrationEnd:   registrationEnd,
		SessionStart:      sessionStart,
		SessionEnd:        sessionEnd,
	}, nil
}
