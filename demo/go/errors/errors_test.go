package errors

import (
	"fmt"
	"testing"

	"github.com/donnol/do"
)

func testErrors() (err error) {
	err = t1()
	if err != nil {
		r, err := t2()
		if err != nil {
			return err
		}
		_ = r
	}
	return
}

func testErrors2() (err error) {
	err = t1()
	if err != nil {
		r, err := t3()
		if err != nil {
			return err
		}
		_ = r
	}
	// 没有用nil覆盖err值，导致t3成功了依然返回错误
	return
}

func testErrors3() (err error) {
	err = t1()
	if err != nil {
		r, err := t3()
		if err != nil {
			return err
		}
		_ = r
	}
	return nil
}

func t1() error {
	return fmt.Errorf("not exists")
}

func t2() (int, error) {
	return 1, fmt.Errorf("request failed")
}

func t3() (int, error) {
	return 2, nil
}

func TestErrors(t *testing.T) {
	{
		err := testErrors()
		do.Assert(t, err.Error(), "request failed")
	}
	{
		err := testErrors2()
		do.Assert(t, err.Error(), "not exists")
	}
	{
		err := testErrors3()
		do.Assert(t, err == nil, true)
	}
}
