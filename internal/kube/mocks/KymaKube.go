// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	corev1 "k8s.io/api/core/v1"
	api "k8s.io/client-go/tools/clientcmd/api"

	dynamic "k8s.io/client-go/dynamic"

	kubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	v1 "k8s.io/api/apps/v1"

	versioned "istio.io/client-go/pkg/clientset/versioned"
)

// KymaKube is an autogenerated mock type for the KymaKube type
type KymaKube struct {
	mock.Mock
}

// Apply provides a mock function with given fields: manifest
func (_m *KymaKube) Apply(manifest []byte) error {
	ret := _m.Called(manifest)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(manifest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DefaultNamespace provides a mock function with given fields:
func (_m *KymaKube) DefaultNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Dynamic provides a mock function with given fields:
func (_m *KymaKube) Dynamic() dynamic.Interface {
	ret := _m.Called()

	var r0 dynamic.Interface
	if rf, ok := ret.Get(0).(func() dynamic.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.Interface)
		}
	}

	return r0
}

// IsPodDeployed provides a mock function with given fields: namespace, name
func (_m *KymaKube) IsPodDeployed(namespace string, name string) (bool, error) {
	ret := _m.Called(namespace, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(namespace, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsPodDeployedByLabel provides a mock function with given fields: namespace, labelName, labelValue
func (_m *KymaKube) IsPodDeployedByLabel(namespace string, labelName string, labelValue string) (bool, error) {
	ret := _m.Called(namespace, labelName, labelValue)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(namespace, labelName, labelValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(namespace, labelName, labelValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Istio provides a mock function with given fields:
func (_m *KymaKube) Istio() versioned.Interface {
	ret := _m.Called()

	var r0 versioned.Interface
	if rf, ok := ret.Get(0).(func() versioned.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(versioned.Interface)
		}
	}

	return r0
}

// KubeConfig provides a mock function with given fields:
func (_m *KymaKube) KubeConfig() *api.Config {
	ret := _m.Called()

	var r0 *api.Config
	if rf, ok := ret.Get(0).(func() *api.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Config)
		}
	}

	return r0
}

// RestConfig provides a mock function with given fields:
func (_m *KymaKube) RestConfig() *rest.Config {
	ret := _m.Called()

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func() *rest.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	return r0
}

// Static provides a mock function with given fields:
func (_m *KymaKube) Static() kubernetes.Interface {
	ret := _m.Called()

	var r0 kubernetes.Interface
	if rf, ok := ret.Get(0).(func() kubernetes.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubernetes.Interface)
		}
	}

	return r0
}

// WaitDeploymentStatus provides a mock function with given fields: namespace, name, cond, status
func (_m *KymaKube) WaitDeploymentStatus(namespace string, name string, cond v1.DeploymentConditionType, status corev1.ConditionStatus) error {
	ret := _m.Called(namespace, name, cond, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, v1.DeploymentConditionType, corev1.ConditionStatus) error); ok {
		r0 = rf(namespace, name, cond, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitPodStatus provides a mock function with given fields: namespace, name, status
func (_m *KymaKube) WaitPodStatus(namespace string, name string, status corev1.PodPhase) error {
	ret := _m.Called(namespace, name, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, corev1.PodPhase) error); ok {
		r0 = rf(namespace, name, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitPodStatusByLabel provides a mock function with given fields: namespace, labelName, labelValue, status
func (_m *KymaKube) WaitPodStatusByLabel(namespace string, labelName string, labelValue string, status corev1.PodPhase) error {
	ret := _m.Called(namespace, labelName, labelValue, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, corev1.PodPhase) error); ok {
		r0 = rf(namespace, labelName, labelValue, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchResource provides a mock function with given fields: res, name, namespace, checkFn
func (_m *KymaKube) WatchResource(res schema.GroupVersionResource, name string, namespace string, checkFn func(*unstructured.Unstructured) (bool, error)) error {
	ret := _m.Called(res, name, namespace, checkFn)

	var r0 error
	if rf, ok := ret.Get(0).(func(schema.GroupVersionResource, string, string, func(*unstructured.Unstructured) (bool, error)) error); ok {
		r0 = rf(res, name, namespace, checkFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewKymaKube interface {
	mock.TestingT
	Cleanup(func())
}

// NewKymaKube creates a new instance of KymaKube. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKymaKube(t mockConstructorTestingTNewKymaKube) *KymaKube {
	mock := &KymaKube{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
