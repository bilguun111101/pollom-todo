package shared

import "time"

func GetMillis() int64 {
	return GetMillisForTime(time.Now())
}

func GetMillisForTime(_time time.Time) int64 {
	return _time.UnixMilli()
}

func GetTimeForMillis(millis int64) time.Time {
	return time.UnixMilli(millis)
}

type AppError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`
	DetailedError string `json:"detailed_error"`
	RequestId     string `json:"request_id,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
	Where         string `json:"-"`
	params        map[string]any
	wrapped       error
}

func NewAppError(where string, id string, params map[string]any, details string, status int) *AppError {
	ap := &AppError{
		Id:            id,
		params:        params,
		Message:       id,
		Where:         where,
		DetailedError: details,
		StatusCode:    status,
	}

	return ap
}
