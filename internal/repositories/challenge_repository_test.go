package repositories

import (
	"errors"
	"oxo_game/internal/models"
	"reflect"
	"testing"
)

func TestInMemoryChallengeRepository_CreateAndGetById(t *testing.T) {
	repo := NewInMemoryChallengeRepository()

	// 创建一个挑战
	challenge := &models.Challenge{
		PlayerID: 1,
		Won:      false,
	}

	id, err := repo.Create(challenge)
	if err != nil {
		t.Fatalf("Error creating challenge: %v", err)
	}

	// 通过ID获取挑战
	createdChallenge, err := repo.GetById(id)
	if err != nil {
		t.Fatalf("Error fetching challenge by ID: %v", err)
	}

	// 检查创建的挑战和获取的挑战是否一致
	if !reflect.DeepEqual(createdChallenge, challenge) {
		t.Errorf("Created challenge does not match expected. Expected %+v, got %+v", challenge, createdChallenge)
	}
}

func TestInMemoryChallengeRepository_GetById_NotFound(t *testing.T) {
	repo := NewInMemoryChallengeRepository()

	_, err := repo.GetById(999) // 999 不存在的ID
	if !errors.Is(err, ErrChallengeNotFound) {
		t.Errorf("Expected ErrChallengeNotFound, got %v", err)
	}
}

func TestInMemoryChallengeRepository_ListByPlayer(t *testing.T) {
	repo := NewInMemoryChallengeRepository()

	// 创建几个挑战，包含一个特定玩家的挑战
	challenges := []*models.Challenge{
		{
			PlayerID: 1,
			Won:      false,
		},
		{
			PlayerID: 1,
			Won:      true,
		},
		{
			PlayerID: 2,
			Won:      false,
		},
	}

	for _, challenge := range challenges {
		_, err := repo.Create(challenge)
		if err != nil {
			t.Fatalf("Error creating challenge: %v", err)
		}
	}

	// 获取特定玩家ID的挑战列表
	playerID := 1
	playerChallenges := repo.ListByPlayer(playerID)

	// 检查结果数量是否符合预期
	if len(playerChallenges) != 2 {
		t.Errorf("Expected 2 challenges for player ID %d, but got %d", playerID, len(playerChallenges))
	}

	// 检查挑战列表是否包含了预期的挑战
	for _, ch := range playerChallenges {
		if ch.PlayerID != playerID {
			t.Errorf("Expected player ID %d, got %d", playerID, ch.PlayerID)
		}
	}
}

func TestInMemoryChallengeRepository_ListLatest(t *testing.T) {
	repo := NewInMemoryChallengeRepository()

	// 创建一些挑战
	challenges := []*models.Challenge{
		{
			PlayerID: 1,
			Won:      false,
		},
		{
			PlayerID: 2,
			Won:      true,
		},
		{
			PlayerID: 3,
			Won:      false,
		},
	}

	for _, challenge := range challenges {
		_, err := repo.Create(challenge)
		if err != nil {
			t.Fatalf("Error creating challenge: %v", err)
		}
	}

	// 获取最新的2个挑战
	n := 2
	latestChallenges := repo.ListLatest(n)

	// 检查结果数量是否符合预期
	if len(latestChallenges) != n {
		t.Errorf("Expected %d latest challenges, but got %d", n, len(latestChallenges))
	}

	// 检查挑战列表是否按照创建时间的倒序排列
	for i := 0; i < len(latestChallenges)-1; i++ {
		if latestChallenges[i].CreatedAt.Before(latestChallenges[i+1].CreatedAt) {
			t.Errorf("Challenges are not in descending order by creation time")
		}
	}
}
