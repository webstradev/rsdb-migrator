package importer

func (data *LoadedData) HandleEdgeCases() {
	// Handle known issue where the date was specified in in the incorrect format for a few articles
	for idx, article := range data.Articles {
		if article.Date == "2017" {
			article.Date = "09/09/2017"
			data.Articles[idx] = article
		}
		if article.Date == "2019" {
			article.Date = "09/09/2017"
			data.Articles[idx] = article
		}
	}

	// Handle known issue where the date was specified in in the incorrect format for a few projects
	for idx, project := range data.Projects {
		if project.Date == "2017" {
			project.Date = "09/09/2017"
			data.Projects[idx] = project
		}
		if project.Date == "2019" {
			project.Date = "09/09/2017"
			data.Projects[idx] = project
		}
	}
}
