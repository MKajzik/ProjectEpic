package epic

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "github.com/onsi/gomega/gstruct"
)

func TestEpic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Epic Suite")
}

var _ = Describe("Testing package epic", func() {
	Context("Function getEpicFreeGame getting first time JSON from Epic", func() {
		It("Should return structure parsed from JSON", func() {
			actual, err := getEpicFreeGame("https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")

			//Expect(actual.GetElementStatus(0)).To(Equal("ACTIVE"))

			/*
							id := func(element interface{}) string {
				     		   return string(element.(string)[0])
				    		}

							Expect(actual).To(MatchFields(IgnoreExtras, Fields{
								"Data": MatchFields(IgnoreExtras, Fields{
									"Catalog": MatchFields(IgnoreExtras, Fields{
										"SearchStore": MatchFields(IgnoreExtras, Fields{
											"Elements" MatchElements(id, IgnoreExtras, Elements{

											})
										})
									})
								})
							}))

			*/

			Expect(actual).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
	Context("Function checkFreeGame take out free game from FreeGame struct", func() {
		It("Should return string with names of free games / game", func() {
			freegame, _ := getEpicFreeGame("https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")
			actual, num := checkFreeGame(freegame)

			Expect(actual).NotTo(BeNil())
			Expect(num).NotTo(BeNil())
		})
	})

	Context("Function prepareJSON should take the FreeGame object and convert it to JSON structure", func() {
		It("Should return JSON for requestBody, name of a game and nil error", func() {
			freegame, _ := getEpicFreeGame("https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")
			actual, name, err := prepareJSON(freegame)

			Expect(actual).NotTo(BeNil())
			Expect(name).NotTo(BeNil())
			Expect(err).To(BeNil())

		})
	})
})
