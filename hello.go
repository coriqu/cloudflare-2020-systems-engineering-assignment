package main
import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "flag"
    "os"
    "time"
    "sort"
)

type mytime struct {
    date        time.Duration
}
type timeArray []mytime
var sizeArray []float64
var timeArraySorted timeArray

func getTime(url string, times int) bool {
    // response, err := http.Get(url);
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
    	fmt.Println(err)
        return false
    }
    // This transport is what's causing unclosed connections.
    transCfg := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Timeout: 2 * time.Second, Transport: transCfg}
    response, err := client.Do(req)
    if err != nil {
    	fmt.Println(err)
        return false
    }
    defer response.Body.Close()
    // Start Time
    start := time.Now();
    // Read Body
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return false
    }
    //End Time
    end := time.Since(start)

    if times==0 {
    	fmt.Println(string(data))
    }
    curTime := mytime{date: end}
    size := float64(len(data))
	timeArraySorted = append(timeArraySorted, curTime)
	sizeArray = append(sizeArray, size)
	return true
}

func makeRequest(url string, profile int) {
    var successRequest, errorRequest int = 0, 0
    // Request
    for i := 0; i < profile; i++ {
        currRequest := getTime(url, i)
        if(currRequest) {
            successRequest++
        } else {
            errorRequest++
        }
    }
    sort.Slice(timeArraySorted, func(i, j int) bool {
        return timeArraySorted[i].date<timeArraySorted[j].date
    })
    sort.Slice(sizeArray, func(i, j int) bool {
        return sizeArray[i]<sizeArray[j]
    })   
    var sum time.Duration
    for _, num := range timeArraySorted {
        sum += num.date
    }
    newSum := int64(sum/time.Microsecond)
    avg := newSum/int64(successRequest)
    newAvg := time.Duration(avg) * time.Microsecond
    var median time.Duration
    if ( profile%2 == 1 ) {
        median = timeArraySorted[(successRequest-1)/2].date
    } else {
        median = (timeArraySorted[(successRequest-2)/2].date + timeArraySorted[(successRequest)/2].date)/2
    }
    ps := float64(successRequest)/float64(profile)*float64(100)
    es := float64(errorRequest)/float64(profile)*float64(100)

    fmt.Println("The requested link:", url)
    fmt.Println("The number of requests:", profile)
    fmt.Println("The fastest time:", timeArraySorted[0].date)
    fmt.Println("The slowest time:", timeArraySorted[successRequest-1].date)
    fmt.Println("The mean time:", newAvg)
    fmt.Println("The median time:", median)
    fmt.Println("The percentage of requests that succeeded:", ps, "%")
    fmt.Println("Any error codes returned that weren't a success:", es, "%")
    fmt.Println("The size in bytes of the smallest response:", sizeArray[0])
    fmt.Println("The size in bytes of the largest response:", sizeArray[successRequest-1])

}
func main() {
	url := flag.String("url", "", "A full URL")
	profile := flag.Int("profile", 0 , "The number of requests")
	flag.Parse()
	if (*url==""||*profile==0) {
        fmt.Println("Please enter the valid URL and profile")
        os.Exit(0)
    }
	fmt.Println("The number of requests:", *profile, "URL", *url)
	makeRequest(*url, *profile)

}

