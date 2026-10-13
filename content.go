package main



type Project struct {
	Name     string   `json:"name"`
	Summary  string   `json:"summary"`
	Status   string   `json:"status"`
	Stack    []string `json:"stack"`
	LiveURL  string   `json:"liveUrl,omitempty"`  // opens in the demo window
	VideoURL string   `json:"videoUrl,omitempty"` // e.g. /static/assets/demo-order.mp4
	CodeURL  string   `json:"codeUrl,omitempty"`  // GitHub repo link
}

type SkillGroup struct {
	Title  string
	Skills []string
}

type Profile struct {
	Name     string
	Handle   string
	Role     string
	Location string
	Tagline  string
	About    []string
	Intro    string // text read aloud if no recorded voice-over exists
	Skills   []SkillGroup
	Projects []Project
	Email    string
	GitHub   string
	LinkedIn string
}

var profile = Profile{
	Name:     "Omale Grace",
	Handle:   "Omale Grace",
	Role:     "Backend Developer",
	Location: "Nigeria",
	Tagline:  "I build fast, reliable backends in Go, and I'm growing into Python for data and prediction work.",
	About: []string{
		"I'm Grace, a developer based in Nigeria. Go is my main language, and I use it to build backends, APIs and full-stack web apps.",
		"I learn by building real things: a food ordering marketplace with live payments, and a prediction platform that pairs a Go server with a Python model.",
	},
	Intro: "Hi, I'm Grace. I'm a backend developer from Nigeria. I build web apps with Go and Python. Take a look at my projects, and try the live demos.",
	Skills: []SkillGroup{
		{"Languages", []string{"Go", "Python", "SQL", "HTML/CSS", "JavaScript"}},
		{"Backend", []string{"REST APIs", "JWT auth", "bcrypt", "html/template", "pgx"}},
		{"Data & Infra", []string{"PostgreSQL", "SQLite", "Neon", "Docker", "Render", "GitHub Codespaces"}},
		{"Integrations", []string{"Paystack payments", "Elo rating models", "Brier scoring"}},
	},
	Projects: []Project{
		{
			Name:    "Order-site",
			Summary: "A full-stack food ordering web app that is growing into a marketplace where local vendors list their food. Includes signup/login, an admin menu manager, favourites, category filters and Paystack payments.",
			Status:  "Live",
			Stack:   []string{"Go", "PostgreSQL (Neon)", "Paystack", "Render"},
			LiveURL: "https://grace-ordering-site.onrender.com",
			CodeURL: "https://github.com/OmaleGrace/Order-site",
			// VideoURL: "/static/assets/demo-order.mp4",
		},
		{
			Name:    "Prediction Site",
			Summary: "A prediction platform for sports and stocks. A Python job fetches data and produces the system's own predictions with an Elo model. The Go server handles accounts, scoring (Brier-based) and the leaderboard, so users can compare their picks to the model.",
			Status:  "In progress",
			Stack:   []string{"Go", "Python", "PostgreSQL", "Docker", "JWT"},
			// LiveURL: "https://your-prediction-site.example",
			// VideoURL: "/static/assets/demo-prediction.mp4",
		},
	},
	Email:    "omalegrace209@gmail.com",
	GitHub:   "https://github.com/OmaleGrace",
	LinkedIn: "https://www.linkedin.com/in/omale-grace",
}
