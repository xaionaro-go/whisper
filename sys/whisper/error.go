package whisper

import "errors"

var (
	ErrTranscriptionFailed = errors.New("whisper_full failed")
	ErrAutoDetectFailed    = errors.New("whisper_lang_auto_detect failed")
)

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return e.Message
}
