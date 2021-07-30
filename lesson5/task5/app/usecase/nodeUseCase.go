package usecase

import (
	"app/repository/httplib"
	"fmt"
)

func getURL(addr, path string) string {
	return fmt.Sprintf("http://%s%s", addr, path)
}

func (uc *DatabaseUseCase) Get(addr, path string, target interface{}) error {

	url := getURL(addr, path)

	if err := uc.rdb.Get(url, &target); err != nil {
		if err := httplib.GetJSON(url, &target); err != nil {
			return err
		}

		if err := uc.rdb.Set(url, target); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (uc *DatabaseUseCase) RebootModem(addr, path string) error {

	url := getURL(addr, path)

	if err := httplib.PostReq(url); err != nil {
		return err
	}

	return nil
}
