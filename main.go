package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Identitas struct {
	tipe, nomor string
}

var (
	nLocker uint8
	lockers []Identitas
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		
		isSuccess, msg := runCommand(cmdString)
		if isSuccess == false {
			fmt.Fprint(os.Stderr, msg, "\n\n")
		} else {
			fmt.Fprint(os.Stdout, "\n")
		}
	}
}

func runCommand(commandStr string) (bool, string) {
	var msg string
	arrCommandStr := strings.Fields(commandStr)
	
	if len(arrCommandStr) < 1 {
		msg = "Mohon tuliskan perintah.."
		return false, msg
	} else if arrCommandStr[0] != "init" && arrCommandStr[0] != "help" && arrCommandStr[0] != "exit" && nLocker == 0 {
		msg = "Mohon tentukan jumlah locker terlebih dulu dengan perintah init [n]."
		return false, msg
	}
	
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	
	case "init":
		if nLocker != 0 {
			msg = fmt.Sprintf("Jumlah locker sudah ditentukan sebelumnya : %d", nLocker)
			return false, msg
		}
		
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [jumlah_locker]"
			return false, msg
		}
		
		n, err := strconv.ParseUint(arrCommandStr[1], 10, 8)
		
		if err != nil || n == 0 {
			msg = "Jumlah locker harus bilangan positif dari 1 sampai 255"
			return false, msg
		}
		
		nLocker = uint8(n)
		lockers = make([]Identitas, nLocker)
		fmt.Fprintln(os.Stdout, "Berhasil membuat locker dengan jumlah", nLocker)
		return true, "Success"
	
	case "input":
		if len(arrCommandStr) < 3 {
			msg = "Perintah ini memerlukan 2 argument : [KTP|SIM|Other] [nomor_ID]"
			return false, msg
		} else if arrCommandStr[1] != "KTP" && arrCommandStr[1] != "SIM" && arrCommandStr[1] != "Other" {
			msg = "Tipe identitas yang dapat digunakan hanya KTP, SIM atau Other."
			return false, msg
		}
		
		tLockerNum := findNomorIdentitas(arrCommandStr[2])
		if tLockerNum != 0 {
			msg = fmt.Sprintf("Kartu identitas dengan nomor %s sebelumnya sudah tersimpan di locker nomor %d", arrCommandStr[2], tLockerNum)
			return false, msg
		}
		
		lockerNum := inputLocker(arrCommandStr[1], arrCommandStr[2])
		
		if lockerNum == 0 {
			msg = "Maaf locker sudah penuh."
			return false, msg
		} else {
			msg = "Success"
			fmt.Fprintln(os.Stdout, "Kartu identitas tersimpan di locker nomor", lockerNum)
			return true, msg
		}
	
	case "status":
		fmt.Fprintln(os.Stdout, "No. Locker   Tipe Identitas   No. Identitas")
		
		for k, v := range lockers {
			if v.tipe == "" || v.nomor == "" {
				row := fmt.Sprintf("%9d.", k+1)
				fmt.Fprintln(os.Stdout, row)
			} else {
				row := fmt.Sprintf("%9d.   %-14s   %s", k+1, v.tipe, v.nomor)
				fmt.Fprintln(os.Stdout, row)
			}
		}
		
		return true, "Success"
	
	case "leave":
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [nomor_locker]"
			return false, msg
		}
		
		n, err := strconv.ParseUint(arrCommandStr[1], 10, 8)
		nl := uint8(n)
		
		if err != nil || nl == 0 || nl > nLocker {
			msg = "Nomor locker tidak tersedia."
			return false, msg
		}
		
		lockers[nl-1] = Identitas{"", ""}
		fmt.Fprintln(os.Stdout, "Loker nomor", n, "berhasil dikosongkan.")
		return true, "Success"
	
	case "find":
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [nomor_identitas]"
			return false, msg
		}
		
		tLockerNum := findNomorIdentitas(arrCommandStr[1])
		
		if tLockerNum != 0 {
			msg = "Found"
			fmt.Fprintln(os.Stdout, "Kartu identitas tersebut berada di locker nomor", tLockerNum)
		} else {
			msg = "Not Found"
			fmt.Fprintln(os.Stdout, "Nomor identitas tidak ditemukan.")
		}
		return true, msg
	
	case "search":
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [tipe_identitas]"
			return false, msg
		} else if arrCommandStr[1] != "KTP" && arrCommandStr[1] != "SIM" && arrCommandStr[1] != "Other" {
			msg = "Tipe identitas yang dapat digunakan hanya KTP, SIM dan Other"
			return false, msg
		}
		
		arrNomorId := []string{}
		isFound := false
		
		for _, v := range lockers {
			if v.tipe == arrCommandStr[1] {
				arrNomorId = append(arrNomorId, v.nomor)
				if isFound == false {
					isFound = true
				}
			}
		}
		
		if isFound {
			strNomorId := strings.Join(arrNomorId, ", ")
			msg = "Found"
			fmt.Fprintln(os.Stdout, "Nomor identitas yang tersimpan:", strNomorId)
		} else {
			msg = "Not Found"
			fmt.Fprintln(os.Stdout, "Tidak ditemukan Nomor Identitas dengan tipe", arrCommandStr[1])
		}
		return true, msg
		
	case "help":
		fmt.Fprintln(os.Stdout, "Daftar perintah yang bisa digunakan:\n\n init [jumlah_locker]\n status\n input [KTP|SIM|Other] [nomor_identitas]\n leave [nomor_locker]\n find [nomor_identitas]\n search [KTP|SIM|Other]\n help\n exit")
		return true, "Success"
		
	default:
		msg = "Perintah tidak dikenali. Gunakan 'help' untuk melihat daftar perintah yang dapat digunakan."
		return false, msg
	}
	
	return false, "Unknown error"
}

func findNomorIdentitas(nomor string) uint8 {
	for k, v := range lockers {
		if v.nomor == nomor {
			return uint8(k+1)
		}
	}
	return 0
}

func inputLocker(tipe, nomor string) uint8 {
	for k, v := range lockers {
		if v.tipe == "" && v.nomor == "" {
			lockers[k] = Identitas{tipe, nomor}
			return uint8(k+1)
		}
	}
	return 0
}