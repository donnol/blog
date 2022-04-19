package generic

type Cond[L, R any] struct {
	// LF keyof L
	// RF keyof R
}

// cond出现的字段必须分别在L和R均出现
// 类型和接口可以，但是类型的字段不行，粒度不够细
func leftJoin[L any, R any, LR any](
	left []L,
	right []R,
	cond Cond[L, R],
	f func(l L, r R) LR,
) []LR {

	return nil
}

// 缩窄范围，条件均使用指定字段Id，其类型为comparable
type IdConstraint[T comparable] interface {
	~struct {
		Id T

		// 怎么表示还有其它任意的字段呢
		Data any
	}
}

func leftJoinById[T comparable, L, R IdConstraint[T], LR any](left []L, right []R, f func(l L, r R) LR) []LR {
	// rm := make(map[T]R)
	// for _, one := range right {
	// 	rm[one.Id] = one // 暂未支持，1.19有望加上
	// }

	res := make([]LR, 0, len(left))
	// for _, one := range left {
	// 	r := rm[one.Id]
	// 	lr := f(one, r)
	// 	res = append(res, lr)
	// }

	return res
}
