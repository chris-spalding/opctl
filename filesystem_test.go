package opspec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "os"
  "path"
  "io/ioutil"
  "github.com/nu7hatch/gouuid"
)

var _ = Describe("_filesystem", func() {

  Context("newFilesystem()", func() {
    It("should return an instance convertable to Filesystem", func() {

      /* arrange/act */
      objectUnderTest := newFilesystem()

      /* assert */
      _, ok := objectUnderTest.(Filesystem)
      if !ok {
        Fail("result not assignable to Filesystem")
      }

    })
  })

  Context("AddDir", func() {

    Context("when passed valid path", func() {

      It("should create dir", func() {

        /* arrange */
        uuid, err := uuid.NewV4()
        if (nil != err) {
          panic(err)
        }
        providedPath := path.Join(os.TempDir(), uuid.String())

        objectUnderTest := newFilesystem()

        /* act */
        objectUnderTest.AddDir(providedPath)

        /* assert */
        Expect(providedPath).To(BeADirectory())

      })

      It("should return nil err", func() {

        /* arrange */
        uuid, err := uuid.NewV4()
        if (nil != err) {
          panic(err)
        }
        providedPath := path.Join(os.TempDir(), uuid.String())

        objectUnderTest := newFilesystem()

        /* act */
        err = objectUnderTest.AddDir(providedPath)

        /* assert */
        Expect(err).To(BeNil())

      })

    })

  })

  Context("GetBytesOfFile", func() {

    Context("when passed path of existing file", func() {

      It("should return expected bytes", func() {

        /* arrange */
        uuid, err := uuid.NewV4()
        if (nil != err) {
          panic(err)
        }
        providedPath := path.Join(os.TempDir(), uuid.String())

        expectedBytes := []byte("dummyBytes")
        err = ioutil.WriteFile(providedPath, expectedBytes, 0644)
        if (nil != err) {
          panic(err)
        }

        objectUnderTest := newFilesystem()

        /* act */
        actualBytes, _ := objectUnderTest.GetBytesOfFile(providedPath)

        /* assert */
        Expect(actualBytes).To(Equal(expectedBytes))

      })

      It("should return expected bytes", func() {

        /* arrange */
        uuid, err := uuid.NewV4()
        if (nil != err) {
          panic(err)
        }
        providedPath := path.Join(os.TempDir(), uuid.String())

        err = ioutil.WriteFile(providedPath, []byte("dummyBytes"), 0644)
        if (nil != err) {
          panic(err)
        }

        objectUnderTest := newFilesystem()

        /* act */
        _, err = objectUnderTest.GetBytesOfFile(providedPath)

        /* assert */
        Expect(err).To(BeNil())

      })

    })

  })

  Context("SaveFile", func() {

    Context("when passed valid path of existing file", func() {

      It("should return expected bytes", func() {

        /* arrange */
        uuid, err := uuid.NewV4()
        if (nil != err) {
          panic(err)
        }
        providedPath := path.Join(os.TempDir(), uuid.String())

        expectedBytes := []byte("dummyBytes")
        err = ioutil.WriteFile(providedPath, expectedBytes, 0644)
        if (nil != err) {
          panic(err)
        }

        objectUnderTest := newFilesystem()

        /* act */
        actualBytes, _ := objectUnderTest.GetBytesOfFile(providedPath)

        /* assert */
        Expect(actualBytes).To(Equal(expectedBytes))

      })

      It("should return expected bytes", func() {

        /* arrange */
        uuid, err := uuid.NewV4()
        if (nil != err) {
          panic(err)
        }
        providedPath := path.Join(os.TempDir(), uuid.String())

        err = ioutil.WriteFile(providedPath, []byte("dummyBytes"), 0644)
        if (nil != err) {
          panic(err)
        }

        objectUnderTest := newFilesystem()

        /* act */
        _, err = objectUnderTest.GetBytesOfFile(providedPath)

        /* assert */
        Expect(err).To(BeNil())

      })

    })

  })

})
