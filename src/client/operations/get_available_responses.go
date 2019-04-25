// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/skycoin/hardware-wallet-daemon/src/models"
)

// GetAvailableReader is a Reader for the GetAvailable structure.
type GetAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAvailableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAvailableOK creates a GetAvailableOK with default headers values
func NewGetAvailableOK() *GetAvailableOK {
	return &GetAvailableOK{}
}

/*GetAvailableOK handles this case with default header values.

success
*/
type GetAvailableOK struct {
	Payload *GetAvailableOKBody
}

func (o *GetAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAvailableOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableDefault creates a GetAvailableDefault with default headers values
func NewGetAvailableDefault(code int) *GetAvailableDefault {
	return &GetAvailableDefault{
		_statusCode: code,
	}
}

/*GetAvailableDefault handles this case with default header values.

error
*/
type GetAvailableDefault struct {
	_statusCode int

	Payload *models.HTTPErrorResponse
}

// Code gets the status code for the get available default response
func (o *GetAvailableDefault) Code() int {
	return o._statusCode
}

func (o *GetAvailableDefault) Error() string {
	return o.Payload.Error.Message
}

func (o *GetAvailableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAvailableOKBody get available o k body
swagger:model GetAvailableOKBody
*/
type GetAvailableOKBody struct {

	// data
	Data bool `json:"data,omitempty"`
}

// Validate validates this get available o k body
func (o *GetAvailableOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAvailableOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAvailableOKBody) UnmarshalBinary(b []byte) error {
	var res GetAvailableOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}