package statistics

func Orchestator(operation StatisticsOperation)(sf StatisticsFunction, err string){
	switch operation{
	case minimum:
		sf = Minimum
	}
}