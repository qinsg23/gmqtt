// Code generated by MockGen. DO NOT EDIT.
// Source: stats.go

// Package gmqtt is a generated GoMock package.
package gmqtt

import (
	packets "github.com/DrmagicE/gmqtt/pkg/packets"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStatsManager is a mock of StatsManager interface
type MockStatsManager struct {
	ctrl     *gomock.Controller
	recorder *MockStatsManagerMockRecorder
}

// MockStatsManagerMockRecorder is the mock recorder for MockStatsManager
type MockStatsManagerMockRecorder struct {
	mock *MockStatsManager
}

// NewMockStatsManager creates a new mock instance
func NewMockStatsManager(ctrl *gomock.Controller) *MockStatsManager {
	mock := &MockStatsManager{ctrl: ctrl}
	mock.recorder = &MockStatsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatsManager) EXPECT() *MockStatsManagerMockRecorder {
	return m.recorder
}

// packetReceived mocks base method
func (m *MockStatsManager) packetReceived(packet packets.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "packetReceived", packet)
}

// packetReceived indicates an expected call of packetReceived
func (mr *MockStatsManagerMockRecorder) packetReceived(packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "packetReceived", reflect.TypeOf((*MockStatsManager)(nil).packetReceived), packet)
}

// packetSent mocks base method
func (m *MockStatsManager) packetSent(packet packets.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "packetSent", packet)
}

// packetSent indicates an expected call of packetSent
func (mr *MockStatsManagerMockRecorder) packetSent(packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "packetSent", reflect.TypeOf((*MockStatsManager)(nil).packetSent), packet)
}

// addClientConnected mocks base method
func (m *MockStatsManager) addClientConnected() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addClientConnected")
}

// addClientConnected indicates an expected call of addClientConnected
func (mr *MockStatsManagerMockRecorder) addClientConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addClientConnected", reflect.TypeOf((*MockStatsManager)(nil).addClientConnected))
}

// addClientDisconnected mocks base method
func (m *MockStatsManager) addClientDisconnected() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addClientDisconnected")
}

// addClientDisconnected indicates an expected call of addClientDisconnected
func (mr *MockStatsManagerMockRecorder) addClientDisconnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addClientDisconnected", reflect.TypeOf((*MockStatsManager)(nil).addClientDisconnected))
}

// addSessionActive mocks base method
func (m *MockStatsManager) addSessionActive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addSessionActive")
}

// addSessionActive indicates an expected call of addSessionActive
func (mr *MockStatsManagerMockRecorder) addSessionActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addSessionActive", reflect.TypeOf((*MockStatsManager)(nil).addSessionActive))
}

// decSessionActive mocks base method
func (m *MockStatsManager) decSessionActive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "decSessionActive")
}

// decSessionActive indicates an expected call of decSessionActive
func (mr *MockStatsManagerMockRecorder) decSessionActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decSessionActive", reflect.TypeOf((*MockStatsManager)(nil).decSessionActive))
}

// addSessionInactive mocks base method
func (m *MockStatsManager) addSessionInactive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addSessionInactive")
}

// addSessionInactive indicates an expected call of addSessionInactive
func (mr *MockStatsManagerMockRecorder) addSessionInactive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addSessionInactive", reflect.TypeOf((*MockStatsManager)(nil).addSessionInactive))
}

// decSessionInactive mocks base method
func (m *MockStatsManager) decSessionInactive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "decSessionInactive")
}

// decSessionInactive indicates an expected call of decSessionInactive
func (mr *MockStatsManagerMockRecorder) decSessionInactive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decSessionInactive", reflect.TypeOf((*MockStatsManager)(nil).decSessionInactive))
}

// addSessionExpired mocks base method
func (m *MockStatsManager) addSessionExpired() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addSessionExpired")
}

// addSessionExpired indicates an expected call of addSessionExpired
func (mr *MockStatsManagerMockRecorder) addSessionExpired() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addSessionExpired", reflect.TypeOf((*MockStatsManager)(nil).addSessionExpired))
}

// messageDropped mocks base method
func (m *MockStatsManager) messageDropped(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageDropped", qos)
}

// messageDropped indicates an expected call of messageDropped
func (mr *MockStatsManagerMockRecorder) messageDropped(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageDropped", reflect.TypeOf((*MockStatsManager)(nil).messageDropped), qos)
}

// messageReceived mocks base method
func (m *MockStatsManager) messageReceived(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageReceived", qos)
}

// messageReceived indicates an expected call of messageReceived
func (mr *MockStatsManagerMockRecorder) messageReceived(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageReceived", reflect.TypeOf((*MockStatsManager)(nil).messageReceived), qos)
}

// messageSent mocks base method
func (m *MockStatsManager) messageSent(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageSent", qos)
}

// messageSent indicates an expected call of messageSent
func (mr *MockStatsManagerMockRecorder) messageSent(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageSent", reflect.TypeOf((*MockStatsManager)(nil).messageSent), qos)
}

// messageEnqueue mocks base method
func (m *MockStatsManager) messageEnqueue(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageEnqueue", delta)
}

// messageEnqueue indicates an expected call of messageEnqueue
func (mr *MockStatsManagerMockRecorder) messageEnqueue(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageEnqueue", reflect.TypeOf((*MockStatsManager)(nil).messageEnqueue), delta)
}

// messageDequeue mocks base method
func (m *MockStatsManager) messageDequeue(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageDequeue", delta)
}

// messageDequeue indicates an expected call of messageDequeue
func (mr *MockStatsManagerMockRecorder) messageDequeue(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageDequeue", reflect.TypeOf((*MockStatsManager)(nil).messageDequeue), delta)
}

// GetStats mocks base method
func (m *MockStatsManager) GetStats() *ServerStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(*ServerStats)
	return ret0
}

// GetStats indicates an expected call of GetStats
func (mr *MockStatsManagerMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockStatsManager)(nil).GetStats))
}

// MockSessionStatsManager is a mock of SessionStatsManager interface
type MockSessionStatsManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionStatsManagerMockRecorder
}

// MockSessionStatsManagerMockRecorder is the mock recorder for MockSessionStatsManager
type MockSessionStatsManagerMockRecorder struct {
	mock *MockSessionStatsManager
}

// NewMockSessionStatsManager creates a new mock instance
func NewMockSessionStatsManager(ctrl *gomock.Controller) *MockSessionStatsManager {
	mock := &MockSessionStatsManager{ctrl: ctrl}
	mock.recorder = &MockSessionStatsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionStatsManager) EXPECT() *MockSessionStatsManagerMockRecorder {
	return m.recorder
}

// messageDropped mocks base method
func (m *MockSessionStatsManager) messageDropped(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageDropped", qos)
}

// messageDropped indicates an expected call of messageDropped
func (mr *MockSessionStatsManagerMockRecorder) messageDropped(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageDropped", reflect.TypeOf((*MockSessionStatsManager)(nil).messageDropped), qos)
}

// messageReceived mocks base method
func (m *MockSessionStatsManager) messageReceived(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageReceived", qos)
}

// messageReceived indicates an expected call of messageReceived
func (mr *MockSessionStatsManagerMockRecorder) messageReceived(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageReceived", reflect.TypeOf((*MockSessionStatsManager)(nil).messageReceived), qos)
}

// messageSent mocks base method
func (m *MockSessionStatsManager) messageSent(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageSent", qos)
}

// messageSent indicates an expected call of messageSent
func (mr *MockSessionStatsManagerMockRecorder) messageSent(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageSent", reflect.TypeOf((*MockSessionStatsManager)(nil).messageSent), qos)
}

// messageEnqueue mocks base method
func (m *MockSessionStatsManager) messageEnqueue(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageEnqueue", delta)
}

// messageEnqueue indicates an expected call of messageEnqueue
func (mr *MockSessionStatsManagerMockRecorder) messageEnqueue(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageEnqueue", reflect.TypeOf((*MockSessionStatsManager)(nil).messageEnqueue), delta)
}

// messageDequeue mocks base method
func (m *MockSessionStatsManager) messageDequeue(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageDequeue", delta)
}

// messageDequeue indicates an expected call of messageDequeue
func (mr *MockSessionStatsManagerMockRecorder) messageDequeue(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageDequeue", reflect.TypeOf((*MockSessionStatsManager)(nil).messageDequeue), delta)
}

// addInflightCurrent mocks base method
func (m *MockSessionStatsManager) addInflightCurrent(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addInflightCurrent", delta)
}

// addInflightCurrent indicates an expected call of addInflightCurrent
func (mr *MockSessionStatsManagerMockRecorder) addInflightCurrent(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addInflightCurrent", reflect.TypeOf((*MockSessionStatsManager)(nil).addInflightCurrent), delta)
}

// decInflightCurrent mocks base method
func (m *MockSessionStatsManager) decInflightCurrent(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "decInflightCurrent", delta)
}

// decInflightCurrent indicates an expected call of decInflightCurrent
func (mr *MockSessionStatsManagerMockRecorder) decInflightCurrent(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decInflightCurrent", reflect.TypeOf((*MockSessionStatsManager)(nil).decInflightCurrent), delta)
}

// addAwaitCurrent mocks base method
func (m *MockSessionStatsManager) addAwaitCurrent(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addAwaitCurrent", delta)
}

// addAwaitCurrent indicates an expected call of addAwaitCurrent
func (mr *MockSessionStatsManagerMockRecorder) addAwaitCurrent(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addAwaitCurrent", reflect.TypeOf((*MockSessionStatsManager)(nil).addAwaitCurrent), delta)
}

// decAwaitCurrent mocks base method
func (m *MockSessionStatsManager) decAwaitCurrent(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "decAwaitCurrent", delta)
}

// decAwaitCurrent indicates an expected call of decAwaitCurrent
func (mr *MockSessionStatsManagerMockRecorder) decAwaitCurrent(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decAwaitCurrent", reflect.TypeOf((*MockSessionStatsManager)(nil).decAwaitCurrent), delta)
}

// GetStats mocks base method
func (m *MockSessionStatsManager) GetStats() *SessionStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(*SessionStats)
	return ret0
}

// GetStats indicates an expected call of GetStats
func (mr *MockSessionStatsManagerMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockSessionStatsManager)(nil).GetStats))
}

// MockpacketStatsManager is a mock of packetStatsManager interface
type MockpacketStatsManager struct {
	ctrl     *gomock.Controller
	recorder *MockpacketStatsManagerMockRecorder
}

// MockpacketStatsManagerMockRecorder is the mock recorder for MockpacketStatsManager
type MockpacketStatsManagerMockRecorder struct {
	mock *MockpacketStatsManager
}

// NewMockpacketStatsManager creates a new mock instance
func NewMockpacketStatsManager(ctrl *gomock.Controller) *MockpacketStatsManager {
	mock := &MockpacketStatsManager{ctrl: ctrl}
	mock.recorder = &MockpacketStatsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockpacketStatsManager) EXPECT() *MockpacketStatsManagerMockRecorder {
	return m.recorder
}

// packetReceived mocks base method
func (m *MockpacketStatsManager) packetReceived(packet packets.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "packetReceived", packet)
}

// packetReceived indicates an expected call of packetReceived
func (mr *MockpacketStatsManagerMockRecorder) packetReceived(packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "packetReceived", reflect.TypeOf((*MockpacketStatsManager)(nil).packetReceived), packet)
}

// packetSent mocks base method
func (m *MockpacketStatsManager) packetSent(packet packets.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "packetSent", packet)
}

// packetSent indicates an expected call of packetSent
func (mr *MockpacketStatsManagerMockRecorder) packetSent(packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "packetSent", reflect.TypeOf((*MockpacketStatsManager)(nil).packetSent), packet)
}

// MockclientStatsManager is a mock of clientStatsManager interface
type MockclientStatsManager struct {
	ctrl     *gomock.Controller
	recorder *MockclientStatsManagerMockRecorder
}

// MockclientStatsManagerMockRecorder is the mock recorder for MockclientStatsManager
type MockclientStatsManagerMockRecorder struct {
	mock *MockclientStatsManager
}

// NewMockclientStatsManager creates a new mock instance
func NewMockclientStatsManager(ctrl *gomock.Controller) *MockclientStatsManager {
	mock := &MockclientStatsManager{ctrl: ctrl}
	mock.recorder = &MockclientStatsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockclientStatsManager) EXPECT() *MockclientStatsManagerMockRecorder {
	return m.recorder
}

// addClientConnected mocks base method
func (m *MockclientStatsManager) addClientConnected() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addClientConnected")
}

// addClientConnected indicates an expected call of addClientConnected
func (mr *MockclientStatsManagerMockRecorder) addClientConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addClientConnected", reflect.TypeOf((*MockclientStatsManager)(nil).addClientConnected))
}

// addClientDisconnected mocks base method
func (m *MockclientStatsManager) addClientDisconnected() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addClientDisconnected")
}

// addClientDisconnected indicates an expected call of addClientDisconnected
func (mr *MockclientStatsManagerMockRecorder) addClientDisconnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addClientDisconnected", reflect.TypeOf((*MockclientStatsManager)(nil).addClientDisconnected))
}

// addSessionActive mocks base method
func (m *MockclientStatsManager) addSessionActive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addSessionActive")
}

// addSessionActive indicates an expected call of addSessionActive
func (mr *MockclientStatsManagerMockRecorder) addSessionActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addSessionActive", reflect.TypeOf((*MockclientStatsManager)(nil).addSessionActive))
}

// decSessionActive mocks base method
func (m *MockclientStatsManager) decSessionActive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "decSessionActive")
}

// decSessionActive indicates an expected call of decSessionActive
func (mr *MockclientStatsManagerMockRecorder) decSessionActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decSessionActive", reflect.TypeOf((*MockclientStatsManager)(nil).decSessionActive))
}

// addSessionInactive mocks base method
func (m *MockclientStatsManager) addSessionInactive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addSessionInactive")
}

// addSessionInactive indicates an expected call of addSessionInactive
func (mr *MockclientStatsManagerMockRecorder) addSessionInactive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addSessionInactive", reflect.TypeOf((*MockclientStatsManager)(nil).addSessionInactive))
}

// decSessionInactive mocks base method
func (m *MockclientStatsManager) decSessionInactive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "decSessionInactive")
}

// decSessionInactive indicates an expected call of decSessionInactive
func (mr *MockclientStatsManagerMockRecorder) decSessionInactive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "decSessionInactive", reflect.TypeOf((*MockclientStatsManager)(nil).decSessionInactive))
}

// addSessionExpired mocks base method
func (m *MockclientStatsManager) addSessionExpired() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addSessionExpired")
}

// addSessionExpired indicates an expected call of addSessionExpired
func (mr *MockclientStatsManagerMockRecorder) addSessionExpired() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addSessionExpired", reflect.TypeOf((*MockclientStatsManager)(nil).addSessionExpired))
}

// MockmessageStatsManager is a mock of messageStatsManager interface
type MockmessageStatsManager struct {
	ctrl     *gomock.Controller
	recorder *MockmessageStatsManagerMockRecorder
}

// MockmessageStatsManagerMockRecorder is the mock recorder for MockmessageStatsManager
type MockmessageStatsManagerMockRecorder struct {
	mock *MockmessageStatsManager
}

// NewMockmessageStatsManager creates a new mock instance
func NewMockmessageStatsManager(ctrl *gomock.Controller) *MockmessageStatsManager {
	mock := &MockmessageStatsManager{ctrl: ctrl}
	mock.recorder = &MockmessageStatsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmessageStatsManager) EXPECT() *MockmessageStatsManagerMockRecorder {
	return m.recorder
}

// messageDropped mocks base method
func (m *MockmessageStatsManager) messageDropped(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageDropped", qos)
}

// messageDropped indicates an expected call of messageDropped
func (mr *MockmessageStatsManagerMockRecorder) messageDropped(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageDropped", reflect.TypeOf((*MockmessageStatsManager)(nil).messageDropped), qos)
}

// messageReceived mocks base method
func (m *MockmessageStatsManager) messageReceived(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageReceived", qos)
}

// messageReceived indicates an expected call of messageReceived
func (mr *MockmessageStatsManagerMockRecorder) messageReceived(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageReceived", reflect.TypeOf((*MockmessageStatsManager)(nil).messageReceived), qos)
}

// messageSent mocks base method
func (m *MockmessageStatsManager) messageSent(qos uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageSent", qos)
}

// messageSent indicates an expected call of messageSent
func (mr *MockmessageStatsManagerMockRecorder) messageSent(qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageSent", reflect.TypeOf((*MockmessageStatsManager)(nil).messageSent), qos)
}

// messageEnqueue mocks base method
func (m *MockmessageStatsManager) messageEnqueue(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageEnqueue", delta)
}

// messageEnqueue indicates an expected call of messageEnqueue
func (mr *MockmessageStatsManagerMockRecorder) messageEnqueue(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageEnqueue", reflect.TypeOf((*MockmessageStatsManager)(nil).messageEnqueue), delta)
}

// messageDequeue mocks base method
func (m *MockmessageStatsManager) messageDequeue(delta uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "messageDequeue", delta)
}

// messageDequeue indicates an expected call of messageDequeue
func (mr *MockmessageStatsManagerMockRecorder) messageDequeue(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "messageDequeue", reflect.TypeOf((*MockmessageStatsManager)(nil).messageDequeue), delta)
}