package main

import "testing"

func TestGetCraneStatus(t *testing.T) {
	status := GetCraneStatus()
	if status.Status != "OPERATIONAL" {
		t.Errorf("Crane bermasalah! Status saat ini: %s", status.Status)
	}
}