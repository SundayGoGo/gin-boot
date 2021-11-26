/**
 * @Author: Tobin
 * @Description:
 * @File:  StructUtil
 * @Version: 1.0.0
 * @Date: 2021/3/19 10:44 上午
 */

package utils

import (
	"encoding/json"
)

func StructCopy(source interface{}, target interface{}) interface{} {

	aj, _ := json.Marshal(source)
	err := json.Unmarshal(aj, &target)

	if err != nil {
		return nil
	}
	return target
}
