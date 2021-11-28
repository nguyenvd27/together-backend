package transfers

import (
	"mime/multipart"
	"strconv"
	"time"
	"together-backend/internal/usecases"
)

func ParseRequestBodyFromMultipartFrom(multipartFrom *multipart.Form) (*usecases.ReqBodyEvent, error) {
	var (
		reqBody usecases.ReqBodyEvent
		err     error
	)
	reqBody.Title = multipartFrom.Value["title"][0]
	reqBody.Content = multipartFrom.Value["content"][0]
	reqBody.CreatedBy, err = strconv.ParseUint(multipartFrom.Value["created_by"][0], 10, 64)
	if err != nil {
		return nil, err
	}
	reqBody.StartTime, err = time.Parse("2006-01-02T15:04:05Z0700", multipartFrom.Value["start_time"][0])
	if err != nil {
		return nil, err
	}
	reqBody.EndTime, err = time.Parse("2006-01-02T15:04:05Z0700", multipartFrom.Value["end_time"][0])
	if err != nil {
		return nil, err
	}
	reqBody.Location, err = strconv.Atoi(multipartFrom.Value["location"][0])
	if err != nil {
		return nil, err
	}
	reqBody.DetailLocation = multipartFrom.Value["detail_location"][0]
	return &reqBody, nil
}

func ParseRequestUpdateEvent(multipartFrom *multipart.Form) (*usecases.ReqBodyEditEvent, error) {
	var (
		reqBody usecases.ReqBodyEditEvent
		err     error
	)
	reqBody.Title = multipartFrom.Value["title"][0]
	reqBody.Content = multipartFrom.Value["content"][0]
	reqBody.Id, err = strconv.ParseUint(multipartFrom.Value["id"][0], 10, 64)
	if err != nil {
		return nil, err
	}
	reqBody.CreatedBy, err = strconv.ParseUint(multipartFrom.Value["created_by"][0], 10, 64)
	if err != nil {
		return nil, err
	}
	reqBody.StartTime, err = time.Parse("2006-01-02T15:04:05Z0700", multipartFrom.Value["start_time"][0])
	if err != nil {
		return nil, err
	}
	reqBody.EndTime, err = time.Parse("2006-01-02T15:04:05Z0700", multipartFrom.Value["end_time"][0])
	if err != nil {
		return nil, err
	}
	reqBody.Location, err = strconv.Atoi(multipartFrom.Value["location"][0])
	if err != nil {
		return nil, err
	}
	reqBody.DetailLocation = multipartFrom.Value["detail_location"][0]
	return &reqBody, nil
}

func ParseRequestUpdateProfile(multipartFrom *multipart.Form) (*usecases.ReqBodyUpdateProfile, error) {
	var (
		reqBody usecases.ReqBodyUpdateProfile
		err     error
	)
	reqBody.Name = multipartFrom.Value["name"][0]

	reqBody.Address, err = strconv.Atoi(multipartFrom.Value["address"][0])
	if err != nil {
		return nil, err
	}

	return &reqBody, nil
}
