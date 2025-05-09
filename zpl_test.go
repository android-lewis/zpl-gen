package zpl

import (
	"os"
	"strings"
	"testing"
)

type LabelDetails struct {
	DistributorCode   string
	FontSize          string
	TitleLines        string
	ShortDescription  string
	CaseQty           string
	ProductSizeOption string
	UnitOfMeasure     string
	UBNCode           string
	EPSCode           string
	Outer1DBarcode    string
	LotNo             string
	ExpiryDate        string
}

func TestLabelGenerationFromFile(t *testing.T) {
	test_zpl := "labels/test/test.zpl"
	expected_zpl := "labels/test/match.zpl"
	unexpected_zpl := "labels/test/no_match.zpl"

	expected_content, err := os.ReadFile(expected_zpl)

	if err != nil {
		t.Error(err)
	}

	unexpected_content, err := os.ReadFile(unexpected_zpl)

	if err != nil {
		t.Error(err)
	}

	details := &LabelDetails{
		DistributorCode:   "UK/Ire",
		FontSize:          "80",
		TitleLines:        "1",
		ShortDescription:  "TEST LABEL",
		CaseQty:           "6",
		ProductSizeOption: "200",
		UnitOfMeasure:     "g",
		UBNCode:           "UBN123456",
		EPSCode:           "EPS123456",
		Outer1DBarcode:    "00112233445566",
		LotNo:             "123",
		ExpiryDate:        "12/12/24",
	}

	detailsMap := GenerateDetailMap(*details)
	generated_content, err := GenerateLabelFile(test_zpl, detailsMap)

	if err != nil {
		t.Error(err)
	}

	if _, _, ok := compareZPLStrings(generated_content, string(unexpected_content)); ok == true {
		t.Logf("generated label matches unexpected content: expected:\n %s \n got:\n %s \n", strings.TrimSpace(string(unexpected_content)), strings.TrimSpace(generated_content))
		t.Fail()
	}

	if exp_line, gen_line, ok := compareZPLStrings(generated_content, string(expected_content)); ok == false {
		t.Logf("generated label does not match expected content: expected:\n\n\t%s\n\ngot:\n\n\t%s\n\n", []byte(exp_line), []byte(gen_line))
		t.FailNow()
	}

}

func TestLabelGenerationFromString(t *testing.T) {
	test_zpl := "labels/test/test.zpl"
	expected_zpl := "labels/test/match.zpl"
	unexpected_zpl := "labels/test/no_match.zpl"

	expected_content, err := os.ReadFile(expected_zpl)

	if err != nil {
		t.Error(err)
	}

	unexpected_content, err := os.ReadFile(unexpected_zpl)

	if err != nil {
		t.Error(err)
	}

	test_file, err := os.ReadFile(test_zpl)

	if err != nil {
		t.Error(err)
	}

	details := &LabelDetails{
		DistributorCode:   "UK/Ire",
		FontSize:          "80",
		TitleLines:        "1",
		ShortDescription:  "TEST LABEL",
		CaseQty:           "6",
		ProductSizeOption: "200",
		UnitOfMeasure:     "g",
		UBNCode:           "UBN123456",
		EPSCode:           "EPS123456",
		Outer1DBarcode:    "00112233445566",
		LotNo:             "123",
		ExpiryDate:        "12/12/24",
	}

	detailsMap := GenerateDetailMap(*details)
	generated_content, err := GenerateLabelString(string(test_file), detailsMap)

	if err != nil {
		t.Error(err)
	}

	if _, _, ok := compareZPLStrings(generated_content, string(unexpected_content)); ok == true {
		t.Logf("generated label matches unexpected content: expected:\n %s \n got:\n %s \n", strings.TrimSpace(string(unexpected_content)), strings.TrimSpace(generated_content))
		t.Fail()
	}

	if exp_line, gen_line, ok := compareZPLStrings(generated_content, string(expected_content)); ok == false {
		t.Logf("generated label does not match expected content: expected:\n\n\t%s\n\ngot:\n\n\t%s\n\n", []byte(exp_line), []byte(gen_line))
		t.FailNow()
	}

}

func compareZPLStrings(gen, exp string) (string, string, bool) {
	gen_s := strings.TrimSpace(gen)
	exp_s := strings.TrimSpace(exp)
	gen_s = cleanString(gen_s)
	exp_s = cleanString(exp_s)
	gen_lines := strings.Split(gen_s, "^FS")
	exp_lines := strings.Split(exp_s, "^FS")

	if len(gen_lines) != len(exp_lines) {
		return "different lengths", "", false
	}

	for i := range gen_lines {
		if gen_lines[i] != exp_lines[i] {
			return exp_lines[i], gen_lines[i], false
		}
	}

	return "", "", true
}
