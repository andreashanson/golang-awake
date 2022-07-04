package mongo

import (
	"fmt"

	"github.com/andreashanson/golang-awake/internal/influencers"
)

type InfluencersRepo struct {
	conn *Connection
}

func NewInfluencersRepo(c *Connection) InfluencersRepo {
	return InfluencersRepo{conn: c}
}

func (mr InfluencersRepo) GetAll() []influencers.Influencer {
	fmt.Println("MONGO")
	var influ []influencers.Influencer
	return influ
}
