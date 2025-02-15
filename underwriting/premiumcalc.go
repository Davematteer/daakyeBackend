package underwriting

func calcPremium(age int, 
				location string,
				smoker bool,
				occupation string, 
				body_mass_index int,
				planType string,
			){
				basePremium := 50
				switch {
				case age > 50:
					basePremium += 30
				}
				case smoker ==
			}