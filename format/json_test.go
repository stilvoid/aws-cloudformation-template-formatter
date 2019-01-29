package format

import (
	"testing"
)

func TestJsonScalars(t *testing.T) {
	cases := []map[string]interface{}{
		{"foo": 1},
		{"foo": 1.0},
		{"foo": 1.234},
		{"foo": "hello"},
		{"foo": true},
		{"foo": false},
	}

	expecteds := []string{
		"{\n    \"foo\": 1\n}",
		"{\n    \"foo\": 1\n}",
		"{\n    \"foo\": 1.234\n}",
		"{\n    \"foo\": \"hello\"\n}",
		"{\n    \"foo\": true\n}",
		"{\n    \"foo\": false\n}",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Json(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

/*
func TestJsonList(t *testing.T) {
	cases := []map[string]interface{}{
		{"foo": []interface{}{}},
		{"foo": []interface{}{1}},
		{"foo": []interface{}{
			1,
			"foo",
			true,
		}},
		{"foo": []interface{}{
			[]interface{}{
				"foo",
				"bar",
			},
			"baz",
		}},
		{"foo": []interface{}{
			[]interface{}{
				[]interface{}{
					"foo",
					"bar",
				},
				"baz",
			},
			"quux",
		}},
		{"foo": []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz":  "quux",
				"mooz": "xyzzy",
			},
		}},
	}

	expecteds := []string{
		"{\n    \"foo\": []\n}",
		"{\n    \"foo\": [\n        1\n    ]\n}",
		"{\n    \"foo\": [\n    1\n  - foo\n  - true",
		"{\n    \"foo\": [\n    - foo\n    - bar\n  - baz",
		"{\n    \"foo\": [\n    - - foo\n      - bar\n    - baz\n  - quux",
		"{\n    \"foo\": [\n    foo: bar\n  - baz: quux\n    mooz: xyzzy",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Json(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestJsonMap(t *testing.T) {
	cases := []map[string]interface{}{
		{},
		{
			"foo": "bar",
		},
		{
			"foo": "bar",
			"baz": "quux",
		},
		{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
			"quux": "mooz",
		},
		{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "quux",
				},
				"mooz": "xyzzy",
			},
			"alpha": "beta",
		},
		{
			"foo": []interface{}{
				"bar",
				"baz",
			},
			"quux": []interface{}{
				"mooz",
			},
		},
	}

	expecteds := []string{
		"{}",
		"foo: bar",
		"baz: quux\n\nfoo: bar",
		"foo:\n  bar: baz\n\nquux: mooz",
		"alpha: beta\n\nfoo:\n  bar:\n    baz: quux\n\n  mooz: xyzzy",
		"foo:\n  - bar\n  - baz\n\nquux:\n  - mooz",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Json(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}
*/

func TestCfnJson(t *testing.T) {
	cases := []map[string]interface{}{
		{
			"Quux":       "mooz",
			"Parameters": "baz",
			"Foo":        "bar",
			"Resources":  "xyzzy",
		},
	}

	expecteds := []string{
		"{\n    \"Parameters\": \"baz\",\n    \"Resources\": \"xyzzy\",\n    \"Foo\": \"bar\",\n    \"Quux\": \"mooz\"\n}",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Json(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestJsonComments(t *testing.T) {
	data := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"quux": "mooz",
		},
		"xyzzy": []interface{}{
			"lorem",
		},
	}

	commentCases := []map[interface{}]interface{}{
		{},
		{"foo": "This is bar"},
		{"baz": "This is baz"},
		{"baz": map[string]interface{}{"quux": "This is quux"}},
	}

	expecteds := []string{
		"{\n    \"baz\": {\n        \"quux\": \"mooz\"\n    },\n    \"foo\": \"bar\",\n    \"xyzzy\": [\n        \"lorem\"\n    ]\n}",
		"{\n    \"baz\": {\n        \"quux\": \"mooz\"\n    },\n    \"foo\": \"bar\",  // This is bar\n    \"xyzzy\": [\n        \"lorem\"\n    ]\n}",
		"{\n    \"baz\": {  // This is baz\n        \"quux\": \"mooz\"\n    },\n    \"foo\": \"bar\",\n    \"xyzzy\": [\n        \"lorem\"\n    ]\n}",
		"{\n    \"baz\": {\n        \"quux\": \"mooz\"  // This is quux\n    },\n    \"foo\": \"bar\",\n    \"xyzzy\": [\n        \"lorem\"\n    ]\n}",
	}

	for i, comments := range commentCases {
		expected := expecteds[i]

		actual := JsonWithComments(data, comments)

		if actual != expected {
			t.Errorf("from %q != %q\n", actual, expected)
		}
	}
}
