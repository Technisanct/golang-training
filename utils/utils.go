package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"golang-training/libs/errors"
)

// RestResponse ...
// Generic response structure for all ReST API calls
type RestResponse struct {
	Message      string `json:"message"`
	Data         any    `json:"data"`
	Next         string `json:"next,omitempty"`
	TotalResults int64  `json:"totalResults,omitempty"`
}

// BuildRestResp ...
// Generic response building for unit test returning RestResponse
func BuildRestResp(req interface{}) interface{} {
	if req == nil {
		return nil
	}

	payload, err := json.Marshal(req)
	if err != nil {
		log.Error().Err(err).Msg(errors.ErrMarshalStruct.Error())
		return err
	}

	var restResponse interface{}
	err = json.Unmarshal(payload, &restResponse)
	if err != nil {
		log.Error().Err(err).Msg(errors.ErrUnMarshalResponseBody.Error())
		return err
	}

	return restResponse
}
