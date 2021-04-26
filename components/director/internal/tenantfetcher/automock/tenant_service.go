// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// TenantService is an autogenerated mock type for the TenantService type
type TenantService struct {
	mock.Mock
}

// CreateManyIfNotExists provides a mock function with given fields: ctx, tenantInputs
func (_m *TenantService) CreateManyIfNotExists(ctx context.Context, tenantInputs []model.BusinessTenantMappingInput) error {
	ret := _m.Called(ctx, tenantInputs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []model.BusinessTenantMappingInput) error); ok {
		r0 = rf(ctx, tenantInputs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMany provides a mock function with given fields: ctx, tenantInputs
func (_m *TenantService) DeleteMany(ctx context.Context, tenantInputs []model.BusinessTenantMappingInput) error {
	ret := _m.Called(ctx, tenantInputs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []model.BusinessTenantMappingInput) error); ok {
		r0 = rf(ctx, tenantInputs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInternalTenant provides a mock function with given fields: ctx, externalTenant
func (_m *TenantService) GetInternalTenant(ctx context.Context, externalTenant string) (string, error) {
	ret := _m.Called(ctx, externalTenant)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, externalTenant)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, externalTenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *TenantService) List(ctx context.Context) ([]*model.BusinessTenantMapping, error) {
	ret := _m.Called(ctx)

	var r0 []*model.BusinessTenantMapping
	if rf, ok := ret.Get(0).(func(context.Context) []*model.BusinessTenantMapping); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BusinessTenantMapping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}