package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

type Magic struct {
	Name            string
	Level           string
	Actions         int
	Image           string
	Keywords        []string
	Tradition       string
	School          string
	Area            string
	Duration        string
	Components      []string
	Cost            string
	SavingThrow     string
	CriticalSuccess string
	CriticalFailure string
	Description     string
	Language        string
}

func loadImage(imagePath string) (image.Image, int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("erro ao abrir a imagem: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("erro ao decodificar a imagem: %v", err)
	}
	return img, img.Bounds().Dx(), img.Bounds().Dy(), nil
}

func drawRoundedBox(dc *gg.Context, x, y, width, height float64, bgColor, borderColor color.Color, borderWidth float64) {
	radius := 20.0
	dc.SetColor(bgColor)
	dc.DrawRoundedRectangle(x, y, width, height, radius)
	dc.Fill()

	dc.SetLineWidth(borderWidth)
	dc.SetColor(borderColor)
	dc.DrawRoundedRectangle(x, y, width, height, radius)
	dc.Stroke()
}

func translateField(field, language string) string {
	translations := map[string]map[string]string{
		"Level": {
			"en": "Level",
			"pt": "Nível",
		},
		"Actions": {
			"en": "Action(s)",
			"pt": "Ação(ões)",
		},
		"Keywords": {
			"en": "Keywords",
			"pt": "Palavras-chave",
		},
		"Area": {
			"en": "Area",
			"pt": "Área",
		},
		"Duration": {
			"en": "Duration",
			"pt": "Duração",
		},
		"Components": {
			"en": "Components",
			"pt": "Componentes",
		},
		"Cost": {
			"en": "Cost",
			"pt": "Custo",
		},
		"Critical Success": {
			"en": "Critical Success",
			"pt": "Sucesso Crítico",
		},
		"Critical Failure": {
			"en": "Critical Failure",
			"pt": "Falha Crítica",
		},
	}
	if langMap, ok := translations[field]; ok {
		if translation, ok := langMap[language]; ok {
			return translation
		}
	}
	return field
}

func resizeImageToFill(img image.Image, width, height int) image.Image {
	dc := gg.NewContext(width, height)
	dc.DrawImageAnchored(img, width/2, height/2, 0.5, 0.5)
	return dc.Image()
}

func drawTextWithShadow(dc *gg.Context, text, fontPath string, fontSize, x, y, maxWidth float64, textColor, shadowColor color.Color, lineSpacing float64, align gg.Align) {
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontPath, err)
	}
	dc.SetColor(shadowColor)
	dc.DrawStringWrapped(text, x+2, y+2, 0.5, 0.0, maxWidth, lineSpacing, align)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(text, x, y, 0.5, 0.0, maxWidth, lineSpacing, align)
}

func createMagicCard(magic Magic, outputFile string) {
	backgroundPath := fmt.Sprintf("./img/backgrounds/%s.png", strings.ToLower(magic.Tradition))
	backgroundImage, bgWidth, bgHeight, err := loadImage(backgroundPath)
	if err != nil {
		log.Fatalf("Erro ao carregar a imagem de fundo '%s': %v", backgroundPath, err)
	}
	dc := gg.NewContext(bgWidth, bgHeight)
	dc.DrawImage(backgroundImage, 0, 0)

	// Proporções fixas para cada seção
	padding := 20.0
	infoBoxHeight := float64(bgHeight) * 0.15
	imageBoxHeight := float64(bgHeight) * 0.40
	detailsBoxHeight := float64(bgHeight) * 0.15
	descriptionBoxHeight := float64(bgHeight) - (infoBoxHeight + imageBoxHeight + detailsBoxHeight + 5*padding)

	fontBoldPath := "fonts/OpenSans-Bold.ttf"
	fontRegularPath := "fonts/OpenSans-Regular.ttf"

	// Caixa de informações (15%)
	drawRoundedBox(dc, padding, padding, float64(bgWidth)-2*padding, infoBoxHeight-padding, color.RGBA{245, 245, 220, 255}, color.RGBA{0, 0, 0, 200}, 2.0)
	name := magic.Name
	levelActions := fmt.Sprintf("%s: %s | %d %s",
		translateField("Level", magic.Language),
		magic.Level,
		magic.Actions,
		translateField("Actions", magic.Language),
	)
	if err := dc.LoadFontFace(fontBoldPath, 70); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontBoldPath, err)
	}
	dc.SetColor(color.Black)
	dc.DrawStringAnchored(name, float64(bgWidth)/2, padding+infoBoxHeight/3, 0.5, 0.5)
	if err := dc.LoadFontFace(fontRegularPath, 32); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontRegularPath, err)
	}
	dc.DrawStringAnchored(levelActions, float64(bgWidth)/2, padding+infoBoxHeight/1.5, 0.5, 0.5)

	// Caixa de imagem (40%)
	imageY := padding + infoBoxHeight + padding
	img, _, _, err := loadImage(magic.Image)
	if err != nil {
		log.Fatalf("Falha ao carregar a imagem da magia: %v", err)
	}
	dc.DrawRoundedRectangle(padding, imageY, float64(bgWidth)-2*padding, imageBoxHeight-padding, 20)
	dc.Clip()
	resizedImage := resizeImageToFill(img, bgWidth-int(2*padding), int(imageBoxHeight-padding))
	dc.DrawImage(resizedImage, int(padding), int(imageY))
	dc.ResetClip()

	// Caixa de detalhes (15%)
	detailsY := imageY + imageBoxHeight + padding
	drawRoundedBox(dc, padding, detailsY, float64(bgWidth)-2*padding, detailsBoxHeight-padding, color.RGBA{245, 245, 220, 255}, color.RGBA{0, 0, 0, 200}, 2.0)
	if err := dc.LoadFontFace(fontBoldPath, 28); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontBoldPath, err)
	}
	boldFont, _ := gg.LoadFontFace(fontBoldPath, 28)
	regularFont, _ := gg.LoadFontFace(fontRegularPath, 26)
	fields := []string{
		fmt.Sprintf("%s: %s", translateField("Keywords", magic.Language), strings.Join(magic.Keywords, ", ")),
		fmt.Sprintf("%s: %s", translateField("Components", magic.Language), strings.Join(magic.Components, ", ")),
		fmt.Sprintf("%s: %s", translateField("Area", magic.Language), magic.Area),
		fmt.Sprintf("%s: %s", translateField("Duration", magic.Language), magic.Duration),
	}
	lineHeight := 45.0 // Aumentando o espaçamento vertical
	currentY := detailsY + 40 // Adicionado padding entre o topo do box e o conteúdo
	for _, field := range fields {
		parts := strings.SplitN(field, ":", 2)
		dc.SetFontFace(boldFont)
		dc.DrawStringAnchored(parts[0]+":", padding+10, currentY, 0, 0)
		dc.SetFontFace(regularFont)
		dc.DrawStringAnchored(parts[1], padding+200, currentY, 0, 0) // Padronizado espaçamento horizontal
		currentY += lineHeight
	}

	// Caixa de descrição (restante)
	descriptionY := detailsY + detailsBoxHeight + padding
	drawRoundedBox(dc, padding, descriptionY, float64(bgWidth)-2*padding, descriptionBoxHeight+padding, color.RGBA{245, 245, 220, 255}, color.RGBA{0, 0, 0, 200}, 2.0)
	if err := dc.LoadFontFace(fontRegularPath, 24); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontRegularPath, err)
	}
	descriptionText := fmt.Sprintf("%s\n\n%s: %s\n%s: %s\n\n%s: %s",
		magic.Description,
		translateField("Critical Success", magic.Language), magic.CriticalSuccess,
		translateField("Critical Failure", magic.Language), magic.CriticalFailure,
		translateField("Saving Throw", magic.Language), magic.SavingThrow,
	)
	dc.DrawStringWrapped(descriptionText, padding+10, descriptionY+30, 0, 0, float64(bgWidth)-2*padding-20, 1.8, gg.AlignLeft)

	// Salvar o card
	outputFilePath := fmt.Sprintf("./img/cards/%s.png", outputFile)
	err = dc.SavePNG(outputFilePath)
	if err != nil {
		log.Fatalf("Falha ao salvar o card: %v", err)
	}
	fmt.Printf("Card salvo em %s\n", outputFilePath)
}

func main() {
	magic := Magic{
		Name:            "Raio de Luz Lunar",
		Level:           "3",
		Actions:         2,
		Image:           "./img/spells/moonlight_ray.png",
		Keywords:        []string{"Ataque", "Frio", "Evocação", "Bom", "Luz"},
		Tradition:       "Primal",
		School:          "Evocação",
		Area:            "Alcance de 36 metros; 1 criatura",
		Duration:        "1 rodada",
		Components:      []string{"Verbal", "Somático"},
		Cost:            "1 ponto de foco",
		SavingThrow:     "Fortitude",
		CriticalSuccess: "O alvo sofre o dobro do dano.",
		CriticalFailure: "O alvo sofre o dano completo e fica atordoado por 1 rodada.",
		Description:     "Você libera um raio sagrado de luz congelante da lua...",
		Language:        "pt",
	}

	createMagicCard(magic, magic.Name+".png")
}
