package sqlc

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var Db *Queries

const (
	db_Driver = "postgres"
	db_Source = "postgresql://postgres:admin@localhost:5432/gateway_router?sslmode=disable"
)

func Init() {
	conn, err := sql.Open(db_Driver, db_Source)
	if err != nil {
		log.Println(err)
		log.Panic("unable to open the database connection")
	}
	if err = conn.Ping(); err != nil {
		log.Panic("unable to rach the database")
	}
	Db = New(conn)
}

type SystemInterface struct {
	Device     string `json:"device"`
	Ipv4       string `json:"ipv4"`
	Macaddress string `json:"macaddress"`
	Type       string `json:"type"`
	Mtu        int64  `json:"mtu"`
}

func getSystemInterfaces() []SystemInterface {
	cmd := exec.Command("ansible", "localhost", "-m", "ansible.builtin.setup", "-c", "local ")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute ansible command: %v\nStderr: %s", err, stderr.String())
	}
	// Split the output into lines and remove the first line
	output := stdout.String()
	jsonStart := strings.Index(output, "{")
	if jsonStart == -1 {
		log.Fatalf("Failed to find JSON in output: %s", output)
	}

	// Extract the JSON part
	cleanedOutput := output[jsonStart:]
	// Parse the JSON output
	var facts map[string]interface{}
	if err := json.Unmarshal([]byte(cleanedOutput), &facts); err != nil {
		log.Fatalf("Failed to parse JSON output: %v\nOutput:", err)
	}
	ansibleFacts, ok := facts["ansible_facts"].(map[string]interface{})
	if !ok {
		log.Fatal("Could not find ansible_facts in the output")
	}
	interfaces, ok := ansibleFacts["ansible_interfaces"].([]interface{})
	if !ok {
		log.Fatal("Could not find ansible_interfaces in ansible_facts")
	}
	systemInterfaces := []SystemInterface{}
	for _, iface := range interfaces {
		log.Println("ansible_" + iface.(string))
		interfaceDetails, ok := ansibleFacts["ansible_"+iface.(string)].(map[string]interface{})
		name := interfaceDetails["device"].(string)
		mac := interfaceDetails["macaddress"].(string)
		kind := interfaceDetails["type"].(string)
		mtu := interfaceDetails["mtu"].(string)
		mtu_int, err := strconv.Atoi(mtu)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		log.Println()
		var ipaddress string
		ip := interfaceDetails["ipv4"].([]interface{})
		if len(ip) != 0 {
			ipmap := ip[0].(map[string]interface{})
			ipaddress = ipmap["address"].(string)
		}
		Iface := SystemInterface{
			Device:     name,
			Macaddress: mac,
			Type:       kind,
			Ipv4:       ipaddress,
			Mtu:        int64(mtu_int),
		}
		systemInterfaces = append(systemInterfaces, Iface)
		if !ok {
			log.Fatalf("Could not find interface details for %s", "ansible_"+iface.(string))
		}
		// log.Println(interfaceDetails)

	}
	return systemInterfaces
}

func InitializeInterfaces() {
	system, err := Db.GetInitialisation(context.Background(), "interfaces")
	if err != nil {
		log.Println(err)
	}
	if !system.Initialised {
		log.Println("Initialising the interfaces...")
		interfaces := getSystemInterfaces()
		log.Println(interfaces)
		for _, i := range interfaces {
			if i.Type != "unknown" {
				args := CreateInterfaceParams{
					Name:       i.Device,
					Ipaddress:  i.Ipv4,
					Macaddress: i.Macaddress,
					// Type:       db.InterfaceTypes(i.Type),
					Mtu: sql.NullInt64{
						Int64: i.Mtu,
						Valid: true,
					},
				}
				_, err := Db.CreateInterface(context.Background(), args)
				if err != nil {
					log.Println(err)
				}
			}
		}
		log.Println("marking infterfaces as initialised..")
		args := MarkInitialisationParams{
			Component:   "interfaces",
			Initialised: true,
		}
		_, err := Db.MarkInitialisation(context.Background(), args)
		if err != nil {
			log.Println(err)
		}
	}
}
