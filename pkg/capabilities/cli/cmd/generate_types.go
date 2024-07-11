package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atombender/go-jsonschema/pkg/generator"
	"github.com/atombender/go-jsonschema/pkg/schemas"
	"github.com/spf13/cobra"
)

var Dir string

// CapabilitySchemaFilePattern is used to extract the package name from the file path.
// This is used as the package name for the generated Go types.
var CapabilitySchemaFilePattern = regexp.MustCompile(`([^/]+)_(action|trigger|consensus|target)\.json$`)

// reg := regexp.MustCompile(`([^/]+)_(trigger|action)\.json$`)

func init() {
	generateTypesCmd.Flags().StringVar(&Dir, "dir", ".", fmt.Sprintf("Directory to search for %s files", CapabilitySchemaFilePattern.String()))
	if err := generateTypesCmd.MarkFlagDirname("dir"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(generateTypesCmd)
}

// Finds all files that match CapabilitySchemaFilePattern in the provided directory and generates Go
// types for each.
var generateTypesCmd = &cobra.Command{
	Use:   "generate-types",
	Short: "Generate Go types from JSON schema capability definitions",
	RunE: func(cmd *cobra.Command, args []string) error {

		dir := cmd.Flag("dir").Value.String()

		var schemaPaths []string

		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Ignore directories and files that don't match the CapabilitySchemaFileExtension
			if info.IsDir() || !CapabilitySchemaFilePattern.MatchString(path) {
				return nil
			}

			schemaPaths = append(schemaPaths, path)
			return nil
		}); err != nil {
			return errors.New(fmt.Sprintf("error walking the directory %v: %v\n", dir, err))
		}

		fmt.Println("Found", schemaPaths, "capability schema files")

		for _, schemaPath := range schemaPaths {
			file, content, err := TypesFromJSONSchema(schemaPath)

			if err != nil {
				return err
			}

			if err = os.WriteFile(file, content, 0644); err != nil {
				return err
			}

			fmt.Println("Generated types for", schemaPath)
		}

		return nil
	},
}

// TypesFromJSONSchema generates Go types from a JSON schema file.
func TypesFromJSONSchema(schemaFilePath string) (outputFilePath string, outputContents []byte, err error) {
	jsonSchema, err := schemas.FromJSONFile(schemaFilePath)
	if err != nil {
		return "", []byte(""), errors.New(fmt.Sprintf("error reading schema file %v:\n\t- %v\n\nTIP: This can happen if the supplied JSON schema is invalid. Try using https://jsonschemalint.com/#!/version/draft-07/markup/json to validate the schema.\n", schemaFilePath, err))
	}

	capabilityInfo := CapabilitySchemaFilePattern.FindStringSubmatch(schemaFilePath)
	packageName := capabilityInfo[1]
	capabilityType := capabilityInfo[2]
	outputName := strings.Replace(schemaFilePath, capabilityType+".json", capabilityType+"_generated.go", 1)
	rootType := capitalize(packageName) + capitalize(capabilityType)

	cfg := generator.Config{
		Warner: func(message string) { fmt.Printf("Warning: %s\n", message) },
		SchemaMappings: []generator.SchemaMapping{
			{
				SchemaID:    jsonSchema.ID,
				PackageName: packageName,
				RootType:    rootType,
				OutputName:  outputName,
			},
		},
	}

	gen, err := generator.New(cfg)
	if err != nil {
		return "", []byte(""), err
	}

	if err = gen.DoFile(schemaFilePath); err != nil {
		return "", []byte(""), err
	}

	generatedContents := gen.Sources()
	content := generatedContents[outputName]

	content = []byte(strings.Replace(string(content), "// Code generated by github.com/atombender/go-jsonschema", "// Code generated by pkg/capabilities/cli", 1))

	return outputName, content, nil
}

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}
