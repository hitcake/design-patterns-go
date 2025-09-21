package db

import "testing"

func TestQuery(t *testing.T) {
	query := NewDefaultQuery(&MySqlDriver{})
	count, err := query.QueryForInt("SELECT count(*) as userCount FROM users")
	if err != nil {
		return
	}
	if count != 100 {
		t.Errorf("count should be 0, but got %d", count)
	}
}
