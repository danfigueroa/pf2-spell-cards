# Pathfinder 2e Spell Card Generator

This project is a **Magic Spell Card Generator** designed to create visually appealing cards for tabletop RPGs. The generator dynamically produces cards with customized backgrounds, formatted text, and descriptive spell details. It supports both English and Portuguese languages.

## Features

- Dynamic backgrounds based on the spell’s tradition:
  - **Primal:** Green Forest
  - **Arcane:** Mystical Blue Energy
  - **Divine:** Ethereal White and Gold
  - **Occult:** Cosmic Purple and Yellow
- Customizable text boxes for:
  - **Spell Name**, **Level**, and **Actions**
  - **Keywords** and **Area**
  - **Spell Description** (with automatic font resizing for large text)
- Multi-language support (**English** and **Portuguese**)
- Auto-saves cards in a structured directory: `img/cards/`
- Support for PNG and JPEG images for both spell artwork and backgrounds

## Example Card

Here’s an example of a card generated for the spell “Moonlight Ray”:

![Moonlight Ray Card](./src/img/cards/Moonlight_Ray.png)

## Requirements

- **Go programming language** installed (version 1.18 or higher)
- **Fonts:**
  - `OpenSans-Bold.ttf`
  - `OpenSans-Regular.ttf`
- A directory structure as follows:

```
├── fonts/
│   ├── OpenSans-Bold.ttf
│   ├── OpenSans-Regular.ttf
├── img/
│   ├── backgrounds/
│   │   ├── primal.png
│   │   ├── arcane.png
│   │   ├── divine.png
│   │   ├── occult.png
│   ├── spells/
│       ├── moonlight_ray.png
    ├── cards/
│       ├── moonlight_ray.png
```

## How to Use

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/magic-card-generator.git
   cd magic-card-generator

### Add Your Assets

1. **Place your spell artwork** in the `img/spells/` directory.  
2. **Place background images** for each tradition in the `img/backgrounds/` directory.  
3. **Ensure the `fonts/` directory** contains the required font files.

### Run the Generator

Run the generator with the following command:

```bash
go run main.go
```

### Find Your Generated Cards

The generated cards will be saved in the `img/cards/` directory.

---

### Customize Your Spell

To customize the spell details, edit the `main` function in `main.go`:

```go
magic := Magic{
    Name:        "Moonlight Ray",
    Level:       "3",
    Actions:     2,
    Image:       "./img/spells/moonlight_ray.png",
    Keywords:    []string{"Attack", "Cold", "Evocation", "Good", "Light"},
    Tradition:   "Primal",
    Area:        "Range 120 ft; 1 creature",
    Description: "You release a sacred ray of freezing moonlight. Make a ranged spell attack. The ray deals 5d6 cold damage; if the target is a demon or undead, it takes an additional 5d6 good damage...",
    Language:    "en",
}
createMagicCard(magic, magic.Name)
```

---

### Contributions

Contributions are welcome! Feel free to:

- Fork this repository  
- Create pull requests  
- Submit issues to suggest improvements  

To showcase your card previews, replace the example image path below with the actual card file path you generated:

```markdown
![Moonlight Ray Card](./src/img/cards/Moonlight_Ray.png)