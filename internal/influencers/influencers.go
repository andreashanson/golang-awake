package influencers

import "time"

type Influencer struct {
	ID        string
	Name      string
	Lastname  string
	Email     string
	CreatedAt time.Time
}
