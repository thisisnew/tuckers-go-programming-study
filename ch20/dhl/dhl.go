package dhl

import "fmt"

type Dhl struct {
}

func (d Dhl) Send(parcel string) {
	fmt.Printf("dhl에서 해외배송택배 %v를 보냅니다.\n", parcel)
}
