package playwright

import (
	"reflect"
)

func newExpectWrapper(f interface{}, args []interface{}, cb func() error) (interface{}, error) {
	val := make(chan interface{}, 1)
	errChan := make(chan error, 1)
	go func() {
		reflectArgs := make([]reflect.Value, 0)
		for i := 0; i < len(args); i++ {
			reflectArgs = append(reflectArgs, reflect.ValueOf(args[i]))
		}
		result := reflect.ValueOf(f).Call(reflectArgs)
		var evVal interface{}
		if len(result) > 0 {
			evVal = result[0].Interface()
		}
		if len(result) > 1 {
			errVal := result[1].Interface()
			err, ok := errVal.(error)
			if ok && err != nil {
				errChan <- err
				return
			}
		}
		val <- evVal
	}()

	if err := cb(); err != nil {
		return nil, err
	}
	select {
	case err := <-errChan:
		return nil, err
	case val := <-val:
		return val, nil
	}
}
