package appErrors

import (
	"reflect"
	"testing"

	"github.com/joomcode/errorx"
)

func TestExtractAppError(t *testing.T) {
	// 1) Custom app error without predefined properties → defaults should apply
	t.Run("default values when no properties set", func(t *testing.T) {
		err := appErrorType.New("Oops") // no HttpCode/Code properties

		out, httpCode := ExtractAppError(err)

		if httpCode != 500 {
			t.Fatalf("expected default httpCode 500, got %d", httpCode)
		}
		if out.Message != "Oops" {
			t.Fatalf("expected message 'Oops', got %q", out.Message)
		}
		if out.Edited {
			t.Fatalf("expected Edited=false by default")
		}
		if out.Code != "" {
			t.Fatalf("expected empty Code by default, got %q", out.Code)
		}
		if out.Error != nil {
			t.Fatalf("expected payload nil by default, got %#v", out.Error)
		}
	})

	// 2) Using predefined InternalError which carries http=500 and code="internal_error"
	t.Run("predefined internal error properties extracted", func(t *testing.T) {
		err := InternalError

		out, httpCode := ExtractAppError(err)

		if httpCode != 500 {
			t.Fatalf("expected httpCode 500, got %d", httpCode)
		}
		if out.Message != "Internal Error" {
			t.Fatalf("unexpected message: %q", out.Message)
		}
		if out.Code != "internal_error" {
			t.Fatalf("expected code 'internal_error', got %q", out.Code)
		}
		if out.Edited {
			t.Fatalf("expected Edited=false, got true")
		}
		if out.Error != nil {
			t.Fatalf("expected payload nil, got %#v", out.Error)
		}
	})

	// 3) NotFoundError with extra properties: Edited=true + payload + override http code
	t.Run("edited and payload and overridden http code are extracted", func(t *testing.T) {
		payload := map[string]any{"detail": "resource missing", "id": 42}
		err := NotFoundError.
			WithProperty(EditedProperty, true).
			WithProperty(errorx.PropertyPayload(), payload).
			WithProperty(HttpCodeProperty, 418) // override

		out, httpCode := ExtractAppError(err)

		if httpCode != 418 {
			t.Fatalf("expected overridden httpCode 418, got %d", httpCode)
		}
		if out.Message != "Not Found" {
			t.Fatalf("unexpected message: %q", out.Message)
		}
		if out.Code != "not_found" {
			t.Fatalf("expected code 'not_found', got %q", out.Code)
		}
		if !out.Edited {
			t.Fatalf("expected Edited=true, got false")
		}
		if !reflect.DeepEqual(out.Error, payload) {
			t.Fatalf("payload mismatch: got %#v", out.Error)
		}
	})
}
