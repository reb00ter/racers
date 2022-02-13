package errors

const (
	InternalError = "internalError"
	RacerNotFound = "racerNotFound"
	BadRequest    = "BadRequest"
)

var errorMessage = map[string]string{
	"internalError": "an internal error occured",
	"racerNotFound": "racer could not be found",
	"BadRequest":    "check your data and try again",
}

// Booms can contain multiple boom errors
type Booms struct {
	Errors []Boom `json:"errors"`
}

func (b *Booms) Add(e Boom) {
	b.Errors = append(b.Errors, e)
}

func NewBooms() Booms {
	return Booms{}
}

// Boom represent the basic structure of a json error
type Boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func NewBoom(code, msg string, details interface{}) Boom {
	return Boom{Code: code, Message: msg, Details: details}
}

func ErrorText(code string) string {
	return errorMessage[code]
}
