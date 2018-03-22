package entity

/*
//                            _ooOoo_
//                           o8888888o
//                           88" . "88
//                           (| -_- |)
//                            O\ = /O
//                        ____/`---'\____
//                      .   ' \\| |// `.
//                       / \\||| : |||// \
//                     / _||||| -:- |||||- \
//                       | | \\\ - /// | |
//                     | \_| ''\---/'' | |
//                      \ .-\__ `-` ___/-. /
//                   ___`. .' /--.--\ `. . __
//                ."" '< `.___\_<|>_/___.' >'"".
//               | | : `- \`.;`\ _ /`;.`/ - ` : | |
//                 \ \ `-. \_ __\ /__ _/ .-` / /
//      ======`-.____`-.___\_____/___.-`____.-*======
//                            `=---='
//
//         .............................................
//                  佛祖保佑             永无BUG
//          佛曰:
//                  写字楼里写字间，写字间里程序员；
//                  程序人员写程序，又拿程序换酒钱。
//                  酒醒只在网上坐，酒醉还来网下眠；
//                  酒醉酒醒日复日，网上网下年复年。
//                  但愿老死电脑间，不愿鞠躬老板前；
//                  奔驰宝马贵者趣，公交自行程序员。
//                  别人笑我忒疯癫，我笑自己命太贱；
//                  不见满街漂亮妹，哪个归得程序员？
//          博主曰:
//                  愿世间再无BUG，祝猿们早日出任CEO，
//                  赢取白富美，走上人生的巅峰！~~~
//      .............................................
*/
import (
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

func Empty(params interface{}) bool {
	//初始化变量
	var (
		flag         bool = true
		defaultValue reflect.Value
	)

	r := reflect.ValueOf(params)

	//获取对应类型默认值
	defaultValue = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), defaultValue.Interface()) {
		flag = false
	}
	return flag
}

//统计数组长度
func Count(info []interface{}) int {
	return len(info)
}

//设置等待时间
func Sleep(s time.Duration) {
	time.Sleep(s * time.Second)
}

//终止程序
func Die(result interface{}) {
	log.Println(result.(string))
	os.Exit(1)
}

//终止程序
func Exit(result interface{}) {
	log.Println(result.(string))
	os.Exit(1)

}

//获取网页数据
func FileGetContent(url string) string {
	var result string
	if Empty(url) == true {
		return result
	}
	h, err := http.Get(url) //获取url资源
	if err != nil { //如果获取失败返回空
		return result
	}
	if h.StatusCode != http.StatusOK { //如果获取状态码不为 200 直接返回空
		return result
	}
	defer h.Body.Close()

	buf := make([]byte, 1024)
	for {
		num, _ := h.Body.Read(buf) //读取http 内容
		if num == 0 {
			break
		}
		result += string(buf)
	}
	return result
}
