package main

import "fmt"
import "conekta.io/api/swagger"

func main() {
    


    address := swagger.BillingAddress{
      Street1:"77 Mystery Lane",
      Street2: "Suite 124",
      Street3: "",
      City: "Darlington",
      State:"NJ",
      Zip: "10192",
      Country: "Mexico",
      TaxId: "xmn671212drx",
      CompanyName:"X-Men Inc.",
      Phone: "77-777-7777",
      Email: "purshasing@x-men.org",
    }

    customer := swagger.Customer{
      LoggedIn: "true",
      SuccessfulPurchases: 14,
      CreatedAt: 1379784950,
      UpdatedAt: 1379784950,
      OfflinePayments: 4,
      Score: 9,
    }

    lineItem := swagger.LineItem{
      Name: "Box of Cohiba S1s",
      Description: "Imported From Mex.",
      UnitPrice: 20000,
      Quantity: 1,
      Sku: "cohb_s1",
      Category: "food",

    }

    details := swagger.Details{
    	Name: "Arnulfo Quimare",
    	Phone: "403-342-0642",
    	Email: "logan@x-men.org",
    	Customer: customer,    	
    	LineItems: []swagger.LineItem{ lineItem },
    	BillingAddress: address,
    }

    charge := swagger.Charge{     	
  		Amount:  20000,
  		Currency: "MXN",
  		ReferenceId: "9839-wolf_pack",
  		Card:  "tok_test_visa_4242",
  		Description: "Test",
  		Details: details,
    }




    defaultApi  := swagger.NewDefaultApi("key_eYvWV7gSDkNYXsmr")
	resp, err := defaultApi.ChargesPost(charge)

	if err != nil {
		fmt.Printf("error")				
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", resp)

    fmt.Printf("done")		
}
