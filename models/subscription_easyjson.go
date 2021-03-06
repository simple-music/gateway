// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels(in *jlexer.Lexer, out *UsersPage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UsersPage, 0, 4)
			} else {
				*out = UsersPage{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 string
			v1 = string(in.String())
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels(out *jwriter.Writer, in UsersPage) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			out.String(string(v3))
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UsersPage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersPage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersPage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersPage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels(l, v)
}
func easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels1(in *jlexer.Lexer, out *UserFull) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "fullName":
			out.FullName = string(in.String())
		case "dateOfBirth":
			if in.IsNull() {
				in.Skip()
				out.DateOfBirth = nil
			} else {
				if out.DateOfBirth == nil {
					out.DateOfBirth = new(string)
				}
				*out.DateOfBirth = string(in.String())
			}
		case "musicalInstruments":
			if in.IsNull() {
				in.Skip()
				out.MusicalInstruments = nil
			} else {
				in.Delim('[')
				if out.MusicalInstruments == nil {
					if !in.IsDelim(']') {
						out.MusicalInstruments = make([]string, 0, 4)
					} else {
						out.MusicalInstruments = []string{}
					}
				} else {
					out.MusicalInstruments = (out.MusicalInstruments)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.MusicalInstruments = append(out.MusicalInstruments, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "numSubscribers":
			out.NumSubscribers = int64(in.Int64())
		case "numSubscriptions":
			out.NumSubscriptions = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels1(out *jwriter.Writer, in UserFull) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"fullName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FullName))
	}
	{
		const prefix string = ",\"dateOfBirth\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.DateOfBirth == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DateOfBirth))
		}
	}
	{
		const prefix string = ",\"musicalInstruments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MusicalInstruments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.MusicalInstruments {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"numSubscribers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NumSubscribers))
	}
	{
		const prefix string = ",\"numSubscriptions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NumSubscriptions))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserFull) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserFull) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserFull) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserFull) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels1(l, v)
}
func easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels2(in *jlexer.Lexer, out *SubscriptionsStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user":
			out.User = string(in.String())
		case "numSubscribers":
			out.NumSubscribers = int64(in.Int64())
		case "numSubscriptions":
			out.NumSubscriptions = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels2(out *jwriter.Writer, in SubscriptionsStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"numSubscribers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NumSubscribers))
	}
	{
		const prefix string = ",\"numSubscriptions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NumSubscriptions))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubscriptionsStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubscriptionsStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFfbd3743EncodeGithubComSimpleMusicGatewayModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubscriptionsStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubscriptionsStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFfbd3743DecodeGithubComSimpleMusicGatewayModels2(l, v)
}
