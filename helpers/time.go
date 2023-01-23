package helpers

import "time"

func ToHumanTimeStamp(timestamp time.Time) string {
	return timestamp.UTC().Format("02 Jan 2006 at 15:04")
}
