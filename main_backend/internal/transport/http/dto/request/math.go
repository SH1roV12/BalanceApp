package request

import "github.com/SH1roV12/balance/internal/gen/pb"

type Sum struct{
	First int32 `json:"first"`
	Second int32 `json:"second"`
}

func ToDomain(sum Sum)*pb.MathRequest{
	return &pb.MathRequest{
		Num1: sum.First,
		Num2: sum.Second,
	}
}