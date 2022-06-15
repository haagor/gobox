package usecase

import (
    "log"
    "testing"
    "time"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    friend "github.com/haagor/gobox/katas/birthdayGreetings/identity/entity"
    mocks "github.com/haagor/gobox/katas/birthdayGreetings/identity/usecase/mocks"
)

func TestGetFriendsBornAt(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockDBAdapter := mocks.NewMockDBAdapter(mockCtrl)

    b, err := time.Parse("2006-01-02", "1993-10-24")
    if err != nil {
        log.Fatal(err)
    }

    mockDBAdapter.EXPECT().GetFriendsByBirthDate(b).Return(nil).Times(1)
    r := GetFriendsBornAt(mockDBAdapter, "1993-10-24")
    assert.Equal(t, 0, len(r))

    var fs []friend.Friend
    fs = append(fs, friend.Friend{FirstName: "Julien", LastName: "Paris", Email: "j.p@wanadoo.fr", Birth: b})
    mockDBAdapter.EXPECT().GetFriendsByBirthDate(b).Return(fs).Times(1)
    r = GetFriendsBornAt(mockDBAdapter, "1993-10-24")
    var e [][3]string
    e = append(e, [3]string{"j.p@wanadoo.fr", "Julien", "Paris"})
    assert.Equal(t, r, e)
}
