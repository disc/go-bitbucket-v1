/*
 * Bitbucket Server API
 *
 * Bitbucket Server API (former stash).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"

	sw "github.com/gfleury/go-bitbucket-v1/test/bb-mock-server/go"
)

func main() {
	log.Fatal(sw.RunServer(7991))
}
