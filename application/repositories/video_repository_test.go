package repositories_test

import (
	"codeflix-encoder/application/repositories"
	"codeflix-encoder/domain"
	"codeflix-encoder/infrastructure/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	actualVideo, err := repo.Find(video.ID)

	require.NotEmpty(t, actualVideo.ID)
	require.Nil(t, err)
	require.Equal(t, actualVideo.ID, video.ID)
}
