package matcher

import "testing"

func Test_match(t *testing.T) {
	type args struct {
		regexp string
		text   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "no ^ * . $", args: args{regexp: "a", text: "a"}, want: true},
		{name: "no ^ * . $", args: args{regexp: "a1", text: "ka1o"}, want: true},
		{name: "no ^ * . $", args: args{regexp: "a1", text: "aka1opa"}, want: true},

		{name: "^", args: args{regexp: "^a1", text: "a1opa"}, want: true},
		{name: "$", args: args{regexp: "lop$", text: "abclop"}, want: true},
		{name: "^$", args: args{regexp: "^a1lop$", text: "a1lop"}, want: true},

		{name: ".", args: args{regexp: "abc.123", text: "abcd123"}, want: true},
		{name: "*", args: args{regexp: "abc*b", text: "321abcccccb123"}, want: true},
		{name: ". *", args: args{regexp: "abc.*b", text: "321abc999b123"}, want: true},

		{name: "中文", args: args{regexp: "這是中文.嗨*", text: "這是中文無嗨嗨嗨"}, want: true},
		{name: "中文123", args: args{regexp: "這是中文.嗨*", text: "這是中文1嗨嗨嗨"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.regexp, tt.args.text); got != tt.want {
				t.Errorf("\nname = %v, args = %+v, matchHere() = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}

func Test_matchHere(t *testing.T) {
	type args struct {
		regexp string
		text   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "no * . $", args: args{regexp: "a", text: "a"}, want: true},
		{name: "no * . $", args: args{regexp: "a", text: "ab"}, want: true},
		{name: "no * . $", args: args{regexp: "ab", text: "ab"}, want: true},

		{name: "no * . $", args: args{regexp: "ab", text: "abc"}, want: true},
		{name: "no * . $", args: args{regexp: "ab", text: "ac"}, want: false},
		{name: "no * . $", args: args{regexp: "123", text: "123ac"}, want: true},

		{name: "with $", args: args{regexp: "123$", text: "123"}, want: true},
		{name: "with $", args: args{regexp: "123$", text: "1234"}, want: false},
		{name: "with $", args: args{regexp: "123$", text: "12"}, want: false},
		{name: "with $", args: args{regexp: "123$4", text: "123"}, want: false},
		{name: "with $", args: args{regexp: "123$4", text: "1234"}, want: false},
		{name: "with $", args: args{regexp: "123$4", text: "123$4"}, want: true},

		{name: "with .", args: args{regexp: ".abc", text: "1abcll"}, want: true},
		{name: "with . *", args: args{regexp: ".*abc", text: "1btaabcll"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchHere(tt.args.regexp, tt.args.text); got != tt.want {
				t.Errorf("\nname = %v, args = %+v, matchHere() = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}

func Test_matchStar(t *testing.T) {
	type args struct {
		c      rune
		regexp string
		text   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{c: 'a', regexp: "bc", text: "abc"}, want: true},
		{name: "2", args: args{c: 'a', regexp: "bc", text: "aabc"}, want: true},
		{name: "3", args: args{c: 'a', regexp: "bc", text: "aaabc"}, want: true},

		{name: "4", args: args{c: 'a', regexp: "bc", text: "bc"}, want: true},
		{name: "5", args: args{c: '1', regexp: "123", text: "123"}, want: true},
		{name: "6", args: args{c: '1', regexp: "123", text: "1123"}, want: true},

		{name: "7", args: args{c: 'a', regexp: "bc", text: "bcaa"}, want: true},
		{name: "8", args: args{c: '1', regexp: "123", text: "123aa"}, want: true},
		{name: "9", args: args{c: '1', regexp: "123", text: "1123aa"}, want: true},

		{name: "10", args: args{c: 'a', regexp: "bc", text: "abaa"}, want: false},
		{name: "11", args: args{c: '1', regexp: "123", text: "1k23aa"}, want: false},
		{name: "12", args: args{c: '1', regexp: "123", text: "1k123aa"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchStar(tt.args.c, tt.args.regexp, tt.args.text); got != tt.want {
				t.Errorf("\nname = %v, args = %+v, matchHere() = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
