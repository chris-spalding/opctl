package core

import (
	"fmt"
	"github.com/opctl/opctl/util/cliexiter"
	"github.com/opctl/opctl/util/cliparamsatisfier"
	"github.com/opspec-io/sdk-golang/model"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (this _core) Run(
	pkgRef string,
	opts *RunOpts,
) {

	// ensure node running
	nodes, err := this.nodeProvider.ListNodes()
	if nil != err {
		panic(err.Error())
	}
	if len(nodes) < 1 {
		this.nodeProvider.CreateNode()
	}

	cwd, err := this.os.Getwd()
	if nil != err {
		this.cliExiter.Exit(cliexiter.ExitReq{Message: err.Error(), Code: 1})
		return // support fake exiter
	}

	startTime := time.Now().UTC()

	pkgPath, ok := this.pkg.Resolve(cwd, pkgRef)
	if !ok {
		msg := fmt.Sprintf("Unable to resolve package '%v' from '%v'; pull required", pkgRef, cwd)
		this.cliExiter.Exit(cliexiter.ExitReq{Message: msg, Code: 1})
		return // support fake exiter
	}

	pkgManifest, err := this.pkg.Get(pkgPath)
	if nil != err {
		this.cliExiter.Exit(cliexiter.ExitReq{Message: err.Error(), Code: 1})
		return // support fake exiter
	}

	argsMap := this.cliParamSatisfier.Satisfy(
		cliparamsatisfier.NewInputSourcer(
			cliparamsatisfier.NewSliceInputSrc(opts.Args, "="),
			cliparamsatisfier.NewYMLFileInputSrc(opts.ArgFile, this.ioutil),
			cliparamsatisfier.NewEnvVarInputSrc(),
			cliparamsatisfier.NewParamDefaultInputSrc(pkgManifest.Inputs),
			cliparamsatisfier.NewCliPromptInputSrc(pkgManifest.Inputs),
		),
		pkgManifest.Inputs,
	)

	// init signal channel
	intSignalsReceived := 0
	signalChannel := make(chan os.Signal, 1)
	defer close(signalChannel)

	signal.Notify(
		signalChannel,
		syscall.SIGINT, //handle SIGINTs
	)

	// start op
	rootOpId, err := this.opspecNodeAPIClient.StartOp(
		model.StartOpReq{
			Args:   argsMap,
			PkgRef: pkgPath,
		},
	)
	if nil != err {
		this.cliExiter.Exit(cliexiter.ExitReq{Message: err.Error(), Code: 1})
		return // support fake exiter
	}

	// start event loop
	eventChannel, err := this.opspecNodeAPIClient.GetEventStream(
		&model.GetEventStreamReq{
			Filter: &model.EventFilter{
				RootOpIds: []string{rootOpId},
				Since:     &startTime,
			},
		},
	)
	if nil != err {
		this.cliExiter.Exit(cliexiter.ExitReq{Message: err.Error(), Code: 1})
		return // support fake exiter
	}

	for {
		select {

		case <-signalChannel:
			if intSignalsReceived == 0 {

				intSignalsReceived++
				fmt.Println(this.cliColorer.Error("Gracefully stopping... (signal Control-C again to force)"))

				this.opspecNodeAPIClient.KillOp(
					model.KillOpReq{
						OpId: rootOpId,
					},
				)
			} else {
				this.cliExiter.Exit(cliexiter.ExitReq{Message: "Terminated by Control-C", Code: 130})
				return // support fake exiter
			}

		case event, isEventChannelOpen := <-eventChannel:
			if !isEventChannelOpen {
				this.cliExiter.Exit(cliexiter.ExitReq{Message: "Event channel closed unexpectedly", Code: 1})
				return // support fake exiter
			}

			this.cliOutput.Event(&event)

			if nil != event.OpEnded {
				if event.OpEnded.OpId == rootOpId {
					switch event.OpEnded.Outcome {
					case model.OpOutcomeSucceeded:
						this.cliExiter.Exit(cliexiter.ExitReq{Code: 0})
					case model.OpOutcomeKilled:
						this.cliExiter.Exit(cliexiter.ExitReq{Code: 137})
					default:
						// treat model.OpOutcomeFailed & unexpected values as errors.
						this.cliExiter.Exit(cliexiter.ExitReq{Code: 1})
					}
					return // support fake exiter
				}
			}
		}

	}

}
