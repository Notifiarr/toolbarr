package app

func (a *App) GetDBTables(dbPath string) ([]string, error) {
	a.log.Tracef("Call:GetDBTables(%s)", dbPath)
	return []string{}, nil
}

func (a *App) GetTableSchema(dbPath, table string) (string, error) {
	a.log.Tracef("Call:GetTableSchema(%s,%s)", dbPath, table)
	return "", nil
}
