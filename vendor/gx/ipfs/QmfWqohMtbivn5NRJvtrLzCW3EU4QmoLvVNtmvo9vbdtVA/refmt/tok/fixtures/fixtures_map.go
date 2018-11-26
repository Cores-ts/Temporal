package fixtures

import (
	. "gx/ipfs/QmfWqohMtbivn5NRJvtrLzCW3EU4QmoLvVNtmvo9vbdtVA/refmt/tok"
)

var sequences_Map = []Sequence{
	{"empty map",
		[]Token{
			{Type: TMapOpen, Length: 0},
			{Type: TMapClose},
		},
	},
	{"single row map",
		[]Token{
			{Type: TMapOpen, Length: 1},
			TokStr("key"),
			TokStr("value"),
			{Type: TMapClose},
		},
	},
	{"duo row map",
		[]Token{
			{Type: TMapOpen, Length: 2},
			TokStr("key"),
			TokStr("value"),
			TokStr("k2"),
			TokStr("v2"),
			{Type: TMapClose},
		},
	},
	{"duo row map alt2",
		// same as previous, but map entries in a different order -- useful to test that unmarshaller can accept that (or, reject non-canonical orders!)
		[]Token{
			{Type: TMapOpen, Length: 2},
			TokStr("k2"),
			TokStr("v2"),
			TokStr("key"),
			TokStr("value"),
			{Type: TMapClose},
		},
	},
	{"quad map default order",
		[]Token{
			{Type: TMapOpen, Length: 4},
			TokStr("1"), TokStr("1"),
			TokStr("b"), TokStr("2"),
			TokStr("bc"), TokStr("3"),
			TokStr("d"), TokStr("4"),
			{Type: TMapClose},
		},
	},
	{"quad map rfc7049 order",
		[]Token{
			{Type: TMapOpen, Length: 4},
			TokStr("1"), TokStr("1"),
			TokStr("b"), TokStr("2"),
			TokStr("d"), TokStr("3"),
			TokStr("bc"), TokStr("4"),
			{Type: TMapClose},
		},
	},
	{"10 map rfc7049 order",
		[]Token{
			{Type: TMapOpen, Length: 10},
			TokStr("1"), TokStr("1"),
			TokStr("2"), TokStr("2"),
			TokStr("b"), TokStr("3"),
			TokStr("d"), TokStr("4"),
			TokStr("11"), TokStr("5"),
			TokStr("bc"), TokStr("6"),
			TokStr("he"), TokStr("7"),
			TokStr("hell"), TokStr("8"),
			TokStr("hello"), TokStr("9"),
			TokStr("bccccc"), TokStr("10"),
			{Type: TMapClose},
		},
	},
	{"7 struct rfc7049 order",
		[]Token{
			{Type: TMapOpen, Length: 7},
			TokStr("g"), TokStr("1"),
			TokStr("ff"), TokStr("2"),
			TokStr("ccc"), TokStr("3"),
			TokStr("ddd"), TokStr("4"),
			TokStr("eee"), TokStr("5"),
			TokStr("bbbb"), TokStr("6"),
			TokStr("aaaaa"), TokStr("7"),
			{Type: TMapClose},
		},
	},
}
