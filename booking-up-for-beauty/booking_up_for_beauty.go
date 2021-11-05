package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	return ConvertToTime(layout, date)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	return time.Now().After(ConvertToTime(layout, date))
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	hour := ConvertToTime(layout, date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "Monday, January 2, 2006, at 15:04"
	return fmt.Sprintf("You have an appointment on %s.",
		Schedule(date).Format(layout))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	layout := "2006-01-2"
	anniversaryDate := fmt.Sprintf("%d-09-15", time.Now().Year())
	time, _ := time.Parse(layout, anniversaryDate)
	return time
}

func ConvertToTime(layout, date string) time.Time {
	var time, _ = time.Parse(layout, date)
	return time
}
