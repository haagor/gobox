package usecase

import (
	"log"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	mocks "github.com/haagor/gobox/katas/birthdayGreetings/identity/usecase/mocks"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBAdapter := mocks.NewMockDBAdapter(mockCtrl)

	b, err := time.Parse("2006-01-02", "1993-10-24")
	if err != nil {
		log.Fatal(err)
	}

	mockDBAdapter.EXPECT().GetFriendsByBirthDate(b).Return(nil).Times(1)
	GetFriendsBornAt(mockDBAdapter, "1993-10-24")
}
