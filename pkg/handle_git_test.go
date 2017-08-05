package pkg

import (
	"errors"
	"fmt"
	"github.com/golang-interfaces/iioutil"
	"github.com/golang-interfaces/ios"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"os"
	"path/filepath"
)

var _ = Context("gitHandle", func() {

	Context("GetContent", func() {

		It("should call os.Open w/ expected args", func() {
			/* arrange */
			providedPkgPath := "dummyPkgPath"
			providedContentPath := "dummyContentPath"

			fakeOS := new(ios.Fake)

			objectUnderTest := gitHandle{
				os:   fakeOS,
				path: providedPkgPath,
			}

			/* act */
			objectUnderTest.GetContent(nil, providedContentPath)

			/* assert */
			Expect(fakeOS.OpenArgsForCall(0)).To(Equal(filepath.Join(providedPkgPath, providedContentPath)))
		})
	})

	wd, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	Context("ListContents", func() {
		It("should call ioutil.ReadDir w/ expected args", func() {
			/* arrange */
			providedPkgPath := "dummyPkgPath"

			fakeIOUtil := new(iioutil.Fake)

			// error to trigger immediate return
			fakeIOUtil.ReadDirReturns(nil, errors.New("dummyError"))

			objectUnderTest := gitHandle{
				ioUtil: fakeIOUtil,
				path:   providedPkgPath,
			}

			/* act */
			objectUnderTest.ListContents(nil)

			/* assert */
			Expect(fakeIOUtil.ReadDirArgsForCall(0)).To(Equal(providedPkgPath))
		})
		Context("ioutil.ReadDir errors", func() {
			It("should be returned", func() {

				/* arrange */
				expectedError := errors.New("dummyError")

				fakeIOUtil := new(iioutil.Fake)
				fakeIOUtil.ReadDirReturns(nil, expectedError)

				objectUnderTest := gitHandle{
					ioUtil: fakeIOUtil,
				}

				/* act */
				_, actualError := objectUnderTest.ListContents(nil)

				/* assert */
				Expect(actualError).To(Equal(expectedError))

			})
		})
		Context("ioutil.ReadDir doesn't error", func() {
			It("should return expected contentList", func() {
				/* arrange */
				rootPkgPath := fmt.Sprintf("%v/testdata/listContents", wd)

				dirStat, err := os.Stat(filepath.Join(rootPkgPath, "/dir1"))
				if nil != err {
					panic(err)
				}

				fileStat, err := os.Stat(filepath.Join(rootPkgPath, "/dir1/file2.txt"))
				if nil != err {
					panic(err)
				}

				expectedContents := []*model.PkgContent{
					{
						Mode: fileStat.Mode(),
						Path: "/dir1/file2.txt",
						Size: 34,
					},
					{
						Mode: dirStat.Mode(),
						Path: "/dir1",
						Size: dirStat.Size(),
					},
					{
						Mode: fileStat.Mode(),
						Path: "/file1.txt",
						Size: 18,
					},
				}

				objectUnderTest := gitHandle{
					ioUtil: iioutil.New(),
					path:   rootPkgPath,
				}

				/* act */
				actualContents, err := objectUnderTest.ListContents(nil)
				if nil != err {
					panic(err)
				}

				/* assert */
				Expect(actualContents).To(Equal(expectedContents))

			})
		})

	})
	Context("Ref", func() {
		It("should return expected ref", func() {
			/* arrange */
			pkgRef := "dummyPkgRef"

			objectUnderTest := gitHandle{
				pkgRef: pkgRef,
			}

			/* act */
			actualRef := objectUnderTest.Ref()

			/* assert */
			Expect(actualRef).To(Equal(pkgRef))
		})
	})
})
