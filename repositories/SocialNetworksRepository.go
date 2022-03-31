package repositories

import (
	db "app/config"
	"app/models"
)

// type socialNetworkRepository interface {
// 	GetByProfileId(profileId int) error
// 	Update(newSN models.SocialNetwork)
// }

type SocialNetworkRepository struct {
	SocialNetworkId int
	ProfileId       int
}

// func GetSN() SocialNetwork {
// 	var sn models.SocialNetwork
// 	return SocialNetwork{
// 		SocialNetwork: sn,
// 	}
// }

func (sn *SocialNetworkRepository) GetById() (models.SocialNetwork, error) {
	var socialNetwork models.SocialNetwork

	err := db.GORM().Where("social_network_id = ?", sn.SocialNetworkId).First(&socialNetwork)
	if err != nil {
		return socialNetwork, err.Error
	}

	return socialNetwork, nil
}

func (sn *SocialNetworkRepository) GetByProfileId() (models.SocialNetwork, error) {
	var socialNetwork models.SocialNetwork

	err := db.GORM().First(&socialNetwork, "profile_id = ?", sn.ProfileId)
	if err != nil {
		return socialNetwork, err.Error
	}

	return socialNetwork, nil
}

func (sn *SocialNetworkRepository) UpdateById(newSN models.SocialNetwork) (models.SocialNetwork, error) {

	socialNetwork, err := sn.GetById()
	if err != nil {
		return socialNetwork, err
	}

	res := db.GORM().Model(&socialNetwork).Updates(newSN)

	return socialNetwork, res.Error
}
