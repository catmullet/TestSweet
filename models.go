package TestSweet

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"testing"
)

var TestSweet *testSweet
var testSweetLogs = make(map[string]string)
var testDataFolder string

func init() {
	TestSweet = new(testSweet)
	TestSweet.Logs = bytes.Buffer{}
}

type bootstrap func()
type teardown func()

type test struct {
	name    string
	results string
	test    func(t *testing.T)
}

func (t *test) Run(name string, f test) bool {
	fmt.Println("Running Test...")
	return t.Run(name, *t)
}

func (ts *testSweet) AddBootstapFunction(f ...bootstrap) {
	for _, val := range f {
		ts.Bootstraps = append(ts.Bootstraps, val)
	}
}

func (ts *testSweet) AddTeardownFunction(f ...teardown) {
	for _, val := range f {
		ts.TearDowns = append(ts.TearDowns, val)
	}
}

func (ts *testSweet) Bootstrap(t *testing.T) func(t *testing.T) {
	bootstrapTest()
	return func(t *testing.T) {
		tearDownTest()
	}
}

func (ts *testSweet) AddTest(t test) {
	ts.SuiteOfTests = append(ts.SuiteOfTests, t)
}

func bootstrapTest() {
	log.SetOutput(&TestSweet.Logs)
	for _, val := range TestSweet.Bootstraps {
		val()
	}
}

func tearDownTest() {
	for _, val := range TestSweet.TearDowns {
		val()
	}
	currentTest := GetCurrentTestInfo()
	testSweetLogs[currentTest] = TestSweet.Logs.String()
	fmt.Println(testSweetLogs[currentTest])
}

func GetCurrentTestInfo() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[1])
	file, line := f.FileLine(pc[1])
	return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
}

type testSweet struct {
	SuiteOfTests []test
	TestMain     func(m *testing.M)
	Bootstraps   []bootstrap
	TearDowns   []teardown
	Logs         bytes.Buffer
}

type stringData struct {
	Types map[string]struct {
		Good []string `json:"good"`
		Bad  []string `json:"bad"`
	} `json:"types"`
}

type intData struct {
	Types map[string]struct {
		Good []int `json:"good"`
		Bad  []int `json:"bad"`
	} `json:"types"`
}

type floatData struct {
	Types map[string]struct {
		Good []float64 `json:"good"`
		Bad  []float64 `json:"bad"`
	} `json:"types"`
}
