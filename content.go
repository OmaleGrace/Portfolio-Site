package main

type Image struct {
	Src     string `json:"src"`
	Caption string `json:"caption"`
}

type Project struct {
	Name     string   `json:"name"`
	Summary  string   `json:"summary"`
	Status   string   `json:"status"`
	Stack    []string `json:"stack"`
	Features []string `json:"features,omitempty"`
	Images   []Image  `json:"images,omitempty"`   
	LiveURL  string   `json:"liveUrl,omitempty"`  
	VideoURL string   `json:"videoUrl,omitempty"` 
	CodeURL  string   `json:"codeUrl,omitempty"`  
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
	Intro    string 
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
			Summary: "A full-stack food ordering platform that is open to any vendor. Local vendors join, list their meals and take their own orders, while customers browse, pay with Paystack and track their orders.",
			Status:  "Live",
			Stack:   []string{"Go", "PostgreSQL (Neon)", "Paystack", "Render"},
			Features: []string{
				"Multi-vendor marketplace: any vendor can list their food",
				"Admin vendor management: add, suspend and extend vendors",
				"Only active, paid-up vendors appear on the menu",
				"Paystack payments, favourites, cart and order tracking",
			},
			Images: []Image{
				{"/static/img/order/home.jpg", "Home page"},
				{"/static/img/order/menu.jpg", "Menu with vendor names and category filters"},
				{"/static/img/order/cart.jpg", "Cart and checkout"},
				{"/static/img/order/account.jpg", "Customer account"},
				{"/static/img/order/vendors.jpg", "Admin: vendor management"},
			},
			LiveURL: "https://grace-ordering-site.onrender.com",
			CodeURL: "https://github.com/OmaleGrace/Order-site",
		},
		{
			Name:    "Prediction Site",
			Summary: "A prediction platform for sports and stocks. A Python job fetches data and produces the system's own predictions with an Elo model. The Go server handles accounts, scoring (Brier-based) and the leaderboard, so users can compare their picks to the model.",
			Status:  "In progress",
			Stack:   []string{"Go", "Python", "PostgreSQL", "Docker", "JWT"},
		},
	},
	Email:    "omalegrace209@gmail.com",
	GitHub:   "https://github.com/OmaleGrace",
	LinkedIn: "https://www.linkedin.com/in/omale-grace",
}