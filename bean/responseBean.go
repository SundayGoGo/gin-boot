/**
 * @Author: Tobin
 * @Description:
 * @File:  ResponseBean
 * @Version: 1.0.0
 * @Date: 2021/3/23 7:26 下午
 */

package bean

type ResponseBean struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type nihao interface {
	buhao()
}
type name struct {

}

func (receiver name)  buhao()  {

}
