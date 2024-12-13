# Pathfinder 2e Spell Card Generator

This project is a Magic Spell Card Generator for creating visually appealing cards for tabletop RPGs. The generator dynamically creates cards with customized backgrounds, formatted text, and descriptive spell details, supporting both English and Portuguese languages.

## Features
	•	Dynamic backgrounds based on the spell’s tradition:
	•	Primal (Green Forest)
	•	Arcane (Mystical Blue Energy)
	•	Divine (Ethereal White and Gold)
	•	Occult (Cosmic Purple and Yellow)
	•	Customizable text boxes for:
	•	Spell Name, Level, and Actions.
	•	Keywords and Area.
	•	Spell Description (with automatic font resizing for large text).
	•	Multi-language support (English and Portuguese).
	•	Auto-saves cards in a structured directory: img/cards/.
	•	Support for PNG and JPEG images for both spell artwork and backgrounds.

## Example Card

Here’s an example of a card generated for the spell “Moonlight Ray”:

## Requirements
	•	Go programming language installed (version 1.18 or higher).
	•	Fonts:
	•	OpenSans-Bold.ttf
	•	OpenSans-Regular.ttf
	•	A directory structure as follows:

.
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



## How to Use
	1.	Clone the repository:

git clone https://github.com/yourusername/magic-card-generator.git
cd magic-card-generator


	2.	Place your spell artwork in the img/spells/ directory.
	3.	Place the background images for each tradition in the img/backgrounds/ directory.
	4.	Make sure the fonts/ directory contains the required font files.
	5.	Run the generator:

go run main.go


	6.	The generated cards will be saved in the img/cards/ directory.

## Customize Your Spell

You can customize the spell details by modifying the main function:

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

## Contributions

Feel free to fork this repository, create pull requests, or submit issues to improve the project!

You can replace the image path in the example (![Moonlight Ray Card](./img/cards/Moonlight_Ray.png)) with the actual card file path you generated to display the correct preview on your GitHub repository.