// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/orderer/consensus/follower"
	"github.com/hyperledger/fabric/protoutil"
)

type BlockPullerSupport struct {
	ChannelIDStub        func() string
	channelIDMutex       sync.RWMutex
	channelIDArgsForCall []struct {
	}
	channelIDReturns struct {
		result1 string
	}
	channelIDReturnsOnCall map[int]struct {
		result1 string
	}
	SerializeStub        func() ([]byte, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct {
	}
	serializeReturns struct {
		result1 []byte
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SignStub        func([]byte) ([]byte, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		arg1 []byte
	}
	signReturns struct {
		result1 []byte
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	VerifyBlockSignatureStub        func([]*protoutil.SignedData, *common.ConfigEnvelope) error
	verifyBlockSignatureMutex       sync.RWMutex
	verifyBlockSignatureArgsForCall []struct {
		arg1 []*protoutil.SignedData
		arg2 *common.ConfigEnvelope
	}
	verifyBlockSignatureReturns struct {
		result1 error
	}
	verifyBlockSignatureReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockPullerSupport) ChannelID() string {
	fake.channelIDMutex.Lock()
	ret, specificReturn := fake.channelIDReturnsOnCall[len(fake.channelIDArgsForCall)]
	fake.channelIDArgsForCall = append(fake.channelIDArgsForCall, struct {
	}{})
	fake.recordInvocation("ChannelID", []interface{}{})
	fake.channelIDMutex.Unlock()
	if fake.ChannelIDStub != nil {
		return fake.ChannelIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelIDReturns
	return fakeReturns.result1
}

func (fake *BlockPullerSupport) ChannelIDCallCount() int {
	fake.channelIDMutex.RLock()
	defer fake.channelIDMutex.RUnlock()
	return len(fake.channelIDArgsForCall)
}

func (fake *BlockPullerSupport) ChannelIDCalls(stub func() string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = stub
}

func (fake *BlockPullerSupport) ChannelIDReturns(result1 string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = nil
	fake.channelIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *BlockPullerSupport) ChannelIDReturnsOnCall(i int, result1 string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = nil
	if fake.channelIDReturnsOnCall == nil {
		fake.channelIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.channelIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *BlockPullerSupport) Serialize() ([]byte, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct {
	}{})
	fake.recordInvocation("Serialize", []interface{}{})
	fake.serializeMutex.Unlock()
	if fake.SerializeStub != nil {
		return fake.SerializeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.serializeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockPullerSupport) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *BlockPullerSupport) SerializeCalls(stub func() ([]byte, error)) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = stub
}

func (fake *BlockPullerSupport) SerializeReturns(result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerSupport) SerializeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerSupport) Sign(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Sign", []interface{}{arg1Copy})
	fake.signMutex.Unlock()
	if fake.SignStub != nil {
		return fake.SignStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.signReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockPullerSupport) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *BlockPullerSupport) SignCalls(stub func([]byte) ([]byte, error)) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = stub
}

func (fake *BlockPullerSupport) SignArgsForCall(i int) []byte {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	argsForCall := fake.signArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BlockPullerSupport) SignReturns(result1 []byte, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerSupport) SignReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.signMutex.Lock()
	defer fake.signMutex.Unlock()
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerSupport) VerifyBlockSignature(arg1 []*protoutil.SignedData, arg2 *common.ConfigEnvelope) error {
	var arg1Copy []*protoutil.SignedData
	if arg1 != nil {
		arg1Copy = make([]*protoutil.SignedData, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.verifyBlockSignatureMutex.Lock()
	ret, specificReturn := fake.verifyBlockSignatureReturnsOnCall[len(fake.verifyBlockSignatureArgsForCall)]
	fake.verifyBlockSignatureArgsForCall = append(fake.verifyBlockSignatureArgsForCall, struct {
		arg1 []*protoutil.SignedData
		arg2 *common.ConfigEnvelope
	}{arg1Copy, arg2})
	fake.recordInvocation("VerifyBlockSignature", []interface{}{arg1Copy, arg2})
	fake.verifyBlockSignatureMutex.Unlock()
	if fake.VerifyBlockSignatureStub != nil {
		return fake.VerifyBlockSignatureStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.verifyBlockSignatureReturns
	return fakeReturns.result1
}

func (fake *BlockPullerSupport) VerifyBlockSignatureCallCount() int {
	fake.verifyBlockSignatureMutex.RLock()
	defer fake.verifyBlockSignatureMutex.RUnlock()
	return len(fake.verifyBlockSignatureArgsForCall)
}

func (fake *BlockPullerSupport) VerifyBlockSignatureCalls(stub func([]*protoutil.SignedData, *common.ConfigEnvelope) error) {
	fake.verifyBlockSignatureMutex.Lock()
	defer fake.verifyBlockSignatureMutex.Unlock()
	fake.VerifyBlockSignatureStub = stub
}

func (fake *BlockPullerSupport) VerifyBlockSignatureArgsForCall(i int) ([]*protoutil.SignedData, *common.ConfigEnvelope) {
	fake.verifyBlockSignatureMutex.RLock()
	defer fake.verifyBlockSignatureMutex.RUnlock()
	argsForCall := fake.verifyBlockSignatureArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *BlockPullerSupport) VerifyBlockSignatureReturns(result1 error) {
	fake.verifyBlockSignatureMutex.Lock()
	defer fake.verifyBlockSignatureMutex.Unlock()
	fake.VerifyBlockSignatureStub = nil
	fake.verifyBlockSignatureReturns = struct {
		result1 error
	}{result1}
}

func (fake *BlockPullerSupport) VerifyBlockSignatureReturnsOnCall(i int, result1 error) {
	fake.verifyBlockSignatureMutex.Lock()
	defer fake.verifyBlockSignatureMutex.Unlock()
	fake.VerifyBlockSignatureStub = nil
	if fake.verifyBlockSignatureReturnsOnCall == nil {
		fake.verifyBlockSignatureReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyBlockSignatureReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BlockPullerSupport) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.channelIDMutex.RLock()
	defer fake.channelIDMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	fake.verifyBlockSignatureMutex.RLock()
	defer fake.verifyBlockSignatureMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockPullerSupport) recordInvocation(key string, args []interface{}) {
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

var _ follower.BlockPullerSupport = new(BlockPullerSupport)
