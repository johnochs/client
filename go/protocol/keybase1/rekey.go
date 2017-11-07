// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/rekey.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type TLF struct {
	Id        TLFID    `codec:"id" json:"id"`
	Name      string   `codec:"name" json:"name"`
	Writers   []string `codec:"writers" json:"writers"`
	Readers   []string `codec:"readers" json:"readers"`
	IsPrivate bool     `codec:"isPrivate" json:"isPrivate"`
}

func (o TLF) DeepCopy() TLF {
	return TLF{
		Id:   o.Id.DeepCopy(),
		Name: o.Name,
		Writers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
		IsPrivate: o.IsPrivate,
	}
}

type ProblemTLF struct {
	Tlf           TLF   `codec:"tlf" json:"tlf"`
	Score         int   `codec:"score" json:"score"`
	Solution_kids []KID `codec:"solution_kids" json:"solution_kids"`
}

func (o ProblemTLF) DeepCopy() ProblemTLF {
	return ProblemTLF{
		Tlf:   o.Tlf.DeepCopy(),
		Score: o.Score,
		Solution_kids: (func(x []KID) []KID {
			if x == nil {
				return nil
			}
			var ret []KID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Solution_kids),
	}
}

// ProblemSet is for a particular (user,kid) that initiated a rekey problem.
// This problem consists of one or more problem TLFs, which are individually scored
// and have attendant solutions --- devices that if they came online can rekey and
// solve the ProblemTLF.
type ProblemSet struct {
	User User         `codec:"user" json:"user"`
	Kid  KID          `codec:"kid" json:"kid"`
	Tlfs []ProblemTLF `codec:"tlfs" json:"tlfs"`
}

func (o ProblemSet) DeepCopy() ProblemSet {
	return ProblemSet{
		User: o.User.DeepCopy(),
		Kid:  o.Kid.DeepCopy(),
		Tlfs: (func(x []ProblemTLF) []ProblemTLF {
			if x == nil {
				return nil
			}
			var ret []ProblemTLF
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Tlfs),
	}
}

type ProblemSetDevices struct {
	ProblemSet ProblemSet `codec:"problemSet" json:"problemSet"`
	Devices    []Device   `codec:"devices" json:"devices"`
}

func (o ProblemSetDevices) DeepCopy() ProblemSetDevices {
	return ProblemSetDevices{
		ProblemSet: o.ProblemSet.DeepCopy(),
		Devices: (func(x []Device) []Device {
			if x == nil {
				return nil
			}
			var ret []Device
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Devices),
	}
}

type Outcome int

const (
	Outcome_NONE    Outcome = 0
	Outcome_FIXED   Outcome = 1
	Outcome_IGNORED Outcome = 2
)

func (o Outcome) DeepCopy() Outcome { return o }

var OutcomeMap = map[string]Outcome{
	"NONE":    0,
	"FIXED":   1,
	"IGNORED": 2,
}

var OutcomeRevMap = map[Outcome]string{
	0: "NONE",
	1: "FIXED",
	2: "IGNORED",
}

func (e Outcome) String() string {
	if v, ok := OutcomeRevMap[e]; ok {
		return v
	}
	return ""
}

type RevokeWarning struct {
	EndangeredTLFs []TLF `codec:"endangeredTLFs" json:"endangeredTLFs"`
}

func (o RevokeWarning) DeepCopy() RevokeWarning {
	return RevokeWarning{
		EndangeredTLFs: (func(x []TLF) []TLF {
			if x == nil {
				return nil
			}
			var ret []TLF
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.EndangeredTLFs),
	}
}

type ShowPendingRekeyStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o ShowPendingRekeyStatusArg) DeepCopy() ShowPendingRekeyStatusArg {
	return ShowPendingRekeyStatusArg{
		SessionID: o.SessionID,
	}
}

type GetPendingRekeyStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o GetPendingRekeyStatusArg) DeepCopy() GetPendingRekeyStatusArg {
	return GetPendingRekeyStatusArg{
		SessionID: o.SessionID,
	}
}

type DebugShowRekeyStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o DebugShowRekeyStatusArg) DeepCopy() DebugShowRekeyStatusArg {
	return DebugShowRekeyStatusArg{
		SessionID: o.SessionID,
	}
}

type RekeyStatusFinishArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o RekeyStatusFinishArg) DeepCopy() RekeyStatusFinishArg {
	return RekeyStatusFinishArg{
		SessionID: o.SessionID,
	}
}

type RekeySyncArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Force     bool `codec:"force" json:"force"`
}

func (o RekeySyncArg) DeepCopy() RekeySyncArg {
	return RekeySyncArg{
		SessionID: o.SessionID,
		Force:     o.Force,
	}
}

type GetRevokeWarningArg struct {
	SessionID    int      `codec:"sessionID" json:"sessionID"`
	ActingDevice DeviceID `codec:"actingDevice" json:"actingDevice"`
	TargetDevice DeviceID `codec:"targetDevice" json:"targetDevice"`
}

func (o GetRevokeWarningArg) DeepCopy() GetRevokeWarningArg {
	return GetRevokeWarningArg{
		SessionID:    o.SessionID,
		ActingDevice: o.ActingDevice.DeepCopy(),
		TargetDevice: o.TargetDevice.DeepCopy(),
	}
}

type RekeyInterface interface {
	// ShowPendingRekeyStatus shows either pending gregor-initiated rekey harassments
	// or nothing if none were pending.
	ShowPendingRekeyStatus(context.Context, int) error
	// GetPendingRekeyStatus returns the pending ProblemSetDevices.
	GetPendingRekeyStatus(context.Context, int) (ProblemSetDevices, error)
	// DebugShowRekeyStatus is used by the CLI to kick off a "ShowRekeyStatus" window for
	// the current user.
	DebugShowRekeyStatus(context.Context, int) error
	// RekeyStatusFinish is called when work is completed on a given RekeyStatus window. The Outcome
	// can be Fixed or Ignored.
	RekeyStatusFinish(context.Context, int) (Outcome, error)
	// RekeySync flushes the current rekey loop and gets to a good stopping point
	// to assert state. Good for race-free testing, not very useful in production.
	// Force overrides a long-snooze.
	RekeySync(context.Context, RekeySyncArg) error
	// GetRevokeWarning computes the TLFs that will be endangered if actingDevice
	// revokes targetDevice.
	GetRevokeWarning(context.Context, GetRevokeWarningArg) (RevokeWarning, error)
}

func RekeyProtocol(i RekeyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.rekey",
		Methods: map[string]rpc.ServeHandlerDescription{
			"showPendingRekeyStatus": {
				MakeArg: func() interface{} {
					ret := make([]ShowPendingRekeyStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ShowPendingRekeyStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]ShowPendingRekeyStatusArg)(nil), args)
						return
					}
					err = i.ShowPendingRekeyStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getPendingRekeyStatus": {
				MakeArg: func() interface{} {
					ret := make([]GetPendingRekeyStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetPendingRekeyStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetPendingRekeyStatusArg)(nil), args)
						return
					}
					ret, err = i.GetPendingRekeyStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"debugShowRekeyStatus": {
				MakeArg: func() interface{} {
					ret := make([]DebugShowRekeyStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DebugShowRekeyStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]DebugShowRekeyStatusArg)(nil), args)
						return
					}
					err = i.DebugShowRekeyStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"rekeyStatusFinish": {
				MakeArg: func() interface{} {
					ret := make([]RekeyStatusFinishArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RekeyStatusFinishArg)
					if !ok {
						err = rpc.NewTypeError((*[]RekeyStatusFinishArg)(nil), args)
						return
					}
					ret, err = i.RekeyStatusFinish(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"rekeySync": {
				MakeArg: func() interface{} {
					ret := make([]RekeySyncArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RekeySyncArg)
					if !ok {
						err = rpc.NewTypeError((*[]RekeySyncArg)(nil), args)
						return
					}
					err = i.RekeySync(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getRevokeWarning": {
				MakeArg: func() interface{} {
					ret := make([]GetRevokeWarningArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetRevokeWarningArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetRevokeWarningArg)(nil), args)
						return
					}
					ret, err = i.GetRevokeWarning(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RekeyClient struct {
	Cli rpc.GenericClient
}

// ShowPendingRekeyStatus shows either pending gregor-initiated rekey harassments
// or nothing if none were pending.
func (c RekeyClient) ShowPendingRekeyStatus(ctx context.Context, sessionID int) (err error) {
	__arg := ShowPendingRekeyStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.showPendingRekeyStatus", []interface{}{__arg}, nil)
	return
}

// GetPendingRekeyStatus returns the pending ProblemSetDevices.
func (c RekeyClient) GetPendingRekeyStatus(ctx context.Context, sessionID int) (res ProblemSetDevices, err error) {
	__arg := GetPendingRekeyStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.getPendingRekeyStatus", []interface{}{__arg}, &res)
	return
}

// DebugShowRekeyStatus is used by the CLI to kick off a "ShowRekeyStatus" window for
// the current user.
func (c RekeyClient) DebugShowRekeyStatus(ctx context.Context, sessionID int) (err error) {
	__arg := DebugShowRekeyStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.debugShowRekeyStatus", []interface{}{__arg}, nil)
	return
}

// RekeyStatusFinish is called when work is completed on a given RekeyStatus window. The Outcome
// can be Fixed or Ignored.
func (c RekeyClient) RekeyStatusFinish(ctx context.Context, sessionID int) (res Outcome, err error) {
	__arg := RekeyStatusFinishArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.rekeyStatusFinish", []interface{}{__arg}, &res)
	return
}

// RekeySync flushes the current rekey loop and gets to a good stopping point
// to assert state. Good for race-free testing, not very useful in production.
// Force overrides a long-snooze.
func (c RekeyClient) RekeySync(ctx context.Context, __arg RekeySyncArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.rekey.rekeySync", []interface{}{__arg}, nil)
	return
}

// GetRevokeWarning computes the TLFs that will be endangered if actingDevice
// revokes targetDevice.
func (c RekeyClient) GetRevokeWarning(ctx context.Context, __arg GetRevokeWarningArg) (res RevokeWarning, err error) {
	err = c.Cli.Call(ctx, "keybase.1.rekey.getRevokeWarning", []interface{}{__arg}, &res)
	return
}
