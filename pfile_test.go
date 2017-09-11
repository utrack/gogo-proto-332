package listingpb

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
)

func TestProtoMarshalling(t *testing.T) {
	so := assert.New(t)
	f := &TestMsg{
		// Foos: []Foo{
		// 	Foo{Bar: "item1"},
		// 	Foo{Bar: "item2"},
		// },
		// Bar: Foo{Bar: "item3"},
		FooMap: map[string]Foo{
			"m1": Foo{"mi1"},
		},
	}

	got, err := (&jsonpb.Marshaler{}).MarshalToString(f)
	so.Nil(err)
	logrus.Info(got)
}
