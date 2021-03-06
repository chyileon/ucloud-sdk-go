// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeBucketDomainRequest is request schema for DescribeBucketDomain action
type DescribeBucketDomainRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 待获取Bucket的名称
	BucketName *string `required:"true"`
}

// DescribeBucketDomainResponse is response schema for DescribeBucketDomain action
type DescribeBucketDomainResponse struct {
	response.CommonBase

	// bucket自定义域名信息
	DataSet []BucketDomainSet
}

// NewDescribeBucketDomainRequest will create request of DescribeBucketDomain action.
func (c *UFileClient) NewDescribeBucketDomainRequest() *DescribeBucketDomainRequest {
	req := &DescribeBucketDomainRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeBucketDomain - 获取bcuket自定义域名描述信息
func (c *UFileClient) DescribeBucketDomain(req *DescribeBucketDomainRequest) (*DescribeBucketDomainResponse, error) {
	var err error
	var res DescribeBucketDomainResponse

	reqCopier := *req
	// In order to ignore the parameters about Region
	reqCopier.Region = ucloud.String("")

	err = c.Client.InvokeAction("DescribeBucketDomain", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
