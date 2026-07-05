# česka slova

Repo for the organisation of my learning stuff, especially vocabs with example sentences.

A small Go-based vocabulary management tool for learning Czech vocabulary from structured YAML files.

It supports multiple output formats (CLI, Markdown, TTS) and filtering by tags to enable focused learning sessions.

### 🚀 Features
- 📂 Load vocabulary from YAML files
- 🧩 Strongly typed Go data model
- 🖥️ Multiple output formats:
- CLI (human readable)
- Markdown (notes / GitHub / Obsidian)
- TTS (spoken practice)
- 🏷️ Tag-based filtering (food, lesson-7, etc.)
- 🔌 Extensible renderer architecture (add new formats easily)

### Project Structure
```
ceska_slova/
├── cmd/
│   └── ceska_slova/
│       └── main.go
│
├── internal/
│   ├── loader/        # YAML loading
│   ├── model/         # Data structures
│   ├── renderer/      # Output formats
│   └── filter/        # Filtering logic
│
├── data/              # Vocabulary files (YAML)
├── schema/            # JSON schema (optional validation)
├── templates/         # YAML templates for new lessons
└── go.mod
```

### 🧪 Usage

#### Run CLI output
`go run ./cmd/ceska_slova --format=cli`

#### Markdown export
`go run ./cmd/ceska_slova --format=markdown`

#### TTS mode (speech training)
`go run ./cmd/ceska_slova --format=tts`

Output format:
```
Muž čte knihu.
---
Muž čte knihu.
---
...
```

### 🏷️ Filtering

```
# Filter vocabulary by tags:

go run ./cmd/ceska_slova --format=cli --tag=food

# Multiple tags:

--tag=food,lesson-1

```

### 🧠 Architecture

This project is built around a simple pipeline:

```
YAML files
   ↓
Loader
   ↓
Filter (optional)
   ↓
Renderer (CLI / Markdown / TTS)
   ↓
Output
```

This makes it easy to extend with new formats or learning modes.

### 🔧 Adding New Output Formats

To add a new format:

1. Create a new file in:
   ```
   internal/renderer/
   ```

2. Implement:
   ```go
   func PrintXYZ(file *model.File)
   ```

3. Register it in main.go:
   ```go
   switch format {
   case "xyz":
     renderer.PrintXYZ(f)
   }
   ```

### 🚀 Roadmap

Planned features:

 - Anki export format
 - CLI subcommands (export markdown, tts, etc.)
 - Random learning mode
 - Spaced repetition support
 - Better schema validation (JSON Schema enforcement)
 - Interactive CLI mode

### 🧭 Philosophy

This tool is designed as:

A data-driven vocabulary engine, not just a vocabulary viewer.

Same data → many learning perspectives.

### 🛠️ Tech
Go 1.22+
YAML for data storage
Simple CLI architecture (no heavy dependencies)
