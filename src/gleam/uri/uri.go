package uri_P

import (
	"net/url"
	"strconv"

	gleam_P "example.com/todo/gleam"
	option_P "example.com/todo/gleam_stdlib/gleam/option"
)

func doParse(uriString string) gleam_P.Result_t[Uri_t, gleam_P.Nil_t] {
	uri, err := url.Parse(uriString)
	if err != nil {
		return gleam_P.Error_c[Uri_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}

	res := Uri_c{
		Scheme:   option_P.None_c[string]{},
		Userinfo: option_P.None_c[string]{},
		Host:     option_P.None_c[string]{},
		Port:     option_P.None_c[int64]{},
		Path:     "",
		Query:    option_P.None_c[string]{},
		Fragment: option_P.None_c[string]{},
	}

	if uri.Scheme != "" {
		res.Scheme = option_P.Some_c[string]{P_0: uri.Scheme}
	}
	if uri.User != nil {
		res.Userinfo = option_P.Some_c[string]{P_0: uri.User.String()}
	}
	host := uri.Hostname()
	if host != "" {
		res.Host = option_P.Some_c[string]{P_0: host}
	}
	port := uri.Port()
	if port != "" {
		port, _ := strconv.ParseInt(port, 10, 64)
		res.Port = option_P.Some_c[int64]{P_0: port}
	}
	res.Path = uri.Path
	if uri.RawQuery != "" {
		res.Query = option_P.Some_c[string]{P_0: uri.RawQuery}
	}
	if uri.Fragment != "" {
		res.Fragment = option_P.Some_c[string]{P_0: uri.Fragment}
	}

	return gleam_P.Ok_c[Uri_t, gleam_P.Nil_t]{P_0: res}
}

func popCodeunit(str string) gleam_P.Tuple2_t[int64, string] {
	if len(str) == 0 {
		return gleam_P.Tuple2_t[int64, string]{P_0: 0, P_1: ""}
	}
	return gleam_P.Tuple2_t[int64, string]{P_0: int64(str[0]), P_1: str[1:]}
}

func codeunitSlice(str string, from, length int64) string {
	if length <= 0 {
		return ""
	}
	strlen := int64(len(str))
	if from >= strlen {
		return ""
	}
	for from < 0 {
		from += strlen
	}
	to := from + length
	if to > strlen {
		to = strlen
	}
	return str[from:to]
}

func ParseQuery(query string) gleam_P.Result_t[gleam_P.List_t[gleam_P.Tuple2_t[string, string]], gleam_P.Nil_t] {
	values, err := url.ParseQuery(query)
	if err != nil {
		return gleam_P.Error_c[gleam_P.List_t[gleam_P.Tuple2_t[string, string]], gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	entries := []gleam_P.Tuple2_t[string, string]{}
	for k, vs := range values {
		for _, v := range vs {
			entries = append(entries, gleam_P.Tuple2_t[string, string]{P_0: k, P_1: v})
		}
	}
	return gleam_P.Ok_c[gleam_P.List_t[gleam_P.Tuple2_t[string, string]], gleam_P.Nil_t]{P_0: gleam_P.ToList(entries...)}
}

func PercentEncode(s string) string {
	return url.QueryEscape(s)
}

func PercentDecode(s string) gleam_P.Result_t[string, gleam_P.Nil_t] {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return gleam_P.Error_c[string, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[string, gleam_P.Nil_t]{P_0: decoded}
}
