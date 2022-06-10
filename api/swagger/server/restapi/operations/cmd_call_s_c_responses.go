// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CmdCallSCOKCode is the HTTP code returned for type CmdCallSCOK
const CmdCallSCOKCode int = 200

/*CmdCallSCOK operation id.

swagger:response cmdCallSCOK
*/
type CmdCallSCOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCmdCallSCOK creates CmdCallSCOK with default headers values
func NewCmdCallSCOK() *CmdCallSCOK {

	return &CmdCallSCOK{}
}

// WithPayload adds the payload to the cmd call s c o k response
func (o *CmdCallSCOK) WithPayload(payload string) *CmdCallSCOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cmd call s c o k response
func (o *CmdCallSCOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CmdCallSCOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}