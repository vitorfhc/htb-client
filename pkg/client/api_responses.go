package client

type APIDataResponse[T any] struct {
	Data T `json:"data"`
}

type APIInfoResponse[T any] struct {
	Info T `json:"info"`
}

type APIMessageResponse struct {
	Message string `json:"message"`
}
