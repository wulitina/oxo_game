package repositories

import (
	"reflect"
	"testing"

	"oxo_game/internal/models"
)

func TestInMemoryLevelRepository_CreateAndGetById(t *testing.T) {
	repo := NewInMemoryLevelRepository()

	// 创建一个等级
	level := &models.Level{
		Name: "Beginner",
	}

	id, err := repo.Create(level)
	if err != nil {
		t.Fatalf("Error creating level: %v", err)
	}

	// 通过ID获取等级
	createdLevel, err := repo.GetById(id)
	if err != nil {
		t.Fatalf("Error fetching level by ID: %v", err)
	}

	// 检查创建的等级和获取的等级是否一致
	if !reflect.DeepEqual(createdLevel, level) {
		t.Errorf("Created level does not match expected. Expected %+v, got %+v", level, createdLevel)
	}
}

func TestInMemoryLevelRepository_GetById_NotFound(t *testing.T) {
	repo := NewInMemoryLevelRepository()

	_, err := repo.GetById(999) // 999 不存在的ID
	if err == nil {
		t.Error("Expected error for level not found, but got nil")
	}
}

func TestInMemoryLevelRepository_List(t *testing.T) {
	repo := NewInMemoryLevelRepository()

	// 创建几个等级
	levels := []*models.Level{
		{Name: "Beginner"},
		{Name: "Intermediate"},
		{Name: "Advanced"},
	}

	for _, level := range levels {
		_, err := repo.Create(level)
		if err != nil {
			t.Fatalf("Error creating level: %v", err)
		}
	}

	// 获取所有等级列表
	listedLevels := repo.List()

	// 检查结果数量是否符合预期
	if len(listedLevels) != len(levels) {
		t.Errorf("Expected %d levels, but got %d", len(levels), len(listedLevels))
	}

	// 检查每个等级是否都在返回的列表中
	for _, expectedLevel := range levels {
		found := false
		for _, listedLevel := range listedLevels {
			if expectedLevel.Name == listedLevel.Name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected level %s not found in listed levels", expectedLevel.Name)
		}
	}
}
