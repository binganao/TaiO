package jobs

var jobs chan string

func InitJobs() {
	jobs = make(chan string, 1000000000)
}

func GetJobs() chan string {
	return jobs
}

func AddJobs(raw string) {
	jobs <- raw
}

func DelJobs(raw string) {
	for i := range jobs {
		if i == raw {
			break
		} else {
			jobs <- i
			continue
		}
	}
}
