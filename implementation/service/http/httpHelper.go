package http

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"github.com/rs/zerolog/log"
// )

// // func httpPost(baseurl string, body string) error {
// // 	resp, err := http.Post(baseurl+path, "application/json", bytes.NewBuffer([]byte(body)))
// // 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return errors.New("Http Request TO Device Manager Failed")
// 	}
// 	print(resp.StatusCode)
// 	return err
// }
// func httpUpdate(baseurl string, body string) error {
// 	client := http.Client{}
// 	request, err := http.NewRequest("PATCH", baseurl+path+"?updateMask="+query, bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return errors.New("Http Request TO Device Manager Failed")
// 	}

// 	return err
// }
// func httpDelete(baseurl string, body string) error {
// 	client := http.Client{}
// 	request, err := http.NewRequest("DELETE", baseurl+path, bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Error().Err(err).Msg("")
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Error().Err(err).Msg(fmt.Sprintf("Response is %d", resp.StatusCode))
// 		return errors.New("Http Request TO Device Manager Failed")
// 	}

// 	return err
// }
