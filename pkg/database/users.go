package database

type Bucket string

const (
	UsersInfo Bucket = "users"
	TasksInfo Bucket = "tasks"
)

type Users interface {
	Save(chatID int64, data string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}

type Tasks interface {
	SaveTasks(chatID int64, tasks string, bucket Bucket) error
	GetTasks(chatID int64, bucket Bucket) (string, error)
}
