package raftstore

import (
	"github.com/pingcap/kvproto/pkg/raft_cmdpb"
)

func respEnsureHeader(resp *raft_cmdpb.RaftCmdResponse) {
	header := resp.GetHeader()
	if header == nil {
		resp.Header = &raft_cmdpb.RaftResponseHeader{}
	}
}

func BindTerm(resp *raft_cmdpb.RaftCmdResponse, term uint64) {
	if term == 0 {
		return
	}
	respEnsureHeader(resp)
	resp.Header.CurrentTerm = term
}

func BindError(resp *raft_cmdpb.RaftCmdResponse, err error) {
	respEnsureHeader(resp)
	resp.Header.Error = RaftstoreErrToPbError(err)
}

func NewError(err error) *raft_cmdpb.RaftCmdResponse {
	resp := &raft_cmdpb.RaftCmdResponse {Header: &raft_cmdpb.RaftResponseHeader{}}
	BindError(resp, err)
	return resp
}

func ErrResp(err error, term uint64) *raft_cmdpb.RaftCmdResponse {
	resp := NewError(err)
	BindTerm(resp, term)
	return resp
}