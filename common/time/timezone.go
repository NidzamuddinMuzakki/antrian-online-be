package time

import (
	"os"
	"time"

	"antrian-golang/common/constant"
)

// LoadTimeZoneFromEnv load timezone from env
// if no default time from env it will return constant.DefaultTimeZone
func LoadTimeZoneFromEnv() string {
	tz := os.Getenv(constant.Timezone)
	if len(tz) <= 0 {
		return constant.DefaultTimeZone
	}
	return tz
}

func LoadTimeZoneAsiaJakarta() *time.Location {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		// we dont want this error to happen
		//
		// this should be never called
		panic(err)
	}
	return loc
}
