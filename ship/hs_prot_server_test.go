package ship

import (
	"testing"

	"github.com/enbility/ship-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestProServerSuite(t *testing.T) {
	suite.Run(t, new(ProServerSuite))
}

type ProServerSuite struct {
	suite.Suite
	role shipRole
}

func (s *ProServerSuite) BeforeTest(suiteName, testName string) {
	s.role = ShipRoleServer
}

func (s *ProServerSuite) Test_Init() {
	sut, data := initTest(s.role)

	sut.setState(model.SmeHelloStateOk, nil)

	sut.handleState(false, nil)

	assert.Equal(s.T(), true, sut.handshakeTimerRunning)

	// the state goes from smeHelloStateOk to smeProtHStateServerInit to smeProtHStateServerListenProposal
	assert.Equal(s.T(), model.SmeProtHStateServerListenProposal, sut.getState())
	assert.Nil(s.T(), data.lastMessage())

	shutdownTest(sut)
}

func (s *ProServerSuite) Test_ListenProposal() {
	sut, data := initTest(s.role)

	sut.setState(model.SmeProtHStateServerListenProposal, nil)

	protMsg := model.MessageProtocolHandshake{
		MessageProtocolHandshake: model.MessageProtocolHandshakeType{
			HandshakeType: model.ProtocolHandshakeTypeTypeAnnounceMax,
			Version:       model.Version{Major: 1, Minor: 0},
			Formats: model.MessageProtocolFormatsType{
				Format: []model.MessageProtocolFormatType{model.MessageProtocolFormatTypeUTF8},
			},
		},
	}

	msg, err := sut.shipMessage(model.MsgTypeControl, protMsg)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), msg)

	sut.handleState(false, msg)

	assert.Equal(s.T(), true, sut.handshakeTimerRunning)

	assert.Equal(s.T(), model.SmeProtHStateServerListenConfirm, sut.getState())
	assert.NotNil(s.T(), data.lastMessage())

	shutdownTest(sut)
}

func (s *ProServerSuite) Test_ListenProposal_Failure() {
	sut, _ := initTest(s.role)

	sut.setState(model.SmeProtHStateServerListenProposal, nil)

	protMsg := model.MessageProtocolHandshake{
		MessageProtocolHandshake: model.MessageProtocolHandshakeType{
			HandshakeType: model.ProtocolHandshakeTypeTypeSelect,
		},
	}

	msg, err := sut.shipMessage(model.MsgTypeControl, protMsg)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), msg)

	sut.handleState(false, msg)

	assert.Equal(s.T(), false, sut.handshakeTimerRunning)

	assert.Equal(s.T(), model.SmeStateError, sut.getState())

	shutdownTest(sut)
}

func (s *ProServerSuite) Test_ListenConfirm() {
	sut, data := initTest(s.role)

	sut.setState(model.SmeProtHStateServerListenConfirm, nil)

	protMsg := model.MessageProtocolHandshake{
		MessageProtocolHandshake: model.MessageProtocolHandshakeType{
			HandshakeType: model.ProtocolHandshakeTypeTypeSelect,
			Version:       model.Version{Major: 1, Minor: 0},
			Formats: model.MessageProtocolFormatsType{
				Format: []model.MessageProtocolFormatType{model.MessageProtocolFormatTypeUTF8},
			},
		},
	}

	msg, err := sut.shipMessage(model.MsgTypeControl, protMsg)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), msg)

	sut.handleState(false, msg)

	assert.Equal(s.T(), false, sut.handshakeTimerRunning)

	// state smeProtHStateServerOk directly goes to smePinStateCheckInit to smePinStateCheckListen
	assert.Equal(s.T(), model.SmePinStateCheckListen, sut.getState())
	assert.NotNil(s.T(), data.lastMessage())

	shutdownTest(sut)
}

func (s *ProServerSuite) Test_ListenConfirm_Failures() {
	sut, data := initTest(s.role)

	sut.setState(model.SmeProtHStateServerListenConfirm, nil)

	protMsg := model.MessageProtocolHandshake{
		MessageProtocolHandshake: model.MessageProtocolHandshakeType{
			HandshakeType: model.ProtocolHandshakeTypeTypeAnnounceMax,
		},
	}

	msg, err := sut.shipMessage(model.MsgTypeControl, protMsg)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), msg)

	sut.handleState(false, msg)

	assert.Equal(s.T(), false, sut.handshakeTimerRunning)

	assert.Equal(s.T(), model.SmeStateError, sut.getState())
	assert.NotNil(s.T(), data.lastMessage())

	shutdownTest(sut)
}
