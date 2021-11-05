package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
	"path/filepath"
)

func getPathDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ".config"), nil
}

func main() {
	word := os.Args[1]

	perms := make([]string,0)

	// common names that are generally present in a bucket name
	envs := []string{"dev","development","stage","stg","s3","staging","prod","production","test"}

	// appending the company name itself
	perms = append(perms, word)

	// allowed special chars according to aws bucket naming rules
	ds := []string{".","-",""} 

	// common_bucket_prefixes.txt contains a list of most common prefixes in a bucket name (of course😅), thanks to @nahamsec
	path, err := getPathDir()
	if err != nil {
		fmt.Println(".config directory not found in the root directory!")
		return
	} 
	file, err := os.Open(path+"/common_bucket_prefixes.txt")
	if err != nil {
		fmt.Println(".config is present but common_bucket_prefixes.txt file not found!")
		return
	} 
	list, err := ioutil.ReadAll(file)
	prefix_list := strings.Split(strings.ReplaceAll(string(list),"\r\n","\n"),"\n")

	// adding (1*(3)*9)*2 = 54 permutations (e.g. uber.dev, stg-uber, s3uber i.e. words from envs)
	for _, env := range envs {
		for _, i := range ds {
			perms = append(perms, word+i+env)
			perms = append(perms, env+i+word)
		}
	}

	// adding (1*(3)*198)*2 = 1188 permutations (e.g. admin-uber, uberbucket, backup.uber i.e. from common_bucket_prefixes.txt)
	for _, env := range prefix_list {
		for _, i := range ds {
			perms = append(perms, word+i+env)
			perms = append(perms, env+i+word)
		}
	}

	// adding (1*(3)*198*(3)*9)*6 = 96228 permutations
	for _, i := range ds {
		for _, prefix := range prefix_list {
			for _, j := range ds {
				for _, env := range envs {
					perms = append(perms, word+i+prefix+j+env)
					perms = append(perms, word+i+env+j+prefix)
					perms = append(perms, prefix+i+word+j+env)
					perms = append(perms, env+i+word+j+prefix)
					perms = append(perms, prefix+i+env+j+word)
					perms = append(perms, env+i+prefix+j+word)
				}
			}
		}
	}

	// removing duplicate permutations and all the words that have length > 63 (not allowed in bucket naming rules)
	check := make(map[string]int)
	for _, i := range perms {
		if len(i) <= 63 {
			check[i] = 1
		}
	}
	res := make([]string,0)
	for str := range check {
		res = append(res, str)
	}

	res_wfile, err := os.Create(word+"-s3prefix.txt")
	if err != nil {
		panic(err)
	}
	defer res_wfile.Close()

	res_file, err := os.OpenFile(word+"-s3prefix.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer res_file.Close()

	for _, prefix := range res {
		if _, err := res_file.WriteString(prefix+"\n"); err != nil {
			panic(err)
		}
	}
}