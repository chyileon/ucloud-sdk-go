//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem CreateURedisBackup

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateURedisBackupRequest is request schema for CreateURedisBackup action
type CreateURedisBackupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 资源id
	GroupId *string `required:"true"`

	// 请求创建组的名称 (范围[6-63],只能包含英文、数字以及符号-和_)
	BackupName *string `required:"true"`
}

// CreateURedisBackupResponse is response schema for CreateURedisBackup action
type CreateURedisBackupResponse struct {
	response.CommonBase

	// 备份id
	BackupId string
}

// NewCreateURedisBackupRequest will create request of CreateURedisBackup action.
func (c *UMemClient) NewCreateURedisBackupRequest() *CreateURedisBackupRequest {
	req := &CreateURedisBackupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// CreateURedisBackup - 创建主备Redis备份
func (c *UMemClient) CreateURedisBackup(req *CreateURedisBackupRequest) (*CreateURedisBackupResponse, error) {
	var err error
	var res CreateURedisBackupResponse

	err = c.client.InvokeAction("CreateURedisBackup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}