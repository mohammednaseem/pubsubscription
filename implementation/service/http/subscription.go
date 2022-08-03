package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
)

func (i *registryService) HttpReq(ctx context.Context, method string, entity string, body string, path string) (model.Response, error) {

	client := http.Client{}
	var dr model.Response
	request, err := http.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	if resp.StatusCode != 200 {
		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
		return dr, errors.New("Http Request TO Device Manager Failed")
	}

	dr.StatusCode = resp.StatusCode
	return dr, err
}

// func (i *registryService) HttpPatch(ctx context.Context, entity string, body string) (model.Response, error) {
// 	path := gjson.Get(string(body), "parent").String()
// 	query := gjson.Get(string(body), "updatemask").String()
// 	client := http.Client{}
// 	var dr model.Response
// 	request, err := http.NewRequest("PATCH", i.baseUrl+entity+"/"+path+"?updateMask="+query, bytes.NewBuffer([]byte(body)))
// 	request.Header.Set("Content-Type", "application/json")
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return dr, errors.New("Http Request TO Device Manager Failed")
// 	}

// 	dr.StatusCode = resp.StatusCode
// 	return dr, err
// }

// func (i *registryService) HttpDelete(ctx context.Context, entity string, body string) (model.Response, error) {
// 	path := gjson.Get(string(body), "parent").String()
// 	client := http.Client{}
// 	var dr model.Response
// 	request, err := http.NewRequest("DELETE", i.baseUrl+entity+"/"+path, bytes.NewBuffer([]byte(body)))
// 	request.Header.Set("Content-Type", "application/json")
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return dr, errors.New("Http Request TO Device Manager Failed")
// 	}

// 	dr.StatusCode = resp.StatusCode

// 	return dr, err
// }

// func (i *registryService) CreateDevice(ctx context.Context, registry string) (model.Response, error) {
// 	// c, cancel := context.WithTimeout(ctx, i.contextTimeout)
// 	// defer cancel()

// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return dr, err

// 	// }
// 	resp, err := http.Get("http://postman-echo.com/get")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var result map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&result)

// 	log.Println(result)
// 	var dr model.Response
// 	return dr, err
// }

// func (i *registryService) UpdateDevice(ctx context.Context, registry string) (model.Response, error) {
// 	// c, cancel := context.WithTimeout(ctx, i.contextTimeout)
// 	// defer cancel()

// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return dr, err

// 	// }
// 	resp, err := http.Get("http://postman-echo.com/get")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var result map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&result)

// 	log.Println(result)
// 	var dr model.Response
// 	return dr, err
// }

// func (i *registryService) DeleteDevice(ctx context.Context, device string) (model.Response, error) {
// 	// c, cancel := context.WithTimeout(ctx, i.contextTimeout)
// 	// defer cancel()

// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return dr, err

// 	// }
// 	path := gjson.Get(string(registry), "parent").String()
// 	var dr model.Response
// 	err := httpDelete(i.baseUrl, registry, path)
// 	return dr, err
// }
// func (i *registryService) httpPost(ctx context.Context, entity string, body string) (model.Response, error) {
// 	path := gjson.Get(string(body), "parent").String()
// 	var dr model.Response
// 	resp, err := http.Post(i.baseUrl+entity+"/"+path, "application/json", bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return dr, errors.New("Http Request TO Device Manager Failed")
// 	}
// 	print(resp.StatusCode)

// 	dr.StatusCode = resp.StatusCode

// 	return dr, err
// }
// func (i *registryService) httpPatch(ctx context.Context, entity string, body string) (model.Response, error) {
// 	path := gjson.Get(string(body), "parent").String()
// 	query := gjson.Get(string(body), "updatemask").String()
// 	client := http.Client{}
// 	var dr model.Response
// 	request, err := http.NewRequest("PATCH", i.baseUrl+entity+"/"+path+"?updateMask="+query, bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return dr, errors.New("Http Request TO Device Manager Failed")
// 	}

// 	dr.StatusCode = resp.StatusCode
// 	return dr, err

// }
// func (i *registryService) httpDelete(ctx context.Context, entity string, body string) (model.Response, error) {
// 	path := gjson.Get(string(body), "parent").String()
// 	client := http.Client{}
// 	var dr model.Response
// 	request, err := http.NewRequest("DELETE", i.baseUrl+entity+"/"+path, bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return dr, errors.New("Http Request TO Device Manager Failed")
// 	}

// 	dr.StatusCode = resp.StatusCode

// 	return dr, err
// }
