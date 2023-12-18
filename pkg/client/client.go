package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vitorfhc/htb-client/pkg/consts"
	"github.com/vitorfhc/htb-client/pkg/models"
	"github.com/vitorfhc/htb-client/pkg/request"
)

type HTBClient struct {
	opts *HTBClientOptions
}

type HTBClientOptions struct {
	AuthToken  string
	HttpClient *http.Client
	Ctx        context.Context
}

type HTBClientOption func(*HTBClientOptions)

func WithAuthToken(token string) HTBClientOption {
	return func(opts *HTBClientOptions) {
		opts.AuthToken = token
	}
}

func WithHttpClient(client *http.Client) HTBClientOption {
	return func(opts *HTBClientOptions) {
		opts.HttpClient = client
	}
}

func WithCtx(ctx context.Context) HTBClientOption {
	return func(opts *HTBClientOptions) {
		opts.Ctx = ctx
	}
}

func NewHtbClient(opts ...HTBClientOption) *HTBClient {
	options := &HTBClientOptions{
		HttpClient: http.DefaultClient,
		Ctx:        context.Background(),
	}

	for _, opt := range opts {
		opt(options)
	}

	return &HTBClient{
		opts: options,
	}
}

// FindActiveMachineByName returns the active machines with the given name.
// If there is no active machine with the given name, it returns ErrMachineNotFound.
func (c *HTBClient) FindActiveMachineByName(name string) (*models.Machine, error) {
	httpRequest, err := request.New(
		c.opts.Ctx,
		request.WithAuthToken(c.opts.AuthToken),
		request.WithPath(consts.HTBPathListActiveLabMachines),
		request.WithQuery(url.Values{
			consts.HTBQueryKeyPerPage: {"100"},
			consts.HTBQueryKeyKeyword: {name},
		}),
	).Build()
	if err != nil {
		return nil, err
	}

	response, err := c.do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	machinesListResponse := APIDataResponse[models.MachinesList]{}
	if err := json.NewDecoder(response.Body).Decode(&machinesListResponse); err != nil {
		return nil, err
	}

	if len(machinesListResponse.Data) == 0 {
		return nil, &ErrMachineNotFound{Name: name}
	}

	if len(machinesListResponse.Data) > 1 {
		return nil, &ErrMultipleMachinesFound{Name: name}
	}

	return machinesListResponse.Data[0], nil
}

// GetAssignedVPNServer returns the assigned vpn server for the given product.
// If there is no assigned vpn server, it returns ErrNoAssignedVPNServer.
func (c *HTBClient) GetAssignedVPNServer(product consts.HTBProduct) (*models.VPNServer, error) {
	httpRequest, err := request.New(
		c.opts.Ctx,
		request.WithAuthToken(c.opts.AuthToken),
		request.WithPath(consts.HTBPathVPNServers),
		request.WithQuery(url.Values{
			consts.HTBQueryKeyProduct: {string(product)},
		}),
	).Build()
	if err != nil {
		return nil, err
	}

	response, err := c.do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	apiResponse := APIDataResponse[*models.VPNServersData]{}
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	if apiResponse.Data.Assigned == nil {
		return nil, &ErrNoAssignedVPNServer{}
	}

	return apiResponse.Data.Assigned, nil
}

// GetActiveLabMachine returns the active lab machine.
// If there is no active lab machine, it returns ErrNoActiveLabMachine.
func (c *HTBClient) GetActiveLabMachine() (*models.Machine, error) {
	httpRequest, err := request.New(
		c.opts.Ctx,
		request.WithAuthToken(c.opts.AuthToken),
		request.WithPath(consts.HTBPathActiveLabMachine),
	).Build()
	if err != nil {
		return nil, err
	}

	response, err := c.do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	apiResponse := APIInfoResponse[*models.Machine]{}

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	if apiResponse.Info == nil {
		return nil, &ErrNoActiveLabMachine{}
	}

	return apiResponse.Info, nil
}

// SpawnLabMachine spawns a lab machine with the given id.
func (c *HTBClient) SpawnLabMachine(id uint) error {
	machine := struct {
		MachineID uint `json:"machine_id"`
	}{
		MachineID: id,
	}

	body, err := json.Marshal(machine)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)

	httpRequest, err := request.New(
		c.opts.Ctx,
		request.WithAuthToken(c.opts.AuthToken),
		request.WithMethod(http.MethodPost),
		request.WithPath(consts.HTBPathSpawnLabMachine),
		request.WithJSONBody(bodyReader),
	).Build()
	if err != nil {
		return err
	}

	_, err = c.do(httpRequest)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer httpRequest.Body.Close()

	return nil
}

// TerminateLabMachine terminates a lab machine with the given id.
func (c *HTBClient) TerminateLabMachine(id uint) error {
	machine := struct {
		MachineID uint `json:"machine_id"`
	}{
		MachineID: id,
	}

	body, err := json.Marshal(machine)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)

	httpRequest, err := request.New(
		c.opts.Ctx,
		request.WithAuthToken(c.opts.AuthToken),
		request.WithMethod(http.MethodPost),
		request.WithPath(consts.HTBPathTerminateLabMachine),
		request.WithJSONBody(bodyReader),
	).Build()
	if err != nil {
		return err
	}

	_, err = c.do(httpRequest)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer httpRequest.Body.Close()

	return nil
}

func (c *HTBClient) GetVPNServers(product consts.HTBProduct) (models.VPNServersList, error) {
	httpRequest, err := request.New(
		c.opts.Ctx,
		request.WithAuthToken(c.opts.AuthToken),
		request.WithPath(consts.HTBPathVPNServers),
		request.WithQuery(url.Values{
			consts.HTBQueryKeyProduct: {string(product)},
		}),
	).Build()
	if err != nil {
		return nil, err
	}

	response, err := c.do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	apiResponse := APIDataResponse[*models.VPNServersData]{}
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	var servers models.VPNServersList
	for _, server := range apiResponse.Data.Options {
		for _, option := range server {
			for _, deepServer := range option.Servers {
				servers = append(servers, deepServer)
			}
		}
	}

	return servers, nil
}

func (c *HTBClient) do(request *http.Request) (*http.Response, error) {
	response, err := c.opts.HttpClient.Do(request.WithContext(c.opts.Ctx))
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 299 {
		jsonDecoder := json.NewDecoder(response.Body)
		apiMsg := APIMessageResponse{}
		if err := jsonDecoder.Decode(&apiMsg); err == nil {
			return nil, &ErrAPIError{
				StatusCode: response.StatusCode,
				Message:    apiMsg.Message,
			}
		}

		return nil, &ErrUnexpectedStatusCode{StatusCode: response.StatusCode}
	}

	if response.Header.Get("Content-Type") != "application/json" {
		return nil, &ErrUnexpectedContentType{
			ExpectedContentType: "application/json",
			ContentType:         response.Header.Get("Content-Type"),
		}
	}

	return response, nil
}
