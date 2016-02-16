package main

import (
	"fmt"
	. "github.com/arsham-f/Ceres/ceres"
	"encoding/json"
	"strings"
	"time"
)


/*
Commands, for now:

	CREATE 	[name] [ttl]
	DEL 	[name]
	INSERT 	[name] [series-value] [data]
	GET 	[name]
	TIME
*/



func handleCommand(input string) string {
	var cmd string;
	fmt.Sscanf(input, "%s %s\n", &cmd)

	cmd = strings.ToUpper(cmd)

	switch cmd {
		case "CREATE":
			return command_create(input)
		case "DEL":
			return command_del(input)
		case "INSERT":
			return command_insert(input)
		case "GET":
			return command_get(input)
		case "TIME":
			return command_time()
	}

	return fmt.Sprintf("Command not found (%s), see HELP.\n", cmd)
}

func command_create(input string) string {
	var cmd, name string
	var ttl int

	fmt.Sscanf(input, "%s %s %d\n", &cmd, &name, &ttl)

	err := CreateBucket(redis_client, name, ttl)

	if err != nil {
		warn("Unable to create bucket")
		return fmt.Sprintf("Unable to create bucket (%s): %s\n", name, err)
	}

	return "Successfully created\n"
}

func command_del(input string) string {
	var cmd, name string
	fmt.Sscanf(input, "%s %s\n", &cmd, &name)

	bucket, err := GetBucket(redis_client, name)

	if err != nil {
		return fmt.Sprintf("Could not find bucket (%s): %s\n", name, err)
	}

	err = bucket.Delete()

	if err != nil {
		return fmt.Sprintf("Could not delete bucket (%s): %s\n", name, err)
	}

	return "Deleted\n"
}

func command_insert(input string) string {
	var cmd, name, data string
	var value int

	fmt.Sscanf(input, "%s %s %d %s\n", &cmd, &name, &value, &data)

	bucket, err := GetBucket(redis_client, name)

	if err != nil {
		return fmt.Sprintf("Could not find bucket (%s): %s\n", name, err)
	}

	datab := []byte(strings.TrimSpace(input))

	i := 0
	spaces := 0
	flag := true
	for spaces = 0; spaces < 3 && i < len(datab); i++ {
		if datab[i] == byte(' ') && flag {
			spaces++
			flag = false
		} else if datab[i] != byte(' ') {
			flag = true
		}
	}

	datab = datab[i:]

	err = bucket.Add(value, strings.TrimSpace(string(datab)))

	if err != nil {
		return fmt.Sprintf("Could not insert data into bucket (%s): %s\n", name, err)
	}

	return "Successfully added\n"
}

func command_get(input string) string {
	var cmd, name string

	fmt.Sscanf(input, "%s %s\n", &cmd, &name)

	bucket, err := GetBucket(redis_client, name)

	if err != nil {
		return fmt.Sprintf("Could not find bucket (%s): %s\n", name, err)
	}

	data, err := bucket.Get()

	if err != nil {
		return fmt.Sprintf("Unable to retrieve data from bucket (%s): %s\n", name, err)
	}

	datab, err := json.Marshal(data)

	if err != nil {
		return fmt.Sprintf("Unable to marshall data: %s\n", err)
	}

	return string(datab) + "\n"
}

func command_time() string {
	return fmt.Sprintf("%d\n", time.Now().UTC().Unix())
}