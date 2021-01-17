package views

const (
	AlertLvlError = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo = "info"
	AlertLvlSuccess = "success"
	AlertMsgGeneric = "Something went wrong. Please try again, and contact us if the problem persists."
)


// Alert is used to render bootstrap alert messages in templates
type Alert struct {
	Level   string
	Message string
}

// Data is the top level structure that views expect data to come in
type Data struct {
	Alert *Alert
	Yield interface{}
}

func (d *Data) SetAlert(err error) {
	if pErr, ok := err.(PublicError); ok {
		d.Alert = &Alert{
			Level: AlertLvlError,
			Message: pErr.Public(),
		}
	} else {
		d.Alert = &Alert{
			Level: AlertLvlError,
			Message: AlertMsgGeneric,
		}
	}
}

type PublicError interface {
	// Error() string
	error
	Public() string
}