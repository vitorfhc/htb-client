package client

import "fmt"

type ErrUnexpectedStatusCode struct {
	StatusCode int
}

func (e *ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.StatusCode)
}

type ErrMachineNotFound struct {
	Name string
}

func (e *ErrMachineNotFound) Error() string {
	return "machine not found: " + e.Name
}

type ErrUnexpectedContentType struct {
	ExpectedContentType string
	ContentType         string
}

func (e *ErrUnexpectedContentType) Error() string {
	return fmt.Sprintf("unexpected content type: %s, expected: %s", e.ContentType, e.ExpectedContentType)
}

type ErrNoActiveLabMachine struct{}

func (e *ErrNoActiveLabMachine) Error() string {
	return "no active lab machine"
}

type ErrNoAssignedVPNServer struct{}

func (e *ErrNoAssignedVPNServer) Error() string {
	return "no assigned vpn server"
}

type ErrMultipleMachinesFound struct {
	Name string
}

func (e *ErrMultipleMachinesFound) Error() string {
	return fmt.Sprintf("multiple machines found for the name: %s", e.Name)
}

type ErrAPIError struct {
	StatusCode int
	Message    string
}

func (e *ErrAPIError) Error() string {
	return fmt.Sprintf("api error: status code %d: %s", e.StatusCode, e.Message)
}
