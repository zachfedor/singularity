// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// UpdateWalletReader is a Reader for the UpdateWallet structure.
type UpdateWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateWalletNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateWalletInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /wallet/{address}/update] UpdateWallet", response, response.Code())
	}
}

// NewUpdateWalletOK creates a UpdateWalletOK with default headers values
func NewUpdateWalletOK() *UpdateWalletOK {
	return &UpdateWalletOK{}
}

/*
UpdateWalletOK describes a response with status code 200, with default header values.

OK
*/
type UpdateWalletOK struct {
	Payload *models.ModelWallet
}

// IsSuccess returns true when this update wallet o k response has a 2xx status code
func (o *UpdateWalletOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update wallet o k response has a 3xx status code
func (o *UpdateWalletOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update wallet o k response has a 4xx status code
func (o *UpdateWalletOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update wallet o k response has a 5xx status code
func (o *UpdateWalletOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update wallet o k response a status code equal to that given
func (o *UpdateWalletOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update wallet o k response
func (o *UpdateWalletOK) Code() int {
	return 200
}

func (o *UpdateWalletOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletOK %s", 200, payload)
}

func (o *UpdateWalletOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletOK %s", 200, payload)
}

func (o *UpdateWalletOK) GetPayload() *models.ModelWallet {
	return o.Payload
}

func (o *UpdateWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelWallet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWalletBadRequest creates a UpdateWalletBadRequest with default headers values
func NewUpdateWalletBadRequest() *UpdateWalletBadRequest {
	return &UpdateWalletBadRequest{}
}

/*
UpdateWalletBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateWalletBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this update wallet bad request response has a 2xx status code
func (o *UpdateWalletBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update wallet bad request response has a 3xx status code
func (o *UpdateWalletBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update wallet bad request response has a 4xx status code
func (o *UpdateWalletBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update wallet bad request response has a 5xx status code
func (o *UpdateWalletBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update wallet bad request response a status code equal to that given
func (o *UpdateWalletBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update wallet bad request response
func (o *UpdateWalletBadRequest) Code() int {
	return 400
}

func (o *UpdateWalletBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletBadRequest %s", 400, payload)
}

func (o *UpdateWalletBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletBadRequest %s", 400, payload)
}

func (o *UpdateWalletBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *UpdateWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWalletNotFound creates a UpdateWalletNotFound with default headers values
func NewUpdateWalletNotFound() *UpdateWalletNotFound {
	return &UpdateWalletNotFound{}
}

/*
UpdateWalletNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateWalletNotFound struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this update wallet not found response has a 2xx status code
func (o *UpdateWalletNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update wallet not found response has a 3xx status code
func (o *UpdateWalletNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update wallet not found response has a 4xx status code
func (o *UpdateWalletNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update wallet not found response has a 5xx status code
func (o *UpdateWalletNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update wallet not found response a status code equal to that given
func (o *UpdateWalletNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update wallet not found response
func (o *UpdateWalletNotFound) Code() int {
	return 404
}

func (o *UpdateWalletNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletNotFound %s", 404, payload)
}

func (o *UpdateWalletNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletNotFound %s", 404, payload)
}

func (o *UpdateWalletNotFound) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *UpdateWalletNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWalletInternalServerError creates a UpdateWalletInternalServerError with default headers values
func NewUpdateWalletInternalServerError() *UpdateWalletInternalServerError {
	return &UpdateWalletInternalServerError{}
}

/*
UpdateWalletInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateWalletInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this update wallet internal server error response has a 2xx status code
func (o *UpdateWalletInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update wallet internal server error response has a 3xx status code
func (o *UpdateWalletInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update wallet internal server error response has a 4xx status code
func (o *UpdateWalletInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update wallet internal server error response has a 5xx status code
func (o *UpdateWalletInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update wallet internal server error response a status code equal to that given
func (o *UpdateWalletInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update wallet internal server error response
func (o *UpdateWalletInternalServerError) Code() int {
	return 500
}

func (o *UpdateWalletInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletInternalServerError %s", 500, payload)
}

func (o *UpdateWalletInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wallet/{address}/update][%d] updateWalletInternalServerError %s", 500, payload)
}

func (o *UpdateWalletInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *UpdateWalletInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
