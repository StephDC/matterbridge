// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// Files returns request builder for AgreementFile collection
func (b *AgreementRequestBuilder) Files() *AgreementFilesCollectionRequestBuilder {
	bb := &AgreementFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/files"
	return bb
}

// AgreementFilesCollectionRequestBuilder is request builder for AgreementFile collection
type AgreementFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AgreementFile collection
func (b *AgreementFilesCollectionRequestBuilder) Request() *AgreementFilesCollectionRequest {
	return &AgreementFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AgreementFile item
func (b *AgreementFilesCollectionRequestBuilder) ID(id string) *AgreementFileRequestBuilder {
	bb := &AgreementFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AgreementFilesCollectionRequest is request for AgreementFile collection
type AgreementFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AgreementFile collection
func (r *AgreementFilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AgreementFile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AgreementFile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []AgreementFile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for AgreementFile collection, max N pages
func (r *AgreementFilesCollectionRequest) GetN(ctx context.Context, n int) ([]AgreementFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AgreementFile collection
func (r *AgreementFilesCollectionRequest) Get(ctx context.Context) ([]AgreementFile, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AgreementFile collection
func (r *AgreementFilesCollectionRequest) Add(ctx context.Context, reqObj *AgreementFile) (resObj *AgreementFile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}