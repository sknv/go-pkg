package problem

import (
	"fmt"
)

// Problem is the struct definition of a problem details object (see RFC-7807).
type Problem struct {
	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the problem when
	// de-referenced.
	Type string `json:"type"`

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the problem, except for
	// purposes of localization.
	Title string `json:"title"`

	// The HTTP status code for this occurrence of the problem.
	Status int `json:"status"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// A URI that identifies the specific occurrence of the problem. This URI
	// may or may not yield further information if de-referenced.
	Instance string `json:"instance,omitempty"`

	// Some additional information.
	Extensions interface{} `json:"extensions,omitempty"`
}

func New(status int, title string) *Problem {
	return &Problem{
		Status: status,
		Title:  title,
	}
}

const (
	_defaultType = "about:blank"
)

func (p *Problem) GetType() string {
	if p.Type == "" {
		return _defaultType
	}
	return p.Type
}

func (p *Problem) SetType(newType string) *Problem {
	p.Type = newType
	return p
}

func (p *Problem) GetTitle() string {
	return p.Title
}

func (p *Problem) SetTitle(title string) *Problem {
	p.Title = title
	return p
}

func (p *Problem) GetStatus() int {
	return p.Status
}

func (p *Problem) SetStatus(status int) *Problem {
	p.Status = status
	return p
}

func (p *Problem) GetDetail() string {
	return p.Detail
}

func (p *Problem) SetDetail(detail string) *Problem {
	p.Detail = detail
	return p
}

func (p *Problem) GetInstance() string {
	return p.Instance
}

func (p *Problem) SetInstance(instance string) *Problem {
	p.Instance = instance
	return p
}

func (p *Problem) GetExtensions() interface{} {
	return p.Extensions
}

func (p *Problem) SetExtensions(ext interface{}) *Problem {
	p.Extensions = ext
	return p
}

func (p *Problem) Error() string {
	return fmt.Sprintf("status = %d, title = %s, detail = %s", p.GetStatus(), p.GetTitle(), p.GetDetail())
}
