// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	config   string
	password string
	username string
)

func main() {
	flag.StringVar(&config, "config", "/etc/secrets/config.json", "The configuration file.")
	flag.StringVar(&password, "password", "", "The password.")
	flag.StringVar(&username, "username", "", "The username.")
	flag.Parse()

	log.SetFlags(0)

	if os.Getenv("OSCON_PASSWORD") != "" {
		password = os.Getenv("OSCON_PASSWORD")
	}

	if os.Getenv("OSCON_USERNAME") != "" {
		username = os.Getenv("OSCON_USERNAME")
	}

	data, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Username: %s", username)
	log.Printf("Password: %s", password)
	log.Println("Config:")
	log.Println(string(data))
}
