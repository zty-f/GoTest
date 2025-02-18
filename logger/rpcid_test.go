package logger

import (
	"context"
	"testing"
)

func TestCurrentRpcId(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	ctx := context.Background()
	ctx2 := context.WithValue(ctx, "rpcid", "3.2")
	tests := []struct {
		name string
		args args
		want string
	}{
		{"rpcid 1", args{ctx}, "1.1"},
		{"rpcid 2", args{ctx2}, "3.2.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrentRpcId(tt.args.ctx); got != tt.want {
				t.Errorf("CurrentRpcId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextRpcId(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKey(MetadataKey), NewTraceNode())
	ret := ""
	ret = NextRpcId(ctx)
	if ret != "1.1.1" {
		t.Errorf("NextRpcId() = %v, want %v", ret, "1.1.1")
	}
	ret = NextRpcId(ctx)
	if ret != "1.1.2" {
		t.Errorf("NextRpcId() = %v, want %v", ret, "1.1.2")
	}
	ret = NextRpcId(ctx)
	if ret != "1.1.3" {
		t.Errorf("NextRpcId() = %v, want %v", ret, "1.1.3")
	}
}
