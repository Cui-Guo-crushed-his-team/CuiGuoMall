// Code generated by Kitex v0.9.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	"errors"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Prepay": kitex.NewMethodInfo(
		prepayHandler,
		newPrepayArgs,
		newPrepayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Repay": kitex.NewMethodInfo(
		repayHandler,
		newRepayArgs,
		newRepayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Finish": kitex.NewMethodInfo(
		finishHandler,
		newFinishArgs,
		newFinishResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetByOutTradeNo": kitex.NewMethodInfo(
		getByOutTradeNoHandler,
		newGetByOutTradeNoArgs,
		newGetByOutTradeNoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	paymentServiceServiceInfo                = NewServiceInfo()
	paymentServiceServiceInfoForClient       = NewServiceInfoForClient()
	paymentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return paymentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PaymentService"
	handlerType := (*payment.PaymentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func prepayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.PrepayReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).Prepay(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PrepayArgs:
		success, err := handler.(payment.PaymentService).Prepay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PrepayResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPrepayArgs() interface{} {
	return &PrepayArgs{}
}

func newPrepayResult() interface{} {
	return &PrepayResult{}
}

type PrepayArgs struct {
	Req *payment.PrepayReq
}

func (p *PrepayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.PrepayReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PrepayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PrepayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PrepayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PrepayArgs) Unmarshal(in []byte) error {
	msg := new(payment.PrepayReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PrepayArgs_Req_DEFAULT *payment.PrepayReq

func (p *PrepayArgs) GetReq() *payment.PrepayReq {
	if !p.IsSetReq() {
		return PrepayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PrepayArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PrepayArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PrepayResult struct {
	Success *payment.PrepayResp
}

var PrepayResult_Success_DEFAULT *payment.PrepayResp

func (p *PrepayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.PrepayResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PrepayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PrepayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PrepayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PrepayResult) Unmarshal(in []byte) error {
	msg := new(payment.PrepayResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PrepayResult) GetSuccess() *payment.PrepayResp {
	if !p.IsSetSuccess() {
		return PrepayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PrepayResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.PrepayResp)
}

func (p *PrepayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrepayResult) GetResult() interface{} {
	return p.Success
}

func repayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.RepayReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).Repay(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RepayArgs:
		success, err := handler.(payment.PaymentService).Repay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RepayResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRepayArgs() interface{} {
	return &RepayArgs{}
}

func newRepayResult() interface{} {
	return &RepayResult{}
}

type RepayArgs struct {
	Req *payment.RepayReq
}

func (p *RepayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.RepayReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RepayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RepayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RepayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RepayArgs) Unmarshal(in []byte) error {
	msg := new(payment.RepayReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RepayArgs_Req_DEFAULT *payment.RepayReq

func (p *RepayArgs) GetReq() *payment.RepayReq {
	if !p.IsSetReq() {
		return RepayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RepayArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RepayArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RepayResult struct {
	Success *payment.RepayResp
}

var RepayResult_Success_DEFAULT *payment.RepayResp

func (p *RepayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.RepayResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RepayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RepayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RepayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RepayResult) Unmarshal(in []byte) error {
	msg := new(payment.RepayResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RepayResult) GetSuccess() *payment.RepayResp {
	if !p.IsSetSuccess() {
		return RepayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RepayResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.RepayResp)
}

func (p *RepayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RepayResult) GetResult() interface{} {
	return p.Success
}

func finishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.FinishReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).Finish(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *FinishArgs:
		success, err := handler.(payment.PaymentService).Finish(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FinishResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newFinishArgs() interface{} {
	return &FinishArgs{}
}

func newFinishResult() interface{} {
	return &FinishResult{}
}

type FinishArgs struct {
	Req *payment.FinishReq
}

func (p *FinishArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.FinishReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FinishArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FinishArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FinishArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *FinishArgs) Unmarshal(in []byte) error {
	msg := new(payment.FinishReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FinishArgs_Req_DEFAULT *payment.FinishReq

func (p *FinishArgs) GetReq() *payment.FinishReq {
	if !p.IsSetReq() {
		return FinishArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FinishArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FinishArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FinishResult struct {
	Success *payment.FinishResp
}

var FinishResult_Success_DEFAULT *payment.FinishResp

func (p *FinishResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.FinishResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FinishResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FinishResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FinishResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *FinishResult) Unmarshal(in []byte) error {
	msg := new(payment.FinishResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FinishResult) GetSuccess() *payment.FinishResp {
	if !p.IsSetSuccess() {
		return FinishResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FinishResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.FinishResp)
}

func (p *FinishResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FinishResult) GetResult() interface{} {
	return p.Success
}

func getByOutTradeNoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.GetByOutTradeNoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).GetByOutTradeNo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetByOutTradeNoArgs:
		success, err := handler.(payment.PaymentService).GetByOutTradeNo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetByOutTradeNoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetByOutTradeNoArgs() interface{} {
	return &GetByOutTradeNoArgs{}
}

func newGetByOutTradeNoResult() interface{} {
	return &GetByOutTradeNoResult{}
}

type GetByOutTradeNoArgs struct {
	Req *payment.GetByOutTradeNoReq
}

func (p *GetByOutTradeNoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.GetByOutTradeNoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetByOutTradeNoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetByOutTradeNoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetByOutTradeNoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetByOutTradeNoArgs) Unmarshal(in []byte) error {
	msg := new(payment.GetByOutTradeNoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetByOutTradeNoArgs_Req_DEFAULT *payment.GetByOutTradeNoReq

func (p *GetByOutTradeNoArgs) GetReq() *payment.GetByOutTradeNoReq {
	if !p.IsSetReq() {
		return GetByOutTradeNoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetByOutTradeNoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetByOutTradeNoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetByOutTradeNoResult struct {
	Success *payment.GetByOutTradeNoResp
}

var GetByOutTradeNoResult_Success_DEFAULT *payment.GetByOutTradeNoResp

func (p *GetByOutTradeNoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.GetByOutTradeNoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetByOutTradeNoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetByOutTradeNoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetByOutTradeNoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetByOutTradeNoResult) Unmarshal(in []byte) error {
	msg := new(payment.GetByOutTradeNoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetByOutTradeNoResult) GetSuccess() *payment.GetByOutTradeNoResp {
	if !p.IsSetSuccess() {
		return GetByOutTradeNoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetByOutTradeNoResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.GetByOutTradeNoResp)
}

func (p *GetByOutTradeNoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetByOutTradeNoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Prepay(ctx context.Context, Req *payment.PrepayReq) (r *payment.PrepayResp, err error) {
	var _args PrepayArgs
	_args.Req = Req
	var _result PrepayResult
	if err = p.c.Call(ctx, "Prepay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Repay(ctx context.Context, Req *payment.RepayReq) (r *payment.RepayResp, err error) {
	var _args RepayArgs
	_args.Req = Req
	var _result RepayResult
	if err = p.c.Call(ctx, "Repay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Finish(ctx context.Context, Req *payment.FinishReq) (r *payment.FinishResp, err error) {
	var _args FinishArgs
	_args.Req = Req
	var _result FinishResult
	if err = p.c.Call(ctx, "Finish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetByOutTradeNo(ctx context.Context, Req *payment.GetByOutTradeNoReq) (r *payment.GetByOutTradeNoResp, err error) {
	var _args GetByOutTradeNoArgs
	_args.Req = Req
	var _result GetByOutTradeNoResult
	if err = p.c.Call(ctx, "GetByOutTradeNo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
