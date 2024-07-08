package main

func main() {
	sites := [...]string{"linkedin.com/jobs/view", "glassdoor.com/job-listing", "boards.greenhouse.io", "jobs.lever.co", "myworkdayjobs.com"}

	for i := 0; i < len(sites); i++ {
		url := `site:${sites[i]} ("software engineer intern" OR "site reliability engineer intern") "United States"`
	}

}
