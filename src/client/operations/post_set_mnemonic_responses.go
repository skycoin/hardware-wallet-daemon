// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/skycoin/hardware-wallet-daemon/src/models"
)

// PostSetMnemonicReader is a Reader for the PostSetMnemonic structure.
type PostSetMnemonicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSetMnemonicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostSetMnemonicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostSetMnemonicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSetMnemonicOK creates a PostSetMnemonicOK with default headers values
func NewPostSetMnemonicOK() *PostSetMnemonicOK {
	return &PostSetMnemonicOK{}
}

/*PostSetMnemonicOK handles this case with default header values.

success
*/
type PostSetMnemonicOK struct {
	Payload *models.HttpsuccessResponse
}

func (o *PostSetMnemonicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttpsuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSetMnemonicDefault creates a PostSetMnemonicDefault with default headers values
func NewPostSetMnemonicDefault(code int) *PostSetMnemonicDefault {
	return &PostSetMnemonicDefault{
		_statusCode: code,
	}
}

/*PostSetMnemonicDefault handles this case with default header values.

error
*/
type PostSetMnemonicDefault struct {
	_statusCode int

	Payload *models.HTTPErrorResponse
}

// Code gets the status code for the post set mnemonic default response
func (o *PostSetMnemonicDefault) Code() int {
	return o._statusCode
}

func (o *PostSetMnemonicDefault) Error() string {
	return o.Payload.Error.Message
}

func (o *PostSetMnemonicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
