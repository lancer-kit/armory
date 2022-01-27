package tools

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// URL is helper for configuration urls.
type URL struct {
	URL      *url.URL
	Str      string
	basePath string
}

const slash = "/"

// SetBasePath add standard path-prefix for URL.
func (j *URL) SetBasePath(path string) {
	j.basePath = path
}

// WithPath returns formatted URL to string with given path-suffix.
func (j *URL) WithPath(path string) string {
	ur := *j.URL
	ur.Path = j.basePath + slash + strings.TrimPrefix(path, slash)

	return ur.String()
}

// WithPathURL returns URL with sanitized and saved path-suffix.
func (j *URL) WithPathURL(path string) url.URL {
	ur := *j.URL
	ur.Path = j.basePath + slash + strings.TrimPrefix(path, slash)

	return ur
}

// UnmarshalYAML is an implementation of yaml.Unmarshaler.
func (j *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}

	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	j.Str = s
	j.URL = u
	j.basePath = strings.TrimSuffix(u.Path, slash)
	return err
}

// UnmarshalJSON is an implementation of json.Unmarshaler.
func (j *URL) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	j.Str = s
	j.URL = u
	j.basePath = strings.TrimSuffix(u.Path, slash)
	return nil
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (j URL) Validate() error {
	return validation.Validate(j.Str, validation.Required, is.URL)
}

// Required is a ozzo-validation rule
var Required = &requiredRule{message: "url cannot be blank", skipNil: false} // nolint:gochecknoglobals

type requiredRule struct {
	message string
	skipNil bool
}

// Validate checks if the given value is valid or not.
func (v *requiredRule) Validate(value interface{}) error {
	j, ok := value.(URL)
	if !ok {
		return errors.New("invalid type")
	}
	return j.Validate()
}

// Error sets the error message for the rule.
func (v *requiredRule) Error(message string) *requiredRule {
	return &requiredRule{
		message: message,
		skipNil: v.skipNil,
	}
}
