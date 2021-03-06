// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/cppforlife/bosh-warden-cpi/disk"
	"github.com/cppforlife/bosh-warden-cpi/vm"
)

type FakeVM struct {
	IDStub        func() apiv1.VMCID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 apiv1.VMCID
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct{}
	deleteReturns     struct {
		result1 error
	}
	AttachDiskStub        func(disk.Disk) (string, error)
	attachDiskMutex       sync.RWMutex
	attachDiskArgsForCall []struct {
		arg1 disk.Disk
	}
	attachDiskReturns struct {
		result1 string
		result2 error
	}
	DetachDiskStub        func(disk.Disk) error
	detachDiskMutex       sync.RWMutex
	detachDiskArgsForCall []struct {
		arg1 disk.Disk
	}
	detachDiskReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVM) ID() apiv1.VMCID {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeVM) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeVM) IDReturns(result1 apiv1.VMCID) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 apiv1.VMCID
	}{result1}
}

func (fake *FakeVM) Delete() error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct{}{})
	fake.recordInvocation("Delete", []interface{}{})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub()
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeVM) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeVM) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVM) AttachDisk(arg1 disk.Disk) (string, error) {
	fake.attachDiskMutex.Lock()
	fake.attachDiskArgsForCall = append(fake.attachDiskArgsForCall, struct {
		arg1 disk.Disk
	}{arg1})
	fake.recordInvocation("AttachDisk", []interface{}{arg1})
	fake.attachDiskMutex.Unlock()
	if fake.AttachDiskStub != nil {
		return fake.AttachDiskStub(arg1)
	} else {
		return fake.attachDiskReturns.result1, fake.attachDiskReturns.result2
	}
}

func (fake *FakeVM) AttachDiskCallCount() int {
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	return len(fake.attachDiskArgsForCall)
}

func (fake *FakeVM) AttachDiskArgsForCall(i int) disk.Disk {
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	return fake.attachDiskArgsForCall[i].arg1
}

func (fake *FakeVM) AttachDiskReturns(result1 string, result2 error) {
	fake.AttachDiskStub = nil
	fake.attachDiskReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVM) DetachDisk(arg1 disk.Disk) error {
	fake.detachDiskMutex.Lock()
	fake.detachDiskArgsForCall = append(fake.detachDiskArgsForCall, struct {
		arg1 disk.Disk
	}{arg1})
	fake.recordInvocation("DetachDisk", []interface{}{arg1})
	fake.detachDiskMutex.Unlock()
	if fake.DetachDiskStub != nil {
		return fake.DetachDiskStub(arg1)
	} else {
		return fake.detachDiskReturns.result1
	}
}

func (fake *FakeVM) DetachDiskCallCount() int {
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	return len(fake.detachDiskArgsForCall)
}

func (fake *FakeVM) DetachDiskArgsForCall(i int) disk.Disk {
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	return fake.detachDiskArgsForCall[i].arg1
}

func (fake *FakeVM) DetachDiskReturns(result1 error) {
	fake.DetachDiskStub = nil
	fake.detachDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVM) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVM) recordInvocation(key string, args []interface{}) {
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

var _ vm.VM = new(FakeVM)
