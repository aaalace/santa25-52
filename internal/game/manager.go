package game

import (
	"github.com/go-redis/redis/v8"
	"math/rand"
	"santa25-52/internal/cache"
	"time"
)

type Manager struct {
	CacheClient *redis.Client
}

func (m *Manager) BuildSantaMap() map[string]string {
	names, _ := m.CacheClient.LRange(m.CacheClient.Context(), cache.TeamsDefaultKey, 0, -1).Result()

	available := append([]string(nil), names...)
	santaMap := make(map[string]string)

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	for _, santa := range names {
		var candidates []string
		for _, candidate := range available {
			if candidate != santa {
				candidates = append(candidates, candidate)
			}
		}

		recipientIndex := randGen.Intn(len(candidates))
		recipient := candidates[recipientIndex]

		santaMap[santa] = recipient

		for i, name := range available {
			if name == recipient {
				available = append(available[:i], available[i+1:]...)
				break
			}
		}
	}

	return santaMap
}
