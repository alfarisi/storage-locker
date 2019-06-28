package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nLocker uint8
	lockers []map[string]string
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
			fmt.Fprintln(os.Stderr, msg)
		}
	}
}

func runCommand(commandStr string) (bool, string) {
	var msg string
	commandStr = strings.TrimSuffix(commandStr, "\n")	
	arrCommandStr := strings.Fields(commandStr)
	
	if len(arrCommandStr) < 1 {
		msg = "Mohon tuliskan perintah.."
		return false, msg
	} else if (arrCommandStr[0] != "init") && (arrCommandStr[0] != "help") && (arrCommandStr[0] != "exit") && (nLocker == 0) {
		msg = "Mohon tentukan jumlah locker terlebih dulu dengan perintah init [n].\n"
		return false, msg
	}
	
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	
	case "init":
		if nLocker != 0 {
			msg = fmt.Sprintf("Jumlah locker sudah ditentukan sebelumnya : %d \n", nLocker)
			return false, msg
		}
		
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [jumlah_locker]"
			return false, msg
		}
		
		n, err := strconv.ParseUint(arrCommandStr[1], 10, 8)
		
		if (err != nil) || (n == 0) {
			msg = "Error: jumlah locker harus bilangan positif dari 1 sampai 255"
			return false, msg
		}
		
		nLocker = uint8(n)
		lockers = make([]map[string]string, nLocker)
		
		msg = "Success"
		fmt.Fprintln(os.Stdout, "Berhasil membuat locker dengan jumlah", nLocker)
		return true, msg
	
	case "input":
		if len(arrCommandStr) < 3 {
			msg = "Perintah ini memerlukan 2 argument : [KTP|SIM|Other] [nomor ID]"
			return false, msg
		} else if (arrCommandStr[1] != "KTP") && (arrCommandStr[1] != "SIM") && (arrCommandStr[1] != "Other") {
			msg = "Error: Tipe identitas yang dapat digunakan hanya KTP, SIM atau Other.\n"
			return false, msg
		}
		
		tLockerNum := findNomorIdentitas(arrCommandStr[2])
		if tLockerNum != 0 {
			msg = fmt.Sprintf("Peringatan: Kartu identitas dengan nomor %s sebelumnya sudah tersimpan di locker nomor %d \n", arrCommandStr[2], tLockerNum)
			return false, msg
		}
		
		lockerNum := inputLocker(arrCommandStr[1], arrCommandStr[2])
		
		if lockerNum == 0 {
			msg = "Maaf locker sudah penuh.\n"
			return false, msg
		} else {
			msg = "Success"
			fmt.Fprintln(os.Stdout, "Kartu identitas tersimpan di locker nomor", lockerNum)
			return true, msg
		}
	
	case "status":
		fmt.Fprintln(os.Stdout, "No. Locker   Tipe Identitas   No. Identitas")
		
		for k, val := range lockers {
			if (val == nil) || (val["status"] == "kosong") {
				row := fmt.Sprintf("%9d.", k+1)
				fmt.Fprintln(os.Stdout, row)
			} else {
				row := fmt.Sprintf("%9d.   %-14s   %s", k+1, val["tipeId"], val["nomorId"])
				fmt.Fprintln(os.Stdout, row)
			}
		}
		
		msg = "Success"
		return true, msg
	
	case "leave":
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [nomor locker]"
			return false, msg
		}
		
		n, err := strconv.ParseUint(arrCommandStr[1], 10, 8)
		nl := uint8(n)
		
		if (err != nil) || (nl == 0) || (nl > nLocker) {
			msg = "Error: Nomor locker tidak tersedia.\n"
			return false, msg
		}
		
		lockers[nl-1] = map[string]string{"tipeId": "", "nomorId": "", "status": "kosong"}
		msg = "Success"
		fmt.Fprintln(os.Stdout, "Loker nomor", n, "berhasil dikosongkan.")
		return true, msg
	
	case "find":
		if len(arrCommandStr) < 2 {
			msg = "Perintah ini memerlukan 1 argument : [nomor identitas]"
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
			msg = "Perintah ini memerlukan 1 argument : [tipe identitas]"
			return false, msg
		} else if (arrCommandStr[1] != "KTP") && (arrCommandStr[1] != "SIM") && (arrCommandStr[1] != "Other") {
			msg = "Tipe identitas yang dapat digunakan hanya KTP, SIM dan Other"
			return false, msg
		}
		
		arrNomorId := []string{}
		isFound := false
		
		for _, val := range lockers {
			if (val["tipeId"] == arrCommandStr[1]) && (val["status"] == "terisi") {
				arrNomorId = append(arrNomorId, val["nomorId"])
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
		msg = "Success"
		return true, msg
		
	default:
		msg = "Perintah tidak dikenali. Gunakan 'help' untuk melihat daftar perintah yang dapat digunakan.\n"
		return false, msg
	}
	
	msg = "Success"
	return true, msg
}

func findNomorIdentitas(nomor string) uint8 {
	for k, val := range lockers {
		if (val["nomorId"] == nomor) && (val["status"] == "terisi") {
			return uint8(k+1)
		}
	}
	return 0
}

func inputLocker(tipe, nomor string) uint8 {
	for k, val := range lockers {
		if (val == nil) || (val["status"] == "kosong") {
			lockers[k] = map[string]string{"tipeId": tipe, "nomorId": nomor, "status": "terisi"}
			return uint8(k+1)
		}
	}
	return 0
}