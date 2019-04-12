// Code generated by mockery v1.0.0. DO NOT EDIT.

package api

import devicewallet "github.com/skycoin/hardware-wallet-go/src/device-wallet"
import messages "github.com/skycoin/hardware-wallet-protob/go"
import mock "github.com/stretchr/testify/mock"
import wire "github.com/skycoin/hardware-wallet-go/src/device-wallet/wire"

// MockGatewayer is an autogenerated mock type for the Gatewayer type
type MockGatewayer struct {
	mock.Mock
}

// AddressGen provides a mock function with given fields: addressN, startIndex, confirmAddress
func (_m *MockGatewayer) AddressGen(addressN int, startIndex int, confirmAddress bool) (wire.Message, error) {
	ret := _m.Called(addressN, startIndex, confirmAddress)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(int, int, bool) wire.Message); ok {
		r0 = rf(addressN, startIndex, confirmAddress)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, bool) error); ok {
		r1 = rf(addressN, startIndex, confirmAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplySettings provides a mock function with given fields: usePassphrase, label, language
func (_m *MockGatewayer) ApplySettings(usePassphrase bool, label string, language string) (wire.Message, error) {
	ret := _m.Called(usePassphrase, label, language)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(bool, string, string) wire.Message); ok {
		r0 = rf(usePassphrase, label, language)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool, string, string) error); ok {
		r1 = rf(usePassphrase, label, language)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backup provides a mock function with given fields:
func (_m *MockGatewayer) Backup() (wire.Message, error) {
	ret := _m.Called()

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func() wire.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ButtonAck provides a mock function with given fields:
func (_m *MockGatewayer) ButtonAck() (wire.Message, error) {
	ret := _m.Called()

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func() wire.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cancel provides a mock function with given fields:
func (_m *MockGatewayer) Cancel() (wire.Message, error) {
	ret := _m.Called()

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func() wire.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangePin provides a mock function with given fields:
func (_m *MockGatewayer) ChangePin() (wire.Message, error) {
	ret := _m.Called()

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func() wire.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckMessageSignature provides a mock function with given fields: message, signature, address
func (_m *MockGatewayer) CheckMessageSignature(message string, signature string, address string) (wire.Message, error) {
	ret := _m.Called(message, signature, address)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(string, string, string) wire.Message); ok {
		r0 = rf(message, signature, address)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(message, signature, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Connected provides a mock function with given fields:
func (_m *MockGatewayer) Connected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FirmwareUpload provides a mock function with given fields: payload, hash
func (_m *MockGatewayer) FirmwareUpload(payload []byte, hash [32]byte) error {
	ret := _m.Called(payload, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, [32]byte) error); ok {
		r0 = rf(payload, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateMnemonic provides a mock function with given fields: wordCount, usePassphrase
func (_m *MockGatewayer) GenerateMnemonic(wordCount uint32, usePassphrase bool) (wire.Message, error) {
	ret := _m.Called(wordCount, usePassphrase)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(uint32, bool) wire.Message); ok {
		r0 = rf(wordCount, usePassphrase)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32, bool) error); ok {
		r1 = rf(wordCount, usePassphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeatures provides a mock function with given fields:
func (_m *MockGatewayer) GetFeatures() (wire.Message, error) {
	ret := _m.Called()

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func() wire.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PassphraseAck provides a mock function with given fields: passphrase
func (_m *MockGatewayer) PassphraseAck(passphrase string) (wire.Message, error) {
	ret := _m.Called(passphrase)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(string) wire.Message); ok {
		r0 = rf(passphrase)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(passphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PinMatrixAck provides a mock function with given fields: p
func (_m *MockGatewayer) PinMatrixAck(p string) (wire.Message, error) {
	ret := _m.Called(p)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(string) wire.Message); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Recovery provides a mock function with given fields: wordCount, usePassphrase, dryRun
func (_m *MockGatewayer) Recovery(wordCount uint32, usePassphrase bool, dryRun bool) (wire.Message, error) {
	ret := _m.Called(wordCount, usePassphrase, dryRun)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(uint32, bool, bool) wire.Message); ok {
		r0 = rf(wordCount, usePassphrase, dryRun)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32, bool, bool) error); ok {
		r1 = rf(wordCount, usePassphrase, dryRun)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAutoPressButton provides a mock function with given fields: simulateButtonPress, simulateButtonType
func (_m *MockGatewayer) SetAutoPressButton(simulateButtonPress bool, simulateButtonType devicewallet.ButtonType) error {
	ret := _m.Called(simulateButtonPress, simulateButtonType)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, devicewallet.ButtonType) error); ok {
		r0 = rf(simulateButtonPress, simulateButtonType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMnemonic provides a mock function with given fields: mnemonic
func (_m *MockGatewayer) SetMnemonic(mnemonic string) (wire.Message, error) {
	ret := _m.Called(mnemonic)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(string) wire.Message); ok {
		r0 = rf(mnemonic)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(mnemonic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignMessage provides a mock function with given fields: addressIndex, message
func (_m *MockGatewayer) SignMessage(addressIndex int, message string) (wire.Message, error) {
	ret := _m.Called(addressIndex, message)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(int, string) wire.Message); ok {
		r0 = rf(addressIndex, message)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(addressIndex, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionSign provides a mock function with given fields: inputs, outputs
func (_m *MockGatewayer) TransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput) (wire.Message, error) {
	ret := _m.Called(inputs, outputs)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func([]*messages.SkycoinTransactionInput, []*messages.SkycoinTransactionOutput) wire.Message); ok {
		r0 = rf(inputs, outputs)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*messages.SkycoinTransactionInput, []*messages.SkycoinTransactionOutput) error); ok {
		r1 = rf(inputs, outputs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Wipe provides a mock function with given fields:
func (_m *MockGatewayer) Wipe() (wire.Message, error) {
	ret := _m.Called()

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func() wire.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WordAck provides a mock function with given fields: word
func (_m *MockGatewayer) WordAck(word string) (wire.Message, error) {
	ret := _m.Called(word)

	var r0 wire.Message
	if rf, ok := ret.Get(0).(func(string) wire.Message); ok {
		r0 = rf(word)
	} else {
		r0 = ret.Get(0).(wire.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(word)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
