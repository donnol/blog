package generic

type LC interface {
}

type RC interface {
}

type LRC interface {
}

type Cond string // 如: Id = UserId

// LC, RC, LRC和Cond之间的关系：
// cond出现的字段必须分别在LC和RC均出现
// 类型和接口可以，但是类型的字段不行，粒度不够细
func leftJoin[L LC, R RC, LR LRC](
	left []L,
	right []R,
	cond Cond,
	f func(l L, r R) LR,
) []LR {

	return nil
}
