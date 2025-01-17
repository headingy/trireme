package monitor

import "github.com/headingy/trireme/policy"

// A Monitor is the interface to implement low level monitoring functions on some well defined primitive.
type Monitor interface {

	// Start starts the monitor.
	Start() error

	// Stop Stops the monitor.
	Stop() error
}

// A ProcessingUnitsHandler is responsible for monitoring creation and deletion of ProcessingUnits.
type ProcessingUnitsHandler interface {

	// SetPURuntime handles the create ProcessingUnit event.
	SetPURuntime(contextID string, runtimeInfo *policy.PURuntime) error

	// HandlePUEvent handles the event generated by the PU.
	HandlePUEvent(contextID string, event Event) error
}

// A SynchronizationHandler can handle a PU synchronization routine.
type SynchronizationHandler interface {

	// HandleSynchronization handles a synchronization routine.
	HandleSynchronization(contextID string, state State, RuntimeReader policy.RuntimeReader, syncType SynchronizationType) error

	// HandleSynchronizationComplete is called when a synchronization job is complete.
	HandleSynchronizationComplete(syncType SynchronizationType)
}
