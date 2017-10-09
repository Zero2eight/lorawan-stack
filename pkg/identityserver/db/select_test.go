// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package db

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func testSelect(t *testing.T, q QueryContext) {
	a := assertions.New(t)

	{
		res := make([]*foo, 0)
		err := q.Select(&res, "SELECT * FROM foo")
		a.So(err, should.BeNil)
		a.So(res, should.HaveLength, len(data))
	}

	// into a slice of struct ptr
	{
		res := make([]*foo, 0)
		err := q.Select(&res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.HaveLength, 1)
	}

	// into a slice of struct
	{
		res := make([]foo, 0)
		err := q.Select(&res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.HaveLength, 1)
	}

	// into a slice of values
	{
		res := make([]string, 0)
		err := q.Select(&res, `SELECT bar FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.HaveLength, 1)
		a.So(res[0], should.Equal, data[1].Bar)
	}

	// cannot use struct directly
	{
		res := foo{}
		err := q.Select(res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.NotBeNil)
	}

	// cannot use slice directly
	{
		res := make([]string, 0)
		err := q.Select(res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.NotBeNil)
	}
}

func testSelectOne(t *testing.T, q QueryContext) {
	a := assertions.New(t)

	// into struct ptr
	{
		res := new(foo)
		err := q.SelectOne(res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.NotBeNil)
		a.So(res.ID, should.NotBeNil)
		a.So(res.Created, should.NotBeNil)
		a.So(res.Bar, should.Equal, data[1].Bar)
		a.So(res.Baz, should.Equal, data[1].Baz)
		a.So(res.Quu, should.Equal, data[1].Quu)
	}

	// into map
	{
		res := make(map[string]interface{})
		err := q.SelectOne(res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.NotBeNil)
		a.So(res["id"], should.NotBeNil)
		a.So(res["created"], should.NotBeNil)
		a.So(res["bar"], should.Equal, data[1].Bar)
		a.So(res["baz"], should.Equal, data[1].Baz)
		a.So(res["quu"], should.Equal, data[1].Quu)
	}

	// into ptr to map
	{
		res := make(map[string]interface{})
		err := q.SelectOne(&res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.NotBeNil)
		a.So(res["id"], should.NotBeNil)
		a.So(res["created"], should.NotBeNil)
		a.So(res["bar"], should.Equal, data[1].Bar)
		a.So(res["baz"], should.Equal, data[1].Baz)
		a.So(res["quu"], should.Equal, data[1].Quu)
	}

	// into value
	{
		res := ""
		err := q.SelectOne(&res, `SELECT bar FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(res, should.Equal, data[1].Bar)
	}

	// into ptr to value
	{
		res := new(string)
		err := q.SelectOne(&res, `SELECT bar FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.BeNil)
		a.So(*res, should.Equal, data[1].Bar)
	}

	// cannot use struct directly
	{
		res := foo{}
		err := q.SelectOne(res, `SELECT * FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.NotBeNil)
	}

	// cannot use value directly
	{
		res := ""
		err := q.SelectOne(res, `SELECT bar FROM foo WHERE bar = $1`, "bar-2")
		a.So(err, should.NotBeNil)
	}
}
