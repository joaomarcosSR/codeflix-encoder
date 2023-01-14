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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.NewVideoRepository(db)
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	actualJob, err := repoJob.Find(job.ID)

	require.NotEmpty(t, actualJob.ID)
	require.Nil(t, err)
	require.Equal(t, actualJob.ID, job.ID)
	require.Equal(t, actualJob.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.NewVideoRepository(db)
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	actualJob, err := repoJob.Find(job.ID)

	require.NotEmpty(t, actualJob.ID)
	require.Nil(t, err)
	require.Equal(t, actualJob.Status, job.Status)
}
