package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jakewarren/hackedemailsapi/api"
	"github.com/jinzhu/now"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
)

var (
	app = kingpin.New("hackedemails", "Un-official API client for hacked-emails.com.")

	debug      = app.Flag("debug", "print debug info").Short('d').Bool()
	filterDate = app.Flag("filter-date", "only print breaches released after specified date").Short('f').String()
	silent     = app.Flag("silent", "suppress response message, only display results").Short('s').Bool()

	email = app.Arg("email", "the email address to lookup.").Required().String()
)

func main() {
	app.Version("0.1.0").VersionFlag.Short('V')
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.ErrorLevel)

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	printResults(*email)

}

func printResults(email string) {
	//query results for the email address
	response, err := api.LookupEmail(email)
	if err != nil {
		fmt.Printf("Decoding api response as JSON failed: %v", err)
		return
	}

	//check if an invalid email was provided
	if response.Status == "badsintax" {
		log.Fatalf("query for %s was rejected. perhaps you did not provide a valid email address?", email)
	}

	breachCount := 0
	var defResponse string

	for _, breach := range response.Breaches {

		if *filterDate != "" {
			filterTime, err := now.Parse(*filterDate)
			if err != nil {
				log.WithError(err).Error("error parsing filter time")
			}

			releaseTime, err := time.Parse(time.RFC3339, breach.DateCreated)
			if err != nil {
				log.WithError(err).Error("error parsing released time")
			}

			if releaseTime.Before(filterTime) {
				log.Debugf("excluding %s (%s)", breach.Title, breach.DateCreated)
				continue
			}

		}

		defResponse += fmt.Sprintf("\n%s\n\tsource_url: %s \n\tdate_released:%s \n\tdate_leaked:%s\n", breach.Title, breach.SourceURL, breach.DateCreated, breach.DateLeaked)
		defResponse += fmt.Sprintf("\tsource_network: %s \n\temail_count: %d\n\tverified: %t\n", breach.SourceNetwork, breach.EmailsCount, breach.Verified)
		if *debug {
			defResponse += fmt.Sprintf("%#+v\n", breach)
		}
		breachCount++
	}

	if !*silent {
		if *filterDate == "" {
			fmt.Printf("%d breaches returned for %s\n", breachCount, response.Query)
		} else {
			fmt.Printf("%d breaches returned for %s (%d filtered out)\n", breachCount, response.Query, (len(response.Breaches) - breachCount))
		}
	}

	fmt.Println(defResponse)
}
