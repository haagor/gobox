package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mocks "github.com/haagor/gobox/katas/birthdayGreetings/identity/usecase/mocks"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBAdapter := mocks.NewMockDBAdapter(mockCtrl)
	GetFriendsBornAt("1993-10-24")

}
