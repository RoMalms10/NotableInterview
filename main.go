package main

import "fmt"

// m = make(map[string]int) - make a map

// For Loop
// sum := 0
// for i := 1; i < 5; i++ {
// 	sum += i
// }

// For-each range loop
// strings := []string{"hello", "world"}
// for i, s := range strings {
// 	fmt.Println(i, s)
// }

// Golang OOP
// package payroll

// type Employee struct {
// 	Id    int
// 	name  string
// 	phone string
// 	email string
// }

// func (e *Employee) GetName() string {
// 	return e.name
// }
// func (e *Employee) SetName(name string) {
// 	e.name = name
// }

// How do we determine what orders should be added based on a set of scenarios
// Fingerprint:
//   reason: annual exam
//   age: 36
//   insurance: medicare

// Scenarios:
//   name: scenario 1
//   reason: annual exam
//   age: 36
//   insurance: medicare
//   orders: folic acid, complete blood count

//   name: scenario 2
//   reason: annual exam
//   age: 41 or older
//   insurance: medicare or aetna
//   orders: complete blood count, cancer screen

//   name: scenario 3
//   reason: annual exam
//   age: between 20 and 40
//   insurance: any insurance except medicare
//   orders: complete blood count, diabetes monitoring, cancer screen

//   name: scenario 4
//   reason: post operation hip
//   age: any age
//   insurance: any insurance
//   orders: xray hip, stitch removal

type Fingerprint struct {
	Reason    string
	Age       int
	Insurance string
}

type Scenario struct {
	Reason    string
	Age       []int
	Insurance []string
	Orders    string
}

func main() {
	scenarioMap := make(map[int]Scenario)
	scenarioMap[1] = Scenario{Reason: "annual exam", Age: []int{36}, Insurance: []string{"medicare"}, Orders: "folic acid, complete blood count"}
	scenarioMap[2] = Scenario{Reason: "annual exam", Age: []int{41, 9999}, Insurance: []string{"medicare", "aetna"}, Orders: "complete blood count, cancer screen"}
	scenarioMap[3] = Scenario{Reason: "annual exam", Age: []int{20, 40}, Insurance: []string{"!medicare", "aetna"}, Orders: "complete blood count, diabetes monitoring, cancer screen"}
	scenarioMap[4] = Scenario{Reason: "post operation hip", Age: []int{0, 9999}, Insurance: []string{"medicare", "aetna"}, Orders: "xray hip, stitch removal"}
	fingerprint := Fingerprint{Reason: "annual exam", Age: 36, Insurance: "medicare"}
	orders := requestedOrders(fingerprint, scenarioMap)
	fmt.Println(orders)

	fingerprint = Fingerprint{Reason: "annual exam", Age: 50, Insurance: "medicare"}
	orders = requestedOrders(fingerprint, scenarioMap)
	fmt.Println(orders)

	fingerprint = Fingerprint{Reason: "annual exam", Age: 30, Insurance: "aetna"}
	orders = requestedOrders(fingerprint, scenarioMap)
	fmt.Println(orders)

	fingerprint = Fingerprint{Reason: "annual exam", Age: 35, Insurance: "medicare"}
	orders = requestedOrders(fingerprint, scenarioMap)
	fmt.Println(orders)

	fingerprint = Fingerprint{Reason: "post operation hip", Age: 10, Insurance: "medicare"}
	orders = requestedOrders(fingerprint, scenarioMap)
	fmt.Println(orders)
}

func requestedOrders(fingerprint Fingerprint, scenarioMap map[int]Scenario) string {
	for _, scenario := range scenarioMap {
		if len(scenario.Age) > 1 {
			// fmt.Println(fingerprint.Age)
			// fmt.Println(scenario)

			// fmt.Println(fingerprint.Age >= scenario.Age[0] && fingerprint.Age <= scenario.Age[1])
			if fingerprint.Age >= scenario.Age[0] && fingerprint.Age <= scenario.Age[1] {
				if ok := checkInsurance(fingerprint.Insurance, scenario.Insurance); ok {
					if scenario.Reason == fingerprint.Reason {
						return scenario.Orders
					}
				}

			}
		} else {
			if scenario.Age[0] == fingerprint.Age && scenario.Reason == fingerprint.Reason {
				if ok := checkInsurance(fingerprint.Insurance, scenario.Insurance); ok {
					return scenario.Orders
				}

			}
		}

	}
	return ""
}

func checkInsurance(fingerprintInsurance string, scenarioInsurance []string) bool {
	for _, insurance := range scenarioInsurance {
		if fingerprintInsurance == insurance {
			return true
		}
	}
	return false
}
