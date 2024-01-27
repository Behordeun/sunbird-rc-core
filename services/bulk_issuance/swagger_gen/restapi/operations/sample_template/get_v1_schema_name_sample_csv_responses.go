// Code generated by go-swagger; DO NOT EDIT.

package sample_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bulk_issuance/swagger_gen/models"
)

// GetV1SchemaNameSampleCsvOKCode is the HTTP code returned for type GetV1SchemaNameSampleCsvOK
const GetV1SchemaNameSampleCsvOKCode int = 200

/*
GetV1SchemaNameSampleCsvOK OK

swagger:response getV1SchemaNameSampleCsvOK
*/
type GetV1SchemaNameSampleCsvOK struct {
	/*

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload models.SampleTemplateResponse `json:"body,omitempty"`
}

// NewGetV1SchemaNameSampleCsvOK creates GetV1SchemaNameSampleCsvOK with default headers values
func NewGetV1SchemaNameSampleCsvOK() *GetV1SchemaNameSampleCsvOK {

	return &GetV1SchemaNameSampleCsvOK{}
}

// WithContentDisposition adds the contentDisposition to the get v1 schema name sample csv o k response
func (o *GetV1SchemaNameSampleCsvOK) WithContentDisposition(contentDisposition string) *GetV1SchemaNameSampleCsvOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get v1 schema name sample csv o k response
func (o *GetV1SchemaNameSampleCsvOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the get v1 schema name sample csv o k response
func (o *GetV1SchemaNameSampleCsvOK) WithPayload(payload models.SampleTemplateResponse) *GetV1SchemaNameSampleCsvOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 schema name sample csv o k response
func (o *GetV1SchemaNameSampleCsvOK) SetPayload(payload models.SampleTemplateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1SchemaNameSampleCsvOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1SchemaNameSampleCsvBadRequestCode is the HTTP code returned for type GetV1SchemaNameSampleCsvBadRequest
const GetV1SchemaNameSampleCsvBadRequestCode int = 400

/*
GetV1SchemaNameSampleCsvBadRequest Bad Request

swagger:response getV1SchemaNameSampleCsvBadRequest
*/
type GetV1SchemaNameSampleCsvBadRequest struct {
	/*

	  Default: "application/json"
	*/
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewGetV1SchemaNameSampleCsvBadRequest creates GetV1SchemaNameSampleCsvBadRequest with default headers values
func NewGetV1SchemaNameSampleCsvBadRequest() *GetV1SchemaNameSampleCsvBadRequest {

	var (
		// initialize headers with default values

		contentTypeDefault = string("application/json")
	)

	return &GetV1SchemaNameSampleCsvBadRequest{

		ContentType: contentTypeDefault,
	}
}

// WithContentType adds the contentType to the get v1 schema name sample csv bad request response
func (o *GetV1SchemaNameSampleCsvBadRequest) WithContentType(contentType string) *GetV1SchemaNameSampleCsvBadRequest {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get v1 schema name sample csv bad request response
func (o *GetV1SchemaNameSampleCsvBadRequest) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get v1 schema name sample csv bad request response
func (o *GetV1SchemaNameSampleCsvBadRequest) WithPayload(payload *models.ErrorPayload) *GetV1SchemaNameSampleCsvBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 schema name sample csv bad request response
func (o *GetV1SchemaNameSampleCsvBadRequest) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1SchemaNameSampleCsvBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1SchemaNameSampleCsvNotFoundCode is the HTTP code returned for type GetV1SchemaNameSampleCsvNotFound
const GetV1SchemaNameSampleCsvNotFoundCode int = 404

/*
GetV1SchemaNameSampleCsvNotFound Not found

swagger:response getV1SchemaNameSampleCsvNotFound
*/
type GetV1SchemaNameSampleCsvNotFound struct {
	/*

	  Default: "application/json"
	*/
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewGetV1SchemaNameSampleCsvNotFound creates GetV1SchemaNameSampleCsvNotFound with default headers values
func NewGetV1SchemaNameSampleCsvNotFound() *GetV1SchemaNameSampleCsvNotFound {

	var (
		// initialize headers with default values

		contentTypeDefault = string("application/json")
	)

	return &GetV1SchemaNameSampleCsvNotFound{

		ContentType: contentTypeDefault,
	}
}

// WithContentType adds the contentType to the get v1 schema name sample csv not found response
func (o *GetV1SchemaNameSampleCsvNotFound) WithContentType(contentType string) *GetV1SchemaNameSampleCsvNotFound {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get v1 schema name sample csv not found response
func (o *GetV1SchemaNameSampleCsvNotFound) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get v1 schema name sample csv not found response
func (o *GetV1SchemaNameSampleCsvNotFound) WithPayload(payload *models.ErrorPayload) *GetV1SchemaNameSampleCsvNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 schema name sample csv not found response
func (o *GetV1SchemaNameSampleCsvNotFound) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1SchemaNameSampleCsvNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
