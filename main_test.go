package main

import (
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	cmdString := "status"
	err := runCommand(cmdString)
	if err == nil {
		t.Error("unexpected results: seharusnya tidak bisa menggunakan perintah lain sebelum dilakukan init.")
	} else {
		fmt.Println("Perintah 'status' tidak bisa dieksekusi : OK")
	}
	
	cmdString = "init 5"
	err = runCommand(cmdString)
	if err != nil {
		t.Error("unexpected results:", err)
	}
	
	cmdString = "input KTP 12345"
	err = runCommand(cmdString)
	if err != nil {
		t.Error("unexpected results:", err)
	}
	
	cmdString = "input SIM 34567"
	err = runCommand(cmdString)
	if err != nil {
		t.Error("unexpected results:", err)
	}
	
	cmdString = "input Other 95445"
	err = runCommand(cmdString)
	if err != nil {
		t.Error("unexpected results:", err)
	}
	
	cmdString = "status"
	err = runCommand(cmdString)
	if err != nil {
		t.Error("unexpected results:", err)
	}
}