package dictionary

import "fmt"

type Alipay struct {

}

func (a *Alipay) canUseFaceId() {

}

type Crash struct {

}

func (c Crash) Stolen() {

}

type CantainCanUseFaceId interface {

}

type CantainCrash interface {

}

func Print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceId:
		fmt.Printf("%T can use faceid\n", payMethod)
	case CantainCrash:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}
