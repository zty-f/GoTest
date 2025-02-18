package logger

import (
	"context"
	"sync"

	"github.com/spf13/cast"
)

const HeaderRpcIDKey = "rpcid"
const HeaderRpcIDSeq = "rpcid_seq"
const DefaultRpcId = "1"
const DefaultRpcSeq = 1
const MetadataKey = "xwx_log_trace_metadata_key"

type TraceNode struct {
	metadata map[string]string
	lock     *sync.RWMutex
}

func NewTraceNode() *TraceNode {
	t := new(TraceNode)
	t.metadata = make(map[string]string)
	t.lock = new(sync.RWMutex)
	return t
}

func (t *TraceNode) Get(key string) string {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.metadata[key]
}

func (t *TraceNode) Set(key, val string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.metadata[key] = val
	return
}

func (t *TraceNode) GetMetaData() map[string]string {
	res := make(map[string]string, 5)
	t.lock.RLock()
	defer t.lock.RUnlock()
	for mk, mv := range t.metadata {
		res[mk] = mv
	}
	return res
}

// GetTraceNodeFromContext 获取一个traceNode
func GetTraceNodeFromContext(ctx context.Context) *TraceNode {
	meta := ctx.Value(ContextKey(MetadataKey))
	if meta == nil {
		return NewTraceNode()
	}

	if val, ok := meta.(*TraceNode); ok {
		return val
	}

	return NewTraceNode()
}

// IncrementRpcId RpcId 1.1.1=>1.1.2
func IncrementRpcId(ctx context.Context) int {
	meta := GetTraceNodeFromContext(ctx)
	rpcSeq := 0
	rpcSeqCtx := meta.Get(HeaderRpcIDSeq)
	if rpcSeqCtx == "" {
		rpcSeq = DefaultRpcSeq
	} else {
		rpcSeq = cast.ToInt(rpcSeqCtx) + 1
	}
	meta.Set(HeaderRpcIDSeq, cast.ToString(rpcSeq))
	return rpcSeq
}

func ReceiveRpcID(ctx context.Context) string {
	rpcid := ctx.Value(HeaderRpcIDKey)
	if rpcid == nil {
		return DefaultRpcId
	}
	return rpcid.(string)
}

// CurrentRpcId 获取当前 rpcid
func CurrentRpcId(ctx context.Context) string {
	return ReceiveRpcID(ctx) + ".1"
}

// NextRpcId 获取下一个 rpcid
func NextRpcId(ctx context.Context) string {
	return CurrentRpcId(ctx) + "." + cast.ToString(IncrementRpcId(ctx))
}
