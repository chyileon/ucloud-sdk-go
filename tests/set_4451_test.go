// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet4451(t *testing.T) {
	t.Skip()
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	testSet4451LoginByPassword00(&ctx)
	testSet4451DescribeUser01(&ctx)
	testSet4451CreateUser02(&ctx)
	testSet4451Recharge03(&ctx)
}

func testSet4451LoginByPassword00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewLoginByPasswordRequest()

	ctx.NoError(utest.SetReqValue(req, "UserEmail", "admin@ucloud.cn"))

	ctx.NoError(utest.SetReqValue(req, "Password", "ucloud.cn"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.LoginByPassword(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "LoginByPasswordResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet4451DescribeUser01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewDescribeUserRequest()

	ctx.NoError(utest.SetReqValue(req, "Offset", 0))

	ctx.NoError(utest.SetReqValue(req, "Limit", 10))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.DescribeUser(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeUserResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet4451CreateUser02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewCreateUserRequest()

	s := fmt.Sprintf("testsdk%s@ucloud.cn", ctx.Must(utest.GetTimestamp("10")))
	ctx.NoError(utest.SetReqValue(req, "UserEmail", s))

	ctx.NoError(utest.SetReqValue(req, "PassWord", "ucloud.cn"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.CreateUser(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateUserResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["UserID"] = ctx.Must(utest.GetValue(resp, "UserID"))
}

func testSet4451Recharge03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ucloudstackClient.NewRechargeRequest()

	ctx.NoError(utest.SetReqValue(req, "UserID", ctx.GetVar("UserID")))

	ctx.NoError(utest.SetReqValue(req, "SerialNo", ctx.Must(utest.GetTimestamp("10"))))

	ctx.NoError(utest.SetReqValue(req, "FromType", "INPOUR_FROM_ALIPAY"))

	ctx.NoError(utest.SetReqValue(req, "Amount", 100))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ucloudstackClient.Recharge(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "RechargeResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}
