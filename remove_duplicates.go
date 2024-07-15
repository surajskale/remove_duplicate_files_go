package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func calculateFileHash(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new SHA-256 hash
	hash := sha256.New()

	// Copy the file content to the hash
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	// Calculate and return the SHA-256 checksum
	return hash.Sum(nil), nil
}

func main() {
	// dir := "/home/suraj_kale/Me/"
	dir := "/media/suraj_kale/Transcend/CameraOri/"

	var duplicate_files []string

	d, err := os.Open(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)

	y, err := d.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	sha_map := make(map[string]string)

	for _, file := range y {
		fmt.Println(file)

		file_details, err := os.Stat(dir + file)
		hash := sha256.New()

		if err != nil {
			fmt.Println(err)
			continue
		}

		if !file_details.IsDir() {

			opened_file, err := os.Open(dir + file)
			if err != nil {
				if _, err := io.Copy(hash, opened_file); err != nil {
					panic(err)
				}
			}

			file_sha, err := calculateFileHash(dir + file)
			if err != nil {
				fmt.Println(err)
				continue
			}

			val, ok := sha_map[string(file_sha)]
			// val, ok := "", false

			if ok {
				fmt.Printf("Warning: %s is a duplicate file with SHA-256 hash %s\n", dir+file, val)
				duplicate_files = append(duplicate_files, dir+file)
			} else {
				sha_map[string(file_sha)] = dir + file
			}

			sha_map[string(file_sha)] = dir + file
			fmt.Println(dir + file + " is a file")
			defer opened_file.Close()

		}
	}

	// fmt.Println("File SHA-256 Hash Map:\n\n")
	// fmt.Println(sha_map)

	// for key, value := range sha_map {
	// 	fmt.Printf("SHA-256 Hash: %s, File Path: %s\n", key, value)
	// }
	fmt.Println("Total duplicate files: \n", len(duplicate_files))

	delete_duplicate_files := false

	if delete_duplicate_files {

		for _, duplicate_file := range duplicate_files {
			os.Remove(duplicate_file)
			fmt.Printf("Duplicate File: %s\n", duplicate_file)
		}
	}
}
