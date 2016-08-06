package main

import "fmt"

type IPInfo struct {
	PublicIP  string
	PrivateIP string
}

func main() {
	countryCapitalMap := make(map[string][]string)
	//var countryCapitalMap map[string][]string
	/* create a map*/
	//countryCapitalMap = make(map[string]string)
	key := "China"
	value := "Beijing"
	countryCapitalMap[key] = append(countryCapitalMap[key], value)
	value = "Shanghai"
	countryCapitalMap[key] = append(countryCapitalMap[key], value)
	for k, v := range countryCapitalMap {
		fmt.Printf("contryCapitalMap[%s]=%d\n", k, len(v))
		for _, capital := range v {
			fmt.Printf("contryCapitalMap[%s]=%s\n", k, capital)
		}
	}
	/*
		weight_map := make(map[int][]IPInfo)
		ipinfo1 := IPInfo{"102.3.4.5", "10.2.3.3"}
		ipinfo11 := IPInfo{"103.3.4.5", "10.2.3.3"}
		ipinfo2 := IPInfo{"109.4.4.2", "10.3.4.5"}
		weight1 := 1
		weight2 := 2
		weight_map[weight1] = append(weight_map[weight1], ipinfo1)
		weight_map[weight1] = append(weight_map[weight1], ipinfo11)
		weight_map[weight2] = append(weight_map[weight2], ipinfo2)
		for k, v := range weight_map {
			fmt.Printf("weight_map[%d]=%d\n", k, len(v))
			for _, ipinfo := range v {
				fmt.Printf("PublicIP = %s, PrivateIP = %s\n", ipinfo.PublicIP, ipinfo.PrivateIP)
			}
		}
	*/
	weight_map := make(map[int][]*IPInfo)
	ipinfo1 := IPInfo{"102.3.4.5", "10.2.3.3"}
	ipinfo11 := IPInfo{"103.3.4.5", "10.2.3.3"}
	ipinfo2 := IPInfo{"109.4.4.2", "10.3.4.5"}
	weight1 := 1
	weight2 := 2
	weight_map[weight1] = append(weight_map[weight1], &ipinfo1)
	weight_map[weight1] = append(weight_map[weight1], &ipinfo11)
	weight_map[weight2] = append(weight_map[weight2], &ipinfo2)
	for k, v := range weight_map {
		fmt.Printf("weight_map[%d]=%d\n", k, len(v))
		for _, ipinfo := range v {
			fmt.Printf("PublicIP = %s, PrivateIP = %s\n", ipinfo.PublicIP, ipinfo.PrivateIP)
		}
	}

}
