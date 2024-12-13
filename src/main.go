package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // Adicionado para suportar JPEG
	_ "image/png"  // Adicionado para suportar PNG
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

type Magic struct {
	Name        string
	Level       string
	Actions     int
	Image       string
	Keywords    []string
	Tradition   string
	Area        string
	Description string
	Language    string
}

func loadImage(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a imagem: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file) // Use o decodificador correto automaticamente
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar a imagem: %v", err)
	}
	return img, nil
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

func resizeImageToFill(img image.Image, width, height int) image.Image {
	dc := gg.NewContext(width, height)
	dc.DrawImageAnchored(img, width/2, height/2, 0.5, 0.5)
	return dc.Image()
}

func drawFullBackground(dc *gg.Context, img image.Image, canvasWidth, canvasHeight int) {
	resizedImg := resizeImageToFill(img, canvasWidth, canvasHeight)
	dc.DrawImage(resizedImg, 0, 0)
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
	}
	if langMap, ok := translations[field]; ok {
		if translation, ok := langMap[language]; ok {
			return translation
		}
	}
	return field
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

func drawInfoBox(dc *gg.Context, magic Magic, x, y, width, height float64, fontBoldPath, fontRegularPath string) {
	drawRoundedBox(dc, x, y, width, height, color.RGBA{245, 245, 220, 255}, color.RGBA{0, 0, 0, 200}, 2.0)

	name := magic.Name
	levelActions := fmt.Sprintf("%s: %s | %d %s",
		translateField("Level", magic.Language),
		magic.Level,
		magic.Actions,
		translateField("Actions", magic.Language),
	)

	if err := dc.LoadFontFace(fontBoldPath, 60); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontBoldPath, err)
	}
	_, nameHeight := dc.MeasureString(name)

	if err := dc.LoadFontFace(fontRegularPath, 24); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontRegularPath, err)
	}
	_, levelActionsHeight := dc.MeasureString(levelActions)

	totalTextHeight := nameHeight + 35.0 + levelActionsHeight
	startingY := y + (height-totalTextHeight)/2

	drawTextWithShadow(dc, name, fontBoldPath, 60, x+width/2, startingY, width-40, color.Black, color.RGBA{0, 0, 0, 150}, 1.5, gg.AlignCenter)
	drawTextWithShadow(dc, levelActions, fontRegularPath, 24, x+width/2, startingY+nameHeight+35.0, width-40, color.Black, color.RGBA{0, 0, 0, 150}, 1.3, gg.AlignCenter)
}

func drawDetailsBox(dc *gg.Context, magic Magic, x, y, width, height float64, fontBoldPath, fontRegularPath string) {
	drawRoundedBox(dc, x, y, width, height, color.RGBA{245, 245, 220, 255}, color.RGBA{0, 0, 0, 200}, 2.0)

	// Labels traduzidos
	keywordsLabel := translateField("Keywords", magic.Language) + ": "
	areaLabel := translateField("Area", magic.Language) + ": "

	// Conteúdo das keywords e área
	keywordsContent := strings.Join(magic.Keywords, ", ")
	areaContent := magic.Area

	// Carregar fontes para os textos regulares
	if err := dc.LoadFontFace(fontRegularPath, 20); err != nil {
		log.Fatalf("Falha ao carregar a fonte %s: %v", fontRegularPath, err)
	}

	// Medir largura dos textos
	keywordsWidth, _ := dc.MeasureString(keywordsLabel + keywordsContent)
	areaWidth, _ := dc.MeasureString(areaLabel + areaContent)

	// Definir posições horizontais para centralizar
	keywordsX := x + (width-keywordsWidth)/2
	areaX := x + (width-areaWidth)/2
	currentY := y + 20.0 // Posição inicial vertical com padding

	// Desenhar keywords
	dc.SetColor(color.Black)
	dc.DrawString(keywordsLabel+keywordsContent, keywordsX, currentY)

	// Desenhar área
	currentY += 30.0 // Espaçamento entre as linhas
	dc.DrawString(areaLabel+areaContent, areaX, currentY)
}

func drawDescriptionBox(dc *gg.Context, description string, x, y, width, height float64, fontPath string) {
	bgColor := color.RGBA{245, 245, 220, 255}
	borderColor := color.RGBA{0, 0, 0, 200}
	borderWidth := 2.0
	drawRoundedBox(dc, x, y, width, height, bgColor, borderColor, borderWidth)

	paddingX := 20.0
	paddingY := 30.0
	textWidth := width - 2*paddingX
	startX := x + paddingX
	startY := y + paddingY

	fontSize := 30.0

	for {
		if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
			log.Fatalf("Falha ao carregar a fonte %s: %v", fontPath, err)
		}

		wrappedText := dc.WordWrap(description, textWidth)
		totalHeight := float64(len(wrappedText)) * (dc.FontHeight() + 1.5)

		if totalHeight <= height-paddingY {
			break
		}
		fontSize -= 1.0 // Diminui o tamanho da fonte
		if fontSize < 10.0 {
			log.Println("Texto muito longo para a caixa, mesmo com fonte mínima.")
			break
		}
	}

	dc.SetColor(color.Black)
	dc.DrawStringWrapped(description, startX, startY, 0, 0, textWidth, 1.5, gg.AlignLeft)
}

func createMagicCard(magic Magic, outputFile string) {
	const canvasWidth int = 1080
	const canvasHeight int = 1920
	dc := gg.NewContext(canvasWidth, canvasHeight)

	// Caminho da imagem de fundo
	backgroundPath := fmt.Sprintf("./img/backgrounds/%s.png", strings.ToLower(magic.Tradition))
	backgroundImage, err := loadImage(backgroundPath)
	if err != nil {
		log.Fatalf("Erro ao carregar a imagem de fundo '%s': %v", backgroundPath, err)
	}
	drawFullBackground(dc, backgroundImage, canvasWidth, canvasHeight)

	// Fontes
	fontBoldPath := "fonts/OpenSans-Bold.ttf"
	fontRegularPath := "fonts/OpenSans-Regular.ttf"

	// Caixa de Informações (Nome, Level e Actions)
	infoBoxX := 50.0
	infoBoxY := 50.0
	infoBoxWidth := float64(canvasWidth) - 2*infoBoxX
	infoBoxHeight := 200.0
	drawInfoBox(dc, magic, infoBoxX, infoBoxY, infoBoxWidth, infoBoxHeight, fontBoldPath, fontRegularPath)

	// Imagem da Magia
	img, err := loadImage(magic.Image)
	if err != nil {
		log.Fatalf("Falha ao carregar a imagem da magia: %v", err)
	}
	imageWidth := 800
	imageHeight := 800
	imageX := (canvasWidth - imageWidth) / 2
	imageY := int(infoBoxY + infoBoxHeight + 30)
	drawRoundedBox(dc, float64(imageX), float64(imageY), float64(imageWidth), float64(imageHeight), color.Transparent, color.RGBA{245, 245, 220, 255}, 10.0)
	dc.DrawImage(resizeImageToFill(img, imageWidth, imageHeight), imageX, imageY)

	// Caixa de Detalhes (Keywords e Área)
	detailsBoxX := 50.0
	detailsBoxY := float64(imageY + imageHeight + 30)
	detailsBoxWidth := float64(canvasWidth) - 2*detailsBoxX
	detailsBoxHeight := 60.0
	drawDetailsBox(dc, magic, detailsBoxX, detailsBoxY, detailsBoxWidth, detailsBoxHeight, fontBoldPath, fontRegularPath)

	// Caixa de Descrição
	descriptionBoxX := 50.0
	descriptionBoxY := detailsBoxY + detailsBoxHeight + 30
	descriptionBoxWidth := float64(canvasWidth) - 2*descriptionBoxX
	descriptionBoxHeight := float64(canvasHeight) - descriptionBoxY - 60.0
	drawDescriptionBox(dc, magic.Description, descriptionBoxX, descriptionBoxY, descriptionBoxWidth, descriptionBoxHeight, fontRegularPath)

	// Caminho do arquivo de saída na pasta "img/cards"
	outputFilePath := fmt.Sprintf("./img/cards/%s.png", outputFile)

	// Salvar o card
	err = dc.SavePNG(outputFilePath)
	if err != nil {
		log.Fatalf("Falha ao salvar o card: %v", err)
	}

	fmt.Printf("Card salvo em %s\n", outputFilePath)
}

func main() {
	magic := Magic{
		Name:        "Raio de Luz Lunar",
		Level:       "3",
		Actions:     2,
		Image:       "./img/spells/moonlight_ray.png",
		Keywords:    []string{"Ataque", "Frio", "Evocação", "Bom", "Luz"},
		Tradition:   "Primal",
		Area:        "Alcance de 36 metros; 1 criatura",
		Description: "Você libera um raio sagrado de luz congelante da lua. Faça um ataque de magia à distância. O raio causa 5d6 de dano frio; se o alvo for um demônio ou morto-vivo, você causa 5d6 de dano bom adicional. O dano frio do raio de luz lunar é considerado dano de prata para fins de resistências, fraquezas e similares. \n\nSucesso Crítico: O alvo sofre o dobro do dano frio e também o dobro do dano bom, se for um demônio ou morto-vivo. \nSucesso: O alvo sofre o dano total. \n\nSe a luz atravessar uma área de escuridão mágica ou atingir uma criatura afetada por escuridão mágica, o raio de luz lunar tenta neutralizar a escuridão. Para determinar se a luz passa por uma área de escuridão, trace uma linha entre você e o alvo da magia. \n\n\n\nAprimorada (+1): O dano frio aumenta em 2d6, e o dano bom contra demônios e mortos-vivos aumenta em 2d6.",
		Language:    "pt",
	}

	createMagicCard(magic, magic.Name+".png")
}
