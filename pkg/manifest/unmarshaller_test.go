package manifest

import (
	"errors"
	"fmt"
	"github.com/golang-interfaces/iioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"gopkg.in/yaml.v2"
)

var _ = Context("_manifestUnmarshaller", func() {

	Context("Unmarshal", func() {

		It("should call validate w/ expected inputs", func() {
			/* arrange */
			providedPath := "/dummy/path"

			fakeValidator := new(fakeValidator)

			// err to cause immediate return
			fakeValidator.ValidateReturns([]error{errors.New("dummyError")})

			objectUnderTest := _Manifest{
				validator: fakeValidator,
			}

			/* act */
			objectUnderTest.Unmarshal(providedPath)

			/* assert */
			Expect(fakeValidator.ValidateArgsForCall(0)).To(Equal(providedPath))
		})
		Context("validator.Validate returns errors", func() {
			It("should return the expected error", func() {
				/* arrange */

				errs := []error{errors.New("dummyErr1"), errors.New("dummyErr2")}
				expectedErr := fmt.Errorf(`
-
  Error(s):
    - %v
    - %v
-`, errs[0], errs[1])

				fakeValidator := new(fakeValidator)

				// err to cause immediate return
				fakeValidator.ValidateReturns(errs)

				objectUnderTest := _Manifest{
					validator: fakeValidator,
				}

				/* act */
				_, actualError := objectUnderTest.Unmarshal("")

				/* assert */
				Expect(actualError).To(Equal(expectedErr))
			})
		})
		Context("validator.Validate doesn't return errors", func() {
			It("should call ioutil.ReadFile w/expected args", func() {
				/* arrange */
				providedPath := "dummyPath"

				fakeIOUtil := new(iioutil.Fake)
				// err to cause immediate return
				fakeIOUtil.ReadFileReturns(nil, errors.New("dummyError"))

				objectUnderTest := _Manifest{
					ioUtil:    fakeIOUtil,
					validator: new(fakeValidator),
				}

				/* act */
				objectUnderTest.Unmarshal(providedPath)

				/* assert */
				Expect(fakeIOUtil.ReadFileArgsForCall(0)).
					To(Equal(providedPath))

			})
			Context("ioutil.ReadFile returns an error", func() {

				It("should return expected error", func() {

					/* arrange */
					expectedError := errors.New("dummyError")

					fakeIOUtil := new(iioutil.Fake)
					fakeIOUtil.ReadFileReturns(nil, expectedError)

					objectUnderTest := _Manifest{
						ioUtil:    fakeIOUtil,
						validator: new(fakeValidator),
					}

					/* act */
					_, actualError := objectUnderTest.Unmarshal("/dummy/path")

					/* assert */
					Expect(actualError).To(Equal(expectedError))

				})

			})

			It("should return expected pkgManifest", func() {

				/* arrange */
				paramDefault := "dummyDefault"
				dummyParams := map[string]*model.Param{
					"dummyName": {
						String: &model.StringParam{
							Constraints: &model.StringConstraints{
								MinLength: 0,
								MaxLength: 1000,
								Pattern:   "dummyPattern",
								Format:    "dummyFormat",
								Enum:      []string{"dummyEnumItem1"},
							},
							Default:     &paramDefault,
							Description: "dummyDescription",
							IsSecret:    true,
						},
					},
				}

				expectedPkgManifest := &model.PkgManifest{
					Description: "dummyDescription",
					Inputs:      dummyParams,
					Name:        "dummyName",
					Outputs:     dummyParams,
					Run: &model.SCG{
						Op: &model.SCGOpCall{
							Pkg: &model.SCGOpCallPkg{
								Ref: "dummyPkgRef",
							},
						},
					},
					Version: "dummyVersion",
				}

				fakeIOUtil := new(iioutil.Fake)
				fakeIOUtil.ReadFileReturns(yaml.Marshal(expectedPkgManifest))

				objectUnderTest := _Manifest{
					ioUtil:    fakeIOUtil,
					validator: new(fakeValidator),
				}

				/* act */
				actualPkgManifest, _ := objectUnderTest.Unmarshal("")

				/* assert */
				Expect(actualPkgManifest).To(Equal(expectedPkgManifest))

			})
		})
	})
})
