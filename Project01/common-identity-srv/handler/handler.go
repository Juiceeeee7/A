package handler

import (
	context "context"
	common_identity "github.com/Juice/Project01/proto/common-identity"
)

type Imp struct {

}

func (i Imp) GeneratorIDCode16Req(ctx context.Context, req *common_identity.StGeneratorIDCode16Req, rsp *common_identity.StGeneratorIDCode16Rsp) error {
	panic("implement me")
}

func (i Imp) GeneratorIDCode32Req(ctx context.Context, req *common_identity.StGeneratorIDCode32Req, rsp *common_identity.StGeneratorIDCode32Rsp) error {
	panic("implement me")
}

