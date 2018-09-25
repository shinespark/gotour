package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// 注意: Error メソッドの中で、 fmt.Sprint(e) を呼び出すことは、
	// 無限ループのプログラムになることでしょう。
	// 最初に fmt.Sprint(float64(e)) として e を変換しておくことで、
	// これを避けることができます。 なぜでしょうか？
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 1; i <= 10; i++ {
		diff := (z*z - x) / (2 * z)

		if z == z-diff {
			return z, nil
		}

		z -= diff
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
