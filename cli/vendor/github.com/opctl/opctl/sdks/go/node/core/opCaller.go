package core

import (
	"context"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/outputs"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
)

//counterfeiter:generate -o internal/fakes/opCaller.go . opCaller
type opCaller interface {
	// Executes an op call
	Call(
		ctx context.Context,
		opCall *model.OpCall,
		inboundScope map[string]*model.Value,
		parentCallID *string,
		opCallSpec *model.OpCallSpec,
	) (
		map[string]*model.Value,
		error,
	)
}

func newOpCaller(
	stateStore stateStore,
	caller caller,
	dataDirPath string,
) opCaller {
	return _opCaller{
		caller:             caller,
		stateStore:         stateStore,
		callScratchDir:     filepath.Join(dataDirPath, "call"),
		outputsInterpreter: outputs.NewInterpreter(),
		opFileGetter:       opfile.NewGetter(),
	}
}

type _opCaller struct {
	stateStore         stateStore
	caller             caller
	callScratchDir     string
	outputsInterpreter outputs.Interpreter
	opFileGetter       opfile.Getter
}

func (oc _opCaller) Call(
	ctx context.Context,
	opCall *model.OpCall,
	inboundScope map[string]*model.Value,
	parentCallID *string,
	opCallSpec *model.OpCallSpec,
) (
	map[string]*model.Value,
	error,
) {
	var err error
	outboundScope := map[string]*model.Value{}

	// form scope for op call by combining defined inputs & op dir
	opCallScope := map[string]*model.Value{}
	for varName, varData := range opCall.Inputs {
		opCallScope[varName] = varData
	}
	opCallScope["/"] = &model.Value{
		Dir: &opCall.OpPath,
	}

	opOutputs, err := oc.caller.Call(
		ctx,
		opCall.ChildCallID,
		opCallScope,
		opCall.ChildCallCallSpec,
		opCall.OpPath,
		&opCall.OpID,
		opCall.RootCallID,
	)
	if nil != err {
		return outboundScope, err
	}

	var opFile *model.OpSpec
	opFile, err = oc.opFileGetter.Get(
		ctx,
		opCall.OpPath,
	)
	if nil != err {
		return outboundScope, err
	}
	opOutputs, err = oc.outputsInterpreter.Interpret(
		opOutputs,
		opFile.Outputs,
		opCall.OpPath,
		filepath.Join(oc.callScratchDir, opCall.OpID),
	)

	// filter op outboundScope to bound call outboundScope
	for boundName, boundValue := range opCallSpec.Outputs {
		// return bound outboundScope
		if "" == boundValue {
			// implicit value
			boundValue = boundName
		}
		for opOutputName, opOutputValue := range opOutputs {
			if boundValue == opOutputName {
				outboundScope[boundName] = opOutputValue
			}
		}
	}

	return outboundScope, err
}
