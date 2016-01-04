package main_test

import (
	"app/typhenapi/type/submarine"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRoom(t *testing.T) {
	Convey("Room", t, func() {
		server := newTestServer()
		s := newClientSession()
		s2 := newClientSession()

		Convey("when a user join to the room", func() {
			Convey("should send a room message", func() {
				done := make(chan *submarine.Room)
				s.api.Battle.RoomHandler = func(m *submarine.Room) { done <- m }
				s.connect(server.URL + "/rooms/1?room_key=key_1")
				m := <-done
				So(m.Id, ShouldEqual, 1)
				So(m.Members, ShouldHaveLength, 1)
			})
		})

		Convey("when a new user join to the room", func() {
			Convey("should send a room message to other users", func() {
				done := make(chan *submarine.Room)
				s.api.Battle.RoomHandler = func(m *submarine.Room) { done <- m }
				s.connect(server.URL + "/rooms/1?room_key=key_1")
				s2.connect(server.URL + "/rooms/1?room_key=key_2")
				<-done
				m := <-done
				So(m.Id, ShouldEqual, 1)
				So(m.Members, ShouldHaveLength, 2)
				So(m.Members[1].Name, ShouldEqual, "I8")
			})
		})

		Reset(func() {
			server.Close()
		})
	})
}