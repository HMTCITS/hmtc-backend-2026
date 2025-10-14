package service

import (
	"context"
	"encoding/json"
	"fmt"

	"strings"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsService interface {
	AppendRow(spreadsheetID string, sheetName string, values []interface{}) error
}

type sheetsService struct{}

func NewSheetsService() SheetsService {
	return &sheetsService{}
}

func getSheetsClient() (*sheets.Service, error) {
	ctx := context.Background()

	refreshToken, err := repository.LoadRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("refresh token kosong: %v", err)
	}

	token := &oauth2.Token{RefreshToken: refreshToken}

	config := &oauth2.Config{
		ClientID:     config.AppConfig.OauthClientID,
		ClientSecret: config.AppConfig.OauthClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{sheets.SpreadsheetsScope},
	}

	ts := config.TokenSource(ctx, token)
	client := oauth2.NewClient(ctx, ts)

	// Buat Sheets service
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat Sheets client: %v", err)
	}

	// log.Printf("service: %+v\n", srv)

	return srv, nil
}

func ensureSheetExists(srv *sheets.Service, spreadsheetID, sheetName string) error {
	resp, err := srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return fmt.Errorf("gagal membaca spreadsheet: %v", err)
	}

	for _, s := range resp.Sheets {
		if s.Properties != nil && s.Properties.Title == sheetName {
			return nil // Sheet sudah ada
		}
	}

	// Sheet belum ada → buat baru
	reqs := []*sheets.Request{
		{
			AddSheet: &sheets.AddSheetRequest{
				Properties: &sheets.SheetProperties{
					Title: sheetName,
				},
			},
		},
	}

	_, err = srv.Spreadsheets.BatchUpdate(spreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: reqs,
	}).Do()
	if err != nil {
		return fmt.Errorf("gagal membuat sheet %s: %v", sheetName, err)
	}

	// log.Println("Sheet baru dibuat:", sheetName)
	return nil
}

func normalizeValue(v interface{}) interface{} {
	switch val := v.(type) {
	case nil:
		return ""
	case string, int, float64, bool:
		return val
	case []string:
		return strings.Join(val, ", ")
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

func (ss *sheetsService) AppendRow(spreadsheetID string, sheetName string, values []interface{}) error {
	srv, err := getSheetsClient()
	if err != nil {
		return err
	}

	if err := ensureSheetExists(srv, spreadsheetID, sheetName); err != nil {
		return err
	}

	// normalize
	nValues := make([]interface{}, len(values))
	for i, v := range values {
		nValues[i] = normalizeValue(v)
	}

	writeRange := fmt.Sprintf("%s!A:Z", sheetName)
	valRange := &sheets.ValueRange{Values: [][]interface{}{nValues}}

	_, err = srv.Spreadsheets.Values.
		Append(spreadsheetID, writeRange, valRange).
		ValueInputOption("USER_ENTERED").
		Do()

	if err != nil {
		return fmt.Errorf("gagal append data ke sheet: %v", err)
	}

	// log.Println("Data berhasil ditambahkan ke sheet:", sheetName)
	return nil
}
