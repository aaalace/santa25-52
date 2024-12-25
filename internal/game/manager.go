package game

import (
	"gorm.io/gorm"
	"math/rand"
	"santa25-52/internal/db"
	"time"
)

type Manager struct {
	DbClient *gorm.DB
}

func (m *Manager) BuildSantaMap() map[db.Member]db.Member {
	// santa:recipient
	santaMap := make(map[db.Member]db.Member)

	var members []db.Member
	_ = m.DbClient.Find(&members)

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	availableMembers := append([]db.Member(nil), members...)
	for _, santa := range members {
		var candidates []db.Member
		for _, candidate := range availableMembers {
			if candidate.ID != santa.ID {
				candidates = append(candidates, candidate)
			}
		}

		recipientIndex := randGen.Intn(len(candidates))
		recipient := candidates[recipientIndex]

		santaMap[santa] = recipient
		for i, name := range availableMembers {
			if name == recipient {
				availableMembers = append(availableMembers[:i], availableMembers[i+1:]...)
				break
			}
		}
	}

	return santaMap
}
