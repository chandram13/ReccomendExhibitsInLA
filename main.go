package ReccomendExhibitsInLA

import "fmt"

func exhibitionRatings();
	exhibit1rating = []
	exhibit2rating = []
	exhibit3rating = []
	exhibit4rating = []
	exhibit5rating = []
	exhibit6rating = []

func LACMA();
	parkingAccessibility = input()
	enoughinterestingDisplays = input()
	relevant = input()
	creativity = input()
	shopping = input()
	freeibies = input()

	if parkingAccessibility == "good":
	exhibit1rating += 1
	if parkingAccessibility == "bad":
	exhibit1rating += -1
	if enoughinterestingDisplays == true:
	exhibit1rating += 1
	if enoughinterestingDisplays == false:
	exhibit1rating += -1
	if relevant == true:
	exhibit1rating += 1
	if relevant == false:
	exhibit1rating += -1
	if creativity == "interesting":
	exhibit1rating += 1
	if creativity == "not interesting":
	exhibit1rating += -1
	if shopping == "present":
	exhibit1rating += 1
	if shopping == "not present":
	exhibit1rating += -1
	if freebies == true:
	exhibit1rating += 1
	if freebies == false:
	exhibit1rating += -1


func CaliforniaScienceCenter();
	parkingAccessibility = input()
	enoughinterestingDisplays = input()
	relevant = input()
	creativity = input()
	shopping = input()
	freeibies = input()

	if parkingAccessibility == "good":
	exhibit2rating += 1
	if parkingAccessibility == "bad":
	exhibit2rating += -1
	if enoughinterestingDisplays == true:
	exhibit2rating += 1
	if enoughinterestingDisplays == false:
	exhibit2rating += -1
	if relevant == true:
	exhibit2rating += 1
	if relevant == false:
	exhibit2rating += -1
	if creativity == "interesting":
	exhibit2rating += 1
	if creativity == "not interesting":
	exhibit2rating += -1
	if shopping == "present":
	exhibit2rating += 1
	if shopping == "not present":
	exhibit2rating += -1
	if freebies == true:
	exhibit2rating += 1
	if freebies == false:
	exhibit2rating += -1

func JPaulGetty();
	parkingAccessibility = input()
	enoughinterestingDisplays = input()
	relevant = input()
	creativity = input()
	shopping = input()
	freeibies = input()

	if parkingAccessibility == "good":
	exhibit3rating += 1
	if parkingAccessibility == "bad":
	exhibit3rating += -1
	if enoughinterestingDisplays == true:
	exhibit3rating += 1
	if enoughinterestingDisplays == false:
	exhibit3rating += -1
	if relevant == true:
	exhibit3rating += 1
	if relevant == false:
	exhibit3rating += -1
	if creativity == "interesting":
	exhibit3rating += 1
	if creativity == "not interesting":
	exhibit3rating += -1
	if shopping == "present":
	exhibit3rating += 1
	if shopping == "not present":
	exhibit3rating += -1
	if freebies == true:
	exhibit3rating += 1
	if freebies == false:
	exhibit3rating += -1	

func NaturalHistoryMuseum();
	parkingAccessibility = input()
	enoughinterestingDisplays = input()
	relevant = input()
	creativity = input()
	shopping = input()
	freeibies = input()

	if parkingAccessibility == "good":
	exhibit4rating += 1
	if parkingAccessibility == "bad":
	exhibit4rating += -1
	if enoughinterestingDisplays == true:
	exhibit4rating += 1
	if enoughinterestingDisplays == false:
	exhibit4rating += -1
	if relevant == true:
	exhibit4rating += 1
	if relevant == false:
	exhibit4rating += -1
	if creativity == "interesting":
	exhibit4rating += 1
	if creativity == "not interesting":
	exhibit4rating += -1
	if shopping == "present":
	exhibit4rating += 1
	if shopping == "not present":
	exhibit4rating += -1
	if freebies == true:
	exhibit4rating += 1
	if freebies == false:
	exhibit4rating += -1	

func HammerMuseum();
	parkingAccessibility = input()
	enoughinterestingDisplays = input()
	relevant = input()
	creativity = input()
	shopping = input()
	freeibies = input()

	if parkingAccessibility == "good":
	exhibit5rating += 1
	if parkingAccessibility == "bad":
	exhibit5rating += -1
	if enoughinterestingDisplays == true:
	exhibit5rating += 1
	if enoughinterestingDisplays == false:
	exhibit5rating += -1
	if relevant == true:
	exhibit5rating += 1
	if relevant == false:
	exhibit5rating += -1
	if creativity == "interesting":
	exhibit5rating += 1
	if creativity == "not interesting":
	exhibit5rating += -1
	if shopping == "present":
	exhibit5rating += 1
	if shopping == "not present":
	exhibit5rating += -1
	if freebies == true:
	exhibit5rating += 1
	if freebies == false:
	exhibit5rating += -1	

func GriffithObservatory();
	parkingAccessibility = input()
	enoughinterestingDisplays = input()
	relevant = input()
	creativity = input()
	shopping = input()
	freeibies = input()

	if parkingAccessibility == "good":
	exhibit6rating += 1
	if parkingAccessibility == "bad":
	exhibit6rating += -1
	if enoughinterestingDisplays == true:
	exhibit6rating += 1
	if enoughinterestingDisplays == false:
	exhibit6rating += -1
	if relevant == true:
	exhibit6rating += 1
	if relevant == false:
	exhibit6rating += -1
	if creativity == "interesting":
	exhibit6rating += 1
	if creativity == "not interesting":
	exhibit6rating += -1
	if shopping == "present":
	exhibit6rating += 1
	if shopping == "not present":
	exhibit6rating += -1
	if freebies == true:
	exhibit6rating += 1
	if freebies == false:
	exhibit6rating += -1	

func main(){
	fmt.PrintIn(exhibit1rating)
	fmt.PrintIn(exhibit2rating)
	fmt.PrintIn(exhibit3rating)
	fmt.PrintIn(exhibit4rating)
	fmt.PrintIn(exhibit5rating)
	fmt.PrintIn(exhibit6rating)
}
