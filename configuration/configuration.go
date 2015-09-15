package configuration

import (
	"net/url"
	"os"
	"time"
)

// Configuration of the application
type Configuration struct {
	headers            *map[string]string
	methods            *[]string
	outputFile         *string
	reportDirectory    *string
	urlResponseTimeout *time.Duration
	workersNumber      *uint64
	workerWaitPeriod   *time.Duration
	fuzzSetFile        **os.File
	baseURL            **url.URL
}

func newConfiguration() *Configuration {
	return new(Configuration)
}

// Headers method returns HTTP headers.
func (c *Configuration) Headers() (result map[string]string, exists bool) {
	if c.headers != nil {
		result = *c.headers
		exists = true
	}
	return result, exists
}

// Methods returns unique HTTP methods.
func (c *Configuration) Methods() []string {
	return *c.methods
}

// OutputFile method returns a path of the text output file.
func (c *Configuration) OutputFile() (result string, exists bool) {
	if c.outputFile != nil {
		result = *c.outputFile
		exists = true
	}
	return result, exists
}

// ReportDirectory method returns a target directory of HTML report files.
func (c *Configuration) ReportDirectory() (result string, exists bool) {
	if c.reportDirectory != nil {
		result = *c.reportDirectory
		exists = true
	}
	return result, exists
}

// URLResponseTimeout method returns fuzzed URL response timeout.
func (c *Configuration) URLResponseTimeout() time.Duration {
	return *c.urlResponseTimeout
}

// WorkersNumber returns a number of active fuzzing workers.
func (c *Configuration) WorkersNumber() uint64 {
	return *c.workersNumber
}

// WorkerWaitPeriod returns a period of time between two fuzzed requests per worker.
func (c *Configuration) WorkerWaitPeriod() time.Duration {
	return *c.workerWaitPeriod
}

// FuzzSetFile returns a file with fuzzed relative URLs.
func (c *Configuration) FuzzSetFile() os.File {
	return **c.fuzzSetFile
}

// BaseURL returns a base URL of the target website.
func (c *Configuration) BaseURL() url.URL {
	return **c.baseURL
}
