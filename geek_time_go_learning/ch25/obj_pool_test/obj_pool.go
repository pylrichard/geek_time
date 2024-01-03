package obj_pool_test

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	//存放任意对象可以使用interface{}，使用时要进行类型断言，不推荐
	//建议不同对象使用单独对象池
	objBuf chan *ReusableObj
}

func NewObjPool(num int) *ObjPool {
	pool := ObjPool{}
	pool.objBuf = make(chan *ReusableObj, num)
	for i := 0; i < num; i++ {
		pool.objBuf<-&ReusableObj{}
	}

	return &pool
}

func (p *ObjPool) GetObj(t time.Duration) (*ReusableObj, error) {
	select {
	case obj := <-p.objBuf:
		return obj, nil
	//超时控制，避免对象池为空时阻塞
	case <-time.After(t):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.objBuf<-obj:
		return nil
	default:
		//当对象池满时，抛出异常
		return errors.New("overflow")
	}
}