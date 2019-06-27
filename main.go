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
	} else if (arrCommandStr[0] != "init") && (nLocker == 0) {
		return errors.New("Mohon tentukan jumlah locker terlebih dulu dengan perintah init [jumlah].")
	}
	
	switch arrCommandStr[0] {
		case "exit":
			os.Exit(0)
			
		case "plus":
			if len(arrCommandStr) < 3 {
				return errors.New("Required for 2 arguments")
			}
			arrNum := []int64{}
			for i, arg := range arrCommandStr {
				if i == 0 {
					continue
				}
				n, _ := strconv.ParseInt(arg, 10, 64)
				arrNum = append(arrNum, n)
			}
			fmt.Fprintln(os.Stdout, sum(arrNum...))
			return nil
		
		case "init":
			if nLocker != 0 {
				msg := fmt.Sprintf("Jumlah locker sudah ditentukan sebelumnya : %d", nLocker)
				return errors.New(msg)
			}
			
			if len(arrCommandStr) < 2 {
				return errors.New("Perintah ini memerlukan 1 argument : jumlah locker")
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
				return errors.New("Error: Tipe identitas yang dapat digunakan hanya KTP, SIM atau Other")
			}
			
			tLockerNum := checkAlreadyStored(arrCommandStr[1], arrCommandStr[2])
			if tLockerNum != 0 {
				msg := fmt.Sprintf("Peringatan: Kartu identitas %s dengan nomor %s sebelumnya sudah tersimpan di locker nomor %d", arrCommandStr[1], arrCommandStr[2], tLockerNum)
				return errors.New(msg)
			}
			
			lockerNum := inputLocker(arrCommandStr[1], arrCommandStr[2])
			
			if lockerNum == 0 {
				return errors.New("Maaf locker sudah penuh")
			} else {
				fmt.Fprintln(os.Stdout, "Kartu identitas tersimpan di locker nomor", lockerNum)
				return nil
			}
		
		case "status":
			fmt.Fprintln(os.Stdout, "No. Locker   Tipe Identitas   No. Identitas")
			
			for k, val := range lockers {
				if (val == nil) || (val["status"] == "kosong") {
					row := fmt.Sprintf("%-3d", k+1)
					fmt.Fprintln(os.Stdout, row)
				} else {
					row := fmt.Sprintf("%-3d       %s        %s", k+1, val["tipeId"], val["nomorId"])
					fmt.Fprintln(os.Stdout, row)
				}
			}
			
			return nil
			
		case "help":
			fmt.Fprintln(os.Stdout, "Jenis perintah\n init [jumlah locker]\n status\n input [tipe identitas] [nomor identitas]\n leave [nomor locker]\n find [nomor identitas]\n search [tipe identitas]\n help\n exit\n")
			return nil
			
		default:
			return errors.New("Perintah tidak dikenali. Gunakan 'help' untuk melihat daftar perintah yang dapat digunakan.\n")
	}
	
	return nil
}

func sum(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}

func checkAlreadyStored(tipe, nomor string) uint8 {
	for k, val := range lockers {
		if (val["tipeId"] == tipe) && (val["nomorId"] == nomor) && (val["status"] == "terisi") {
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