package wiimdev

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newCommandTestDevice(t *testing.T, body string) (*WiimDevice, func()) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(body))
	}))

	return &WiimDevice{
		url:    server.URL + "/?command=",
		client: server.Client(),
	}, server.Close
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want time.Duration
	}{
		{name: "valid", in: "01:02:03", want: time.Hour + 2*time.Minute + 3*time.Second},
		{name: "invalid", in: "bad", want: 0},
		{name: "negative", in: "-01:00:00", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDuration(tt.in); got != tt.want {
				t.Fatalf("parseDuration(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCommandDecodesPrimitiveResponses(t *testing.T) {
	tests := []struct {
		name string
		body string
		run  func(*WiimDevice) error
	}{
		{
			name: "string",
			body: "OK",
			run: func(device *WiimDevice) error {
				var got string
				if err := device.command("status", &got); err != nil {
					return err
				}
				if got != "OK" {
					t.Fatalf("string result = %q, want OK", got)
				}
				return nil
			},
		},
		{
			name: "int",
			body: "7\n",
			run: func(device *WiimDevice) error {
				var got int
				if err := device.command("level", &got); err != nil {
					return err
				}
				if got != 7 {
					t.Fatalf("int result = %d, want 7", got)
				}
				return nil
			},
		},
		{
			name: "float",
			body: "0.25\n",
			run: func(device *WiimDevice) error {
				var got float64
				if err := device.command("balance", &got); err != nil {
					return err
				}
				if got != 0.25 {
					t.Fatalf("float result = %f, want 0.25", got)
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device, closeServer := newCommandTestDevice(t, tt.body)
			defer closeServer()

			if err := tt.run(device); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestCommandDecodesJSONIntoTarget(t *testing.T) {
	device, closeServer := newCommandTestDevice(t, `{"status":"OK","Name":"Rock"}`)
	defer closeServer()

	var got struct {
		Status string `json:"status"`
		Name   string `json:"Name"`
	}
	if err := device.command("EQGetBand", &got); err != nil {
		t.Fatal(err)
	}
	if got.Status != "OK" || got.Name != "Rock" {
		t.Fatalf("JSON result = %+v, want status OK and name Rock", got)
	}
}

func TestAudioInputOptionsShareDefinitions(t *testing.T) {
	if got := audioInputIndex("40"); got != 2 {
		t.Fatalf("audioInputIndex(40) = %d, want 2", got)
	}
	if got := audioInputIndex("missing"); got != -1 {
		t.Fatalf("audioInputIndex(missing) = %d, want -1", got)
	}

	device, closeServer := newCommandTestDevice(t, "OK")
	defer closeServer()

	if err := device.SetAudioInput(2); err != nil {
		t.Fatal(err)
	}
	if err := device.SetAudioInput(-1); err == nil {
		t.Fatal("SetAudioInput(-1) succeeded, want error")
	}

	got, err := device.GetAudioInputList()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"Network", "Bluetooth", "Line In", "Optical In"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("GetAudioInputList()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
