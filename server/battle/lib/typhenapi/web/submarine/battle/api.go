// This file was generated by typhen-api

package battle

import (
	"bytes"
	"github.com/shiwano/submarine/server/battle/lib/typhenapi/core"
	submarine "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	submarine_battle "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"io"
	"io/ioutil"
	"net/http"
)

// WebAPI sends a request.
type WebAPI struct {
	baseURI              string
	serializer           typhenapi.Serializer
	beforeRequestHandler func(*http.Request)
	Client               *http.Client
}

// New creates a WebAPI.
func New(baseURI string, serializer typhenapi.Serializer, client *http.Client) *WebAPI {
	if client == nil {
		client = &http.Client{}
	}
	api := &WebAPI{}
	api.baseURI = baseURI
	api.serializer = serializer
	api.Client = client
	return api
}

// FindRoom send a findRoom request.
func (api *WebAPI) FindRoom(roomId int64) (*submarine_battle.FindRoomObject, error) {
	reqBody := &FindRoomRequestBody{}
	reqBody.RoomId = roomId

	reqBodyData, err := reqBody.Bytes(api.serializer)
	if err != nil {
		return nil, err
	}

	req, err := api.createRequest("POST", api.baseURI+"/battle/find_room", bytes.NewReader(reqBodyData))
	if err != nil {
		return nil, err
	}

	res, data, err := api.sendRequest(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, api.tryToDeserializeAPIError(data)
	}

	result := new(submarine_battle.FindRoomObject)
	if err := api.serializer.Deserialize(data, result); err != nil {
		return nil, err
	}
	if err := result.Coerce(); err != nil {
		return nil, err
	}
	return result, nil
}

// FindRoomMember send a findRoomMember request.
func (api *WebAPI) FindRoomMember(roomKey string) (*submarine_battle.FindRoomMemberObject, error) {
	reqBody := &FindRoomMemberRequestBody{}
	reqBody.RoomKey = roomKey

	reqBodyData, err := reqBody.Bytes(api.serializer)
	if err != nil {
		return nil, err
	}

	req, err := api.createRequest("POST", api.baseURI+"/battle/find_room_member", bytes.NewReader(reqBodyData))
	if err != nil {
		return nil, err
	}

	res, data, err := api.sendRequest(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, api.tryToDeserializeAPIError(data)
	}

	result := new(submarine_battle.FindRoomMemberObject)
	if err := api.serializer.Deserialize(data, result); err != nil {
		return nil, err
	}
	if err := result.Coerce(); err != nil {
		return nil, err
	}
	return result, nil
}

// CloseRoom send a closeRoom request.
func (api *WebAPI) CloseRoom(roomId int64) (*typhenapi.Void, error) {
	reqBody := &CloseRoomRequestBody{}
	reqBody.RoomId = roomId

	reqBodyData, err := reqBody.Bytes(api.serializer)
	if err != nil {
		return nil, err
	}

	req, err := api.createRequest("POST", api.baseURI+"/battle/close_room", bytes.NewReader(reqBodyData))
	if err != nil {
		return nil, err
	}

	res, data, err := api.sendRequest(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, api.tryToDeserializeAPIError(data)
	}

	result := new(typhenapi.Void)
	if err := api.serializer.Deserialize(data, result); err != nil {
		return nil, err
	}
	if err := result.Coerce(); err != nil {
		return nil, err
	}
	return result, nil
}

func (api *WebAPI) SetBeforeRequestHandler(handler func(*http.Request)) {
	api.beforeRequestHandler = handler
}

func (api *WebAPI) tryToDeserializeAPIError(data []byte) error {
	apiError := new(submarine.Error)
	if err := api.serializer.Deserialize(data, apiError); err != nil {
		return err
	}

	if err := apiError.Coerce(); err != nil {
		return err
	}

	return apiError
}

func (api *WebAPI) createRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if api.beforeRequestHandler != nil {
		api.beforeRequestHandler(req)
	}

	return req, nil
}

func (api *WebAPI) sendRequest(req *http.Request) (*http.Response, []byte, error) {
	res, err := api.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, data, nil
}
