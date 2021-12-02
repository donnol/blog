# etcd

## raft

### message type

```go
// For description of different message types, see:
// https://pkg.go.dev/go.etcd.io/etcd/raft/v3#hdr-MessageType
type MessageType int32

const (
    // 选举时使用；
    // 如果节点是一个follower或candidate，它在选举超时前没有收到任何心跳，它就回传递MsgHup消息给它自己的Step方法，然后成为（或保持）一个candidate从而开启一个新的选举
	MsgHup            MessageType = 0
    // 一个内部类型，它向leader发送一个类型为“MsgHeartbeat”的心跳信号
    // 如果节点是一个leader，raft里的tick函数将会是“tickHeartbeat”，触发leader周期性地发送“MsgHeartbeat”消息给它的followers
	MsgBeat           MessageType = 1
    // 提议往它的日志条目里追加数据；
    // 这是一个特别的类型，由follower反推提议给leader（正常是leader提议，follower执行）；
    // 发给leader的话，leader调用“appendEntry”方法追加条目到它的日志里，然后调用“bcastAppend”方法发送这些条目给它的远端节点；
    // 发给candidate的话，它们直接丢弃该消息
    // 发给follower的话，follower会将消息存储到它们的信箱里。会把发送者的id一起存储，然后转发给leader。
	MsgProp           MessageType = 2
    // 包含了要复制的日志条目
    // leader调用“bcastAppend”（里面调用“sendAppend”），发送“一会要被复制的日志”消息；
    // 当candidate收到消息后，在它的Step方法里，它马上回退为follower，因为这条消息表明已经存在一个有效leader了。
    // candidate和follower均会返回一条“MsgAppResp”类型消息以作响应。
	MsgApp            MessageType = 3
    // 调用“handlerAppendEntries”方法
	MsgAppResp        MessageType = 4
    // 请求集群中的节点给自己投票；
    // 当节点是follower或candidate，并且它们的Step方法收到了“MsgHup”消息，节点调用“campaign”方法去提议自己成为一个leader。一旦“campaign”方法被调用，节点成为candidate，并发送“MsgVote”给集群中的远端节点请求投票。
    // 当leader或candidate的Step方法收到该消息，并且消息的Term比它们的Term小，“MsgVote”将被拒绝。
    // 当leader或candidate收到的消息的Term要更大时，它会回退为follower。
    // 当follower收到该消息，仅当发送者的最后的term比“MsgVote”的term要大，或发送者的最后term等于“MsgVote”的term（但发送者的最后提交index大于等于follower的），
	MsgVote           MessageType = 5
    // 投票响应；
    // 当candidate收到后，它会统计选票，如果大于majority（quorum），它成为leader并调用“bcastAppend”。如果candidate收到大量的否决票，它将回退到follower
	MsgVoteResp       MessageType = 6
    // 请求安装一个快照消息；
    // 当一个节点刚成为leader，或者leader收到了“MsgProp”消息，它调用“bcastAppend”方法（里面再调用“sendAppend”）方法到每个follower。在“sendAppend”方法里，如果一个leader获取term或条目失败了，leader通过"MsgSnap"消息请求快照。
	MsgSnap           MessageType = 7
    // leader发送心跳；
    // 当candidate收到“MsgHeartbeat”，并且消息的term比candidate的大，candidate回退到follower并且更新它的提交index为这次心跳里的值。然后candidate发送消息到它的信箱。
    // 当消息发送到follower的Step方法，并且消息的term比follower的大，follower更新它的leader id
	MsgHeartbeat      MessageType = 8
    // 心跳响应；
    // leader收到后就知道有哪些follower响应了。
    // 只有当leader的最后提交index比follower的Match index大时，leader执行“sendAppend”方法
	MsgHeartbeatResp  MessageType = 9
    // 表明请求没有被交付；
    // 当“MsgUnreachable”被传送到leader的Step方法，leader发现follower无法到达，很有可能“MsgApp”都丢失了。当follower的进度状态为复制时，leader设置它回probe（哨兵）
	MsgUnreachable    MessageType = 10
    // 表明快照安装消息的结果
    // 当一个follower拒绝了“MsgSnap”，这显示快照请求失败了--因为网络原因；**leader认为follower成为哨兵了**?(Then leader considers follower's progress as probe.)；
    // 当“MsgSnap”没有被拒绝，它表明快照成功了，leader设置follower的进度为哨兵，并恢复它的日志复制
	MsgSnapStatus     MessageType = 11
	MsgCheckQuorum    MessageType = 12
	MsgTransferLeader MessageType = 13
	MsgTimeoutNow     MessageType = 14
	MsgReadIndex      MessageType = 15
	MsgReadIndexResp  MessageType = 16
    // "MsgPreVote"和“MsgPreVoteResp”用在可选的两阶段选举协议上；
    // 当Config.PreVote为true，将会进行一次预选举，除非预选举表明竞争节点会赢，否则没有节点会增加它们的term值。
    // 这最小化了**一个发生了分区的节点重新加入到集群时**会带来的中断/干扰
	MsgPreVote        MessageType = 17
	MsgPreVoteResp    MessageType = 18
)
```

## 实现

## 使用
