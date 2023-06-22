package players

import "testing"

func TestTaxiDriver(t *testing.T) {
	driver := TaxiDriver{
		ID: 1,
	}

	t.Log("Given the need to test TaxiDriver's behavior at different time.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen working in the morning.", testID)
		{
			...
		}

		testID++
		t.Logf("\tTest %d:\tWhen working in the evening.", testID)
		{
			...
		}
	}
}