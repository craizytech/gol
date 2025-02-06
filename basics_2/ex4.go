// encoding interfaces to JSON without fixing the number of fields
// in your structs

package main

type Customer map[string]interface{}
type Name map[string]interface{}
type Address map[string]interface{}
