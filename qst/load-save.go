package qst

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var questPath = "."

// Load loads a questionaire from a JSON file.
func Load(fn string) (*QuestionaireT, error) {

	q := QuestionaireT{}

	bts, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Printf("Could not read file %v : %v", fn, err)
		return &q, err
	}

	err = json.Unmarshal(bts, &q)
	if err != nil {
		log.Printf("Unmarshal %v: %v", fn, err)
		return &q, err
	}

	// Checksum
	bts = bytes.Replace(bts, []byte(q.MD5), []byte("md5dummy"), 1) // replace once to save memory
	got := md5Str(bts)
	if got != q.MD5 {
		return &q, fmt.Errorf("MD5 hashes differ; want - got\n%v\n%v", q.MD5, got)
	}

	return &q, nil
}

// Save a questionaire to JSON
func (q *QuestionaireT) Save(fn ...string) error {

	q.MD5 = "md5dummy"

	firstColLeftMostPrefix := " "
	bts, err := json.MarshalIndent(q, firstColLeftMostPrefix, "\t")
	if err != nil {
		return err
	}

	// The MD5 value is set *after* serialization, through bytes.Replace
	hsh := md5Str(bts)
	bts = bytes.Replace(bts, []byte(q.MD5), []byte(hsh), 1) // replace once to save memory
	q.MD5 = hsh

	saveDir := path.Dir(questPath)
	err = os.Chmod(saveDir, 0755)
	if err != nil {
		return err
	}

	questFile := path.Base(questPath)
	if len(fn) > 0 {
		questFile = fn[0]
	}

	pthOld := path.Join(saveDir, questFile)

	keepBackup := false
	if keepBackup {
		fileBackup := strings.Replace(questFile, ".json", fmt.Sprintf("_%v.json", time.Now().Unix()), 1)
		pthBackup := path.Join(saveDir, fileBackup)
		if questFile != "questionaire-example.json" {
			err = os.Rename(pthOld, pthBackup)
			if err != nil {
				return err
			}
		}
	}

	err = ioutil.WriteFile(pthOld, bts, 0644)
	if err != nil {
		return err
	}
	log.Printf("Saved questionaire file to %v", pthOld)
	return nil
}

// Md5Str computes the md5 hash of a byte slice.
func md5Str(buf []byte) string {
	hasher := md5.New()
	hasher.Write(buf)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
