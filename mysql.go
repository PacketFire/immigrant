package main

type MysqlDriver struct {
}

func (this *MysqlDriver) Init(context map[string]string) {

}

func (this *MysqlDriver) Migrate(r Revision, c chan error) {

}

func (this *MysqlDriver) Rollback(r Revision, c chan error) {

}

func (this *MysqlDriver) State() *Revision {

	return new(Revision)
}

func (This *MysqlDriver) Close() {

}
