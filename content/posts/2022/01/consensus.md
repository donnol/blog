---
author: "jdlau"
date: 2022-01-17
linktitle: consensus
menu:
next:
prev:
title: consensus
weight: 10
categories: ['consensus']
tags: ['deal']
---

在中文里，共表示共同（至少两个人？一个人行不行？），识表示认识，组合一起成为共识，共同的认识，引申出共同的想法、共同的行为。

在英语里，con是一个词根--表示"共同"，sensus表示感觉，加在一起组成**consensus**。

人类社会的发展催生了交易，交易的前提是双方达成共识，比如油换盐，比如钱换粮。如果你不承认我的油，不愿意与我交易，那就没办法了。

人与人之间的共识是非常难以达成的，不像歌里唱的：我说一，你说一。很多时候，我说一，他也承诺他会说一，但他没说--可能因为一些事忘了说，可能因为他突然不想说了，也有可能他被胁迫了不能说。反正就是不一而足的情况导致了意见/行为不一。

在日常生活中，特别是集市上，往往都是一手交钱、一手交货，交易完成就完成了，如果后面出现了问题--比如货不对版、钱有真伪，那就是另外的问题了。

那如果我们分别在不同的地方，没法面对面交易呢；又或者交易的东西不方便马上拿到面前来交易呢；又或者交易之后发现货不对版不想要了呢？

这时候，为了解决这些问题，某种机构应运而生。结合现在网购流行的社会，大家不难发现有哪些这类的机构。

目前的社会除了网购流行之外，是不是机器也很流行呢。那机器又是什么呢？机器能做什么，从而在这个社会如此流行呢？机器又能不能充当某类机构来完成某些事呢？

## 共识要素

某件事，事的主体，事的具体。比如购物，买卖双方、以钱易物。

## 机器共识

### 拜占庭将军问题

[wiki](https://zh.wikipedia.org/wiki/%E6%8B%9C%E5%8D%A0%E5%BA%AD%E5%B0%86%E5%86%9B%E9%97%AE%E9%A2%98)

> 拜占庭将军问题（Byzantine Generals Problem），是由莱斯利·兰波特在其同名论文中提出的**分布式对等**网络**通信容错**问题。
>
> 在分布式计算中，**不同的计算机通过通讯交换信息达成共识**而按照同一套协作策略行动。但有时候，系统中的成员计算机可能出错而发送错误的信息，用于传递信息的通讯网络也可能导致信息损坏，使得网络中不同的成员关于全体协作的策略得出不同结论，从而**破坏系统一致性**。拜占庭将军问题被认为是容错性问题中最难的问题类型之一。 

关键词：分布式对等、通信容错、不同计算机通过通讯交换信息从而达成共识、共识达成失败会导致系统一致性被破坏。

问题描述：

> 一组拜占庭将军分别各率领一支军队共同围困一座城市。
>
> 为了简化问题，将各支军队的行动策略限定为进攻或撤离两种。
>
> 因为部分军队进攻部分军队撤离可能会造成灾难性后果，因此各位将军必须通过投票来达成一致策略，即所有军队一起进攻或所有军队一起撤离。
>
> 因为各位将军分处城市不同方向，他们只能通过信使互相联系。
>
> 在投票过程中每位将军都将自己投票给进攻还是撤退的信息通过信使分别通知其他所有将军，这样一来**每位将军根据自己的投票和其他所有将军送来的信息**就可以知道共同的投票结果而决定行动策略。

面临问题：

> 系统的问题在于，可能将军中**出现叛徒**，他们不仅可能向较为糟糕的策略投票，还可能**选择性地发送**投票信息。-- 出现叛徒，半真半假，选择性投票，（一人投两票） -- 控制投票时间，只要不在其他人都投完之后再投，他就没法知道别人投的什么票；一人投一票，投票之后不能再投；
>
>> 假设有9位将军投票，其中1名叛徒。8名忠诚的将军中出现了4人投进攻，4人投撤离的情况。这时候叛徒可能故意给4名投进攻的将领送信表示投票进攻，而给4名投撤离的将领送信表示投撤离。这样一来在4名投进攻的将领看来，投票结果是5人投进攻，从而发起进攻；而在4名投撤离的将军看来则是5人投撤离。这样各支军队的一致协同就遭到了破坏。
>
> 由于将军之间需要通过信使通讯，叛变将军可能通过**伪造信件**来以其他将军的身份发送假投票。而即使在保证所有将军忠诚的情况下，也不能排除信使被敌人截杀，甚至被敌人间谍替换等情况。因此很难通过保证**人员可靠性及通讯可靠性**来解决问题。
>
>> 人可能是假的，信可能是假，空气都可能是假的；
>
> 假使那些忠诚（或是没有出错）的将军仍然能通过多数决定来决定他们的战略，便称达到了拜占庭容错。在此，票都会有一个默认值，若消息（票）没有被收到，则使用此默认值来投票。 

应用：

> 在点对点式数字货币系统比特币里，比特币网络的运作是平行的（parallel）。各节点与终端都运算著区块链来达成工作量证明（PoW）。**工作量证明的链接是解决比特币系统中拜占庭问题的关键**，避免有问题的节点（即前文提到的“反叛的将军”）破坏数字货币系统里交易帐的正确性，是对整个系统的运行状态有着重要的意义。
>
> 在一些飞行器（如波音777）的系统中也有使用拜占庭容错。而且由于是即时系统，容错的功能也要能尽快回复，比如即使系统中有错误发生，容错系统也只能做出一微秒以内的延迟。
>
> 一些航天飞机的飞行系统甚至将容错功能放到整个系统的设计之中。
>
> 拜占庭容错机制是将收到的消息（或是收到的消息的签名）转交到其他的接收者。这类机制都假设它们转交的消息都可能念有拜占庭问题。在高度安全要求的系统中，这些假设甚至要求证明**错误能在一个合理的等级下被排除**。当然，要证明这点，首先遇到的问题就是如何有效的找出所有可能的、应能被容错的错误。这时候会试着在系统中加入错误插入器。 

### eth共识

#### Beacon

Beacon：信标

eth2将要升级的共识机制，即将使用的基于eth1和PoS算法的共识。

信标链不支持叔块了。

信标链和经典链在校验header时的不同：

>(a) The following fields are expected to be constants:
> - difficulty is expected to be 0 -- 难度固定为0
> - nonce is expected to be 0 -- 随机数固定为0
> - unclehash is expected to be Hash(emptyHeader)
> to be the desired constants -- 叔块哈希固定为空值
>
>(b) the timestamp is not verified anymore
>(c) the extradata is limited to 32 bytes

切换点：`TerminalTotalDifficulty is the amount of total difficulty reached by the network that triggers the consensus upgrade.`

[eth2升级](https://ethereum.org/zh/eth2/)

引入质押

计算当前块的basefee:

> If the parent gasUsed is the same as the target, the baseFee remains unchanged.
> 
> If the parent block used more gas than its target, the baseFee should increase.
>
> Otherwise if the parent block used less gas than its target, the baseFee should decrease.

其中，target：`parentGasTarget = parent.GasLimit / params.ElasticityMultiplier`，**params.ElasticityMultiplier**是常量，值为2.

具体计算过程在`consensus/misc/eip1559.go`的`CalcBaseFee`函数里。

#### PoW

`consensus/ethash`

未来15秒以内的块都算是正常的块，每个块最多2个叔块。

校验叔块时，获取最近6个高度的块信息，作为检验叔块是否合法的依据。

1.如何挖出一个新块？

`FinalizeAndAssemble`

```go
// FinalizeAndAssemble implements consensus.Engine, accumulating the block and
// uncle rewards, setting the final state and assembling the block.
func (ethash *Ethash) FinalizeAndAssemble(
    chain consensus.ChainHeaderReader,
    header *types.Header,
    state *state.StateDB,
    txs []*types.Transaction,
    uncles []*types.Header,
    receipts []*types.Receipt,
) (
    *types.Block,
    error,
)
```

2.如何将这个块上链？

```go
// Seal generates a new sealing request for the given input block and pushes
// the result into the given channel.
//
// Note, the method returns immediately and will send the result async. More
// than one result may also be returned depending on the consensus algorithm.
// 
// Seal implements consensus.Engine, attempting to find a nonce that satisfies
// the block's difficulty requirements.
func (ethash *Ethash) Seal(
    chain consensus.ChainHeaderReader,
    block *types.Block,
    results chan<- *types.Block,
    stop <-chan struct{},
) error

// mine is the actual proof-of-work miner that searches for a nonce starting from
// seed that results in correct final block difficulty.
func (ethash *Ethash) mine(
    block *types.Block,
    id int,
    seed uint64,
    abort chan struct{},
    found chan *types.Block,
)

// hashimotoFull aggregates data from the full dataset (using the full in-memory
// dataset) in order to produce our final value for a particular header hash and
// nonce.
func hashimotoFull(
    dataset []uint32,
    hash []byte,
    nonce uint64,
) (
    digest []byte,
    result []byte,
)
```

先获得一个随机数，然后递增该随机数，直到计算出符合要求的结果。-- 除结果外，还有一个`digest`。

最后将该随机数和`digest`保存到块的header里。

```go
func (w *worker) resultLoop() {
	defer w.wg.Done()
	for {
		select {
		case block := <-w.resultCh:
			// Short circuit when receiving empty result.
			if block == nil {
				continue
			}
			// Short circuit when receiving duplicate result caused by resubmitting.
			if w.chain.HasBlock(block.Hash(), block.NumberU64()) {
				continue
			}
			var (
				sealhash = w.engine.SealHash(block.Header())
				hash     = block.Hash()
			)
			w.pendingMu.RLock()
			task, exist := w.pendingTasks[sealhash]
			w.pendingMu.RUnlock()
			if !exist {
				log.Error("Block found but no relative pending task", "number", block.Number(), "sealhash", sealhash, "hash", hash)
				continue
			}
			// Different block could share same sealhash, deep copy here to prevent write-write conflict.
			var (
				receipts = make([]*types.Receipt, len(task.receipts))
				logs     []*types.Log
			)
			for i, taskReceipt := range task.receipts {
				receipt := new(types.Receipt)
				receipts[i] = receipt
				*receipt = *taskReceipt

				// add block location fields
				receipt.BlockHash = hash
				receipt.BlockNumber = block.Number()
				receipt.TransactionIndex = uint(i)

				// Update the block hash in all logs since it is now available and not when the
				// receipt/log of individual transactions were created.
				receipt.Logs = make([]*types.Log, len(taskReceipt.Logs))
				for i, taskLog := range taskReceipt.Logs {
					log := new(types.Log)
					receipt.Logs[i] = log
					*log = *taskLog
					log.BlockHash = hash
				}
				logs = append(logs, receipt.Logs...)
			}

            // ===== 提交块，保存状态到数据库 =====

			// Commit block and state to database.
			_, err := w.chain.WriteBlockAndSetHead(block, receipts, logs, task.state, true)
			if err != nil {
				log.Error("Failed writing block to chain", "err", err)
				continue
			}
			log.Info("Successfully sealed new block", "number", block.Number(), "sealhash", sealhash, "hash", hash,
				"elapsed", common.PrettyDuration(time.Since(task.createdAt)))

			// Broadcast the block and announce chain insertion event
			w.mux.Post(core.NewMinedBlockEvent{Block: block})

			// Insert the block into the set of pending ones to resultLoop for confirmations
			w.unconfirmed.Insert(block.NumberU64(), block.Hash())

		case <-w.exitCh:
			return
		}
	}
}
```

#### PoA

`consensus/clique`

### filecoin

在某段时间里存储着某些内容。

需要证明存储了内容，并且存储了一段时间。

PoST

[关键过程：P1, P2, C1 C2](https://new.qq.com/omn/20210116/20210116A06N5600.html)

[密封接口定义](https://github.com/filecoin-project/specs-storage)

```go
type Sealer interface {
	SealPreCommit1(ctx context.Context, sector SectorRef, ticket abi.SealRandomness, pieces []abi.PieceInfo) (PreCommit1Out, error)
	SealPreCommit2(ctx context.Context, sector SectorRef, pc1o PreCommit1Out) (SectorCids, error)

	SealCommit1(ctx context.Context, sector SectorRef, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, cids SectorCids) (Commit1Out, error)
	SealCommit2(ctx context.Context, sector SectorRef, c1o Commit1Out) (Proof, error)

	FinalizeSector(ctx context.Context, sector SectorRef, keepUnsealed []Range) error

	// ReleaseUnsealed marks parts of the unsealed sector file as safe to drop
	//  (called by the fsm on restart, allows storage to keep no persistent
	//   state about unsealed fast-retrieval copies)
	ReleaseUnsealed(ctx context.Context, sector SectorRef, safeToFree []Range) error
	ReleaseSectorKey(ctx context.Context, sector SectorRef) error
	ReleaseReplicaUpgrade(ctx context.Context, sector SectorRef) error

	// Removes all data associated with the specified sector
	Remove(ctx context.Context, sector SectorRef) error

	// Generate snap deals replica update
	ReplicaUpdate(ctx context.Context, sector SectorRef, pieces []abi.PieceInfo) (ReplicaUpdateOut, error)

	// Prove that snap deals replica was done correctly
	ProveReplicaUpdate1(ctx context.Context, sector SectorRef, sectorKey, newSealed, newUnsealed cid.Cid) (ReplicaVanillaProofs, error)
	ProveReplicaUpdate2(ctx context.Context, sector SectorRef, sectorKey, newSealed, newUnsealed cid.Cid, vanillaProofs ReplicaVanillaProofs) (ReplicaUpdateProof, error)

	// GenerateSectorKeyFromData computes sector key given unsealed data and updated replica
	GenerateSectorKeyFromData(ctx context.Context, sector SectorRef, unsealed cid.Cid) error
}
```

[密封接口实现](https://github.com/filecoin-project/lotus/tree/master/extern/sector-storage)

```go
type Manager struct {
	ls         stores.LocalStorage
	storage    *stores.Remote
	localStore *stores.Local
	remoteHnd  *stores.FetchHandler
	index      stores.SectorIndex

	sched *scheduler

	storage.Prover

	workLk sync.Mutex
	work   *statestore.StateStore

	callToWork map[storiface.CallID]WorkID
	// used when we get an early return and there's no callToWork mapping
	callRes map[storiface.CallID]chan result

	results map[WorkID]result
	waitRes map[WorkID]chan struct{}
}
// Manager实现了以下两个接口
var (
	_ storage.Sealer = (*Manager)(nil)
	_ SectorManager = (*Manager)(nil)
)
```

[有限状态机FSM](https://github.com/filecoin-project/lotus/tree/master/extern/storage-sealing)

```go
type Sealing struct {
	Api      SealingAPI
	DealInfo *CurrentDealInfoManager

	feeCfg config.MinerFeeConfig
	events Events

	startupWait sync.WaitGroup

	maddr address.Address

	sealer  sectorstorage.SectorManager // 上述Manager实现了SectorManager接口
	sectors *statemachine.StateGroup
	sc      SectorIDCounter
	verif   ffiwrapper.Verifier
	pcp     PreCommitPolicy

	inputLk        sync.Mutex
	openSectors    map[abi.SectorID]*openSector
	sectorTimers   map[abi.SectorID]*time.Timer
	pendingPieces  map[cid.Cid]*pendingPiece
	assignedPieces map[abi.SectorID][]cid.Cid
	creating       *abi.SectorNumber // used to prevent a race where we could create a new sector more than once

	upgradeLk sync.Mutex
	toUpgrade map[abi.SectorNumber]struct{}

	notifee SectorStateNotifee
	addrSel AddrSel

	stats SectorStats

	terminator  *TerminateBatcher
	precommiter *PreCommitBatcher
	commiter    *CommitBatcher

	getConfig GetSealingConfigFunc
}

// 有限状态机规则
// 
// 一个map，key是状态，value是一个函数，函数接受一系列事件和当前的扇区信息，对扇区执行指定操作，最后返回
var fsmPlanners = map[SectorState]func(events []statemachine.Event, state *SectorInfo) (uint64, error) {
	// ...
}

// Sealing实现了go-statemachine包里的StateHandler接口
func (m *Sealing) Plan(events []statemachine.Event, user interface{}) (interface{}, uint64, error) 

// 根据上面定义的有限状态转换规则执行
func (m *Sealing) plan(events []statemachine.Event, state *SectorInfo) (func(statemachine.Context, SectorInfo) error, uint64, error) {
	// ...
}
```

[go-statemachine](github.com/filecoin-project/go-statemachine)

```go
type StateHandler interface {
	// returns
	Plan(events []Event, user interface{}) (interface{}, uint64, error)
}

// StateGroup 返回StateMachine
func (s *StateGroup) loadOrCreate(name interface{}, userState interface{}) (*StateMachine, error) {
	// ...

	res := &StateMachine{
		planner:  s.hnd.Plan,
		eventsIn: make(chan Event),

		name:      name,
		st:        s.sts.Get(name),
		stateType: s.stateType,

		stageDone: make(chan struct{}),
		closing:   make(chan struct{}),
		closed:    make(chan struct{}),
	}

	go res.run() // 启动状态机

	return res, nil
}
```

[共识库： rust-fil-proofs](https://github.com/filecoin-project/rust-fil-proofs)

> Storage Proofs (storage-proofs) A library for constructing storage proofs – including non-circuit proofs, corresponding SNARK circuits, and a method of combining them.
>
> Storage Proofs Core (storage-proofs-core) A set of common primitives used throughout the other storage-proofs sub-crates, including crypto, merkle tree, hashing and gadget interfaces.
>
> Storage Proofs PoRep (storage-proofs-porep) storage-proofs-porep is intended to serve as a reference implementation for Proof-of-Replication (PoRep), while also performing the heavy lifting for filecoin-proofs.
>
> Primary Components:
>
>> PoR (Proof-of-Retrievability: Merkle inclusion proof)
>>
>> DrgPoRep (Depth Robust Graph Proof-of-Replication)
>>
>> StackedDrgPoRep
>
> Storage Proofs PoSt (storage-proofs-post) storage-proofs-post is intended to serve as a reference implementation for Proof-of-Space-time (PoSt), for filecoin-proofs.
>
> Primary Components:
>
>>    PoSt (Proof-of-Spacetime)
>
> Filecoin Proofs (filecoin-proofs) A wrapper around storage-proofs, providing an FFI-exported API callable from C (and in practice called by lotus via cgo). Filecoin-specific values of setup parameters are included here.
