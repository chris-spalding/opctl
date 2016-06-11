package opspec

//go:generate counterfeiter -o ./fakeCreateOpUseCase.go --fake-name fakeCreateOpUseCase ./ createOpUseCase

import (
  "github.com/opspec-io/sdk-golang/models"
  "path"
)

type createOpUseCase interface {
  Execute(
  req models.CreateOpReq,
  ) (err error)
}

func newCreateOpUseCase(
filesystem Filesystem,
yamlCodec yamlCodec,
) createOpUseCase {

  return &_createOpUseCase{
    filesystem:filesystem,
    yamlCodec:yamlCodec,
  }

}

type _createOpUseCase struct {
  filesystem Filesystem
  yamlCodec  yamlCodec
}

func (this _createOpUseCase) Execute(
req models.CreateOpReq,
) (err error) {

  err = this.filesystem.AddDir(
    req.Path,
  )
  if (nil != err) {
    return
  }

  var opFile = models.OpFile{
    Description:req.Description,
    Name:req.Name,
  }

  opFileBytes, err := this.yamlCodec.toYaml(&opFile)
  if (nil != err) {
    return
  }

  err = this.filesystem.SaveFile(
    path.Join(req.Path, "op.yml"),
    opFileBytes,
  )

  return

}
