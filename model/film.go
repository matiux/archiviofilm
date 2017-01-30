package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/djherbis/times"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"os"
	"time"
)

/**
 * atime = Access time
 * mtime = Modify time
 * ctime = Change time
 */
type Film struct {
	Id                                               string
	File, Checksum                                   string
	Seen                                             bool
	ATime, MTime, CTime, BTime, CreatedAt, UpdatedAt time.Time
}

func hash_file_md5(filePath string) (string, error) {

	var returnMD5String string

	file, err := os.Open(filePath)

	// info, err := os.Stat(filePath)
	// fmt.Println(info.ModTime().Format("2006-01-02 15:04:05"))

	if err != nil {
		return returnMD5String, err
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]

	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func NewFilm(path string, seen bool, withHash bool) Film {

	var fileHash string
	var cTime, bTime time.Time

	mTime := time.Now()

	t, err := times.Stat(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	if t.HasChangeTime() {
		cTime = t.ChangeTime()
	}

	if t.HasBirthTime() {
		bTime = t.BirthTime()
	}

	if withHash {

		fileHash, _ = hash_file_md5(path)
	}

	film := Film{
		Id:        uuid.NewV4().String(),
		File:      path,
		Checksum:  fileHash,
		Seen:      seen,
		ATime:     t.AccessTime(),
		MTime:     t.ModTime(),
		CTime:     cTime,
		BTime:     bTime,
		CreatedAt: mTime,
		UpdatedAt: mTime,
	}

	return film
}

func (f *Film) ToString() string {

	stringFilm := fmt.Sprintf("Id: %s - File: %v - Checksum: %v - Created: %v - Updated: %v", f.Id, f.File, f.Checksum, f.CreatedAt.Format("2006-01-02 15:04:05"), f.UpdatedAt.Format("2006-01-02 15:04:05"))

	return stringFilm
}
