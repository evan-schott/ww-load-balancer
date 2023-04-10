// Code generated by capnpc-go. DO NOT EDIT.

package arith

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	fmt "fmt"
)

type Arith capnp.Client

// Arith_TypeID is the unique identifier for the type Arith.
const Arith_TypeID = 0xf77c0d9e18b30d80

func (c Arith) Multiply(ctx context.Context, params func(Arith_multiply_Params) error) (Arith_multiply_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xf77c0d9e18b30d80,
			MethodID:      0,
			InterfaceName: "arith.capnp:Arith",
			MethodName:    "multiply",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 16, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Arith_multiply_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Arith_multiply_Results_Future{Future: ans.Future()}, release
}
func (c Arith) Divide(ctx context.Context, params func(Arith_divide_Params) error) (Arith_divide_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xf77c0d9e18b30d80,
			MethodID:      1,
			InterfaceName: "arith.capnp:Arith",
			MethodName:    "divide",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 16, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Arith_divide_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Arith_divide_Results_Future{Future: ans.Future()}, release
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Arith) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Arith) AddRef() Arith {
	return Arith(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Arith) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Arith) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Arith) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Arith) DecodeFromPtr(p capnp.Ptr) Arith {
	return Arith(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Arith) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Arith) IsSame(other Arith) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Arith) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Arith) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
} // A Arith_Server is a Arith with a local implementation.
type Arith_Server interface {
	Multiply(context.Context, Arith_multiply) error

	Divide(context.Context, Arith_divide) error
}

// Arith_NewServer creates a new Server from an implementation of Arith_Server.
func Arith_NewServer(s Arith_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Arith_Methods(nil, s), s, c)
}

// Arith_ServerToClient creates a new Client from an implementation of Arith_Server.
// The caller is responsible for calling Release on the returned Client.
func Arith_ServerToClient(s Arith_Server) Arith {
	return Arith(capnp.NewClient(Arith_NewServer(s)))
}

// Arith_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Arith_Methods(methods []server.Method, s Arith_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf77c0d9e18b30d80,
			MethodID:      0,
			InterfaceName: "arith.capnp:Arith",
			MethodName:    "multiply",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Multiply(ctx, Arith_multiply{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf77c0d9e18b30d80,
			MethodID:      1,
			InterfaceName: "arith.capnp:Arith",
			MethodName:    "divide",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Divide(ctx, Arith_divide{call})
		},
	})

	return methods
}

// Arith_multiply holds the state for a server call to Arith.multiply.
// See server.Call for documentation.
type Arith_multiply struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Arith_multiply) Args() Arith_multiply_Params {
	return Arith_multiply_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Arith_multiply) AllocResults() (Arith_multiply_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Arith_multiply_Results(r), err
}

// Arith_divide holds the state for a server call to Arith.divide.
// See server.Call for documentation.
type Arith_divide struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Arith_divide) Args() Arith_divide_Params {
	return Arith_divide_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Arith_divide) AllocResults() (Arith_divide_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_divide_Results(r), err
}

// Arith_List is a list of Arith.
type Arith_List = capnp.CapList[Arith]

// NewArith creates a new list of Arith.
func NewArith_List(s *capnp.Segment, sz int32) (Arith_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Arith](l), err
}

type Arith_multiply_Params capnp.Struct

// Arith_multiply_Params_TypeID is the unique identifier for the type Arith_multiply_Params.
const Arith_multiply_Params_TypeID = 0x937c04ee0557da9a

func NewArith_multiply_Params(s *capnp.Segment) (Arith_multiply_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_multiply_Params(st), err
}

func NewRootArith_multiply_Params(s *capnp.Segment) (Arith_multiply_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_multiply_Params(st), err
}

func ReadRootArith_multiply_Params(msg *capnp.Message) (Arith_multiply_Params, error) {
	root, err := msg.Root()
	return Arith_multiply_Params(root.Struct()), err
}

func (s Arith_multiply_Params) String() string {
	str, _ := text.Marshal(0x937c04ee0557da9a, capnp.Struct(s))
	return str
}

func (s Arith_multiply_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Arith_multiply_Params) DecodeFromPtr(p capnp.Ptr) Arith_multiply_Params {
	return Arith_multiply_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Arith_multiply_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Arith_multiply_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Arith_multiply_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Arith_multiply_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Arith_multiply_Params) A() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Arith_multiply_Params) SetA(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

func (s Arith_multiply_Params) B() int64 {
	return int64(capnp.Struct(s).Uint64(8))
}

func (s Arith_multiply_Params) SetB(v int64) {
	capnp.Struct(s).SetUint64(8, uint64(v))
}

// Arith_multiply_Params_List is a list of Arith_multiply_Params.
type Arith_multiply_Params_List = capnp.StructList[Arith_multiply_Params]

// NewArith_multiply_Params creates a new list of Arith_multiply_Params.
func NewArith_multiply_Params_List(s *capnp.Segment, sz int32) (Arith_multiply_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return capnp.StructList[Arith_multiply_Params](l), err
}

// Arith_multiply_Params_Future is a wrapper for a Arith_multiply_Params promised by a client call.
type Arith_multiply_Params_Future struct{ *capnp.Future }

func (f Arith_multiply_Params_Future) Struct() (Arith_multiply_Params, error) {
	p, err := f.Future.Ptr()
	return Arith_multiply_Params(p.Struct()), err
}

type Arith_multiply_Results capnp.Struct

// Arith_multiply_Results_TypeID is the unique identifier for the type Arith_multiply_Results.
const Arith_multiply_Results_TypeID = 0xde53484dbde6376e

func NewArith_multiply_Results(s *capnp.Segment) (Arith_multiply_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Arith_multiply_Results(st), err
}

func NewRootArith_multiply_Results(s *capnp.Segment) (Arith_multiply_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Arith_multiply_Results(st), err
}

func ReadRootArith_multiply_Results(msg *capnp.Message) (Arith_multiply_Results, error) {
	root, err := msg.Root()
	return Arith_multiply_Results(root.Struct()), err
}

func (s Arith_multiply_Results) String() string {
	str, _ := text.Marshal(0xde53484dbde6376e, capnp.Struct(s))
	return str
}

func (s Arith_multiply_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Arith_multiply_Results) DecodeFromPtr(p capnp.Ptr) Arith_multiply_Results {
	return Arith_multiply_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Arith_multiply_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Arith_multiply_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Arith_multiply_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Arith_multiply_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Arith_multiply_Results) Product() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Arith_multiply_Results) SetProduct(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

// Arith_multiply_Results_List is a list of Arith_multiply_Results.
type Arith_multiply_Results_List = capnp.StructList[Arith_multiply_Results]

// NewArith_multiply_Results creates a new list of Arith_multiply_Results.
func NewArith_multiply_Results_List(s *capnp.Segment, sz int32) (Arith_multiply_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[Arith_multiply_Results](l), err
}

// Arith_multiply_Results_Future is a wrapper for a Arith_multiply_Results promised by a client call.
type Arith_multiply_Results_Future struct{ *capnp.Future }

func (f Arith_multiply_Results_Future) Struct() (Arith_multiply_Results, error) {
	p, err := f.Future.Ptr()
	return Arith_multiply_Results(p.Struct()), err
}

type Arith_divide_Params capnp.Struct

// Arith_divide_Params_TypeID is the unique identifier for the type Arith_divide_Params.
const Arith_divide_Params_TypeID = 0x899d1df16a063b91

func NewArith_divide_Params(s *capnp.Segment) (Arith_divide_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_divide_Params(st), err
}

func NewRootArith_divide_Params(s *capnp.Segment) (Arith_divide_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_divide_Params(st), err
}

func ReadRootArith_divide_Params(msg *capnp.Message) (Arith_divide_Params, error) {
	root, err := msg.Root()
	return Arith_divide_Params(root.Struct()), err
}

func (s Arith_divide_Params) String() string {
	str, _ := text.Marshal(0x899d1df16a063b91, capnp.Struct(s))
	return str
}

func (s Arith_divide_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Arith_divide_Params) DecodeFromPtr(p capnp.Ptr) Arith_divide_Params {
	return Arith_divide_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Arith_divide_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Arith_divide_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Arith_divide_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Arith_divide_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Arith_divide_Params) Num() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Arith_divide_Params) SetNum(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

func (s Arith_divide_Params) Denom() int64 {
	return int64(capnp.Struct(s).Uint64(8))
}

func (s Arith_divide_Params) SetDenom(v int64) {
	capnp.Struct(s).SetUint64(8, uint64(v))
}

// Arith_divide_Params_List is a list of Arith_divide_Params.
type Arith_divide_Params_List = capnp.StructList[Arith_divide_Params]

// NewArith_divide_Params creates a new list of Arith_divide_Params.
func NewArith_divide_Params_List(s *capnp.Segment, sz int32) (Arith_divide_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return capnp.StructList[Arith_divide_Params](l), err
}

// Arith_divide_Params_Future is a wrapper for a Arith_divide_Params promised by a client call.
type Arith_divide_Params_Future struct{ *capnp.Future }

func (f Arith_divide_Params_Future) Struct() (Arith_divide_Params, error) {
	p, err := f.Future.Ptr()
	return Arith_divide_Params(p.Struct()), err
}

type Arith_divide_Results capnp.Struct

// Arith_divide_Results_TypeID is the unique identifier for the type Arith_divide_Results.
const Arith_divide_Results_TypeID = 0xc3b2f98a3be2a268

func NewArith_divide_Results(s *capnp.Segment) (Arith_divide_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_divide_Results(st), err
}

func NewRootArith_divide_Results(s *capnp.Segment) (Arith_divide_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Arith_divide_Results(st), err
}

func ReadRootArith_divide_Results(msg *capnp.Message) (Arith_divide_Results, error) {
	root, err := msg.Root()
	return Arith_divide_Results(root.Struct()), err
}

func (s Arith_divide_Results) String() string {
	str, _ := text.Marshal(0xc3b2f98a3be2a268, capnp.Struct(s))
	return str
}

func (s Arith_divide_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Arith_divide_Results) DecodeFromPtr(p capnp.Ptr) Arith_divide_Results {
	return Arith_divide_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Arith_divide_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Arith_divide_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Arith_divide_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Arith_divide_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Arith_divide_Results) Quo() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Arith_divide_Results) SetQuo(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

func (s Arith_divide_Results) Rem() int64 {
	return int64(capnp.Struct(s).Uint64(8))
}

func (s Arith_divide_Results) SetRem(v int64) {
	capnp.Struct(s).SetUint64(8, uint64(v))
}

// Arith_divide_Results_List is a list of Arith_divide_Results.
type Arith_divide_Results_List = capnp.StructList[Arith_divide_Results]

// NewArith_divide_Results creates a new list of Arith_divide_Results.
func NewArith_divide_Results_List(s *capnp.Segment, sz int32) (Arith_divide_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return capnp.StructList[Arith_divide_Results](l), err
}

// Arith_divide_Results_Future is a wrapper for a Arith_divide_Results promised by a client call.
type Arith_divide_Results_Future struct{ *capnp.Future }

func (f Arith_divide_Results_Future) Struct() (Arith_divide_Results, error) {
	p, err := f.Future.Ptr()
	return Arith_divide_Results(p.Struct()), err
}

const schema_f454c62f08bc504b = "x\xda\x8c\x92\xbfk\xd4`\x1c\xc6\x9f\xe7\xfb&\xcd!" +
	"\xde\xf0^:\x88p\x14\xca\x0d\x16\xfc\xd1\xaa \xb6C" +
	"\xafN\xa2\x08\xf7\xaa\xe0\x1c{\x81\x8b\xdc\xe5b~(" +
	"\x85\x82n\xa2\x9b\xba\x09u\x10\xfc\x07t\xd5I\x14\x1c" +
	"\x1c]\x04\x1d\x1c\x1c\x05A\x10\xa7H\xd2&\x06<\xd4" +
	"\xf1\x0b\x1f\x9eO\x9e\xbc\xcf\xb2\xc5\xbe\xb5\xd2^P\x10" +
	"\xd3\xb3\xe7\xf2\xfbks\xd7\xbeuw\xee\xc2t(\x80" +
	"\xe5\x00'2\x0aAw\x8b7\xc1\xfc\xd1\x87+\xf6W" +
	"k\xfba\x13x\xcfN\x01|,\x81\xd1\x93\xcfk\xf7" +
	"~>{\xd5\x04N\xcb\xbe\x02\xd8\x90\x02\x08O}y" +
	"y\xe1\xec\xa5O\x05\xc0=`G\x0e\x16\xc0SY\x07" +
	"\xf3\xdb\xed\xe7\x07\x1e\xb7\xb7\x7f@\xefW\xf9\xf9\xc1\x8b" +
	"\xd6\xb17\x97\xbf\x03t_\xcb\x03\xf7\x9d8\x80\xfbV" +
	"\xee\xb8]\xe5\xe0p\xee\xc5A::\xba\xe9\xa9(\x8c" +
	"V7\xcac\x18\xdc\x08\x86~o\xe0\xc5\xde$\x81i" +
	")\x0b\xb0\x08\xe8\xa5E\xc0\xf4\x14\xcd\xb2P\x93\xf3\x85" +
	"^\x1f9\x0e\x98C\x8a\xe6\xa4\xd0\x09\xb3\x09m\x08m" +
	"pa\xe8\x87\xd3\xfa\x9a\xe5\x99d\xe34\x88\xc6[\xbb" +
	"&&MSg\x96\xa9\xf3\xdbD\xafJ\xe6\xd5\xbf9" +
	"\xf6\xba\\\xf4\x93l\x9c&\xf8w\x9b\xc5F\x9b\xeb\xd9" +
	"\xb4\xcavb\xff\xff\xba\x94&\x95&\xc6\xaaE\xed3" +
	"\xa5\x97f^x+\x8a\xa7\xc3l3\xfd#\x8bU\x16" +
	"G\x03\xd2\xb4\x94\x0d\xd4ca\xf5\xe6z\xe5\x1cD/" +
	"9d\xbd4V\x83\xd1\xddU\x88\xd6N^}\x0b\x80" +
	">\xd7w\xff@\x9f\x03\xf2W\x00\x00\x00\xff\xff\xc1V" +
	"\x9c\xa3"

func init() {
	schemas.Register(schema_f454c62f08bc504b,
		0x899d1df16a063b91,
		0x937c04ee0557da9a,
		0xc3b2f98a3be2a268,
		0xde53484dbde6376e,
		0xf77c0d9e18b30d80)
}
