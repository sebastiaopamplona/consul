// Code generated by mockery v2.12.2. DO NOT EDIT.

package pbdns

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockDNSServiceClient is an autogenerated mock type for the DNSServiceClient type
type MockDNSServiceClient struct {
	mock.Mock
}

// Query provides a mock function with given fields: ctx, in, opts
func (_m *MockDNSServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryRequest, ...grpc.CallOption) *QueryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockDNSServiceClient creates a new instance of MockDNSServiceClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDNSServiceClient(t testing.TB) *MockDNSServiceClient {
	mock := &MockDNSServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}