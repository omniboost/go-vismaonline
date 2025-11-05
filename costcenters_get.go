package vismaonline

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-vismaonline/utils"
)

func (c *Client) NewCostCentersGet() CostCenterGet {
	r := CostCenterGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CostCenterGet struct {
	client      *Client
	queryParams *CostCenterGetQueryParams
	pathParams  *CostCenterGetPathParams
	method      string
	headers     http.Header
	requestBody CostCenterGetBody
}

func (r CostCenterGet) NewQueryParams() *CostCenterGetQueryParams {
	return &CostCenterGetQueryParams{}
}

type CostCenterGetQueryParams struct {
}

func (p CostCenterGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CostCenterGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r CostCenterGet) NewPathParams() *CostCenterGetPathParams {
	return &CostCenterGetPathParams{}
}

type CostCenterGetPathParams struct {
}

func (p *CostCenterGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostCenterGet) PathParams() *CostCenterGetPathParams {
	return r.pathParams
}

func (r *CostCenterGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CostCenterGet) SetMethod(method string) {
	r.method = method
}

func (r *CostCenterGet) Method() string {
	return r.method
}

func (r CostCenterGet) NewRequestBody() CostCenterGetBody {
	return CostCenterGetBody{}
}

type CostCenterGetBody struct {
}

func (r *CostCenterGet) RequestBody() *CostCenterGetBody {
	return nil
}

func (r *CostCenterGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CostCenterGet) SetRequestBody(body CostCenterGetBody) {
	r.requestBody = body
}

func (r *CostCenterGet) NewResponseBody() *CostCenterGetResponseBody {
	return &CostCenterGetResponseBody{}
}

type CostCenterGetResponseBody struct {
	Meta Meta        `json:"Meta"`
	Data CostCenters `json:"Data"`
}

func (r *CostCenterGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/costcenters", r.PathParams())
	return &u
}

func (r *CostCenterGet) Do() (CostCenterGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
