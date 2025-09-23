package result

import "log"

func PanicOnError(r *Result) {
	if r.Error != nil {
		log.Panicf(
			"exit: %d\nerror: %s\nstdout: %s\nstderr: %s\n",
			r.Exit,
			r.Error,
			r.OutputString,
			r.ErrorString,
		)
	}
}
