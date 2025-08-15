package rabbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	mockDB := InitMock()
	mockDB.Connect.On("Connect").Return(nil)

	err := mockDB.Connect.Connect()

	assert.NoError(t, err)
	mockDB.Connect.AssertExpectations(t)
}

func TestChannel(t *testing.T) {
	mockDB := InitMock()
	mockDB.Connect.On("Channel").Return(nil)

	err := mockDB.Connect.Channel()

	assert.NoError(t, err)
	mockDB.Connect.AssertExpectations(t)
}

func TestCloseConnection(t *testing.T) {
	mockDB := InitMock()
	mockDB.Connect.On("CloseConnection").Return(nil)

	err := mockDB.Connect.CloseConnection()

	assert.NoError(t, err)
	mockDB.Connect.AssertExpectations(t)
}

func TestCloseChannel(t *testing.T) {
	mockDB := InitMock()
	mockDB.Connect.On("CloseChannel").Return(nil)

	err := mockDB.Connect.CloseChannel()

	assert.NoError(t, err)
	mockDB.Connect.AssertExpectations(t)
}

func TestIsConnected(t *testing.T) {
	mockDB := InitMock()
	mockDB.Connect.On("IsConnected").Return(nil)

	err := mockDB.Connect.IsConnected()

	assert.NoError(t, err)
	mockDB.Connect.AssertExpectations(t)
}
