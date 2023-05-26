package bowlertests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/markoxley/bowler"
	"github.com/markoxley/bowler/where"
)

func TestConnection(t *testing.T) {
	conf := getConnectionDetails()
	if err := bowler.Configure(conf); err != nil {
		t.Errorf("Unable to connect to database: %v", err)
	}
}

func TestNewTableCreation(t *testing.T) {
	reset()
	m := &TestModel{}
	bowler.Save(m)
	if !testTableExists("TestModel") {
		t.Errorf("Error testing for %v", "TestModel")
	}
}

func TestCount(t *testing.T) {
	reset()
	tm := &TestModel{}
	bowler.Save(tm)
	bowler.RawExecute("delete from TestModel")
	tm2 := &TestModel{}
	bowler.Save(tm2)
	i := bowler.Count[TestModel](nil)
	if i != 1 {
		t.Errorf("Expected 1, got %d", i)
	}
}

func TestGetRecord(t *testing.T) {
	reset()
	tm1 := &TestModel{
		Name: "Test1",
		Age:  42,
	}
	bowler.Save(tm1)
	cl := where.Equal("id", *tm1.ID).String()
	c := &bowler.Criteria{
		Where: cl,
	}
	tm3, err := bowler.First[TestModel](c)
	if err != nil {
		t.Error(err)
	}
	if *tm3.ID != *tm1.ID {
		t.Errorf("Expected ID %v, got %v", *tm1.ID, *tm3.ID)
	}
	if compareDates(tm3.CreateDate, tm1.CreateDate) {
		t.Errorf("Expected CreateDate %v, got %v", tm1.CreateDate, tm3.CreateDate)
	}
	if compareDates(tm3.LastUpdate, tm1.LastUpdate) {
		t.Errorf("Expected LastUpdate %v, got %v", tm1.LastUpdate, tm3.LastUpdate)
	}
}

func TestUpdateRecord(t *testing.T) {
	reset()
	tm1 := &TestModel{
		Name: "Test1",
		Age:  42,
	}
	bowler.Save(tm1)

	tm3, err := bowler.First[TestModel]()

	if err != nil {
		t.Error(err)
	}

	tm3.Age = 18
	tm3.Name = "David"

	bowler.Save(tm3)

	i := bowler.Count[TestModel]()
	if i != 1 {
		t.Errorf("Expected 1 record, found %d", i)
	}

	tm5, _ := bowler.First[TestModel]()

	if tm5.Age != tm3.Age {
		t.Errorf("Expected Age of %d, got %d", tm3.Age, tm5.Age)
	}
	if tm5.Name != tm3.Name {
		t.Errorf("Expected Name of %s, got %s", tm3.Name, tm5.Name)
	}
}
