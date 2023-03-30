package postgres_service

type Service struct {
	connectionString string

	maxConnections int
	maxLifetime    int
}

func New(connectionString string) *Service {
	return &Service{
		connectionString: connectionString,

		maxConnections: maxConnections,
		maxLifetime:    maxLifetime,
	}
}
