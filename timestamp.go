package whmcsgo

import (
	"strconv"
	"strings"
	"time"
)

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an RFC3339 or Unix timestamp. This is necessary for some
// fields since the GitHub API is inconsistent in how it represents times. All
// exported methods of time.Time can be called on Timestamp.
type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in RFC3339 or Unix format.
func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		(*t).Time = time.Unix(i, 0)
	} else {
		(*t).Time, err = time.Parse(`"`+time.RFC3339+`"`, str)
	}
	return
}

// Equal reports whether t and u are equal based on time.Equal
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}

// WHCMSdate allows the JSON string to be Unmarshaled
type WHCMSdate struct {
	time.Time
}

// UnmarshalJSON interface, we need a function UnmarshalJSON on the WHCMSdate type.
func (wd *WHCMSdate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)

	layout := "2006-01-02 -0700 MST"
	if strInput == "0000-00-00 00:00:00" {
		wd.Time = time.Time{}
		return nil
	}

	if len(strInput) > 10 {
		layout = "2006-01-02 15:04:05 -0700 MST"
	}

	newTime, err := time.Parse(layout, strInput+" +1000 UTC")
	if err != nil {
		wd.Time = time.Time{}
		return err
	}

	wd.Time = newTime
	return nil
}
