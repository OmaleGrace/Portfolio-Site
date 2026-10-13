# Portfolio

My personal portfolio website, built with Go. It introduces who I am, shows my projects with live demos and screenshots, and lists my skills.

**Live site:** https://omale-grace-portfolio.onrender.com  

## Features

- Dark, terminal-inspired design that works on phones and desktops
- Sections for About, Projects, Skills and Contact
- Project cards with feature lists, a screenshot gallery and a live demo window
- A voice intro: plays my own recording, with a browser voice as a fallback
- Typing effect on the hero line, respecting reduced-motion settings
- All content lives in one Go file, so updating the site means editing one place
- A small JSON API at `/api/projects` and a health check at `/healthz`

## Built with

- **Go** (standard library only): `net/http`, `html/template`, `embed`
- **HTML, CSS and vanilla JavaScript**, with no frameworks and no build step
- **Render** for hosting

## Project structure

```
.
├── main.go          # server, routes, template rendering
├── content.go       # all site content (name, bio, skills, projects, links)
├── templates/
│   └── index.html   # page template
└── static/
    ├── css/style.css
    ├── js/app.js
    ├── img/         # my photo and project screenshots
    └── audio/       # voice intro (intro.mp3)
```

Templates and static files are embedded into the program with `go:embed`, so the whole site is a single binary.

## Run locally

You need Go 1.22 or newer.

```bash
git clone https://github.com/OmaleGrace/Portfolio-Site.git
cd Portfolio-Site
go run .
```

Then open http://localhost:8080. Set the `PORT` environment variable to use a different port.

Because files are embedded when the program starts, stop it (Ctrl+C) and run `go run .` again after any change.


## Projects featured

- **Order-site:** a multi-vendor food ordering platform built with Go and PostgreSQL, with Paystack payments and an admin panel for managing vendors.
- **Prediction Site:** a sports and stocks prediction platform (in progress), pairing a Go server with a Python model.

## Author

**Omale Grace** (Omahlay), backend developer from Nigeria.
[GitHub](https://github.com/OmaleGrace) · [LinkedIn](https://www.linkedin.com/in/omale-grace)