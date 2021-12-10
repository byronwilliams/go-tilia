package tilia

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/byronwilliams/go-tilia/libraries"
	"github.com/byronwilliams/go-tilia/projects"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

func NewUnexpectedResponseError(expectedStatusCode, actualStatusCode int) error {
	return fmt.Errorf("expected %d, got %d", expectedStatusCode, actualStatusCode)
}

type TiliaClient struct {
	cl      *http.Client
	baseURL string
}

func NewTiliaClient(baseURL string) *TiliaClient {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.RetryWaitMax = time.Minute * 3
	retryClient.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, n int) {
		if n > 0 {
			logger.Printf("retrying request %s, attempt %d", req.URL.String(), n)
		}
	}

	stdClient := retryClient.StandardClient()

	return &TiliaClient{cl: stdClient, baseURL: baseURL}
}

func (tc *TiliaClient) get(ctx context.Context, urlPath string, response interface{}, expectedStatusCodes int) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tc.baseURL+urlPath, nil)

	if err != nil {
		return err
	}

	resp, err := tc.cl.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// fmt.Println(string(b))

	// if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
	// 	return err
	// }

	if err = json.Unmarshal(b, &response); err != nil {
		return err
	}

	if resp.StatusCode != expectedStatusCodes {
		return NewUnexpectedResponseError(expectedStatusCodes, resp.StatusCode)
	}

	return nil
}

func (tc *TiliaClient) post(ctx context.Context, urlPath string, body interface{}, expectedStatusCodes int) (projects.StandardResponse, error) {
	var stdResp projects.StandardResponse
	var r io.Reader

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			return stdResp, err
		}

		fmt.Println(string(b))

		r = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tc.baseURL+urlPath, r)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return stdResp, err
	}

	resp, err := tc.cl.Do(req)

	if err != nil {
		return stdResp, err
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return stdResp, err
	}

	fmt.Println("post resp", resp.Header.Get("content-type"), string(b))

	if strings.HasPrefix(resp.Header.Get("content-type"), "application/json") {
		defer resp.Body.Close()

		if err = json.Unmarshal(b, &stdResp); err != nil {
			return stdResp, err
		}
	}

	if resp.StatusCode != expectedStatusCodes {
		return stdResp, NewUnexpectedResponseError(expectedStatusCodes, resp.StatusCode)
	}

	return stdResp, nil
}

func (tc *TiliaClient) put(ctx context.Context, urlPath string, body interface{}, expectedStatusCodes int) (projects.StandardResponse, error) {
	var stdResp projects.StandardResponse
	var r io.Reader

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			return stdResp, err
		}

		fmt.Println(string(b))

		r = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, tc.baseURL+urlPath, r)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return stdResp, err
	}

	resp, err := tc.cl.Do(req)

	if err != nil {
		return stdResp, err
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return stdResp, err
	}

	fmt.Println("post resp", resp.Header.Get("content-type"), string(b))

	if strings.HasPrefix(resp.Header.Get("content-type"), "application/json") {
		defer resp.Body.Close()

		if err = json.Unmarshal(b, &stdResp); err != nil {
			return stdResp, err
		}
	}

	if resp.StatusCode != expectedStatusCodes {
		return stdResp, NewUnexpectedResponseError(expectedStatusCodes, resp.StatusCode)
	}

	return stdResp, nil
}

func (tc *TiliaClient) delete(ctx context.Context, urlPath string, expectedStatusCodes int) (projects.StandardResponse, error) {
	var stdResp projects.StandardResponse

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, tc.baseURL+urlPath, nil)

	if err != nil {
		return stdResp, err
	}

	resp, err := tc.cl.Do(req)

	if err != nil {
		return stdResp, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&stdResp); err != nil {
		return stdResp, err
	}

	if resp.StatusCode != expectedStatusCodes {
		return stdResp, NewUnexpectedResponseError(expectedStatusCodes, resp.StatusCode)
	}

	return stdResp, nil
}

func (tc *TiliaClient) marshalQueryString(opt interface{}) url.Values {
	v, _ := query.Values(opt)

	return v
}

func (tc *TiliaClient) NewProject(ctx context.Context, id string, opts *projects.CreateProjectRequest) (projects.StandardResponse, error) {
	var body *projects.CreateProjectRequest

	if opts == nil {
		body = &projects.CreateProjectRequest{ID: id}
	} else if opts != nil && id != opts.ID {
		return projects.StandardResponse{}, errors.New("id must match opts")
	} else {
		body = opts
	}

	resp, err := tc.post(ctx, "/jobs", body, http.StatusCreated)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (tc *TiliaClient) DeleteProject(ctx context.Context, id string) (projects.StandardResponse, error) {
	resp, err := tc.delete(ctx, fmt.Sprintf("/jobs/%s", id), http.StatusOK)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (tc *TiliaClient) UploadFileFromURL(ctx context.Context, projectId, filename, downloadFromUrl string) (string, error) {
	cl := &http.Client{Timeout: time.Second * 10}
	resp, err := cl.Get(downloadFromUrl)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected 200, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body := &bytes.Buffer{}
	nw := multipart.NewWriter(body)
	part, err := nw.CreateFormFile("file", filename)

	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, resp.Body)

	if err != nil {
		return "", err
	}

	if err = nw.Close(); err != nil {
		return "", err
	}

	urlPath := fmt.Sprintf("/jobs/%s/files/upload", projectId)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tc.baseURL+urlPath, body)
	req.Header.Set("content-type", nw.FormDataContentType())

	if err != nil {
		return "", err
	}

	resp, err = tc.cl.Do(req)

	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	fmt.Println("post resp", string(b))

	defer resp.Body.Close()

	var stdResp projects.StandardResponse

	if err = json.Unmarshal(b, &stdResp); err != nil {
		return "", err
	}

	expectedStatusCodes := http.StatusOK
	if resp.StatusCode != expectedStatusCodes {
		return "", NewUnexpectedResponseError(expectedStatusCodes, resp.StatusCode)
	}

	fmt.Println(stdResp)

	return stdResp.Resources[0], err
}

func (tc *TiliaClient) AddProductToProject(ctx context.Context, projectId string, body projects.AddProductToProjectRequest) (projects.StandardResponse, error) {
	resp, err := tc.post(ctx, fmt.Sprintf("/jobs/%s/products", projectId), body, http.StatusOK)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (tc *TiliaClient) ExportProject(ctx context.Context, projectId string, format projects.ExportType, opts *projects.ExportRequest) (projects.StandardResponse, error) {
	resp, err := tc.post(ctx, fmt.Sprintf("/jobs/%s/export/%s", projectId, format), opts, http.StatusOK)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (tc *TiliaClient) ExportProjectToBytes(ctx context.Context, projectId string, format projects.ExportType, opts *projects.ExportRequest) ([]byte, error) {
	resp, err := tc.ExportProject(ctx, projectId, format, opts)

	if err != nil {
		return nil, err
	}

	if len(resp.Resources) == 0 {
		return nil, errors.New("not enough resources exported")
	}

	fmt.Println("exportResources", resp.Resources)

	fileResp, err := tc.cl.Get(resp.Resources[0])

	if err != nil {
		return nil, err
	}

	if fileResp.StatusCode != http.StatusOK {
		return nil, errors.New("status code was not 200")
	}

	b, err := ioutil.ReadAll(fileResp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (tc *TiliaClient) StartPlanProject(ctx context.Context, projectId string, body projects.PlanProjectRequest) (projects.StandardResponse, error) {
	resp, err := tc.post(ctx, fmt.Sprintf("/jobs/%s/plan/start", projectId), body, http.StatusCreated)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (tc *TiliaClient) StopPlanProject(ctx context.Context, projectId string) (projects.StandardResponse, error) {
	resp, err := tc.post(ctx, fmt.Sprintf("/jobs/%s/plan/stop", projectId), nil, http.StatusOK)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (tc *TiliaClient) GetPlanStatus(ctx context.Context, projectId string) (projects.PlanStatusResponse, error) {
	var data projects.PlanStatusResponse

	err := tc.get(ctx, fmt.Sprintf("/jobs/%s/plan/status", projectId), &data, http.StatusOK)

	if err != nil {
		return data, err
	}

	return data, nil
}

func (tc *TiliaClient) ListPlanResults(ctx context.Context, projectId string, opts *projects.ListPlanResultsRequestOpts) ([]projects.ListPlanResultsResponse, error) {
	var data []projects.ListPlanResultsResponse
	var qs url.Values

	if opts != nil {
		qs = tc.marshalQueryString(opts)
	}

	fmt.Println("encoded", qs.Encode())

	err := tc.get(ctx, fmt.Sprintf("/jobs/%s/plan/results?%s", projectId, qs.Encode()), &data, http.StatusOK)

	if err != nil {
		return data, err
	}

	return data, nil
}

func (tc *TiliaClient) ApplyPlanResult(ctx context.Context, projectId string, resultId int) (projects.StandardResponse, error) {
	data, err := tc.post(ctx, fmt.Sprintf("/jobs/%s/plan/results/%d/apply", projectId, resultId), nil, http.StatusOK)

	return data, err
}

func (tc *TiliaClient) ListStocks(ctx context.Context) ([]libraries.StockV2, error) {
	var stocks []libraries.StockV2

	err := tc.get(ctx, ("/libraries/v2/stocks"), &stocks, http.StatusOK)

	return stocks, err
}

func (tc *TiliaClient) CreateStock(ctx context.Context, stock libraries.Stock) (projects.StandardResponse, error) {
	resp, err := tc.post(ctx, ("/libraries/stocks"), stock, http.StatusOK)

	return resp, err
}

func (tc *TiliaClient) DeleteStock(ctx context.Context, id string) (projects.StandardResponse, error) {
	return tc.delete(ctx, fmt.Sprintf("/libraries/stocks/%s", id), http.StatusOK)
}

func (tc *TiliaClient) ListStockTypes(ctx context.Context) ([]libraries.StockType, error) {
	var stockTypes []libraries.StockType

	err := tc.get(ctx, ("/libraries/stocktypes"), &stockTypes, http.StatusOK)

	return stockTypes, err
}

func (tc *TiliaClient) CreateGrade(ctx context.Context, stockId string, grade libraries.Grades) (projects.StandardResponse, error) {
	return tc.post(ctx, fmt.Sprintf("/libraries/stocks/%s/grades", stockId), grade, http.StatusOK)
}

func (tc *TiliaClient) CreateRoll(ctx context.Context, stockId, gradeId string, roll libraries.Rolls) (projects.StandardResponse, error) {
	return tc.post(ctx, fmt.Sprintf("/libraries/stocks/%s/grades/%s/rolls", stockId, gradeId), roll, http.StatusOK)
}

func (tc *TiliaClient) ListMarks(ctx context.Context) ([]libraries.Mark, error) {
	var marks []libraries.Mark

	err := tc.get(ctx, ("/libraries/marks"), &marks, http.StatusOK)

	return marks, err
}

func (tc *TiliaClient) GetThingByName(ctx context.Context, name string) (libraries.Thing, error) {
	var things []libraries.Thing
	err := tc.get(ctx, ("/libraries/things"), &things, http.StatusOK)

	if err != nil {
		return libraries.Thing{}, err
	}

	for _, thing := range things {
		if thing.Name == name {
			return thing, nil
		}
	}

	return libraries.Thing{}, sql.ErrNoRows
}

func (tc *TiliaClient) UpdateThing(ctx context.Context, id string, thing libraries.UpdateThing) (projects.StandardResponse, error) {
	return tc.put(ctx, fmt.Sprintf("/libraries/things/%s", id), thing, http.StatusOK)
}
