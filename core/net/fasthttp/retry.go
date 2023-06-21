package fasthttp

import (
	"github.com/nyuncloud/ncloud-go/core/net/consts"
	"math/rand"
	"time"
)

type BodyFunc func() ([]byte, error)

func RetryDoWithErr(fn BodyFunc, retries int, sleep time.Duration) ([]byte, error) {
	if sleep == 0 {
		sleep = consts.DefaultSleep
	}
	body, err := fn()
	if err != nil {
		retries--
		if retries <= 0 {
			return nil, err
		}
		sleep += (time.Duration(rand.Int63n(int64(sleep)))) / 2
		time.Sleep(sleep)
		return RetryDoWithErr(fn, retries, 2*sleep)
	}

	return body, nil
}

func RetryDoHTTPWithErr(fn BodyFunc, retries int, sleep time.Duration) ([]byte, error) {
	var res []byte
	_, err := RetryDoWithErr(func() ([]byte, error) {
		var err error
		res, err = fn()
		return res, err
	}, retries, sleep)

	return res, err
}
