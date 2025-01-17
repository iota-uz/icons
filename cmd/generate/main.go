package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	baseDir   = "phosphor-icons/SVGs"
	outputDir = "./phosphor"
)

// Types
type IconProcessingError struct {
	msg string
}

func (e *IconProcessingError) Error() string {
	return e.msg
}

type SVG struct {
	XMLName xml.Name
	Content string     `xml:",innerxml"`
	Attrs   []xml.Attr `xml:",any,attr"`
}

type IconVariant map[string]string
type IconGroups map[string][]string

type ProcessingStats struct {
	totalFiles     int
	processedFiles int
	skippedFiles   int
	failedFiles    int
}

type IconProcessor struct {
	stats        *ProcessingStats
	iconVariants map[string]IconVariant
	iconGroups   IconGroups
}

// Icon variants mapping
var variants = map[string]string{
	"Filled":  "fill",
	"DuoTone": "duotone",
	"Bold":    "bold",
	"Thin":    "thin",
	"Regular": "regular",
	"Light":   "light",
}

// String manipulation functions
func toCamelCase(str string) string {
	str = regexp.MustCompile(`-+`).ReplaceAllString(strings.Trim(str, "-"), "-")
	components := strings.Split(str, "-")
	var result strings.Builder
	for _, comp := range components {
		if comp == "" {
			continue
		}
		result.WriteString(strings.Title(comp))
	}
	return result.String()
}

func cleanIconName(fileName string) string {
	suffixes := []string{"fill", "bold", "duotone", "thin", "regular", "light"}
	for _, suffix := range suffixes {
		if strings.HasSuffix(fileName, "-"+suffix) {
			return fileName[:len(fileName)-len(suffix)-1]
		}
	}
	return fileName
}

func getIconGroup(fileName string) string {
	parts := strings.Split(fileName, "-")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// SVG processing functions
func extractSVGContents(content string) (string, error) {
	var svg SVG
	if err := xml.Unmarshal([]byte(content), &svg); err != nil {
		return "", &IconProcessingError{msg: fmt.Sprintf("Failed to parse SVG content: %v", err)}
	}
	cleaned := regexp.MustCompile(`\sxmlns:[\w\d]+="[^"]*"`).ReplaceAllString(svg.Content, "")
	cleaned = regexp.MustCompile(`\sxmlns="[^"]*"`).ReplaceAllString(cleaned, "")
	return strings.TrimSpace(cleaned), nil
}

// Template generation functions
func generateTemplFunction(iconName string, variants IconVariant) (string, error) {
	if len(variants) == 0 {
		return "", &IconProcessingError{msg: fmt.Sprintf("No variants found for icon: %s", iconName)}
	}

	functionName := toCamelCase(iconName)
	defaultContent := variants["Regular"]
	if defaultContent == "" {
		for _, content := range variants {
			defaultContent = content
			break
		}
	}

	var variantCases strings.Builder
	first := true
	for variant, content := range variants {
		if variant == "Regular" {
			continue
		}
		if first {
			variantCases.WriteString(fmt.Sprintf("\nif props.Variant == %s {\n%s\n}", variant, content))
			first = false
		} else {
			variantCases.WriteString(fmt.Sprintf("\nelse if props.Variant == \"%s\" {\n%s\n}", variant, content))
		}
	}

	if variantCases.Len() > 0 {
		variantCases.WriteString(fmt.Sprintf("\nelse {\n%s\n}", defaultContent))
	} else {
		variantCases.WriteString(fmt.Sprintf("\n%s\n", defaultContent))
	}

	return fmt.Sprintf("templ %s(props Props) {@svg(props) {%s}}", functionName, variantCases.String()), nil
}

// File operations
func (p *IconProcessor) collectIconNames(variantPath string) error {
	files, err := filepath.Glob(filepath.Join(variantPath, "*.svg"))
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	p.stats.totalFiles += len(files)

	for _, file := range files {
		baseName := filepath.Base(file)
		iconName := cleanIconName(strings.TrimSuffix(baseName, ".svg"))
		if _, exists := p.iconVariants[iconName]; !exists {
			p.iconVariants[iconName] = make(IconVariant)
		}
	}

	return nil
}

func (p *IconProcessor) processSVGFiles(variant, variantPath string) error {
	files, err := filepath.Glob(filepath.Join(variantPath, "*.svg"))
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		baseName := filepath.Base(file)
		iconName := cleanIconName(strings.TrimSuffix(baseName, ".svg"))

		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			p.stats.failedFiles++
			continue
		}

		svgContent, err := extractSVGContents(string(content))
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", file, err)
			p.stats.failedFiles++
			continue
		}

		if svgContent == "" {
			fmt.Printf("Warning: No content found in %s\n", file)
			p.stats.skippedFiles++
			continue
		}

		p.iconVariants[iconName][variant] = svgContent
	}

	return nil
}

func (p *IconProcessor) generateIconGroups() {
	p.iconGroups = make(IconGroups)
	for iconName := range p.iconVariants {
		groupName := getIconGroup(iconName)
		p.iconGroups[groupName] = append(p.iconGroups[groupName], iconName)
	}
}

func (p *IconProcessor) generateListFile() error {
	listFile := filepath.Join(outputDir, "list.go")
	file, err := os.Create(listFile)
	if err != nil {
		return fmt.Errorf("failed to create list.go file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString("package phosphor\n\nimport \"github.com/a-h/templ\"\n\ntype Icon func(props Props) templ.Component\n\nvar Icons = map[string]Icon{\n"); err != nil {
		return fmt.Errorf("failed to write list.go header: %v", err)
	}

	for iconName := range p.iconVariants {
		functionName := toCamelCase(iconName)
		if _, err := file.WriteString(fmt.Sprintf("\t\"%s\": %s,\n", functionName, functionName)); err != nil {
			return fmt.Errorf("failed to write icon mapping: %v", err)
		}
	}

	if _, err := file.WriteString("}\n"); err != nil {
		return fmt.Errorf("failed to write list.go footer: %v", err)
	}

	return nil
}

func (p *IconProcessor) generateTemplFiles() error {
	for groupName, iconNames := range p.iconGroups {
		fileName := filepath.Join(outputDir, groupName+".templ")
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", fileName, err)
		}

		if _, err := file.WriteString("package phosphor\n\n"); err != nil {
			return fmt.Errorf("failed to write to file %s: %v", fileName, err)
		}

		for _, iconName := range iconNames {
			templFunction, err := generateTemplFunction(iconName, p.iconVariants[iconName])
			if err != nil {
				fmt.Printf("Error generating template for icon %s: %v\n", iconName, err)
				p.stats.failedFiles++
				continue
			}

			if _, err := file.WriteString(templFunction + "\n\n"); err != nil {
				return fmt.Errorf("failed to write to file %s: %v", fileName, err)
			}
			p.stats.processedFiles++
		}

		file.Close()
		fmt.Printf("Generated %s.templ with %d icons\n", groupName, len(iconNames))
	}

	return nil
}

func (p *IconProcessor) printSummary() {
	fmt.Printf("\nProcessing Summary:\n")
	fmt.Printf("Total SVG files found: %d\n", p.stats.totalFiles)
	fmt.Printf("Total groups generated: %d\n", len(p.iconGroups))
	fmt.Printf("Successfully processed icons: %d\n", p.stats.processedFiles)
	fmt.Printf("Skipped files: %d\n", p.stats.skippedFiles)
	fmt.Printf("Failed files: %d\n", p.stats.failedFiles)
	fmt.Printf("\nOutput directory: %s\n", outputDir)
}

// Main processing function
func processIcons() error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	processor := &IconProcessor{
		stats:        &ProcessingStats{},
		iconVariants: make(map[string]IconVariant),
	}

	// First pass: collect all icon names
	for _, folder := range variants {
		variantPath := filepath.Join(baseDir, folder)
		if _, err := os.Stat(variantPath); os.IsNotExist(err) {
			fmt.Printf("Skipping non-existent folder: %s\n", folder)
			continue
		}

		if err := processor.collectIconNames(variantPath); err != nil {
			return err
		}
	}

	// Generate list.go with all discovered icons
	if err := processor.generateListFile(); err != nil {
		return err
	}

	// Second pass: process SVG contents
	for variant, folder := range variants {
		variantPath := filepath.Join(baseDir, folder)
		if _, err := os.Stat(variantPath); os.IsNotExist(err) {
			continue
		}

		if err := processor.processSVGFiles(variant, variantPath); err != nil {
			return err
		}
	}

	// Generate icon groups and templ files
	processor.generateIconGroups()
	if err := processor.generateTemplFiles(); err != nil {
		return err
	}

	processor.printSummary()
	return nil
}

func main() {
	if err := processIcons(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
