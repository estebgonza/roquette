package roqclient

import "testing"

func TestDriverExist(t *testing.T) {
	type driver struct {
		name     string
		expected string
	}
	drivers := []driver{
		driver{
			name:     "hive",
			expected: "",
		},
		driver{
			name:     "postgres",
			expected: "",
		},
		driver{
			name:     "mysql",
			expected: "type does'nt exist or not supported",
		},
		driver{
			name:     "oracle",
			expected: "type does'nt exist or not supported",
		},
		driver{
			name:     "",
			expected: "driver name not specified",
		},
	}

	for _, driver := range drivers {
		_, err := getClientByType(driver.name)
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		if driver.expected != errorMessage {
			t.Errorf("expected: %s, actual: %s", driver.expected, errorMessage)
		}
	}
}
