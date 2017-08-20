package number

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/pkg/errors"
	"strconv"
)

var _ = Context("Interpret", func() {
	Context("expression is ref", func() {
		Context("ref isn't in scope", func() {
			It("should return expected result", func() {
				/* arrange */
				providedRef := "dummyRef"
				providedExpression := fmt.Sprintf("$(%v)", providedRef)

				objectUnderTest := _interpreter{}

				/* act */
				actualNumber, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					providedExpression,
				)

				/* assert */
				Expect(actualNumber).To(Equal(float64(0)))
				Expect(actualErr).To(Equal(fmt.Errorf("Unable to interpret number; %v not in scope", providedRef)))
			})
		})
		Context("ref is in scope", func() {
			It("should call coercer.Coerce w/ expected args", func() {
				/* arrange */
				providedRef := "dummyRef"
				providedExpression := fmt.Sprintf("$(%v)", providedRef)

				providedScopeValue := &model.Value{}
				providedScope := map[string]*model.Value{
					providedRef: providedScopeValue,
				}

				fakeCoercer := new(fakeCoercer)
				// err to trigger immediate return
				fakeCoercer.CoerceReturns(float64(0), errors.New("dummyError"))

				objectUnderTest := _interpreter{
					coercer: fakeCoercer,
				}

				/* act */
				objectUnderTest.Interpret(providedScope, providedExpression)

				/* assert */
				Expect(fakeCoercer.CoerceArgsForCall(0)).To(Equal(providedScopeValue))
			})
			Context("coercer.Coerce errs", func() {
				It("should return expected result", func() {
					/* arrange */
					providedRef := "dummyRef"
					providedExpression := fmt.Sprintf("$(%v)", providedRef)

					fakeCoercer := new(fakeCoercer)

					coerceError := errors.New("dummyError")
					fakeCoercer.CoerceReturns(float64(1), coerceError)

					objectUnderTest := _interpreter{
						coercer: fakeCoercer,
					}

					/* act */
					_, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{
							providedRef: nil,
						},
						providedExpression,
					)

					/* assert */
					Expect(actualErr).To(Equal(fmt.Errorf("Unable to interpret number; error was: %v", coerceError.Error())))
				})
			})
			Context("coercer.Coerce doesn't err", func() {
				It("should return expected result", func() {
					/* arrange */
					providedRef := "dummyRef"
					providedExpression := fmt.Sprintf("$(%v)", providedRef)

					fakeCoercer := new(fakeCoercer)

					coercedNumber := float64(1)
					fakeCoercer.CoerceReturns(coercedNumber, nil)

					objectUnderTest := _interpreter{
						coercer: fakeCoercer,
					}

					/* act */
					actualNumber, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{
							providedRef: nil,
						},
						providedExpression,
					)

					/* assert */
					Expect(actualNumber).To(Equal(coercedNumber))
					Expect(actualErr).To(BeNil())
				})
			})
		})
	})
	Context("expression is refs", func() {
		Context("first ref not in scope", func() {
			It("should return expected result", func() {
				/* arrange */
				providedRef1 := "dummyRef1"
				providedExpression := fmt.Sprintf("$(%v)$(dummyRef2)", providedRef1)

				objectUnderTest := _interpreter{}

				/* act */
				actualNumber, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					providedExpression,
				)

				/* assert */
				Expect(actualNumber).To(Equal(float64(0)))
				Expect(actualErr).To(Equal(fmt.Errorf("Unable to interpret number; %v not in scope", providedRef1)))
			})
		})
		Context("second ref not in scope", func() {
			It("should return expected result", func() {
				/* arrange */
				providedRef1 := "dummyRef1"
				providedRef2 := "dummyRef2"
				providedExpression := fmt.Sprintf("$(%v)$(%v)", providedRef1, providedRef2)

				fakeCoercer := new(fakeCoercer)

				objectUnderTest := _interpreter{
					coercer: fakeCoercer,
				}

				/* act */
				actualNumber, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{
						providedRef1: nil,
					},
					providedExpression,
				)

				/* assert */
				Expect(actualNumber).To(Equal(float64(0)))
				Expect(actualErr).To(Equal(fmt.Errorf("Unable to interpret number; %v not in scope", providedRef2)))
			})
		})
		Context("refs in scope", func() {
			It("should call coercer.Coerce w/ expected args", func() {
				/* arrange */
				providedRef1 := "dummyRef1"
				providedRef2 := "dummyRef2"
				providedExpression := fmt.Sprintf("$(%v)$(%v)", providedRef1, providedRef2)

				providedScopeValue1 := &model.Value{}
				providedScopeValue2 := &model.Value{}
				providedScope := map[string]*model.Value{
					providedRef1: providedScopeValue1,
					providedRef2: providedScopeValue2,
				}

				fakeCoercer := new(fakeCoercer)

				objectUnderTest := _interpreter{
					coercer: fakeCoercer,
				}

				/* act */
				objectUnderTest.Interpret(providedScope, providedExpression)

				/* assert */
				Expect(fakeCoercer.CoerceArgsForCall(0)).To(Equal(providedScopeValue1))
				Expect(fakeCoercer.CoerceArgsForCall(1)).To(Equal(providedScopeValue2))
			})
			Context("first coercer.Coerce errs", func() {
				It("should return expected result", func() {
					/* arrange */
					providedRef1 := "dummyRef1"
					providedExpression := fmt.Sprintf("$(%v)$(dummyRef2)", providedRef1)

					fakeCoercer := new(fakeCoercer)

					coerceError := errors.New("dummyError")
					fakeCoercer.CoerceReturns(float64(1), coerceError)

					objectUnderTest := _interpreter{
						coercer: fakeCoercer,
					}

					/* act */
					_, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{
							providedRef1: nil,
						},
						providedExpression,
					)

					/* assert */
					Expect(actualErr).To(Equal(fmt.Errorf("Unable to interpret number; error was: %v", coerceError.Error())))
				})
			})
			Context("second coercer.Coerce errs", func() {
				It("should return expected result", func() {
					/* arrange */
					providedRef1 := "dummyRef1"
					providedRef2 := "dummyRef2"
					providedExpression := fmt.Sprintf("$(%v)$(%v)", providedRef1, providedRef2)

					fakeCoercer := new(fakeCoercer)

					coerceError := errors.New("dummyError")
					fakeCoercer.CoerceReturnsOnCall(1, float64(1), coerceError)

					objectUnderTest := _interpreter{
						coercer: fakeCoercer,
					}

					/* act */
					_, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{
							providedRef1: nil,
							providedRef2: nil,
						},
						providedExpression,
					)

					/* assert */
					Expect(actualErr).To(Equal(fmt.Errorf("Unable to interpret number; error was: %v", coerceError.Error())))
				})
			})
			Context("coercer.Coerce doesn't err", func() {
				It("should return expected result", func() {
					/* arrange */
					providedRef1 := "dummyRef1"
					providedRef2 := "dummyRef2"
					providedExpression := fmt.Sprintf("$(%v)$(%v)", providedRef1, providedRef2)

					fakeCoercer := new(fakeCoercer)

					coercedNumber1 := float64(1)
					fakeCoercer.CoerceReturnsOnCall(0, coercedNumber1, nil)
					coercedNumber2 := float64(2)
					fakeCoercer.CoerceReturnsOnCall(1, coercedNumber2, nil)

					expectedNumber, err := strconv.ParseFloat(
						strconv.FormatFloat(coercedNumber1, 'f', -1, 64)+
							strconv.FormatFloat(coercedNumber2, 'f', -1, 64),
						64,
					)
					if nil != err {
						panic(err)
					}

					objectUnderTest := _interpreter{
						coercer: fakeCoercer,
					}

					/* act */
					actualNumber, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{
							providedRef1: nil,
							providedRef2: nil,
						},
						providedExpression,
					)

					/* assert */
					Expect(actualNumber).To(Equal(expectedNumber))
					Expect(actualErr).To(BeNil())
				})
			})
		})
	})
	Context("expression doesn't contain ref", func() {
		It("should return input unmodified", func() {
			/* arrange */
			providedExpression := "123423.2"
			expectedNumber, err := strconv.ParseFloat(providedExpression, 64)
			if nil != err {
				panic(err)
			}

			objectUnderTest := _interpreter{}

			/* act */
			actualNumber, actualErr := objectUnderTest.Interpret(
				map[string]*model.Value{},
				providedExpression,
			)

			/* assert */
			Expect(actualNumber).To(Equal(expectedNumber))
			Expect(actualErr).To(BeNil())
		})
	})
	Context("expression contains refs", func() {
		Context("at beginning", func() {
			It("should return expected result", func() {
				/* arrange */
				providedRef := "dummyRef"
				providedSuffix := float64(2)
				providedExpression := fmt.Sprintf("$(%v)%v", providedRef, providedSuffix)

				fakeCoercer := new(fakeCoercer)

				coercedNumber := float64(1)
				fakeCoercer.CoerceReturns(coercedNumber, nil)

				expectedNumber, err := strconv.ParseFloat(
					strconv.FormatFloat(coercedNumber, 'f', -1, 64)+
						strconv.FormatFloat(providedSuffix, 'f', -1, 64),
					64,
				)
				if nil != err {
					panic(err)
				}

				objectUnderTest := _interpreter{
					coercer: fakeCoercer,
				}

				/* act */
				actualNumber, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{
						providedRef: nil,
					},
					providedExpression,
				)

				/* assert */
				Expect(actualNumber).To(Equal(expectedNumber))
				Expect(actualErr).To(BeNil())
			})
		})
		Context("in middle", func() {
			It("should return expected result", func() {
				/* arrange */
				providedRef := "dummyRef"
				providedPrefix := float64(1)
				providedSuffix := float64(3)
				providedExpression := fmt.Sprintf("%v$(%v)%v", providedPrefix, providedRef, providedSuffix)

				fakeCoercer := new(fakeCoercer)

				coercedNumber := float64(2)
				fakeCoercer.CoerceReturns(coercedNumber, nil)

				expectedNumber, err := strconv.ParseFloat(
					strconv.FormatFloat(providedPrefix, 'f', -1, 64)+
						strconv.FormatFloat(coercedNumber, 'f', -1, 64)+
						strconv.FormatFloat(providedSuffix, 'f', -1, 64),
					64,
				)
				if nil != err {
					panic(err)
				}

				objectUnderTest := _interpreter{
					coercer: fakeCoercer,
				}

				/* act */
				actualNumber, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{
						providedRef: nil,
					},
					providedExpression,
				)

				/* assert */
				Expect(actualNumber).To(Equal(expectedNumber))
				Expect(actualErr).To(BeNil())
			})
		})
		Context("at end", func() {
			It("should return expected result", func() {
				/* arrange */
				providedRef := "dummyRef"
				providedPrefix := float64(1)
				providedExpression := fmt.Sprintf("%v$(%v)", providedPrefix, providedRef)

				fakeCoercer := new(fakeCoercer)

				coercedNumber := float64(2)
				fakeCoercer.CoerceReturns(coercedNumber, nil)

				expectedNumber, err := strconv.ParseFloat(
					strconv.FormatFloat(providedPrefix, 'f', -1, 64)+
						strconv.FormatFloat(coercedNumber, 'f', -1, 64),
					64,
				)
				if nil != err {
					panic(err)
				}

				objectUnderTest := _interpreter{
					coercer: fakeCoercer,
				}

				/* act */
				actualNumber, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{
						providedRef: nil,
					},
					providedExpression,
				)

				/* assert */
				Expect(actualNumber).To(Equal(expectedNumber))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
