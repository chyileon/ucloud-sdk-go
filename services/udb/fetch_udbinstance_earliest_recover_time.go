//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB FetchUDBInstanceEarliestRecoverTime

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// FetchUDBInstanceEarliestRecoverTimeRequest is request schema for FetchUDBInstanceEarliestRecoverTime action
type FetchUDBInstanceEarliestRecoverTimeRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// DB实例Id
	DBId *string `required:"true"`
}

// FetchUDBInstanceEarliestRecoverTimeResponse is response schema for FetchUDBInstanceEarliestRecoverTime action
type FetchUDBInstanceEarliestRecoverTimeResponse struct {
	response.CommonBase

	// 获取最早可回档时间点
	EarliestTime int
}

// NewFetchUDBInstanceEarliestRecoverTimeRequest will create request of FetchUDBInstanceEarliestRecoverTime action.
func (c *UDBClient) NewFetchUDBInstanceEarliestRecoverTimeRequest() *FetchUDBInstanceEarliestRecoverTimeRequest {
	req := &FetchUDBInstanceEarliestRecoverTimeRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// FetchUDBInstanceEarliestRecoverTime - 获取UDB最早可回档的时间点
func (c *UDBClient) FetchUDBInstanceEarliestRecoverTime(req *FetchUDBInstanceEarliestRecoverTimeRequest) (*FetchUDBInstanceEarliestRecoverTimeResponse, error) {
	var err error
	var res FetchUDBInstanceEarliestRecoverTimeResponse

	err = c.client.InvokeAction("FetchUDBInstanceEarliestRecoverTime", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}