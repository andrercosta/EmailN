package campaign

import "github.com/andrercosta/EmailN/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(contract.NewCampaign) error {

	return nil
}
