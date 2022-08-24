package main

import "testing"

func TestSum(t *testing.T) {
	t_result_1 := sum("192780", "33221")
	a_result_1 := "226001"

	t_result_2 := sum("1", "2")
	a_result_2 := "3"

	t_result_3 := sum("1", "200")
	a_result_3 := "201"

	t_result_4 := sum("1", "20")
	a_result_4 := "21"

	t_result_5 := sum("77378279320090", "61813294987513")
	a_result_5 := "139191574307603"

	if t_result_1 != a_result_1 {
		t.Errorf("got %s, wanted %s", t_result_1, a_result_1)
	}
	if t_result_2 != a_result_2 {
		t.Errorf("got %s, wanted %s", t_result_2, a_result_2)
	}
	if t_result_3 != a_result_3 {
		t.Errorf("got %s, wanted %s", t_result_3, a_result_3)
	}
	if t_result_4 != a_result_4 {
		t.Errorf("got %s, wanted %s", t_result_4, a_result_4)
	}
	if t_result_5 != a_result_5 {
		t.Errorf("got %s, wanted %s", t_result_5, a_result_5)
	}
}
