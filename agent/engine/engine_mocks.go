// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/engine (interfaces: TaskEngine,DockerClient,ImageManager)

package engine

import (
	context0 "context"
	io "io"
	time "time"

	api "github.com/aws/amazon-ecs-agent/agent/api"
	dockerclient "github.com/aws/amazon-ecs-agent/agent/engine/dockerclient"
	image "github.com/aws/amazon-ecs-agent/agent/engine/image"
	statechange "github.com/aws/amazon-ecs-agent/agent/statechange"
	statemanager "github.com/aws/amazon-ecs-agent/agent/statemanager"
	go_dockerclient "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of TaskEngine interface
type MockTaskEngine struct {
	ctrl     *gomock.Controller
	recorder *_MockTaskEngineRecorder
}

// Recorder for MockTaskEngine (not exported)
type _MockTaskEngineRecorder struct {
	mock *MockTaskEngine
}

func NewMockTaskEngine(ctrl *gomock.Controller) *MockTaskEngine {
	mock := &MockTaskEngine{ctrl: ctrl}
	mock.recorder = &_MockTaskEngineRecorder{mock}
	return mock
}

func (_m *MockTaskEngine) EXPECT() *_MockTaskEngineRecorder {
	return _m.recorder
}

func (_m *MockTaskEngine) AddTask(_param0 *api.Task) error {
	ret := _m.ctrl.Call(_m, "AddTask", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTaskEngineRecorder) AddTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTask", arg0)
}

func (_m *MockTaskEngine) Disable() {
	_m.ctrl.Call(_m, "Disable")
}

func (_mr *_MockTaskEngineRecorder) Disable() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disable")
}

func (_m *MockTaskEngine) GetTaskByArn(_param0 string) (*api.Task, bool) {
	ret := _m.ctrl.Call(_m, "GetTaskByArn", _param0)
	ret0, _ := ret[0].(*api.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineRecorder) GetTaskByArn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTaskByArn", arg0)
}

func (_m *MockTaskEngine) Init(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Init", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTaskEngineRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockTaskEngine) ListTasks() ([]*api.Task, error) {
	ret := _m.ctrl.Call(_m, "ListTasks")
	ret0, _ := ret[0].([]*api.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskEngineRecorder) ListTasks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks")
}

func (_m *MockTaskEngine) MarshalJSON() ([]byte, error) {
	ret := _m.ctrl.Call(_m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskEngineRecorder) MarshalJSON() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarshalJSON")
}

func (_m *MockTaskEngine) MustInit(_param0 context.Context) {
	_m.ctrl.Call(_m, "MustInit", _param0)
}

func (_mr *_MockTaskEngineRecorder) MustInit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MustInit", arg0)
}

func (_m *MockTaskEngine) SetSaver(_param0 statemanager.Saver) {
	_m.ctrl.Call(_m, "SetSaver", _param0)
}

func (_mr *_MockTaskEngineRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSaver", arg0)
}

func (_m *MockTaskEngine) StateChangeEvents() chan statechange.Event {
	ret := _m.ctrl.Call(_m, "StateChangeEvents")
	ret0, _ := ret[0].(chan statechange.Event)
	return ret0
}

func (_mr *_MockTaskEngineRecorder) StateChangeEvents() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StateChangeEvents")
}

func (_m *MockTaskEngine) UnmarshalJSON(_param0 []byte) error {
	ret := _m.ctrl.Call(_m, "UnmarshalJSON", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTaskEngineRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnmarshalJSON", arg0)
}

func (_m *MockTaskEngine) Version() (string, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskEngineRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

// Mock of DockerClient interface
type MockDockerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockDockerClientRecorder
}

// Recorder for MockDockerClient (not exported)
type _MockDockerClientRecorder struct {
	mock *MockDockerClient
}

func NewMockDockerClient(ctrl *gomock.Controller) *MockDockerClient {
	mock := &MockDockerClient{ctrl: ctrl}
	mock.recorder = &_MockDockerClientRecorder{mock}
	return mock
}

func (_m *MockDockerClient) EXPECT() *_MockDockerClientRecorder {
	return _m.recorder
}

func (_m *MockDockerClient) APIVersion() (dockerclient.DockerVersion, error) {
	ret := _m.ctrl.Call(_m, "APIVersion")
	ret0, _ := ret[0].(dockerclient.DockerVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) APIVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "APIVersion")
}

func (_m *MockDockerClient) ContainerEvents(_param0 context0.Context) (<-chan DockerContainerChangeEvent, error) {
	ret := _m.ctrl.Call(_m, "ContainerEvents", _param0)
	ret0, _ := ret[0].(<-chan DockerContainerChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) ContainerEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerEvents", arg0)
}

func (_m *MockDockerClient) CreateContainer(_param0 *go_dockerclient.Config, _param1 *go_dockerclient.HostConfig, _param2 string, _param3 time.Duration) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "CreateContainer", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

func (_mr *_MockDockerClientRecorder) CreateContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateContainer", arg0, arg1, arg2, arg3)
}

func (_m *MockDockerClient) DescribeContainer(_param0 string) (api.ContainerStatus, DockerContainerMetadata) {
	ret := _m.ctrl.Call(_m, "DescribeContainer", _param0)
	ret0, _ := ret[0].(api.ContainerStatus)
	ret1, _ := ret[1].(DockerContainerMetadata)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) DescribeContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeContainer", arg0)
}

func (_m *MockDockerClient) ImportLocalEmptyVolumeImage() DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "ImportLocalEmptyVolumeImage")
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

func (_mr *_MockDockerClientRecorder) ImportLocalEmptyVolumeImage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImportLocalEmptyVolumeImage")
}

func (_m *MockDockerClient) InspectContainer(_param0 string, _param1 time.Duration) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "InspectContainer", _param0, _param1)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) InspectContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectContainer", arg0, arg1)
}

func (_m *MockDockerClient) InspectImage(_param0 string) (*go_dockerclient.Image, error) {
	ret := _m.ctrl.Call(_m, "InspectImage", _param0)
	ret0, _ := ret[0].(*go_dockerclient.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) InspectImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectImage", arg0)
}

func (_m *MockDockerClient) KnownVersions() []dockerclient.DockerVersion {
	ret := _m.ctrl.Call(_m, "KnownVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

func (_mr *_MockDockerClientRecorder) KnownVersions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KnownVersions")
}

func (_m *MockDockerClient) ListContainers(_param0 bool, _param1 time.Duration) ListContainersResponse {
	ret := _m.ctrl.Call(_m, "ListContainers", _param0, _param1)
	ret0, _ := ret[0].(ListContainersResponse)
	return ret0
}

func (_mr *_MockDockerClientRecorder) ListContainers(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainers", arg0, arg1)
}

func (_m *MockDockerClient) LoadImage(_param0 io.Reader, _param1 time.Duration) error {
	ret := _m.ctrl.Call(_m, "LoadImage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerClientRecorder) LoadImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadImage", arg0, arg1)
}

func (_m *MockDockerClient) PullImage(_param0 string, _param1 *api.RegistryAuthenticationData) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "PullImage", _param0, _param1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

func (_mr *_MockDockerClientRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PullImage", arg0, arg1)
}

func (_m *MockDockerClient) RemoveContainer(_param0 string, _param1 time.Duration) error {
	ret := _m.ctrl.Call(_m, "RemoveContainer", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerClientRecorder) RemoveContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContainer", arg0, arg1)
}

func (_m *MockDockerClient) RemoveImage(_param0 string, _param1 time.Duration) error {
	ret := _m.ctrl.Call(_m, "RemoveImage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerClientRecorder) RemoveImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveImage", arg0, arg1)
}

func (_m *MockDockerClient) StartContainer(_param0 string, _param1 time.Duration) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "StartContainer", _param0, _param1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

func (_mr *_MockDockerClientRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartContainer", arg0, arg1)
}

func (_m *MockDockerClient) Stats(_param0 string, _param1 context0.Context) (<-chan *go_dockerclient.Stats, error) {
	ret := _m.ctrl.Call(_m, "Stats", _param0, _param1)
	ret0, _ := ret[0].(<-chan *go_dockerclient.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stats", arg0, arg1)
}

func (_m *MockDockerClient) StopContainer(_param0 string, _param1 time.Duration) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "StopContainer", _param0, _param1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

func (_mr *_MockDockerClientRecorder) StopContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopContainer", arg0, arg1)
}

func (_m *MockDockerClient) SupportedVersions() []dockerclient.DockerVersion {
	ret := _m.ctrl.Call(_m, "SupportedVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

func (_mr *_MockDockerClientRecorder) SupportedVersions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SupportedVersions")
}

func (_m *MockDockerClient) Version() (string, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerClientRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

func (_m *MockDockerClient) WithVersion(_param0 dockerclient.DockerVersion) DockerClient {
	ret := _m.ctrl.Call(_m, "WithVersion", _param0)
	ret0, _ := ret[0].(DockerClient)
	return ret0
}

func (_mr *_MockDockerClientRecorder) WithVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithVersion", arg0)
}

// Mock of ImageManager interface
type MockImageManager struct {
	ctrl     *gomock.Controller
	recorder *_MockImageManagerRecorder
}

// Recorder for MockImageManager (not exported)
type _MockImageManagerRecorder struct {
	mock *MockImageManager
}

func NewMockImageManager(ctrl *gomock.Controller) *MockImageManager {
	mock := &MockImageManager{ctrl: ctrl}
	mock.recorder = &_MockImageManagerRecorder{mock}
	return mock
}

func (_m *MockImageManager) EXPECT() *_MockImageManagerRecorder {
	return _m.recorder
}

func (_m *MockImageManager) AddAllImageStates(_param0 []*image.ImageState) {
	_m.ctrl.Call(_m, "AddAllImageStates", _param0)
}

func (_mr *_MockImageManagerRecorder) AddAllImageStates(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddAllImageStates", arg0)
}

func (_m *MockImageManager) GetImageStateFromImageName(_param0 string) *image.ImageState {
	ret := _m.ctrl.Call(_m, "GetImageStateFromImageName", _param0)
	ret0, _ := ret[0].(*image.ImageState)
	return ret0
}

func (_mr *_MockImageManagerRecorder) GetImageStateFromImageName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetImageStateFromImageName", arg0)
}

func (_m *MockImageManager) RecordContainerReference(_param0 *api.Container) error {
	ret := _m.ctrl.Call(_m, "RecordContainerReference", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImageManagerRecorder) RecordContainerReference(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecordContainerReference", arg0)
}

func (_m *MockImageManager) RemoveContainerReferenceFromImageState(_param0 *api.Container) error {
	ret := _m.ctrl.Call(_m, "RemoveContainerReferenceFromImageState", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImageManagerRecorder) RemoveContainerReferenceFromImageState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContainerReferenceFromImageState", arg0)
}

func (_m *MockImageManager) SetSaver(_param0 statemanager.Saver) {
	_m.ctrl.Call(_m, "SetSaver", _param0)
}

func (_mr *_MockImageManagerRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSaver", arg0)
}

func (_m *MockImageManager) StartImageCleanupProcess(_param0 context.Context) {
	_m.ctrl.Call(_m, "StartImageCleanupProcess", _param0)
}

func (_mr *_MockImageManagerRecorder) StartImageCleanupProcess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartImageCleanupProcess", arg0)
}
