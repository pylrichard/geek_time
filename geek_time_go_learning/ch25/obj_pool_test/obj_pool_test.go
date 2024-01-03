package obj_pool_test

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//对象数量超出对象池大小，报溢出异常
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	//	t.Error(err)
	//}
	for i := 0; i < 12; i++ {
		if v, err := pool.GetObj(1 * time.Second); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			//如果不回收对象到对象池，获取不到对象会报超时异常
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}