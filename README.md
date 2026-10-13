# Portfolio (Go)

Run: `go run .` then open http://localhost:8080  (set PORT to change)

## Make it yours
- Edit `content.go`: name, bio, skills, projects, email, GitHub, LinkedIn.
- Photo: save as `static/img/photo.jpg` (square works best). Until then a "G" avatar shows.
- Voice-over: record yourself and save as `static/audio/intro.mp3`. Until then the button reads the `Intro` text aloud with the browser's built-in voice.
- Demos: set `LiveURL` (opens the live site in a window) and/or `VideoURL` (e.g. `/static/assets/demo-prediction.mp4`) on a project in `content.go`.

## Deploy
Build a single binary with `go build -o portfolio .` (all files are embedded). Works on Render, Fly.io or any host that runs Go; it reads the PORT env var.
