// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-cli/pkg/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

type ResourceManagement struct {
	CreateConfigSignatureStub        func(msp.SigningIdentity, string) (*common.ConfigSignature, error)
	createConfigSignatureMutex       sync.RWMutex
	createConfigSignatureArgsForCall []struct {
		arg1 msp.SigningIdentity
		arg2 string
	}
	createConfigSignatureReturns struct {
		result1 *common.ConfigSignature
		result2 error
	}
	createConfigSignatureReturnsOnCall map[int]struct {
		result1 *common.ConfigSignature
		result2 error
	}
	CreateConfigSignatureDataStub        func(msp.SigningIdentity, string) (resource.ConfigSignatureData, error)
	createConfigSignatureDataMutex       sync.RWMutex
	createConfigSignatureDataArgsForCall []struct {
		arg1 msp.SigningIdentity
		arg2 string
	}
	createConfigSignatureDataReturns struct {
		result1 resource.ConfigSignatureData
		result2 error
	}
	createConfigSignatureDataReturnsOnCall map[int]struct {
		result1 resource.ConfigSignatureData
		result2 error
	}
	InstallCCStub        func(resmgmt.InstallCCRequest, ...resmgmt.RequestOption) ([]resmgmt.InstallCCResponse, error)
	installCCMutex       sync.RWMutex
	installCCArgsForCall []struct {
		arg1 resmgmt.InstallCCRequest
		arg2 []resmgmt.RequestOption
	}
	installCCReturns struct {
		result1 []resmgmt.InstallCCResponse
		result2 error
	}
	installCCReturnsOnCall map[int]struct {
		result1 []resmgmt.InstallCCResponse
		result2 error
	}
	InstantiateCCStub        func(string, resmgmt.InstantiateCCRequest, ...resmgmt.RequestOption) (resmgmt.InstantiateCCResponse, error)
	instantiateCCMutex       sync.RWMutex
	instantiateCCArgsForCall []struct {
		arg1 string
		arg2 resmgmt.InstantiateCCRequest
		arg3 []resmgmt.RequestOption
	}
	instantiateCCReturns struct {
		result1 resmgmt.InstantiateCCResponse
		result2 error
	}
	instantiateCCReturnsOnCall map[int]struct {
		result1 resmgmt.InstantiateCCResponse
		result2 error
	}
	JoinChannelStub        func(string, ...resmgmt.RequestOption) error
	joinChannelMutex       sync.RWMutex
	joinChannelArgsForCall []struct {
		arg1 string
		arg2 []resmgmt.RequestOption
	}
	joinChannelReturns struct {
		result1 error
	}
	joinChannelReturnsOnCall map[int]struct {
		result1 error
	}
	QueryChannelsStub        func(...resmgmt.RequestOption) (*peer.ChannelQueryResponse, error)
	queryChannelsMutex       sync.RWMutex
	queryChannelsArgsForCall []struct {
		arg1 []resmgmt.RequestOption
	}
	queryChannelsReturns struct {
		result1 *peer.ChannelQueryResponse
		result2 error
	}
	queryChannelsReturnsOnCall map[int]struct {
		result1 *peer.ChannelQueryResponse
		result2 error
	}
	QueryCollectionsConfigStub        func(string, string, ...resmgmt.RequestOption) (*common.CollectionConfigPackage, error)
	queryCollectionsConfigMutex       sync.RWMutex
	queryCollectionsConfigArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []resmgmt.RequestOption
	}
	queryCollectionsConfigReturns struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	queryCollectionsConfigReturnsOnCall map[int]struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	QueryConfigFromOrdererStub        func(string, ...resmgmt.RequestOption) (fab.ChannelCfg, error)
	queryConfigFromOrdererMutex       sync.RWMutex
	queryConfigFromOrdererArgsForCall []struct {
		arg1 string
		arg2 []resmgmt.RequestOption
	}
	queryConfigFromOrdererReturns struct {
		result1 fab.ChannelCfg
		result2 error
	}
	queryConfigFromOrdererReturnsOnCall map[int]struct {
		result1 fab.ChannelCfg
		result2 error
	}
	QueryInstalledChaincodesStub        func(...resmgmt.RequestOption) (*peer.ChaincodeQueryResponse, error)
	queryInstalledChaincodesMutex       sync.RWMutex
	queryInstalledChaincodesArgsForCall []struct {
		arg1 []resmgmt.RequestOption
	}
	queryInstalledChaincodesReturns struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}
	queryInstalledChaincodesReturnsOnCall map[int]struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}
	QueryInstantiatedChaincodesStub        func(string, ...resmgmt.RequestOption) (*peer.ChaincodeQueryResponse, error)
	queryInstantiatedChaincodesMutex       sync.RWMutex
	queryInstantiatedChaincodesArgsForCall []struct {
		arg1 string
		arg2 []resmgmt.RequestOption
	}
	queryInstantiatedChaincodesReturns struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}
	queryInstantiatedChaincodesReturnsOnCall map[int]struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}
	SaveChannelStub        func(resmgmt.SaveChannelRequest, ...resmgmt.RequestOption) (resmgmt.SaveChannelResponse, error)
	saveChannelMutex       sync.RWMutex
	saveChannelArgsForCall []struct {
		arg1 resmgmt.SaveChannelRequest
		arg2 []resmgmt.RequestOption
	}
	saveChannelReturns struct {
		result1 resmgmt.SaveChannelResponse
		result2 error
	}
	saveChannelReturnsOnCall map[int]struct {
		result1 resmgmt.SaveChannelResponse
		result2 error
	}
	UpgradeCCStub        func(string, resmgmt.UpgradeCCRequest, ...resmgmt.RequestOption) (resmgmt.UpgradeCCResponse, error)
	upgradeCCMutex       sync.RWMutex
	upgradeCCArgsForCall []struct {
		arg1 string
		arg2 resmgmt.UpgradeCCRequest
		arg3 []resmgmt.RequestOption
	}
	upgradeCCReturns struct {
		result1 resmgmt.UpgradeCCResponse
		result2 error
	}
	upgradeCCReturnsOnCall map[int]struct {
		result1 resmgmt.UpgradeCCResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ResourceManagement) CreateConfigSignature(arg1 msp.SigningIdentity, arg2 string) (*common.ConfigSignature, error) {
	fake.createConfigSignatureMutex.Lock()
	ret, specificReturn := fake.createConfigSignatureReturnsOnCall[len(fake.createConfigSignatureArgsForCall)]
	fake.createConfigSignatureArgsForCall = append(fake.createConfigSignatureArgsForCall, struct {
		arg1 msp.SigningIdentity
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateConfigSignature", []interface{}{arg1, arg2})
	fake.createConfigSignatureMutex.Unlock()
	if fake.CreateConfigSignatureStub != nil {
		return fake.CreateConfigSignatureStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createConfigSignatureReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) CreateConfigSignatureCallCount() int {
	fake.createConfigSignatureMutex.RLock()
	defer fake.createConfigSignatureMutex.RUnlock()
	return len(fake.createConfigSignatureArgsForCall)
}

func (fake *ResourceManagement) CreateConfigSignatureCalls(stub func(msp.SigningIdentity, string) (*common.ConfigSignature, error)) {
	fake.createConfigSignatureMutex.Lock()
	defer fake.createConfigSignatureMutex.Unlock()
	fake.CreateConfigSignatureStub = stub
}

func (fake *ResourceManagement) CreateConfigSignatureArgsForCall(i int) (msp.SigningIdentity, string) {
	fake.createConfigSignatureMutex.RLock()
	defer fake.createConfigSignatureMutex.RUnlock()
	argsForCall := fake.createConfigSignatureArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) CreateConfigSignatureReturns(result1 *common.ConfigSignature, result2 error) {
	fake.createConfigSignatureMutex.Lock()
	defer fake.createConfigSignatureMutex.Unlock()
	fake.CreateConfigSignatureStub = nil
	fake.createConfigSignatureReturns = struct {
		result1 *common.ConfigSignature
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) CreateConfigSignatureReturnsOnCall(i int, result1 *common.ConfigSignature, result2 error) {
	fake.createConfigSignatureMutex.Lock()
	defer fake.createConfigSignatureMutex.Unlock()
	fake.CreateConfigSignatureStub = nil
	if fake.createConfigSignatureReturnsOnCall == nil {
		fake.createConfigSignatureReturnsOnCall = make(map[int]struct {
			result1 *common.ConfigSignature
			result2 error
		})
	}
	fake.createConfigSignatureReturnsOnCall[i] = struct {
		result1 *common.ConfigSignature
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) CreateConfigSignatureData(arg1 msp.SigningIdentity, arg2 string) (resource.ConfigSignatureData, error) {
	fake.createConfigSignatureDataMutex.Lock()
	ret, specificReturn := fake.createConfigSignatureDataReturnsOnCall[len(fake.createConfigSignatureDataArgsForCall)]
	fake.createConfigSignatureDataArgsForCall = append(fake.createConfigSignatureDataArgsForCall, struct {
		arg1 msp.SigningIdentity
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateConfigSignatureData", []interface{}{arg1, arg2})
	fake.createConfigSignatureDataMutex.Unlock()
	if fake.CreateConfigSignatureDataStub != nil {
		return fake.CreateConfigSignatureDataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createConfigSignatureDataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) CreateConfigSignatureDataCallCount() int {
	fake.createConfigSignatureDataMutex.RLock()
	defer fake.createConfigSignatureDataMutex.RUnlock()
	return len(fake.createConfigSignatureDataArgsForCall)
}

func (fake *ResourceManagement) CreateConfigSignatureDataCalls(stub func(msp.SigningIdentity, string) (resource.ConfigSignatureData, error)) {
	fake.createConfigSignatureDataMutex.Lock()
	defer fake.createConfigSignatureDataMutex.Unlock()
	fake.CreateConfigSignatureDataStub = stub
}

func (fake *ResourceManagement) CreateConfigSignatureDataArgsForCall(i int) (msp.SigningIdentity, string) {
	fake.createConfigSignatureDataMutex.RLock()
	defer fake.createConfigSignatureDataMutex.RUnlock()
	argsForCall := fake.createConfigSignatureDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) CreateConfigSignatureDataReturns(result1 resource.ConfigSignatureData, result2 error) {
	fake.createConfigSignatureDataMutex.Lock()
	defer fake.createConfigSignatureDataMutex.Unlock()
	fake.CreateConfigSignatureDataStub = nil
	fake.createConfigSignatureDataReturns = struct {
		result1 resource.ConfigSignatureData
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) CreateConfigSignatureDataReturnsOnCall(i int, result1 resource.ConfigSignatureData, result2 error) {
	fake.createConfigSignatureDataMutex.Lock()
	defer fake.createConfigSignatureDataMutex.Unlock()
	fake.CreateConfigSignatureDataStub = nil
	if fake.createConfigSignatureDataReturnsOnCall == nil {
		fake.createConfigSignatureDataReturnsOnCall = make(map[int]struct {
			result1 resource.ConfigSignatureData
			result2 error
		})
	}
	fake.createConfigSignatureDataReturnsOnCall[i] = struct {
		result1 resource.ConfigSignatureData
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) InstallCC(arg1 resmgmt.InstallCCRequest, arg2 ...resmgmt.RequestOption) ([]resmgmt.InstallCCResponse, error) {
	fake.installCCMutex.Lock()
	ret, specificReturn := fake.installCCReturnsOnCall[len(fake.installCCArgsForCall)]
	fake.installCCArgsForCall = append(fake.installCCArgsForCall, struct {
		arg1 resmgmt.InstallCCRequest
		arg2 []resmgmt.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("InstallCC", []interface{}{arg1, arg2})
	fake.installCCMutex.Unlock()
	if fake.InstallCCStub != nil {
		return fake.InstallCCStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.installCCReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) InstallCCCallCount() int {
	fake.installCCMutex.RLock()
	defer fake.installCCMutex.RUnlock()
	return len(fake.installCCArgsForCall)
}

func (fake *ResourceManagement) InstallCCCalls(stub func(resmgmt.InstallCCRequest, ...resmgmt.RequestOption) ([]resmgmt.InstallCCResponse, error)) {
	fake.installCCMutex.Lock()
	defer fake.installCCMutex.Unlock()
	fake.InstallCCStub = stub
}

func (fake *ResourceManagement) InstallCCArgsForCall(i int) (resmgmt.InstallCCRequest, []resmgmt.RequestOption) {
	fake.installCCMutex.RLock()
	defer fake.installCCMutex.RUnlock()
	argsForCall := fake.installCCArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) InstallCCReturns(result1 []resmgmt.InstallCCResponse, result2 error) {
	fake.installCCMutex.Lock()
	defer fake.installCCMutex.Unlock()
	fake.InstallCCStub = nil
	fake.installCCReturns = struct {
		result1 []resmgmt.InstallCCResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) InstallCCReturnsOnCall(i int, result1 []resmgmt.InstallCCResponse, result2 error) {
	fake.installCCMutex.Lock()
	defer fake.installCCMutex.Unlock()
	fake.InstallCCStub = nil
	if fake.installCCReturnsOnCall == nil {
		fake.installCCReturnsOnCall = make(map[int]struct {
			result1 []resmgmt.InstallCCResponse
			result2 error
		})
	}
	fake.installCCReturnsOnCall[i] = struct {
		result1 []resmgmt.InstallCCResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) InstantiateCC(arg1 string, arg2 resmgmt.InstantiateCCRequest, arg3 ...resmgmt.RequestOption) (resmgmt.InstantiateCCResponse, error) {
	fake.instantiateCCMutex.Lock()
	ret, specificReturn := fake.instantiateCCReturnsOnCall[len(fake.instantiateCCArgsForCall)]
	fake.instantiateCCArgsForCall = append(fake.instantiateCCArgsForCall, struct {
		arg1 string
		arg2 resmgmt.InstantiateCCRequest
		arg3 []resmgmt.RequestOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("InstantiateCC", []interface{}{arg1, arg2, arg3})
	fake.instantiateCCMutex.Unlock()
	if fake.InstantiateCCStub != nil {
		return fake.InstantiateCCStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instantiateCCReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) InstantiateCCCallCount() int {
	fake.instantiateCCMutex.RLock()
	defer fake.instantiateCCMutex.RUnlock()
	return len(fake.instantiateCCArgsForCall)
}

func (fake *ResourceManagement) InstantiateCCCalls(stub func(string, resmgmt.InstantiateCCRequest, ...resmgmt.RequestOption) (resmgmt.InstantiateCCResponse, error)) {
	fake.instantiateCCMutex.Lock()
	defer fake.instantiateCCMutex.Unlock()
	fake.InstantiateCCStub = stub
}

func (fake *ResourceManagement) InstantiateCCArgsForCall(i int) (string, resmgmt.InstantiateCCRequest, []resmgmt.RequestOption) {
	fake.instantiateCCMutex.RLock()
	defer fake.instantiateCCMutex.RUnlock()
	argsForCall := fake.instantiateCCArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ResourceManagement) InstantiateCCReturns(result1 resmgmt.InstantiateCCResponse, result2 error) {
	fake.instantiateCCMutex.Lock()
	defer fake.instantiateCCMutex.Unlock()
	fake.InstantiateCCStub = nil
	fake.instantiateCCReturns = struct {
		result1 resmgmt.InstantiateCCResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) InstantiateCCReturnsOnCall(i int, result1 resmgmt.InstantiateCCResponse, result2 error) {
	fake.instantiateCCMutex.Lock()
	defer fake.instantiateCCMutex.Unlock()
	fake.InstantiateCCStub = nil
	if fake.instantiateCCReturnsOnCall == nil {
		fake.instantiateCCReturnsOnCall = make(map[int]struct {
			result1 resmgmt.InstantiateCCResponse
			result2 error
		})
	}
	fake.instantiateCCReturnsOnCall[i] = struct {
		result1 resmgmt.InstantiateCCResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) JoinChannel(arg1 string, arg2 ...resmgmt.RequestOption) error {
	fake.joinChannelMutex.Lock()
	ret, specificReturn := fake.joinChannelReturnsOnCall[len(fake.joinChannelArgsForCall)]
	fake.joinChannelArgsForCall = append(fake.joinChannelArgsForCall, struct {
		arg1 string
		arg2 []resmgmt.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("JoinChannel", []interface{}{arg1, arg2})
	fake.joinChannelMutex.Unlock()
	if fake.JoinChannelStub != nil {
		return fake.JoinChannelStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.joinChannelReturns
	return fakeReturns.result1
}

func (fake *ResourceManagement) JoinChannelCallCount() int {
	fake.joinChannelMutex.RLock()
	defer fake.joinChannelMutex.RUnlock()
	return len(fake.joinChannelArgsForCall)
}

func (fake *ResourceManagement) JoinChannelCalls(stub func(string, ...resmgmt.RequestOption) error) {
	fake.joinChannelMutex.Lock()
	defer fake.joinChannelMutex.Unlock()
	fake.JoinChannelStub = stub
}

func (fake *ResourceManagement) JoinChannelArgsForCall(i int) (string, []resmgmt.RequestOption) {
	fake.joinChannelMutex.RLock()
	defer fake.joinChannelMutex.RUnlock()
	argsForCall := fake.joinChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) JoinChannelReturns(result1 error) {
	fake.joinChannelMutex.Lock()
	defer fake.joinChannelMutex.Unlock()
	fake.JoinChannelStub = nil
	fake.joinChannelReturns = struct {
		result1 error
	}{result1}
}

func (fake *ResourceManagement) JoinChannelReturnsOnCall(i int, result1 error) {
	fake.joinChannelMutex.Lock()
	defer fake.joinChannelMutex.Unlock()
	fake.JoinChannelStub = nil
	if fake.joinChannelReturnsOnCall == nil {
		fake.joinChannelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.joinChannelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ResourceManagement) QueryChannels(arg1 ...resmgmt.RequestOption) (*peer.ChannelQueryResponse, error) {
	fake.queryChannelsMutex.Lock()
	ret, specificReturn := fake.queryChannelsReturnsOnCall[len(fake.queryChannelsArgsForCall)]
	fake.queryChannelsArgsForCall = append(fake.queryChannelsArgsForCall, struct {
		arg1 []resmgmt.RequestOption
	}{arg1})
	fake.recordInvocation("QueryChannels", []interface{}{arg1})
	fake.queryChannelsMutex.Unlock()
	if fake.QueryChannelsStub != nil {
		return fake.QueryChannelsStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryChannelsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) QueryChannelsCallCount() int {
	fake.queryChannelsMutex.RLock()
	defer fake.queryChannelsMutex.RUnlock()
	return len(fake.queryChannelsArgsForCall)
}

func (fake *ResourceManagement) QueryChannelsCalls(stub func(...resmgmt.RequestOption) (*peer.ChannelQueryResponse, error)) {
	fake.queryChannelsMutex.Lock()
	defer fake.queryChannelsMutex.Unlock()
	fake.QueryChannelsStub = stub
}

func (fake *ResourceManagement) QueryChannelsArgsForCall(i int) []resmgmt.RequestOption {
	fake.queryChannelsMutex.RLock()
	defer fake.queryChannelsMutex.RUnlock()
	argsForCall := fake.queryChannelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ResourceManagement) QueryChannelsReturns(result1 *peer.ChannelQueryResponse, result2 error) {
	fake.queryChannelsMutex.Lock()
	defer fake.queryChannelsMutex.Unlock()
	fake.QueryChannelsStub = nil
	fake.queryChannelsReturns = struct {
		result1 *peer.ChannelQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryChannelsReturnsOnCall(i int, result1 *peer.ChannelQueryResponse, result2 error) {
	fake.queryChannelsMutex.Lock()
	defer fake.queryChannelsMutex.Unlock()
	fake.QueryChannelsStub = nil
	if fake.queryChannelsReturnsOnCall == nil {
		fake.queryChannelsReturnsOnCall = make(map[int]struct {
			result1 *peer.ChannelQueryResponse
			result2 error
		})
	}
	fake.queryChannelsReturnsOnCall[i] = struct {
		result1 *peer.ChannelQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryCollectionsConfig(arg1 string, arg2 string, arg3 ...resmgmt.RequestOption) (*common.CollectionConfigPackage, error) {
	fake.queryCollectionsConfigMutex.Lock()
	ret, specificReturn := fake.queryCollectionsConfigReturnsOnCall[len(fake.queryCollectionsConfigArgsForCall)]
	fake.queryCollectionsConfigArgsForCall = append(fake.queryCollectionsConfigArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []resmgmt.RequestOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("QueryCollectionsConfig", []interface{}{arg1, arg2, arg3})
	fake.queryCollectionsConfigMutex.Unlock()
	if fake.QueryCollectionsConfigStub != nil {
		return fake.QueryCollectionsConfigStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryCollectionsConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) QueryCollectionsConfigCallCount() int {
	fake.queryCollectionsConfigMutex.RLock()
	defer fake.queryCollectionsConfigMutex.RUnlock()
	return len(fake.queryCollectionsConfigArgsForCall)
}

func (fake *ResourceManagement) QueryCollectionsConfigCalls(stub func(string, string, ...resmgmt.RequestOption) (*common.CollectionConfigPackage, error)) {
	fake.queryCollectionsConfigMutex.Lock()
	defer fake.queryCollectionsConfigMutex.Unlock()
	fake.QueryCollectionsConfigStub = stub
}

func (fake *ResourceManagement) QueryCollectionsConfigArgsForCall(i int) (string, string, []resmgmt.RequestOption) {
	fake.queryCollectionsConfigMutex.RLock()
	defer fake.queryCollectionsConfigMutex.RUnlock()
	argsForCall := fake.queryCollectionsConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ResourceManagement) QueryCollectionsConfigReturns(result1 *common.CollectionConfigPackage, result2 error) {
	fake.queryCollectionsConfigMutex.Lock()
	defer fake.queryCollectionsConfigMutex.Unlock()
	fake.QueryCollectionsConfigStub = nil
	fake.queryCollectionsConfigReturns = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryCollectionsConfigReturnsOnCall(i int, result1 *common.CollectionConfigPackage, result2 error) {
	fake.queryCollectionsConfigMutex.Lock()
	defer fake.queryCollectionsConfigMutex.Unlock()
	fake.QueryCollectionsConfigStub = nil
	if fake.queryCollectionsConfigReturnsOnCall == nil {
		fake.queryCollectionsConfigReturnsOnCall = make(map[int]struct {
			result1 *common.CollectionConfigPackage
			result2 error
		})
	}
	fake.queryCollectionsConfigReturnsOnCall[i] = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryConfigFromOrderer(arg1 string, arg2 ...resmgmt.RequestOption) (fab.ChannelCfg, error) {
	fake.queryConfigFromOrdererMutex.Lock()
	ret, specificReturn := fake.queryConfigFromOrdererReturnsOnCall[len(fake.queryConfigFromOrdererArgsForCall)]
	fake.queryConfigFromOrdererArgsForCall = append(fake.queryConfigFromOrdererArgsForCall, struct {
		arg1 string
		arg2 []resmgmt.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("QueryConfigFromOrderer", []interface{}{arg1, arg2})
	fake.queryConfigFromOrdererMutex.Unlock()
	if fake.QueryConfigFromOrdererStub != nil {
		return fake.QueryConfigFromOrdererStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryConfigFromOrdererReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) QueryConfigFromOrdererCallCount() int {
	fake.queryConfigFromOrdererMutex.RLock()
	defer fake.queryConfigFromOrdererMutex.RUnlock()
	return len(fake.queryConfigFromOrdererArgsForCall)
}

func (fake *ResourceManagement) QueryConfigFromOrdererCalls(stub func(string, ...resmgmt.RequestOption) (fab.ChannelCfg, error)) {
	fake.queryConfigFromOrdererMutex.Lock()
	defer fake.queryConfigFromOrdererMutex.Unlock()
	fake.QueryConfigFromOrdererStub = stub
}

func (fake *ResourceManagement) QueryConfigFromOrdererArgsForCall(i int) (string, []resmgmt.RequestOption) {
	fake.queryConfigFromOrdererMutex.RLock()
	defer fake.queryConfigFromOrdererMutex.RUnlock()
	argsForCall := fake.queryConfigFromOrdererArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) QueryConfigFromOrdererReturns(result1 fab.ChannelCfg, result2 error) {
	fake.queryConfigFromOrdererMutex.Lock()
	defer fake.queryConfigFromOrdererMutex.Unlock()
	fake.QueryConfigFromOrdererStub = nil
	fake.queryConfigFromOrdererReturns = struct {
		result1 fab.ChannelCfg
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryConfigFromOrdererReturnsOnCall(i int, result1 fab.ChannelCfg, result2 error) {
	fake.queryConfigFromOrdererMutex.Lock()
	defer fake.queryConfigFromOrdererMutex.Unlock()
	fake.QueryConfigFromOrdererStub = nil
	if fake.queryConfigFromOrdererReturnsOnCall == nil {
		fake.queryConfigFromOrdererReturnsOnCall = make(map[int]struct {
			result1 fab.ChannelCfg
			result2 error
		})
	}
	fake.queryConfigFromOrdererReturnsOnCall[i] = struct {
		result1 fab.ChannelCfg
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryInstalledChaincodes(arg1 ...resmgmt.RequestOption) (*peer.ChaincodeQueryResponse, error) {
	fake.queryInstalledChaincodesMutex.Lock()
	ret, specificReturn := fake.queryInstalledChaincodesReturnsOnCall[len(fake.queryInstalledChaincodesArgsForCall)]
	fake.queryInstalledChaincodesArgsForCall = append(fake.queryInstalledChaincodesArgsForCall, struct {
		arg1 []resmgmt.RequestOption
	}{arg1})
	fake.recordInvocation("QueryInstalledChaincodes", []interface{}{arg1})
	fake.queryInstalledChaincodesMutex.Unlock()
	if fake.QueryInstalledChaincodesStub != nil {
		return fake.QueryInstalledChaincodesStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryInstalledChaincodesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) QueryInstalledChaincodesCallCount() int {
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	return len(fake.queryInstalledChaincodesArgsForCall)
}

func (fake *ResourceManagement) QueryInstalledChaincodesCalls(stub func(...resmgmt.RequestOption) (*peer.ChaincodeQueryResponse, error)) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = stub
}

func (fake *ResourceManagement) QueryInstalledChaincodesArgsForCall(i int) []resmgmt.RequestOption {
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	argsForCall := fake.queryInstalledChaincodesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ResourceManagement) QueryInstalledChaincodesReturns(result1 *peer.ChaincodeQueryResponse, result2 error) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = nil
	fake.queryInstalledChaincodesReturns = struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryInstalledChaincodesReturnsOnCall(i int, result1 *peer.ChaincodeQueryResponse, result2 error) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = nil
	if fake.queryInstalledChaincodesReturnsOnCall == nil {
		fake.queryInstalledChaincodesReturnsOnCall = make(map[int]struct {
			result1 *peer.ChaincodeQueryResponse
			result2 error
		})
	}
	fake.queryInstalledChaincodesReturnsOnCall[i] = struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryInstantiatedChaincodes(arg1 string, arg2 ...resmgmt.RequestOption) (*peer.ChaincodeQueryResponse, error) {
	fake.queryInstantiatedChaincodesMutex.Lock()
	ret, specificReturn := fake.queryInstantiatedChaincodesReturnsOnCall[len(fake.queryInstantiatedChaincodesArgsForCall)]
	fake.queryInstantiatedChaincodesArgsForCall = append(fake.queryInstantiatedChaincodesArgsForCall, struct {
		arg1 string
		arg2 []resmgmt.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("QueryInstantiatedChaincodes", []interface{}{arg1, arg2})
	fake.queryInstantiatedChaincodesMutex.Unlock()
	if fake.QueryInstantiatedChaincodesStub != nil {
		return fake.QueryInstantiatedChaincodesStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryInstantiatedChaincodesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) QueryInstantiatedChaincodesCallCount() int {
	fake.queryInstantiatedChaincodesMutex.RLock()
	defer fake.queryInstantiatedChaincodesMutex.RUnlock()
	return len(fake.queryInstantiatedChaincodesArgsForCall)
}

func (fake *ResourceManagement) QueryInstantiatedChaincodesCalls(stub func(string, ...resmgmt.RequestOption) (*peer.ChaincodeQueryResponse, error)) {
	fake.queryInstantiatedChaincodesMutex.Lock()
	defer fake.queryInstantiatedChaincodesMutex.Unlock()
	fake.QueryInstantiatedChaincodesStub = stub
}

func (fake *ResourceManagement) QueryInstantiatedChaincodesArgsForCall(i int) (string, []resmgmt.RequestOption) {
	fake.queryInstantiatedChaincodesMutex.RLock()
	defer fake.queryInstantiatedChaincodesMutex.RUnlock()
	argsForCall := fake.queryInstantiatedChaincodesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) QueryInstantiatedChaincodesReturns(result1 *peer.ChaincodeQueryResponse, result2 error) {
	fake.queryInstantiatedChaincodesMutex.Lock()
	defer fake.queryInstantiatedChaincodesMutex.Unlock()
	fake.QueryInstantiatedChaincodesStub = nil
	fake.queryInstantiatedChaincodesReturns = struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) QueryInstantiatedChaincodesReturnsOnCall(i int, result1 *peer.ChaincodeQueryResponse, result2 error) {
	fake.queryInstantiatedChaincodesMutex.Lock()
	defer fake.queryInstantiatedChaincodesMutex.Unlock()
	fake.QueryInstantiatedChaincodesStub = nil
	if fake.queryInstantiatedChaincodesReturnsOnCall == nil {
		fake.queryInstantiatedChaincodesReturnsOnCall = make(map[int]struct {
			result1 *peer.ChaincodeQueryResponse
			result2 error
		})
	}
	fake.queryInstantiatedChaincodesReturnsOnCall[i] = struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) SaveChannel(arg1 resmgmt.SaveChannelRequest, arg2 ...resmgmt.RequestOption) (resmgmt.SaveChannelResponse, error) {
	fake.saveChannelMutex.Lock()
	ret, specificReturn := fake.saveChannelReturnsOnCall[len(fake.saveChannelArgsForCall)]
	fake.saveChannelArgsForCall = append(fake.saveChannelArgsForCall, struct {
		arg1 resmgmt.SaveChannelRequest
		arg2 []resmgmt.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("SaveChannel", []interface{}{arg1, arg2})
	fake.saveChannelMutex.Unlock()
	if fake.SaveChannelStub != nil {
		return fake.SaveChannelStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.saveChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) SaveChannelCallCount() int {
	fake.saveChannelMutex.RLock()
	defer fake.saveChannelMutex.RUnlock()
	return len(fake.saveChannelArgsForCall)
}

func (fake *ResourceManagement) SaveChannelCalls(stub func(resmgmt.SaveChannelRequest, ...resmgmt.RequestOption) (resmgmt.SaveChannelResponse, error)) {
	fake.saveChannelMutex.Lock()
	defer fake.saveChannelMutex.Unlock()
	fake.SaveChannelStub = stub
}

func (fake *ResourceManagement) SaveChannelArgsForCall(i int) (resmgmt.SaveChannelRequest, []resmgmt.RequestOption) {
	fake.saveChannelMutex.RLock()
	defer fake.saveChannelMutex.RUnlock()
	argsForCall := fake.saveChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ResourceManagement) SaveChannelReturns(result1 resmgmt.SaveChannelResponse, result2 error) {
	fake.saveChannelMutex.Lock()
	defer fake.saveChannelMutex.Unlock()
	fake.SaveChannelStub = nil
	fake.saveChannelReturns = struct {
		result1 resmgmt.SaveChannelResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) SaveChannelReturnsOnCall(i int, result1 resmgmt.SaveChannelResponse, result2 error) {
	fake.saveChannelMutex.Lock()
	defer fake.saveChannelMutex.Unlock()
	fake.SaveChannelStub = nil
	if fake.saveChannelReturnsOnCall == nil {
		fake.saveChannelReturnsOnCall = make(map[int]struct {
			result1 resmgmt.SaveChannelResponse
			result2 error
		})
	}
	fake.saveChannelReturnsOnCall[i] = struct {
		result1 resmgmt.SaveChannelResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) UpgradeCC(arg1 string, arg2 resmgmt.UpgradeCCRequest, arg3 ...resmgmt.RequestOption) (resmgmt.UpgradeCCResponse, error) {
	fake.upgradeCCMutex.Lock()
	ret, specificReturn := fake.upgradeCCReturnsOnCall[len(fake.upgradeCCArgsForCall)]
	fake.upgradeCCArgsForCall = append(fake.upgradeCCArgsForCall, struct {
		arg1 string
		arg2 resmgmt.UpgradeCCRequest
		arg3 []resmgmt.RequestOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpgradeCC", []interface{}{arg1, arg2, arg3})
	fake.upgradeCCMutex.Unlock()
	if fake.UpgradeCCStub != nil {
		return fake.UpgradeCCStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.upgradeCCReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ResourceManagement) UpgradeCCCallCount() int {
	fake.upgradeCCMutex.RLock()
	defer fake.upgradeCCMutex.RUnlock()
	return len(fake.upgradeCCArgsForCall)
}

func (fake *ResourceManagement) UpgradeCCCalls(stub func(string, resmgmt.UpgradeCCRequest, ...resmgmt.RequestOption) (resmgmt.UpgradeCCResponse, error)) {
	fake.upgradeCCMutex.Lock()
	defer fake.upgradeCCMutex.Unlock()
	fake.UpgradeCCStub = stub
}

func (fake *ResourceManagement) UpgradeCCArgsForCall(i int) (string, resmgmt.UpgradeCCRequest, []resmgmt.RequestOption) {
	fake.upgradeCCMutex.RLock()
	defer fake.upgradeCCMutex.RUnlock()
	argsForCall := fake.upgradeCCArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ResourceManagement) UpgradeCCReturns(result1 resmgmt.UpgradeCCResponse, result2 error) {
	fake.upgradeCCMutex.Lock()
	defer fake.upgradeCCMutex.Unlock()
	fake.UpgradeCCStub = nil
	fake.upgradeCCReturns = struct {
		result1 resmgmt.UpgradeCCResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) UpgradeCCReturnsOnCall(i int, result1 resmgmt.UpgradeCCResponse, result2 error) {
	fake.upgradeCCMutex.Lock()
	defer fake.upgradeCCMutex.Unlock()
	fake.UpgradeCCStub = nil
	if fake.upgradeCCReturnsOnCall == nil {
		fake.upgradeCCReturnsOnCall = make(map[int]struct {
			result1 resmgmt.UpgradeCCResponse
			result2 error
		})
	}
	fake.upgradeCCReturnsOnCall[i] = struct {
		result1 resmgmt.UpgradeCCResponse
		result2 error
	}{result1, result2}
}

func (fake *ResourceManagement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createConfigSignatureMutex.RLock()
	defer fake.createConfigSignatureMutex.RUnlock()
	fake.createConfigSignatureDataMutex.RLock()
	defer fake.createConfigSignatureDataMutex.RUnlock()
	fake.installCCMutex.RLock()
	defer fake.installCCMutex.RUnlock()
	fake.instantiateCCMutex.RLock()
	defer fake.instantiateCCMutex.RUnlock()
	fake.joinChannelMutex.RLock()
	defer fake.joinChannelMutex.RUnlock()
	fake.queryChannelsMutex.RLock()
	defer fake.queryChannelsMutex.RUnlock()
	fake.queryCollectionsConfigMutex.RLock()
	defer fake.queryCollectionsConfigMutex.RUnlock()
	fake.queryConfigFromOrdererMutex.RLock()
	defer fake.queryConfigFromOrdererMutex.RUnlock()
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	fake.queryInstantiatedChaincodesMutex.RLock()
	defer fake.queryInstantiatedChaincodesMutex.RUnlock()
	fake.saveChannelMutex.RLock()
	defer fake.saveChannelMutex.RUnlock()
	fake.upgradeCCMutex.RLock()
	defer fake.upgradeCCMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ResourceManagement) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ fabric.ResourceManagement = new(ResourceManagement)
