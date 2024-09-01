package repo

import "time"

type Article struct {
	Title     string
	Content   string
	Created   string
	Updated   string
	Author    string
	Published bool
}

func GetArticles() ([]Article, error) {
	articles := make([]Article, 0)
	created := time.Now().AddDate(0, 0, 1)
	updated := time.Now()
	articles = append(articles, Article{
		Title: "Why turtles live so long?",
		Content: `Turtles are known for their impressively long lifespans, and there are several factors that 
contribute to this longevity:<\br> 
1. Slow Metabolism<\br> 
2. Adaptation to Environmental Stress<\br> 
3. Efficient Cellular Repair`,
		Created:   created.Format(time.RFC822),
		Updated:   updated.Format(time.RFC822),
		Author:    "Turtle lover",
		Published: true,
	})
	articles = append(articles, Article{
		Title:     "Why sky is blue?",
		Content:   "TBA",
		Created:   created.Format(time.RFC822),
		Updated:   updated.Format(time.RFC822),
		Author:    "Sky dreamer",
		Published: false,
	})
	articles = append(articles, Article{
		Title:     "Two particles had collision in CERN",
		Content:   "TBA",
		Created:   created.Format(time.RFC822),
		Updated:   updated.Format(time.RFC822),
		Author:    "Nephew physician",
		Published: false,
	})

	return articles, nil
}
