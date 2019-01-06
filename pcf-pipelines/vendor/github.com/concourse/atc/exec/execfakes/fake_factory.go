// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/exec"
)

type FakeFactory struct {
	GetStub        func(lager.Logger, atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.ActionsBuildEventsDelegate, exec.BuildStepDelegate) exec.StepFactory
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.Plan
		arg3 db.Build
		arg4 exec.StepMetadata
		arg5 db.ContainerMetadata
		arg6 exec.ActionsBuildEventsDelegate
		arg7 exec.BuildStepDelegate
	}
	getReturns struct {
		result1 exec.StepFactory
	}
	getReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	PutStub        func(lager.Logger, atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.ActionsBuildEventsDelegate, exec.BuildStepDelegate) exec.StepFactory
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.Plan
		arg3 db.Build
		arg4 exec.StepMetadata
		arg5 db.ContainerMetadata
		arg6 exec.ActionsBuildEventsDelegate
		arg7 exec.BuildStepDelegate
	}
	putReturns struct {
		result1 exec.StepFactory
	}
	putReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	TaskStub        func(lager.Logger, atc.Plan, db.Build, db.ContainerMetadata, exec.TaskBuildEventsDelegate, exec.ActionsBuildEventsDelegate, exec.BuildStepDelegate) exec.StepFactory
	taskMutex       sync.RWMutex
	taskArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.Plan
		arg3 db.Build
		arg4 db.ContainerMetadata
		arg5 exec.TaskBuildEventsDelegate
		arg6 exec.ActionsBuildEventsDelegate
		arg7 exec.BuildStepDelegate
	}
	taskReturns struct {
		result1 exec.StepFactory
	}
	taskReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactory) Get(arg1 lager.Logger, arg2 atc.Plan, arg3 db.Build, arg4 exec.StepMetadata, arg5 db.ContainerMetadata, arg6 exec.ActionsBuildEventsDelegate, arg7 exec.BuildStepDelegate) exec.StepFactory {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.Plan
		arg3 db.Build
		arg4 exec.StepMetadata
		arg5 db.ContainerMetadata
		arg6 exec.ActionsBuildEventsDelegate
		arg7 exec.BuildStepDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeFactory) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFactory) GetArgsForCall(i int) (lager.Logger, atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.ActionsBuildEventsDelegate, exec.BuildStepDelegate) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3, fake.getArgsForCall[i].arg4, fake.getArgsForCall[i].arg5, fake.getArgsForCall[i].arg6, fake.getArgsForCall[i].arg7
}

func (fake *FakeFactory) GetReturns(result1 exec.StepFactory) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) GetReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Put(arg1 lager.Logger, arg2 atc.Plan, arg3 db.Build, arg4 exec.StepMetadata, arg5 db.ContainerMetadata, arg6 exec.ActionsBuildEventsDelegate, arg7 exec.BuildStepDelegate) exec.StepFactory {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.Plan
		arg3 db.Build
		arg4 exec.StepMetadata
		arg5 db.ContainerMetadata
		arg6 exec.ActionsBuildEventsDelegate
		arg7 exec.BuildStepDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.putReturns.result1
}

func (fake *FakeFactory) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeFactory) PutArgsForCall(i int) (lager.Logger, atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.ActionsBuildEventsDelegate, exec.BuildStepDelegate) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3, fake.putArgsForCall[i].arg4, fake.putArgsForCall[i].arg5, fake.putArgsForCall[i].arg6, fake.putArgsForCall[i].arg7
}

func (fake *FakeFactory) PutReturns(result1 exec.StepFactory) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) PutReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Task(arg1 lager.Logger, arg2 atc.Plan, arg3 db.Build, arg4 db.ContainerMetadata, arg5 exec.TaskBuildEventsDelegate, arg6 exec.ActionsBuildEventsDelegate, arg7 exec.BuildStepDelegate) exec.StepFactory {
	fake.taskMutex.Lock()
	ret, specificReturn := fake.taskReturnsOnCall[len(fake.taskArgsForCall)]
	fake.taskArgsForCall = append(fake.taskArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.Plan
		arg3 db.Build
		arg4 db.ContainerMetadata
		arg5 exec.TaskBuildEventsDelegate
		arg6 exec.ActionsBuildEventsDelegate
		arg7 exec.BuildStepDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Task", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.taskMutex.Unlock()
	if fake.TaskStub != nil {
		return fake.TaskStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.taskReturns.result1
}

func (fake *FakeFactory) TaskCallCount() int {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return len(fake.taskArgsForCall)
}

func (fake *FakeFactory) TaskArgsForCall(i int) (lager.Logger, atc.Plan, db.Build, db.ContainerMetadata, exec.TaskBuildEventsDelegate, exec.ActionsBuildEventsDelegate, exec.BuildStepDelegate) {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return fake.taskArgsForCall[i].arg1, fake.taskArgsForCall[i].arg2, fake.taskArgsForCall[i].arg3, fake.taskArgsForCall[i].arg4, fake.taskArgsForCall[i].arg5, fake.taskArgsForCall[i].arg6, fake.taskArgsForCall[i].arg7
}

func (fake *FakeFactory) TaskReturns(result1 exec.StepFactory) {
	fake.TaskStub = nil
	fake.taskReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) TaskReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.TaskStub = nil
	if fake.taskReturnsOnCall == nil {
		fake.taskReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.taskReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFactory) recordInvocation(key string, args []interface{}) {
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

var _ exec.Factory = new(FakeFactory)