package test

import (
	"testing"

	"github.com/benny502/gorm-helper/builder"
	"github.com/benny502/gorm-helper/mocks"
	"github.com/benny502/gorm-helper/where"

	gomock "github.com/golang/mock/gomock"
)

func TestBuilder(t *testing.T) {
	where := where.NewWhere()
	where.Add("m_user.id in ?", []int{1, 2, 3})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mocks.NewMockDB(ctrl)

	db.EXPECT().Where(gomock.Eq("m_user.id in ?"), gomock.Eq([]int{1, 2, 3})).Return(db)
	builder.NewBuilder().WithWhere(where).Build(db)

	//panic("doom")
}
