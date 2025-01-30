package main

import (
	"context"
	"database/sql"
	db "gateway-router-db/sqlc"
	"gateway-router-db/utils"
	"log"
)

func main() {
	db.Init()
	
	system, err := db.Db.GetInitialisation(context.Background(), "interfaces")
	if err != nil {
		log.Println(err)
	}
	if !system.Initialised {
		log.Println("Initialising the interfaces...")
		interfaces := utils.GetSystemInterfaces()
		for _, i := range interfaces {
			if i.Type != "unknown" {
				args := db.CreateInterfaceParams{
					Name: sql.NullString{
						String: i.Device,
						Valid:  true,
					},
					Ipaddress:  i.Ipv4,
					Macaddress: i.Macaddress,
					Type:       db.InterfaceTypes(i.Type),
					Mtu: sql.NullInt64{
						Int64: i.Mtu,
						Valid: true,
					},
				}
				_, err := db.Db.CreateInterface(context.Background(), args)
				if err != nil {
					log.Println(err)
				}
			}
		}
		log.Println("marking infterfaces as initialised..")
		args := db.MarkInitialisationParams{
			Component:   "interfaces",
			Initialised: true,
		}
		_, err := db.Db.MarkInitialisation(context.Background(), args)
		if err != nil {
			log.Println(err)
		}
	}

	// log.Println(interfaces)
}

// func main() {
// 	cmd := exec.Command("ansible", "localhost", "-m", "ansible.builtin.setup", "-c", "local ")
// 	var stdout, stderr bytes.Buffer
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("Failed to execute ansible command: %v\nStderr: %s", err, stderr.String())
// 	}
// 	// Split the output into lines and remove the first line
// 	output := stdout.String()
// 	// lines := strings.Split(output, "\n")
// 	// if len(lines) > 1 {
// 	// 	output = strings.Join(lines[1:], "\n") // Join everything except the first line
// 	// }

// 	// Find where the JSON starts (the first `{`)
// 	jsonStart := strings.Index(output, "{")
// 	if jsonStart == -1 {
// 		log.Fatalf("Failed to find JSON in output: %s", output)
// 	}

// 	// Extract the JSON part
// 	cleanedOutput := output[jsonStart:]
// 	// Parse the JSON output
// 	var facts map[string]interface{}
// 	if err := json.Unmarshal([]byte(cleanedOutput), &facts); err != nil {
// 		log.Fatalf("Failed to parse JSON output: %v\nOutput:", err)
// 	}
// 	ansibleFacts, ok := facts["ansible_facts"].(map[string]interface{})
// 	if !ok {
// 		log.Fatal("Could not find ansible_facts in the output")
// 	}
// 	interfaces, ok := ansibleFacts["ansible_interfaces"].([]interface{})
// 	if !ok {
// 		log.Fatal("Could not find ansible_interfaces in ansible_facts")
// 	}
// 	for _, iface := range interfaces {
// 		log.Println("ansible_" + iface.(string))
// 		interfaceDetails, ok := ansibleFacts["ansible_"+iface.(string)].(map[string]interface{})
// 		if !ok {
// 			log.Fatalf("Could not find interface details for %s", "ansible_"+iface.(string))
// 		}
// 		log.Println(interfaceDetails)
// 	}
// 	// Print the gathered facts
// }
