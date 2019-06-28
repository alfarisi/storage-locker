package main

import (
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	var (
		isSuccess bool
		msg string
	)
	
	isSuccess, msg = runCommand("status")
	if isSuccess == true {
		t.Error("unexpected results: seharusnya tidak boleh menggunakan perintah 'status' atau lainnya sebelum dilakukan init.")
	} else {
		fmt.Println("[OK]", msg)
	}
	
	isSuccess, msg = runCommand("init 3")
	if isSuccess == false {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("doAnything")
	if isSuccess == true {
		t.Error("unexpected results: seharusnya tidak bisa mengeksekusi perintah 'doAnything'.")
	} else {
		fmt.Println("[OK]", msg)
	}
	
	isSuccess, msg = runCommand("input KTP 12345")
	if isSuccess == false {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("input KTPx 12345032423943")
	if isSuccess == true {
		t.Error("unexpected results: seharusnya tidak bisa menggunakan tipe identitas 'KTPx'.")
	} else {
		fmt.Println("[OK]", msg)
	}
	
	isSuccess, msg = runCommand("input KTP 12345")
	if isSuccess == true {
		t.Error("unexpected results: seharusnya tidak bisa menginput nomor identitas yang sama.")
	} else {
		fmt.Println("[OK]", msg)
	}
	
	isSuccess, msg = runCommand("input SIM 34567")
	if isSuccess == false {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("input Other 95445")
	if isSuccess == false {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("input KTP 92341")
	if isSuccess == true {
		t.Error("unexpected results: seharusnya locker sudah penuh.")
	} else {
		fmt.Println("[OK]", msg)
	}
	
	isSuccess, msg = runCommand("find 95445")
	if (isSuccess == false) || (isSuccess == true && msg == "Not Found") {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("search KTP")
	if (isSuccess == false) || (isSuccess == true && msg == "Not Found") {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("status")
	if isSuccess == false {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("leave 1")
	if isSuccess == false {
		t.Error("unexpected results:", msg)
	}
	
	isSuccess, msg = runCommand("leave 300")
	if isSuccess == true {
		t.Error("unexpected results: seharusnya tidak bisa dikosongkan karena nomor melebihi jumlah locker.")
	} else {
		fmt.Println("[OK]", msg)
	}
}