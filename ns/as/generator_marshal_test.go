package as_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-act"
	"github.com/reiver/go-act/ns/as"
)

func TestGenerator_marshal(t *testing.T) {

	tests := []struct{
		Value as.Generator
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Value: as.Generator{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"name":"apple banana cherry"}`),
		},
		{
			Value: as.Generator{
				Type: opt.Something("bot"),
			},
			Expected: []byte(`{"type":"bot"}`),
		},
		{
			Value: as.Generator{
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"url":"http://example.com/basket-bot.html"}`),
		},



		{
			Value: as.Generator{
				Type: opt.Something("bot"),
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"type":"bot","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: as.Generator{
				Name: opt.Something("apple banana cherry"),
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"name":"apple banana cherry","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: as.Generator{
				Name: opt.Something("apple banana cherry"),
				Type: opt.Something("bot"),
			},
			Expected: []byte(`{"name":"apple banana cherry","type":"bot"}`),
		},



		{
			Value: as.Generator{
				Name: opt.Something("apple banana cherry"),
				Type: opt.Something("bot"),
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"name":"apple banana cherry","type":"bot","url":"http://example.com/basket-bot.html"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := act.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled %T is not what was expected", testNumber, test.Value)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}