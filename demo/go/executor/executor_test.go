package executor

import (
	"fmt"
	"testing"
)

func TestExecutorSimple(t *testing.T) {
	type obj struct {
	}
	execer := New[*obj](true)

	for _, tc := range []struct {
		name string
		ops  [][]Operation[*obj]
	}{
		{name: "serial",
			ops: [][]Operation[*obj]{
				{
					func(t *obj) error {
						fmt.Println("first")
						return nil
					},
				},
				{
					func(t *obj) error {
						fmt.Println("second")
						return nil
					},
				},
				{
					func(t *obj) error {
						fmt.Println("third")
						return nil
					},
				},
			},
		},
		{name: "alone",
			ops: [][]Operation[*obj]{
				{
					func(t *obj) error {
						fmt.Println("first")
						return nil
					},
					func(t *obj) error {
						fmt.Println("second")
						return nil
					},
					func(t *obj) error {
						fmt.Println("third")
						return nil
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := execer.Exec(&obj{}, tc.ops...); err != nil {
				t.Fatal(err)
			}
		})
	}
}

// === service layer ===

// 每个接口/功能方法里对应一个Tie，贯穿本方法所需要执行的包括db在内的多种操作；
// 后续Domain执行时需要使用它来获取数据
type UserBookTie struct {
	BookId   uint
	BookName string
	UserId   uint
	UserName string
}

// === db layer ===

// 假设要对数据库表操作，每个表对应一个结构体，其中包含对该表的各种操作（每种操作一个方法）
// 在实际的执行过程中，需要操作多个表，
// 根据某个参数获取某个表的数据，再根据获取的数据获取其它表里的更多数据
// 怎么把上一个操作拿到的结果给到下一个操作呢？ -- 结构体指针
// 通过泛型，添加`T any`约束参数来收集中间结果
type (
	user struct {
		id   uint
		name string
	}

	book struct {
		id     uint
		name   string
		userId uint // 关联`user.id`
	}
)

func (u *user) ById() error {
	// before
	fmt.Println(u.id)

	// do something by id and put the result to u

	// after
	u.name = "jd"
	fmt.Println(u)

	return nil
}

func (u *user) Create() error {
	// create

	return nil
}

func (b *book) Create() error {
	// create

	return nil
}

func (b *book) ById() error {
	// before
	fmt.Println(b.id)

	// do something by id and put the result to b

	// after
	b.name = "coding"
	b.userId = 1
	fmt.Println(b)

	return nil
}

func TestExecutorArgument(t *testing.T) {

	execer := New[*UserBookTie](true)

	// 此时，两个对象之间还没法通过`userId`关联上
	b := &book{id: 1}
	u := &user{}

	ubTie := &UserBookTie{}

	for _, tc := range []struct {
		name string
		ops  [][]Operation[*UserBookTie]
	}{
		{name: "serial-getByBookId",
			ops: [][]Operation[*UserBookTie]{
				// 在执行过程中，才能拿到`userId`
				{
					func(tie *UserBookTie) error {
						// 执行操作
						if err := b.ById(); err != nil {
							return err
						}

						// 往tie写入中间结果
						tie.BookId = b.id
						tie.BookName = b.name
						tie.UserId = b.userId
						fmt.Println("book tie: ", tie)

						return nil
					},
				},
				// 拿到`userId`后怎么给到`user`呢，通过Tie
				{
					func(tie *UserBookTie) error {
						// 从tie获取中间结果
						fmt.Println("user tie: ", tie)
						u.id = tie.UserId

						// 执行操作
						if err := u.ById(); err != nil {
							return err
						}

						tie.UserName = u.name

						return nil
					},
				},
			},
		},
		{name: "serial-createBook",
			ops: [][]Operation[*UserBookTie]{
				{
					func(tie *UserBookTie) error {
						u.id = 2
						u.name = "jn"

						// 执行操作
						if err := u.Create(); err != nil {
							return err
						}

						tie.UserId = u.id

						return nil
					},
				},
				{
					func(tie *UserBookTie) error {
						// 执行操作
						b := &book{
							id:     2,
							name:   "coding 2",
							userId: tie.UserId,
						}
						if err := b.Create(); err != nil {
							return err
						}

						fmt.Println("book: ", b)

						return nil
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := execer.Exec(ubTie, tc.ops...); err != nil {
				t.Fatal(err)
			}
			t.Logf("tie: %+v\n", ubTie)
		})
	}
}
