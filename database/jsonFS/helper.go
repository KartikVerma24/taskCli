package jsonfs

import (
	"os"
	"path/filepath"
)

func writeFileAtomic(filename string, data []byte) error {
	dir := filepath.Dir(filename)

	// here we will create a tmp file in the same directory and then write all the data in tmp file, later rename it. 
	// This is being done because of the json structure we have is bit complex then normal json. we have to append in the task array and thus have to keep in mind the commas and json format.
	// hence copying the current file data to new with the updates is much more simple
	// additionally it does not strain the memory as for a single user even if there are 1000s of task, it would hardly take 100s of KB 

	tmpFile, createTmpFileErr := os.CreateTemp(dir, "tmp-*")
	if createTmpFileErr != nil {
		return createTmpFileErr
	}
	tmpName := tmpFile.Name()

	// Ensure cleanup on failure
	defer os.Remove(tmpName)

	_, tmpWriteErr := tmpFile.Write(data) 
	if tmpWriteErr != nil {
		tmpFile.Close()
		return tmpWriteErr
	}

	syncErr := tmpFile.Sync(); 
	if syncErr != nil {
		tmpFile.Close()
		return syncErr
	}

	if err := tmpFile.Close(); err != nil {
		return err
	}

	// Atomic replace
	return os.Rename(tmpName, filename)
}
