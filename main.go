package main

import (
	"bufio"
	"errors"
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
		
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")	
	arrCommandStr := strings.Fields(commandStr)
	
	if len(arrCommandStr) < 1 {
		return errors.New("Mohon tuliskan perintah..")
	} else if (arrCommandStr[0] != "init") && (arrCommandStr[0] != "help") && (arrCommandStr[0] != "exit") && (nLocker == 0) {
		return errors.New("Mohon tentukan jumlah locker terlebih dulu dengan perintah init [n].\n")
	}
	
	switch arrCommandStr[0] {
		case "exit":
			os.Exit(0)
		
		case "init":
			if nLocker != 0 {
				msg := fmt.Sprintf("Jumlah locker sudah ditentukan sebelumnya : %d \n", nLocker)
				return errors.New(msg)
			}
			
			if len(arrCommandStr) < 2 {
				return errors.New("Perintah ini memerlukan 1 argument : [jumlah locker]")
			}
			
			n, err := strconv.ParseUint(arrCommandStr[1], 10, 8)
			
			if (err != nil) || (n == 0) {
				return errors.New("Error: jumlah locker harus bilangan positif dari 1 sampai 255")
			}
			
			nLocker = uint8(n)
			lockers = make([]map[string]string, nLocker)
			
			fmt.Fprintln(os.Stdout, "Berhasil membuat locker dengan jumlah", nLocker)
			return nil
		
		case "input":
			if len(arrCommandStr) < 3 {
				return errors.New("Perintah ini memerlukan 2 argument : [KTP|SIM|Other] [nomor ID]")
			} else if (arrCommandStr[1] != "KTP") && (arrCommandStr[1] != "SIM") && (arrCommandStr[1] != "Other") {
				return errors.New("Error: Tipe identitas yang dapat digunakan hanya KTP, SIM atau Other.\n")
			}
			
			tLockerNum := findNomorIdentitas(arrCommandStr[2])
			if tLockerNum != 0 {
				msg := fmt.Sprintf("Peringatan: Kartu identitas dengan nomor %s sebelumnya sudah tersimpan di locker nomor %d \n", arrCommandStr[2], tLockerNum)
				return errors.New(msg)
			}
			
			lockerNum := inputLocker(arrCommandStr[1], arrCommandStr[2])
			
			if lockerNum == 0 {
				return errors.New("Maaf locker sudah penuh.\n")
			} else {
				fmt.Fprintln(os.Stdout, "Kartu identitas tersimpan di locker nomor", lockerNum)
				return nil
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
			
			return nil
		
		case "leave":
			if len(arrCommandStr) < 2 {
				return errors.New("Perintah ini memerlukan 1 argument : [nomor locker]")
			}
			
			n, err := strconv.ParseUint(arrCommandStr[1], 10, 8)
			nl := uint8(n)
			
			if (err != nil) || (nl == 0) || (nl > nLocker) {
				return errors.New("Error: Nomor locker tidak tersedia.\n")
			}
			
			lockers[nl-1] = map[string]string{"tipeId": "", "nomorId": "", "status": "kosong"}
			fmt.Fprintln(os.Stdout, "Loker nomor", n, "berhasil dikosongkan.")
			return nil
		
		case "find":
			if len(arrCommandStr) < 2 {
				return errors.New("Perintah ini memerlukan 1 argument : [nomor identitas]")
			}
			
			tLockerNum := findNomorIdentitas(arrCommandStr[1])
			
			if tLockerNum != 0 {
				fmt.Fprintln(os.Stdout, "Kartu identitas tersebut berada di locker nomor", tLockerNum)
				return nil
			}
			return errors.New("Nomor identitas tidak ditemukan.\n")
		
		case "search":
			if len(arrCommandStr) < 2 {
				return errors.New("Perintah ini memerlukan 1 argument : [tipe identitas]")
			} else if (arrCommandStr[1] != "KTP") && (arrCommandStr[1] != "SIM") && (arrCommandStr[1] != "Other") {
				return errors.New("Tipe identitas yang dapat digunakan hanya KTP, SIM dan Other")
			}
			
			arrNomorId := []string{}
			isFound := false
			
			for _, val := range lockers {
				if (val["tipeId"] == arrCommandStr[1]) && (val["status"] == "terisi") {
					arrNomorId = append(arrNomorId, val["nomorId"])
					isFound = true
				}
			}
			
			if isFound {
				strNomorId := strings.Join(arrNomorId, ", ")
				fmt.Fprintln(os.Stdout, "Nomor identitas yang tersimpan:", strNomorId)
			} else {
				fmt.Fprintln(os.Stdout, "Tidak ditemukan Nomor Identitas dengan tipe", arrCommandStr[1])
			}
			return nil
			
		case "help":
			fmt.Fprintln(os.Stdout, "Daftar perintah yang bisa digunakan:\n\n init [jumlah_locker]\n status\n input [KTP|SIM|Other] [nomor_identitas]\n leave [nomor_locker]\n find [nomor_identitas]\n search [KTP|SIM|Other]\n help\n exit")
			return nil
			
		default:
			return errors.New("Perintah tidak dikenali. Gunakan 'help' untuk melihat daftar perintah yang dapat digunakan.\n")
	}
	
	return nil
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