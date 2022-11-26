// rpc 远程过程调用 c/s场景 仅当程序运行在不同机器上时，这项技术才实用。实现了自动编码/解码传输的跨网络方法调用
package rpc_objects

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
