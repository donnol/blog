
// Cond出现的字段必须分别在L和R均出现
export type Cond<L, R> = {
    l: keyof L
    r: keyof R
}

export function leftJoin<L, R, LR>(
    left: L[],
    right: R[],
    cond: Cond<L, R>,
    f: (l: L, r: R) => LR,
): LR[] {
    const rm = new Map()
    right.forEach((value: R, _index: number, _array: R[]) => {
        const rv = getProperty(value, cond.r)
        rm.set(rv, value)
    })

    const res = new Array(left.length)
    left.forEach((value: L, _index: number, _array: L[]) => {
        const lv = getProperty(value, cond.l)
        const rv = rm.get(lv)
        const lr = f(value, rv)
        res.push(lr)
    })

    return res
}

export function getProperty<T, K extends keyof T>(o: T, propertyName: K): T[K] {
    return o[propertyName]; // o[propertyName] is of type T[K]
}

const users = [
    {
        id: 1,
        name: "jd",
    },
    {
        id: 2,
        name: "jk",
    },
]

const articles = [
    {
        id: 1,
        userId: 1,
        name: "join"
    },
    {
        id: 2,
        userId: 2,
        name: "join 2"
    },
]

const res = leftJoin(users, articles, { l: "id", r: "userId" }, (l, r) => {
    return {
        id: r.id,
        name: r.name,
        userId: l.id,
        userName: l.name,
    }
})
console.log("res: ", res)
