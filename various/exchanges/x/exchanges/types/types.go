package types

import "fmt"

func (m *Pool) ShareCoinDenom() string {
	return fmt.Sprintf("share-%s", m.Denom)
}
