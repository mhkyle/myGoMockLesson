package myGoProjects

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
	"myProject.io/pkg/db"
)

func Test_getFromDBAndHandleIt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	mockedDB := db.NewMockDB(ctrl)
	// 打桩1：DB.Get("Tom") 方法返回值为 ""，错误为"not exist"
	mockedDB.EXPECT().Get(gomock.Eq("Tom")).Return("", errors.New("not exist"))
	// DB exists, 返回结果为 "Jerry"，错误为 nil
	mockedDB.EXPECT().Get(gomock.Eq("Jerry")).Return("Jerry", nil)

	// 测试用例1：
	// 输入："Tom"
	// 预期输出：""，错误："not exist"
	if v, err := getFromDBAndHandleIt(mockedDB, "Tom"); v != "" || err == nil {
		t.Fatal("expected \"\" and not exist error, but got", v, err)
	}

	// 测试用例2：
	// 输入："Jerry"
	// 预期输出："Jerry result"，错误：nil
	if v, err := getFromDBAndHandleIt(mockedDB, "Jerry"); v != "Jerry" || err != nil {
		t.Fatal("expected Jerry and nil error, but got", v, err)
	}

}
